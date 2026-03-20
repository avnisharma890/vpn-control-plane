package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB() (*sql.DB, error) {

	connStr := "postgres://vpnadmin:strongpassword@localhost:5432/vpnmanager"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL")

	return db, nil
}

func CreateDevice(db *sql.DB, publicKey string, vpnIP string) error {

	query := `
	INSERT INTO devices (public_key, vpn_ip)
	VALUES ($1, $2)
	`

	_, err := db.Exec(query, publicKey, vpnIP)

	return err
}