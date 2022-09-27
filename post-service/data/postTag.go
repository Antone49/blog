package data

import (
	"context"
	"log"
)

// PostTag is the structure which holds one post from the database.
type PostTag struct {
	PostId int      `json:"postId"`
	TagId  int      `json:"tagId"`
}

// GetAll returns a slice of all PostTag, sorted by created
func (p *PostTag) GetAll(search string) ([]*PostTag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Default query all posts
	query := `select postId, tagId from post`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var postTags []*PostTag

	for rows.Next() {
		var postTag PostTag
		err := rows.Scan(
			&postTag.PostId,
			&postTag.TagId,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		postTags = append(postTags, &postTag)
	}

	return postTags, nil
}

func (t *PostTag) RemoveByTagId(tagId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from postTag where tagId = $1`

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

func (t *PostTag) RemoveByPostId(postId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from PostTag where postId = $1`

	_, err := db.ExecContext(ctx, stmt, postId)

	if err != nil {
		return err
	}

	return nil
}

// GetTagsFromPostId returns the tag post
func (p *PostTag) GetTagsFromPostId(idPost int) ([]*Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name 
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

func (p *PostTag) InsertFromPostId(idPost int, idTag []int) (error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	for _, element := range idTag {
		stmt := `insert into PostTag (postId, tagId)
			values ($1, $2)`

		_, err := db.ExecContext(ctx, stmt, idPost, element)
		if err != nil {
			return err
		}
	}

	return nil
}

