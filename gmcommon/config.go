package gmcommon

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"path"
)

// Configuration defines the configuration for the engine,
// it is loaded from a config.json file
type Configuration struct {
	Fullscreen bool
	Scale      float64
	Width      int
	Height     int
}

// ConfigBasePath stores the base path of where to find
// the config.json file
var ConfigBasePath string

func init() {
	path := filepath.Dir(os.Args[0])
	var err error
	ConfigBasePath, err = filepath.Abs(path)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Configuration base path:", ConfigBasePath)
}

// LoadConfigurations loads all configurations from config.json
// and makes a Configuration struct from it
func LoadConfigurations() Configuration {
	log.Println("Loading configurations")

	filePath := path.Join(ConfigBasePath, "config.json")
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	var config Configuration
	err = json.Unmarshal(file, &config)

	if err != nil {
		log.Fatal(err)
	}

	return config
}
