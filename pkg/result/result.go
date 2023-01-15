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

import "github.com/MovieStoreGuy/benchmarkit/pkg/result/internal/proto"

type Result struct {
	orig *proto.Result
}

func NewResult() Result {
	return Result{orig: &proto.Result{}}
}

func (r *Result) SetPlatform(platform string) {
	r.orig.Platform = platform
}

func (r *Result) SetArch(arch string) {
	r.orig.Arch = arch
}

func (r *Result) SetName(name string) {
	r.orig.Name = name
}

func (r *Result) SetExecutions(n uint64) {
	r.orig.Executions = n
}

func (r *Result) SetValue(value float64) {
	r.orig.Value = value
}

func (r *Result) SetMetric(metric string) {
	r.orig.Metric = metric
}

func (r *Result) Platform() string {
	return r.orig.GetPlatform()
}

func (r *Result) Arch() string {
	return r.orig.GetArch()
}

func (r *Result) Name() string {
	return r.orig.GetName()
}

func (r *Result) Executions() uint64 {
	return r.orig.GetExecutions()
}

func (r *Result) Value() float64 {
	return r.orig.GetValue()
}

func (r *Result) Metric() string {
	return r.orig.GetMetric()
}

func (r *Result) original() *proto.Result {
	return r.orig
}
