package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

// Accepts a string path, returns json data
func readJSON(path string) (*map[string]interface{}, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, errors.Wrap(err, "Can't open the file")
	}

	contents := make(map[string]interface{})
	err = json.Unmarshal(data, &contents)

	if err != nil {
		err = errors.Wrap(err, "Can't unmarshal file")
	}

	return &contents, err
}
