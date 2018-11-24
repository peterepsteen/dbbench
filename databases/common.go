package databases

import (
	"database/sql"
	"log"
)

type BenchType int

const (
	// Loop executes the benchmark several times.
	Loop BenchType = iota
	// Single executes the benchmark once.
	Single BenchType = iota
)

// Benchmark contains a benchmark func and its name.
type Benchmark struct {
	Name string
	Type BenchType
	Func func(int)
}

func mustExec(result sql.Result, err error, name string) {
	if err != nil {
		log.Fatalf("failed to %v: %v\n", name, err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("failed to get %v rows: %v\n", name, err)
	}
	if rows == 0 {
		log.Fatalf("no rows %ved\n", name)
	}
}
