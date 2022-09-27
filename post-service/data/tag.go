package data

import (
	"context"
	"log"
)

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// GetAllTags returns a slice of all tags
func (t *Tag) GetAll() ([]*Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name from tag`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []*Tag

	for rows.Next() {
		var tag Tag
		err := rows.Scan(
			&tag.Id,
			&tag.Name,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		tags = append(tags, &tag)
	}

	return tags, nil
}

// Insert inserts a new tag into the database
func (t *Tag) Insert(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `insert into tag (name)
		values ($1)`

	_, err := db.ExecContext(ctx, stmt, name)

	if err != nil {
		return err
	}

	return nil
}

func (t *Tag) Update(id int, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update tag set
		name = $1 
		where id = $2
	`
	_, err := db.ExecContext(ctx, stmt, name, id)

	if err != nil {
		return err
	}

	return nil
}

func (t *Tag) Remove(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from tag where id = $1`

	_, err := db.ExecContext(ctx, stmt, id)

	if err != nil {
		return err
	}

	return nil
}

func (t *Tag) RemoveByName(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from tag where name = $1`

	_, err := db.ExecContext(ctx, stmt, name)

	if err != nil {
		return err
	}

	return nil
}
