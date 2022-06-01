package log

import (
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const SecurityMark = "[Security Mark]"

// function variables for all field types
// in github.com/uber-go/zap/field.go

type Core = zapcore.Core
type Level = zapcore.Level
type Field = zap.Field
type Option = zap.Option
type Config = zap.Config
type EncoderConfig = zapcore.EncoderConfig
type EncodeType = int
type LevelEnablerFunc func(lvl Level) bool
type RotateOptions struct {
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
	LocalTime  bool
}
type TeeOption struct {
	Filename string
	Ropt     RotateOptions
	Lef      LevelEnablerFunc
}
type Logger struct {
	zaplog *zap.Logger // zap ensure that zap.Logger is safe for concurrent use
	level  Level
}

const (
	InfoLevel   Level = zap.InfoLevel   // 0, default level
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log
	PanicLevel  Level = zap.PanicLevel  // 4, PanicLevel logs a message, then panics
	FatalLevel  Level = zap.FatalLevel  // 5, FatalLevel logs a message, then calls os.Exit(1).
	DebugLevel  Level = zap.DebugLevel  // -1
)
const (
	ConsoleEncoder EncodeType = iota // 0, default 控制台编码器
	JSONEncoder
)

var (
	std = New(os.Stderr, DebugLevel, WithCaller(true), AddCallerSkip(0))
	// defined
	Skip        = zap.Skip
	Binary      = zap.Binary
	Bool        = zap.Bool
	Boolp       = zap.Boolp
	ByteString  = zap.ByteString
	Complex128  = zap.Complex128
	Complex128p = zap.Complex128p
	Complex64   = zap.Complex64
	Complex64p  = zap.Complex64p
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Int         = zap.Int
	Intp        = zap.Intp
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	String      = zap.String
	Stringp     = zap.Stringp
	Uint        = zap.Uint
	Uintp       = zap.Uintp
	Uint64      = zap.Uint64
	Uint64p     = zap.Uint64p
	Uint32      = zap.Uint32
	Uint32p     = zap.Uint32p
	Uint16      = zap.Uint16
	Uint16p     = zap.Uint16p
	Uint8       = zap.Uint8
	Uint8p      = zap.Uint8p
	Uintptr     = zap.Uintptr
	Uintptrp    = zap.Uintptrp
	Reflect     = zap.Reflect
	Namespace   = zap.Namespace
	Stringer    = zap.Stringer
	Time        = zap.Time
	Timep       = zap.Timep
	Stack       = zap.Stack
	StackSkip   = zap.StackSkip
	Errors      = zap.Errors
	ErrorField  = zap.Error
	Duration    = zap.Duration
	Durationp   = zap.Durationp
	Any         = zap.Any

	WithCaller    = zap.WithCaller
	AddCallerSkip = zap.AddCallerSkip
	AddStacktrace = zap.AddStacktrace

	// core
	NewCore           = zapcore.NewCore
	NewNopCore        = zapcore.NewNopCore
	AddSync           = zapcore.AddSync
	NewJSONEncoder    = zapcore.NewJSONEncoder
	NewConsoleEncoder = zapcore.NewConsoleEncoder

	// encoder config
	NewDevelopmentEncoderConfig = zap.NewDevelopmentEncoderConfig
	NewProductionEncoderConfig  = zap.NewProductionEncoderConfig

	Named  = std.zaplog.Named
	Info   = std.zaplog.Info
	Warn   = std.zaplog.Warn
	Error  = std.zaplog.Error
	DPanic = std.zaplog.DPanic
	Panic  = std.zaplog.Panic
	Fatal  = std.zaplog.Fatal
	Debug  = std.zaplog.Debug
	Sugar  = std.zaplog.Sugar
	Debugf = std.zaplog.Sugar().Debugf
	Infof  = std.zaplog.Sugar().Infof
	Warnf  = std.zaplog.Sugar().Warnf
	Errorf = std.zaplog.Sugar().Errorf
	Fatalf = std.zaplog.Sugar().Fatalf
)

func Default() *Logger {
	return std
}
func ResetDefault(newStd *Logger) {
	std = newStd

	Named = std.zaplog.Named
	Info = std.zaplog.Info
	Warn = std.zaplog.Warn
	Error = std.zaplog.Error
	DPanic = std.zaplog.DPanic
	Panic = std.zaplog.Panic
	Fatal = std.zaplog.Fatal
	Debug = std.zaplog.Debug
	Sugar = std.zaplog.Sugar
	Debugf = std.zaplog.Sugar().Debugf
	Infof = std.zaplog.Sugar().Infof
	Warnf = std.zaplog.Sugar().Warnf
	Errorf = std.zaplog.Sugar().Errorf
	Fatalf = std.zaplog.Sugar().Fatalf
}

func StringSecurity(key string, v string) Field {
	return Field{Key: key, Type: zapcore.StringType, String: v}
}

func SetLevel(level string) {
	std.level.Set(level)
}

// New create a new logger (not support log rotating).
func New(writer io.Writer, level Level, opts ...Option) *Logger {
	checkWriter(writer)

	encoder := zap.NewDevelopmentEncoderConfig()
	encoder.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("15:04:05.000"))
	}
	encoder.EncodeLevel = func(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {
		zapcore.CapitalColorLevelEncoder(l, pae)
	}
	encoder.EncodeCaller = func(c zapcore.EntryCaller, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(c.TrimmedPath())
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder),
		zapcore.AddSync(writer),
		zapcore.Level(level),
	)
	logger := &Logger{
		zaplog: zap.New(core, opts...),
		level:  level,
	}
	return logger
}

// NewWithEncoder New Logger with custom EncoderConfig
func NewWithEncoder(writer io.Writer, level Level, encoder EncoderConfig, opts ...Option) *Logger {
	checkWriter(writer)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		zapcore.AddSync(writer),
		zapcore.Level(level),
	)

	logger := &Logger{
		zaplog: zap.New(core, opts...),
		level:  level,
	}
	return logger
}

// NewCustom New Logger with any config
func NewCustom(core Core, opts ...Option) *Logger {
	z := zap.New(core, opts...)

	logger := &Logger{
		zaplog: z,
		level:  DPanicLevel,
	}

	return logger
}

// NewTeeWithRotate New Prod Logger and rotate logger files
func NewTeeWithRotate(tops []TeeOption, opts ...Option) *Logger {
	var cores []zapcore.Core
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	cfg.EncoderConfig.EncodeCaller = func(ec zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(ec.TrimmedPath())
	}

	for _, top := range tops {
		top := top

		lv := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return top.Lef(Level(lvl))
		})

		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   top.Filename,
			MaxSize:    top.Ropt.MaxSize,
			MaxBackups: top.Ropt.MaxBackups,
			MaxAge:     top.Ropt.MaxAge,
			Compress:   top.Ropt.Compress,
			LocalTime:  top.Ropt.LocalTime,
		})

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg.EncoderConfig),
			zapcore.AddSync(w),
			lv,
		)
		cores = append(cores, core)
	}

	logger := &Logger{
		zaplog: zap.New(zapcore.NewTee(cores...), opts...),
	}
	return logger
}

// NewNop New nil output Logger
func NewNop() *Logger {
	logger := &Logger{
		zaplog: zap.NewNop(),
		level:  zap.PanicLevel,
	}
	return logger
}

func (l *Logger) Sync() error {
	return l.zaplog.Sync()
}

func Sync() error {
	if std != nil {
		return std.Sync()
	}
	return nil
}

func checkWriter(writer io.Writer) {
	if writer == nil {
		panic("the writer is nil, if you need nop log, you can try NewNop()")
	}
}
