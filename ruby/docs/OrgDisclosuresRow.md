# PharmaObservatory::OrgDisclosuresRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **org_id** | **String** | Legal-entity level: the organization record. | [optional] |
| **org_name** | **String** | The organization&#39;s name. | [optional] |
| **company_id** | **String** | Company level: the company (null when unresolved). | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **claim_category** | **String** | The claim&#39;s domain: climate or diversity. | [optional] |
| **claim_type** | **String** | Which claim the row is about. Climate: net_zero_target, near_term_target, sbti_status_claimed, target_temperature_claimed. Diversity: board_pct_women, workforce_pct_women, leadership_pct_women, workforce_pct_underrepresented, gender_pay_gap_pct, has_dei_targets, external_recognition_claimed. | [optional] |
| **disclosure_source** | **String** | Where the claim comes from: sec_edgar_extraction (figures extracted from the company&#39;s own US Securities and Exchange Commission filings, open) or sbti (the Science Based Targets initiative dashboard, licensable). | [optional] |
| **disclosure_state** | **String** | disclosed: the source states a value or a positive claim. true_negative: the source states a real negative (false or none). not_disclosed: the source covers this organization and states no value for the claim. A claim withheld from a caller who lacks the licence is not a state: the row is absent for that caller and the response names the withheld tier. | [optional] |
| **claim_value** | **String** | The claimed value as text: true or false for a yes-or-no claim, a number for a percentage, the source&#39;s own word for a status. Null when disclosure_state is not_disclosed. | [optional] |
| **claim_detail** | **String** | A qualifier of the claim, such as the target year. Null when the source gives none. | [optional] |
| **disclosure_date** | **Date** | The date of the disclosure: the filing date for an extracted claim, the dashboard&#39;s last-updated date for an SBTi claim. Null when the source gives none. | [optional] |
| **reporting_year** | **String** | The reporting year the company states for an extracted claim. Null for an SBTi claim. | [optional] |
| **evidence_quote** | **String** | The sentence of the filing the figure was extracted from. Null for an SBTi claim. | [optional] |
| **confidence** | **Float** | The extraction confidence, 0 to 1. Null for an SBTi claim, which carries none. | [optional] |
| **source_record_id** | **String** | The source record: the filing document id for an extracted claim, the SBTi company id for an SBTi claim. | [optional] |
| **license_name** | **String** | The licence of the claim&#39;s source. | [optional] |
| **attribution** | **String** | The credit line for the claim&#39;s source. Show it wherever the claim is shown. | [optional] |
| **restriction_tier** | **String** | none for an extracted claim, licensable for an SBTi claim. The view withholds a licensable row from a caller who does not hold the licence. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgDisclosuresRow.new(
  org_id: null,
  org_name: null,
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  claim_category: null,
  claim_type: null,
  disclosure_source: null,
  disclosure_state: null,
  claim_value: null,
  claim_detail: null,
  disclosure_date: null,
  reporting_year: null,
  evidence_quote: null,
  confidence: null,
  source_record_id: null,
  license_name: null,
  attribution: null,
  restriction_tier: null
)
```

