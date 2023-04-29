package log

type ZapLogConfig struct {
	Filepath     string // 日志文件夹路径，默认："./log"
	MaxSize      int    // 单文件最大大小，默认：20mb
	MaxBackups   int    // 日志文件的最大数量，默认：5个
	MaxAge       int    // 日志文件保留天数，默认：30天
	Compress     bool   // 启用gzip压缩，默认：true
	HighLogLevel int    // 高等级日志级别，大于等于该级别的日志，会存入高等级日志文件。1=Debug 2=Info 3=Warn 4=Error（默认值） 5=DPanic 6=Panic 7=Fatal
	HighLogName  string // 高等级日志文件名称，默认："error.log"
	LowLogName   string // 低等级日志文件名称，默认："info.log"
}

var defaultZapLogConfig = ZapLogConfig{
	Filepath:     "./log",
	MaxSize:      20,
	MaxBackups:   5,
	MaxAge:       30,
	Compress:     true,
	HighLogLevel: 4,
	HighLogName:  "error.log",
	LowLogName:   "info.log",
}

func configDefault(config ...ZapLogConfig) ZapLogConfig {
	if len(config) < 1 {
		return defaultZapLogConfig
	}

	cfg := config[0]
	if cfg.Filepath == "" {
		cfg.Filepath = defaultZapLogConfig.Filepath
	}
	if cfg.MaxSize == 0 {
		cfg.MaxSize = defaultZapLogConfig.MaxSize
	}
	if cfg.MaxBackups == 0 {
		cfg.MaxBackups = defaultZapLogConfig.MaxBackups
	}
	if cfg.MaxAge == 0 {
		cfg.MaxAge = defaultZapLogConfig.MaxAge
	}
	if cfg.Compress {
		cfg.Compress = defaultZapLogConfig.Compress
	}
	if cfg.HighLogLevel == 0 {
		cfg.HighLogLevel = defaultZapLogConfig.HighLogLevel
	}
	if cfg.HighLogName == "" {
		cfg.HighLogName = defaultZapLogConfig.HighLogName
	}
	if cfg.LowLogName == "" {
		cfg.LowLogName = defaultZapLogConfig.LowLogName
	}
	return cfg
}
