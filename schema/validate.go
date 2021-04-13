package schema

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/qri-io/jsonschema"
	"log"
)

func validate(schema []byte, data []byte) error {
	ctx := context.Background()

	rs := &jsonschema.Schema{}
	if err := json.Unmarshal(schema, rs); err != nil {
		return err
	}

	if jsErrors, _ := rs.ValidateBytes(ctx, data); len(jsErrors) > 0 {
		for _, err := range jsErrors {
			log.Println(err)
		}

		return errors.New("")
	}

	return nil
}
