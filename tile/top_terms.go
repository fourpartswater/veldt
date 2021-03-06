package tile

import (
	"fmt"

	"github.com/unchartedsoftware/veldt/util/json"
)

// TopTerms represents a tile which returns counts for the top most occurring
// terms in a provided field.
// fieldType is an optional value representing the type of the field.  Currently only 'string' is supported, and all
// other fieldType values default to an array of strings.
type TopTerms struct {
	TermsField string
	TermsCount int
	FieldType  string
}

// Parse parses the provided JSON object and populates the tiles attributes.
func (t *TopTerms) Parse(params map[string]interface{}) error {
	termsField, ok := json.GetString(params, "termsField")
	if !ok {
		return fmt.Errorf("`termsField` parameter missing from tile")
	}
	termsCount, ok := json.GetInt(params, "termsCount")
	if !ok {
		return fmt.Errorf("`termsCount` parameter missing from tile")
	}
	fieldType, ok := json.GetString(params, "fieldType")
	if ok && fieldType == "string" {
		t.FieldType = fieldType
	}
	t.TermsField = termsField
	t.TermsCount = termsCount
	return nil
}
