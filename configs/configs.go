package configs

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Config defines valid configuration parameters for spot app
type Config struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_url"`
}

// Load creates a new Config struct with the values stored in the given file.
// Currently, only supports json, will modify in the future to support more
// formats.
func Load(configFile string) (Config, error) {
	configJson, err := os.Open(configFile)
	if err != nil {
		return Config{}, err
	}
	defer configJson.Close()

	configBytes, err := ioutil.ReadAll(configJson)
	if err != nil {
		return Config{}, err
	}

	var c Config
	json.Unmarshal(configBytes, &c)

	return c, nil
}

// DefaultPath defines the system default configuration file location
func DefaultPath() (string, error) {
	// TODO: refactor this to support more platforms
	//configDir, err := os.UserConfigDir()
	//if err != nil {
	//	return "", err
	//}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".config", "spot", "config.json"), nil
}
