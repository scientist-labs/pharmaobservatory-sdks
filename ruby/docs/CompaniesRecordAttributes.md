# PharmaObservatory::CompaniesRecordAttributes

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_firmographics** | [**CompanyFirmographicsRow**](CompanyFirmographicsRow.md) |  | [optional] |
| **company_funding_summary** | [**CompanyFundingSummaryRow**](CompanyFundingSummaryRow.md) |  | [optional] |
| **company_peer_profile** | [**CompanyPeerProfileRow**](CompanyPeerProfileRow.md) |  | [optional] |
| **company_trial_geography** | [**CompanyTrialGeographyRow**](CompanyTrialGeographyRow.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompaniesRecordAttributes.new(
  company_firmographics: null,
  company_funding_summary: null,
  company_peer_profile: null,
  company_trial_geography: null
)
```

