# PharmaObservatory::ImportRefusalsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **refusal_id** | **String** | The refusal line id. | [optional] |
| **fei** | **String** | The FDA establishment identifier (FEI) of the refused firm. | [optional] |
| **firm_name** | **String** | The firm&#39;s name as filed with the regulator (the firm named on the refusal). | [optional] |
| **firm_org_id** | **String** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] |
| **firm_company_id** | **String** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] |
| **firm_company_name** | **String** | Display name of the firm&#39;s company. | [optional] |
| **firm_group_id** | **String** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **firm_group_name** | **String** | Display name of the firm&#39;s group. | [optional] |
| **is_org_resolved** | **Boolean** | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional] |
| **match_tier** | **String** | How well the firm was attached to an organization: high, medium, low, registry, ambiguous or unresolved. | [optional] |
| **org_resolution_method** | **String** | The route that attached the firm: fei_bridge (the FDA establishment identifier, FEI, bridge), fei_firm_mint or fei_firm_fold (the operating firm&#39;s registry record), firm_name_geo (name and geography), firm_name_only (name alone) or unresolved. | [optional] |
| **refusal_date** | **Date** | The date the import was refused. | [optional] |
| **product_code** | **String** | The FDA product code of the refused product. | [optional] |
| **product_description** | **String** | The refused product as described in the report. | [optional] |
| **industry_code** | **String** | The FDA industry code of the product. | [optional] |
| **refusal_charges** | **String** | The charges (violations) the refusal cites. | [optional] |
| **city** | **String** | The firm&#39;s city. | [optional] |
| **state** | **String** | The firm&#39;s state or province. | [optional] |
| **postal_code** | **String** | The firm&#39;s postal code. | [optional] |
| **country_name** | **String** | The country the shipment came from. | [optional] |
| **country_code** | **String** | The FDA country code of the shipment. | [optional] |
| **country_iso2** | **String** | The two-letter country code of the shipment. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ImportRefusalsRow.new(
  refusal_id: null,
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
  refusal_date: null,
  product_code: null,
  product_description: null,
  industry_code: null,
  refusal_charges: null,
  city: null,
  state: null,
  postal_code: null,
  country_name: null,
  country_code: null,
  country_iso2: null
)
```

