package database

import (
	"database/sql"

	models "github.com/anilkusc/BullsAndCows/models"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	*models.User
}

func (u *User) ReadUser(db *sql.DB, id int) (models.User, error) {
	//	var query string
	var user models.User
	/*
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
	*/
	return user, nil

}
