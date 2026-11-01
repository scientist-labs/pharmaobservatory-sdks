# InvestorPortfolioRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**investor_id** | Option<**String**> | The investor key. | [optional]
**company_key** | Option<**String**> | The company's key on this dataset: its organization id where it resolves, else the SEC CIK, else investor_site: followed by the normalized name as printed on the investor's page. | [optional]
**company_key_kind** | Option<**String**> | What company_key is: org, cik or investor_site. | [optional]
**sources** | Option<**Vec<String>**> | Where the row comes from: sec_edgar (S-1 and 424B holder tables), investor_site (the investor's own portfolio page), or both. | [optional]
**issuer_cik** | Option<**String**> | The issuer's SEC Central Index Key (CIK), ten digits. Null for a row only the investor's website names. | [optional]
**issuer_name** | Option<**String**> | The company's name: its organization record's name, else the SEC company name, else the name as printed on the investor's page. | [optional]
**company_name_as_listed** | Option<**String**> | The company's name as printed on the investor's own portfolio page. Null for a row only SEC filings name. | [optional]
**company_website_url** | Option<**String**> | The company's website as linked on the investor's own portfolio page. Null when the page gave none. | [optional]
**issuer_org_id** | Option<**String**> | Legal-entity level: the company's organization record (null when unresolved). | [optional]
**issuer_company_id** | Option<**String**> | Company level: the company's record (null when unresolved). | [optional]
**issuer_company_name** | Option<**String**> | Display name of the company. | [optional]
**issuer_group_id** | Option<**String**> | Group level: the company's current ultimate parent; a root's group is itself. | [optional]
**issuer_group_name** | Option<**String**> | Display name of the company's group. | [optional]
**relationship_types** | Option<**Vec<String>**> | How the investor is tied to the company: holder_5pct (a 5 percent holder) or selling_stockholder from filings, and portfolio_listed_current, portfolio_listed_realized, portfolio_listed_exited or portfolio_listed (no status printed) from the investor's own page. | [optional]
**edge_count** | Option<**i64**> | Filing mentions behind the row, across every filing and relationship type. 0 for a row only the investor's website names. | [optional]
**first_observed_date** | Option<[**String**](string.md)> | The earliest filing date the investor was named against the issuer. Null for a row only the investor's website names. | [optional]
**last_observed_date** | Option<[**String**](string.md)> | The latest filing date the investor was named against the issuer. Null for a row only the investor's website names. | [optional]
**listed_from** | Option<[**String**](string.md)> | The retrieval date of the first crawl of the investor's own portfolio page that listed the company (a crawl date, not an investment date). Null for a row only SEC filings name. | [optional]
**listed_to** | Option<[**String**](string.md)> | The retrieval date of the first later crawl that did not list the company, or the date the page went away. Null while the company is still listed, and for a row only SEC filings name. | [optional]
**removal_reason** | Option<**String**> | Why listed_to is set: absent_from_later_version (a later crawl of the page omitted the company) or page_gone (the page went away). Null exactly when listed_to is null. | [optional]
**is_removed** | Option<**bool**> | True when the investor's own page no longer lists the company (listed_to is set). SEC filings do not reopen it. | [optional]
**status_on_page** | Option<**String**> | The status the investor's page states for the company, verbatim: current, realized, exited or unknown (the page listed it with no status). Null for a row only SEC filings name. | [optional]
**status_date** | Option<**String**> | The date the page itself prints for the status, as printed (for example 2024 or March 2024). Null unless the page printed one. | [optional]
**status_first_observed_at** | Option<[**String**](string.md)> | The date we first saw this status for the company on the investor's page: an observation date, never the date the event happened. | [optional]
**source_url** | Option<**String**> | The address of the investor's portfolio page the listing was read from. Null for a row only SEC filings name. | [optional]
**site_last_retrieved_at** | Option<**String**> | When the investor's portfolio page was last retrieved with the company listed on it, in UTC. | [optional]
**is_life_sciences_issuer** | Option<**bool**> | True when the issuer's SIC (Standard Industrial Classification) code is a life-sciences code, false when it is another code, null when the issuer has no SIC evidence. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


