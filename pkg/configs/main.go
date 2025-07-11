package configs

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
	"github.com/tidwall/pretty"
)

// New reads configuration with koanf.
func New(path, prefix string, instance interface{}) (interface{}, error) {
	k := koanf.New(".")

	// load default configuration from file
	err := k.Load(structs.Provider(instance, "koanf"), nil)
	if err != nil {
		return nil, err
	}

	// load configuration from file
	err = k.Load(file.Provider(path), yaml.Parser())
	if err != nil {
		log.Printf("error loading %s: %s", path, err)
	}

	// load environment variables
	err = k.Load(env.Provider(prefix, ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, prefix)), "__", ".")
	}), nil)
	if err != nil {
		log.Printf("error loading environment variables: %s", err)
	}

	err = k.Unmarshal("", &instance)
	if err != nil {
		return nil, err
	}

	indent, err := json.MarshalIndent(instance, "", "\t")
	if err != nil {
		return nil, err
	}

	indent = pretty.Color(indent, nil)
	tmpl := `
	================ Loaded Configuration ================
	%s
	======================================================
	`
	log.Printf(tmpl, string(indent))

	return instance, nil
}
