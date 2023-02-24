package stream

import (
	"context"
)

type Result[T any] struct {
	Val T
	Err error
}

type Entry[T any] struct {
	Ctx context.Context
	ch  chan<- Result[T]
}

func (e *Entry[T]) Send(val T) {
	e.ch <- Result[T]{Val: val}
}

func (e *Entry[T]) Panic(err error) {
	e.ch <- Result[T]{Err: err}
	panic(err)
}

func (e *Entry[T]) Error(err error) {
	e.ch <- Result[T]{Err: err}
}

func (e *Entry[T]) Close() {
	close(e.ch)
}

// WithContext 替换 ctx
func (e *Entry[T]) WithContext(ctx context.Context) *Entry[T] {
	if ctx == nil {
		panic("nil context")
	}
	r := &Entry[T]{
		Ctx: ctx,
		ch:  e.ch,
	}
	return r
}

func AsyncReflow[T any](ctx context.Context, fun func(*Entry[T])) <-chan Result[T] {
	ch := make(chan Result[T], 1)
	entry := &Entry[T]{
		Ctx: ctx,
		ch:  ch,
	}
	go func() {
		defer func() {
			_ = recover()
			close(ch)
		}()
		fun(entry)
	}()
	return ch
}
