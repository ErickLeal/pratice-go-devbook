package services

import (
	models "api/app/models/user"
	"api/app/repositories"
	"api/config"
	"log"
)

func CreateUser(userRequest models.UserCreateRequest) (models.UserModel, error) {
	println("CRETETETEEEE")
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println("error to connect database - ", err)
		return models.UserModel{}, err
	}

	repo := repositories.NewUserRepository(db)
	userID, err := repo.CreateUser(userRequest)
	if err != nil {
		return models.UserModel{}, err
	}
	println("userID")
	println(userID)
	createdUser, err := repo.GetUserById(uint64(userID))
	if err != nil {
		return models.UserModel{}, err
	}

	return createdUser, nil
}

func GetUserById(userID uint64) (models.UserModel, error) {

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println("error to connect database - ", err)
		return models.UserModel{}, err
	}

	repo := repositories.NewUserRepository(db)

	user, err := repo.GetUserById(userID)
	if err != nil {
		return models.UserModel{}, err
	}

	return user, nil
}

func UpdateUser(userID uint64, userRequest models.UserUpdateRequest) (models.UserModel, error) {

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println("error to connect database - ", err)
		return models.UserModel{}, err
	}

	repo := repositories.NewUserRepository(db)

	err = repo.UpdateUser(userID, userRequest)
	if err != nil {
		return models.UserModel{}, err
	}

	user, err := repo.GetUserById(userID)
	if err != nil {
		return models.UserModel{}, err
	}

	return user, nil
}

func DeleteUser(userID uint64) error {

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println("error to connect database - ", err)
		return err
	}

	repo := repositories.NewUserRepository(db)

	err = repo.DeleteUser(userID)
	if err != nil {
		return err
	}

	return nil
}

func ListAllUsers() ([]models.UserModel, error) {

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println("error to connect database - ", err)
		return []models.UserModel{}, err
	}

	repo := repositories.NewUserRepository(db)

	users, err := repo.GetAllUsers()
	if err != nil {
		return []models.UserModel{}, err
	}

	return users, nil
}
