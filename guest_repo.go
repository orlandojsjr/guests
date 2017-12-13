package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

func FindGuests(InvitedBy string) Guest {
    db, err := sql.Open("mysql", "root:abc123@tcp(127.0.0.1:3306)/guests?parseTime=true")

    rows, err := db.Query("SELECT name, email, COALESCE(message,''), amount, invited_by, created_at FROM guests where id = ?", 1)

    if err != nil {
         log.Fatal(err)
    }
    defer rows.Close()
    defer db.Close()

    for rows.Next() {
        result := new(Guest)
        err := rows.Scan(&result.Name, &result.Email, &result.Message, &result.Amount, &result.InvitedBy, &result.CreatedAt)
        if err != nil {
            log.Fatal(err)
        }
        return Guest{Name: result.Name, Email: result.Email, Message: result.Message, Amount: result.Amount, InvitedBy: result.InvitedBy, CreatedAt: result.CreatedAt}
    }

    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    return Guest{}
}

func CreateGuest(guest Guest) Guest {
    return guest
}

