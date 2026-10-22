# CompanyRareDiseaseDeclarationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**assertion_id** | Option<**String**> | The declaration assertion id. | [optional]
**organization_name** | Option<**String**> | The organization's name as the source filed it. | [optional]
**organization_org_id** | Option<**String**> | Legal-entity level: the organization record of the organization. | [optional]
**organization_company_id** | Option<**String**> | Company level: the company the organization belongs to (null when unresolved). | [optional]
**organization_company_name** | Option<**String**> | Display name of the organization's company. | [optional]
**organization_group_id** | Option<**String**> | Group level: the organization's current ultimate parent company; a root company's group is itself. | [optional]
**organization_group_name** | Option<**String**> | Display name of the organization's group. | [optional]
**declared_function** | Option<**String**> | The function the source states for the organization. | [optional]
**evidence_type** | Option<**String**> | The kind of evidence the source gives. | [optional]
**authority_class** | Option<**String**> | How authoritative the source is for the claim. | [optional]
**source_key** | Option<**String**> | The source the assertion comes from. | [optional]
**source_record_id** | Option<**String**> | The source's own record id. | [optional]
**source_url** | Option<**String**> | A link to the source record. | [optional]
**observed_at** | Option<**String**> | When the assertion was observed (UTC). | [optional]
**last_verified_at** | Option<**String**> | When the assertion was last verified (UTC). | [optional]
**confidence** | Option<**f64**> | The source's or the extraction's own confidence; null when none is stated. | [optional]
**license_name** | Option<**String**> | The licence of the source. | [optional]
**attribution** | Option<**String**> | The credit line for the source. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


