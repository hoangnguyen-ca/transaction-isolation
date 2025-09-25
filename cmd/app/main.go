package main

import (
	"github.com/hoangnguyen-ca/transaction-isolation/internal/db"
	"github.com/hoangnguyen-ca/transaction-isolation/internal/examples"
	"github.com/hoangnguyen-ca/transaction-isolation/internal/examples/counter"
)

func main() {
	pool := db.GetDBPool()
	myExamples := []examples.Example{counter.NewCounterExample(pool)}

	examples.InitExamples(myExamples)
	examples.RunExamples(myExamples)
}
