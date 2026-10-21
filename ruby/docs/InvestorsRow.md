# PharmaObservatory::InvestorsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **investor_id** | **String** | The investor key, the id the web page uses: adv:&lt;CRD&gt; for a Form ADV adviser, lei:&lt;LEI&gt; where one is known, agency:&lt;slug&gt; for an NIH RePORTER funding agency, investor:&lt;hash&gt; for a name-resolved investor. | [optional] |
| **legal_name** | **String** | The investor&#39;s legal or display name. | [optional] |
| **investor_type** | **String** | venture_capital, private_equity, hedge_fund (from the plurality of the adviser&#39;s Form ADV Item 7.B fund types), government_agency or other. | [optional] |
| **source** | **String** | The source lane: sec_adv (Form ADV advisers) or reporter (NIH RePORTER funding agencies). | [optional] |
| **country** | **String** | The adviser&#39;s main-office country, or United States for an NIH funding agency. | [optional] |
| **raum_usd** | **Float** | Regulatory assets under management (RAUM) in whole US dollars, as the adviser reports it on Form ADV. Null for an agency or an adviser that reports none. | [optional] |
| **raum_band** | **String** | The RAUM band the amount falls in. | [optional] |
| **is_era** | **Boolean** | True for an exempt reporting adviser (ERA), false for a registered adviser, null for an agency. | [optional] |
| **is_life_sciences_focused** | **Boolean** | True when the investor is life-sciences focused, false when it is not, null when no signal exists. Read it with is_life_sciences_focused_basis. | [optional] |
| **is_life_sciences_focused_basis** | **String** | Where is_life_sciences_focused came from: human or agent (a recorded classification), override (a curator&#39;s call), edge_share (the share of its portfolio edges in life-sciences issuers) or classifier (the Form 13F biomedical value share). Null exactly when the flag is null. | [optional] |
| **private_fund_count** | **Integer** | Private funds the adviser reports on Form ADV Item 7.B. | [optional] |
| **total_private_fund_gav** | **Float** | The adviser&#39;s total private-fund gross asset value (GAV) in whole US dollars, a firm-level figure. | [optional] |
| **portfolio_co_count** | **Integer** | Distinct portfolio companies named against the investor in S-1 and 424B offering holder tables. 0 usually means no named-holder link was observed. | [optional] |
| **holdings_position_count** | **Integer** | Current Form 13F positions reported by the managers that resolve to this investor, summed over its manager keys. 0 when none resolves. | [optional] |
| **org_id** | **String** | Legal-entity level: the investor&#39;s organization record where the link resolves (null otherwise). An advfirm:&lt;CRD&gt; shell resolves to no legal-entity record and no company. | [optional] |
| **company_id** | **String** | Company level: the investor&#39;s company (null when unresolved). | [optional] |
| **company_name** | **String** | Display name of the investor&#39;s company. | [optional] |
| **group_id** | **String** | Group level: the investor&#39;s company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the investor&#39;s group. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestorsRow.new(
  investor_id: null,
  legal_name: null,
  investor_type: null,
  source: null,
  country: null,
  raum_usd: null,
  raum_band: null,
  is_era: null,
  is_life_sciences_focused: null,
  is_life_sciences_focused_basis: null,
  private_fund_count: null,
  total_private_fund_gav: null,
  portfolio_co_count: null,
  holdings_position_count: null,
  org_id: null,
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null
)
```

