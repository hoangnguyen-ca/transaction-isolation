// Package counterex
package counterex

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CounterExample struct {
	pool *pgxpool.Pool
}

func NewCounterExample(pool *pgxpool.Pool) CounterExample {
	return CounterExample{pool: pool}
}

func (ce CounterExample) Init() {
	query :=
		`
		CREATE TABLE IF NOT EXISTS counter (
		    id BOOLEAN PRIMARY KEY DEFAULT TRUE,
			value BIGINT NOT NULL DEFAULT 0
		);

		INSERT INTO counter (id, value)
		VALUES (TRUE, 0)
		ON CONFLICT (id) DO NOTHING;
		`
	_, err := ce.pool.Exec(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
}

// Run the example
func (ce CounterExample) Run() {
	// spin up threads and run example
	iterations := 100
	var wg sync.WaitGroup
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ce.incrementCounter()
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			v1, v2 := ce.doubleReadCounter()
			fmt.Printf("Counter read: %v %v\n", v1, v2)
		}()
	}
	wg.Wait()
}

func (ce CounterExample) incrementCounter() {
	ctx := context.Background()
	query :=
		`
		UPDATE counter 
		SET value = value + 1;
		`
	_, err := ce.pool.Exec(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
}

func (ce CounterExample) doubleReadCounter() (int, int) {
	ctx := context.Background()
	tx, err := ce.pool.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}
	query :=
		`
		SELECT value from counter;
		`
	var val1 int
	var val2 int
	err = tx.QueryRow(ctx, query).Scan(&val1)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.QueryRow(ctx, query).Scan(&val2)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit(ctx)
	return val1, val2
}
