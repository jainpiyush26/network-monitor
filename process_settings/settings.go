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
	interval int64
	overall  int64
	servers  map[string]interface{}
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

	var monitor_settings Settings

	json_settings := gjson.Parse(string(settings_json))

	monitor_settings.output = (json_settings.Get("output path").Get("value")).String()
	monitor_settings.pattern = (json_settings.Get("output pattern").Get("value")).String()
	monitor_settings.interval = (json_settings.Get("runtime interval").Get("value")).Int()
	monitor_settings.overall = (json_settings.Get("runtime overall").Get("value")).Int()
	monitor_settings.servers = json_settings.Get("servers").Get("value").Value().(map[string]interface{})

	fmt.Println(monitor_settings.interval)
	fmt.Println(monitor_settings.pattern)
	for _key,_value := range monitor_settings.servers{
		fmt.Println("Name:", _key, "=>", "Connect to:", _value)
	}
	fmt.Println(monitor_settings.output)

}
