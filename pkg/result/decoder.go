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
	"errors"
	"fmt"
	"io"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/MovieStoreGuy/benchmarkit/pkg/result/internal/encoded"
)

var (
	ErrUnknownDecodeFormat = errors.New("unknown decode format")
)

// Decoder is used to read the wire format
// to convert into a Benchmark type
type Decoder interface {
	Decode() (Benchmark, error)
}

// NewDecoderFunc unifies and inlines encoding type Unmarshalers
type NewDecoderFunc func(in io.Reader) (Decoder, error)

// DecoderFunc will process the data into the wrapped type Benchmark
type DecodeFunc func() (Benchmark, error)

func (fn DecodeFunc) Decode() (Benchmark, error) {
	return fn()
}

// DecodeFactory allows for a simple means of associating
// the encoding format to the Decoder type.
type DecodeFactory interface {
	NewDecoder(format Format, in io.Reader) (Decoder, error)
}

type decoderMap map[Format]NewDecoderFunc

var (
	_ Decoder       = (DecodeFunc)(nil)
	_ DecodeFactory = (*decoderMap)(nil)
)

func NewDecoderFactory() DecodeFactory {
	return decoderMap{
		FormatProtobuf: NewDecoderFunc(func(in io.Reader) (Decoder, error) {
			return DecodeFunc(func() (Benchmark, error) {
				buf, err := io.ReadAll(in)
				if err != nil {
					return NewBenchmark(), err
				}
				var b encoded.Benchmark
				if err = proto.Unmarshal(buf, &b); err != nil {
					return NewBenchmark(), err
				}
				return Benchmark{orig: &b}, nil
			}), nil
		}),
		FormatJSON: NewDecoderFunc(func(in io.Reader) (Decoder, error) {
			return DecodeFunc(func() (Benchmark, error) {
				buf, err := io.ReadAll(in)
				if err != nil {
					return NewBenchmark(), err
				}
				var b encoded.Benchmark
				if err = protojson.Unmarshal(buf, &b); err != nil {
					return NewBenchmark(), err
				}
				return Benchmark{orig: &b}, nil
			}), nil
		}),
	}
}

func (fm decoderMap) NewDecoder(format Format, in io.Reader) (Decoder, error) {
	fn, ok := fm[format]
	if !ok {
		return nil, fmt.Errorf("format %s: %w", format, ErrUnknownDecodeFormat)
	}
	return fn(in)
}
