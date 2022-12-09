/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package conv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/TwiN/go-color"
	"github.com/sirupsen/logrus"
	chatgpt_go "github.com/zhan3333/chatgpt-go"
	"github.com/inancgumus/screen"
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

// ConvCmd represents the conv command
var ConvCmd = &cobra.Command{
	Use:   "conv",
	Short: "Converse with OpenAI's ChatGPT",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		screen.Clear()
		screen.MoveTopLeft()

		fmt.Println(color.Ize(color.Green,"\nConversation with ChatGPT ( type 'bye' to exit )"))
		fmt.Println(color.Ize(color.Green,"================================================="))

		loginData := getLoginData()
		//fmt.Println("Current Login Data: (email) %s (password) %s (apikey) %s", loginData.Email, loginData.Password, loginData.ApiKey)
		bye := false
		reader := bufio.NewScanner(os.Stdin)
		timeout := time.Second * 60
		client, err := chatgpt_go.NewChatGPT(loginData.ApiKey, chatgpt_go.ChatGPTOptions{
			Log:     logrus.NewEntry(logrus.StandardLogger()),
			Timeout: &timeout,
		})
		if err != nil {
			panic(err)
		}
		conversation := client.NewConversation("", "")

		for !bye {
			fmt.Print("\nMe: ")

			reader.Scan()

			userText := reader.Text()

			if userText == "bye" {
				bye = true
			} else {
				resp, err := conversation.SendMessage(userText)
				if err != nil {
					panic(err)
				}
				fmt.Println("ChatGPT: \n", resp)
			}

		}

		//cmd.Help()
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
