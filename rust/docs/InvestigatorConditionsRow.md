# InvestigatorConditionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**investigator_person_id** | Option<**String**> | The investigator's resolved person key, the id in the web page address (/people/<id>). The entry key of the investigators object type. | [optional]
**investigator_person_name** | Option<**String**> | The investigator's name as the registry filed it, resolved across the person's merged name fragments. | [optional]
**investigator_affiliation** | Option<**String**> | The investigator's affiliation as filed (the most-trialled affiliation across the merged fragments). | [optional]
**condition_id** | Option<**String**> | The condition's MeSH descriptor id, the id in the conditions page address (D001943). | [optional]
**condition_name** | Option<**String**> | The condition's MeSH name. | [optional]
**n_trials** | Option<**i64**> | Distinct trials the person is named on in the condition. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


