package config

type ZapConfig struct {
	Prefix     string         `yaml:"prefix" mapstructure:"prefix"`
	TimeFormat string         `yaml:"timeFormat" mapstructure:"timeFormat"`
	Level      string         `yaml:"level" mapstructure:"level"`
	Caller     bool           `yaml:"caller" mapstructure:"caller"`
	StackTrace bool           `yaml:"stackTrace" mapstructure:"stackTrace"`
	Writer     string         `yaml:"writer" mapstructure:"writer"`
	Encode     string         `yaml:"encode" mapstructure:"encode"`
	LogFile    *LogFileConfig `yaml:"logFile" mapstructure:"logFile"`
}

type LogFileConfig struct {
	MaxSize  int      `yaml:"maxSize" mapstructure:"maxSize"`
	MaxAge   int      `yaml:"maxAge" mapstructure:"maxAge"`
	BackUps  int      `yaml:"backups" mapstructure:"backups"`
	Compress bool     `yaml:"compress" mapstructure:"compress"`
	Output   []string `yaml:"output" mapstructure:"output"`
}

const DebugLevel = "debug"
const InfoLevel = "info"
const ErrorLevel = "error"
const PanicLevel = "panic"
const FatalLevel = "fatal"
const WriteBoth = "both"
const WriteConsole = "console"
const WriteFile = "file"
