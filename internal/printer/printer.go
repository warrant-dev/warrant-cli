// Copyright 2023 Forerunner Labs, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package printer

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/muesli/termenv"
)

var Purple = termenv.ColorProfile().Color("#6310FF")
var Red = termenv.ColorProfile().Color("#FF0000")
var Green = termenv.ColorProfile().Color("#00FF00")
var Checkmark = "✔"
var Cross = "✖"

func init() {
	if runtime.GOOS == "windows" {
		Checkmark = "√"
		Cross = "×"
	}
}

func PrintJson(val any) {
	bytes, err := json.MarshalIndent(val, "", "    ")
	if err != nil {
		PrintErrAndExit(err.Error())
	}
	fmt.Printf("%s\n", string(bytes))
}

func PrintErrAndExit(msg string) {
	fmt.Fprintln(os.Stderr, "Error:", msg)
	os.Exit(1)
}
