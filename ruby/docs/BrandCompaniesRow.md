# PharmaObservatory::BrandCompaniesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **brand_key** | **String** | The brand key: an upper-case alphanumeric trade-name token (the brand&#39;s leading name tokens), the same key the brands pages use. Not a display name; brand_name is the display form. | [optional] |
| **company_role** | **String** | The applicant&#39;s role for the brand: approval_holder or generic_applicant. | [optional] |
| **applicant_name_key** | **String** | The normalised applicant name, the grain key that separates two spellings of one applicant. | [optional] |
| **applicant_name** | **String** | The applicant name FDA published, as filed. | [optional] |
| **applicant_org_id** | **String** | Legal-entity level: the applicant&#39;s organization record (null when unresolved). | [optional] |
| **applicant_company_id** | **String** | Company level: the applicant&#39;s company (null when unresolved). | [optional] |
| **applicant_company_name** | **String** | Display name of the applicant&#39;s company. | [optional] |
| **applicant_group_id** | **String** | Group level: the applicant&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **applicant_group_name** | **String** | Display name of the applicant&#39;s group. | [optional] |
| **match_scope** | **String** | How firmly the applicant name resolved to the organization (spine_unique, anchored and establishment_registry are folded onto an organization the spine already held; anchored_scoped, registry_scoped, alias_scoped and none are newly minted organizations). Null when the applicant did not resolve. | [optional] |
| **n_applications** | **Integer** | Distinct FDA applications the applicant holds for the brand. | [optional] |
| **first_approval_date** | **Date** | The earliest approval date among those applications. | [optional] |
| **latest_approval_date** | **Date** | The latest approval date among those applications. | [optional] |
| **is_marketed_now** | **Boolean** | True when at least one of the applicant&#39;s products is currently marketed. | [optional] |
| **source_systems** | **Array&lt;String&gt;** | The FDA records the applicant was read from (Drugs@FDA, Orange Book, Purple Book). | [optional] |
| **as_of** | **Date** | The ingest snapshot date the row was read from. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BrandCompaniesRow.new(
  molecule_id: null,
  brand_key: null,
  company_role: null,
  applicant_name_key: null,
  applicant_name: null,
  applicant_org_id: null,
  applicant_company_id: null,
  applicant_company_name: null,
  applicant_group_id: null,
  applicant_group_name: null,
  match_scope: null,
  n_applications: null,
  first_approval_date: null,
  latest_approval_date: null,
  is_marketed_now: null,
  source_systems: null,
  as_of: null
)
```

