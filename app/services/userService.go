package services

import (
	models "api/app/models/user"
	"api/app/repositories"
	"api/config"
	"log"
)

func CreateUser(userRequest models.UserCreateRequest) (models.UserModel, error) {

	db, err := config.ConnectMysql()
	if err != nil {
		log.Println("error to connect mysql - ", err)
		return models.UserModel{}, err
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	userID, err := repo.Create(userRequest)
	if err != nil {
		return models.UserModel{}, err
	}

	createdUser, err := repo.GetById(uint64(userID))
	if err != nil {
		return models.UserModel{}, err
	}

	return createdUser, nil
}

func GetUserById(userID uint64) (models.UserModel, error) {

	db, err := config.ConnectMysql()
	if err != nil {
		log.Println("error to connect mysql - ", err)
		return models.UserModel{}, err
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	user, err := repo.GetById(userID)
	if err != nil {
		return models.UserModel{}, err
	}

	return user, nil
}
