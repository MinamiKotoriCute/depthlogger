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

func Default() *DepthLogger {
	return &DepthLogger{Logger: slog.Default()}
}

func Debugf(msg string, args ...any) {
	Default().LogDepthf(context.Background(), slog.LevelDebug, 1, msg, args...)
}

func DebugDepth(msg string, skip int, args ...any) {
	Default().LogDepth(context.Background(), slog.LevelDebug, skip+1, msg, args...)
}

func DebugDepthf(msg string, skip int, args ...any) {
	Default().LogDepthf(context.Background(), slog.LevelDebug, skip+1, msg, args...)
}

func DebugContextf(ctx context.Context, msg string, args ...any) {
	Default().LogDepthf(ctx, slog.LevelDebug, 1, msg, args...)
}

func DebugContextDepth(ctx context.Context, skip int, msg string, args ...any) {
	Default().LogDepth(ctx, slog.LevelDebug, skip+1, msg, args...)
}

func DebugContextDepthf(ctx context.Context, skip int, msg string, args ...any) {
	Default().LogDepthf(ctx, slog.LevelDebug, skip+1, msg, args...)
}

func Infof(msg string, args ...any) {
	Default().LogDepthf(context.Background(), slog.LevelInfo, 1, msg, args...)
}

func InfoDepthDepth(skip int, msg string, args ...any) {
	Default().LogDepth(context.Background(), slog.LevelInfo, skip+1, msg, args...)
}

func InfoDepthDepthf(skip int, msg string, args ...any) {
	Default().LogDepthf(context.Background(), slog.LevelInfo, skip+1, msg, args...)
}

func InfoContextf(ctx context.Context, msg string, args ...any) {
	Default().LogDepthf(ctx, slog.LevelInfo, 1, msg, args...)
}

func InfoContextDepth(ctx context.Context, skip int, msg string, args ...any) {
	Default().LogDepth(ctx, slog.LevelInfo, skip+1, msg, args...)
}

func InfoContextDepthf(ctx context.Context, skip int, msg string, args ...any) {
	Default().LogDepthf(ctx, slog.LevelInfo, skip+1, msg, args...)
}

func Warnf(msg string, args ...any) {
	Default().LogDepthf(context.Background(), slog.LevelWarn, 1, msg, args...)
}

func WarnDepth(skip int, msg string, args ...any) {
	Default().LogDepth(context.Background(), slog.LevelWarn, skip+1, msg, args...)
}

func WarnDepthf(skip int, msg string, args ...any) {
	Default().LogDepthf(context.Background(), slog.LevelWarn, skip+1, msg, args...)
}

func WarnContextf(ctx context.Context, msg string, args ...any) {
	Default().LogDepthf(ctx, slog.LevelWarn, 1, msg, args...)
}

func WarnContextDepth(ctx context.Context, skip int, msg string, args ...any) {
	Default().LogDepth(ctx, slog.LevelWarn, skip+1, msg, args...)
}

func WarnContextDepthf(ctx context.Context, skip int, msg string, args ...any) {
	Default().LogDepthf(ctx, slog.LevelWarn, skip+1, msg, args...)
}

func Errorf(msg string, args ...any) {
	Default().LogDepthf(context.Background(), slog.LevelError, 1, msg, args...)
}

func ErrorDepth(skip int, msg string, args ...any) {
	Default().LogDepth(context.Background(), slog.LevelError, skip+1, msg, args...)
}

func ErrorDepthf(skip int, msg string, args ...any) {
	Default().LogDepthf(context.Background(), slog.LevelError, skip+1, msg, args...)
}

func ErrorContextf(ctx context.Context, msg string, args ...any) {
	Default().LogDepthf(ctx, slog.LevelError, 1, msg, args...)
}

func ErrorContextDepth(ctx context.Context, skip int, msg string, args ...any) {
	Default().LogDepth(ctx, slog.LevelError, skip+1, msg, args...)
}

func ErrorContextDepthf(ctx context.Context, skip int, msg string, args ...any) {
	Default().LogDepthf(ctx, slog.LevelError, skip+1, msg, args...)
}

func Logf(ctx context.Context, level slog.Level, msg string, args ...any) {
	Default().LogDepthf(ctx, level, 1, msg, args...)
}

func LogDepth(ctx context.Context, level slog.Level, skip int, msg string, args ...any) {
	Default().LogDepth(ctx, level, skip+1, msg, args...)
}

func LogDepthAttrs(ctx context.Context, level slog.Level, skip int, msg string, attrs ...slog.Attr) {
	Default().LogDepthAttrs(ctx, level, skip+1, msg, attrs...)
}

func LogDepthf(ctx context.Context, level slog.Level, skip int, msg string, args ...any) {
	Default().LogDepthf(ctx, level, skip+1, msg, args...)
}

func LogPCFunc(ctx context.Context, level slog.Level, pcf PCGetterFunc, msg string, args ...any) {
	Default().LogPCFunc(ctx, level, pcf, msg, args...)
}

func LogPCFuncAttrs(ctx context.Context, level slog.Level, pcf PCGetterFunc, msg string, attrs ...slog.Attr) {
	Default().LogPCFuncAttrs(ctx, level, pcf, msg, attrs...)
}

func LogPCFuncf(ctx context.Context, level slog.Level, pcf PCGetterFunc, msg string, args ...any) {
	Default().LogPCFuncf(ctx, level, pcf, msg, args...)
}

func LogPC(ctx context.Context, level slog.Level, pcGetter PCGetter, msg string, args ...any) {
	Default().LogPC(ctx, level, pcGetter, msg, args...)
}

func LogPCAttrs(ctx context.Context, level slog.Level, pcGetter PCGetter, msg string, attrs ...slog.Attr) {
	Default().LogPCAttrs(ctx, level, pcGetter, msg, attrs...)
}

func LogPCf(ctx context.Context, level slog.Level, pcGetter PCGetter, msg string, args ...any) {
	Default().LogPCf(ctx, level, pcGetter, msg, args...)
}
