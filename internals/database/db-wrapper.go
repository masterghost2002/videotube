package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	config "github.com/masterghost2002/videotube/configs"
)

var Storage *Queries

func InitDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s sslrootcert=%s ", config.ENVS.DBAddress, config.ENVS.DBUser, config.ENVS.DBPassword, config.ENVS.DBName, config.ENVS.DBPort, config.ENVS.SSLMODE, config.ENVS.CRT_PATH)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	if err := conn.Ping(); err != nil {
		log.Fatal("failed to connect with database: ", err)
	}

	//wrap up the database in the new of queries
	Storage = New(conn)
	return nil

}
func GetDB() *Queries {
	return Storage
}
