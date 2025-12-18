package utils

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetOutput(os.Stdout)

	// Define formatação como JSON estruturado
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     false,
	})

	// Define nível padrão (pode ser ajustado por env depois)
	logger.SetLevel(logrus.InfoLevel)
}

// GetLogger retorna instância do logger principal
func GetLogger() *logrus.Logger {
	return logger
}

// LogError é um atalho para logar erros com contexto
func LogError(err error, context map[string]interface{}) {
	if err == nil {
		return
	}
	entry := logger.WithFields(logrus.Fields(context))
	entry.Error(err)
}

// LogInfo loga uma mensagem informativa com contexto
func LogInfo(msg string, context map[string]interface{}) {
	entry := logger.WithFields(logrus.Fields(context))
	entry.Info(msg)
}
