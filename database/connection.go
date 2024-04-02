package database

import (
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
		RwDSN: "host=myhost user=myuser dbname=mydb sslmode=disable password=mypassword",
		RoDSN: "host=myhost user=myuser dbname=mydb sslmode=disable password=mypassword",
	}

	return cfg
}
