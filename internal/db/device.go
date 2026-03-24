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

func DeleteDevice(db *sql.DB, publicKey string) error {

	query := `
	DELETE FROM devices
	WHERE public_key = $1
	`

	_, err := db.Exec(query, publicKey)
	return err
}

func GetDevices(db *sql.DB) ([]map[string]interface{}, error) {

	rows, err := db.Query("SELECT id, public_key, vpn_ip, created_at FROM devices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []map[string]interface{}

	for rows.Next() {
		var id int
		var publicKey, vpnIP, createdAt string

		err := rows.Scan(&id, &publicKey, &vpnIP, &createdAt)
		if err != nil {
			return nil, err
		}

		device := map[string]interface{}{
			"id":         id,
			"public_key": publicKey,
			"vpn_ip":     vpnIP,
			"created_at": createdAt,
		}

		devices = append(devices, device)
	}

	return devices, nil
}

func GetDeviceByID(db *sql.DB, id int) (string, error) {

	var publicKey string

	query := `
	SELECT public_key FROM devices
	WHERE id = $1
	`

	err := db.QueryRow(query, id).Scan(&publicKey)
	if err != nil {
		return "", err
	}

	return publicKey, nil
}