/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package code

import (
	"fmt"
	"time"

	"github.com/TwiN/go-color"
	"github.com/inancgumus/screen"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	chatgpt_go "github.com/zhan3333/chatgpt-go"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		screen.Clear()
		screen.MoveTopLeft()

		fmt.Println(color.Ize(color.Green,"\nHere is ChatGPT's code"))
		fmt.Println(color.Ize(color.Green,"========================="))

		loginData := getLoginData()
		timeout := time.Second * 60
		client, err := chatgpt_go.NewChatGPT(loginData.ApiKey, chatgpt_go.ChatGPTOptions{
			Log:     logrus.NewEntry(logrus.StandardLogger()),
			Timeout: &timeout,
		})
		if err != nil {
			panic(err)
		}
		conversation := client.NewConversation("", "")

		userText := prompt
		//fmt.Println(userText)

		resp, err := conversation.SendMessage(userText)
		if err != nil {
			panic(err)
		}

		fmt.Println("\n", resp)

		//cmd.Help()
	},
}

func init() {
	//Cmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
