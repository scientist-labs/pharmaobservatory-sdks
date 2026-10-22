# RareDiseaseSponsorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**orpha_code** | Option<**String**> | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional]
**sponsor_name** | Option<**String**> | The sponsor's name as filed in the registry. | [optional]
**sponsor_org_id** | Option<**String**> | Legal-entity level: the organization record the sponsor resolved to. | [optional]
**sponsor_company_id** | Option<**String**> | Company level: the sponsor's company (null when unresolved). | [optional]
**sponsor_company_name** | Option<**String**> | Display name of the sponsor's company. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the sponsor's current ultimate parent company; a root company's group is itself. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the sponsor's group. | [optional]
**registry_sponsor_name** | Option<**String**> | The sponsor's most common name as filed across the disease's trials. | [optional]
**org_type** | Option<**String**> | The organization's type. | [optional]
**sponsor_agency_class** | Option<**String**> | The funder class the registry declared for the sponsor. | [optional]
**sponsor_resolution_confidence** | Option<**String**> | How firmly the sponsor resolved to an organization record (high, medium, low or unresolved). | [optional]
**n_trials** | Option<**i64**> | Distinct exact direct trials of the disease the sponsor leads or collaborates on. | [optional]
**n_trials_lead** | Option<**i64**> | Those trials the sponsor leads. | [optional]
**n_trials_collaborator** | Option<**i64**> | Those trials the sponsor collaborates on. | [optional]
**n_completed** | Option<**i64**> | Those trials the registry marks completed. | [optional]
**n_terminated** | Option<**i64**> | Those trials marked terminated. | [optional]
**n_withdrawn** | Option<**i64**> | Those trials marked withdrawn. | [optional]
**n_suspended** | Option<**i64**> | Those trials marked suspended. | [optional]
**n_open** | Option<**i64**> | Those trials recruiting or not yet recruiting. | [optional]
**n_ongoing** | Option<**i64**> | Those trials active but not recruiting, or enrolling by invitation. | [optional]
**n_other** | Option<**i64**> | Those trials in any other status. | [optional]
**n_unverified_3y** | Option<**i64**> | Those trials whose record was last updated three or more years before the snapshot. | [optional]
**first_trial_start** | Option<[**String**](string.md)> | The earliest start date among those trials. | [optional]
**last_trial_start** | Option<[**String**](string.md)> | The latest start date among those trials. | [optional]
**first_trial_start_year** | Option<**i32**> | The year of the earliest start. | [optional]
**last_trial_start_year** | Option<**i32**> | The year of the latest start. | [optional]
**last_update_posted_date** | Option<[**String**](string.md)> | The latest registry update among those trials. | [optional]
**is_dormant** | Option<**bool**> | True when the sponsor's latest trial start is more than five years before the registry snapshot. | [optional]
**mapping_arm** | Option<**String**> | The mapping arm the dataset is built at (always exact). | [optional]
**edge_scope** | Option<**String**> | The edge scope the dataset is built at (always direct). | [optional]
**freshness_as_of** | Option<[**String**](string.md)> | The registry's latest update date the snapshot reflects. | [optional]
**license_name** | Option<**String**> | The licence of the trial registry data. | [optional]
**attribution** | Option<**String**> | The attribution line the source requires; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


