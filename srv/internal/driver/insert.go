package driver

import (
	"errors"

	"github.com/clarke94/weigh-me/srv/models"
)

// InsertUser will insert one user
func InsertUser(user models.User) (userID int64, err error) {
	db, err := Connect()
	if err != nil {
		return 0, errors.New("Unable to execute the query")
	}
	defer db.Close()

	sqlStatement := "INSERT INTO users (auth_id, first_name, last_name, email) VALUES ($1, $2, $3, $4) RETURNING id"

	var id int64

	err = db.QueryRow(sqlStatement, user.AuthID, user.FirstName, user.LastName, user.Email).Scan(&id)
	if err != nil {
		return 0, errors.New("Unable to execute the query")
	}

	return id, nil
}

// InsertWeight will insert one weight for a given user
func InsertWeight(weight models.Weight) (userID int64, err error) {
	db, err := Connect()
	if err != nil {
		return 0, errors.New("Unable to execute the query")
	}
	defer db.Close()

	sqlStatement := "INSERT INTO weights (user_id, name, value) VALUES ($1, $2, $3) RETURNING id"

	var id int64

	err = db.QueryRow(sqlStatement, weight.UserID, weight.Name, weight.Value).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
