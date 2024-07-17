package depthlogger

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"time"
)

type PCGetterFunc func() uintptr

type PCGetter interface {
	GetPC() uintptr
}

type DepthLogger struct {
	*slog.Logger
}

func NewDepthLogger(h slog.Handler) *DepthLogger {
	return &DepthLogger{slog.New(h)}
}

func (o *DepthLogger) Debugf(msg string, args ...any) {
	o.LogDepthf(context.Background(), slog.LevelDebug, 1, msg, args...)
}

func (o *DepthLogger) DebugDepth(msg string, skip int, args ...any) {
	o.LogDepth(context.Background(), slog.LevelDebug, skip+1, msg, args...)
}

func (o *DepthLogger) DebugDepthf(msg string, skip int, args ...any) {
	o.LogDepthf(context.Background(), slog.LevelDebug, skip+1, msg, args...)
}

func (o *DepthLogger) DebugContextf(ctx context.Context, msg string, args ...any) {
	o.LogDepthf(ctx, slog.LevelDebug, 1, msg, args...)
}

func (o *DepthLogger) DebugContextDepth(ctx context.Context, skip int, msg string, args ...any) {
	o.LogDepth(ctx, slog.LevelDebug, skip+1, msg, args...)
}

func (o *DepthLogger) DebugContextDepthf(ctx context.Context, skip int, msg string, args ...any) {
	o.LogDepthf(ctx, slog.LevelDebug, skip+1, msg, args...)
}

func (o *DepthLogger) Infof(msg string, args ...any) {
	o.LogDepthf(context.Background(), slog.LevelInfo, 1, msg, args...)
}

func (o *DepthLogger) InfoDepthDepth(skip int, msg string, args ...any) {
	o.LogDepth(context.Background(), slog.LevelInfo, skip+1, msg, args...)
}

func (o *DepthLogger) InfoDepthDepthf(skip int, msg string, args ...any) {
	o.LogDepthf(context.Background(), slog.LevelInfo, skip+1, msg, args...)
}

func (o *DepthLogger) InfoContextf(ctx context.Context, msg string, args ...any) {
	o.LogDepthf(ctx, slog.LevelInfo, 1, msg, args...)
}

func (o *DepthLogger) InfoContextDepth(ctx context.Context, skip int, msg string, args ...any) {
	o.LogDepth(ctx, slog.LevelInfo, skip+1, msg, args...)
}

func (o *DepthLogger) InfoContextDepthf(ctx context.Context, skip int, msg string, args ...any) {
	o.LogDepthf(ctx, slog.LevelInfo, skip+1, msg, args...)
}

func (o *DepthLogger) Warnf(msg string, args ...any) {
	o.LogDepthf(context.Background(), slog.LevelWarn, 1, msg, args...)
}

func (o *DepthLogger) WarnDepth(skip int, msg string, args ...any) {
	o.LogDepth(context.Background(), slog.LevelWarn, skip+1, msg, args...)
}

func (o *DepthLogger) WarnDepthf(skip int, msg string, args ...any) {
	o.LogDepthf(context.Background(), slog.LevelWarn, skip+1, msg, args...)
}

func (o *DepthLogger) WarnContextf(ctx context.Context, msg string, args ...any) {
	o.LogDepthf(ctx, slog.LevelWarn, 1, msg, args...)
}

func (o *DepthLogger) WarnContextDepth(ctx context.Context, skip int, msg string, args ...any) {
	o.LogDepth(ctx, slog.LevelWarn, skip+1, msg, args...)
}

func (o *DepthLogger) WarnContextDepthf(ctx context.Context, skip int, msg string, args ...any) {
	o.LogDepthf(ctx, slog.LevelWarn, skip+1, msg, args...)
}

func (o *DepthLogger) Errorf(msg string, args ...any) {
	o.LogDepthf(context.Background(), slog.LevelError, 1, msg, args...)
}

