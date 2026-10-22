# PharmaObservatory::LegalEntitiesRecordAttributes

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ownership_percentages** | [**OwnershipPercentagesRow**](OwnershipPercentagesRow.md) |  | [optional] |
| **sponsor_recruitment_duration** | [**SponsorRecruitmentDurationRow**](SponsorRecruitmentDurationRow.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::LegalEntitiesRecordAttributes.new(
  ownership_percentages: null,
  sponsor_recruitment_duration: null
)
```

