# UnmatchedProductRevenueRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevenueLineId** | Pointer to **NullableString** | The revenue line id: a hash of the document, product name, fiscal year, basis and geography. | [optional] 
**FilerCik** | Pointer to **NullableString** | The filer&#39;s SEC Central Index Key, ten digits. | [optional] 
**FilerName** | Pointer to **NullableString** | The filer&#39;s name as filed (the reporting company of the document). | [optional] 
**FilerOrgId** | Pointer to **NullableString** | Legal-entity level: the filer&#39;s organization record, edgar:&lt;cik&gt; when that record exists, else the single record carrying the CIK (null when none; the CIK is kept). | [optional] 
**FilerCompanyId** | Pointer to **NullableString** | Company level: the filer&#39;s company (null when the organization does not resolve to a company). | [optional] 
**FilerCompanyName** | Pointer to **NullableString** | Display name of the filer&#39;s company. | [optional] 
**FilerGroupId** | Pointer to **NullableString** | Group level: the filer&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**FilerGroupName** | Pointer to **NullableString** | Display name of the filer&#39;s group. | [optional] 
**DocumentId** | Pointer to **NullableString** | The extracted filing document id. | [optional] 
**AccessionNumber** | Pointer to **NullableString** | The SEC accession number of the filing. | [optional] 
**Form** | Pointer to **NullableString** | The SEC form the line was extracted from (10-K or 20-F). | [optional] 
**FilingDate** | Pointer to **NullableString** | The filing date of the document. | [optional] 
**FiscalYearCovered** | Pointer to **NullableInt32** | The fiscal year the document covers, as the extraction states it. | [optional] 
**ProductName** | Pointer to **NullableString** | The product name as filed, never a drug name: the brand gate did not claim it, and devices, services and apparel that share a word with a drug land here by design. | [optional] 
**CleanedKey** | Pointer to **NullableString** | The product name reduced to its brand-alias key, the form the nearest-name pointer compares. | [optional] 
**FiscalYear** | Pointer to **NullableInt32** | The fiscal year the line reports revenue for. | [optional] 
**RevenueBasis** | Pointer to **NullableString** | What the line&#39;s revenue covers, as the extraction states it: worldwide, us, ex_us, region or country. | [optional] 
**Geography** | Pointer to **NullableString** | The region or country the line covers, as filed; null for a worldwide or us line. | [optional] 
**RevenueAmount** | Pointer to **NullableFloat32** | The revenue as filed, in the currency and unit the line states (revenue_currency and reporting_unit). | [optional] 
**RevenueCurrency** | Pointer to **NullableString** | The currency the line states its revenue in. | [optional] 
**ReportingUnit** | Pointer to **NullableString** | The unit the document states its amounts in (units, thousands, millions or billions). | [optional] 
**RevenueUsdMillions** | Pointer to **NullableFloat32** | The line&#39;s revenue in US dollars millions: the amount normalized by reporting_unit, then converted at the annual average exchange rate of the filing&#39;s fiscal year. Null when the unit is unreadable or no rate exists for that currency and year. The extractor&#39;s own USD figure is not published, because it reads 0 for a non-USD line. | [optional] 
**IsLatestWorldwideLine** | Pointer to **NullableBool** | True for the single worldwide line per filer and product with a USD figure from the latest fiscal year (the line the residual mart takes as the product&#39;s latest worldwide figure; ties broken by the headline filing, then the earlier filing date). | [optional] 
**ProductLineCount** | Pointer to **NullableInt32** | How many extracted lines the product has, repeated on every line of the product. | [optional] 
**BestAliasKey** | Pointer to **NullableString** | The nearest trade-name alias by trigram similarity: a pointer for a person to decide on, not a match. Null when no alias is close enough. | [optional] 
**BestAliasBrandKey** | Pointer to **NullableString** | The brand key of the nearest alias. | [optional] 
**BestAliasSimilarity** | Pointer to **NullableFloat32** | The trigram similarity between cleaned_key and best_alias_key, 0 to 1. | [optional] 
**BrandOrGeneric** | Pointer to **NullableString** | The extraction&#39;s own label for the product (brand, generic, franchise and so on). | [optional] 
**Indication** | Pointer to **NullableString** | The indication the extraction attached to the product, as filed. | [optional] 
**EvidenceQuote** | Pointer to **NullableString** | The filing&#39;s own text the line was extracted from. | [optional] 
**ExtractionConfidence** | Pointer to **NullableFloat32** | The language model&#39;s self-reported confidence for the line, 0 to 1; not a reliability measure. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The link to the filing on SEC EDGAR. | [optional] 

## Methods

### NewUnmatchedProductRevenueRow

`func NewUnmatchedProductRevenueRow() *UnmatchedProductRevenueRow`

NewUnmatchedProductRevenueRow instantiates a new UnmatchedProductRevenueRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnmatchedProductRevenueRowWithDefaults

`func NewUnmatchedProductRevenueRowWithDefaults() *UnmatchedProductRevenueRow`

NewUnmatchedProductRevenueRowWithDefaults instantiates a new UnmatchedProductRevenueRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevenueLineId

`func (o *UnmatchedProductRevenueRow) GetRevenueLineId() string`

GetRevenueLineId returns the RevenueLineId field if non-nil, zero value otherwise.

### GetRevenueLineIdOk

`func (o *UnmatchedProductRevenueRow) GetRevenueLineIdOk() (*string, bool)`

GetRevenueLineIdOk returns a tuple with the RevenueLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueLineId

`func (o *UnmatchedProductRevenueRow) SetRevenueLineId(v string)`

SetRevenueLineId sets RevenueLineId field to given value.

### HasRevenueLineId

`func (o *UnmatchedProductRevenueRow) HasRevenueLineId() bool`

HasRevenueLineId returns a boolean if a field has been set.

### SetRevenueLineIdNil

`func (o *UnmatchedProductRevenueRow) SetRevenueLineIdNil(b bool)`

 SetRevenueLineIdNil sets the value for RevenueLineId to be an explicit nil

### UnsetRevenueLineId
`func (o *UnmatchedProductRevenueRow) UnsetRevenueLineId()`

UnsetRevenueLineId ensures that no value is present for RevenueLineId, not even an explicit nil
### GetFilerCik

`func (o *UnmatchedProductRevenueRow) GetFilerCik() string`

GetFilerCik returns the FilerCik field if non-nil, zero value otherwise.

### GetFilerCikOk

`func (o *UnmatchedProductRevenueRow) GetFilerCikOk() (*string, bool)`

GetFilerCikOk returns a tuple with the FilerCik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilerCik

`func (o *UnmatchedProductRevenueRow) SetFilerCik(v string)`

SetFilerCik sets FilerCik field to given value.

### HasFilerCik

`func (o *UnmatchedProductRevenueRow) HasFilerCik() bool`

HasFilerCik returns a boolean if a field has been set.

### SetFilerCikNil

`func (o *UnmatchedProductRevenueRow) SetFilerCikNil(b bool)`

 SetFilerCikNil sets the value for FilerCik to be an explicit nil

### UnsetFilerCik
`func (o *UnmatchedProductRevenueRow) UnsetFilerCik()`

UnsetFilerCik ensures that no value is present for FilerCik, not even an explicit nil
### GetFilerName

`func (o *UnmatchedProductRevenueRow) GetFilerName() string`

GetFilerName returns the FilerName field if non-nil, zero value otherwise.

### GetFilerNameOk

`func (o *UnmatchedProductRevenueRow) GetFilerNameOk() (*string, bool)`

GetFilerNameOk returns a tuple with the FilerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilerName

`func (o *UnmatchedProductRevenueRow) SetFilerName(v string)`

SetFilerName sets FilerName field to given value.

### HasFilerName

`func (o *UnmatchedProductRevenueRow) HasFilerName() bool`

HasFilerName returns a boolean if a field has been set.

### SetFilerNameNil

`func (o *UnmatchedProductRevenueRow) SetFilerNameNil(b bool)`

 SetFilerNameNil sets the value for FilerName to be an explicit nil

### UnsetFilerName
`func (o *UnmatchedProductRevenueRow) UnsetFilerName()`

UnsetFilerName ensures that no value is present for FilerName, not even an explicit nil
### GetFilerOrgId

`func (o *UnmatchedProductRevenueRow) GetFilerOrgId() string`

GetFilerOrgId returns the FilerOrgId field if non-nil, zero value otherwise.

### GetFilerOrgIdOk

`func (o *UnmatchedProductRevenueRow) GetFilerOrgIdOk() (*string, bool)`

GetFilerOrgIdOk returns a tuple with the FilerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilerOrgId

`func (o *UnmatchedProductRevenueRow) SetFilerOrgId(v string)`

SetFilerOrgId sets FilerOrgId field to given value.

### HasFilerOrgId

`func (o *UnmatchedProductRevenueRow) HasFilerOrgId() bool`

HasFilerOrgId returns a boolean if a field has been set.

### SetFilerOrgIdNil

`func (o *UnmatchedProductRevenueRow) SetFilerOrgIdNil(b bool)`

 SetFilerOrgIdNil sets the value for FilerOrgId to be an explicit nil

### UnsetFilerOrgId
`func (o *UnmatchedProductRevenueRow) UnsetFilerOrgId()`

UnsetFilerOrgId ensures that no value is present for FilerOrgId, not even an explicit nil
### GetFilerCompanyId

`func (o *UnmatchedProductRevenueRow) GetFilerCompanyId() string`

GetFilerCompanyId returns the FilerCompanyId field if non-nil, zero value otherwise.

### GetFilerCompanyIdOk

`func (o *UnmatchedProductRevenueRow) GetFilerCompanyIdOk() (*string, bool)`

GetFilerCompanyIdOk returns a tuple with the FilerCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilerCompanyId

`func (o *UnmatchedProductRevenueRow) SetFilerCompanyId(v string)`

SetFilerCompanyId sets FilerCompanyId field to given value.

### HasFilerCompanyId

`func (o *UnmatchedProductRevenueRow) HasFilerCompanyId() bool`

HasFilerCompanyId returns a boolean if a field has been set.

### SetFilerCompanyIdNil

`func (o *UnmatchedProductRevenueRow) SetFilerCompanyIdNil(b bool)`

 SetFilerCompanyIdNil sets the value for FilerCompanyId to be an explicit nil

### UnsetFilerCompanyId
`func (o *UnmatchedProductRevenueRow) UnsetFilerCompanyId()`

UnsetFilerCompanyId ensures that no value is present for FilerCompanyId, not even an explicit nil
### GetFilerCompanyName

`func (o *UnmatchedProductRevenueRow) GetFilerCompanyName() string`

GetFilerCompanyName returns the FilerCompanyName field if non-nil, zero value otherwise.

### GetFilerCompanyNameOk

`func (o *UnmatchedProductRevenueRow) GetFilerCompanyNameOk() (*string, bool)`

GetFilerCompanyNameOk returns a tuple with the FilerCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilerCompanyName

`func (o *UnmatchedProductRevenueRow) SetFilerCompanyName(v string)`

SetFilerCompanyName sets FilerCompanyName field to given value.

### HasFilerCompanyName

`func (o *UnmatchedProductRevenueRow) HasFilerCompanyName() bool`

HasFilerCompanyName returns a boolean if a field has been set.

### SetFilerCompanyNameNil

`func (o *UnmatchedProductRevenueRow) SetFilerCompanyNameNil(b bool)`

 SetFilerCompanyNameNil sets the value for FilerCompanyName to be an explicit nil

### UnsetFilerCompanyName
`func (o *UnmatchedProductRevenueRow) UnsetFilerCompanyName()`

UnsetFilerCompanyName ensures that no value is present for FilerCompanyName, not even an explicit nil
### GetFilerGroupId

`func (o *UnmatchedProductRevenueRow) GetFilerGroupId() string`

GetFilerGroupId returns the FilerGroupId field if non-nil, zero value otherwise.

### GetFilerGroupIdOk

`func (o *UnmatchedProductRevenueRow) GetFilerGroupIdOk() (*string, bool)`

GetFilerGroupIdOk returns a tuple with the FilerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilerGroupId

`func (o *UnmatchedProductRevenueRow) SetFilerGroupId(v string)`

SetFilerGroupId sets FilerGroupId field to given value.

### HasFilerGroupId

`func (o *UnmatchedProductRevenueRow) HasFilerGroupId() bool`

HasFilerGroupId returns a boolean if a field has been set.

### SetFilerGroupIdNil

`func (o *UnmatchedProductRevenueRow) SetFilerGroupIdNil(b bool)`

 SetFilerGroupIdNil sets the value for FilerGroupId to be an explicit nil

### UnsetFilerGroupId
`func (o *UnmatchedProductRevenueRow) UnsetFilerGroupId()`

UnsetFilerGroupId ensures that no value is present for FilerGroupId, not even an explicit nil
### GetFilerGroupName

`func (o *UnmatchedProductRevenueRow) GetFilerGroupName() string`

GetFilerGroupName returns the FilerGroupName field if non-nil, zero value otherwise.

### GetFilerGroupNameOk

`func (o *UnmatchedProductRevenueRow) GetFilerGroupNameOk() (*string, bool)`

GetFilerGroupNameOk returns a tuple with the FilerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilerGroupName

`func (o *UnmatchedProductRevenueRow) SetFilerGroupName(v string)`

SetFilerGroupName sets FilerGroupName field to given value.

### HasFilerGroupName

`func (o *UnmatchedProductRevenueRow) HasFilerGroupName() bool`

HasFilerGroupName returns a boolean if a field has been set.

### SetFilerGroupNameNil

`func (o *UnmatchedProductRevenueRow) SetFilerGroupNameNil(b bool)`

 SetFilerGroupNameNil sets the value for FilerGroupName to be an explicit nil

### UnsetFilerGroupName
`func (o *UnmatchedProductRevenueRow) UnsetFilerGroupName()`

UnsetFilerGroupName ensures that no value is present for FilerGroupName, not even an explicit nil
### GetDocumentId

`func (o *UnmatchedProductRevenueRow) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *UnmatchedProductRevenueRow) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *UnmatchedProductRevenueRow) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *UnmatchedProductRevenueRow) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### SetDocumentIdNil

`func (o *UnmatchedProductRevenueRow) SetDocumentIdNil(b bool)`

 SetDocumentIdNil sets the value for DocumentId to be an explicit nil

### UnsetDocumentId
`func (o *UnmatchedProductRevenueRow) UnsetDocumentId()`

UnsetDocumentId ensures that no value is present for DocumentId, not even an explicit nil
### GetAccessionNumber

`func (o *UnmatchedProductRevenueRow) GetAccessionNumber() string`

GetAccessionNumber returns the AccessionNumber field if non-nil, zero value otherwise.

### GetAccessionNumberOk

`func (o *UnmatchedProductRevenueRow) GetAccessionNumberOk() (*string, bool)`

GetAccessionNumberOk returns a tuple with the AccessionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessionNumber

`func (o *UnmatchedProductRevenueRow) SetAccessionNumber(v string)`

SetAccessionNumber sets AccessionNumber field to given value.

### HasAccessionNumber

`func (o *UnmatchedProductRevenueRow) HasAccessionNumber() bool`

HasAccessionNumber returns a boolean if a field has been set.

### SetAccessionNumberNil

`func (o *UnmatchedProductRevenueRow) SetAccessionNumberNil(b bool)`

 SetAccessionNumberNil sets the value for AccessionNumber to be an explicit nil

### UnsetAccessionNumber
`func (o *UnmatchedProductRevenueRow) UnsetAccessionNumber()`

UnsetAccessionNumber ensures that no value is present for AccessionNumber, not even an explicit nil
### GetForm

`func (o *UnmatchedProductRevenueRow) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *UnmatchedProductRevenueRow) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *UnmatchedProductRevenueRow) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *UnmatchedProductRevenueRow) HasForm() bool`

HasForm returns a boolean if a field has been set.

### SetFormNil

`func (o *UnmatchedProductRevenueRow) SetFormNil(b bool)`

 SetFormNil sets the value for Form to be an explicit nil

### UnsetForm
`func (o *UnmatchedProductRevenueRow) UnsetForm()`

UnsetForm ensures that no value is present for Form, not even an explicit nil
### GetFilingDate

`func (o *UnmatchedProductRevenueRow) GetFilingDate() string`

GetFilingDate returns the FilingDate field if non-nil, zero value otherwise.

### GetFilingDateOk

`func (o *UnmatchedProductRevenueRow) GetFilingDateOk() (*string, bool)`

GetFilingDateOk returns a tuple with the FilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingDate

`func (o *UnmatchedProductRevenueRow) SetFilingDate(v string)`

SetFilingDate sets FilingDate field to given value.

### HasFilingDate

`func (o *UnmatchedProductRevenueRow) HasFilingDate() bool`

HasFilingDate returns a boolean if a field has been set.

### SetFilingDateNil

`func (o *UnmatchedProductRevenueRow) SetFilingDateNil(b bool)`

 SetFilingDateNil sets the value for FilingDate to be an explicit nil

### UnsetFilingDate
`func (o *UnmatchedProductRevenueRow) UnsetFilingDate()`

UnsetFilingDate ensures that no value is present for FilingDate, not even an explicit nil
### GetFiscalYearCovered

`func (o *UnmatchedProductRevenueRow) GetFiscalYearCovered() int32`

GetFiscalYearCovered returns the FiscalYearCovered field if non-nil, zero value otherwise.

### GetFiscalYearCoveredOk

`func (o *UnmatchedProductRevenueRow) GetFiscalYearCoveredOk() (*int32, bool)`

GetFiscalYearCoveredOk returns a tuple with the FiscalYearCovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearCovered

`func (o *UnmatchedProductRevenueRow) SetFiscalYearCovered(v int32)`

SetFiscalYearCovered sets FiscalYearCovered field to given value.

### HasFiscalYearCovered

`func (o *UnmatchedProductRevenueRow) HasFiscalYearCovered() bool`

HasFiscalYearCovered returns a boolean if a field has been set.

### SetFiscalYearCoveredNil

`func (o *UnmatchedProductRevenueRow) SetFiscalYearCoveredNil(b bool)`

 SetFiscalYearCoveredNil sets the value for FiscalYearCovered to be an explicit nil

### UnsetFiscalYearCovered
`func (o *UnmatchedProductRevenueRow) UnsetFiscalYearCovered()`

UnsetFiscalYearCovered ensures that no value is present for FiscalYearCovered, not even an explicit nil
### GetProductName

`func (o *UnmatchedProductRevenueRow) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UnmatchedProductRevenueRow) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UnmatchedProductRevenueRow) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *UnmatchedProductRevenueRow) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *UnmatchedProductRevenueRow) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *UnmatchedProductRevenueRow) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetCleanedKey

