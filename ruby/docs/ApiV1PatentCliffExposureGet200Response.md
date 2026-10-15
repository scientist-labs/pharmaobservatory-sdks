# PharmaObservatory::ApiV1PatentCliffExposureGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **drug_ids** | **Array&lt;String&gt;** |  |  |
| **company_ids** | **Array&lt;String&gt;** |  |  |
| **exposures** | [**Array&lt;ApiV1PatentCliffExposureGet200ResponseExposuresInner&gt;**](ApiV1PatentCliffExposureGet200ResponseExposuresInner.md) |  |  |
| **build_state** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **error** | **String** |  | [optional] |
| **message** | **String** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1PatentCliffExposureGet200Response.new(
  drug_ids: null,
  company_ids: null,
  exposures: null,
  build_state: null,
  error: null,
  message: null
)
```

