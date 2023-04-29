package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewZapLog(config ...ZapLogConfig) *zap.SugaredLogger {
	var coreArr []zapcore.Core
	var cfg = configDefault(config...)

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zapcore.Level(cfg.HighLogLevel-2)
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zapcore.Level(cfg.HighLogLevel-2) && lev >= zap.DebugLevel
	})

	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.Filepath + "/" + cfg.LowLogName,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	})
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority)
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.Filepath + "/" + cfg.HighLogName,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	})
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority)

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)
	return zap.New(zapcore.NewTee(coreArr...)).Sugar()
}