`func (o *UnmatchedProductRevenueRow) GetCleanedKey() string`

GetCleanedKey returns the CleanedKey field if non-nil, zero value otherwise.

### GetCleanedKeyOk

`func (o *UnmatchedProductRevenueRow) GetCleanedKeyOk() (*string, bool)`

GetCleanedKeyOk returns a tuple with the CleanedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanedKey

`func (o *UnmatchedProductRevenueRow) SetCleanedKey(v string)`

SetCleanedKey sets CleanedKey field to given value.

### HasCleanedKey

`func (o *UnmatchedProductRevenueRow) HasCleanedKey() bool`

HasCleanedKey returns a boolean if a field has been set.

### SetCleanedKeyNil

`func (o *UnmatchedProductRevenueRow) SetCleanedKeyNil(b bool)`

 SetCleanedKeyNil sets the value for CleanedKey to be an explicit nil

### UnsetCleanedKey
`func (o *UnmatchedProductRevenueRow) UnsetCleanedKey()`

UnsetCleanedKey ensures that no value is present for CleanedKey, not even an explicit nil
### GetFiscalYear

`func (o *UnmatchedProductRevenueRow) GetFiscalYear() int32`

GetFiscalYear returns the FiscalYear field if non-nil, zero value otherwise.

### GetFiscalYearOk

