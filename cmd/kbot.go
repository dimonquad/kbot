/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	tele "gopkg.in/telebot.v4"
)

var (
	teleToken = os.Getenv("Tele_Token")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kbot %s started", appVersion)
		kbot, err := tele.NewBot(tele.Settings{
			URL:    "",
			Token:  teleToken,
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Check systemEnv: %v", err)
			return
		}

		kbot.Handle(tele.OnText, func(c tele.Context) error {
			log.Print(c.Message().Payload, c.Text())
			payload := c.Message().Payload
			switch payload {
			case "hello":
				return c.Send("Hello! How can I help you?")
			}
			return err
		})
		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
