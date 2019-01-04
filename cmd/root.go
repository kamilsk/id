package cmd

import "github.com/spf13/cobra"

// RootCmd is the entry point of command-line interface.
var RootCmd = &cobra.Command{Use: "passport", Short: "Passport"}

func init() { RootCmd.AddCommand(completionCmd, migrateCmd, runCmd) }
