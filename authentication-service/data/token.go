package data

import (
	"context"
	"log"
	"time"
)

// User is the structure which holds one user from the database.
type Token struct {
	ID        int       `json:"id"`
	Token     string    `json:"token"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"updated_at"`
}

// GetByToken returns one user by token
func (u *Token) GetByToken(tokenStr string) (*Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, token, userId, created_at, expired_at from token where token = $1`

	var token Token
	row := db.QueryRowContext(ctx, query, tokenStr)

	err := row.Scan(
		&token.ID,
		&token.Token,
		&token.UserId,
		&token.CreatedAt,
		&token.ExpiredAt,
	)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

// GetOne returns one user by id
func (u *Token) GetById(id int) (*Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, token, userId, created_at, expired_at from token where id = $1`

	var token Token
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&token.ID,
		&token.Token,
		&token.UserId,
		&token.CreatedAt,
		&token.ExpiredAt,
	)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

// GetAllByUser returns a slice of all token for an user
func (u *Token) GetAllByUser(userId int) ([]*Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, token, userId, created_at, expired_at from token where userId = $1`

	rows, err := db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokens []*Token

	for rows.Next() {
		var token Token
		err := rows.Scan(
			&token.ID,
			&token.Token,
			&token.UserId,
			&token.CreatedAt,
			&token.ExpiredAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		tokens = append(tokens, &token)
	}

	return tokens, nil
}

// Delete deletes one user from the database, by User.ID
func (u *Token) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from token where id = $1`

	_, err := db.ExecContext(ctx, stmt, u.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes one user from the database, by ID
func (u *Token) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from token where id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes one token from the database
func (u *Token) DeleteByToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from token where token = $1`

	_, err := db.ExecContext(ctx, stmt, token)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes one token from the database by userId
func (u *Token) DeleteByUserId(userId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from token where userId = $1`

	_, err := db.ExecContext(ctx, stmt, userId)
	if err != nil {
		return err
	}

	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (u *Token) Insert(token string, userId int, duration time.Duration) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int
	stmt := `insert into token (token, userId, created_at, expired_at)
		values ($1, $2, $3, $4) returning id`

	timeNow := time.Now()

	err := db.QueryRowContext(ctx, stmt,
		token,
		userId,
		timeNow,
		timeNow.Add(duration),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}
