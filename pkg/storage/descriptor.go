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

import "github.com/MovieStoreGuy/benchmarkit/pkg/result"

// Descriptor allows for storage solutions to define
// how to reference the returned benchmark in future.
type Descriptor[Reference any] interface {
	// Reference is the type used by the storage solution
	// to be able to reference the benchmark.
	Reference() Reference
	// Benchmark is the value provided
	Benchmark() result.Benchmark
}

type descriptor[ref any] struct {
	reference ref
	bench     result.Benchmark
}

func NewDescriptor[Reference any](ref Reference, bench result.Benchmark) Descriptor[Reference] {
	return descriptor[Reference]{
		reference: ref,
		bench:     bench,
	}
}

func (d descriptor[ref]) Reference() ref {
	return d.reference
}

func (d descriptor[ref]) Benchmark() result.Benchmark {
	return d.bench
}
