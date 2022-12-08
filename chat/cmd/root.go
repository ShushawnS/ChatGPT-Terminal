/*
Copyright © 2022 Shushawn Saha s62saha@uwaterloo.ca

*/
package cmd

import (
	"os"
	"github.com/shushawns/ChatGPT-Terminal/chat/cmd/auth"
	"github.com/shushawns/ChatGPT-Terminal/chat/cmd/code"
	"github.com/shushawns/ChatGPT-Terminal/chat/cmd/conv"
	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ChatGPT",
	Short: "Use ChatGPT from OpenAI in the terminal!",
	Long: `[Add epic intro]`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommandPalettes(subCommand *cobra.Command) {
	rootCmd.AddCommand(subCommand)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ChatGPT-Terminal.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommandPalettes(auth.AuthCmd)
	addSubCommandPalettes(code.CodeCmd)
	addSubCommandPalettes(conv.ConvCmd)
}


