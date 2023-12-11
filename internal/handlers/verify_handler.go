package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

type VerifyHandler struct {
	DB *sql.DB
}

func NewVerifyHandler(db *sql.DB) *VerifyHandler {
	return &VerifyHandler{DB: db}
}