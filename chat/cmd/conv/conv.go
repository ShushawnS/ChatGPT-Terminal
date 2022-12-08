/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package conv

import (
	//"fmt"

	"github.com/spf13/cobra"
)

// ConvCmd represents the conv command
var ConvCmd = &cobra.Command{
	Use:   "conv",
	Short: "Converse with OpenAI's ChatGPT",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	//ootCmd.AddCommand(convCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
