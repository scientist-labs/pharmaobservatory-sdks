# RareDiseaseInvestigatorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**orpha_code** | Option<**String**> | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional]
**disease_name** | Option<**String**> | The disease's preferred English name. | [optional]
**investigator_person_id** | Option<**String**> | The investigator's resolved person key. | [optional]
**investigator_person_name** | Option<**String**> | The investigator's name as the registry filed it. | [optional]
**identity_basis** | Option<**String**> | What the person's identity rests on: curated, npi, orcid, reporter, affinity or self. | [optional]
**merged_fragment_count** | Option<**i64**> | Name fragments merged into this person. | [optional]
**name_is_ambiguous_in_disease** | Option<**bool**> | True when another investigator of the same disease shares the name. | [optional]
**pi_trial_count** | Option<**i64**> | Exact direct trials of the disease that name the person as principal investigator. | [optional]
**first_trial_start_date** | Option<[**String**](string.md)> | The earliest start date among those trials. | [optional]
**last_trial_start_date** | Option<[**String**](string.md)> | The latest start date among those trials. | [optional]
**last_trial_start_year** | Option<**i32**> | The year of the latest start. | [optional]
**evidence_trial_id** | Option<**String**> | One of the trials that names the person (the latest started). | [optional]
**last_record_update_date** | Option<[**String**](string.md)> | The latest registry update among those trials. | [optional]
**registry_reference_date** | Option<[**String**](string.md)> | The registry's latest update date the snapshot reflects. | [optional]
**recency_state** | Option<**String**> | How recent the latest trial start is against the registry snapshot: within_5_years, 5_to_10_years, over_10_years or no_start_date. | [optional]
**trial_site_names** | Option<**Vec<String>**> | The sites of the person's trials as the registry lists them. | [optional]
**trial_site_count** | Option<**i64**> | How many such sites. | [optional]
**stated_affiliations** | Option<**Vec<String>**> | The affiliations the registry states for the person. | [optional]
**stated_affiliation_count** | Option<**i64**> | How many stated affiliations. | [optional]
**affiliation_evidence** | Option<**String**> | trial_site when a trial site is listed, stated_affiliation_only when only an affiliation is stated, none otherwise. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


