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

	"github.com/pkg/errors"
	"github.com/warrant-dev/warrant-cli/internal/config"
	"github.com/warrant-dev/warrant-go/v6"
)

// Read an objectType and objectId from string
func ReadObjectArg(arg string) (string, string, error) {
	typeAndId := strings.Split(arg, ":")
	if len(typeAndId) != 2 {
		return "", "", fmt.Errorf("invalid object")
	}

	return typeAndId[0], typeAndId[1], nil
}

// Read object metadata from json string
func ReadObjectMetaArg(arg string) (map[string]interface{}, error) {
	var meta map[string]interface{}
	err := json.Unmarshal([]byte(arg), &meta)
	if err != nil {
		return meta, errors.Wrap(err, "invalid object meta")
	}
	return meta, nil
}

// Read subjectType, subjectId and optional relation from a subject string
func ReadSubjectArg(arg string) (string, string, string, error) {
	subjectAndRelation := strings.Split(arg, "#")
	if len(subjectAndRelation) > 2 {
		return "", "", "", fmt.Errorf("invalid subject")
	}
	typeAndId := strings.Split(subjectAndRelation[0], ":")
	if len(typeAndId) != 2 {
		return "", "", "", fmt.Errorf("invalid subject")
	}
	if len(subjectAndRelation) == 1 {
		return typeAndId[0], typeAndId[1], "", nil
	}
	if len(subjectAndRelation) == 2 {
		return typeAndId[0], typeAndId[1], subjectAndRelation[1], nil
	}

	return "", "", "", fmt.Errorf("invalid subject")
}

// Read subject, relation, object and optional context from args
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

// Read subject, relation, object and optional policy from args
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
