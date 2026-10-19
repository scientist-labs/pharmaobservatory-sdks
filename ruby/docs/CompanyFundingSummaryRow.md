# PharmaObservatory::CompanyFundingSummaryRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | The company key (company:&lt;key&gt;). | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **n_rounds** | **Integer** | Funding events of the company, disclosed or not. | [optional] |
| **total_non_dilutive_usd** | **Float** | Disclosed non-dilutive amounts (grants, debt, royalty and similar), in whole US dollars. | [optional] |
| **total_dilutive_usd** | **Float** | Disclosed dilutive amounts (equity-issuing rounds), in whole US dollars. | [optional] |
| **total_unclassified_usd** | **Float** | Disclosed amounts of events whose instrument is not classified, kept apart from the other two totals. | [optional] |
| **n_unclassified_rounds** | **Integer** | Events whose instrument is not classified. | [optional] |
| **total_raised_usd** | **Float** | All disclosed amounts, the sum of the three totals, in whole US dollars. | [optional] |
| **first_funded_date** | **Date** | Date of the earliest event. | [optional] |
| **last_funded_date** | **Date** | Date of the latest event. | [optional] |
| **latest_round_type** | **String** | The round type of the latest event. | [optional] |
| **latest_round_date** | **Date** | The date of the latest event. | [optional] |
| **stage_label** | **String** | A coarse reading of the round types: public, growth_private, seed_private, unclassified or non_dilutive_only. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyFundingSummaryRow.new(
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  n_rounds: null,
  total_non_dilutive_usd: null,
  total_dilutive_usd: null,
  total_unclassified_usd: null,
  n_unclassified_rounds: null,
  total_raised_usd: null,
  first_funded_date: null,
  last_funded_date: null,
  latest_round_type: null,
  latest_round_date: null,
  stage_label: null
)
```

