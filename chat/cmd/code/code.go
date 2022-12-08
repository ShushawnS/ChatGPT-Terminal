/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package code

import (
	//"fmt"

	"github.com/spf13/cobra"
)

// CodeCmd represents the code command
var CodeCmd = &cobra.Command{
	Use:   "code",
	Short: "Interact with OpenAi's ChatGPT using code",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	//rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
