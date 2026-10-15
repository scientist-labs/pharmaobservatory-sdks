# PharmaObservatory::ApiV1OrgFinancingAndOwnershipGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **org_name** | **String** |  | [optional] |
| **resolution** | [**ApiV1InvestorProfileGet200ResponseResolution**](ApiV1InvestorProfileGet200ResponseResolution.md) |  | [optional] |
| **org_id** | **String** |  |  |
| **requested_org_id** | **String** |  |  |
| **rolled_up_to_parent** | **Boolean** |  |  |
| **financing** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing**](ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing.md) |  |  |
| **investors** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseInvestors**](ApiV1OrgFinancingAndOwnershipGet200ResponseInvestors.md) |  |  |
| **ownership_positions** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPositions**](ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPositions.md) |  |  |
| **semantics** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseSemantics**](ApiV1OrgFinancingAndOwnershipGet200ResponseSemantics.md) |  |  |
| **build_state** | [**ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState**](ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState.md) |  |  |
| **sections** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseSections**](ApiV1OrgFinancingAndOwnershipGet200ResponseSections.md) |  |  |
| **error** | **String** |  | [optional] |
| **message** | **String** |  | [optional] |
| **ownership_percentages** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages**](ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages.md) |  | [optional] |
| **name_withheld_note** | **String** |  | [optional] |
| **unmatched_product_revenue** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseUnmatchedProductRevenue**](ApiV1OrgFinancingAndOwnershipGet200ResponseUnmatchedProductRevenue.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1OrgFinancingAndOwnershipGet200Response.new(
  org_name: null,
  resolution: null,
  org_id: null,
  requested_org_id: null,
  rolled_up_to_parent: null,
  financing: null,
  investors: null,
  ownership_positions: null,
  semantics: null,
  build_state: null,
  sections: null,
  error: null,
  message: null,
  ownership_percentages: null,
  name_withheld_note: null,
  unmatched_product_revenue: null
)
```

