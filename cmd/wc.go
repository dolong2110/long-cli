package cmd

import "github.com/spf13/cobra"

var wcCmd = &cobra.Command{
	Use:     "wc",
	Version: "",
	Short:   "",
	Long:    ``,
	Args:    cobra.MinimumNArgs(1),
	Run:     wc,
}

func wc(cmd *cobra.Command, args []string) {

}
