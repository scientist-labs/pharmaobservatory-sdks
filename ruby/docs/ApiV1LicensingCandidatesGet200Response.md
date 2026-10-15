# PharmaObservatory::ApiV1LicensingCandidatesGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **condition_id** | **String** |  |  |
| **condition_name** | **String** |  |  |
| **candidates** | [**Array&lt;ApiV1LicensingCandidatesGet200ResponseCandidatesInner&gt;**](ApiV1LicensingCandidatesGet200ResponseCandidatesInner.md) |  |  |
| **build_state** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **restriction** | [**ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction**](ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1LicensingCandidatesGet200Response.new(
  condition_id: null,
  condition_name: null,
  candidates: null,
  build_state: null,
  restriction: null
)
```

