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

func (repo Users) CreateUser(user models.UserCreateRequest) (uint64, error) {
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

func (repo Users) UpdateUser(userID uint64, user models.UserUpdateRequest) error {
	statement, err := repo.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, userID); err != nil {
		return err
	}

	return nil
}

func (repo Users) GetUserById(userID uint64) (models.UserModel, error) {
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

func (repo Users) DeleteUser(userID uint64) error {
	statement, err := repo.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID); err != nil {
		return err
	}

	return nil
}

func (repo Users) GetAllUsers() ([]models.UserModel, error) {
	lines, err := repo.db.Query(
		"select id, name, nick, email, created_at from users",
	)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var result []models.UserModel

	for lines.Next() {
		var currentUser models.UserModel

		if err = lines.Scan(
			&currentUser.ID,
			&currentUser.Name,
			&currentUser.Nick,
			&currentUser.Email,
			&currentUser.CreatedAt,
		); err != nil {
			return nil, err
		}

		result = append(result, currentUser)
	}

	return result, nil
}
