# CompanyMembersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**company_id** | Option<**String**> | The company the legal entity is a member of. | [optional]
**company_name** | Option<**String**> | The company's display name. | [optional]
**legal_entity_id** | Option<**String**> | A registry-scoped legal entity id (edgar:, ror:, lei: or rx_provider:). A registration, not an organization record id. | [optional]
**link_basis** | Option<**String**> | The source that placed the legal entity in the company (gleif, edgar, ror, name_key and so on). | [optional]
**parent_company_id** | Option<**String**> | The company's direct owner, null for a root. | [optional]
**parent_company_name** | Option<**String**> | Display name of the direct owner. | [optional]
**parent_link_basis** | Option<**String**> | The authority of the first ownership hop. | [optional]
**since** | Option<[**String**](string.md)> | Date the ownership link began, when every hop is dated. | [optional]
**since_precision** | Option<**String**> | Precision of since. | [optional]
**since_source** | Option<**String**> | Where the date came from. | [optional]
**group_id** | Option<**String**> | Group level: the current ultimate parent; a root's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]
**hop_depth** | Option<**i32**> | Hops from the company to its ultimate parent (0 for a root). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


