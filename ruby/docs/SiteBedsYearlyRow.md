# PharmaObservatory::SiteBedsYearlyRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **site_id** | **String** | The site key: one facility spelling from the registry. | [optional] |
| **year** | **Integer** | The registry year. | [optional] |
| **ccn** | **String** | CMS Certification Number of the matched hospital. | [optional] |
| **bed_count** | **Integer** | Certified bed count that year. | [optional] |
| **is_teaching** | **Boolean** | Whether the hospital was a teaching hospital that year. | [optional] |
| **pos_vintage** | **String** | The Provider of Services file vintage the row comes from. | [optional] |
| **as_of** | **Date** | The as-of date of that file. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SiteBedsYearlyRow.new(
  site_id: null,
  year: null,
  ccn: null,
  bed_count: null,
  is_teaching: null,
  pos_vintage: null,
  as_of: null
)
```

