package main

import (
    "fmt"
    "time"
    "database/sql"
    _ "database/sql/driver"
    _ "github.com/go-sql-driver/mysql"
)

func dbConnection() *sql.DB {
    db, err := sql.Open("mysql", "root@(127.0.0.1:3306)/hello_go?parseTime=true")
    if err != nil {
        panic(err)
    }
    return db
}

func getDB() error {
    db := dbConnection()
    err := db.Ping()
    if err != nil {
        return err
    }
    return nil
}

func createTable() {
    db := dbConnection()
    query := `
        CREATE TABLE users (
            id INT AUTO_INCREMENT,
            username TEXT NOT NULL,
            password TEXT NOT NULL,
            created_at DATETIME,
            PRIMARY KEY (id)
        );`

    _, err := db.Exec(query)
    if err != nil {
        panic(err)
    }
}

func getUsers(db *sql.DB) []string {

    type user struct {
        id int
        username string
        password string
        createdAt time.Time
    }

    rows, err := db.Query("SELECT id, username, password, created_at FROM USERS")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    var names []string
    for rows.Next() {
        var u user
        err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
        if err != nil {
            panic(err)
        }
        names = append(names, u.username)
    }
    return names
}

func getUser(db *sql.DB, id int64) string {
    query := `
        SELECT username
        FROM users
        WHERE id = ?
        ;`

    var username string
    err := db.QueryRow(query, 1).Scan(&username)
    if err != nil {
        panic(err)
    }
    fmt.Printf("id: %d, username: %s\n", id, username)
    return username
}

func createUser(db *sql.DB, username string, password string) int64 {
    createdAt := time.Now()
    result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
    if err != nil {
        panic(err)
    }

    userId, err := result.LastInsertId()
    return userId
}
