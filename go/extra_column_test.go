package pharmaobservatory

import (
	"encoding/json"
	"os"
	"testing"
)

// Hand-written (the generator never touches it): a dataset row keeps a column the spec does not
// list. Needs disallowAdditionalPropertiesIfNotPresent=false in config/openapi-config.json.
func TestDatasetRowKeepsUnrecordedColumn(t *testing.T) {
	raw, err := os.ReadFile("../testdata/biomarkers_extra_column.json")
	if err != nil {
		t.Fatal(err)
	}
	var resp BiomarkersResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	row := resp.Rows[0]
	if row.GetRouteKey() != "braf-v600e" {
		t.Errorf("route_key = %q", row.GetRouteKey())
	}
	if got := row.AdditionalProperties["unrecorded_column"]; got != "kept-by-the-client" {
		t.Errorf("unrecorded_column = %v, want it kept", got)
	}
}
