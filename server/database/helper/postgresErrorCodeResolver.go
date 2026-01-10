package helper

import (
	"errors"

	"github.com/jackc/pgconn"
)

var (
	ErrUnResolveable   = errors.New("Could not resolve as Postgres error")
	ErrUniqueViolation = errors.New("unique kwy violation")
	ErrUnResolved      = errors.New("unresolved error")
)

func ResolvePostgresError(err error) error {
	if err == nil {
		return nil
	}
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return ErrUnResolveable
	}

	switch pgErr.Code {
	case "23505":
		return ErrUniqueViolation
	default:
		return ErrUnResolved
	}
}
