package mapper

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/types/status"
)

type Status struct {
}

func (s Status) FromInternal(data map[string]interface{}) {
	status.Set(data)
}

func (s Status) ToInternal(data map[string]interface{}) {
}

func (s Status) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	schema.ResourceFields["state"] = types.Field{
		Type: "string",
	}
	schema.ResourceFields["transitioning"] = types.Field{
		Type: "enum",
		Options: []string{
			"yes",
			"no",
			"error",
		},
	}
	schema.ResourceFields["transitioningMessage"] = types.Field{
		Type: "string",
	}
	return nil
}
