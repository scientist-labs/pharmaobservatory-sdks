# RareDiseaseTrialCoverageRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**orpha_code** | Option<**String**> | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional]
**coverage_state** | Option<**String**> | no_bridge, bridged_no_trials or trials. | [optional]
**mapping_arms** | Option<**Vec<String>**> | The mapping arms the disease has a condition under (exact, unspecified). | [optional]
**mapping_target_curie** | Option<**String**> | The MeSH condition or conditions the disease maps to (prefixed identifiers, comma separated), the exact ones where any exist. | [optional]
**mapping_target_name** | Option<**String**> | The names of those MeSH conditions. | [optional]
**mapping_sources** | Option<**Vec<String>**> | The files that assert the mapping. | [optional]
**mapping_source_releases** | Option<**Vec<String>**> | The releases of those files. | [optional]
**n_trials_exact_direct** | Option<**i64**> | Trials linked through an exact mapping whose own condition list names the disease. Count within this column only. | [optional]
**n_trials_exact_broader** | Option<**i64**> | Trials linked through an exact mapping to a broader condition than the trial studies. Count within this column only. | [optional]
**n_trials_unspecified** | Option<**i64**> | Trials linked through a MONDO cross-reference that states no relation. Count within this column only. | [optional]
**n_open_direct** | Option<**i64**> | Exact direct trials that are recruiting or not yet recruiting. | [optional]
**n_ongoing_direct** | Option<**i64**> | Exact direct trials that are active but not recruiting, or enrolling by invitation. | [optional]
**n_unverified_2y_direct** | Option<**i64**> | Exact direct trials whose record was last updated two or more years before the snapshot. | [optional]
**direct_share** | Option<**f64**> | The exact direct trials as a share of the exact direct and exact broader trials; null when both are zero. | [optional]
**n_sponsors** | Option<**i64**> | Sponsoring organizations of the exact direct trials. | [optional]
**snapshot_date** | Option<[**String**](string.md)> | The date of the registry snapshot. | [optional]
**freshness_as_of** | Option<[**String**](string.md)> | The registry's latest update date the snapshot reflects. | [optional]
**license_name** | Option<**String**> | The licence of the Orphadata Science release the mapping rests on. | [optional]
**attribution** | Option<**String**> | The attribution line the source requires; quote it with any figure taken from the row. | [optional]
**trials_license_name** | Option<**String**> | The licence of the trial registry data. | [optional]
**trials_attribution** | Option<**String**> | The attribution line the trial registry requires. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


