# BrandCompaniesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**brand_key** | Option<**String**> | The brand key: an upper-case alphanumeric trade-name token (the brand's leading name tokens), the same key the brands pages use. Not a display name; brand_name is the display form. | [optional]
**company_role** | Option<**String**> | The applicant's role for the brand: approval_holder or generic_applicant. | [optional]
**applicant_name_key** | Option<**String**> | The normalised applicant name, the grain key that separates two spellings of one applicant. | [optional]
**applicant_name** | Option<**String**> | The applicant name FDA published, as filed. | [optional]
**applicant_org_id** | Option<**String**> | Legal-entity level: the applicant's organization record (null when unresolved). | [optional]
**applicant_company_id** | Option<**String**> | Company level: the applicant's company (null when unresolved). | [optional]
**applicant_company_name** | Option<**String**> | Display name of the applicant's company. | [optional]
**applicant_group_id** | Option<**String**> | Group level: the applicant's current ultimate parent company; a root company's group is itself. | [optional]
**applicant_group_name** | Option<**String**> | Display name of the applicant's group. | [optional]
**match_scope** | Option<**String**> | How firmly the applicant name resolved to the organization (spine_unique, anchored and establishment_registry are folded onto an organization the spine already held; anchored_scoped, registry_scoped, alias_scoped and none are newly minted organizations). Null when the applicant did not resolve. | [optional]
**n_applications** | Option<**i32**> | Distinct FDA applications the applicant holds for the brand. | [optional]
**first_approval_date** | Option<[**String**](string.md)> | The earliest approval date among those applications. | [optional]
**latest_approval_date** | Option<[**String**](string.md)> | The latest approval date among those applications. | [optional]
**is_marketed_now** | Option<**bool**> | True when at least one of the applicant's products is currently marketed. | [optional]
**source_systems** | Option<**Vec<String>**> | The FDA records the applicant was read from (Drugs@FDA, Orange Book, Purple Book). | [optional]
**as_of** | Option<[**String**](string.md)> | The ingest snapshot date the row was read from. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


