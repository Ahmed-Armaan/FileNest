package helper

import (
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
)

var (
	ErrUnResolveable   = errors.New("Could not resolve as Postgres error")
	ErrUniqueViolation = errors.New("unique kwy violation")
	ErrUnResolved      = errors.New("unresolved error")
	ErrDuplicateObject = errors.New("Object already exist")
)

func ResolvePostgresError(err error) error {
	fmt.Println("Validating error")
	if err == nil {
		return nil
	}
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		fmt.Println("Woah its truly unresolvable")
		return ErrUnResolveable
	}
	fmt.Printf("Code: %s\n", pgErr.Code)

	switch pgErr.Code {
	case "23505":
		return ErrUniqueViolation
	case "42710":
		return ErrDuplicateObject
	default:
		return ErrUnResolved
	}
}
