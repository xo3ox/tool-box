package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	logger *zap.Logger // 指向zap.   logger   的结构体

}

// ------------------------------------------------------------------ 日志级别 ------------------------------------------------------------------
// Debug 输出 Debug 级别日志
func (l *logger) Debug(message string, fields ...zapcore.Field) {
	l.logger.Debug(message, fields...)
}

// Debug 输出带Caller信息的Debug级别日志.
func (l *logger) DebugWithStack(fields ...zapcore.Field) func(...zapcore.Field) {
	fields = append(fields, zap.StackSkip("stack", 1))
	l.logger.Debug("method start", fields...)
	return func(fields1 ...zapcore.Field) {
		fields1 = append(fields1, zap.StackSkip("stack", 1))
		l.logger.Debug("method end", fields1...)
	}
}

// Info 输出 Info 级别日志
func (l *logger) Info(message string, fields ...zapcore.Field) {
	l.logger.Info(message, fields...)
}

// Warn 输出 Warn 级别日志
func (l *logger) Warn(message string, fields ...zapcore.Field) {
	l.logger.Warn(message, fields...)
}

// Error 输出 Error 级别日志
func (l *logger) Error(message string, fields ...zapcore.Field) {
	l.logger.Error(message, fields...)
}

// Fatal 输出 Fatal 级别日志
func (l *logger) Fatal(message string, fields ...zapcore.Field) {
	l.logger.Fatal(message, fields...)
}

// ------------------------------------------------------------------ key-value ------------------------------------------------------------------
// AnyField 任意类型
func (l *logger) AnyField(key string, v interface{}) zap.Field {
	return zap.Any(key, v)
}

// BinaryField base64编码的二进制字符串
func (l *logger) BinaryField(key string, v []byte) zap.Field {
	return zap.Binary(key, v)
}

// IntField int类型
func (l *logger) IntField(k string, v int) zap.Field {
	return (zap.Int(k, v))
}

// StringField字串类型
func (l *logger) StringField(k, v string) zap.Field {
	return zap.String(k, v)
}

// BoolField 布尔型
func (l *logger) BoolField(key string, v bool) zap.Field {
	return zap.Bool(key, v)
}

// TimeField 时间戳类型
func (l *logger) TimeField(key string, v time.Time) zap.Field {
	return zap.Time(key, v)
}

// ErrorField 错误类型
func (l *logger) ErrorField(v error) zap.Field {
	return zap.Error(v)
}
