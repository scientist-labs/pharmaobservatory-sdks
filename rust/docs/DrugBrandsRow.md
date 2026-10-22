# DrugBrandsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**brand_key** | Option<**String**> | The brand key: an upper-case alphanumeric trade-name token (the brand's leading name tokens), the same key the brands pages use. Not a display name; brand_name is the display form. | [optional]
**brand_name** | Option<**String**> | The brand's display name (the brand key in display case). | [optional]
**molecule_name** | Option<**String**> | The display name of the drug the brand is a brand of. | [optional]
**approved_anywhere** | Option<**bool**> | True when the drug is approved in at least one jurisdiction. | [optional]
**is_marketed** | Option<**bool**> | True when the drug, or a record folded under it, has a package in the US National Drug Code (NDC) directory. | [optional]
**atc_class** | Option<**String**> | The top-level Anatomical Therapeutic Chemical (ATC) class of the drug. | [optional]
**is_biologic** | Option<**bool**> | True when the drug is a biologic. | [optional]
**first_approval_year** | Option<**i32**> | The earliest approval year ChEMBL records for the drug. | [optional]
**holder_name** | Option<**String**> | The FDA approval holder's name as filed, present only when the brand has exactly one holder name. | [optional]
**holder_org_id** | Option<**String**> | Legal-entity level: the approval holder's organization record (null when the brand has no holder, more than one, or an unresolved one). | [optional]
**holder_company_id** | Option<**String**> | Company level: the approval holder's company (null when unresolved). | [optional]
**holder_company_name** | Option<**String**> | Display name of the approval holder's company. | [optional]
**holder_group_id** | Option<**String**> | Group level: the approval holder's current ultimate parent company; a root company's group is itself. | [optional]
**holder_group_name** | Option<**String**> | Display name of the approval holder's group. | [optional]
**holder_match_scope** | Option<**String**> | How firmly the holder name resolved to the organization (spine_unique, anchored and establishment_registry are folded onto an organization the spine already held; anchored_scoped, registry_scoped, alias_scoped and none are newly minted organizations). Null when the brand has no single holder organization. | [optional]
**n_holder_names** | Option<**i32**> | Distinct approval-holder names FDA publishes for the brand (0 for a brand with no FDA applicant, an RxNorm-only brand). | [optional]
**n_holder_orgs** | Option<**i32**> | Distinct organizations those names resolve to; above 1 with a null holder_org_id means more than one company holds the brand. | [optional]
**holder_as_of** | Option<[**String**](string.md)> | The ingest snapshot date the holder was read from. | [optional]
**headline_indication** | Option<**String**> | The brand's headline indication: the one most often named in the filings that disclose its revenue. | [optional]
**latest_revenue_usd_millions** | Option<**f64**> | The brand's latest disclosed worldwide revenue in USD millions (null when none is disclosed). | [optional]
**latest_revenue_year** | Option<**i32**> | The fiscal year of latest_revenue_usd_millions. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


