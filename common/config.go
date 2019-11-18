package common

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// Configuration defines the configuration for the engine,
// it is loaded from a config.json file
type Configuration struct {
	Fullscreen bool    // if should run in fullscreen mode
	Scale      float64 // window's scale
	Width      int     // window's width
	Height     int     // window's height
}

// ConfigBasePath stores the base path of where to find
// the config.json file
var ConfigBasePath string

func init() {
	path := filepath.Dir(os.Args[0])
	var err error
	ConfigBasePath, err = filepath.Abs(path)

	if err != nil {
		logrus.Fatal(err)
	}
}

// LoadConfigurations loads all configurations from config.json
// and makes a Configuration struct from it
func LoadConfigurations() *Configuration {
	logrus.Trace("load configurations")

	// Reads the file
	filePath := path.Join(ConfigBasePath, "config.json")
	logrus.WithField("config-file", filePath).Debug("configuration load")
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		logrus.Fatal(err)
	}

	// Convert from what is in the file to Configuration
	var config Configuration
	err = json.Unmarshal(file, &config)

	if err != nil {
		logrus.Fatal(err)
	}

	fields := logrus.Fields{
		"fullscreen": config.Fullscreen,
		"width":      config.Width,
		"height":     config.Height,
		"scale":      config.Scale,
	}
	logrus.WithFields(fields).Debug("configuration loaded")

	return &config
}