`func (o *UnmatchedProductRevenueRow) GetFiscalYearOk() (*int32, bool)`

GetFiscalYearOk returns a tuple with the FiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYear

`func (o *UnmatchedProductRevenueRow) SetFiscalYear(v int32)`

SetFiscalYear sets FiscalYear field to given value.

### HasFiscalYear

`func (o *UnmatchedProductRevenueRow) HasFiscalYear() bool`

HasFiscalYear returns a boolean if a field has been set.

### SetFiscalYearNil

`func (o *UnmatchedProductRevenueRow) SetFiscalYearNil(b bool)`

 SetFiscalYearNil sets the value for FiscalYear to be an explicit nil

### UnsetFiscalYear
`func (o *UnmatchedProductRevenueRow) UnsetFiscalYear()`

UnsetFiscalYear ensures that no value is present for FiscalYear, not even an explicit nil
### GetRevenueBasis

`func (o *UnmatchedProductRevenueRow) GetRevenueBasis() string`

GetRevenueBasis returns the RevenueBasis field if non-nil, zero value otherwise.

### GetRevenueBasisOk

`func (o *UnmatchedProductRevenueRow) GetRevenueBasisOk() (*string, bool)`

GetRevenueBasisOk returns a tuple with the RevenueBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueBasis

