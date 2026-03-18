package launcher

import (
	"Actium_Todo/internal/config"
	db "Actium_Todo/internal/database"
	"Actium_Todo/internal/service"
	"Actium_Todo/internal/transport/api"
	"log"
)

func Start() {

	//Reading config files
	err := config.ReadConfig("internal/config/config.json")
	if err != nil {
		log.Fatal("Не удалось прочитать конфиг файл!", err)
	}

	log.Println(*config.GetConf())

	//Connecting Database
	dbData := config.GetConf().Database
	err = db.ConnectDB(dbData.Username, dbData.Password, dbData.DBName, dbData.Address)
	if err != nil {
		log.Fatal("Ошибка при подключении к БД!", err)
	}

	defer db.CloseDB()

	//Run the application
	service.Run()
}

func StartWeb() {

	//Reading Config files
	err := config.ReadConfig("internal/config/config.json")
	if err != nil {
		log.Fatal("Не удалось прочитать конфиг файл!", err)
	}
	config.LoadEnv()

	//Connecting to the BD
	dbData := config.GetConf().Database
	err = db.ConnectDB(dbData.Username, dbData.Password, dbData.DBName, dbData.Address)
	if err != nil {
		log.Fatal("Ошибка при подключении к БД!", err)
	}

	defer db.CloseDB()

	//Launching the server
	err = api.InitRouter()
	if err != nil {
		log.Fatal("Ошибка при запуске сервера!", err)
	}
}
