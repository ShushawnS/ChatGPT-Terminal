/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package code

import (
	"fmt"
	"os"
	"time"

	"github.com/inancgumus/screen"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	chatgpt_go "github.com/zhan3333/chatgpt-go"
)

func getCodeExplain(filePath string) (code string) {
	file, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	return "Hi, can you please explain the following code: ``` " + string(file) + "```"

}

// explainCmd represents the explain command
var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		screen.Clear()
		screen.MoveTopLeft()

		fmt.Println("\nHere is ChatGPT's explanation of the code")
		fmt.Println("============================================")

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

		userText := getCodeExplain(filePath)

		resp, err := conversation.SendMessage(userText)
		if err != nil {
			panic(err)
		}

		fmt.Println("\n", resp)

		cmd.Help()
	},
}

func init() {
	//rootCmd.AddCommand(explainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// explainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// explainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
