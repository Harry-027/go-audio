package cmd

import (
	"github.com/spf13/cobra"
)

// Cobra root command ...
var RootCmd = &cobra.Command{
	Use:   "go-audio",
	Short: "go-audio is a go CLI client that connects with opentts to convert pdf content into audiofile.",
}
