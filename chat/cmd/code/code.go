/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package code

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"

	"github.com/spf13/cobra"
)
type loginData struct {
	Email    string `json:email`
	Password string `json:password`
	ApiKey   string `json:apikey`
}

func getLoginData() (lData *loginData) {
	filePath := "../config.json"
	jsonFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var loginData loginData

	json.Unmarshal(byteValue, &loginData)

	return &loginData

}

var (
	filePath	string
	problem 	string
	prompt		string
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
	explainCmd.Flags().StringVarP(&filePath,"filePath", "f", "", "The file path to the code you want explained")
	fixCmd.Flags().StringVarP(&filePath,"filePath", "f", "", "The file path to the code you want fixed")
	fixCmd.Flags().StringVarP(&problem,"problem", "p", "", "Problem you are facing in the code")
	generateCmd.Flags().StringVarP(&prompt,"prompt", "p", "", "The prompt of the code you want generated")
	// Here you will define your flags and configuration settings.
	CodeCmd.AddCommand(explainCmd)
	CodeCmd.AddCommand(fixCmd)
	CodeCmd.AddCommand(generateCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
