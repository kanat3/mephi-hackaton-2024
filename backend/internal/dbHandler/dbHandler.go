package database

import (
	"fmt"

	"backend/internal/config"
	"backend/internal/dbHandler/interfaces"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbHandler struct {
	db *gorm.DB
}

func getDbHandler() interfaces.DbHandler {
	cfg := config.GetConfig()
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBname)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	DbHandler := new(DbHandler)
	DbHandler.db = db
	return DbHandler
}

func (handler *DbHandler) Create(obj interface{}) {
	handler.db.Create(obj)
}

func (handler *DbHandler) FindAll(obj interface{}) {
	handler.db.Find(obj)
}

func (handler *DbHandler) DeleteById(obj interface{}, id string) {
	handler.db.Delete(obj, id)
}

func (handler *DbHandler) SelectById(obj interface{}, id string) {
	handler.db.Select(obj, id)
}

func (handler *DbHandler) Where(object interface{}, args ...interface{}) (tx *gorm.DB) {
	return handler.db.Where(object, args)
}

func (handler *DbHandler) Preload(query string, args ...interface{}) (tx *gorm.DB) {
	return handler.db.Preload(query, args)
}
