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
	"os"
	"path"
)

type osFS string

var (
	_ ManagedFS = (*osFS)(nil)
)

func NewOS(root string) (ManagedFS, error) {
	if _, err := os.Stat(path.Clean(root)); err != nil {
		return nil, err
	}
	return osFS(path.Clean(root)), nil
}

func (rfs osFS) Open(name string) (File, error) {
	return os.Open(path.Clean(path.Join(string(rfs), name)))
}

func (rfs osFS) Delete(name string) error {
	return os.Remove(path.Clean(path.Join(string(rfs), name)))
}

func (rfs osFS) Create(name string) (File, error) {
	return os.Create(path.Clean(path.Join(string(rfs), name)))
}
