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