package repositories

import (
	models "api/app/models/user"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repo Users) Create(user models.UserCreateRequest) (uint64, error) {
	statement, err := repo.db.Prepare(
		"INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (repo Users) GetById(userID uint64) (models.UserModel, error) {
	var user models.UserModel
	err := repo.db.QueryRow(
		"SELECT id, name, nick, email, created_at FROM users WHERE id = ?",
		userID,
	).Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)

	if err != nil {
		return models.UserModel{}, err
	}

	return user, nil
}
