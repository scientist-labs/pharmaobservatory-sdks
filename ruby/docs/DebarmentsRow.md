# PharmaObservatory::DebarmentsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **debarment_id** | **String** | The debarment id. | [optional] |
| **firm_name** | **String** | The firm&#39;s name as filed with the regulator (the debarred entity). | [optional] |
| **firm_org_id** | **String** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] |
| **firm_company_id** | **String** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] |
| **firm_company_name** | **String** | Display name of the firm&#39;s company. | [optional] |
| **firm_group_id** | **String** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **firm_group_name** | **String** | Display name of the firm&#39;s group. | [optional] |
| **is_org_resolved** | **Boolean** | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional] |
| **match_tier** | **String** | How well the firm name was attached to an organization: high, medium or low name-and-geography matches, registry (the operating firm of a registered facility), ambiguous (several candidates, none chosen) or unresolved. Warning letters, untitled letters and debarments attach by firm name alone. | [optional] |
| **effective_date** | **Date** | The date the debarment took effect. | [optional] |
| **term** | **String** | The debarment term (for example permanent or a number of years). | [optional] |
| **fr_citation** | **String** | The Federal Register citation of the notice. | [optional] |
| **fr_date** | **Date** | The Federal Register publication date. | [optional] |
| **source_url** | **String** | The FDA debarment list page. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DebarmentsRow.new(
  debarment_id: null,
  firm_name: null,
  firm_org_id: null,
  firm_company_id: null,
  firm_company_name: null,
  firm_group_id: null,
  firm_group_name: null,
  is_org_resolved: null,
  match_tier: null,
  effective_date: null,
  term: null,
  fr_citation: null,
  fr_date: null,
  source_url: null
)
```

