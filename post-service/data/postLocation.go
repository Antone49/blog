package data

import (
	"context"
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