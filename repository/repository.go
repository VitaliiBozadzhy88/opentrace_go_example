package repository

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/opentracing/opentracing-go"
	"traceWithGoV1/model"
)

const dburl = "root:vitalik88@tcp(localhost:3306)/usersdb"

// Repository retrieves information about users.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Repository
func NewRepository() *Repository {
	db, err := sql.Open("mysql", dburl)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot ping the db: %v", err)
	}
	return &Repository{
		db: db,
	}
}

// GetDataByEmail tries to find the person in the database by email.
func (r *Repository) GetDataByEmail(ctx context.Context, email string) (model.User, error) {
	query := "select email, name, activation_code from usersdb.Users where email = ?"

	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"get-person",
		opentracing.Tag{
			Key: "db.statement", Value: query},
		opentracing.Tag{
			Key: "email:", Value: email,
		},
	)
	defer span.Finish()

	rows, err := r.db.QueryContext(ctx, query, email)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		var name string
		var code string

		err := rows.Scan(&email, &name, &code)
		if err != nil {
			return model.User{}, err
		}
		return model.User{
			Name:           name,
			ActivationCode: code,
		}, nil
	}
	return model.User{
		Name: email,
	}, nil
}

// Close db connection.
func (r *Repository) Close() {
	r.db.Close()
}
