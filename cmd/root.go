package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/a-yee/spot/auth"
	"github.com/a-yee/spot/configs"
	"github.com/spf13/cobra"
)

var configFile string

func init() {
	var err error
	configFile, err = configs.DefaultPath()
	if err != nil {
		log.Fatal("Failed to load config file")
		os.Exit(1)
	}
}

var rootCmd = cobra.Command{
	Use:   "spot",
	Short: "Spotify TUI",
	Long:  "Music Player TUI for Spotify",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := configs.Load(configFile)
		if err != nil {
			return err
		}

		client := auth.NewAPIClient(c)
		// use the client to make calls that require authorization
		user, err := client.CurrentUser(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("You are logged in as:", user.ID)

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
