package main

import (
	"fmt"
	"errors"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("psql-tmp: missing file operand\n")
		fmt.Printf("Try 'psql-tmp --help' for more information\n")
		os.Exit(1)
	}

	if os.Args[1] == "--help" {
		fmt.Printf("psql-tmp is a tool for executing .sql files in temporary PostgreSQL database.\n\n")
		fmt.Printf("Usage: psql-tmp FILE\n")
		os.Exit(1)
	}

	sqlScript := os.Args[1]

	if _, err := os.Stat(sqlScript); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("psql-tmp: cannot access '%s': No such file or directory\n", sqlScript)
		os.Exit(1)
	}

	pgTmpRun := exec.Command("pg_tmp")
	uri, _ := pgTmpRun.Output()

	psqlRun := exec.Command("psql", string(uri), "-f", sqlScript, "-o", "./psql-output")
	psqlRun.Run()

	cat := exec.Command("cat", "./psql-output")
	output, _ := cat.Output()
	fmt.Println(string(output))

	exec.Command("rm", "./psql-output").Run()
}
