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

package filesystem

import (
	"context"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FileSystemLifecycle(t *testing.T, factory func() ManagedFS) {
	t.Run("Open Files", func(t *testing.T) {
		const filename = `file-name-test`

		fsys := factory()

		f, err := fsys.Open(context.Background(), filename)
		assert.Nil(t, f, "Must be nil when opening file")
		assert.ErrorIs(t, err, fs.ErrNotExist, "Must error when opening file")

		f, err = fsys.Create(context.Background(), filename)
		assert.NoError(t, err, "Must not error when creating filename")
		assert.NotNil(t, f, "Must have a valid file")

		if _, ok := f.(NopFile); f != nil && !ok {
			f, err = fsys.Open(context.Background(), filename)
			assert.NoError(t, err, "Must not error when creating filename")
			assert.NotNil(t, f, "Must have a valid file")
		}
	})

	t.Run("Delete Files", func(t *testing.T) {
		const filename = `file-name-test`

		fsys := factory()

		f, err := fsys.Create(context.Background(), filename)
		assert.NoError(t, err, "Must not error when creating filename")
		assert.NotNil(t, f, "Must have a valid file")

		err = fsys.Delete(context.Background(), filename)
		assert.NoError(t, err, "Must not error when creating filename")
	})
}
