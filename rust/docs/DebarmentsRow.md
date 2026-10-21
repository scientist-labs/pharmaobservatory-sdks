# DebarmentsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**debarment_id** | Option<**String**> | The debarment id. | [optional]
**firm_name** | Option<**String**> | The firm's name as filed with the regulator (the debarred entity). | [optional]
**firm_org_id** | Option<**String**> | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional]
**firm_company_id** | Option<**String**> | Company level: the firm's company (null when unresolved); the entry key for the companies object type. | [optional]
**firm_company_name** | Option<**String**> | Display name of the firm's company. | [optional]
**firm_group_id** | Option<**String**> | Group level: the firm's current ultimate parent; a root's group is itself. | [optional]
**firm_group_name** | Option<**String**> | Display name of the firm's group. | [optional]
**is_org_resolved** | Option<**bool**> | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional]
**match_tier** | Option<**String**> | How well the firm name was attached to an organization: high, medium or low name-and-geography matches, registry (the operating firm of a registered facility), ambiguous (several candidates, none chosen) or unresolved. Warning letters, untitled letters and debarments attach by firm name alone. | [optional]
**effective_date** | Option<[**String**](string.md)> | The date the debarment took effect. | [optional]
**term** | Option<**String**> | The debarment term (for example permanent or a number of years). | [optional]
**fr_citation** | Option<**String**> | The Federal Register citation of the notice. | [optional]
**fr_date** | Option<[**String**](string.md)> | The Federal Register publication date. | [optional]
**source_url** | Option<**String**> | The FDA debarment list page. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


