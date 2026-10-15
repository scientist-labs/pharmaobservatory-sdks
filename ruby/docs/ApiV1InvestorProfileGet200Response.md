# PharmaObservatory::ApiV1InvestorProfileGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **error** | **String** |  | [optional] |
| **message** | **String** |  | [optional] |
| **investor_id** | **String** |  |  |
| **profile** | [**ApiV1InvestorProfileGet200ResponseProfile**](ApiV1InvestorProfileGet200ResponseProfile.md) |  |  |
| **funds** | [**ApiV1InvestorProfileGet200ResponseFunds**](ApiV1InvestorProfileGet200ResponseFunds.md) |  |  |
| **ownership** | [**ApiV1InvestorProfileGet200ResponseOwnership**](ApiV1InvestorProfileGet200ResponseOwnership.md) |  |  |
| **semantics** | [**ApiV1InvestorProfileGet200ResponseSemantics**](ApiV1InvestorProfileGet200ResponseSemantics.md) |  |  |
| **build_state** | [**ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState**](ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState.md) |  |  |
| **sections** | [**ApiV1InvestorProfileGet200ResponseSections**](ApiV1InvestorProfileGet200ResponseSections.md) |  |  |
| **investor_name** | **String** |  | [optional] |
| **resolution** | [**ApiV1InvestorProfileGet200ResponseResolution**](ApiV1InvestorProfileGet200ResponseResolution.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1InvestorProfileGet200Response.new(
  error: null,
  message: null,
  investor_id: null,
  profile: null,
  funds: null,
  ownership: null,
  semantics: null,
  build_state: null,
  sections: null,
  investor_name: null,
  resolution: null
)
```

