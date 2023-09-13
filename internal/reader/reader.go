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

package reader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/warrant-dev/warrant-cli/internal/config"
	"github.com/warrant-dev/warrant-go/v5"
)

func ReadObjectArg(arg string) (string, string, error) {
	if strings.Contains(arg, "#") {
		return "", "", fmt.Errorf("invalid object: cannot contain '#'")
	}
	typeAndId := strings.Split(arg, ":")
	if len(typeAndId) == 1 {
		return typeAndId[0], "", nil
	}
	if len(typeAndId) == 2 {
		return typeAndId[0], typeAndId[1], nil
	}
	return "", "", fmt.Errorf("invalid object")
}

func ReadObjectMetaArg(arg string) (map[string]interface{}, error) {
	var meta map[string]interface{}
	err := json.Unmarshal([]byte(arg), &meta)
	if err != nil {
		return meta, err
	}
	return meta, nil
}

func ReadSubjectArg(arg string) (string, string, string, error) {
	objAndRelation := strings.Split(arg, "#")
	objType, id, err := ReadObjectArg(objAndRelation[0])
	if len(objAndRelation) == 1 {
		return objType, id, "", err
	}
	if len(objAndRelation) == 2 {
		return objType, id, objAndRelation[1], err
	}
	return "", "", "", fmt.Errorf("invalid subject")
}

func ReadCheckArgs(args []string) (*warrant.WarrantCheckParams, error) {
	if len(args) < 3 || len(args) > 4 {
		return nil, fmt.Errorf("invalid check: %s", args)
	}
	subjectType, subjectId, subjectRelation, err := ReadSubjectArg(args[0])
	if err != nil {
		return nil, err
	}
	relation := args[1]
	objectType, objectId, err := ReadObjectArg(args[2])
	if err != nil {
		return nil, err
	}

	var context warrant.PolicyContext
	if len(args) == 4 {
		contextStr := args[3]
		json.Unmarshal([]byte(contextStr), &context)
	}

	return &warrant.WarrantCheckParams{
		WarrantCheck: warrant.WarrantCheck{
			Object: warrant.Object{
				ObjectType: objectType,
				ObjectId:   objectId,
			},
			Relation: relation,
			Subject: warrant.Subject{
				ObjectType: subjectType,
				ObjectId:   subjectId,
				Relation:   subjectRelation,
			},
			Context: context,
		},
	}, nil
}

func ReadWarrantArgs(args []string) (*warrant.WarrantParams, error) {
	if len(args) < 3 || len(args) > 4 {
		return nil, fmt.Errorf("invalid warrant: %s", args)
	}
	subjectType, subjectId, subjectRelation, err := ReadSubjectArg(args[0])
	if err != nil {
		return nil, err
	}
	relation := args[1]
	objectType, objectId, err := ReadObjectArg(args[2])
	if err != nil {
		return nil, err
	}

	var policy string
	if len(args) == 4 {
		policy = args[3]
	}

	return &warrant.WarrantParams{
		ObjectType: objectType,
		ObjectId:   objectId,
		Relation:   relation,
		Subject: warrant.Subject{
			ObjectType: subjectType,
			ObjectId:   subjectId,
			Relation:   subjectRelation,
		},
		Policy: policy,
	}, nil
}

func PromptAndReadFromStdIn(prompt string) (string, error) {
	fmt.Println(prompt + ":")
	buf := bufio.NewReader(os.Stdin)
	input, err := buf.ReadBytes('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(input)), nil
}

func ReadEnvFromConsole() (string, *config.Environment, error) {
	envName, err := PromptAndReadFromStdIn("Enter environment name")
	if err != nil {
		return "", nil, err
	}

	apiKey, err := PromptAndReadFromStdIn("Enter API key")
	if err != nil {
		return "", nil, err
	}

	apiEndpoint, err := PromptAndReadFromStdIn("Warrant endpoint override (leave blank to use default https://api.warrant.dev)")
	if err != nil {
		return "", nil, err
	}
	if apiEndpoint == "" {
		apiEndpoint = "https://api.warrant.dev"
	}

	return envName, &config.Environment{
		ApiKey:      apiKey,
		ApiEndpoint: apiEndpoint,
	}, nil
}
