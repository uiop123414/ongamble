package dbrepo

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"ongambl/internal/models"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeOut = time.Second * 3

func (m *PostgresDBRepo) Conncetion() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) CreateUser(user models.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	tx, err := m.DB.Begin()
	if err != nil {
		return 0, err
	}

	stmt := `INSERT INTO users (name, email, password, activated, created_at,
		updated_at)
		values ($1, $2, $3, $4, $5, $6) returning id`

	var newID int

	err = tx.QueryRowContext(ctx, stmt,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		false,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return newID, nil
}

func (m *PostgresDBRepo) GetUserByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	tx, err := m.DB.Begin()
	if err != nil {
		return nil, err
	}

	query := `SELECT 
				id, name, email, password, activated, created_at, updated_at, version 
			  FROM
			  	users
			  WHERE
				name = $1`

	var u models.User

	err = tx.QueryRowContext(ctx, query, username).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.PasswordHash,
		&u.Activated,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Version,
	)

	if err != nil {
		tx.Rollback()
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, models.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	tx.Commit()

	return &u, nil
}

func (m *PostgresDBRepo) GetUserByID(id int64) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	tx, err := m.DB.Begin()
	if err != nil {
		return nil, err
	}

	query := `SELECT 
				id, name, email, password, activated, created_at, updated_at, version
			  FROM 
			  	users
			  WHERE
			   id = $1`

	var u models.User
	err = tx.QueryRowContext(ctx, query, id).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.PasswordHash,
		&u.Activated,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Version,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &u, nil
}

func (m *PostgresDBRepo) NewToken(userID int64, ttl time.Duration, scope string) (*models.Token, error) {
	token, err := models.GenerateToken(userID, ttl, scope)
	fmt.Println("Old token ", token)
	if err != nil {
		return nil, err
	}

	err = m.InsertToken(token)
	return token, err
}

func (m *PostgresDBRepo) InsertToken(token *models.Token) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	tokenHash := sha256.Sum256(token.Hash)

	query := `INSERT INTO tokens (hash, user_id, expiry, scope)
			  VALUES ($1, $2, $3, $4);`

	args := []interface{}{tokenHash[:], token.UserID, token.Expiry, token.Scope}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (m *PostgresDBRepo) DeleteAllTokensForUser(scope string, userID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	query := `DELETE FROM tokens
			  WHERE scope = $1 AND user_id = $2`

	_, err = tx.ExecContext(ctx, query, scope, userID)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (m *PostgresDBRepo) GetUserByToken(tokenScope, tokenPlainText string) (*models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.Begin()
	if err != nil {
		return nil, err
	}

	query := `
			SELECT users.id, users.created_at, users.name, users.email, users.password, users.activated, users.version
			FROM users
			INNER JOIN tokens
			ON users.id = tokens.user_id
			WHERE tokens.hash = $1
			AND tokens.scope = $2
			AND tokens.expiry > $3`

	tokenHash := sha256.Sum256([]byte(tokenPlainText))

	args := []interface{}{tokenHash[:], tokenScope, time.Now()}

	var user models.User

	err = tx.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Activated,
		&user.Version,
	)
	if err != nil {
		tx.Rollback()

		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, models.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	tx.Commit()

	return &user, nil
}

func (m *PostgresDBRepo) NewArticle(article *models.Article) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO articles (name, publish, reading_time, username, html_list)
	VALUES ($1, $2, $3, $4, $5);`

	_, err = tx.ExecContext(ctx, query, article.Name, article.Publish, article.ReadingTime, article.Username, article.HtmlList)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (m *PostgresDBRepo) GetArticle(id int) (*models.Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.Begin()
	if err != nil {
		return nil, err
	}

	query := `SELECT name, publish, reading_time, username, html_list, created_at, updated_at, version 
				FROM articles
				WHERE id = $1`

	var article models.Article

	err = tx.QueryRowContext(ctx, query, id).Scan(
		&article.Name,
		&article.Publish,
		&article.ReadingTime,
		&article.Username,
		&article.HtmlList,
		&article.CreatedAt,
		&article.UpdatedAt,
		&article.Version,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &article, nil
}

func (m *PostgresDBRepo) GetNews(page int) ([]models.News, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.Begin()
	if err != nil {
		return nil, err
	}

	query := `SELECT id, name, created_at FROM articles WHERE publish=true ORDER BY created_at, id LIMIT $1 OFFSET $2`

	var newsArr []models.News

	limit := 9
	offset := (page - 1) * 9

	rows, err := tx.QueryContext(ctx, query, limit, offset)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var news models.News

		err = rows.Scan(
			&news.ID,
			&news.Name,
			&news.CreatedAt,
		)

		if err != nil {
			tx.Rollback()
			return nil, err
		}

		newsArr = append(newsArr, news)
	}

	if err = rows.Err(); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return newsArr, err
}

func (m *PostgresDBRepo) GetUserPermissions(token string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.Begin()
	if err != nil {
		return nil, err
	}

	query := `  
	SELECT 
		code 
	FROM 
		permissions 
	LEFT JOIN (
		SELECT 
			users_permissions.permission_id as id 
		FROM 
			users_permissions 
		LEFT JOIN (
			SELECT 
				user_id 
			FROM 
				tokens 
			WHERE 
				hash = $1 
				AND scope = $2
		) as tokens ON users_permissions.user_id = tokens.user_id
	) as tmp ON permissions.id = tmp.id`

	rows, err := tx.QueryContext(ctx, query, token, models.ScopeActivation)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var permissions []string

	for rows.Next() {
		var permission string
		err = rows.Scan(
			&permission,
		)

		if err != nil {
			tx.Rollback()
			return nil, err
		}

		permissions = append(permissions, permission)
	}

	if err = rows.Err(); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return permissions, nil
}
