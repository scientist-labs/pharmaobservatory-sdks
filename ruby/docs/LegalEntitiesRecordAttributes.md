# PharmaObservatory::LegalEntitiesRecordAttributes

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **legal_entity_sec_filings** | [**LegalEntitySecFilingsRow**](LegalEntitySecFilingsRow.md) |  | [optional] |
| **ownership_percentages** | [**OwnershipPercentagesRow**](OwnershipPercentagesRow.md) |  | [optional] |
| **sbti_crosscheck** | [**SbtiCrosscheckRow**](SbtiCrosscheckRow.md) |  | [optional] |
| **sponsor_delivery** | [**SponsorDeliveryRow**](SponsorDeliveryRow.md) |  | [optional] |
| **sponsor_recruitment_duration** | [**SponsorRecruitmentDurationRow**](SponsorRecruitmentDurationRow.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::LegalEntitiesRecordAttributes.new(
  legal_entity_sec_filings: null,
  ownership_percentages: null,
  sbti_crosscheck: null,
  sponsor_delivery: null,
  sponsor_recruitment_duration: null
)
```

