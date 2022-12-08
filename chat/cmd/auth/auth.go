/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package auth

import (
	//"fmt"

	"github.com/spf13/cobra"
)

var (
	email string
	password string
	apikey string
)

// AuthCmd represents the auth command
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Provide credentials to authenticate with OpenAI's ChatGPT",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	//rootCmd.AddCommand(authCmd)
	loginCmd.Flags().StringVarP(&email,"email", "e", "", "The email to your OpenAI account")
	loginCmd.Flags().StringVarP(&password,"password", "p", "", "The password to your OpenAI account")
	loginCmd.Flags().StringVarP(&apikey,"apikey", "k", "", "The apikey to your OpenAI account")
	// Here you will define your flags and configuration settings.
	AuthCmd.AddCommand(loginCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
