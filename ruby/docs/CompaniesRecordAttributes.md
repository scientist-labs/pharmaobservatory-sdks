# PharmaObservatory::CompaniesRecordAttributes

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_funding_summary** | [**CompanyFundingSummaryRow**](CompanyFundingSummaryRow.md) |  | [optional] |
| **company_trial_geography** | [**CompanyTrialGeographyRow**](CompanyTrialGeographyRow.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompaniesRecordAttributes.new(
  company_funding_summary: null,
  company_trial_geography: null
)
```

