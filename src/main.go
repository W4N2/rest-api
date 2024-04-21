package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/w4n2/rest-api/src/api"
	"github.com/w4n2/rest-api/src/config"
	"github.com/w4n2/rest-api/src/internal/db"
	"github.com/w4n2/rest-api/src/store"
)

func main() {
	cfg := mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStorage := db.NewMySQLStorage(cfg)

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := store.NewStore(db)
	api := api.NewAPIServer(":3000", store)

	api.Serve()
}
