# RareDiseaseTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**orpha_code** | Option<**String**> | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional]
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**mapping_arm** | Option<**String**> | exact when the link rests on an asserted equivalence, unspecified when it rests on a MONDO cross-reference that states no relation. | [optional]
**edge_scope** | Option<**String**> | direct when the trial's own condition list names the mapped condition, broader_family when only a broader condition does. | [optional]
**via_exact** | Option<**bool**> | True when an exact mapping links the trial to the disease. | [optional]
**via_unspecified** | Option<**bool**> | True when an unspecified-relation mapping links the trial to the disease. | [optional]
**via_direct** | Option<**bool**> | True when a direct edge links the trial to the disease. | [optional]
**via_broader** | Option<**bool**> | True when a broader-family edge links the trial to the disease. | [optional]
**mapping_target_curie** | Option<**String**> | The MeSH condition (as a prefixed identifier) the mapping reaches the trial through. | [optional]
**mapping_sources** | Option<**Vec<String>**> | The files that assert the mapping (ordo, orphadata_product1 or mondo_legacy_bridge). | [optional]
**trial_title** | Option<**String**> | The trial's brief title. | [optional]
**overall_status** | Option<**String**> | The registry's overall recruitment status (COMPLETED, RECRUITING, TERMINATED and so on). | [optional]
**study_type** | Option<**String**> | The registry's study type (INTERVENTIONAL, OBSERVATIONAL or EXPANDED_ACCESS). | [optional]
**phase_label** | Option<**String**> | The trial phase as a display label (Phase 2, Phase 3, Observational and so on). | [optional]
**enrollment** | Option<**i32**> | The trial's total enrollment as registered. | [optional]
**enrollment_type** | Option<**String**> | ACTUAL when the registered enrollment is the number enrolled, ESTIMATED when it is the planned number. | [optional]
**start_date** | Option<[**String**](string.md)> | The trial's start date as registered. | [optional]
**start_year** | Option<**i32**> | The year of start_date. | [optional]
**completion_date** | Option<[**String**](string.md)> | The trial's completion date as registered. | [optional]
**why_stopped** | Option<**String**> | The reason the sponsor gave for stopping the trial, where it stopped early. | [optional]
**were_results_reported** | Option<**bool**> | True when the sponsor has posted results to the registry. | [optional]
**last_update_posted_date** | Option<[**String**](string.md)> | The date the registry record was last updated. | [optional]
**facility_count** | Option<**i32**> | Facilities the registry lists for the trial. | [optional]
**facility_country_count** | Option<**i32**> | Countries the listed facilities are in. | [optional]
**facility_countries** | Option<**Vec<String>**> | The countries the listed facilities are in. | [optional]
**days_since_update** | Option<**i32**> | Days from the last registry update to the registry snapshot date. | [optional]
**is_unverified** | Option<**bool**> | True when the registry record was last updated two or more years before the snapshot date. | [optional]
**sponsor_name** | Option<**String**> | The sponsor's name as filed in the registry. | [optional]
**sponsor_org_id** | Option<**String**> | Legal-entity level: the organization record the sponsor resolved to (null when unresolved). | [optional]
**sponsor_company_id** | Option<**String**> | Company level: the sponsor's company (null when unresolved). | [optional]
**sponsor_company_name** | Option<**String**> | Display name of the sponsor's company. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the sponsor's current ultimate parent company; a root company's group is itself. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the sponsor's group. | [optional]
**sponsor_agency_class** | Option<**String**> | The funder class the registry declared for the lead sponsor; null when the trial has no lead sponsor row. | [optional]
**sponsor_resolution_confidence** | Option<**String**> | How firmly the lead sponsor resolved to an organization record (high, medium or unresolved). | [optional]
**license_name** | Option<**String**> | The licence of the trial registry data. | [optional]
**attribution** | Option<**String**> | The attribution line the source requires; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


