package logger

import (
	"os"
	"time"

	"github.com/leeseika/feature-show/settings"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const debugLevel = "debug"
const infoLevel = "info"
const errorLevel = "error"
const panicLevel = "panic"
const fatalLevel = "fatal"
const writeBoth = "both"
const writeConsole = "console"
const writeFile = "file"

func Init() error {
	config := settings.Conf.Log
	// 构建编码器
	encoder := zapEncoder(config)
	// 构建日志级别
	levelEnabler := zapLevelEnabler(config)
	// 最后获得Core和Options
	subCore, options := tee(config, encoder, levelEnabler)
	// 创建Logger
	logger := zap.New(subCore, options...)

	zap.ReplaceGlobals(logger)

	return nil
}

// 将所有合并
func tee(cfg *settings.Log, encoder zapcore.Encoder, levelEnabler zapcore.LevelEnabler) (core zapcore.Core, options []zap.Option) {
	sink := zapWriteSyncer(cfg)
	return zapcore.NewCore(encoder, sink, levelEnabler), buildOptions(cfg, levelEnabler)
}

// 构建Option
func buildOptions(cfg *settings.Log, levelEnabler zapcore.LevelEnabler) (options []zap.Option) {
	if cfg.Caller {
		options = append(options, zap.AddCaller())
	}

	if cfg.Stacktrace {
		options = append(options, zap.AddStacktrace(levelEnabler))
	}
	return
}
func zapEncoder(config *settings.Log) zapcore.Encoder {
	// 新建一个配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "Time",
		LevelKey:      "Level",
		NameKey:       "Logger",
		CallerKey:     "Caller",
		MessageKey:    "Message",
		StacktraceKey: "StackTrace",
		LineEnding:    zapcore.DefaultLineEnding,
		FunctionKey:   zapcore.OmitKey,
	}
	// 自定义时间格式
	encoderConfig.EncodeTime = CustomTimeFormatEncoder
	// 日志级别大写
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// 秒级时间间隔
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	// 简短的调用者输出
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 完整的序列化logger名称
	encoderConfig.EncodeName = zapcore.FullNameEncoder
	// 最终的日志编码 json或者console
	switch config.Encode {
	case "json":
		{
			return zapcore.NewJSONEncoder(encoderConfig)
		}
	case "console":
		{
			return zapcore.NewConsoleEncoder(encoderConfig)
		}
	}
	// 默认console
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func zapLevelEnabler(cfg *settings.Log) zapcore.LevelEnabler {
	switch cfg.Level {
	case debugLevel:
		return zap.DebugLevel
	case infoLevel:
		return zap.InfoLevel
	case errorLevel:
		return zap.ErrorLevel
	case panicLevel:
		return zap.PanicLevel
	case fatalLevel:
		return zap.FatalLevel
	}
	// 默认Debug级别
	return zap.DebugLevel
}
func zapWriteSyncer(cfg *settings.Log) zapcore.WriteSyncer {
	syncers := make([]zapcore.WriteSyncer, 0, 2)
	// 如果开启了日志控制台输出，就加入控制台书写器
	if cfg.Writer == writeBoth || cfg.Writer == writeConsole {
		syncers = append(syncers, zapcore.AddSync(os.Stdout))
	}

	// 如果开启了日志文件存储，就根据文件路径切片加入书写器
	if cfg.Writer == writeBoth || cfg.Writer == writeFile {
		// 添加日志输出器
		for _, path := range cfg.LogFile.Output {
			logger := &lumberjack.Logger{
				Filename:   path,                 //文件路径
				MaxSize:    cfg.LogFile.MaxSize,  //分割文件的大小
				MaxAge:     cfg.LogFile.MaxAge,   //已经被分割存储的日志文件最大的留存时间，单位是天
				MaxBackups: cfg.LogFile.Backups,  //分割存储的日志文件最多的留存个数
				Compress:   cfg.LogFile.Compress, //是否压缩
				LocalTime:  true,                 //使用本地时间
			}

			syncers = append(syncers, zapcore.Lock(zapcore.AddSync(logger)))
		}
	}

	return zap.CombineWriteSyncers(syncers...)
}

func CustomTimeFormatEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString("feature-show" + "\t" + t.Format(settings.Conf.Log.TimeFormat))
}
