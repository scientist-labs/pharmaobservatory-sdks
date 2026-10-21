# InspectionFindingsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**citation_id** | Option<**String**> | The citation id. | [optional]
**inspection_id** | Option<**String**> | The inspection the citation belongs to; join to the inspections dataset. | [optional]
**fei** | Option<**String**> | The FDA establishment identifier (FEI) of the inspected site. | [optional]
**firm_name** | Option<**String**> | The firm's name as filed with the regulator. | [optional]
**firm_org_id** | Option<**String**> | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional]
**firm_company_id** | Option<**String**> | Company level: the firm's company (null when unresolved); the entry key for the companies object type. | [optional]
**firm_company_name** | Option<**String**> | Display name of the firm's company. | [optional]
**firm_group_id** | Option<**String**> | Group level: the firm's current ultimate parent; a root's group is itself. | [optional]
**firm_group_name** | Option<**String**> | Display name of the firm's group. | [optional]
**is_org_resolved** | Option<**bool**> | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional]
**match_tier** | Option<**String**> | How well the firm was attached to an organization: high, medium, low, registry, ambiguous or unresolved. | [optional]
**inspection_date** | Option<[**String**](string.md)> | The date of the inspection. | [optional]
**classification** | Option<**String**> | The outcome of the inspection: NAI, VAI or OAI. | [optional]
**act_cfr_number** | Option<**String**> | The cited regulation (a Code of Federal Regulations section or Act citation). | [optional]
**short_description** | Option<**String**> | The FDA's short description of the observation. | [optional]
**program_area** | Option<**String**> | The FDA program area of the citation. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


