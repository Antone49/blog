package data

import (
	"context"
	"log"
	"strconv"
	"time"
)

// Post is the structure which holds one post from the database.
type Post struct {
	Id             int       `json:"id"`
	Title          string    `json:"title",omitempty`
	Image          string    `json:"image",omitempty`
	Content        string    `json:"content",omitempty`
	CreatedAt      time.Time `json:"created_at",omitempty`
	UpdatedAt      time.Time `json:"updated_at",omitempty`
}

// GetAllPosts returns a slice of all posts, sorted by created
func (p *Post) GetAll(search string) ([]*Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Default query all posts
	query := `select id, title, image, content, created_at, updated_at
	from post order by created_at`

	if len(search) > 0 {
		query = `SELECT DISTINCT post.id, post.title, post.image, post.content, post.created_at, post.updated_at
					FROM post
				LEFT JOIN PostTag
					ON PostTag.postId = post.Id
				LEFT JOIN Tag
					ON PostTag.tagId = Tag.Id
				WHERE Tag.Id = '` + search + `'
					OR post.title like '%` + search + `%'
					OR post.content like '%` + search + `%';`
	}

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post

	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Image,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		posts = append(posts, &post)
	}

	return posts, nil
}

// GetLastest returns a slice of latest posts, sorted by created
func (p *Post) GetLastest(latest int) ([]*Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, title, image, content, created_at, updated_at
	from post order by created_at limit ` + strconv.Itoa(latest)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post

	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Image,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		posts = append(posts, &post)
	}

	return posts, nil
}

// Get returns one post by id
func (p *Post) Get(id int) (*Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, title, image, content, created_at, updated_at from post where id = $1`

	row := db.QueryRowContext(ctx, query, id)

	var post Post
	err := row.Scan(
		&post.Id,
		&post.Title,
		&post.Image,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

// Update updates one post in the database, using the information
// stored in the receiver u
func (p *Post) Update(post Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update post set
		title = $1,
		content = $2,
		updated_at = $3
		where id = $4
	`

	_, err := db.ExecContext(ctx, stmt,
		post.Title,
		post.Content,
		time.Now(),
		post.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

// Update updates one post in the database, using the information
// stored in the receiver u
func (p *Post) UpdateImage(post Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update post set
		image = $1,
		updated_at = $2
		where id = $3
	`

	_, err := db.ExecContext(ctx, stmt,
		post.Image,
		time.Now(),
		post.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes one post from the database, by Post.ID
func (p *Post) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from post where id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (p *Post) Insert() (*Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `insert into post (title, image, content, created_at, updated_at)
		values ($1, $2, $3, $4, $5) RETURNING id;`

	row := db.QueryRowContext(ctx, query,
		"",
		"",
		"",
		time.Now(),
		time.Now(),
	)

	var post Post
	err := row.Scan(
		&post.Id,
	)

	if err != nil {
		return nil, err
	}

	return &post, nil
}
