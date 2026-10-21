# PharmaObservatory::InvestorPrivateFundsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **fund_id** | **String** | The fund key, advfund:&lt;PF-ID&gt;. | [optional] |
| **pf_id** | **String** | The SEC private fund identification number (PF-ID). | [optional] |
| **fund_name** | **String** | The fund&#39;s name as the adviser filed it. | [optional] |
| **fund_type** | **String** | The fund type the adviser reports on Form ADV Schedule D. | [optional] |
| **gross_asset_value** | **Float** | The fund&#39;s gross asset value in whole US dollars, as filed. | [optional] |
| **beneficial_owner_count** | **Integer** | The number of beneficial owners the fund reports. It counts the fund&#39;s own investors, not investors in this directory. | [optional] |
| **is_master** | **Boolean** | True when the fund is a master fund. | [optional] |
| **is_feeder** | **Boolean** | True when the fund is a feeder fund. | [optional] |
| **master_fund_name** | **String** | The master fund&#39;s name when this fund is a feeder. | [optional] |
| **master_pf_id** | **String** | The master fund&#39;s PF-ID when this fund is a feeder. | [optional] |
| **adviser_count** | **Integer** | Advisers that report this PF-ID. | [optional] |
| **investor_id** | **String** | The investor key of the adviser the fund is listed under. | [optional] |
| **as_of_date** | **Date** | The date of the filing the fund row comes from. | [optional] |
| **source_url** | **String** | A link to the adviser&#39;s Form ADV filing. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestorPrivateFundsRow.new(
  fund_id: null,
  pf_id: null,
  fund_name: null,
  fund_type: null,
  gross_asset_value: null,
  beneficial_owner_count: null,
  is_master: null,
  is_feeder: null,
  master_fund_name: null,
  master_pf_id: null,
  adviser_count: null,
  investor_id: null,
  as_of_date: null,
  source_url: null
)
```

