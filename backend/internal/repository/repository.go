package repository

import (
	"database/sql"
	"ongambl/internal/models"
	"time"
)

type DatabaseRepo interface {
	Conncetion() *sql.DB
	CreateUser(user models.User) (int, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(id int64) (*models.User, error)
	NewToken(userID int64, ttl time.Duration, scope string) (*models.Token, error)
	InsertToken(token *models.Token) error
	DeleteAllTokensForUser(scope string, userID int64) error
	GetUserByToken(tokenScope, tokenPlainText string) (*models.User, error)
}
