package main

import "time"

type Guest struct {
    Name        string    `json:"name"`
    Email       string    `json:"email"`
    Message     string    `json:"message"`
    Amount      string    `json:"amount"`
    InvitedBy   string    `json:"invited_by"`
    CreatedAt   time.Time `json:"created_at"`
}

type Guests []Guest

