# PharmaObservatory::InvestorPortfolioRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **investor_id** | **String** | The investor key. | [optional] |
| **company_key** | **String** | The company&#39;s key on this dataset: its organization id where it resolves, else the SEC CIK, else investor_site: followed by the normalized name as printed on the investor&#39;s page. | [optional] |
| **company_key_kind** | **String** | What company_key is: org, cik or investor_site. | [optional] |
| **sources** | **Array&lt;String&gt;** | Where the row comes from: sec_edgar (S-1 and 424B holder tables), investor_site (the investor&#39;s own portfolio page), or both. | [optional] |
| **issuer_cik** | **String** | The issuer&#39;s SEC Central Index Key (CIK), ten digits. Null for a row only the investor&#39;s website names. | [optional] |
| **issuer_name** | **String** | The company&#39;s name: its organization record&#39;s name, else the SEC company name, else the name as printed on the investor&#39;s page. | [optional] |
| **company_name_as_listed** | **String** | The company&#39;s name as printed on the investor&#39;s own portfolio page. Null for a row only SEC filings name. | [optional] |
| **company_website_url** | **String** | The company&#39;s website as linked on the investor&#39;s own portfolio page. Null when the page gave none. | [optional] |
| **issuer_org_id** | **String** | Legal-entity level: the company&#39;s organization record (null when unresolved). | [optional] |
| **issuer_company_id** | **String** | Company level: the company&#39;s record (null when unresolved). | [optional] |
| **issuer_company_name** | **String** | Display name of the company. | [optional] |
| **issuer_group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **issuer_group_name** | **String** | Display name of the company&#39;s group. | [optional] |
| **relationship_types** | **Array&lt;String&gt;** | How the investor is tied to the company: holder_5pct (a 5 percent holder) or selling_stockholder from filings, and portfolio_listed (status-free) from the investor&#39;s own page. While website status is gated the website lane is always portfolio_listed. | [optional] |
| **edge_count** | **Integer** | Filing mentions behind the row, across every filing and relationship type. 0 for a row only the investor&#39;s website names. | [optional] |
| **first_observed_date** | **Date** | The earliest filing date the investor was named against the issuer. Null for a row only the investor&#39;s website names. | [optional] |
| **last_observed_date** | **Date** | The latest filing date the investor was named against the issuer. Null for a row only the investor&#39;s website names. | [optional] |
| **listed_from** | **Date** | The retrieval date of the first crawl of the investor&#39;s own portfolio page that listed the company (a crawl date, not an investment date). Null for a row only SEC filings name. | [optional] |
| **listed_to** | **Date** | Always null while website-lane removal is gated: a crawl-derived absence is not served as a removal. Null for a row only SEC filings name. | [optional] |
| **removal_reason** | **String** | Always null while website-lane removal is gated. When a gated version is re-enabled: absent_from_later_version (a later crawl of the page omitted the company) or page_gone (the page went away), null exactly when listed_to is null. | [optional] |
| **is_removed** | **Boolean** | Always false while website-lane removal is gated; no row is served as removed. When re-enabled: true when the investor&#39;s own page no longer lists the company (listed_to is set). | [optional] |
| **status_on_page** | **String** | Always unknown for a website-listed company while website status is gated (the v1 status is not reliable enough to show). When re-enabled: the status the page states, current, realized or exited. Null for a row only SEC filings name. | [optional] |
| **status_date** | **String** | Always null while website status is gated. When re-enabled: the date the page itself prints for the status, as printed (for example 2024 or March 2024). | [optional] |
| **status_first_observed_at** | **Date** | Always null while website status is gated. When re-enabled: the date we first saw this status on the investor&#39;s page, an observation date, never the date the event happened. | [optional] |
| **site_evidence_verified** | **Boolean** | True when the quote behind a website listing was found on the investor&#39;s page, false when it was not, null when no verification flag was recorded and for a row only SEC filings name. | [optional] |
| **source_url** | **String** | The address of the investor&#39;s portfolio page the listing was read from. Null for a row only SEC filings name. | [optional] |
| **site_last_retrieved_at** | **Time** | When the investor&#39;s portfolio page was last retrieved with the company listed on it, in UTC. | [optional] |
| **is_life_sciences_issuer** | **Boolean** | True when the issuer&#39;s SIC (Standard Industrial Classification) code is a life-sciences code, false when it is another code, null when the issuer has no SIC evidence. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestorPortfolioRow.new(
  investor_id: null,
  company_key: null,
  company_key_kind: null,
  sources: null,
  issuer_cik: null,
  issuer_name: null,
  company_name_as_listed: null,
  company_website_url: null,
  issuer_org_id: null,
  issuer_company_id: null,
  issuer_company_name: null,
  issuer_group_id: null,
  issuer_group_name: null,
  relationship_types: null,
  edge_count: null,
  first_observed_date: null,
  last_observed_date: null,
  listed_from: null,
  listed_to: null,
  removal_reason: null,
  is_removed: null,
  status_on_page: null,
  status_date: null,
  status_first_observed_at: null,
  site_evidence_verified: null,
  source_url: null,
  site_last_retrieved_at: null,
  is_life_sciences_issuer: null
)
```

