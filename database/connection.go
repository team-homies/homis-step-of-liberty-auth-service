package database

import (
	"fmt"
	"main/config"
	"main/database/entity"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

type DBConfig struct {
	RwDSN string
	RoDSN string
}

func InitDB() (*gorm.DB, error) {
	dns := generateDialector()
	db, err := gorm.Open(postgres.Open(dns.RwDSN), &gorm.Config{})
	db.AutoMigrate(&entity.User{}, &entity.Visual{})
	if err != nil {
		return nil, err

	}

	err = db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{postgres.Open(dns.RoDSN)},
		Policy:   dbresolver.RandomPolicy{},
	}).SetMaxOpenConns(50).SetMaxIdleConns(10))
	if err != nil {
		return nil, err
	}

	DB = db

	return DB, nil
}

func generateDialector() DBConfig {
	cfg := DBConfig{
		// viper에서 dns 읽어오기
		RwDSN: fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			viper.GetString(config.DATABASE_RW_USER),
			viper.GetString(config.DATABASE_RW_PASSWORD),
			viper.GetString(config.DATABASE_HOST),
			viper.GetInt(config.DATABASE_PORT),
			viper.GetString(config.DATABASE_DATABASE),
		),
		RoDSN: fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			viper.GetString(config.DATABASE_RO_USER),
			viper.GetString(config.DATABASE_RO_PASSWORD),
			viper.GetString(config.DATABASE_HOST),
			viper.GetInt(config.DATABASE_PORT),
			viper.GetString(config.DATABASE_DATABASE),
		),
	}

	return cfg
}
