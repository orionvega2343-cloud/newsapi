package internal

import (
	"github.com/jmoiron/sqlx"
)

func NewDB(connStr string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Repository struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateUser(u User) error {
	_, err := r.DB.Exec(`INSERT INTO users (email,password_hash) VALUES ($1, $2)`, u.Email, u.PasswordHash)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	var user User
	err := r.DB.Get(&user, `SELECT * FROM users WHERE email = $1`, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) SaveArticle(a Article) error {
	_, err := r.DB.Exec(`INSERT INTO articles (title,link,source) VALUES ($1, $2, $3)`, a.Title, a.Link, a.Source)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetArticles() ([]Article, error) {
	var articles []Article
	err := r.DB.Select(&articles, `SELECT * FROM articles`)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
