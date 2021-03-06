/*
Copyright 2018 Heptio Inc.

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

package app

import (
	"os"

	ops "github.com/heptio/sonobuoy/cmd/sonobuoy/app/operations"
	"github.com/heptio/sonobuoy/pkg/errlog"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var path string

func init() {
	cmd := &cobra.Command{
		Use:   "cp",
		Short: "Copies the results to a specified path",
		Run:   copyResults,
	}
	cmd.PersistentFlags().StringVar(
		&path, "path", "./",
		"TBD: location to output",
	)
	// TODO: Other Options.?.?
	RootCmd.AddCommand(cmd)
}

func copyResults(cmd *cobra.Command, args []string) {
	code := 0
	if err := ops.CopyResults(path); err != nil {
		errlog.LogError(errors.Wrap(err, "error attempting to copy sonobuoy results"))
		code = 1
	}
	os.Exit(code)
}
