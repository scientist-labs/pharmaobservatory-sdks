# InspectionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**inspection_id** | Option<**String**> | The inspection id. | [optional]
**fei** | Option<**String**> | The FDA establishment identifier (FEI) of the inspected site. | [optional]
**firm_name** | Option<**String**> | The firm's name as filed with the regulator (the inspected establishment's firm). | [optional]
**firm_org_id** | Option<**String**> | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional]
**firm_company_id** | Option<**String**> | Company level: the firm's company (null when unresolved); the entry key for the companies object type. | [optional]
**firm_company_name** | Option<**String**> | Display name of the firm's company. | [optional]
**firm_group_id** | Option<**String**> | Group level: the firm's current ultimate parent; a root's group is itself. | [optional]
**firm_group_name** | Option<**String**> | Display name of the firm's group. | [optional]
**is_org_resolved** | Option<**bool**> | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional]
**match_tier** | Option<**String**> | How well the firm was attached to an organization: high, medium, low, registry, ambiguous or unresolved. | [optional]
**org_resolution_method** | Option<**String**> | The route that attached the firm: fei_bridge (the FDA establishment identifier, FEI, bridge), fei_firm_mint or fei_firm_fold (the operating firm's registry record), firm_name_geo (name and geography), firm_name_only (name alone) or unresolved. | [optional]
**inspection_date** | Option<[**String**](string.md)> | The date the inspection ended. | [optional]
**fiscal_year** | Option<**String**> | The federal fiscal year of the inspection. | [optional]
**classification** | Option<**String**> | The outcome: NAI (no action indicated), VAI (voluntary action indicated) or OAI (official action indicated). | [optional]
**had_483** | Option<**bool**> | True when the inspection ended with a Form 483 of observations. | [optional]
**product_type** | Option<**String**> | The product type inspected (Drugs, Devices, Biologics, Veterinary, Tobacco or Food/Cosmetics). | [optional]
**project_area** | Option<**String**> | The FDA project area of the inspection. | [optional]
**city** | Option<**String**> | The establishment's city. | [optional]
**state** | Option<**String**> | The establishment's state or province. | [optional]
**postal_code** | Option<**String**> | The establishment's postal code. | [optional]
**country_name** | Option<**String**> | The establishment's country name. | [optional]
**country_iso2** | Option<**String**> | The establishment's two-letter country code. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


