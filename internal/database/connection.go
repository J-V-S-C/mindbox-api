package database

import (
	"database/sql"
	"fmt"

	"github.com/J-V-S-C/MindBox/internal/config"
	"github.com/J-V-S-C/MindBox/internal/utils"
)

func Connect() *sql.DB{
	cfg := config.New()
	psqlconn := cfg.GetDSN()
	checkError := utils.CheckError

	db, err := sql.Open("postgres", psqlconn)
	checkError(err)

	err = db.Ping()
	checkError(err)

	fmt.Println("Connected!")
	return db
}
