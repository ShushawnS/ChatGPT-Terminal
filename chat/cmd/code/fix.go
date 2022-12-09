/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package code

import (
	"fmt"
	"os"
	"time"

	"github.com/TwiN/go-color"
	"github.com/sirupsen/logrus"
	chatgpt_go "github.com/zhan3333/chatgpt-go"
	"github.com/inancgumus/screen"
	"github.com/spf13/cobra"
)


func getCodeFix (filePath string, problem string) (code string) {
	file, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	return "Hi, can you please fix the following code: ``` " + string(file) + "```" + " The problem is: " + problem

}

// fixCmd represents the fix command
var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		screen.Clear()
		screen.MoveTopLeft()

		fmt.Println(color.Ize(color.Green,"\nHere is ChatGPT's fixed solution of the code"))
		fmt.Println(color.Ize(color.Green,"============================================"))

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

		userText := getCodeFix(filePath, problem)

		resp, err := conversation.SendMessage(userText)
		if err != nil {
			panic(err)
		}

		fmt.Println("\n", resp)

		//cmd.Help()
	},
}

func init() {
	//rootCmd.AddCommand(fixCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fixCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fixCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
