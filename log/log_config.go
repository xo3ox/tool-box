package log

import (
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type LogConfig struct {
	File               string `yaml:"file"`           // 日志路径
	Level              string `yaml:"level"`          // 日志等级
	MaxSize            int    `yaml:"maxSize"`        // 日志最大100M(megabytes)
	MaxBackups         int    `yaml:"maxBackups"`     // 日志最大备份数
	MaxAge             int    `yaml:"maxAge"`         // 日志最长时间days
	AbledWebManage     bool   `yaml:"webManageAbled"` // 是否开启日志管理服务
	WebManagePort      string `yaml:"webManagePort"`  // 日志服务端口
	AbledConsoleOutput bool   `yaml:"consoleOutput"`  // 控制台输出，true输出，false不输出
}

// NewLog 初始化log
func (logConfig *LogConfig) NewLog() *logger {
	var _aLevel zap.AtomicLevel
	switch logConfig.Level {
	case "fatal", "FATAL", "Fatal":
		_aLevel = zap.NewAtomicLevelAt(zap.FatalLevel)
	case "error", "ERROR", "Error":
		_aLevel = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "warn", "WARN", "Warn":
		_aLevel = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "info", "INFO", "Info":
		_aLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	default:
		_aLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	if logConfig.AbledWebManage {
		http.HandleFunc("/loglevel", _aLevel.ServeHTTP)
		go func() {
			if err := http.ListenAndServe(logConfig.WebManagePort, nil); err != nil {
				log.Panic(err.Error())
			}
		}()
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logConfig.File,
		MaxSize:    logConfig.MaxSize,
		MaxBackups: logConfig.MaxBackups,
		MaxAge:     logConfig.MaxAge,
	})
	encoder := zap.NewProductionConfig() // new生产者配置
	// 屏蔽调用位置
	// encoder.EncoderConfig.CallerKey = ""
	// 格式化时间显示方式.
	encoder.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//文件和控制台同时输出.
	var syncConfig zapcore.WriteSyncer
	if logConfig.AbledConsoleOutput { // 如果控制台输出.
		syncConfig = zapcore.NewMultiWriteSyncer(w, zapcore.AddSync(os.Stdout))
	} else {
		syncConfig = zapcore.WriteSyncer(w)
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder.EncoderConfig), syncConfig, _aLevel)
	return &logger{
		logger: zap.New(core),
	}
}
