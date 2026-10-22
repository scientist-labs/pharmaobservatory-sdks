# PharmaObservatory::DrugsRecordAttributes

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **drug_label_safety** | [**DrugLabelSafetyRow**](DrugLabelSafetyRow.md) |  | [optional] |
| **drug_regulatory** | [**DrugRegulatoryRow**](DrugRegulatoryRow.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugsRecordAttributes.new(
  drug_label_safety: null,
  drug_regulatory: null
)
```

