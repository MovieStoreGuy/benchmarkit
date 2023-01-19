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

package encode // import "github.com/MovieStoreGuy/benchmarkit/pkg/encode"

import (
	"errors"
	"fmt"

	"github.com/MovieStoreGuy/benchmarkit/pkg/result"
)

var (
	ErrFormatNotDefined = errors.New("format not defined")
	ErrFormatClashed    = errors.New("format already defined")
)

// Factory maps named formats to result encoders.
type Factory interface {
	NewEncoder(name string) (result.Encoder, error)
}

// NewMethodFunc is used to unify to creating new encoders
type NewMethodFunc func() result.Encoder

type factory map[string]NewMethodFunc

var (
	EncoderFactory = (*factory)(nil)
)

// NewFactory creates a factory with extra encoders provided.
func NewFactory(extras map[string]NewMethodFunc) (Factory, error) {
	f := factory{
		"proto": result.NewProtoEncoder,
		"json":  result.NewJSONEncoder,
	}
	for name, newer := range extras {
		if _, ok := f[name]; ok {
			return nil, fmt.Errorf("name %s: %w", name, ErrFormatClashed)
		}
		f[name] = newer
	}
	return f, nil
}

func (f factory) NewEncoder(name string) (result.Encoder, error) {
	newer, ok := f[name]
	if !ok {
		return nil, fmt.Errorf("name %s : %w", name, ErrFormatNotDefined)
	}
	return newer(), nil
}
