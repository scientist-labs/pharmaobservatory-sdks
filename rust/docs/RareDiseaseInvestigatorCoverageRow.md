# RareDiseaseInvestigatorCoverageRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**orpha_code** | Option<**String**> | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional]
**coverage_state** | Option<**String**> | no_exact_mapping, no_direct_trials, no_publishable_investigator or computed. | [optional]
**n_direct_trials** | Option<**i64**> | Exact direct trials of the disease. | [optional]
**n_direct_trials_with_named_pi** | Option<**i64**> | Those trials that name a principal investigator. | [optional]
**n_pi_trial_edges** | Option<**i64**> | Trial and principal investigator pairs among them. | [optional]
**n_pi_people** | Option<**i64**> | Distinct people named as principal investigator. | [optional]
**n_investigators** | Option<**i64**> | Rows of rare_disease_investigators for the disease. | [optional]
**n_name_ambiguous** | Option<**i64**> | Investigators whose name another investigator of the disease shares. | [optional]
**registry_reference_date** | Option<[**String**](string.md)> | The registry's latest update date the snapshot reflects. | [optional]
**license_name** | Option<**String**> | The licence of the Orphadata Science release the disease list rests on. | [optional]
**attribution** | Option<**String**> | The attribution line the source requires; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


