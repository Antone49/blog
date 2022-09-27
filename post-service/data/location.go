package data

import (
	"context"
	"log"
)

type Location struct {
	Id 			int 	`json:"id"`
	Name 		string 	`json:"name"`
	Latitude   	float64	`json:"latitude"`
	Longitude  	float64	`json:"longitude"`
	Image      	string	`json:"image"`
}

// GetAll returns a slice of all post locations, sorted by created
func (l *Location) GetAll() ([]*Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, latitude, longitude, name, image
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
			&location.Id,
			&location.Latitude,
			&location.Longitude,
			&location.Name,
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

// Get returns one post by id
func (l *Location) Get(id int) (*Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, latitude, longitude, name, image from location where id = $1`

	var location Location
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&location.Id,
		&location.Latitude,
		&location.Longitude,
		&location.Name,
		&location.Image,
	)

	if err != nil {
		return nil, err
	}

	return &location, nil
}

func (l *Location) Insert(location Location) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `insert into location (name, latitude, longitude, image)
		values ($1, $2, $3, $4)`

	_, err := db.ExecContext(ctx, stmt, location.Name, location.Latitude, location.Longitude, location.Image)

	if err != nil {
		return err
	}

	return nil
}

func (l *Location) Update(location Location) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update location set
		name = $1, 
		latitude = $2, 
		longitude = $3, 
		image = $4 
		where id = $5
	`
	_, err := db.ExecContext(ctx, stmt, location.Name, location.Latitude, location.Longitude, location.Image, location.Id)

	if err != nil {
		return err
	}

	return nil
}

func (t *Location) Remove(location Location) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from location where id = $1`

	_, err := db.ExecContext(ctx, stmt, location.Id)

	if err != nil {
		return err
	}

	return nil
}

