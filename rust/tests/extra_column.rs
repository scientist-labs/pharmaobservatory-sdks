// Hand-written (the generator never touches it): a section response that carries a column the spec
// does not list still decodes. serde ignores unknown fields on these models.
use pharmaobservatory::models::BiomarkersList200Response;

#[test]
fn section_row_with_unrecorded_column_decodes() {
    let raw = include_str!("../../testdata/biomarkers_list_extra_column.json");
    let resp: BiomarkersList200Response = serde_json::from_str(raw).expect("decode");
    assert_eq!(resp.rows[0].route_key, "braf-v600e");
}
