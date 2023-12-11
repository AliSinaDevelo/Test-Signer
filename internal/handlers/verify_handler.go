package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
	"github.com/AliSinaDevelo/Test-Signer/internal/model"
)

type VerifyHandler struct {
	DB *sql.DB
}

func NewVerifyHandler(db *sql.DB) *VerifyHandler {
	return &VerifyHandler{DB: db}
}

func (vh *VerifyHandler) Verify(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	signature := r.URL.Query().Get("signature")

	var fetchedSig Signature
	row := vh.DB.QueryRow("SELECT signature FROM signatures WHERE username = $1 AND signaturee = $2 ", user, signature)
	err := row.Scan(&fetchedSignature.User, &fetchedSignature.Signature, &fetchedSignature.Answers, &fetchedSignature.Timestamp)
	if err != nil {
		http.Error(w, "Invalid signature opr user", http.StatusBadRequest)
		return
	}


	// return the signature
	response := map[string]interface{}{
		"status": "success",
		"user": fetchedSignature.User,
		"answers": fetchedSignature.Answers,
		"timestamp": fetchedSignature.Timestamp.Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}