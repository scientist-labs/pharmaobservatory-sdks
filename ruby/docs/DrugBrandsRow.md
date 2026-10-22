# PharmaObservatory::DrugBrandsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **brand_key** | **String** | The brand key: an upper-case alphanumeric trade-name token (the brand&#39;s leading name tokens), the same key the brands pages use. Not a display name; brand_name is the display form. | [optional] |
| **brand_name** | **String** | The brand&#39;s display name (the brand key in display case). | [optional] |
| **molecule_name** | **String** | The display name of the drug the brand is a brand of. | [optional] |
| **approved_anywhere** | **Boolean** | True when the drug is approved in at least one jurisdiction. | [optional] |
| **is_marketed** | **Boolean** | True when the drug, or a record folded under it, has a package in the US National Drug Code (NDC) directory. | [optional] |
| **atc_class** | **String** | The top-level Anatomical Therapeutic Chemical (ATC) class of the drug. | [optional] |
| **is_biologic** | **Boolean** | True when the drug is a biologic. | [optional] |
| **first_approval_year** | **Integer** | The earliest approval year ChEMBL records for the drug. | [optional] |
| **holder_name** | **String** | The FDA approval holder&#39;s name as filed, present only when the brand has exactly one holder name. | [optional] |
| **holder_org_id** | **String** | Legal-entity level: the approval holder&#39;s organization record (null when the brand has no holder, more than one, or an unresolved one). | [optional] |
| **holder_company_id** | **String** | Company level: the approval holder&#39;s company (null when unresolved). | [optional] |
| **holder_company_name** | **String** | Display name of the approval holder&#39;s company. | [optional] |
| **holder_group_id** | **String** | Group level: the approval holder&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **holder_group_name** | **String** | Display name of the approval holder&#39;s group. | [optional] |
| **holder_match_scope** | **String** | How firmly the holder name resolved to the organization (spine_unique, anchored and establishment_registry are folded onto an organization the spine already held; anchored_scoped, registry_scoped, alias_scoped and none are newly minted organizations). Null when the brand has no single holder organization. | [optional] |
| **n_holder_names** | **Integer** | Distinct approval-holder names FDA publishes for the brand (0 for a brand with no FDA applicant, an RxNorm-only brand). | [optional] |
| **n_holder_orgs** | **Integer** | Distinct organizations those names resolve to; above 1 with a null holder_org_id means more than one company holds the brand. | [optional] |
| **holder_as_of** | **Date** | The ingest snapshot date the holder was read from. | [optional] |
| **headline_indication** | **String** | The brand&#39;s headline indication: the one most often named in the filings that disclose its revenue. | [optional] |
| **latest_revenue_usd_millions** | **Float** | The brand&#39;s latest disclosed worldwide revenue in USD millions (null when none is disclosed). | [optional] |
| **latest_revenue_year** | **Integer** | The fiscal year of latest_revenue_usd_millions. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugBrandsRow.new(
  molecule_id: null,
  brand_key: null,
  brand_name: null,
  molecule_name: null,
  approved_anywhere: null,
  is_marketed: null,
  atc_class: null,
  is_biologic: null,
  first_approval_year: null,
  holder_name: null,
  holder_org_id: null,
  holder_company_id: null,
  holder_company_name: null,
  holder_group_id: null,
  holder_group_name: null,
  holder_match_scope: null,
  n_holder_names: null,
  n_holder_orgs: null,
  holder_as_of: null,
  headline_indication: null,
  latest_revenue_usd_millions: null,
  latest_revenue_year: null
)
```

