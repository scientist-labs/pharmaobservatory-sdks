# ConditionLiteratureOrgsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**org_id** | Option<**String**> | Legal-entity level: the resolved parent organization record the byline names. | [optional]
**org_name** | Option<**String**> | The organization's name. | [optional]
**org_type** | Option<**String**> | The organization's type. | [optional]
**org_country** | Option<**String**> | The organization's country code (null when not recorded). | [optional]
**n_papers** | Option<**i32**> | Distinct papers naming the organization on a byline. | [optional]
**n_trial_linked_papers** | Option<**i32**> | The part of n_papers a registered clinical trial cites. | [optional]
**latest_pub_year** | Option<**i32**> | The newest publication year among those papers. | [optional]
**org_arm** | Option<**String**> | company when the organization is company-typed; institution for every other organization. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


