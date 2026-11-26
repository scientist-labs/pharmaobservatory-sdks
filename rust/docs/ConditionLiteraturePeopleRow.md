# ConditionLiteraturePeopleRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**investigator_person_id** | Option<**String**> | The author's resolved person key, the id in the web page address (/people/<id>). The entry key of the investigators object type. | [optional]
**person_name** | Option<**String**> | The person's name, resolved across the person's merged name fragments. | [optional]
**person_affiliation** | Option<**String**> | The person's affiliation as the person record holds it (null when not recorded). | [optional]
**orcid** | Option<**String**> | The person's ORCID iD (Open Researcher and Contributor ID) when the person record carries one (null when none). | [optional]
**n_papers** | Option<**i32**> | Distinct papers on which a ledger decision links a byline to the person. | [optional]
**n_trial_linked_papers** | Option<**i32**> | The part of n_papers a registered clinical trial cites. | [optional]
**latest_pub_year** | Option<**i32**> | The newest publication year among those papers. | [optional]
**n_first_author** | Option<**i32**> | The part of n_papers on which the person's linked byline is the first author position. | [optional]
**n_last_author** | Option<**i32**> | The part of n_papers on which the person's linked byline is the last author position. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


