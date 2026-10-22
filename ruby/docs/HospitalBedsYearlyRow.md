# PharmaObservatory::HospitalBedsYearlyRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ccn** | **String** | CMS Certification Number of the hospital. | [optional] |
| **year** | **Integer** | The calendar year. | [optional] |
| **bed_count** | **Integer** | Certified bed count that year. | [optional] |
| **is_teaching** | **Boolean** | Whether the hospital was a teaching hospital that year. | [optional] |
| **pos_vintage** | **String** | The Provider of Services file vintage the row comes from. | [optional] |
| **as_of** | **Date** | The as-of date of that file. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::HospitalBedsYearlyRow.new(
  ccn: null,
  year: null,
  bed_count: null,
  is_teaching: null,
  pos_vintage: null,
  as_of: null
)
```

