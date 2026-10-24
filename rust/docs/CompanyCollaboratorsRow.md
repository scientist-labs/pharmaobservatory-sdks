# CompanyCollaboratorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**company_id** | Option<**String**> | Company level: the company the row is about; the entry key for the companies object type. | [optional]
**company_name** | Option<**String**> | Display name of the company. | [optional]
**group_id** | Option<**String**> | Group level: the company's current ultimate parent; a root's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**condition_name** | Option<**String**> | The MeSH descriptor name of the condition. | [optional]
**disease_rank** | Option<**i32**> | The condition's rank among the company's disease areas by trial count (1 is the busiest). Only the top 10 are kept. | [optional]
**collaborator_company_id** | Option<**String**> | Company level: the company that co-sponsors the trials. | [optional]
**collaborator_company_name** | Option<**String**> | Display name of the collaborator's company. | [optional]
**collaborator_group_id** | Option<**String**> | Group level: the collaborator's current ultimate parent; a root's group is itself. | [optional]
**collaborator_group_name** | Option<**String**> | Display name of the collaborator's group. | [optional]
**collaborator_rank** | Option<**i32**> | 0 for the company itself, 1 to 6 for its collaborators ranked by shared trials across its top disease areas. | [optional]
**shared_trials** | Option<**i32**> | Distinct trials the company shares with the collaborator in this disease area (for rank 0, the company's own trials in it). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


