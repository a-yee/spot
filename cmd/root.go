package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/a-yee/spot/app"
	"github.com/a-yee/spot/auth"
	"github.com/a-yee/spot/configs"
	"github.com/a-yee/spot/ui"
	"github.com/a-yee/spot/ui/component/footer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	api "github.com/zmb3/spotify/v2"
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
		players, err := client.PlayerDevices(context.Background())
		if err != nil {
			return err
		}

		var spotDeviceID api.ID
		for _, device := range players {
			if device.Name == c.DeviceName {
				spotDeviceID = device.ID
				break
			}
		}
		if spotDeviceID == "" {
			return fmt.Errorf("Either need to restart spotifyd or " +
				"missing/not recognized device_name in configs")
		}

		playerState, err := client.PlayerState(context.Background())
		if err != nil {
			return err
		}

		if playerState.Playing {
			err = client.Pause(context.Background())
			if err != nil {
				return err
			}
		}

		err = client.TransferPlayback(
			context.Background(),
			spotDeviceID,
			false)
		if err != nil {
			return err
		}

		// client.Play(context.Background())

		ai := app.NewAppInfo(
			context.Background(),
			client,
			0,
			0,
		)

		spot := ui.NewAppModel(ai)
		spot.SetFooter(footer.New(ai, spot))
		_, err = tea.NewProgram(
			spot,
			tea.WithAltScreen(),
			//tea.WithMouseCellMotion(),
		).Run()
		return err
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
