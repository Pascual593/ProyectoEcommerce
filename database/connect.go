package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Driver para conectar a MySQL
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	//vamos a configurar el enlace con mysql
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	//probar la conexion con bd, haciendo un ping a la bd
	if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Conexion exitosa a la base de datos")
	return db, nil

}
