# ApiV1TrialsRecordsGet200ResponseTrialsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**canonical_trial_id** | **String** |  | 
**nct_id** | **String** |  | 
**source_registry** | **String** |  | 
**registry** | **String** |  | 
**brief_title** | **String** |  | 
**official_title** | Option<**String**> |  | 
**study_type** | **String** |  | 
**overall_status** | **String** |  | 
**overall_status_summary** | **String** |  | 
**phase** | Option<**String**> |  | 
**phase_by_status** | **String** |  | 
**enrollment** | Option<**i32**> |  | 
**enrollment_type** | Option<**String**> |  | 
**enrollment_flag_reason** | Option<[**serde_json::Value**](.md)> |  | 
**start_date** | **String** |  | 
**start_year** | **i32** |  | 
**primary_completion_date** | Option<**String**> |  | 
**primary_completion_date_type** | Option<**String**> |  | 
**completion_date** | Option<**String**> |  | 
**completion_date_type** | Option<**String**> |  | 
**study_first_posted_date** | **String** |  | 
**last_update_posted_date** | **String** |  | 
**why_stopped** | Option<**String**> |  | 
**number_of_arms** | Option<[**serde_json::Value**](.md)> |  | 
**number_of_facilities** | **i32** |  | 
**were_results_reported** | **bool** |  | 
**has_us_facility** | Option<**bool**> |  | 
**gender** | **String** |  | 
**minimum_age** | Option<**String**> |  | 
**maximum_age** | Option<**String**> |  | 
**healthy_volunteers** | Option<**bool**> |  | 
**criteria** | **String** |  | 
**lead_sponsor_name** | **String** |  | 
**lead_sponsor_id** | **String** |  | 
**collaborators** | **Vec<String>** |  | 
**participating_sponsors** | **Vec<String>** |  | 
**conditions** | **Vec<String>** |  | 
**condition_mesh** | **Vec<String>** |  | 
**interventions** | **Vec<String>** |  | 
**intervention_types** | **Vec<String>** |  | 
**intervention_mesh** | Option<**Vec<String>**> |  | 
**facility_count** | Option<**i32**> |  | 
**facility_countries** | Option<**Vec<String>**> |  | 
**facility_country_count** | Option<**i32**> |  | 
**enrollment_is_plausible** | **bool** |  | 
**phase_label** | **String** |  | 
**extracted** | Option<[**models::ApiV1TrialsRecordsGet200ResponseTrialsInnerExtracted**](_api_v1_trials_records_get_200_response_trials_inner_extracted.md)> |  | 
**linked_drugs** | Option<[**Vec<models::ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner>**](_api_v1_trials_records_get_200_response_trials_inner_linked_drugs_inner.md)> |  | [optional]
**linked_drugs_built** | Option<**bool**> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


