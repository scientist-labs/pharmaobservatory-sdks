# PharmaObservatory::InstitutionalHoldingsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **holding_id** | **String** | The holding line id. | [optional] |
| **issuer_name** | **String** | The security issuer as filed. | [optional] |
| **issuer_org_id** | **String** | Legal-entity level: the issuer&#39;s organization record (null when unresolved). | [optional] |
| **issuer_company_id** | **String** | Company level: the issuer&#39;s company (null when unresolved). | [optional] |
| **issuer_company_name** | **String** | Display name of the issuer&#39;s company. | [optional] |
| **issuer_group_id** | **String** | Group level: the issuer&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **issuer_group_name** | **String** | Display name of the issuer&#39;s group. | [optional] |
| **security_ref** | **String** | The security&#39;s CUSIP reference. | [optional] |
| **holder_name** | **String** | The institutional manager&#39;s name. | [optional] |
| **holder_kind** | **String** | Manager kind. | [optional] |
| **holder_org_id** | **String** | Legal-entity level: the manager&#39;s organization record (null when unresolved). | [optional] |
| **holder_company_id** | **String** | Company level: the manager&#39;s company (null when unresolved). | [optional] |
| **holder_company_name** | **String** | Display name of the manager&#39;s company. | [optional] |
| **holder_group_id** | **String** | Group level: the manager&#39;s current ultimate parent. | [optional] |
| **holder_group_name** | **String** | Display name of the manager&#39;s group. | [optional] |
| **reported_value_amount** | **Float** | Reported value of the position. | [optional] |
| **reported_value_currency** | **String** | Currency of the reported value. | [optional] |
| **shares** | **Float** | Shares standing. | [optional] |
| **percent_basis** | **String** | The basis the percent is measured on. | [optional] |
| **reported_percent** | **Float** | Reported share of the issuer. | [optional] |
| **edge_state** | **String** | State of the position (standing, ceased, no longer reported). | [optional] |
| **first_observed_date** | **Date** | First filing date the position was observed. | [optional] |
| **last_filing_date** | **Date** | Latest filing date the position was reported. | [optional] |
| **contains_personal_data** | **Boolean** | Whether the holder is a natural person. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InstitutionalHoldingsRow.new(
  holding_id: null,
  issuer_name: null,
  issuer_org_id: null,
  issuer_company_id: null,
  issuer_company_name: null,
  issuer_group_id: null,
  issuer_group_name: null,
  security_ref: null,
  holder_name: null,
  holder_kind: null,
  holder_org_id: null,
  holder_company_id: null,
  holder_company_name: null,
  holder_group_id: null,
  holder_group_name: null,
  reported_value_amount: null,
  reported_value_currency: null,
  shares: null,
  percent_basis: null,
  reported_percent: null,
  edge_state: null,
  first_observed_date: null,
  last_filing_date: null,
  contains_personal_data: null
)
```

