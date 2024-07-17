package depthlogger

import (
	"context"
	"log/slog"
	"runtime"
	"time"
)

type DepthLogger struct {
	*slog.Logger
}

func NewDepthLogger(h slog.Handler) *DepthLogger {
	return &DepthLogger{slog.New(h)}
}

func (o *DepthLogger) DebugDepth(msg string, skip int, args ...any) {
	o.LogDepth(context.Background(), slog.LevelDebug, skip+1, msg, args...)
}

func (o *DepthLogger) DebugDepthContext(ctx context.Context, skip int, msg string, args ...any) {
	o.LogDepth(ctx, slog.LevelDebug, skip+1, msg, args...)
}

func (o *DepthLogger) InfoDepthDepth(skip int, msg string, args ...any) {
	o.LogDepth(context.Background(), slog.LevelInfo, skip+1, msg, args...)
}

func (o *DepthLogger) InfoDepthContext(ctx context.Context, skip int, msg string, args ...any) {
	o.LogDepth(ctx, slog.LevelInfo, skip+1, msg, args...)
}

func (o *DepthLogger) WarnDepth(skip int, msg string, args ...any) {
	o.LogDepth(context.Background(), slog.LevelWarn, skip+1, msg, args...)
}

func (o *DepthLogger) WarnDepthContext(ctx context.Context, skip int, msg string, args ...any) {
	o.LogDepth(ctx, slog.LevelWarn, skip+1, msg, args...)
}

func (o *DepthLogger) ErrorDepth(skip int, msg string, args ...any) {
	o.LogDepth(context.Background(), slog.LevelError, skip+1, msg, args...)
}

func (o *DepthLogger) ErrorDepthContext(ctx context.Context, skip int, msg string, args ...any) {
	o.LogDepth(ctx, slog.LevelError, skip+1, msg, args...)
}

func (o *DepthLogger) LogDepth(ctx context.Context, level slog.Level, skip int, msg string, args ...any) {
	if !o.Enabled(ctx, level) {
		return
	}
	var pc uintptr
	{
		var pcs [1]uintptr
		// skip [runtime.Callers, this function]
		runtime.Callers(skip+2, pcs[:])
		pc = pcs[0]
	}
	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.Add(args...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = o.Handler().Handle(ctx, r)
}

func (o *DepthLogger) LogAttrsDepth(ctx context.Context, level slog.Level, skip int, msg string, attrs ...slog.Attr) {
	if !o.Enabled(ctx, level) {
		return
	}
	var pc uintptr
	{
		var pcs [1]uintptr
		// skip [runtime.Callers, this function]
		runtime.Callers(skip+2, pcs[:])
		pc = pcs[0]
	}
	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.AddAttrs(attrs...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = o.Handler().Handle(ctx, r)
}

func (o *DepthLogger) LogPc(ctx context.Context, level slog.Level, pc uintptr, msg string, args ...any) {
	if !o.Enabled(ctx, level) {
		return
	}
	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.Add(args...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = o.Handler().Handle(ctx, r)
}
