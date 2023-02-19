// Copyright 2023 Sean (MovieStoreGuy) Marciniak
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filesystem // import "github.com/MovieStoreGuy/benchmarkit/pkg/storage/fsstorage/internal/filesystem"

import (
	"context"
	"io/fs"
)

type NopFS struct{}

type NopFile struct{}

var (
	_ File      = (*NopFile)(nil)
	_ ManagedFS = (*NopFS)(nil)
)

func (NopFS) Create(_ context.Context, _ string) (File, error) {
	return NopFile{}, nil
}

func (NopFS) Open(_ context.Context, _ string) (File, error) {
	return nil, fs.ErrNotExist
}

func (NopFS) Delete(_ context.Context, _ string) error {
	return nil
}

func (NopFile) Read(in []byte) (int, error) {
	return len(in), nil
}

func (NopFile) Write(in []byte) (int, error) {
	return len(in), nil
}

func (NopFile) Close() error {
	return nil
}
