package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Declaração de variáveis de ambiente para o logger
var (
	log        *zap.Logger    // Variável privada para o logger
	LOG_OUTPUT = "LOG_OUTPUT" // Variável de ambiente para o output do log
	LOG_LEVEL  = "LOG_LEVEL"  // Variável de ambiente para o nível do log
)

func init() { // Função para inicializar o logger
	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogs()},            // Obtém o output do log
		Level:       zap.NewAtomicLevelAt(getLevelLogs()), // Obtém o nível do log
		Encoding:    "json",                               // Define o formato do log como JSON
		EncoderConfig: zapcore.EncoderConfig{ // Configuração do encoder
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build() // Constroi o logger
}

// Função para registrar logs de informação
func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...) // Registra o log de informação
	log.Sync()                 // Garante que os logs sejam escritos imediatamente
}

// Função para registrar logs de erro
func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	log.Sync() 
}

// Função para obter o output dos logs
func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout" // Se a variável de ambiente estiver vazia, retorna "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
