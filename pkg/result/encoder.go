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
	"bytes"
	"errors"
	"fmt"
	"io"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	ErrUndefinedNamedDecoder = errors.New("undefined named decoder")
)

// Encoder is used to write the result data
// out as a specific format.
type Encoder interface {
	Encode(bench Benchmark) error
}

// NewEncoderFunc allows to create an inline encoder
// to help unify how encoders are implemented
type NewEncoderFunc func(out io.Writer) (Encoder, error)

// EncoderFunc allows for a function definition to replace
// how the Benchmark is being consumed
type EncoderFunc func(bench Benchmark) error

// EncoderFactory allows for a different types of
// export formats to be defined and used.
type EncoderFactory interface {
	NamedEncoder(name string, out io.Writer) (Encoder, error)
}

func (fn EncoderFunc) Encode(bench Benchmark) error {
	return fn(bench)
}

type factoryMap map[string]NewEncoderFunc

var (
	_ Encoder        = (EncoderFunc)(nil)
	_ EncoderFactory = (*factoryMap)(nil)
)

func NewEncoderFactory() EncoderFactory {
	return factoryMap{
		"proto": NewEncoderFunc(func(out io.Writer) (Encoder, error) {
			return EncoderFunc(func(bench Benchmark) error {
				buf, err := proto.Marshal(bench.original())
				if err != nil {
					return err
				}
				_, err = io.Copy(out, bytes.NewBuffer(buf))
				return err
			}), nil
		}),
		"json": NewEncoderFunc(func(out io.Writer) (Encoder, error) {
			return EncoderFunc(func(bench Benchmark) error {
				buf, err := protojson.Marshal(bench.original())
				if err != nil {
					return err
				}
				_, err = io.Copy(out, bytes.NewReader(buf))
				return err
			}), nil
		}),
	}
}

func (fm factoryMap) NamedEncoder(name string, out io.Writer) (Encoder, error) {
	fn, ok := fm[name]
	if !ok {
		return nil, fmt.Errorf("encoder %s : %w", name, ErrUndefinedNamedDecoder)
	}
	return fn(out)
}
