# BrandRevenueRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**brand_key** | Option<**String**> | The brand key: an upper-case alphanumeric trade-name token (the brand's leading name tokens), the same key the brands pages use. Not a display name; brand_name is the display form. | [optional]
**product_name** | Option<**String**> | The product name as filed. | [optional]
**reporter_org_id** | Option<**String**> | Legal-entity level: the SEC filer that reported the figure, as filed (edgar:<CIK>). | [optional]
**reporter_company_id** | Option<**String**> | Company level: the reporting company's company (null when unresolved). | [optional]
**reporter_company_name** | Option<**String**> | Display name of the reporting company's company. | [optional]
**reporter_group_id** | Option<**String**> | Group level: the reporting company's current ultimate parent company; a root company's group is itself. | [optional]
**reporter_group_name** | Option<**String**> | Display name of the reporting company's group. | [optional]
**fiscal_year** | Option<**i32**> | The fiscal year of the figure. | [optional]
**revenue_basis** | Option<**String**> | The standardised basis: worldwide, us, ex_us, region or country. | [optional]
**geography** | Option<**String**> | The normalised geography: worldwide, us or ex_us for the spine (equal to the basis), the filer's own label for a region or country line. | [optional]
**revenue_usd_millions** | Option<**f64**> | The figure in USD millions (null where no unit or exchange rate is available). | [optional]
**revenue_amount** | Option<**f64**> | The figure as reported, in reporting_unit and revenue_currency; can be negative for adjustment or return lines. | [optional]
**revenue_currency** | Option<**String**> | The filing's currency (ISO code). | [optional]
**reporting_unit** | Option<**String**> | The filing's reporting unit (units, thousands, millions or billions); null when the filing did not state a readable one. | [optional]
**as_reported** | Option<**bool**> | True when the figure comes from the filing for its own fiscal year, false when it is a comparative year read from a later filing. | [optional]
**n_filings** | Option<**i32**> | Distinct filings that report the figure. | [optional]
**restated** | Option<**bool**> | True when several filings report the figure and differ by more than one percent. | [optional]
**confidence** | Option<**f64**> | The extraction's confidence in the line, 0 to 1. | [optional]
**source_url** | Option<**String**> | The link to the SEC filing the figure was extracted from. | [optional]
**evidence_quote** | Option<**String**> | The sentence of the filing the figure was extracted from. | [optional]
**match_method** | Option<**String**> | How the product name matched the brand: alias_prefix, product_prefix or curated_alias. A label of the mechanism, not a confidence. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


