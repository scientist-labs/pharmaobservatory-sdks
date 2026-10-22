# PharmaObservatory::InstitutionalQuarterTotalsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **quarter_total_id** | **String** | The issuer and report period id. | [optional] |
| **issuer_org_id** | **String** | Legal-entity level: the issuer organization as the ownership identity seam resolves it, a resolution root; a listed subsidiary&#39;s securities may be keyed at its resolved parent. | [optional] |
| **issuer_company_id** | **String** | Company level: the issuer&#39;s company (null when the organization does not resolve to a company). | [optional] |
| **issuer_company_name** | **String** | Display name of the issuer&#39;s company. | [optional] |
| **issuer_group_id** | **String** | Group level: the issuer&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **issuer_group_name** | **String** | Display name of the issuer&#39;s group. | [optional] |
| **report_period** | **Date** | The Form 13F report period (quarter end) the total is measured at. | [optional] |
| **is_reference_period** | **Boolean** | True for the one quarter the ownership percentage reads: the most recent in which at least 80 percent of the prior quarter&#39;s filing managers had filed. | [optional] |
| **manager_count** | **Integer** | Managers reporting a counted equity position in the issuer at the period. | [optional] |
| **position_count** | **Integer** | Counted equity positions reported against the issuer at the period. | [optional] |
| **lane_position_count** | **Integer** | Every position reported against the issuer at the period, counted or excluded. | [optional] |
| **shares_reported** | **Float** | Shares reported on Form 13F for the issuer at the period: counted equity positions only, summed within this issuer and period. | [optional] |
| **shares_otr** | **Float** | Of shares_reported, the shares under other (OTR) discretion. | [optional] |
| **shares_excluded_non_equity_issue** | **Float** | Shares reported against convertible notes, debt and preferred issues, excluded from shares_reported. | [optional] |
| **excluded_non_equity_issue_position_count** | **Integer** | Positions excluded as non-equity issues. | [optional] |
| **excluded_derivative_position_count** | **Integer** | Positions excluded because they are puts or calls. | [optional] |
| **excluded_principal_position_count** | **Integer** | Positions excluded because they are reported as a principal amount rather than shares. | [optional] |
| **excluded_non_equity_cusip_count** | **Integer** | Distinct CUSIPs excluded as non-equity issues. | [optional] |
| **equity_cusip_count** | **Integer** | Distinct equity CUSIPs counted at the period. | [optional] |
| **latest_filing_date** | **Date** | The latest Form 13F filing date among the period&#39;s positions. | [optional] |
| **aggregation_rule** | **String** | The rule the total was built under; always sum_within_issuer_period_only. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InstitutionalQuarterTotalsRow.new(
  quarter_total_id: null,
  issuer_org_id: null,
  issuer_company_id: null,
  issuer_company_name: null,
  issuer_group_id: null,
  issuer_group_name: null,
  report_period: null,
  is_reference_period: null,
  manager_count: null,
  position_count: null,
  lane_position_count: null,
  shares_reported: null,
  shares_otr: null,
  shares_excluded_non_equity_issue: null,
  excluded_non_equity_issue_position_count: null,
  excluded_derivative_position_count: null,
  excluded_principal_position_count: null,
  excluded_non_equity_cusip_count: null,
  equity_cusip_count: null,
  latest_filing_date: null,
  aggregation_rule: null
)
```

