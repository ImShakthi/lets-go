package snippets

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"lets-go/models"
)

func TomlParserInit() {
	var config models.Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)
}
