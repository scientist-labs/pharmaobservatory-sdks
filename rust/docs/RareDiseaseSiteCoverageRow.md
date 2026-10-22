# RareDiseaseSiteCoverageRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**orpha_code** | Option<**String**> | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional]
**concept_type** | Option<**String**> | Whether the concept is a group of disorders, a disorder or a subtype. | [optional]
**snapshot_date** | Option<[**String**](string.md)> | The date of the registry snapshot. | [optional]
**exact_state** | Option<**String**> | How far the exact arm reaches: no_mapping, no_trials, no_site_data or computed. | [optional]
**exact_n_conditions** | Option<**i32**> | MeSH conditions the disease maps to under the exact arm. | [optional]
**exact_n_trials** | Option<**i32**> | Trials linked under the exact arm. | [optional]
**exact_n_trials_with_site** | Option<**i32**> | Trials of the exact arm with at least one geocoded site. | [optional]
**exact_n_sites_total** | Option<**i32**> | Distinct sites of the exact arm. | [optional]
**exact_n_sites_emitted** | Option<**i32**> | Sites of the exact arm published in rare_disease_trial_sites (capped at 500 per disease and arm). | [optional]
**exact_n_sites_emitted_geocoded** | Option<**i32**> | Published sites of the exact arm that carry a coordinate. | [optional]
**exact_n_facilities_ungeocoded** | Option<**i32**> | Facility listings of the exact arm with no coordinate. | [optional]
**unspecified_state** | Option<**String**> | How far the unspecified arm reaches: no_mapping, no_trials, no_site_data or computed. | [optional]
**unspecified_n_conditions** | Option<**i32**> | MeSH conditions the disease maps to under the unspecified arm. | [optional]
**unspecified_n_trials** | Option<**i32**> | Trials linked under the unspecified arm. | [optional]
**unspecified_n_trials_with_site** | Option<**i32**> | Trials of the unspecified arm with at least one geocoded site. | [optional]
**unspecified_n_sites_total** | Option<**i32**> | Distinct sites of the unspecified arm. | [optional]
**unspecified_n_sites_emitted** | Option<**i32**> | Sites of the unspecified arm published in rare_disease_trial_sites (capped at 500 per disease and arm). | [optional]
**unspecified_n_sites_emitted_geocoded** | Option<**i32**> | Published sites of the unspecified arm that carry a coordinate. | [optional]
**unspecified_n_facilities_ungeocoded** | Option<**i32**> | Facility listings of the unspecified arm with no coordinate. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


