package Log

import (
	"log"
	"os"
)

func WriteLog(message interface{}) {
	FileLog, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Ошибка в чтении файла логов FILE_LOG_NOT_READED", err)
	}
	defer FileLog.Close()

	log.SetOutput(FileLog)

	switch v := message.(type) {
	case error:
		log.Println("ERROR:", v)
	case string:
		log.Println("INFO:", v)
	default:
		log.Println(v)
	}
}
