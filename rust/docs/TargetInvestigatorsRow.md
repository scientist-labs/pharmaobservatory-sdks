# TargetInvestigatorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_id** | Option<**String**> | The target's Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. | [optional]
**investigator_person_id** | Option<**String**> | The investigator's resolved person key, the id in the web page address (/people/<id>). The entry key of the investigators object type. | [optional]
**investigator_person_name** | Option<**String**> | The investigator's name as the registry filed it, resolved across the person's merged name fragments. | [optional]
**investigator_affiliation** | Option<**String**> | The investigator's affiliation as filed (the most-trialled affiliation across the merged fragments). | [optional]
**n_trials** | Option<**i64**> | Distinct trials the person is named on that are in scope of the row (the target's trials, or the drug's). | [optional]
**is_principal_investigator** | Option<**bool**> | True when the person holds the principal investigator role on any registered trial (the person-level flag the investigators dataset and the people page publish), not only the trials of this row. | [optional]
**roles** | Option<**Vec<String>**> | The distinct roles the person holds on the trials of this row, as the registry filed them (PRINCIPAL_INVESTIGATOR, STUDY_DIRECTOR, STUDY_CHAIR, SUB_INVESTIGATOR). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


