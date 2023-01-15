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

// Package Encoded is the raw wire format of the project
// and intentionally not directly exposed to simplify interactions.
// Any protobuf definitions are not publically acessible that the provided
// encoders are used.

//go:generate protoc --go_opt=paths=source_relative --go_out=. message.proto
package encoded // import "github.com/MovieStoreGuy/benchmarkit/pkg/result/internal/encoded"
