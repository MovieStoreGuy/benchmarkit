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
	"io"
	"io/fs"
	"path"
	"time"
)

type memFS struct {
	root string
	data map[string][]byte
}

type memFile struct {
	io.ReadWriter
	info memInfo
}

type memInfo struct {
	name  string
	size  int64
	mtime time.Time
}

var (
	_ ManagedFS = (*memFS)(nil)

	_ File     = (*memFile)(nil)
	_ FileInfo = (*memInfo)(nil)
)

func NewMemory(root string) (ManagedFS, error) {
	return &memFS{
		root: path.Clean(root),
		data: make(map[string][]byte),
	}, nil
}

func (mem *memFS) Open(name string) (File, error) {
	fp := path.Clean(path.Join(mem.root, name))
	if content, ok := mem.data[fp]; ok {
		return &memFile{
			ReadWriter: bytes.NewBuffer(content),
			info: memInfo{
				name:  path.Base(fp),
				size:  int64(len(content)),
				mtime: time.Now(),
			},
		}, nil
	}
	return nil, &fs.PathError{
		Op:   "open",
		Path: fp,
		Err:  fs.ErrNotExist,
	}
}

func (mem *memFS) Delete(name string) error {
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

func (mem *memFS) Create(name string) (File, error) {
	fp := path.Clean(path.Join(mem.root, name))
	if content, ok := mem.data[fp]; !ok {
		return &memFile{
			ReadWriter: bytes.NewBuffer(content[:]),
			info: memInfo{
				name:  path.Base(fp),
				size:  -1,
				mtime: time.Now(),
			},
		}, nil
	}
	return nil, &fs.PathError{
		Op:   "open",
		Path: fp,
		Err:  fs.ErrExist,
	}
}

func (mf *memFile) Stat() (FileInfo, error) { return &mf.info, nil }
func (mf *memFile) Close() error            { return nil }

func (ms *memInfo) Name() string       { return ms.name }
func (ms *memInfo) Size() int64        { return ms.size }
func (ms *memInfo) Mode() FileMode     { return FileMode(0) }
func (ms *memInfo) ModTime() time.Time { return ms.mtime }
func (ms *memInfo) IsDir() bool        { return ms.Mode().IsDir() }
func (ms *memInfo) Sys() any           { return nil }
