// Hand-written (the generator never touches it): a dataset response that carries a column the spec
// does not list still decodes. serde ignores unknown fields on these models.
use pharmaobservatory::models::BiomarkersResponse;

#[test]
fn dataset_row_with_unrecorded_column_decodes() {
    let raw = include_str!("../../testdata/biomarkers_extra_column.json");
    let resp: BiomarkersResponse = serde_json::from_str(raw).expect("decode");
    assert_eq!(resp.rows[0].route_key, Some(Some("braf-v600e".to_string())));
}
