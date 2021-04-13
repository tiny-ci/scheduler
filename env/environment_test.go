package env

import (
	"os"
	"testing"
)

func setup(name string, content string, type_ EnvDataType, required bool) (Environment, error) {
	os.Setenv(name, content)

	return New(map[string]Var{
		"var": makeEV(name, type_, required),
	})
}

func TestStringEnv(t *testing.T) {
	name := "STRING_VAR"
	value := "value"

	environment, err := setup(name, value, DTString, true)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if v := environment.StringEnv(name); v != value {
		t.Fatalf("expected value: '%s'; got: '%s'", value, v)
	}
}

func TestBooleanEnv(t *testing.T) {
	name := "BOOLEAN_VAR"
	value := "true"

	environment, err := setup(name, value, DTBoolean, true)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if v := environment.BoolEnv(name); v != true {
		t.Fatalf("expected value: '%s'; got: '%t'", value, v)
	}
}

func TestNumberEnv(t *testing.T) {
	name := "NUMBER_VAR"
	value := "12345"

	environment, err := setup(name, value, DTNumber, true)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if v := environment.NumberEnv(name); v != 12345 {
		t.Fatalf("expected value: '%s'; got: '%d'", value, v)
	}
}

func TestListEnv(t *testing.T) {
	name := "LIST_ENV"
	value := "a,b,c,d"

	environment, err := setup(name, value, DTList, true)
	if err != nil {
		t.Fatalf(err.Error())
	}

	list := environment.ListEnv(name)
	if list[0] != "a" || list[1] != "b" || list[2] != "c" || list[3] != "d" {
		t.Fatalf("expected value: '%s'; got: '%v'", value, list)
	}
}

func TestWrongDataTypeEnv(t *testing.T) {
	name := "NUMBER_VAR"
	value := "abc"

	_, err := setup(name, value, DTNumber, true)
	if err == nil {
		t.Fatalf("data type validation should have failed")
	}

	t.Logf("got: %s", err.Error())
}

func TestMissingEnvs(t *testing.T) {
	os.Setenv("REQUIRED_ONE", "a")
	os.Setenv("REQUIRED_FOU", "a")

	_, err := New(map[string]Var{
		"reqOne": makeEV("REQUIRED_ONE", DTString, false),
		"reqTwo": makeEV("REQUIRED_TWO", DTString, false),
		"reqThr": makeEV("REQUIRED_THR", DTString, false),
	})

	if err == nil {
		t.Fatalf("required env checking should have failed")
	}

	t.Logf("got: %s", err.Error())
}
