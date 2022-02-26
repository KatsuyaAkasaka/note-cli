package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "nt",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Cmd() *cobra.Command {
	rootCmd.AddCommand(
		setTodo(),
	)
	return rootCmd
}
