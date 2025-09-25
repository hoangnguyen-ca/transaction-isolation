package counter

import (
	"context"
	"log"

	"github.com/hoangnguyen-ca/transaction-isolation/internal/examples"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CounterExample struct {
	pool *pgxpool.Pool
}

func NewCounterExample(pool *pgxpool.Pool) CounterExample {
	return CounterExample{pool: pool}
}

var _ examples.Example = (*CounterExample)(nil)

func (ce CounterExample) Init() {
	_, err := ce.pool.Exec(context.Background(),
		`
        CREATE TABLE IF NOT EXISTS counters (
            id BIGSERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            value BIGINT NOT NULL DEFAULT 0
        )
    `)
	if err != nil {
		log.Fatal(err)
	}
}

func (ce CounterExample) Run() {
	// spin up threads and run example
}

func incrementCounter() {

}

func readCounter() {

}
