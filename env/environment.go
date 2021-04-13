package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Var struct {
	name     string
	dtype    EnvDataType
	optional bool
}

type Environment struct {
	values map[string]interface{}
}

func makeEV(name string, dtype EnvDataType, optional bool) Var {
	return Var{name, dtype, optional}
}

func parseData(dtype EnvDataType, value string) (interface{}, error) {
	switch dtype {
	case DTString:
		return value, nil

	case DTNumber:
		number, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}

		return number, nil

	case DTBoolean:
		val, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}

		return val, nil

	case DTList:
		return strings.Split(value, ","), nil

	default:
		return nil, errors.New(fmt.Sprintf("unknown dtype '%s'", dtype))
	}
}

func New(coreVars map[string]Var) (Environment, error) {
	env := Environment{
		values: make(map[string]interface{}),
	}

	var missing []string

	for _, v := range coreVars {
		value := os.Getenv(v.name)

		if value == "" {
			if v.optional {
				env.values[v.name] = ""
			} else {
				missing = append(missing, v.name)
			}

			continue
		}

		content, err := parseData(v.dtype, value)
		if err != nil {
			return env, errors.New(fmt.Sprintf("environment variable %s is not of type %s", v.name, v.dtype))
		}

		env.values[v.name] = content
	}

	if len(missing) > 0 {
		return env, errors.New(
			fmt.Sprintf("the following environment variables are missing: %s", strings.Join(missing, ", ")))
	}

	return env, nil
}

func (e Environment) StringEnv(name string) string {
	return e.values[name].(string)
}

func (e Environment) ListEnv(name string) []string {
	return e.values[name].([]string)
}

func (e Environment) NumberEnv(name string) int {
	return e.values[name].(int)
}

func (e Environment) BoolEnv(name string) bool {
	return e.values[name].(bool)
}
