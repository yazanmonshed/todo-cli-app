/*
Copyright © 2020  yazanmonshed.com

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
	// "github.com/spf13/cli/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/cobra/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo app already build",
	Long:  "todo app build for imporved you work and be superpower user",
}

func main() {
	cmd.Execute()
}
