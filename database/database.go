package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() error {
    var err error
    dsn := "root:@tcp(localhost:3306)/singlemarket"
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        return fmt.Errorf("error opening database connection: %v", err)
    }

    err = db.Ping()
    if err != nil {
        return fmt.Errorf("error pinging database: %v", err)
    }

    err = migrate()
    if err != nil {
        return fmt.Errorf("error during migration: %v", err)
    }

    return nil
}

func GetDB() *sql.DB {
    return db
}

func migrate() error {
    queries := []string{
        `CREATE TABLE IF NOT EXISTS clients (
            id INT AUTO_INCREMENT PRIMARY KEY,
            first_name VARCHAR(255) NOT NULL,
            last_name VARCHAR(255) NOT NULL,
            phone VARCHAR(20),
            address TEXT,
            email VARCHAR(255)
        )`,
        `CREATE TABLE IF NOT EXISTS products (
            id INT AUTO_INCREMENT PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            description TEXT,
            price DECIMAL(10, 2) NOT NULL,
            quantity INT NOT NULL,
            active BOOLEAN NOT NULL DEFAULT true
        )`,
        `CREATE TABLE IF NOT EXISTS orders (
            id INT AUTO_INCREMENT PRIMARY KEY,
            client_id INT,
            product_id INT,
            quantity INT NOT NULL,
            price DECIMAL(10, 2) NOT NULL,
            purchase_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (client_id) REFERENCES clients(id),
            FOREIGN KEY (product_id) REFERENCES products(id)
        )`,
    }

    for _, query := range queries {
        _, err := db.Exec(query)
        if err != nil {
            return err
        }
    }
    return nil
}
