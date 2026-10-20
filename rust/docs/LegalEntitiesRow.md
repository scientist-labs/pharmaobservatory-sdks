# LegalEntitiesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**org_id** | Option<**String**> | The organization record id: a registered legal entity (edgar:, ror:) or a name-resolved sponsor (sponsor:). | [optional]
**org_name** | Option<**String**> | The organization's name. | [optional]
**is_name_resolved_sponsor** | Option<**bool**> | True for a sponsor resolved from a filed name alone rather than from a registration. | [optional]
**org_origin** | Option<**String**> | How the organization record was minted. | [optional]
**country** | Option<**String**> | Country as the source states it. | [optional]
**org_type** | Option<**String**> | Organization type: Academic, Government, Hospital, Industry, Nonprofit or Other (from the ROR record or the filing), or the lowercase tags investor and manufacturer that mark a minted adviser or manufacturer record. | [optional]
**edgar_cik** | Option<**String**> | Central index key (CIK) in EDGAR, the US Securities and Exchange Commission (SEC) filing system, when it files. | [optional]
**edgar_ticker** | Option<**String**> | Exchange ticker, when listed. | [optional]
**ror_id** | Option<**String**> | Research Organization Registry id, when it has one. | [optional]
**lei** | Option<**String**> | Legal Entity Identifier, when it has one. | [optional]
**lifecycle_status** | Option<**String**> | Lifecycle status from the registry sources. | [optional]
**num_lead_trials** | Option<**i64**> | Trials the organization leads. | [optional]
**num_trials** | Option<**i64**> | Trials the organization leads or collaborates on. | [optional]
**total_lead_enrollment** | Option<**f64**> | Planned enrollment across led trials. | [optional]
**num_recent_lead_trials** | Option<**i64**> | Led trials started in the recent window. | [optional]
**activity_score** | Option<**f64**> | Trial activity score. | [optional]
**company_id** | Option<**String**> | Company level: the company the organization belongs to (null when unresolved). | [optional]
**company_name** | Option<**String**> | Display name of the company. | [optional]
**group_id** | Option<**String**> | Group level: the company's current ultimate parent; a root's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