`func (o *UnmatchedProductRevenueRow) SetRevenueBasis(v string)`

SetRevenueBasis sets RevenueBasis field to given value.

### HasRevenueBasis

`func (o *UnmatchedProductRevenueRow) HasRevenueBasis() bool`

HasRevenueBasis returns a boolean if a field has been set.

### SetRevenueBasisNil

`func (o *UnmatchedProductRevenueRow) SetRevenueBasisNil(b bool)`

 SetRevenueBasisNil sets the value for RevenueBasis to be an explicit nil

### UnsetRevenueBasis
`func (o *UnmatchedProductRevenueRow) UnsetRevenueBasis()`

UnsetRevenueBasis ensures that no value is present for RevenueBasis, not even an explicit nil
### GetGeography

`func (o *UnmatchedProductRevenueRow) GetGeography() string`

GetGeography returns the Geography field if non-nil, zero value otherwise.

### GetGeographyOk

`func (o *UnmatchedProductRevenueRow) GetGeographyOk() (*string, bool)`

GetGeographyOk returns a tuple with the Geography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeography

`func (o *UnmatchedProductRevenueRow) SetGeography(v string)`

SetGeography sets Geography field to given value.

### HasGeography

`func (o *UnmatchedProductRevenueRow) HasGeography() bool`

