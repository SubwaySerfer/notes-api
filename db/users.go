package db

import (
    "database/sql"
)

func CreateTable(db *sql.DB) error {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        age INTEGER NOT NULL
    );`
    _, err := db.Exec(query)
    return err
}

func InsertUser(db *sql.DB, name string, age int) error {
    query := `INSERT INTO users (name, age) VALUES (?, ?)`
    _, err := db.Exec(query, name, age)
    return err
}

func GetUsers(db *sql.DB) ([]map[string]interface{}, error) {
    rows, err := db.Query("SELECT id, name, age FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []map[string]interface{}
    for rows.Next() {
        var id int
        var name string
        var age int

        err := rows.Scan(&id, &name, &age)
        if err != nil {
            return nil, err
        }

        users = append(users, map[string]interface{}{
            "id":   id,
            "name": name,
            "age":  age,
        })
    }
    return users, nil
}
