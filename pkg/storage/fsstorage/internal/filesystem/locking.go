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
	"errors"
	"sync"
)

type lockingFS struct {
	locks   sync.Map
	wrapped ManagedFS
}

var (
	_ ManagedFS = (*lockingFS)(nil)
)

func ApplyFSLock(fs ManagedFS) ManagedFS {
	return &lockingFS{wrapped: fs}
}

func (lfs *lockingFS) Open(name string) (File, error) {
	v, _ := lfs.locks.LoadOrStore(name, new(sync.Mutex))
	l, ok := v.(*sync.Mutex)
	if !ok {
		return nil, errors.New("can not type cast lock")
	}
	l.Lock()
	defer l.Unlock()
	return lfs.wrapped.Open(name)
}

func (lfs *lockingFS) Delete(name string) error {
	v, _ := lfs.locks.LoadOrStore(name, new(sync.Mutex))
	l, ok := v.(*sync.Mutex)
	if !ok {
		return errors.New("can not type cast lock")
	}
	l.Lock()
	defer l.Unlock()
	return lfs.wrapped.Delete(name)
}

func (lfs *lockingFS) Create(name string) (File, error) {
	v, _ := lfs.locks.LoadOrStore(name, new(sync.Mutex))
	l, ok := v.(*sync.Mutex)
	if !ok {
		return nil, errors.New("can not type cast lock")
	}
	l.Lock()
	defer l.Unlock()
	return lfs.wrapped.Create(name)
}