HasGeography returns a boolean if a field has been set.

### SetGeographyNil

`func (o *UnmatchedProductRevenueRow) SetGeographyNil(b bool)`

 SetGeographyNil sets the value for Geography to be an explicit nil

### UnsetGeography
`func (o *UnmatchedProductRevenueRow) UnsetGeography()`

UnsetGeography ensures that no value is present for Geography, not even an explicit nil
### GetRevenueAmount

`func (o *UnmatchedProductRevenueRow) GetRevenueAmount() float32`

GetRevenueAmount returns the RevenueAmount field if non-nil, zero value otherwise.

### GetRevenueAmountOk

`func (o *UnmatchedProductRevenueRow) GetRevenueAmountOk() (*float32, bool)`

GetRevenueAmountOk returns a tuple with the RevenueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueAmount

`func (o *UnmatchedProductRevenueRow) SetRevenueAmount(v float32)`

SetRevenueAmount sets RevenueAmount field to given value.

### HasRevenueAmount

`func (o *UnmatchedProductRevenueRow) HasRevenueAmount() bool`

HasRevenueAmount returns a boolean if a field has been set.

### SetRevenueAmountNil

`func (o *UnmatchedProductRevenueRow) SetRevenueAmountNil(b bool)`

 SetRevenueAmountNil sets the value for RevenueAmount to be an explicit nil

### UnsetRevenueAmount
`func (o *UnmatchedProductRevenueRow) UnsetRevenueAmount()`

UnsetRevenueAmount ensures that no value is present for RevenueAmount, not even an explicit nil
### GetRevenueCurrency

`func (o *UnmatchedProductRevenueRow) GetRevenueCurrency() string`

GetRevenueCurrency returns the RevenueCurrency field if non-nil, zero value otherwise.

