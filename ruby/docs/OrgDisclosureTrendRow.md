# PharmaObservatory::OrgDisclosureTrendRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **org_id** | **String** | Legal-entity level: the organization record of the SEC registrant that filed. | [optional] |
| **org_name** | **String** | Display name of the filing organization record. | [optional] |
| **company_id** | **String** | Company level: the company the filing organization belongs to (null when unresolved). | [optional] |
| **company_name** | **String** | Display name of the filing organization&#39;s company. | [optional] |
| **group_id** | **String** | Group level: the filing organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the filing organization&#39;s group. | [optional] |
| **reporting_year** | **Integer** | The fiscal year the filing reports. | [optional] |
| **reporting_company** | **String** | The company name as the filing states it. | [optional] |
| **cik** | **String** | The registrant&#39;s SEC Central Index Key (CIK). | [optional] |
| **ticker** | **String** | The stock ticker the filing states. | [optional] |
| **board_pct_women** | **Float** | Women&#39;s share of the board, as the filing states it. | [optional] |
| **workforce_pct_women** | **Float** | Women&#39;s share of the workforce. | [optional] |
| **leadership_pct_women** | **Float** | Women&#39;s share of leadership roles. | [optional] |
| **workforce_pct_underrepresented** | **Float** | Share of the workforce from underrepresented groups. | [optional] |
| **gender_pay_gap_pct** | **Float** | The stated gender pay gap. | [optional] |
| **has_dei_targets** | **Boolean** | Whether the filing states diversity, equity and inclusion targets. | [optional] |
| **has_net_zero_target** | **Boolean** | Whether the filing states a net-zero target. | [optional] |
| **net_zero_target_year** | **Integer** | The year of the net-zero target. | [optional] |
| **reduction_pct** | **Float** | The emissions reduction the filing commits to. | [optional] |
| **sbti_status_claimed** | **String** | The Science Based Targets initiative (SBTi) status the filing claims for itself (a claim in the filing, not the SBTi list). | [optional] |
| **target_temperature_claimed** | **String** | The warming limit the filing says its targets align with. | [optional] |
| **climate_document_id** | **String** | The filing document the climate measures were extracted from. | [optional] |
| **diversity_document_id** | **String** | The filing document the diversity measures were extracted from. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgDisclosureTrendRow.new(
  org_id: null,
  org_name: null,
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  reporting_year: null,
  reporting_company: null,
  cik: null,
  ticker: null,
  board_pct_women: null,
  workforce_pct_women: null,
  leadership_pct_women: null,
  workforce_pct_underrepresented: null,
  gender_pay_gap_pct: null,
  has_dei_targets: null,
  has_net_zero_target: null,
  net_zero_target_year: null,
  reduction_pct: null,
  sbti_status_claimed: null,
  target_temperature_claimed: null,
  climate_document_id: null,
  diversity_document_id: null
)
```

