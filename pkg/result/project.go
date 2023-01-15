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
	"github.com/MovieStoreGuy/benchmarkit/pkg/result/internal/encoded"
)

type Project struct {
	orig *encoded.Project
}

func NewProject() Project {
	return Project{orig: new(encoded.Project)}
}

func (p *Project) SetName(name string) {
	p.orig.Name = name
}

func (p *Project) SetTag(tag string) {
	p.orig.Tag = tag
}

func (p *Project) SetCommitID(id string) {
	p.orig.CommitId = id
}

func (p *Project) Name() string {
	return p.orig.GetName()
}

func (p *Project) Tag() string {
	return p.orig.GetTag()
}

func (p *Project) CommitID() string {
	return p.orig.GetCommitId()
}

func (p *Project) original() *encoded.Project {
	return p.orig
}
