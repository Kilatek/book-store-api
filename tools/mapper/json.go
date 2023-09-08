package mapper

import "encoding/json"

// MapStructsWithJSONTags maps two structs based on their JSON tags.
func MapStructsWithJSONTags(source any, dest any) error {
	raw, err := json.Marshal(source)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(raw, dest); err != nil {
		return err
	}

	return nil
}
