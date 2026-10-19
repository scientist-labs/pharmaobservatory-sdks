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

// Hand-written: bin/go-spec-patch leaves TrialDeliveryRow.has_enrollment_delta undeclared on the Go
// copy (it collides with the generated HasEnrollmentDelta() method), so it must still decode, as an
// additional property, beside the typed enrollment_delta.
func TestCollidingHasColumnIsKeptAsAnAdditionalProperty(t *testing.T) {
	var row TrialDeliveryRow
	if err := json.Unmarshal([]byte(`{"enrollment_delta":12,"has_enrollment_delta":true}`), &row); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if got := row.GetEnrollmentDelta(); got != 12 {
		t.Errorf("enrollment_delta = %d", got)
	}
	if !row.HasEnrollmentDelta() {
		t.Error("HasEnrollmentDelta() is the presence check of enrollment_delta")
	}
	if got := row.AdditionalProperties["has_enrollment_delta"]; got != true {
		t.Errorf("has_enrollment_delta = %v, want it kept", got)
	}
}
