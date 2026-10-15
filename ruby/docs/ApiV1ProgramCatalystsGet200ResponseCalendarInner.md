# PharmaObservatory::ApiV1ProgramCatalystsGet200ResponseCalendarInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **event_type** | **String** |  |  |
| **id** | **String** |  |  |
| **date** | **String** |  |  |
| **nct_id** | **String** |  | [optional] |
| **title** | **String** |  |  |
| **phase** | **String** |  | [optional] |
| **matched_drug_ids** | **Array&lt;Object&gt;** |  | [optional] |
| **matched_company_ids** | **Array&lt;String&gt;** |  | [optional] |
| **canonical_drug_id** | **String** |  | [optional] |
| **canonical_org_id** | **String** |  | [optional] |
| **drug_name** | **String** |  | [optional] |
| **cliff_year** | **Integer** |  | [optional] |
| **cliff_basis** | **String** |  | [optional] |
| **org_name** | **String** |  | [optional] |
| **source** | **String** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1ProgramCatalystsGet200ResponseCalendarInner.new(
  event_type: null,
  id: null,
  date: null,
  nct_id: null,
  title: null,
  phase: null,
  matched_drug_ids: null,
  matched_company_ids: null,
  canonical_drug_id: null,
  canonical_org_id: null,
  drug_name: null,
  cliff_year: null,
  cliff_basis: null,
  org_name: null,
  source: null
)
```

