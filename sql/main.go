package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type UserRec struct {
	User        string
	DisplayName string
	ID          int
}

type Storage struct {
	conn        *sql.DB
	getUserStmt *sql.Stmt
}

func main() {

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)
}

func NewStorage(ctx context.Context, conn *sql.DB) *Storage {
	c, _ := conn.PrepareContext(ctx, `SELECT "User", "DisplayName" FROM users WHERE "ID" = $1`)
	return &Storage{
		getUserStmt: c,
	}
}

func (s *Storage) GetUser(ctx context.Context, conn *sql.DB, id int) (UserRec, error) {

	u := UserRec{ID: id}
	err := s.getUserStmt.QueryRow(id).Scan(&u)

	return u, err

}
