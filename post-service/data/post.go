package data

import (
	"context"
	"strconv"
	"log"
	"time"
)

// Post is the structure which holds one post from the database.
type Post struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	ThumbnailImage string    `json:"thumbnailImage"`
	Image          string    `json:"image"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GetAllPosts returns a slice of all posts, sorted by created
func (p *Post) GetAll(search string) ([]*Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Default query all posts
	query := `select id, title, thumbnailImage, content, created_at, updated_at
	from post order by created_at`

	if len(search) > 0 {
		query = `SELECT DISTINCT post.id, post.title, post.thumbnailImage, post.content, post.created_at, post.updated_at
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
			&post.ID,
			&post.Title,
			&post.ThumbnailImage,
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

	query := `select id, title, thumbnailImage, content, created_at, updated_at
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
			&post.ID,
			&post.Title,
			&post.ThumbnailImage,
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

	var post Post
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&post.ID,
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

// GetTags returns the tag post
func (p *Post) GetTags(id int) ([]*Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select tagId from postTag where postId = $1`

	rows, err := db.QueryContext(ctx, query, id)
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

// Update updates one post in the database, using the information
// stored in the receiver u
func (p *Post) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update post set
		title = $1,
		thumbnail = $2,
		content = $3,
		updated_at = $4
		where id = $5
	`

	_, err := db.ExecContext(ctx, stmt,
		p.Title,
		p.ThumbnailImage,
		p.Content,
		time.Now(),
		p.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes one post from the database, by Post.ID
func (p *Post) DeleteByID() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from post where id = $1`

	_, err := db.ExecContext(ctx, stmt, p.ID)
	if err != nil {
		return err
	}

	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (p *Post) Insert(user Post) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int
	stmt := `insert into post (title, thumbnail, content, created_at, updated_at)
		values ($1, $2, $3, $4, $5) returning id`

	err := db.QueryRowContext(ctx, stmt,
		user.Title,
		user.ThumbnailImage,
		user.Content,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}
