/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rlmcpherson/kops/upup/pkg/api"
	"github.com/spf13/cobra"
)

// GetCmd represents the get command
type GetCmd struct {
	output string

	cobraCommand *cobra.Command
}

var getCmd = GetCmd{
	cobraCommand: &cobra.Command{
		Use:        "get",
		SuggestFor: []string{"list"},
		Short:      "list or get objects",
		Long:       `list or get objects`,
	},
}

const (
	OutputYaml  = "yaml"
	OutputTable = "table"
	OutputJSON  = "json"
)

func init() {
	cmd := getCmd.cobraCommand

	rootCommand.AddCommand(cmd)

	cmd.PersistentFlags().StringVarP(&getCmd.output, "output", "o", OutputTable, "output format.  One of: table, yaml, json")
}

type marshalFunc func(v interface{}) ([]byte, error)

func marshalToStdout(item interface{}, marshal marshalFunc) error {
	b, err := marshal(item)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(b)
	if err != nil {
		return fmt.Errorf("error writing to stdout: %v", err)
	}
	return nil
}

// v must be a pointer to a marshalable object
func marshalYaml(v interface{}) ([]byte, error) {
	y, err := api.ToYaml(v)
	if err != nil {
		return nil, fmt.Errorf("error marshaling yaml: %v", err)
	}
	return y, nil
}

// v must be a pointer to a marshalable object
func marshalJSON(v interface{}) ([]byte, error) {
	j, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json: %v", err)
	}
	return j, nil
}
