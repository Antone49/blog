package data

import (
	"context"
	"log"
)

type PostLocation struct {
	PostId	int       
	LocationId	int
}

func (t *PostLocation) RemoveByPostId(postId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from postLocation where postId = $1`

	_, err := db.ExecContext(ctx, stmt, postId)

	if err != nil {
		return err
	}

	return nil
}

func (t *PostLocation) RemoveByLocationId(locationId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from postLocation where locationId = $1`

	_, err := db.ExecContext(ctx, stmt, locationId)

	if err != nil {
		return err
	}

	return nil
}

// GetTagsFromPostId returns the tag post
func (p *PostLocation) GetLocationsFromPostId(idPost int) ([]*Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, latitude, longitude, name, image
				from Location 
			LEFT JOIN PostLocation
				ON PostLocation.postId = $1
			WHERE PostLocation.locationId = Location.id`

	rows, err := db.QueryContext(ctx, query, idPost)
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

func (p *PostLocation) InsertFromPostId(idPost int, idLocation []int) (error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	for _, element := range idLocation {
		stmt := `insert into PostLocation (postId, locationId)
			values ($1, $2)`

		_, err := db.ExecContext(ctx, stmt, idPost, element)
		if err != nil {
			return err
		}
	}

	return nil
}

