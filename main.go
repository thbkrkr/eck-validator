package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/elastic/cloud-on-k8s/pkg/apis/elasticsearch/v1beta1"
	"github.com/fatih/color"
	"sigs.k8s.io/yaml"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

func main() {
	in, err := readStdin()
	exitIf(err)

	jsonBytes, err := yaml.YAMLToJSON(in)
	exitIf(err)

	var es v1beta1.Elasticsearch
	dec := json.NewDecoder(bytes.NewReader(jsonBytes))
	dec.DisallowUnknownFields()

	err = dec.Decode(&es)
	if err != nil {
		fmt.Printf(red("KO")+" %v", cleanErr(err))
	} else {
		fmt.Printf(green("OK"))
	}
}

func readStdin() ([]byte, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return nil, errors.New("stdin is empty")
	}

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}

	return in, nil
}

func cleanErr(err error) error {
	// Remove the 'json: ' part from the error message
	return errors.New(strings.Replace(err.Error(), "json: ", "", -1))
}

func exitIf(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
