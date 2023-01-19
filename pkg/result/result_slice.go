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

import (
	"go.uber.org/multierr"

	"github.com/MovieStoreGuy/benchmarkit/pkg/result/internal/encoded"
)

type ResultSlice struct {
	orig *[]*encoded.Result
}

func NewResultSlice() ResultSlice {
	return ResultSlice{orig: new([]*encoded.Result)}
}

func (rs ResultSlice) At(i int) *Result {
	return &Result{orig: (*rs.orig)[i]}
}

func (rs ResultSlice) Append(results ...*Result) {
	rs.EnsureCapacity(rs.Len() + len(results))
	for i := 0; i < len(results); i++ {
		(*rs.orig)[i] = results[i].original()
	}
}

func (rs ResultSlice) AppendEmpty() *Result {
	(*rs.orig) = append((*rs.orig), &encoded.Result{})
	return &Result{orig: (*rs.orig)[rs.Len()-1]}
}

func (rs ResultSlice) EnsureCapacity(newcap int) {
	if current := cap(*rs.orig); current >= newcap {
		return
	}
	neworig := make([]*encoded.Result, len(*rs.orig), newcap)
	copy(neworig, (*rs.orig))
	rs.orig = &neworig
}

func (rs ResultSlice) Len() int {
	return len(*rs.orig)
}

func (rs ResultSlice) MoveAndAppend(slice *ResultSlice) {
	(*rs.orig) = append((*rs.orig), slice.original()...)
	slice.orig = new([]*encoded.Result)
}

func (rs ResultSlice) Range(fn func(r *Result) error) (errs error) {
	for i, len := 0, rs.Len(); i < len; i++ {
		errs = multierr.Append(errs, fn(rs.At(i)))
	}
	return errs
}

func (rs ResultSlice) SetAt(i int, r *Result) {
	(*rs.orig)[i] = r.original()
}

func (rs ResultSlice) original() []*encoded.Result {
	return *rs.orig
}
