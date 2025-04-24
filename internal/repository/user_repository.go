package repository

import (
	"database/sql"
	"simple-crud/internal/models"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateNewUser(user *models.User) error {
	query := `
        INSERT INTO users (username, email, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, updated_at
    `
	return r.db.QueryRow(query, user.Username, user.Email, time.Now(), time.Now()).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}


func (r *UserRepository) GetUserById(id int64) (*models.User, error) {
	query := `
		SELECT id, username, email, created_at, updated_at
		FROM users
		WHERE id = $1;
	`

	var user models.User
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}


func (r *UserRepository) UpdateUser(user *models.User) error {
	query := `
        UPDATE users
        SET username = $1, email = $2, updated_at = $3
        WHERE id = $4
    `

	_, err := r.db.Exec(query, user.Username, user.Email, time.Now(), user.ID)
	return err
}

func (r *UserRepository) DeleteUser(id int64) error {
	query := `
        DELETE FROM users
        WHERE id = $1
    `
	_, err := r.db.Exec(query, id)
	return err
}


func (r *UserRepository) GetAllUsers() ([]*models.User, error) {
	query := `
        SELECT id, username, email, created_at, updated_at
        FROM users
    `

	rows, err := r.db.Query(query)
	if err!= nil {
		return nil, err
	}

	defer rows.Close()
	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err!= nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}