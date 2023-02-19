// Copyright 2023 Sean (MovieStoreGuy) Marciniak
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filesystem // import "github.com/MovieStoreGuy/benchmarkit/pkg/storage/fsstorage/internal/filesystem"

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/trace"
)

const (
	InstrumentationName    = "filesystem"
	InstrumentationVersion = "1.0.0"

	operationCreate = `Create file`
	operationOpen   = `Open file`
	operationDelete = `Delete file`
)

type instrumentedFS struct {
	tracer  sdktrace.Tracer
	wrapped ManagedFS
}

type InstrumentationOptions struct {
	TracerProvider sdktrace.TracerProvider
}

// InstrumentationOptionFunc allows for providing dynamic options
// instrumentation as project becomes more stable
type InstrumentationOptionFunc func(io *InstrumentationOptions)

type operationFunc func(ctx context.Context, name string) (File, error)

var (
	_ ManagedFS = (*instrumentedFS)(nil)
)

func WithTracerProvider(tp sdktrace.TracerProvider) InstrumentationOptionFunc {
	return func(io *InstrumentationOptions) {
		io.TracerProvider = tp
	}
}

func ApplyInstrumented(fs ManagedFS, opts ...InstrumentationOptionFunc) ManagedFS {
	o := &InstrumentationOptions{
		TracerProvider: sdktrace.NewNoopTracerProvider(),
	}
	for _, opt := range opts {
		opt(o)
	}
	return &instrumentedFS{
		wrapped: fs,
		tracer: o.TracerProvider.Tracer(
			InstrumentationName,
			sdktrace.WithInstrumentationVersion(InstrumentationVersion),
		),
	}
}

func (in *instrumentedFS) operation(ctx context.Context, operation string, name string, op operationFunc) (File, error) {
	ctx, span := in.tracer.Start(
		ctx,
		operation,
		sdktrace.WithSpanKind(sdktrace.SpanKindInternal),
		sdktrace.WithAttributes(attribute.String("name", name)),
	)
	f, err := op(ctx, name)
	if err != nil {
		span.RecordError(err)
	}
	span.End()
	return f, err
}

func (in *instrumentedFS) Create(ctx context.Context, name string) (File, error) {
	return in.operation(ctx, operationCreate, name, in.wrapped.Create)
}

func (in *instrumentedFS) Open(ctx context.Context, name string) (File, error) {
	return in.operation(ctx, operationOpen, name, in.wrapped.Open)
}

func (in *instrumentedFS) Delete(ctx context.Context, name string) error {
	_, err := in.operation(
		ctx,
		operationDelete,
		name,
		func(ctx context.Context, name string) (File, error) {
			return NopFile{}, in.wrapped.Delete(ctx, name)
		},
	)
	return err
}
