# InvestigatorTopicsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**investigator_person_id** | Option<**String**> | The investigator's resolved person key, the id in the web page address (/people/<id>). The entry key of the investigators object type. | [optional]
**descriptor_id** | Option<**String**> | The MeSH descriptor id. | [optional]
**descriptor_name** | Option<**String**> | The MeSH descriptor name. | [optional]
**is_disease** | Option<**bool**> | True when the descriptor is also a condition. | [optional]
**condition_id** | Option<**String**> | The descriptor id as a conditions id when the descriptor is a condition (null otherwise). | [optional]
**n_publications** | Option<**i64**> | Distinct papers the person authored on the descriptor. | [optional]
**n_major_topic_publications** | Option<**i64**> | Those papers where the descriptor is a major topic. | [optional]
**topic_weight** | Option<**i64**> | n_publications plus n_major_topic_publications, so a major-topic paper counts twice. | [optional]
**topic_rank** | Option<**i32**> | The person's own topic order by weight (1 is their strongest topic); not a ranking of people. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