### GetRevenueCurrencyOk

`func (o *UnmatchedProductRevenueRow) GetRevenueCurrencyOk() (*string, bool)`

GetRevenueCurrencyOk returns a tuple with the RevenueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueCurrency

`func (o *UnmatchedProductRevenueRow) SetRevenueCurrency(v string)`

SetRevenueCurrency sets RevenueCurrency field to given value.

### HasRevenueCurrency

`func (o *UnmatchedProductRevenueRow) HasRevenueCurrency() bool`

HasRevenueCurrency returns a boolean if a field has been set.

### SetRevenueCurrencyNil

`func (o *UnmatchedProductRevenueRow) SetRevenueCurrencyNil(b bool)`

 SetRevenueCurrencyNil sets the value for RevenueCurrency to be an explicit nil

### UnsetRevenueCurrency
`func (o *UnmatchedProductRevenueRow) UnsetRevenueCurrency()`

UnsetRevenueCurrency ensures that no value is present for RevenueCurrency, not even an explicit nil
### GetReportingUnit

`func (o *UnmatchedProductRevenueRow) GetReportingUnit() string`

GetReportingUnit returns the ReportingUnit field if non-nil, zero value otherwise.

### GetReportingUnitOk

`func (o *UnmatchedProductRevenueRow) GetReportingUnitOk() (*string, bool)`

GetReportingUnitOk returns a tuple with the ReportingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingUnit

`func (o *UnmatchedProductRevenueRow) SetReportingUnit(v string)`

SetReportingUnit sets ReportingUnit field to given value.

### HasReportingUnit

`func (o *UnmatchedProductRevenueRow) HasReportingUnit() bool`

HasReportingUnit returns a boolean if a field has been set.

### SetReportingUnitNil

`func (o *UnmatchedProductRevenueRow) SetReportingUnitNil(b bool)`

 SetReportingUnitNil sets the value for ReportingUnit to be an explicit nil

### UnsetReportingUnit
`func (o *UnmatchedProductRevenueRow) UnsetReportingUnit()`

UnsetReportingUnit ensures that no value is present for ReportingUnit, not even an explicit nil
### GetRevenueUsdMillions

`func (o *UnmatchedProductRevenueRow) GetRevenueUsdMillions() float32`

GetRevenueUsdMillions returns the RevenueUsdMillions field if non-nil, zero value otherwise.

### GetRevenueUsdMillionsOk

`func (o *UnmatchedProductRevenueRow) GetRevenueUsdMillionsOk() (*float32, bool)`

GetRevenueUsdMillionsOk returns a tuple with the RevenueUsdMillions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueUsdMillions

`func (o *UnmatchedProductRevenueRow) SetRevenueUsdMillions(v float32)`

SetRevenueUsdMillions sets RevenueUsdMillions field to given value.

### HasRevenueUsdMillions

`func (o *UnmatchedProductRevenueRow) HasRevenueUsdMillions() bool`

HasRevenueUsdMillions returns a boolean if a field has been set.

### SetRevenueUsdMillionsNil

`func (o *UnmatchedProductRevenueRow) SetRevenueUsdMillionsNil(b bool)`

 SetRevenueUsdMillionsNil sets the value for RevenueUsdMillions to be an explicit nil

### UnsetRevenueUsdMillions
`func (o *UnmatchedProductRevenueRow) UnsetRevenueUsdMillions()`

UnsetRevenueUsdMillions ensures that no value is present for RevenueUsdMillions, not even an explicit nil
### GetIsLatestWorldwideLine

`func (o *UnmatchedProductRevenueRow) GetIsLatestWorldwideLine() bool`

GetIsLatestWorldwideLine returns the IsLatestWorldwideLine field if non-nil, zero value otherwise.

### GetIsLatestWorldwideLineOk

`func (o *UnmatchedProductRevenueRow) GetIsLatestWorldwideLineOk() (*bool, bool)`

GetIsLatestWorldwideLineOk returns a tuple with the IsLatestWorldwideLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLatestWorldwideLine

`func (o *UnmatchedProductRevenueRow) SetIsLatestWorldwideLine(v bool)`

SetIsLatestWorldwideLine sets IsLatestWorldwideLine field to given value.

### HasIsLatestWorldwideLine

`func (o *UnmatchedProductRevenueRow) HasIsLatestWorldwideLine() bool`

HasIsLatestWorldwideLine returns a boolean if a field has been set.

### SetIsLatestWorldwideLineNil

`func (o *UnmatchedProductRevenueRow) SetIsLatestWorldwideLineNil(b bool)`

 SetIsLatestWorldwideLineNil sets the value for IsLatestWorldwideLine to be an explicit nil

### UnsetIsLatestWorldwideLine
`func (o *UnmatchedProductRevenueRow) UnsetIsLatestWorldwideLine()`

UnsetIsLatestWorldwideLine ensures that no value is present for IsLatestWorldwideLine, not even an explicit nil
### GetProductLineCount

