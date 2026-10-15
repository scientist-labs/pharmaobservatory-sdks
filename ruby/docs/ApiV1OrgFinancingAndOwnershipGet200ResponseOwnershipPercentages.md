# PharmaObservatory::ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **org_id** | **String** |  |  |
| **requested_org_id** | **String** |  |  |
| **any** | **Boolean** |  |  |
| **shares_outstanding** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesSharesOutstanding**](ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesSharesOutstanding.md) |  |  |
| **insider** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInsider**](ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInsider.md) |  |  |
| **institutional** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInstitutional**](ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesInstitutional.md) |  |  |
| **notes** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesNotes**](ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentagesNotes.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages.new(
  org_id: null,
  requested_org_id: null,
  any: null,
  shares_outstanding: null,
  insider: null,
  institutional: null,
  notes: null
)
```

