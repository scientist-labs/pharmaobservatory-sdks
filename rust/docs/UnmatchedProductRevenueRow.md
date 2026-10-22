# UnmatchedProductRevenueRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**revenue_line_id** | Option<**String**> | The revenue line id: a hash of the document, product name, fiscal year, basis and geography. | [optional]
**filer_cik** | Option<**String**> | The filer's SEC Central Index Key, ten digits. | [optional]
**filer_name** | Option<**String**> | The filer's name as filed (the reporting company of the document). | [optional]
**filer_org_id** | Option<**String**> | Legal-entity level: the filer's organization record, edgar:<cik> when that record exists, else the single record carrying the CIK (null when none; the CIK is kept). | [optional]
**filer_company_id** | Option<**String**> | Company level: the filer's company (null when the organization does not resolve to a company). | [optional]
**filer_company_name** | Option<**String**> | Display name of the filer's company. | [optional]
**filer_group_id** | Option<**String**> | Group level: the filer's current ultimate parent company; a root company's group is itself. | [optional]
**filer_group_name** | Option<**String**> | Display name of the filer's group. | [optional]
**document_id** | Option<**String**> | The extracted filing document id. | [optional]
**accession_number** | Option<**String**> | The SEC accession number of the filing. | [optional]
**form** | Option<**String**> | The SEC form the line was extracted from (10-K or 20-F). | [optional]
**filing_date** | Option<[**String**](string.md)> | The filing date of the document. | [optional]
**fiscal_year_covered** | Option<**i32**> | The fiscal year the document covers, as the extraction states it. | [optional]
**product_name** | Option<**String**> | The product name as filed, never a drug name: the brand gate did not claim it, and devices, services and apparel that share a word with a drug land here by design. | [optional]
**cleaned_key** | Option<**String**> | The product name reduced to its brand-alias key, the form the nearest-name pointer compares. | [optional]
**fiscal_year** | Option<**i32**> | The fiscal year the line reports revenue for. | [optional]
**revenue_basis** | Option<**String**> | What the line's revenue covers, as the extraction states it: worldwide, us, ex_us, region or country. | [optional]
**geography** | Option<**String**> | The region or country the line covers, as filed; null for a worldwide or us line. | [optional]
**revenue_amount** | Option<**f64**> | The revenue as filed, in the currency and unit the line states (revenue_currency and reporting_unit). | [optional]
**revenue_currency** | Option<**String**> | The currency the line states its revenue in. | [optional]
**reporting_unit** | Option<**String**> | The unit the document states its amounts in (units, thousands, millions or billions). | [optional]
**revenue_usd_millions** | Option<**f64**> | The line's revenue in US dollars millions: the amount normalized by reporting_unit, then converted at the annual average exchange rate of the filing's fiscal year. Null when the unit is unreadable or no rate exists for that currency and year. The extractor's own USD figure is not published, because it reads 0 for a non-USD line. | [optional]
**is_latest_worldwide_line** | Option<**bool**> | True for the single worldwide line per filer and product with a USD figure from the latest fiscal year (the line the residual mart takes as the product's latest worldwide figure; ties broken by the headline filing, then the earlier filing date). | [optional]
**product_line_count** | Option<**i32**> | How many extracted lines the product has, repeated on every line of the product. | [optional]
**best_alias_key** | Option<**String**> | The nearest trade-name alias by trigram similarity: a pointer for a person to decide on, not a match. Null when no alias is close enough. | [optional]
**best_alias_brand_key** | Option<**String**> | The brand key of the nearest alias. | [optional]
**best_alias_similarity** | Option<**f64**> | The trigram similarity between cleaned_key and best_alias_key, 0 to 1. | [optional]
**brand_or_generic** | Option<**String**> | The extraction's own label for the product (brand, generic, franchise and so on). | [optional]
**indication** | Option<**String**> | The indication the extraction attached to the product, as filed. | [optional]
**evidence_quote** | Option<**String**> | The filing's own text the line was extracted from. | [optional]
**extraction_confidence** | Option<**f64**> | The language model's self-reported confidence for the line, 0 to 1; not a reliability measure. | [optional]
**source_url** | Option<**String**> | The link to the filing on SEC EDGAR. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


