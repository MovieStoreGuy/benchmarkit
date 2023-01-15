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

package result // import "github.com/MovieStoreGuy/benchmarkit/pkg/result"

import "github.com/MovieStoreGuy/benchmarkit/pkg/result/internal/encoded"

type Benchmark struct {
	orig *encoded.Benchmark
}

func NewBenchmark() Benchmark {
	return Benchmark{
		orig: &encoded.Benchmark{
			Project: new(encoded.Project),
			Results: []*encoded.Result{},
		},
	}
}

func (b *Benchmark) SetProject(p *Project) {
	b.orig.Project = p.original()
}

func (b *Benchmark) SetResults(results *ResultSlice) {
	b.orig.Results = results.original()
}

func (b *Benchmark) Project() *Project {
	return &Project{orig: b.orig.Project}
}

func (b *Benchmark) Results() *ResultSlice {
	return &ResultSlice{orig: &b.orig.Results}
}

func (b *Benchmark) original() *encoded.Benchmark {
	return b.orig
}
