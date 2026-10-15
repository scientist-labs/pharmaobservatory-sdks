# PharmaObservatory::ApiV1InvestorProfileGet200ResponseProfile

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **canonical_investor_id** | **String** |  |  |
| **legal_name** | **String** |  |  |
| **investor_type** | **String** |  |  |
| **source** | **String** |  |  |
| **country** | **String** |  |  |
| **raum_usd** | **Float** |  |  |
| **raum_band** | **String** |  |  |
| **is_era** | **Boolean** |  |  |
| **is_life_sciences_focused** | **Boolean** |  |  |
| **canonical_org_id** | **String** |  |  |
| **private_fund_count** | **Integer** |  |  |
| **total_private_fund_gav** | **Float** |  |  |
| **portfolio_co_count** | **Integer** |  |  |
| **portfolio_co_count_note** | **String** |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1InvestorProfileGet200ResponseProfile.new(
  canonical_investor_id: null,
  legal_name: null,
  investor_type: null,
  source: null,
  country: null,
  raum_usd: null,
  raum_band: null,
  is_era: null,
  is_life_sciences_focused: null,
  canonical_org_id: null,
  private_fund_count: null,
  total_private_fund_gav: null,
  portfolio_co_count: null,
  portfolio_co_count_note: null
)
```

