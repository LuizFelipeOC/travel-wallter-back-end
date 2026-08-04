package auth

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	CreateUser(name, email, password string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	VerifyPassword(hashedPassword, password string) bool
}

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(name, email, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO usuarios (nome, email, senha, criado_em, atualizado_em)
	          VALUES (:1, :2, :3, SYSTIMESTAMP, SYSTIMESTAMP)`

	_, err = r.db.Exec(query, name, email, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	user, err := r.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *AuthRepository) GetUserByEmail(email string) (*User, error) {
	query := `SELECT id, nome, email, senha, criado_em, atualizado_em FROM usuarios WHERE email = :1`

	var user User
	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) GetUserByID(id int) (*User, error) {
	query := `SELECT id, nome, email, senha, criado_em, atualizado_em FROM usuarios WHERE id = :1`

	var user User
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	user.Password = ""
	return &user, nil
}

func (r *AuthRepository) VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
