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

package storage // import "github.com/MovieStoreGuy/benchmarkit/pkg/storage"

import (
	"context"

	"github.com/MovieStoreGuy/benchmarkit/pkg/result"
)

type Storage[Reference any] interface {
	// Init creates any required entries or structures for the storage type
	Init(ctx context.Context) error

	// Create takes the benchmark results and stores it to the storage
	// solution returning the storage descriptor or an error.
	Create(ctx context.Context, bench ...result.Benchmark) ([]Descriptor[Reference], error)

	// Update modifies the reference with exact value provided.
	Update(ctx context.Context, references ...Descriptor[Reference]) error

	// Delete removes the references from the storage
	Delete(ctx context.Context, references ...Descriptor[Reference]) error
}
