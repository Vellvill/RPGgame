package postgres

import (
	"Consumer/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

const pgxConnErr = "can not connect to postgres"

type connErr struct {
	pgxConnErr string
	attempts   int
	time       int
}

func (c connErr) Error() string {
	return fmt.Sprintf("%d attempts, %d time, %s", c.attempts, c.time, pgxConnErr)
}

type Client interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	Ping(ctx context.Context) error
}

func NewClient(config *config.Config, attempts, timeToConnect int) (Client, error) {
	dsn := fmt.Sprintf("%s:%s/%s:%s/%s")
	return ReConnect(timeToConnect, attempts, dsn)
}

func ReConnect(t, att int, dsn string) (Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
	defer cancel()
	for i := att; i > 0; i-- {
		pool, err := pgxpool.Connect(ctx, dsn)
		if err != nil {
			continue
		}
		return pool, nil
	}
	return nil, connErr{
		pgxConnErr: pgxConnErr,
		attempts:   att,
		time:       t,
	}
}
