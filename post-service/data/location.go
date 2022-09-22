package data

import (
	"context"
	"log"
)

type Location struct {
	ID 			string 	`json:"id"`
	Title 		string 	`json:"title"`
	Latitude   	float64	`json:"latitude"`
	Longitude  	float64	`json:"longitude"`
	Image      	string	`json:"image"`
}

// GetAll returns a slice of all post locations, sorted by created
func (l *Location) GetAll() ([]*Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select postId, latitude, longitude, name, image
	from location`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []*Location

	for rows.Next() {
		var location Location
		err := rows.Scan(
			&location.ID,
			&location.Latitude,
			&location.Longitude,
			&location.Title,
			&location.Image,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		locations = append(locations, &location)
	}

	return locations, nil
}

