package schema

import (
    "errors"
    "encoding/json"
    "log"
    "github.com/qri-io/jsonschema"
)

func validate(schema []byte, data []byte) error {
    rs := &jsonschema.RootSchema{}
    if err := json.Unmarshal(schema, rs); err != nil {
        return err
    }

    if jsErrors, _ := rs.ValidateBytes(data); len(jsErrors) > 0 {
        for _, err := range jsErrors {
            log.Println(err)
        }

        return errors.New("")
    }

    return nil
}
