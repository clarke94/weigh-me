package driver

import (
	"errors"
)

type Weight struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

// GetAllWeights will get all the weights of a given user
func GetAllWeights() ([]Weight, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var weights []Weight
	sqlStatement := "SELECT id, name, value FROM weights"

	// execute the sql statement
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, errors.New("Unable to execute the query")
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var weight Weight

		// unmarshal the row object to user
		err = rows.Scan(&weight.ID, &weight.Name, &weight.Value)

		if err != nil {
			return nil, errors.New("Unable to execute the row scan")
		}

		// append the user in the users slice
		weights = append(weights, weight)
	}

	return weights, nil
}

// GetAllWeightsByUserID will return all weight values for that user
func GetAllWeightsByUserID(id int64) ([]Weight, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var weights []Weight
	sqlStatement := "SELECT id, name, value FROM weights WHERE user_id=$1"

	// execute the sql statement
	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		return nil, errors.New("Unable to execute the query")
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var weight Weight

		// unmarshal the row object to user
		err = rows.Scan(&weight.ID, &weight.Name, &weight.Value)

		if err != nil {
			return nil, errors.New("No rows were returned")
		}

		// append the user in the users slice
		weights = append(weights, weight)
	}

	return weights, nil
}
