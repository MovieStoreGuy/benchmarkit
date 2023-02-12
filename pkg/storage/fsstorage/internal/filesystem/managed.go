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

import "io/fs"

// ManagedFS is an interface that allows CRUD interactions
// with a abstracted filesystem.
type ManagedFS interface {
	Create(name string) (File, error)
	Delete(name string) error
	Open(name string) (File, error)
}

type (
	FS       = fs.FS
	FileInfo = fs.FileInfo
	FileMode = fs.FileMode

	File interface {
		fs.File

		Write(p []byte) (int, error)
	}
)