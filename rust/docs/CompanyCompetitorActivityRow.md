# CompanyCompetitorActivityRow

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
**competitor_company_id** | Option<**String**> | Company level: the company active in the same disease area. | [optional]
**competitor_company_name** | Option<**String**> | Display name of the competitor's company. | [optional]
**competitor_group_id** | Option<**String**> | Group level: the competitor's current ultimate parent; a root's group is itself. | [optional]
**competitor_group_name** | Option<**String**> | Display name of the competitor's group. | [optional]
**competitor_rank** | Option<**i32**> | 0 for the company itself, 1 to 6 for the companies ranked by trial count across the company's top disease areas. | [optional]
**trials** | Option<**i32**> | Distinct trials the company in this row has in the disease area. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


