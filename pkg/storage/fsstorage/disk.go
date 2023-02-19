// Copyright 2023 Sean (MovieStoreGuy) Marciniak
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fsstorage // import "github.com/MovieStoreGuy/benchmarkit/pkg/storage/fsstorage"

import (
	"bytes"
	"context"
	"errors"
	"io"
	"strings"

	sdktrace "go.opentelemetry.io/otel/trace"

	"github.com/MovieStoreGuy/benchmarkit/pkg/result"
	"github.com/MovieStoreGuy/benchmarkit/pkg/storage"
	"github.com/MovieStoreGuy/benchmarkit/pkg/storage/fsstorage/internal/filesystem"
)

type storer struct {
	fs      filesystem.ManagedFS
	encoder result.Encoder
}

type FileReference struct {
	Name string
}

var (
	_ storage.Storage[FileReference] = (*storer)(nil)
)

func WithRootDir(root string) storage.Option[FileReference] {
	return storage.OptionFunc[FileReference](func(s storage.Storage[FileReference]) error {
		store, ok := s.(*storer)
		if !ok {
			return errors.New("unable to cast to storer")
		}
		fs, err := filesystem.NewOS(root)
		if err != nil {
			return err
		}
		store.fs = fs
		return nil
	})
}

func WithFSLock() storage.Option[FileReference] {
	return storage.OptionFunc[FileReference](func(s storage.Storage[FileReference]) error {
		store, ok := s.(*storer)
		if !ok {
			return errors.New("unable to cast to storer")
		}
		store.fs = filesystem.ApplyFSLock(store.fs)
		return nil
	})
}

func WithInMemoryFS(root string) storage.Option[FileReference] {
	return storage.OptionFunc[FileReference](func(s storage.Storage[FileReference]) error {
		store, ok := s.(*storer)
		if !ok {
			return errors.New("unable to cast to storer")
		}
		fs, err := filesystem.NewMemory(root)
		if err != nil {
			return err
		}
		store.fs = fs
		return nil
	})
}

func WithInstrumentedFS(tp sdktrace.TracerProvider) storage.Option[FileReference] {
	return storage.OptionFunc[FileReference](func(s storage.Storage[FileReference]) error {
		store, ok := s.(*storer)
		if !ok {
			return errors.New("unable to cast to storer")
		}
		store.fs = filesystem.ApplyInstrumented(
			store.fs,
			filesystem.WithTracerProvider(tp),
		)
		return nil
	})
}

func New(opts ...storage.Option[FileReference]) (storage.Storage[FileReference], error) {
	fs, err := filesystem.NewOS(".")
	if err != nil {
		return nil, err
	}
	s := &storer{
		fs:      fs,
		encoder: result.NewJSONEncoder(),
	}
	for _, opt := range opts {
		if err := opt.Apply(s); err != nil {
			return nil, err
		}
	}
	return s, err
}

func hashName(vals ...string) string {
	return strings.Join(vals, "_")
}

func (s *storer) Init(_ context.Context) error {
	return nil
}

func (s *storer) Create(ctx context.Context, benckmarks ...result.Benchmark) ([]storage.Descriptor[FileReference], error) {
	var (
		descriptors = make([]storage.Descriptor[FileReference], 0, len(benckmarks))
	)
	for _, b := range benckmarks {
		content, err := s.encoder.Encode(b)
		if err != nil {
			return nil, err
		}
		ref := FileReference{
			Name: hashName(
				b.Project().Name(),
				b.Project().CommitID(),
				b.Project().Tag(),
			),
		}
		f, err := s.fs.Create(ctx, ref.Name)
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(f, bytes.NewReader(content)); err != nil {
			return nil, err
		}
		descriptors = append(descriptors, storage.NewDescriptor(ref, b))
	}
	return descriptors, nil
}

func (s *storer) Delete(ctx context.Context, references ...storage.Descriptor[FileReference]) error {
	for _, ref := range references {
		err := s.fs.Delete(ctx, hashName(
			ref.Benchmark().Project().Name(),
			ref.Benchmark().Project().CommitID(),
			ref.Benchmark().Project().Tag(),
		))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *storer) Update(ctx context.Context, references ...storage.Descriptor[FileReference]) error {
	for _, ref := range references {
		name := hashName(
			ref.Benchmark().Project().Name(),
			ref.Benchmark().Project().CommitID(),
			ref.Benchmark().Project().Tag(),
		)
		f, err := s.fs.Open(ctx, name)
		if err != nil {
			return err
		}
		d, err := s.encoder.Encode(ref.Benchmark())
		if err != nil {
			return err
		}
		if _, err = io.Copy(f, bytes.NewReader(d)); err != nil {
			return err
		}
	}
	return nil
}
