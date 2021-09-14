package yamlloader

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

// Load - loads yaml file to object or returns error
func Load(path string, o interface{}) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, o)
	if err != nil {
		return err
	}

	return nil
}

// MustLoad - load yaml file to object or throw fatal on error
func MustLoad(path string, o interface{}) {
	if err := Load(path, o); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
