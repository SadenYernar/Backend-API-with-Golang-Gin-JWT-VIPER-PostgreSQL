package repository

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"project/internal/model"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (u *UserStorage) CreateUser(user model.User) (int, model.User, error) {
	records := `INSERT INTO users(uuid,name,username,email,password) VALUES ($1,$2,$3,$4,$5)`
	fmt.Println(user)

	query, err := u.db.Prepare(records)
	if err != nil {
		return http.StatusInternalServerError, user, fmt.Errorf("Error in CreateUser method in repository: %w", err)
	}
	_, err = query.Exec(user.Uuid, user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		fmt.Print(err)
		return http.StatusInternalServerError, user, fmt.Errorf("Error in CreateUser method in repository: %w", err)
	}

	fmt.Println("There")
	fmt.Println("User created successfully!")
	return http.StatusOK, user, err
}

func (u *UserStorage) GetUserInfoByEmail(user model.User) (model.User, error) {
	row := u.db.QueryRow("SELECT uuid,username,email,password FROM users WHERE email=$1", user.Email)
	temp := model.User{}
	err := row.Scan(&temp.Uuid, &temp.Username, &temp.Email, &temp.Password)
	if err != nil {
		log.Printf("Error with GetUserInfo in repository: %v\n", err)
		return model.User{}, err
	}
	return temp, nil
}

func (u *UserStorage) SetSession(user model.User, token string) error {
	records := `UPDATE users SET token=$1 WHERE uuid=$2`
	query, err := u.db.Prepare(records)
	if err != nil {
		return fmt.Errorf("Error in SetSession method in repository: %w", err)
	}
	_, err = query.Exec(token, user.Uuid)
	if err != nil {
		return fmt.Errorf("Error in SetSession method in repository: %w", err)
	}
	fmt.Println("Session created successfully!")
	return nil
}

func (u *UserStorage) GetUserInfoByUsername(username string) (model.User, error) {
	row := u.db.QueryRow("SELECT uuid,email FROM users WHERE username=$1", username)
	temp := model.User{}
	err := row.Scan(&temp.Uuid, &temp.Email)
	if err != nil {
		log.Printf("Error with GetUserInfo in repository: %v\n", err)
		return model.User{}, err
	}
	return temp, nil
}

// func (u *UserStorage) CheckUserByEmail(user model.User) bool {
// 	stmt := `SELECT email FROM users WHERE email == $1`
// 	query, err := u.db.Prepare(stmt)
// 	if err != nil {
// 		return false, err
// 	}
// 	row := query.QueryRow(email)
// 	var mail string
// 	err = row.Scan(&mail)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return false, nil
// 		}
// 		return false, err
// 	}
// 	return true, nil
// }
