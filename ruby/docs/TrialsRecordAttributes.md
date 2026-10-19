# PharmaObservatory::TrialsRecordAttributes

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_delivery** | [**TrialDeliveryRow**](TrialDeliveryRow.md) |  | [optional] |
| **trial_demographics** | [**TrialDemographicsRow**](TrialDemographicsRow.md) |  | [optional] |
| **trial_enrollment_rate** | [**TrialEnrollmentRateRow**](TrialEnrollmentRateRow.md) |  | [optional] |
| **trial_extracted_endpoints** | [**TrialExtractedEndpointsRow**](TrialExtractedEndpointsRow.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialsRecordAttributes.new(
  trial_delivery: null,
  trial_demographics: null,
  trial_enrollment_rate: null,
  trial_extracted_endpoints: null
)
```

