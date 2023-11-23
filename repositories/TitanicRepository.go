package repositories

import (
	"fmt"
	"log"
	"parquet_1/infrastructures"
	"parquet_1/interfaces"
	"parquet_1/models"
)

func NewTitanicRepository() *TitanicRepository {
	return &TitanicRepository{infrastructures.NewMySQLHandler()}
}

type TitanicRepository struct {
	DBHandler interfaces.IDbHandler
}

func (repo *TitanicRepository) TruncateRecords() {
	repo.DBHandler.Execute("TRUNCATE TABLE titanic")
}

func (repo *TitanicRepository) AddPassenger(passenger models.TitanicPassenger) (int64, error) {

	//url := article.Url[:min(250, len(article.Url))]
	//title := article.Title[:min(250, len(article.Title))]

	Result, err := repo.DBHandler.Execute("insert into titanic (passenger_id, survived, pclass, name, sex, age, sib_sp, parch, ticket, fare, cabin, embarked) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
		passenger.Id,
		passenger.Survived,
		passenger.PClass,
		passenger.Name,
		passenger.Sex,
		passenger.Age,
		passenger.SibSp,
		passenger.Parch,
		passenger.Ticket,
		passenger.Fare,
		passenger.Cabin,
		passenger.Embarked,
	)

	if err != nil {
		log.Fatal(err)

		return 0, err
	}

	lastInsertId, err := Result.LastInsertId()
	if err != nil {
		log.Fatal(err)

		return 0, err
	}

	return lastInsertId, nil
}

func (repo *TitanicRepository) AddPassengersBatch(data []models.TitanicPassenger) {
	for _, el := range data {
		repo.AddPassenger(el)
	}
}

func (repo *TitanicRepository) GetRowsCount() int {
	var numRows int = 0
	row := repo.DBHandler.QueryRow("SELECT count(1) FROM titanic WHERE 1")
	errScan := row.Scan(&numRows)
	if errScan != nil {
		return 0
	}

	return numRows
}

func (repo *TitanicRepository) GetPassengersBatch(start, limit int) []models.TitanicPassenger {
	rows, err := repo.DBHandler.Query("SELECT passenger_id, survived, pclass, name, sex, age, sib_sp, parch, ticket, fare, cabin, embarked FROM titanic LIMIT ?, ?", start, limit)
	if err != nil {
		fmt.Errorf("Error GetPassengersBatch %s", err)
	}

	passengers := []models.TitanicPassenger{}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var passenger models.TitanicPassenger
		if err := rows.Scan(&passenger.Id, &passenger.Survived, &passenger.PClass, &passenger.Name, &passenger.Sex, &passenger.Age,
			&passenger.SibSp, &passenger.Parch, &passenger.Ticket, &passenger.Fare, &passenger.Cabin, &passenger.Embarked); err != nil {

			fmt.Errorf("GetPassengersBatch failed %v", err)
			return []models.TitanicPassenger{}
		}
		passengers = append(passengers, passenger)
	}

	return passengers
}
