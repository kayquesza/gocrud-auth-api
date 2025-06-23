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

func Info(message string, tags ...zap.Field) { // Função para registrar logs de informação
	log.Info(message, tags...) // Registra o log de informação
	log.Sync()                 // Garante que os logs sejam escritos imediatamente
}

func Error(message string, err error, tags ...zap.Field) { // Função para registrar logs de erro
	tags = append(tags, zap.NamedError("error", err)) // Criar um campo de erro no log
	log.Info(message, tags...)
	log.Sync() // Garante que os logs sejam escritos imediatamente
}

func getOutputLogs() string { // Função que lê a v.A LOG_OUTPUT
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT))) // Remove espaços em branco e converte para minúsculas
	if output == "" {
		return "stdout" // Se a v.A estiver vazia, retorna "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level { // Função que lê a v.A LOG_LEVEL
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) { // Remove espaços em branco e converte para minúsculas
	case "info":
		return zapcore.InfoLevel // Se a v.A for "info", retorna o nível de informação
	case "error":
		return zapcore.ErrorLevel // Se a v.A for "error", retorna o nível de erro
	case "debug":
		return zapcore.DebugLevel // Se a v.A for "debug", retorna o nível de depuração
	default:
		return zapcore.InfoLevel // Se a v.A for "info", retorna como padrão
	}
}
