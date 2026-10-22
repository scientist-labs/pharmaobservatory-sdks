# PharmaObservatory::OwnershipPercentagesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **subject_org_id** | **String** | Legal-entity level: the subject organization&#39;s record, a resolution root (the identity every ownership dataset keys on). One row per organization; an insider row with no organization has no row here. | [optional] |
| **subject_company_id** | **String** | Company level: the subject organization&#39;s company (null when the organization does not resolve to a company). | [optional] |
| **subject_company_name** | **String** | Display name of the subject organization&#39;s company. | [optional] |
| **subject_group_id** | **String** | Group level: the subject organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **subject_group_name** | **String** | Display name of the subject organization&#39;s group. | [optional] |
| **any_ownership_source** | **Boolean** | True when an insider row or a Form 13F quarter exists for the organization. | [optional] |
| **shares_outstanding_state** | **String** | reported when the organization has a cover-page share count on file, not_reported when it has none (then no percentage is taken and the share figures are shown instead). | [optional] |
| **shares_outstanding** | **Float** | The denominator both percentages divide into: the latest cover-page share count (dei:EntityCommonStockSharesOutstanding) from an annual form. It is measured on the filing&#39;s cover date, 0 to 3 months after the fiscal period end, and the cover date is not published, so its vintage is the fiscal year in shares_outstanding_fy. | [optional] |
| **shares_outstanding_fy** | **Integer** | The fiscal year of the filing the share count comes from. | [optional] |
| **shares_outstanding_concept** | **String** | The XBRL (eXtensible Business Reporting Language) concept the share count is read from. Always dei:EntityCommonStockSharesOutstanding. | [optional] |
| **insider_state** | **String** | How the insider figure was reached. reported: The sum of the current non-derivative balance each reporting owner last filed on a Form 3, 4 or 5, divided by the share count on file. no_insider_filing: No Form 3, 4 or 5 in our data names an insider of this company. That is a statement about coverage, not a statement that insiders hold nothing. all_positions_closed: Every insider position on file has been reported closed. This is a filed zero, not missing data. derivative_only: The insiders on file hold only derivative positions (options and similar), which are not outstanding shares. This is a filed zero, not missing data. no_denominator: No percentage: this company&#39;s SEC filings in our data carry no cover-page share count, so there is no denominator to divide by. The share count on file is shown instead. not_computable: No percentage is taken here. The reported share total is given instead, with the reason no ratio is taken. | [optional] |
| **insider_reason** | **String** | Why an insider ratio was not taken when insider_state is not_computable. denominator_stale: The share count on file is more than two years older than the positions it would be divided into, so the ratio would not describe this company. exceeds_shares_outstanding: The insider total on file is larger than the company&#39;s whole share count, which is arithmetically impossible for a single source and means the filings behind it have not consolidated correctly. Null otherwise. | [optional] |
| **insider_percent** | **Float** | The share of shares outstanding held by Form 3, 4 and 5 insiders (insider_shares divided by shares_outstanding, times 100, unrounded); null unless insider_state is reported. | [optional] |
| **insider_shares** | **Float** | The sum of the current non-derivative balance each reporting insider last filed. A ceased or options-only insider adds nothing to it. Null when no insider holds standing shares. | [optional] |
| **insider_owner_count** | **Integer** | Insiders on file for the organization (0 when none). | [optional] |
| **insider_standing_owner_count** | **Integer** | Insiders with a positive standing equity balance, the only population insider_shares describes. | [optional] |
| **insider_closed_owner_count** | **Integer** | Insiders whose standing equity balance is zero or absent: closed positions plus options-only insiders. | [optional] |
| **insider_derivative_only_owner_count** | **Integer** | Of the closed insiders, those whose only remaining positions are derivatives (options and similar), which are not outstanding shares. | [optional] |
| **insider_last_filing_date** | **Date** | The latest filing date among the organization&#39;s insider rows. | [optional] |
| **institutional_state** | **String** | How the institutional figure was reached. reported: Every position reported on Form 13F for the quarter shown, summed once per filing manager, against the same share count as the insider figure. exceeds_outstanding: Reported positions exceed the shares outstanding on file. Shown as filed. implausible: No percentage: reported positions exceed 1.5 times the shares outstanding on file, which indicates a corporate action or a denominator mismatch rather than ownership. no_13f_issuer: No Form 13F issuer resolved to this company in any quarter. Positions are reported against a SECURITY, so this is a statement about the CUSIP-to-company bridge, not about institutional interest. no_reference_period: No manager reported this company for the most recent fully-reported quarter. Positions in earlier quarters are on file and are reported per manager, but a manager whose newest filing is for an older quarter is not a current holder and is not counted here. no_reference_shares: No manager reported an equity position in this company for the most recent fully-reported quarter. Only options, debt or non-equity issues were reported, and none of those is a share count. no_denominator: No percentage: this company&#39;s SEC filings in our data carry no cover-page share count, so there is no denominator to divide by. The reported share total is shown instead. not_computable: No percentage is taken here. The reported share total is given instead, with the reason no ratio is taken. | [optional] |
| **institutional_reason** | **String** | Why an institutional ratio was not taken when institutional_state is not_computable. denominator_stale: The share count on file is more than two years older than the positions it would be divided into, so the ratio would not describe this company. Null otherwise. | [optional] |
| **institutional_percent** | **Float** | The share of shares outstanding under reported institutional discretion at the reference quarter (institutional_shares_reported divided by shares_outstanding, times 100, unrounded, as filed and never capped); null unless institutional_state is reported or exceeds_outstanding. | [optional] |
| **institutional_shares_reported** | **Float** | The shares reported on Form 13F at the reference quarter, counted equity positions only, summed once per filing manager. Null when the organization has no reference-quarter row. | [optional] |
| **institutional_manager_count** | **Integer** | Managers reporting a counted equity position at the reference quarter (0 when there is no reference-quarter row). | [optional] |
| **institutional_report_period** | **Date** | The reference quarter end the institutional figure is measured at: the most recent quarter in which at least 80 percent of the prior quarter&#39;s filing managers had filed. | [optional] |
| **institutional_shares_otr** | **Float** | Of institutional_shares_reported, the shares under other (OTR) discretion. | [optional] |
| **institutional_shares_excluded_non_equity_issue** | **Float** | Shares reported against convertible notes, debt and preferred issues at the reference quarter, excluded from institutional_shares_reported. | [optional] |
| **institutional_equity_cusip_count** | **Integer** | Distinct equity CUSIPs (Committee on Uniform Securities Identification Procedures identifiers) counted at the reference quarter (0 when there is no reference-quarter row). | [optional] |
| **institutional_latest_filing_date** | **Date** | The latest Form 13F filing date among the reference quarter&#39;s positions. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OwnershipPercentagesRow.new(
  subject_org_id: null,
  subject_company_id: null,
  subject_company_name: null,
  subject_group_id: null,
  subject_group_name: null,
  any_ownership_source: null,
  shares_outstanding_state: null,
  shares_outstanding: null,
  shares_outstanding_fy: null,
  shares_outstanding_concept: null,
  insider_state: null,
  insider_reason: null,
  insider_percent: null,
  insider_shares: null,
  insider_owner_count: null,
  insider_standing_owner_count: null,
  insider_closed_owner_count: null,
  insider_derivative_only_owner_count: null,
  insider_last_filing_date: null,
  institutional_state: null,
  institutional_reason: null,
  institutional_percent: null,
  institutional_shares_reported: null,
  institutional_manager_count: null,
  institutional_report_period: null,
  institutional_shares_otr: null,
  institutional_shares_excluded_non_equity_issue: null,
  institutional_equity_cusip_count: null,
  institutional_latest_filing_date: null
)
```

