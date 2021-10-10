package settings

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/tidwall/gjson"
)

type Settings struct {
	output   string
	pattern  string
	interval int
	overall  int
	servers  map[string]string
}

func Run() {
	settings_path, ok := os.LookupEnv("MONITOR_SETTINGS_PATH")
	if !ok {
		settings_path, _ = filepath.Abs("./config/settings.json")
	}
	_, err := os.Stat(settings_path)
	if os.IsNotExist(err) {
		log.Fatal("Monitor settings path do not exist!")
	}

	settings_json, err := ioutil.ReadFile(settings_path)

	if err != nil {
		log.Fatal(err)
	}

	monitor_settings := gjson.Get(string(settings_json), "servers.value")
	fmt.Println(monitor_settings.Map())

}
