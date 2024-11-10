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

	stmt := `INSERT INTO users (name, email, password, activated, created_at,
		updated_at)
		values ($1, $2, $3, $4, $5, $6) returning id`

	var newID int

	err := m.DB.QueryRowContext(ctx, stmt,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		false,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (m *PostgresDBRepo) GetUserByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT 
				id, name, email, password, activated, created_at, updated_at, version 
			  FROM
			  	users
			  WHERE
				name = $1`

	var u models.User

	err := m.DB.QueryRowContext(ctx, query, username).Scan(
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
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, models.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &u, err
}

func (m *PostgresDBRepo) GetUserByID(id int64) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT 
				id, name, email, password, activated, created_at, updated_at, version
			  FROM 
			  	users
			  WHERE
			   id = $1`

	var u models.User
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
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
		return nil, err
	}

	return &u, err
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

	tokenHash := sha256.Sum256(token.Hash)

	query := `INSERT INTO tokens (hash, user_id, expiry, scope)
			  VALUES ($1, $2, $3, $4);`

	args := []interface{}{tokenHash[:], token.UserID, token.Expiry, token.Scope}

	_, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return err
}

func (m *PostgresDBRepo) DeleteAllTokensForUser(scope string, userID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM tokens
			  WHERE scope = $1 AND user_id = $2`

	_, err := m.DB.ExecContext(ctx, query, scope, userID)

	return err
}

func (m *PostgresDBRepo) GetUserByToken(tokenScope, tokenPlainText string) (*models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

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

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Activated,
		&user.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, models.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m *PostgresDBRepo) NewArticle(article *models.Article) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO articles (name, publish, reading_time, username, html_list)
	VALUES ($1, $2, $3, $4, $5);`

	_, err := m.DB.ExecContext(ctx, query, article.Name, article.Publish, article.ReadingTime, article.Username, article.HtmlList)

	if err != nil {
		return err
	}

	return nil
}

func (m *PostgresDBRepo) GetArticle(id int) (*models.Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT name, publish, reading_time, username, html_list, created_at, updated_at, version 
				FROM articles
				WHERE id = $1`

	var article models.Article

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
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
		return nil, err
	}

	return &article, nil
}

func (m *PostgresDBRepo) GetNews(page int) (*[]models.News, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, name, created_at FROM articles WHERE publish=true ORDER BY created_at, id LIMIT $1 OFFSET $2`

	var newsArr []models.News

	limit := 9
	offset := (page - 1) * 9

	rows, err := m.DB.QueryContext(ctx, query, limit, offset)
	if err != nil {
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
			return nil, err
		}

		newsArr = append(newsArr, news)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &newsArr, err
}

func (m *PostgresDBRepo) GetUserPermissions(token string) (*[]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
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
	
	rows, err := m.DB.QueryContext(ctx, query, token, models.ScopeActivation)
	if err != nil {
		return nil, err
	}

	var permissions []string

	for rows.Next() {
		var permission string
		err = rows.Scan(
			&permission,
		)

		if err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return &permissions, nil
}
