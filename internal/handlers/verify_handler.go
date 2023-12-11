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

	var fetchedSig model.Signature
	row := vh.DB.QueryRow("SELECT user_id, signature, answers, timestamp FROM signatures WHERE user_id = $1 AND signature = $2", user, signature)
	err := row.Scan(&fetchedSig.User, &fetchedSig.Signature, &fetchedSig.Answers, &fetchedSig.Timestamp)
	if err != nil {
		http.Error(w, "Invalid signature or user", http.StatusBadRequest)
		return
	}


	// return the signature
	response := map[string]interface{}{
		"status": "success",
		"user": fetchedSig.User,
		"answers": fetchedSig.Answers,
		"timestamp": fetchedSig.Timestamp.Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}