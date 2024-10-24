package fw

import (
	"context"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type BootTask[T any] struct {
	ctx  context.Context
	cfg  T
	name string
	fn   func(ctx context.Context, cfg T) error
}

func NewBootTask[T any](ctx context.Context, cfg T, name string, fn func(ctx context.Context, cfg T) error) *BootTask[T] {
	return &BootTask[T]{
		ctx:  logx.ContextWithFields(ctx, logc.Field("boot_task", name)),
		cfg:  cfg,
		name: name,
		fn:   fn,
	}
}

func (t *BootTask[T]) Start() error {
	logc.Info(t.ctx, "Starting")
	if err := t.fn(t.ctx, t.cfg); err != nil {
		logc.Errorf(t.ctx, "Failed: %+v", err)
		return err
	} else {
		logc.Infof(t.ctx, "Done")
		return nil
	}
}

func (t *BootTask[T]) StartAsync() (done chan error) {
	done = make(chan error)
	go func() {
		defer close(done)
		if err := t.fn(t.ctx, t.cfg); err != nil {
			logc.Errorf(t.ctx, "Failed: %+v", err)
			done <- err
		} else {
			logc.Infof(t.ctx, "Done")
			done <- nil
		}
	}()
	return done
}
