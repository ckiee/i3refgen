package main

import (
	"github.com/cep21/xdgbasedir"
	"io/ioutil"
	"log"
	"strings"
)

func readParse() []keybinding {
	xdg, err := xdgbasedir.GetConfigFileLocation("i3")
	cfg := xdg + "/config"
	if err != nil {
		log.Fatal("Error while getting XDG Config location", err)
	}
	str, err2 := read(cfg)
	if err2 != nil {
		log.Fatal("Error while reading file", err2)
	}
	return parse(*str)
}
func getFilePath() (string, error) {
	xdg, err := xdgbasedir.GetConfigFileLocation("i3")
	cfg := xdg + "/config"
	return cfg, err
}
func read(path string) (*string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	str := string(file)
	return &str, nil
}

func parse(data string) []keybinding {
	arr := strings.Split(data, "\n")
	var res []keybinding
	for _, v := range arr {
		if strings.HasPrefix(v, "bindsym") {
			v = strings.TrimPrefix(v, "bindsym ")
			v = strings.TrimPrefix(v, "--release ")
			// log.Println(v)
			key := strings.Split(v, " ")[0]
			cmd := strings.TrimPrefix(v, key)
			kb := keybinding{key, cmd}
			res = append(res, kb)
		}
	}
	return res

}
