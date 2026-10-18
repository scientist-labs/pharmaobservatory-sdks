# PharmaObservatory::OrgsDiversity200ResponseRowsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **reporting_year** | **String** |  |  |
| **workforce_pct_women** | **Float** |  |  |
| **board_pct_women** | **Float** |  |  |
| **canonical_org_id** | **String** |  | [optional] |
| **reporting_company** | **String** |  | [optional] |
| **cik** | **String** |  | [optional] |
| **ticker** | **String** |  | [optional] |
| **leadership_pct_women** | **Float** |  | [optional] |
| **workforce_pct_underrepresented** | **Object** |  | [optional] |
| **gender_pay_gap_pct** | **Object** |  | [optional] |
| **gender_pay_gap_basis** | **String** |  | [optional] |
| **has_dei_targets** | **Boolean** |  | [optional] |
| **dei_target_text** | **String** |  | [optional] |
| **external_recognition_claimed** | **String** |  | [optional] |
| **document_id** | **String** |  | [optional] |
| **evidence_quote** | **String** |  | [optional] |
| **confidence** | **Float** |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsDiversity200ResponseRowsInner.new(
  reporting_year: null,
  workforce_pct_women: null,
  board_pct_women: null,
  canonical_org_id: null,
  reporting_company: null,
  cik: null,
  ticker: null,
  leadership_pct_women: null,
  workforce_pct_underrepresented: null,
  gender_pay_gap_pct: null,
  gender_pay_gap_basis: null,
  has_dei_targets: null,
  dei_target_text: null,
  external_recognition_claimed: null,
  document_id: null,
  evidence_quote: null,
  confidence: null
)
```

