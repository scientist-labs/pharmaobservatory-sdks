# PdufaDatesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**pdufa_date_id** | Option<**String**> | The calendar entry id. | [optional]
**source** | Option<**String**> | Which lane the row came from: benzinga, sec_extraction or both. | [optional]
**restriction_tier** | Option<**String**> | none for a row from SEC filings; licensable for a row from the Benzinga calendar (the filings are those made to the US Securities and Exchange Commission). | [optional]
**decision_date** | Option<[**String**](string.md)> | The decision date of the US Food and Drug Administration (FDA); a PDUFA date is the goal date under the Prescription Drug User Fee Act. | [optional]
**date_start** | Option<[**String**](string.md)> | Start of the date interval the source states. | [optional]
**date_end** | Option<[**String**](string.md)> | End of the date interval the source states. | [optional]
**date_precision** | Option<**String**> | Precision of the date (day, month, quarter). | [optional]
**date_phrase** | Option<**String**> | The date as the source words it. | [optional]
**event_type** | Option<**String**> | Event type. | [optional]
**event_category** | Option<**String**> | Event category (pdufa, adcomm and so on). | [optional]
**drug_name** | Option<**String**> | The drug name as filed. | [optional]
**drug_id** | Option<**String**> | The drug key. | [optional]
**molecule_id** | Option<**String**> | The molecule key, the resolved drug of the calendar entry and the entry key for the drugs object type (null when the drug name did not resolve to a drug). | [optional]
**sponsor_name** | Option<**String**> | The sponsor name as filed. | [optional]
**sponsor_org_id** | Option<**String**> | Legal-entity level: the organization record the sponsor resolved to (null when unresolved). | [optional]
**sponsor_company_id** | Option<**String**> | Company level: the company the sponsor belongs to (null when unresolved). | [optional]
**sponsor_company_name** | Option<**String**> | Display name of the sponsor's company. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the sponsor's current ultimate parent company; a root company's group is itself. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the sponsor's group. | [optional]
**ticker** | Option<**String**> | Exchange ticker. | [optional]
**condition_id** | Option<**String**> | The condition key. | [optional]
**condition_name** | Option<**String**> | The condition's name. | [optional]
**indication_text** | Option<**String**> | The indication as the source words it. | [optional]
**source_type** | Option<**String**> | Kind of source document. | [optional]
**source_link** | Option<**String**> | Link to the source document. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


