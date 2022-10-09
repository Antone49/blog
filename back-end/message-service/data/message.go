package data

import (
	"context"
	"log"
	"time"
	"strconv"
)

type Message struct {
	Id 				int 		`json:"id"`
	Message 		string 		`json:"message"`
	Username 		string 		`json:"username"`
	PostId   		int			`json:"postId"`
	Valid   		bool		`json:"valid"`
	CreatedAt      	time.Time 	`json:"created_at",omitempty`
	UpdatedAt      	time.Time 	`json:"updated_at",omitempty`
}

// GetAll returns a slice of all post locations, sorted by created
func (l *Message) GetAll() ([]*Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, message, username, postId, valid, created_at, update_at
	from message`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*Message

	for rows.Next() {
		var message Message
		err := rows.Scan(
			&message.Id,
			&message.Message,
			&message.Username,
			&message.PostId,
			&message.Valid,
			&message.CreatedAt,
			&message.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		messages = append(messages, &message)
	}

	return messages, nil
}

func (p *Message) GetLastest(latest int) ([]*Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, message, username, postId, valid, created_at, update_at
	from message order by created_at limit ` + strconv.Itoa(latest)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*Message

	for rows.Next() {
		var message Message
		err := rows.Scan(
			&message.Id,
			&message.Message,
			&message.Username,
			&message.PostId,
			&message.Valid,
			&message.CreatedAt,
			&message.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		messages = append(messages, &message)
	}

	return messages, nil
}

// Get returns one post by id
func (l *Message) GetAllFromPostId(postId int) ([]*Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select message, username, created_at from message where postId = $1 and valid=true`

	rows, err := db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*Message

	for rows.Next() {
		var message Message
		err := rows.Scan(
			&message.Message,
			&message.Username,
			&message.CreatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		messages = append(messages, &message)
	}

	return messages, nil
}

func (l *Message) Insert(message Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `insert into message (message, username, postId, valid, created_at, update_at)
		values ($1, $2, $3, $4, $5, $6)`

	_, err := db.ExecContext(ctx, stmt, message.Message, message.Username, message.PostId, true, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (l *Message) Update(message Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update message set
		message = $1, 
		username = $2, 
		postId = $3, 
		created_at = $4,
		update_at = $5 
		where id = $6
	`
	_, err := db.ExecContext(ctx, stmt, message.Message, message.Username, message.PostId, message.CreatedAt, message.UpdatedAt, message.Id)

	if err != nil {
		return err
	}

	return nil
}

func (l *Message) UpdateValidity(message Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update message set
		valid = $1, 
		update_at = $2 
		where id = $3
	`
	_, err := db.ExecContext(ctx, stmt, message.Valid, time.Now(), message.Id)

	if err != nil {
		return err
	}

	return nil
}


func (t *Message) Delete(message Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from message where id = $1`

	_, err := db.ExecContext(ctx, stmt, message.Id)

	if err != nil {
		return err
	}

	return nil
}

