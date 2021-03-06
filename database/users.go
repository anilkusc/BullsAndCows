package database

import (
	"database/sql"
	"strconv"

	models "github.com/anilkusc/BullsAndCows/models"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	*models.User
}

func (u *User) CreateUser(db *sql.DB, user models.User) (models.User, error) {

	statement, err := db.Prepare("INSERT INTO Users (Id,Name) VALUES(?,?)")
	if err != nil {
		return user, err
	}
	statement.Exec(user.Id, user.Name)
	statement.Close()
	return user, nil

}

func (u *User) ReadUser(db *sql.DB, id int) (models.User, error) {
	var query string
	var user models.User

	query = "SELECT * FROM Users where Id=" + strconv.Itoa(id)
	rows, err := db.Query(query)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return user, err
		}
	}
	if err = rows.Err(); err != nil {
		return user, err
	}
	return user, nil

}
func (u *User) UpdateUser(db *sql.DB, user models.User) (models.User, error) {

	statement, err := db.Prepare("UPDATE Users SET Id=?,Name=?")
	if err != nil {
		return user, err
	}
	statement.Exec(user.Id, user.Name)
	statement.Close()

	return user, nil

}

func (u *User) DeleteUser(db *sql.DB, id int) (models.User, error) {
	user, err := u.ReadUser(db, id)
	if err != nil {
		return user, err
	}

	statement, err := db.Prepare("DELETE FROM Users where Id=?")
	if err != nil {
		return user, err
	}
	statement.Exec(user.Id)
	statement.Close()

	return user, nil

}
func (u *User) ListUsers(db *sql.DB) ([]models.User, error) {

	var users []models.User

	query := "SELECT * FROM Users"
	rows, err := db.Query(query)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil

}
