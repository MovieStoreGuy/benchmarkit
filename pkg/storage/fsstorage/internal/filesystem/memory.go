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

package filesystem // import "github.com/MovieStoreGuy/benchmarkit/pkg/storage/fsstorage/internal/filesystem"

import (
	"bytes"
	"context"
	"io"
	"io/fs"
	"path"
)

type memFS struct {
	root string
	data map[string][]byte
}

type memFile struct {
	io.ReadWriter
}

var (
	_ ManagedFS = (*memFS)(nil)

	_ File = (*memFile)(nil)
)

func NewMemory(root string) (ManagedFS, error) {
	return &memFS{
		root: path.Clean(root),
		data: make(map[string][]byte),
	}, nil
}

func (mem *memFS) Open(ctx context.Context, name string) (File, error) {
	fp := path.Clean(path.Join(mem.root, name))
	if content, ok := mem.data[fp]; ok {
		return &memFile{
			ReadWriter: bytes.NewBuffer(content),
		}, nil
	}
	return nil, &fs.PathError{
		Op:   "open",
		Path: fp,
		Err:  fs.ErrNotExist,
	}
}

func (mem *memFS) Delete(ctx context.Context, name string) error {
	fp := path.Clean(path.Join(mem.root, name))
	if _, ok := mem.data[fp]; ok {
		delete(mem.data, fp)
		return nil
	}
	return &fs.PathError{
		Op:   "remove",
		Path: fp,
		Err:  fs.ErrNotExist,
	}
}

func (mem *memFS) Create(ctx context.Context, name string) (File, error) {
	fp := path.Clean(path.Join(mem.root, name))
	if _, ok := mem.data[fp]; ok {
		return nil, &fs.PathError{
			Op:   "open",
			Path: fp,
			Err:  fs.ErrExist,
		}
	}
	content := make([]byte, 0, 100)
	mem.data[fp] = content
	return &memFile{
		ReadWriter: bytes.NewBuffer(content[:]),
	}, nil
}

func (f *memFile) Close() error { return nil }
