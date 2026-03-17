package launcher

import (
	"Actium_Todo/internal/config"
	db "Actium_Todo/internal/database"
	"Actium_Todo/internal/service"
	"log"
)

func Start() {

	//Чтение конфиг-файла
	err := config.ReadConfig("internal/config/config.json")
	if err != nil {
		log.Fatal("Не удалось прочитать конфиг файл!", err)
	}

	log.Println(*config.GetConf())

	//Подключение к базе данных
	dbData := config.GetConf().Database
	err = db.ConnectDB(dbData.Username, dbData.Password, dbData.DBName, dbData.Address)
	if err != nil {
		log.Fatal("Ошибка при подключении к БД!", err)
	}

	defer db.CloseDB()

	//Run the application
	service.Run()
}
