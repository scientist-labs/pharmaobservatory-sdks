# PharmaObservatory::ApiV1ProgramCatalystsGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **drug_ids** | **Array&lt;String&gt;** |  |  |
| **company_ids** | **Array&lt;String&gt;** |  |  |
| **calendar** | [**Array&lt;ApiV1ProgramCatalystsGet200ResponseCalendarInner&gt;**](ApiV1ProgramCatalystsGet200ResponseCalendarInner.md) |  |  |
| **pdufa** | [**ApiV1ProgramCatalystsGet200ResponsePdufa**](ApiV1ProgramCatalystsGet200ResponsePdufa.md) |  |  |
| **build_state** | [**ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState**](ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState.md) |  |  |
| **sections** | [**ApiV1ProgramCatalystsGet200ResponseSections**](ApiV1ProgramCatalystsGet200ResponseSections.md) |  |  |
| **error** | **String** |  | [optional] |
| **message** | **String** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1ProgramCatalystsGet200Response.new(
  drug_ids: null,
  company_ids: null,
  calendar: null,
  pdufa: null,
  build_state: null,
  sections: null,
  error: null,
  message: null
)
```

