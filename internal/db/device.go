package db

import (
	"database/sql"
	"fmt"
	"vpn-manager/internal/model"

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

func GetDevices(db *sql.DB) ([]model.Device, error) {

	rows, err := db.Query("SELECT id, public_key, vpn_ip, created_at FROM devices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []model.Device

	for rows.Next() {
		var d model.Device

		err := rows.Scan(&d.ID, &d.PublicKey, &d.VPNIP, &d.CreatedAt)
		if err != nil {
			return nil, err
		}

		devices = append(devices, d)
	}

	return devices, nil
}

func GetDeviceByID(db *sql.DB, id int) (string, error) {

	var publicKey string

	err := db.QueryRow(
		"SELECT public_key FROM devices WHERE id = $1",
		id,
	).Scan(&publicKey)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", err
	}

	return publicKey, nil
}