func (o *DepthLogger) ErrorDepth(skip int, msg string, args ...any) {
	o.LogDepth(context.Background(), slog.LevelError, skip+1, msg, args...)
}

func (o *DepthLogger) ErrorDepthf(skip int, msg string, args ...any) {
	o.LogDepthf(context.Background(), slog.LevelError, skip+1, msg, args...)
}

func (o *DepthLogger) ErrorContextf(ctx context.Context, msg string, args ...any) {
	o.LogDepthf(ctx, slog.LevelError, 1, msg, args...)
}

func (o *DepthLogger) ErrorContextDepth(ctx context.Context, skip int, msg string, args ...any) {
	o.LogDepth(ctx, slog.LevelError, skip+1, msg, args...)
}

func (o *DepthLogger) ErrorContextDepthf(ctx context.Context, skip int, msg string, args ...any) {
	o.LogDepthf(ctx, slog.LevelError, skip+1, msg, args...)
}

func (o *DepthLogger) Logf(ctx context.Context, level slog.Level, msg string, args ...any) {
	o.LogDepthf(ctx, level, 1, msg, args...)
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

func (o *DepthLogger) LogDepthAttrs(ctx context.Context, level slog.Level, skip int, msg string, attrs ...slog.Attr) {
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

func (o *DepthLogger) LogDepthf(ctx context.Context, level slog.Level, skip int, msg string, args ...any) {
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
	r := slog.NewRecord(time.Now(), level, fmt.Sprintf(msg, args...), pc)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = o.Handler().Handle(ctx, r)
}

func (o *DepthLogger) LogPCFunc(ctx context.Context, level slog.Level, pcf PCGetterFunc, msg string, args ...any) {
	if !o.Enabled(ctx, level) {
		return
	}
	r := slog.NewRecord(time.Now(), level, msg, pcf())
	r.Add(args...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = o.Handler().Handle(ctx, r)
}

func (o *DepthLogger) LogPCFuncAttrs(ctx context.Context, level slog.Level, pcf PCGetterFunc, msg string, attrs ...slog.Attr) {
	if !o.Enabled(ctx, level) {
		return
	}
	r := slog.NewRecord(time.Now(), level, msg, pcf())
	r.AddAttrs(attrs...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = o.Handler().Handle(ctx, r)
}

func (o *DepthLogger) LogPCFuncf(ctx context.Context, level slog.Level, pcf PCGetterFunc, msg string, args ...any) {
	if !o.Enabled(ctx, level) {
		return
	}
	r := slog.NewRecord(time.Now(), level, fmt.Sprintf(msg, args...), pcf())
	if ctx == nil {
		ctx = context.Background()
	}
	_ = o.Handler().Handle(ctx, r)
}

func (o *DepthLogger) LogPC(ctx context.Context, level slog.Level, pcGetter PCGetter, msg string, args ...any) {
	if !o.Enabled(ctx, level) {
		return
	}
	r := slog.NewRecord(time.Now(), level, msg, pcGetter.GetPC())
	r.Add(args...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = o.Handler().Handle(ctx, r)
}

func (o *DepthLogger) LogPCAttrs(ctx context.Context, level slog.Level, pcGetter PCGetter, msg string, attrs ...slog.Attr) {
	if !o.Enabled(ctx, level) {
		return
	}
	r := slog.NewRecord(time.Now(), level, msg, pcGetter.GetPC())
	r.AddAttrs(attrs...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = o.Handler().Handle(ctx, r)
}

func (o *DepthLogger) LogPCf(ctx context.Context, level slog.Level, pcGetter PCGetter, msg string, args ...any) {
	if !o.Enabled(ctx, level) {
		return
	}
	r := slog.NewRecord(time.Now(), level, fmt.Sprintf(msg, args...), pcGetter.GetPC())
	if ctx == nil {
		ctx = context.Background()
	}
	_ = o.Handler().Handle(ctx, r)
}
