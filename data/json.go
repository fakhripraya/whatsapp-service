package data

import (
	"encoding/json"
)

// ExistanceResult is a struct to store the result of existance method
type ExistanceResult struct {
	Status int
	Jid    string
}

// UnmarshalJSON unmarshal the object from JSON formated string
func UnmarshalJSON(jsonString string, exResult *ExistanceResult) error {
	// unmarshall to the given instance
	err := json.Unmarshal([]byte(jsonString), &exResult)

	if err != nil {
		return err
	}

	return nil
}
