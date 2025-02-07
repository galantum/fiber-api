package repositories

import (
	"context"
	"fiber-api/infrastructure"
	"fiber-api/models"

	"github.com/gofiber/fiber/v2/log"
)

type UserRepository interface {
	GetUsers() ([]models.User, error)
	GetUser(id int) (models.User, error)

	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	DeleteUser(user models.User) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUsers() ([]models.User, error) {

	rows, err := infrastructure.DB.Query(context.Background(), "SELECT id, user_name, user_email FROM account.users")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name, &user.Email)
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) GetUser(id int) (models.User, error) {
	var user models.User
	err := infrastructure.DB.QueryRow(context.Background(), "SELECT id, user_name, user_email FROM account.users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		log.Error(err)
		return user, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user models.User) error {
	_, err := infrastructure.DB.Exec(context.Background(), "INSERT INTO account.users (user_name, user_email) VALUES ($1, $2)", user.Name, user.Email)

	return err
}

func (r *userRepository) UpdateUser(user models.User) error {

	_, err := infrastructure.DB.Exec(context.Background(), "UPDATE account.users SET user_name = $1, user_email = $2 WHERE id = $3", user.Name, user.Email, user.ID)

	return err
}

func (r *userRepository) DeleteUser(user models.User) error {

	_, err := infrastructure.DB.Exec(context.Background(), "DELETE FROM account.users WHERE id = $1", user.ID)
	return err
}
