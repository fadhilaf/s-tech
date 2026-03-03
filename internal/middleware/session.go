package middleware

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware(db *sql.DB) gin.HandlerFunc {
	sessionKey := os.Getenv("APP_SESSION_KEY")
	store, err := postgres.NewStore(db, []byte(sessionKey)) //bagusny ambil dari .env
	if err != nil {
		log.Fatalf("Terjadi error dalam pembuatan session store: %v", err)
	}

	return sessions.Sessions("mysession", store)
}
