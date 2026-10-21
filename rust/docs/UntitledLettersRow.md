# UntitledLettersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**letter_id** | Option<**String**> | The letter id. | [optional]
**letter_url** | Option<**String**> | The FDA page of the letter. | [optional]
**firm_name** | Option<**String**> | The firm's name as filed with the regulator. | [optional]
**firm_org_id** | Option<**String**> | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional]
**firm_company_id** | Option<**String**> | Company level: the firm's company (null when unresolved); the entry key for the companies object type. | [optional]
**firm_company_name** | Option<**String**> | Display name of the firm's company. | [optional]
**firm_group_id** | Option<**String**> | Group level: the firm's current ultimate parent; a root's group is itself. | [optional]
**firm_group_name** | Option<**String**> | Display name of the firm's group. | [optional]
**is_org_resolved** | Option<**bool**> | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional]
**match_tier** | Option<**String**> | How well the firm name was attached to an organization: high, medium or low name-and-geography matches, registry (the operating firm of a registered facility), ambiguous (several candidates, none chosen) or unresolved. Warning letters, untitled letters and debarments attach by firm name alone. | [optional]
**product_issue** | Option<**String**> | The product and the issue the letter raises (for example the promotional claim). | [optional]
**issue_date** | Option<[**String**](string.md)> | The date the letter was issued. | [optional]
**response_letter_url** | Option<**String**> | The firm's response letter on the FDA site, where one is posted. | [optional]
**closeout_letter_url** | Option<**String**> | The FDA close-out letter, where one is posted. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


