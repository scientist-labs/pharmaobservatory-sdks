# PharmaObservatory::ApiV1RwdConductorTrackRecordGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **institution_name** | **String** |  |  |
| **hma_ema_id** | **String** |  |  |
| **canonical_org_id** | **String** |  |  |
| **found** | **Boolean** |  |  |
| **candidates** | **Array&lt;Object&gt;** |  |  |
| **conductor** | [**ApiV1RwdConductorTrackRecordGet200ResponseConductor**](ApiV1RwdConductorTrackRecordGet200ResponseConductor.md) |  |  |
| **delivery_by_mandate** | [**ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandate**](ApiV1RwdConductorTrackRecordGet200ResponseDeliveryByMandate.md) |  |  |
| **footprint** | [**ApiV1RwdConductorTrackRecordGet200ResponseFootprint**](ApiV1RwdConductorTrackRecordGet200ResponseFootprint.md) |  |  |
| **coverage_caveat** | **String** |  |  |
| **tier_provisional_note** | **String** |  |  |
| **build_state** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **restriction** | [**ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction**](ApiV1ConditionsConditionIdBurdenVsEvidenceGet200ResponseRestriction.md) |  |  |
| **resolution_confidence** | [**ApiV1ResolveEntityGet200ResponseResolutionConfidence**](ApiV1ResolveEntityGet200ResponseResolutionConfidence.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1RwdConductorTrackRecordGet200Response.new(
  institution_name: null,
  hma_ema_id: null,
  canonical_org_id: null,
  found: null,
  candidates: null,
  conductor: null,
  delivery_by_mandate: null,
  footprint: null,
  coverage_caveat: null,
  tier_provisional_note: null,
  build_state: null,
  restriction: null,
  resolution_confidence: null
)
```

