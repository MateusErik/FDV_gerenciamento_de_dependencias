package internal

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	// Cria/abre o arquivo de log
	file, err := os.OpenFile("logs/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Não foi possível abrir ou criar o arquivo de log: " + err.Error())
	}

	// Inicializa o logger
	Log = logrus.New()
	Log.SetOutput(file)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Log.SetLevel(logrus.InfoLevel)
}
