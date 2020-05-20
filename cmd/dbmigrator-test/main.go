package main

import (
	"context"
	"github.com/Kalinin-Andrey/dbmigrator-test/internal/app/cmd"
)

func main() {
	cmd.Execute(context.Background())
}
