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
	"encoding/json"
	"fmt"
	"strings"

	"github.com/warrant-dev/warrant-go/v5"
)

type Object struct {
	Type     string
	Id       string
	Relation string
}

func ParseObject(arg string) (Object, error) {
	arr := strings.Split(arg, ":")
	if len(arr) != 2 {
		return Object{}, fmt.Errorf("Invalid object provided")
	}
	idAndRelation := strings.Split(arr[1], "#")
	if len(idAndRelation) == 2 {
		return Object{
			Type:     arr[0],
			Id:       idAndRelation[0],
			Relation: idAndRelation[1],
		}, nil
	} else {
		return Object{
			Type: arr[0],
			Id:   arr[1],
		}, nil
	}
}

func ReadCheckArgs(args []string) (*warrant.WarrantCheckParams, error) {
	subject, err := ParseObject(args[0])
	if err != nil {
		return nil, err
	}
	relation := args[1]
	object, err := ParseObject(args[2])
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
				ObjectType: object.Type,
				ObjectId:   object.Id,
			},
			Relation: relation,
			Subject: warrant.Subject{
				ObjectType: subject.Type,
				ObjectId:   subject.Id,
				Relation:   subject.Relation,
			},
			Context: context,
		},
	}, nil
}

func ReadWarrantArgs(args []string) (*warrant.WarrantParams, error) {
	subject, err := ParseObject(args[0])
	if err != nil {
		return nil, err
	}
	relation := args[1]
	object, err := ParseObject(args[2])
	if err != nil {
		return nil, err
	}

	var policy string
	if len(args) == 4 {
		policy = args[3]
	}

	return &warrant.WarrantParams{
		ObjectType: object.Type,
		ObjectId:   object.Id,
		Relation:   relation,
		Subject: warrant.Subject{
			ObjectType: subject.Type,
			ObjectId:   subject.Id,
			Relation:   subject.Relation,
		},
		Policy: policy,
	}, nil
}
