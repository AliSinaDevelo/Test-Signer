package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
	"github.com/AliSinaDevelo/Test-Signer/internal/model"
	"github.com/dgrijalva/jwt-go"
)

type SignHandler struct {
	DB *sql.DB
}

func NewSignHandler(db *sql.DB) *SignHandler {
	return &SignHandler{DB: db}
}

func (h *SignHandler) Sign(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Username string `json:"username"`
		Answer  []model.QuestionAnswer `json:"answer"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": requestData.Username,
		"answer": requestData.Answer,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// save token to database

	_, err = h.DB.Exec("INSERT INTO signatures (user_id, signature, answer, timestamp) VALUES ($1, $2, $3, $4)", requestData.Username, tokenString, requestData.Answer, time.Now())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"token": tokenString,
	}
	json.NewEncoder(w).Encode(response)
}
