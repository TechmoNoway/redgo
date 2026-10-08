package command

import (
	"errors"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func Parse(input string) (Command, error) {
	fields := strings.Fields(input)

	if len(fields) == 0 {
		return Command{}, errors.New("empty command")
	}

	return Command{
		Name: strings.ToUpper(fields[0]),
		Args: fields[1:],
	}, nil
}
