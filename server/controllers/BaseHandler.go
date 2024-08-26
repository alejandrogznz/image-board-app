package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
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