`func (o *UnmatchedProductRevenueRow) GetProductLineCount() int32`

GetProductLineCount returns the ProductLineCount field if non-nil, zero value otherwise.

### GetProductLineCountOk

`func (o *UnmatchedProductRevenueRow) GetProductLineCountOk() (*int32, bool)`

GetProductLineCountOk returns a tuple with the ProductLineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductLineCount

`func (o *UnmatchedProductRevenueRow) SetProductLineCount(v int32)`

SetProductLineCount sets ProductLineCount field to given value.

### HasProductLineCount

`func (o *UnmatchedProductRevenueRow) HasProductLineCount() bool`

HasProductLineCount returns a boolean if a field has been set.

### SetProductLineCountNil

`func (o *UnmatchedProductRevenueRow) SetProductLineCountNil(b bool)`

 SetProductLineCountNil sets the value for ProductLineCount to be an explicit nil

### UnsetProductLineCount
`func (o *UnmatchedProductRevenueRow) UnsetProductLineCount()`

UnsetProductLineCount ensures that no value is present for ProductLineCount, not even an explicit nil
### GetBestAliasKey

`func (o *UnmatchedProductRevenueRow) GetBestAliasKey() string`

GetBestAliasKey returns the BestAliasKey field if non-nil, zero value otherwise.

### GetBestAliasKeyOk

`func (o *UnmatchedProductRevenueRow) GetBestAliasKeyOk() (*string, bool)`

GetBestAliasKeyOk returns a tuple with the BestAliasKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestAliasKey

`func (o *UnmatchedProductRevenueRow) SetBestAliasKey(v string)`

SetBestAliasKey sets BestAliasKey field to given value.

### HasBestAliasKey

`func (o *UnmatchedProductRevenueRow) HasBestAliasKey() bool`

HasBestAliasKey returns a boolean if a field has been set.

### SetBestAliasKeyNil

`func (o *UnmatchedProductRevenueRow) SetBestAliasKeyNil(b bool)`

 SetBestAliasKeyNil sets the value for BestAliasKey to be an explicit nil

### UnsetBestAliasKey
`func (o *UnmatchedProductRevenueRow) UnsetBestAliasKey()`

UnsetBestAliasKey ensures that no value is present for BestAliasKey, not even an explicit nil
### GetBestAliasBrandKey

`func (o *UnmatchedProductRevenueRow) GetBestAliasBrandKey() string`

GetBestAliasBrandKey returns the BestAliasBrandKey field if non-nil, zero value otherwise.

### GetBestAliasBrandKeyOk

`func (o *UnmatchedProductRevenueRow) GetBestAliasBrandKeyOk() (*string, bool)`

GetBestAliasBrandKeyOk returns a tuple with the BestAliasBrandKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestAliasBrandKey

`func (o *UnmatchedProductRevenueRow) SetBestAliasBrandKey(v string)`

SetBestAliasBrandKey sets BestAliasBrandKey field to given value.

### HasBestAliasBrandKey

`func (o *UnmatchedProductRevenueRow) HasBestAliasBrandKey() bool`

HasBestAliasBrandKey returns a boolean if a field has been set.

### SetBestAliasBrandKeyNil

`func (o *UnmatchedProductRevenueRow) SetBestAliasBrandKeyNil(b bool)`

 SetBestAliasBrandKeyNil sets the value for BestAliasBrandKey to be an explicit nil

### UnsetBestAliasBrandKey
`func (o *UnmatchedProductRevenueRow) UnsetBestAliasBrandKey()`

UnsetBestAliasBrandKey ensures that no value is present for BestAliasBrandKey, not even an explicit nil
### GetBestAliasSimilarity

`func (o *UnmatchedProductRevenueRow) GetBestAliasSimilarity() float32`

GetBestAliasSimilarity returns the BestAliasSimilarity field if non-nil, zero value otherwise.

### GetBestAliasSimilarityOk

`func (o *UnmatchedProductRevenueRow) GetBestAliasSimilarityOk() (*float32, bool)`

GetBestAliasSimilarityOk returns a tuple with the BestAliasSimilarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestAliasSimilarity

`func (o *UnmatchedProductRevenueRow) SetBestAliasSimilarity(v float32)`

SetBestAliasSimilarity sets BestAliasSimilarity field to given value.

### HasBestAliasSimilarity

`func (o *UnmatchedProductRevenueRow) HasBestAliasSimilarity() bool`

HasBestAliasSimilarity returns a boolean if a field has been set.

### SetBestAliasSimilarityNil

`func (o *UnmatchedProductRevenueRow) SetBestAliasSimilarityNil(b bool)`

 SetBestAliasSimilarityNil sets the value for BestAliasSimilarity to be an explicit nil

### UnsetBestAliasSimilarity
`func (o *UnmatchedProductRevenueRow) UnsetBestAliasSimilarity()`

UnsetBestAliasSimilarity ensures that no value is present for BestAliasSimilarity, not even an explicit nil
### GetBrandOrGeneric

