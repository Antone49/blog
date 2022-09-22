package data

import (
	"context"
	"log"
)

// PostTag is the structure which holds one post from the database.
type PostTag struct {
	PostId	int       
	TagId	int
}

// GetTagsFromPostId returns the tag post
func (p *PostTag) GetTagsFromPostId(idPost int) ([]*Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select name 
				from Tag 
			LEFT JOIN PostTag
				ON PostTag.postId = $1
			WHERE PostTag.tagId = Tag.id`

	rows, err := db.QueryContext(ctx, query, idPost)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []*Tag

	for rows.Next() {
		var tag Tag
		err := rows.Scan(
			&tag.ID,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		tags = append(tags, &tag)
	}

	return tags, nil
}

func (t *PostTag) RemoveByTagId(tagId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from PostTag where tagId = $1`

	_, err := db.ExecContext(ctx, stmt, tagId)

	if err != nil {
		return err
	}

	return nil
}

func (t *PostTag) RemoveByTagName(tagName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from PostTag using PostTag as PT 
			LEFT JOIN Tag as T
				ON T.name = $1 
	 		WHERE PostTag.tagId = T.id`

	_, err := db.ExecContext(ctx, stmt, tagName)

	if err != nil {
		return err
	}

	return nil
}