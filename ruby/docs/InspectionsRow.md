# PharmaObservatory::InspectionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **inspection_id** | **String** | The inspection id. | [optional] |
| **fei** | **String** | The FDA establishment identifier (FEI) of the inspected site. | [optional] |
| **firm_name** | **String** | The firm&#39;s name as filed with the regulator (the inspected establishment&#39;s firm). | [optional] |
| **firm_org_id** | **String** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] |
| **firm_company_id** | **String** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] |
| **firm_company_name** | **String** | Display name of the firm&#39;s company. | [optional] |
| **firm_group_id** | **String** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **firm_group_name** | **String** | Display name of the firm&#39;s group. | [optional] |
| **is_org_resolved** | **Boolean** | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional] |
| **match_tier** | **String** | How well the firm was attached to an organization: high, medium, low, registry, ambiguous or unresolved. | [optional] |
| **org_resolution_method** | **String** | The route that attached the firm: fei_bridge (the FDA establishment identifier, FEI, bridge), fei_firm_mint or fei_firm_fold (the operating firm&#39;s registry record), firm_name_geo (name and geography), firm_name_only (name alone) or unresolved. | [optional] |
| **inspection_date** | **Date** | The date the inspection ended. | [optional] |
| **fiscal_year** | **String** | The federal fiscal year of the inspection. | [optional] |
| **classification** | **String** | The outcome: NAI (no action indicated), VAI (voluntary action indicated) or OAI (official action indicated). | [optional] |
| **had_483** | **Boolean** | True when the inspection ended with a Form 483 of observations. | [optional] |
| **product_type** | **String** | The product type inspected (Drugs, Devices, Biologics, Veterinary, Tobacco or Food/Cosmetics). | [optional] |
| **project_area** | **String** | The FDA project area of the inspection. | [optional] |
| **city** | **String** | The establishment&#39;s city. | [optional] |
| **state** | **String** | The establishment&#39;s state or province. | [optional] |
| **postal_code** | **String** | The establishment&#39;s postal code. | [optional] |
| **country_name** | **String** | The establishment&#39;s country name. | [optional] |
| **country_iso2** | **String** | The establishment&#39;s two-letter country code. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InspectionsRow.new(
  inspection_id: null,
  fei: null,
  firm_name: null,
  firm_org_id: null,
  firm_company_id: null,
  firm_company_name: null,
  firm_group_id: null,
  firm_group_name: null,
  is_org_resolved: null,
  match_tier: null,
  org_resolution_method: null,
  inspection_date: null,
  fiscal_year: null,
  classification: null,
  had_483: null,
  product_type: null,
  project_area: null,
  city: null,
  state: null,
  postal_code: null,
  country_name: null,
  country_iso2: null
)
```

