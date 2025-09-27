package main

import (
	"github.com/hoangnguyen-ca/transaction-isolation/internal/db"
	"github.com/hoangnguyen-ca/transaction-isolation/internal/examples"
	"github.com/hoangnguyen-ca/transaction-isolation/internal/examples/counterex"
	"github.com/jackc/pgx/v5"
)

func main() {
	pool := db.GetDBPool()
	myExamples := []examples.Example{counterex.NewCounterExample(pool, &pgx.TxOptions{IsoLevel: pgx.ReadCommitted})}

	examples.InitExamples(myExamples)
	examples.RunExamples(myExamples, 100)
}
