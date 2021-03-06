// +build linux darwin
// +build !cobra

package main

import (
	"errors"
	"syscall"
)

func init() {
	signals = append(signals, syscall.SIGQUIT, syscall.SIGABRT)
}

// AdminChecker is a simple interface used to determine if admin is required.
type AdminChecker interface {
	IsAdmin() (bool, error)
}

// AdminRequired parses through an arg list and requires an admin (sudo)
// for the specified commands.
//
// Note that if the command is required, *and* we failed to get admin
// permission information, let the user run the command anyway.
// Better UX than failing for a possible non-issue.
func AdminRequired(args, reqs []string, p AdminChecker) error {
	// Ignore the permErr in the beginning. If the arg command
	isAdmin, permErr := p.IsAdmin()

	// If the user is admin, any admin requiring command is already
	// satisfied.
	if isAdmin {
		return nil
	}

	// At the moment, we're only checking the first level of commands.
	// Subcommands are ignored.
	if len(args) < 2 {
		return nil
	}

	c := args[1]

	var err error
	for _, r := range reqs {
		if c == r {
			// Use sudo terminology, for unix
			err = errors.New("Command requires sudo")
			break
		}
	}

	// If the command is required, *and* we failed to get admin permission
	// information, let the user run the command anyway. Better UX than
	// failing for a possible non-issue.
	if err != nil && permErr != nil {
		return nil
	}

	return err
}
