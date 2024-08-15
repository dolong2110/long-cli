package cmd

import "github.com/spf13/cobra"

var Root = &cobra.Command{
	Use:     "long",
	Version: "",
	Short:   "",
	Long:    ``,
}

func init() {
	Root.AddCommand(
		wcCmd,
	)
}
