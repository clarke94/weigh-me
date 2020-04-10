package driver

import (
	"github.com/clarke94/weigh-me/srv/models"
)

// GetUserByAuthID driver that connects to the Database and returns a user
func GetUserByAuthID(id string) (models.User, error) {
	// create a user of models.User type
	var user models.User

	db, err := Connect()
	if err != nil {
		return user, err
	}
	defer db.Close()

	// create the select sql query
	sqlStatement := "SELECT id, auth_id, first_name, last_name, email FROM users WHERE auth_id=$1"

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err = row.Scan(&user.ID, &user.AuthID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}
