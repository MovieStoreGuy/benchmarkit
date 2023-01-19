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
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/MovieStoreGuy/benchmarkit/pkg/result/internal/encoded"
)

func NewJSONEncoder() Encoder {
	return EncoderFunc(func(bench Benchmark) ([]byte, error) {
		return protojson.Marshal(bench.original())
	})
}

func NewProtoEncoder() Encoder {
	return EncoderFunc(func(bench Benchmark) ([]byte, error) {
		return proto.Marshal(bench.original())
	})
}

func NewJSONDecoder() Decoder {
	return DecoderFunc(func(data []byte) (Benchmark, error) {
		var orig encoded.Benchmark
		if err := protojson.Unmarshal(data, &orig); err != nil {
			return Benchmark{}, err
		}
		return Benchmark{orig: &orig}, nil
	})
}

func NewProtoDecoder() Decoder {
	return DecoderFunc(func(data []byte) (Benchmark, error) {
		var orig encoded.Benchmark
		if err := proto.Unmarshal(data, &orig); err != nil {
			return Benchmark{}, err
		}
		return Benchmark{orig: &orig}, nil
	})
}
