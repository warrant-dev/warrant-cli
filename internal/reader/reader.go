package reader

import (
	"fmt"
	"strings"
)

type Object struct {
	Type string
	Id   string
}

func ParseObject(arg string) (Object, error) {
	arr := strings.Split(arg, ":")
	if len(arr) != 2 {
		return Object{}, fmt.Errorf("Invalid object provided")
	}
	return Object{
		Type: arr[0],
		Id:   arr[1],
	}, nil
}
