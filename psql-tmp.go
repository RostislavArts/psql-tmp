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
	uri, pgTmpErr := pgTmpRun.Output()
	if pgTmpErr != nil {
		fmt.Println("pg_tmp:", pgTmpErr)
		os.Exit(1)
	}

	psqlRun := exec.Command("psql", string(uri), "-f", sqlScript, "-o", "./psql-output")
	if err := psqlRun.Run(); err != nil {
		fmt.Println("psql:", err)
		os.Exit(1)
	}

	cat := exec.Command("cat", "./psql-output")
	output, catErr := cat.Output()
	if pgTmpErr != nil {
		fmt.Println("cat:", catErr)
		os.Exit(1)
	}

	fmt.Println(string(output))

	remove := exec.Command("rm", "./psql-output")
	if err := remove.Run(); err != nil {
		fmt.Println("rm:", err)
		os.Exit(1)
	}
}