`func (o *UnmatchedProductRevenueRow) GetBrandOrGeneric() string`

GetBrandOrGeneric returns the BrandOrGeneric field if non-nil, zero value otherwise.

### GetBrandOrGenericOk

`func (o *UnmatchedProductRevenueRow) GetBrandOrGenericOk() (*string, bool)`

GetBrandOrGenericOk returns a tuple with the BrandOrGeneric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandOrGeneric

`func (o *UnmatchedProductRevenueRow) SetBrandOrGeneric(v string)`

SetBrandOrGeneric sets BrandOrGeneric field to given value.

### HasBrandOrGeneric

`func (o *UnmatchedProductRevenueRow) HasBrandOrGeneric() bool`

HasBrandOrGeneric returns a boolean if a field has been set.

### SetBrandOrGenericNil

`func (o *UnmatchedProductRevenueRow) SetBrandOrGenericNil(b bool)`

 SetBrandOrGenericNil sets the value for BrandOrGeneric to be an explicit nil

### UnsetBrandOrGeneric
`func (o *UnmatchedProductRevenueRow) UnsetBrandOrGeneric()`

UnsetBrandOrGeneric ensures that no value is present for BrandOrGeneric, not even an explicit nil
### GetIndication

`func (o *UnmatchedProductRevenueRow) GetIndication() string`

GetIndication returns the Indication field if non-nil, zero value otherwise.

### GetIndicationOk

`func (o *UnmatchedProductRevenueRow) GetIndicationOk() (*string, bool)`

GetIndicationOk returns a tuple with the Indication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndication

`func (o *UnmatchedProductRevenueRow) SetIndication(v string)`

SetIndication sets Indication field to given value.

### HasIndication

`func (o *UnmatchedProductRevenueRow) HasIndication() bool`

HasIndication returns a boolean if a field has been set.

### SetIndicationNil

`func (o *UnmatchedProductRevenueRow) SetIndicationNil(b bool)`

 SetIndicationNil sets the value for Indication to be an explicit nil

### UnsetIndication
`func (o *UnmatchedProductRevenueRow) UnsetIndication()`

UnsetIndication ensures that no value is present for Indication, not even an explicit nil
### GetEvidenceQuote

`func (o *UnmatchedProductRevenueRow) GetEvidenceQuote() string`

GetEvidenceQuote returns the EvidenceQuote field if non-nil, zero value otherwise.

### GetEvidenceQuoteOk

`func (o *UnmatchedProductRevenueRow) GetEvidenceQuoteOk() (*string, bool)`

GetEvidenceQuoteOk returns a tuple with the EvidenceQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceQuote

`func (o *UnmatchedProductRevenueRow) SetEvidenceQuote(v string)`

SetEvidenceQuote sets EvidenceQuote field to given value.

### HasEvidenceQuote

`func (o *UnmatchedProductRevenueRow) HasEvidenceQuote() bool`

HasEvidenceQuote returns a boolean if a field has been set.

### SetEvidenceQuoteNil

`func (o *UnmatchedProductRevenueRow) SetEvidenceQuoteNil(b bool)`

 SetEvidenceQuoteNil sets the value for EvidenceQuote to be an explicit nil

### UnsetEvidenceQuote
`func (o *UnmatchedProductRevenueRow) UnsetEvidenceQuote()`

UnsetEvidenceQuote ensures that no value is present for EvidenceQuote, not even an explicit nil
### GetExtractionConfidence

`func (o *UnmatchedProductRevenueRow) GetExtractionConfidence() float32`

GetExtractionConfidence returns the ExtractionConfidence field if non-nil, zero value otherwise.

### GetExtractionConfidenceOk

`func (o *UnmatchedProductRevenueRow) GetExtractionConfidenceOk() (*float32, bool)`

GetExtractionConfidenceOk returns a tuple with the ExtractionConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionConfidence

`func (o *UnmatchedProductRevenueRow) SetExtractionConfidence(v float32)`

SetExtractionConfidence sets ExtractionConfidence field to given value.

### HasExtractionConfidence

`func (o *UnmatchedProductRevenueRow) HasExtractionConfidence() bool`

HasExtractionConfidence returns a boolean if a field has been set.

### SetExtractionConfidenceNil

`func (o *UnmatchedProductRevenueRow) SetExtractionConfidenceNil(b bool)`

 SetExtractionConfidenceNil sets the value for ExtractionConfidence to be an explicit nil

### UnsetExtractionConfidence
`func (o *UnmatchedProductRevenueRow) UnsetExtractionConfidence()`

UnsetExtractionConfidence ensures that no value is present for ExtractionConfidence, not even an explicit nil
### GetSourceUrl

`func (o *UnmatchedProductRevenueRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *UnmatchedProductRevenueRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *UnmatchedProductRevenueRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *UnmatchedProductRevenueRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *UnmatchedProductRevenueRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *UnmatchedProductRevenueRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


