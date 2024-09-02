package controllers

import (
	"encoding/json"
	"database/sql"
	"fmt"
	"net/http"

	"image-board/model"
)

// BaseHandler will hold everything that controller needs

type BaseHandler struct {
	db *sql.DB
}

func NewBaseHandler (db *sql.DB) *BaseHandler {
	return &BaseHandler {
		db: db,
	}
}

func (h *BaseHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	err := h.db.Ping()
	if err != nil {
		fmt.Println("DB Error")
	}
	w.Write([]byte("HelloWorldHandler!"))
}

func (h *BaseHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query("SELECT * FROM USERS")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	
	// Iterate through all the rows
	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.LastName, &user.FirstName)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	j, _ := json.Marshal(users)
	w.Header().Add("status", "200")
	w.Header().Add("Content-Type", "application/json")
	w.Write(j)

}
