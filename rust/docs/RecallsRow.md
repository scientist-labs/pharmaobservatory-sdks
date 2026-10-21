# RecallsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**recall_number** | Option<**String**> | The FDA recall number. | [optional]
**domain** | Option<**String**> | The regulatory domain of the recalled product (drug or device here). | [optional]
**status** | Option<**String**> | The recall status (Ongoing, Completed or Terminated). | [optional]
**recall_class** | Option<**String**> | The severity class: Class I, II or III. | [optional]
**product_type** | Option<**String**> | The product type the FDA reports (Drugs, Devices and so on). | [optional]
**firm_name** | Option<**String**> | The firm's name as filed with the regulator (the recalling firm). | [optional]
**firm_org_id** | Option<**String**> | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional]
**firm_company_id** | Option<**String**> | Company level: the firm's company (null when unresolved); the entry key for the companies object type. | [optional]
**firm_company_name** | Option<**String**> | Display name of the firm's company. | [optional]
**firm_group_id** | Option<**String**> | Group level: the firm's current ultimate parent; a root's group is itself. | [optional]
**firm_group_name** | Option<**String**> | Display name of the firm's group. | [optional]
**is_org_resolved** | Option<**bool**> | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional]
**match_tier** | Option<**String**> | How well the firm name was attached to an organization: high, medium or low name-and-geography matches, registry (the operating firm of a registered facility), ambiguous (several candidates, none chosen) or unresolved. Warning letters, untitled letters and debarments attach by firm name alone. | [optional]
**org_resolution_method** | Option<**String**> | The route that attached the firm: fei_bridge (the FDA establishment identifier, FEI, bridge), fei_firm_mint or fei_firm_fold (the operating firm's registry record), firm_name_geo (name and geography), firm_name_only (name alone) or unresolved. | [optional]
**product_description** | Option<**String**> | The recalled product as described in the report. | [optional]
**reason_for_recall** | Option<**String**> | The reason the FDA gives for the recall. | [optional]
**recall_initiation_date** | Option<[**String**](string.md)> | The date the firm began the recall. | [optional]
**report_date** | Option<[**String**](string.md)> | The date the FDA reported the recall. | [optional]
**termination_date** | Option<[**String**](string.md)> | The date the recall was terminated, where it was. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


