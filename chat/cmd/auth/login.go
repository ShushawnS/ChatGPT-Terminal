/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package auth

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


func saveToFile(email string, password string, apikey string) {
	filePath := "../config.json"

	fo, err := os.Create(filePath)
	if err != nil{
		panic(err)
	} 

	data := map[string] interface{} {
		"email":	email,
		"password": password,
		"apikey":	apikey,
	}

	defer fo.Close()

	e := json.NewEncoder(fo)
	if err := e.Encode(data); err != nil {
		panic(err)
	}

	fmt.Println(fo)
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login into your OpenAI account",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		saveToFile(email, password, apikey)

		//cmd.Help()
	},
}

func init() {
	//rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
