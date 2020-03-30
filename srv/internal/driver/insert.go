package driver

import (
	"errors"

	"github.com/clarke94/weigh-me/srv/models"
)

// InsertUser will insert one user
func InsertUser(user models.User) (returnedUser models.User, err error) {
	db, err := Connect()
	if err != nil {
		return returnedUser, errors.New("Unable to execute the query")
	}
	defer db.Close()

	sqlStatement := "INSERT INTO users (auth_id, first_name, last_name, email) VALUES ($1, $2, $3, $4) RETURNING id, auth_id, first_name, last_name, email"

	err = db.QueryRow(sqlStatement, user.AuthID, user.FirstName, user.LastName, user.Email).Scan(&returnedUser.ID, &returnedUser.AuthID, &returnedUser.FirstName, &returnedUser.LastName, &returnedUser.Email)
	if err != nil {
		return returnedUser, errors.New("Unable to execute the query")
	}

	return returnedUser, nil
}

// InsertWeight will insert one weight for a given user
func InsertWeight(weight models.Weight) (returnedWeight models.Weight, err error) {
	db, err := Connect()
	if err != nil {
		return returnedWeight, errors.New("Unable to execute the query")
	}
	defer db.Close()

	sqlStatement := "INSERT INTO weights (user_id, name, value) VALUES ($1, $2, $3) RETURNING id, user_id, name, value"

	err = db.QueryRow(sqlStatement, weight.UserID, weight.Name, weight.Value).Scan(&returnedWeight.ID, &returnedWeight.UserID, &returnedWeight.Name, &returnedWeight.Value)
	if err != nil {
		return returnedWeight, err
	}

	return returnedWeight, nil
}
