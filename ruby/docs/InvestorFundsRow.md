# PharmaObservatory::InvestorFundsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **investor_id** | **String** | The investor key. | [optional] |
| **fund_type** | **String** | The Form ADV Item 7.B fund type: hedge, liquidity, pe (private equity), real_estate, securitized, vc (venture capital) or other. | [optional] |
| **fund_count** | **Integer** | Funds of this type the adviser reports. | [optional] |
| **total_private_fund_gav** | **Float** | The adviser&#39;s firm-level total private-fund gross asset value in whole US dollars, repeated on every row of the investor. Do not add it across rows. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestorFundsRow.new(
  investor_id: null,
  fund_type: null,
  fund_count: null,
  total_private_fund_gav: null
)
```

