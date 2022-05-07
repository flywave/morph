package main

import (
	"os"

	"github.com/flywave/morph/drivers"
	"github.com/pkg/errors"

	"github.com/flywave/morph"

	"github.com/flywave/morph/commands"
)

func main() {
	if err := commands.RootCmd().Execute(); err != nil {
		var databaseErr *drivers.DatabaseError
		if errors.As(errors.Cause(err), &databaseErr) {
			morph.ErrorLogger.Fprintf(os.Stderr, "An Error Occurred: This and all later migrations have been cancelled\n")
		} else {
			morph.ErrorLogger.Fprintf(os.Stderr, "An Error Occurred:\n")
		}
		_, _ = morph.ErrorLoggerLight.Fprintf(os.Stderr, "--> %v\n", err)
	}
}
