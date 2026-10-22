# BrandRevenueRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**BrandKey** | Pointer to **NullableString** | The brand key: an upper-case alphanumeric trade-name token (the brand&#39;s leading name tokens), the same key the brands pages use. Not a display name; brand_name is the display form. | [optional] 
**ProductName** | Pointer to **NullableString** | The product name as filed. | [optional] 
**ReporterOrgId** | Pointer to **NullableString** | Legal-entity level: the SEC filer that reported the figure, as filed (edgar:&lt;CIK&gt;). | [optional] 
**ReporterCompanyId** | Pointer to **NullableString** | Company level: the reporting company&#39;s company (null when unresolved). | [optional] 
**ReporterCompanyName** | Pointer to **NullableString** | Display name of the reporting company&#39;s company. | [optional] 
**ReporterGroupId** | Pointer to **NullableString** | Group level: the reporting company&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**ReporterGroupName** | Pointer to **NullableString** | Display name of the reporting company&#39;s group. | [optional] 
**FiscalYear** | Pointer to **NullableInt32** | The fiscal year of the figure. | [optional] 
**RevenueBasis** | Pointer to **NullableString** | The standardised basis: worldwide, us, ex_us, region or country. | [optional] 
**Geography** | Pointer to **NullableString** | The normalised geography: worldwide, us or ex_us for the spine (equal to the basis), the filer&#39;s own label for a region or country line. | [optional] 
**RevenueUsdMillions** | Pointer to **NullableFloat64** | The figure in USD millions (null where no unit or exchange rate is available). | [optional] 
**RevenueAmount** | Pointer to **NullableFloat64** | The figure as reported, in reporting_unit and revenue_currency; can be negative for adjustment or return lines. | [optional] 
**RevenueCurrency** | Pointer to **NullableString** | The filing&#39;s currency (ISO code). | [optional] 
**ReportingUnit** | Pointer to **NullableString** | The filing&#39;s reporting unit (units, thousands, millions or billions); null when the filing did not state a readable one. | [optional] 
**AsReported** | Pointer to **NullableBool** | True when the figure comes from the filing for its own fiscal year, false when it is a comparative year read from a later filing. | [optional] 
**NFilings** | Pointer to **NullableInt32** | Distinct filings that report the figure. | [optional] 
**Restated** | Pointer to **NullableBool** | True when several filings report the figure and differ by more than one percent. | [optional] 
**Confidence** | Pointer to **NullableFloat64** | The extraction&#39;s confidence in the line, 0 to 1. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The link to the SEC filing the figure was extracted from. | [optional] 
**EvidenceQuote** | Pointer to **NullableString** | The sentence of the filing the figure was extracted from. | [optional] 
**MatchMethod** | Pointer to **NullableString** | How the product name matched the brand: alias_prefix, product_prefix or curated_alias. A label of the mechanism, not a confidence. | [optional] 

## Methods

### NewBrandRevenueRow

`func NewBrandRevenueRow() *BrandRevenueRow`

NewBrandRevenueRow instantiates a new BrandRevenueRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandRevenueRowWithDefaults

`func NewBrandRevenueRowWithDefaults() *BrandRevenueRow`

NewBrandRevenueRowWithDefaults instantiates a new BrandRevenueRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *BrandRevenueRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *BrandRevenueRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *BrandRevenueRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *BrandRevenueRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *BrandRevenueRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *BrandRevenueRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetBrandKey

`func (o *BrandRevenueRow) GetBrandKey() string`

GetBrandKey returns the BrandKey field if non-nil, zero value otherwise.

### GetBrandKeyOk

`func (o *BrandRevenueRow) GetBrandKeyOk() (*string, bool)`

GetBrandKeyOk returns a tuple with the BrandKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandKey

`func (o *BrandRevenueRow) SetBrandKey(v string)`

SetBrandKey sets BrandKey field to given value.

### HasBrandKey

`func (o *BrandRevenueRow) HasBrandKey() bool`

HasBrandKey returns a boolean if a field has been set.

### SetBrandKeyNil

`func (o *BrandRevenueRow) SetBrandKeyNil(b bool)`

 SetBrandKeyNil sets the value for BrandKey to be an explicit nil

### UnsetBrandKey
`func (o *BrandRevenueRow) UnsetBrandKey()`

UnsetBrandKey ensures that no value is present for BrandKey, not even an explicit nil
### GetProductName

`func (o *BrandRevenueRow) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *BrandRevenueRow) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *BrandRevenueRow) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *BrandRevenueRow) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *BrandRevenueRow) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *BrandRevenueRow) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetReporterOrgId

`func (o *BrandRevenueRow) GetReporterOrgId() string`

GetReporterOrgId returns the ReporterOrgId field if non-nil, zero value otherwise.

### GetReporterOrgIdOk

`func (o *BrandRevenueRow) GetReporterOrgIdOk() (*string, bool)`

GetReporterOrgIdOk returns a tuple with the ReporterOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterOrgId

`func (o *BrandRevenueRow) SetReporterOrgId(v string)`

SetReporterOrgId sets ReporterOrgId field to given value.

### HasReporterOrgId

`func (o *BrandRevenueRow) HasReporterOrgId() bool`

HasReporterOrgId returns a boolean if a field has been set.

### SetReporterOrgIdNil

`func (o *BrandRevenueRow) SetReporterOrgIdNil(b bool)`

 SetReporterOrgIdNil sets the value for ReporterOrgId to be an explicit nil

### UnsetReporterOrgId
`func (o *BrandRevenueRow) UnsetReporterOrgId()`

UnsetReporterOrgId ensures that no value is present for ReporterOrgId, not even an explicit nil
### GetReporterCompanyId

`func (o *BrandRevenueRow) GetReporterCompanyId() string`

GetReporterCompanyId returns the ReporterCompanyId field if non-nil, zero value otherwise.

### GetReporterCompanyIdOk

`func (o *BrandRevenueRow) GetReporterCompanyIdOk() (*string, bool)`

GetReporterCompanyIdOk returns a tuple with the ReporterCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterCompanyId

`func (o *BrandRevenueRow) SetReporterCompanyId(v string)`

SetReporterCompanyId sets ReporterCompanyId field to given value.

### HasReporterCompanyId

`func (o *BrandRevenueRow) HasReporterCompanyId() bool`

HasReporterCompanyId returns a boolean if a field has been set.

### SetReporterCompanyIdNil

`func (o *BrandRevenueRow) SetReporterCompanyIdNil(b bool)`

 SetReporterCompanyIdNil sets the value for ReporterCompanyId to be an explicit nil

### UnsetReporterCompanyId
`func (o *BrandRevenueRow) UnsetReporterCompanyId()`

UnsetReporterCompanyId ensures that no value is present for ReporterCompanyId, not even an explicit nil
### GetReporterCompanyName

`func (o *BrandRevenueRow) GetReporterCompanyName() string`

GetReporterCompanyName returns the ReporterCompanyName field if non-nil, zero value otherwise.

### GetReporterCompanyNameOk

`func (o *BrandRevenueRow) GetReporterCompanyNameOk() (*string, bool)`

GetReporterCompanyNameOk returns a tuple with the ReporterCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterCompanyName

`func (o *BrandRevenueRow) SetReporterCompanyName(v string)`

SetReporterCompanyName sets ReporterCompanyName field to given value.

### HasReporterCompanyName

`func (o *BrandRevenueRow) HasReporterCompanyName() bool`

HasReporterCompanyName returns a boolean if a field has been set.

### SetReporterCompanyNameNil

`func (o *BrandRevenueRow) SetReporterCompanyNameNil(b bool)`

 SetReporterCompanyNameNil sets the value for ReporterCompanyName to be an explicit nil

### UnsetReporterCompanyName
`func (o *BrandRevenueRow) UnsetReporterCompanyName()`

UnsetReporterCompanyName ensures that no value is present for ReporterCompanyName, not even an explicit nil
### GetReporterGroupId

`func (o *BrandRevenueRow) GetReporterGroupId() string`

GetReporterGroupId returns the ReporterGroupId field if non-nil, zero value otherwise.

### GetReporterGroupIdOk

`func (o *BrandRevenueRow) GetReporterGroupIdOk() (*string, bool)`

GetReporterGroupIdOk returns a tuple with the ReporterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterGroupId

`func (o *BrandRevenueRow) SetReporterGroupId(v string)`

SetReporterGroupId sets ReporterGroupId field to given value.

### HasReporterGroupId

`func (o *BrandRevenueRow) HasReporterGroupId() bool`

HasReporterGroupId returns a boolean if a field has been set.

### SetReporterGroupIdNil

`func (o *BrandRevenueRow) SetReporterGroupIdNil(b bool)`

 SetReporterGroupIdNil sets the value for ReporterGroupId to be an explicit nil

### UnsetReporterGroupId
`func (o *BrandRevenueRow) UnsetReporterGroupId()`

UnsetReporterGroupId ensures that no value is present for ReporterGroupId, not even an explicit nil
### GetReporterGroupName

`func (o *BrandRevenueRow) GetReporterGroupName() string`

GetReporterGroupName returns the ReporterGroupName field if non-nil, zero value otherwise.

### GetReporterGroupNameOk

`func (o *BrandRevenueRow) GetReporterGroupNameOk() (*string, bool)`

GetReporterGroupNameOk returns a tuple with the ReporterGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterGroupName

`func (o *BrandRevenueRow) SetReporterGroupName(v string)`

SetReporterGroupName sets ReporterGroupName field to given value.

### HasReporterGroupName

`func (o *BrandRevenueRow) HasReporterGroupName() bool`

HasReporterGroupName returns a boolean if a field has been set.

### SetReporterGroupNameNil

`func (o *BrandRevenueRow) SetReporterGroupNameNil(b bool)`

 SetReporterGroupNameNil sets the value for ReporterGroupName to be an explicit nil

### UnsetReporterGroupName
`func (o *BrandRevenueRow) UnsetReporterGroupName()`

UnsetReporterGroupName ensures that no value is present for ReporterGroupName, not even an explicit nil
### GetFiscalYear

`func (o *BrandRevenueRow) GetFiscalYear() int32`

GetFiscalYear returns the FiscalYear field if non-nil, zero value otherwise.

### GetFiscalYearOk

`func (o *BrandRevenueRow) GetFiscalYearOk() (*int32, bool)`

GetFiscalYearOk returns a tuple with the FiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYear

`func (o *BrandRevenueRow) SetFiscalYear(v int32)`

SetFiscalYear sets FiscalYear field to given value.

### HasFiscalYear

`func (o *BrandRevenueRow) HasFiscalYear() bool`

HasFiscalYear returns a boolean if a field has been set.

### SetFiscalYearNil

`func (o *BrandRevenueRow) SetFiscalYearNil(b bool)`

 SetFiscalYearNil sets the value for FiscalYear to be an explicit nil

### UnsetFiscalYear
`func (o *BrandRevenueRow) UnsetFiscalYear()`

UnsetFiscalYear ensures that no value is present for FiscalYear, not even an explicit nil
### GetRevenueBasis

`func (o *BrandRevenueRow) GetRevenueBasis() string`

GetRevenueBasis returns the RevenueBasis field if non-nil, zero value otherwise.

### GetRevenueBasisOk

`func (o *BrandRevenueRow) GetRevenueBasisOk() (*string, bool)`

GetRevenueBasisOk returns a tuple with the RevenueBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueBasis

`func (o *BrandRevenueRow) SetRevenueBasis(v string)`

SetRevenueBasis sets RevenueBasis field to given value.

### HasRevenueBasis

`func (o *BrandRevenueRow) HasRevenueBasis() bool`

HasRevenueBasis returns a boolean if a field has been set.

### SetRevenueBasisNil

`func (o *BrandRevenueRow) SetRevenueBasisNil(b bool)`

 SetRevenueBasisNil sets the value for RevenueBasis to be an explicit nil

### UnsetRevenueBasis
`func (o *BrandRevenueRow) UnsetRevenueBasis()`

UnsetRevenueBasis ensures that no value is present for RevenueBasis, not even an explicit nil
### GetGeography

`func (o *BrandRevenueRow) GetGeography() string`

GetGeography returns the Geography field if non-nil, zero value otherwise.

### GetGeographyOk

`func (o *BrandRevenueRow) GetGeographyOk() (*string, bool)`

GetGeographyOk returns a tuple with the Geography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeography

`func (o *BrandRevenueRow) SetGeography(v string)`

SetGeography sets Geography field to given value.

### HasGeography

`func (o *BrandRevenueRow) HasGeography() bool`

HasGeography returns a boolean if a field has been set.

### SetGeographyNil

`func (o *BrandRevenueRow) SetGeographyNil(b bool)`

 SetGeographyNil sets the value for Geography to be an explicit nil

### UnsetGeography
`func (o *BrandRevenueRow) UnsetGeography()`

UnsetGeography ensures that no value is present for Geography, not even an explicit nil
### GetRevenueUsdMillions

`func (o *BrandRevenueRow) GetRevenueUsdMillions() float64`

GetRevenueUsdMillions returns the RevenueUsdMillions field if non-nil, zero value otherwise.

### GetRevenueUsdMillionsOk

`func (o *BrandRevenueRow) GetRevenueUsdMillionsOk() (*float64, bool)`

GetRevenueUsdMillionsOk returns a tuple with the RevenueUsdMillions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueUsdMillions

`func (o *BrandRevenueRow) SetRevenueUsdMillions(v float64)`

SetRevenueUsdMillions sets RevenueUsdMillions field to given value.

### HasRevenueUsdMillions

`func (o *BrandRevenueRow) HasRevenueUsdMillions() bool`

HasRevenueUsdMillions returns a boolean if a field has been set.

### SetRevenueUsdMillionsNil

`func (o *BrandRevenueRow) SetRevenueUsdMillionsNil(b bool)`

 SetRevenueUsdMillionsNil sets the value for RevenueUsdMillions to be an explicit nil

### UnsetRevenueUsdMillions
`func (o *BrandRevenueRow) UnsetRevenueUsdMillions()`

UnsetRevenueUsdMillions ensures that no value is present for RevenueUsdMillions, not even an explicit nil
### GetRevenueAmount

`func (o *BrandRevenueRow) GetRevenueAmount() float64`

GetRevenueAmount returns the RevenueAmount field if non-nil, zero value otherwise.

### GetRevenueAmountOk

`func (o *BrandRevenueRow) GetRevenueAmountOk() (*float64, bool)`

GetRevenueAmountOk returns a tuple with the RevenueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueAmount

`func (o *BrandRevenueRow) SetRevenueAmount(v float64)`

SetRevenueAmount sets RevenueAmount field to given value.

### HasRevenueAmount

`func (o *BrandRevenueRow) HasRevenueAmount() bool`

HasRevenueAmount returns a boolean if a field has been set.

### SetRevenueAmountNil

`func (o *BrandRevenueRow) SetRevenueAmountNil(b bool)`

 SetRevenueAmountNil sets the value for RevenueAmount to be an explicit nil

### UnsetRevenueAmount
`func (o *BrandRevenueRow) UnsetRevenueAmount()`

UnsetRevenueAmount ensures that no value is present for RevenueAmount, not even an explicit nil
### GetRevenueCurrency

`func (o *BrandRevenueRow) GetRevenueCurrency() string`

GetRevenueCurrency returns the RevenueCurrency field if non-nil, zero value otherwise.

### GetRevenueCurrencyOk

`func (o *BrandRevenueRow) GetRevenueCurrencyOk() (*string, bool)`

GetRevenueCurrencyOk returns a tuple with the RevenueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueCurrency

`func (o *BrandRevenueRow) SetRevenueCurrency(v string)`

SetRevenueCurrency sets RevenueCurrency field to given value.

### HasRevenueCurrency

`func (o *BrandRevenueRow) HasRevenueCurrency() bool`

HasRevenueCurrency returns a boolean if a field has been set.

### SetRevenueCurrencyNil

`func (o *BrandRevenueRow) SetRevenueCurrencyNil(b bool)`

 SetRevenueCurrencyNil sets the value for RevenueCurrency to be an explicit nil

### UnsetRevenueCurrency
`func (o *BrandRevenueRow) UnsetRevenueCurrency()`

UnsetRevenueCurrency ensures that no value is present for RevenueCurrency, not even an explicit nil
### GetReportingUnit

`func (o *BrandRevenueRow) GetReportingUnit() string`

GetReportingUnit returns the ReportingUnit field if non-nil, zero value otherwise.

### GetReportingUnitOk

`func (o *BrandRevenueRow) GetReportingUnitOk() (*string, bool)`

GetReportingUnitOk returns a tuple with the ReportingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingUnit

`func (o *BrandRevenueRow) SetReportingUnit(v string)`

SetReportingUnit sets ReportingUnit field to given value.

### HasReportingUnit

`func (o *BrandRevenueRow) HasReportingUnit() bool`

HasReportingUnit returns a boolean if a field has been set.

### SetReportingUnitNil

`func (o *BrandRevenueRow) SetReportingUnitNil(b bool)`

 SetReportingUnitNil sets the value for ReportingUnit to be an explicit nil

### UnsetReportingUnit
`func (o *BrandRevenueRow) UnsetReportingUnit()`

UnsetReportingUnit ensures that no value is present for ReportingUnit, not even an explicit nil
### GetAsReported

`func (o *BrandRevenueRow) GetAsReported() bool`

GetAsReported returns the AsReported field if non-nil, zero value otherwise.

### GetAsReportedOk

`func (o *BrandRevenueRow) GetAsReportedOk() (*bool, bool)`

GetAsReportedOk returns a tuple with the AsReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsReported

`func (o *BrandRevenueRow) SetAsReported(v bool)`

SetAsReported sets AsReported field to given value.

### HasAsReported

`func (o *BrandRevenueRow) HasAsReported() bool`

HasAsReported returns a boolean if a field has been set.

### SetAsReportedNil

`func (o *BrandRevenueRow) SetAsReportedNil(b bool)`

 SetAsReportedNil sets the value for AsReported to be an explicit nil

### UnsetAsReported
`func (o *BrandRevenueRow) UnsetAsReported()`

UnsetAsReported ensures that no value is present for AsReported, not even an explicit nil
### GetNFilings

`func (o *BrandRevenueRow) GetNFilings() int32`

GetNFilings returns the NFilings field if non-nil, zero value otherwise.

### GetNFilingsOk

`func (o *BrandRevenueRow) GetNFilingsOk() (*int32, bool)`

GetNFilingsOk returns a tuple with the NFilings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFilings

`func (o *BrandRevenueRow) SetNFilings(v int32)`

SetNFilings sets NFilings field to given value.

### HasNFilings

`func (o *BrandRevenueRow) HasNFilings() bool`

HasNFilings returns a boolean if a field has been set.

### SetNFilingsNil

`func (o *BrandRevenueRow) SetNFilingsNil(b bool)`

 SetNFilingsNil sets the value for NFilings to be an explicit nil

### UnsetNFilings
`func (o *BrandRevenueRow) UnsetNFilings()`

UnsetNFilings ensures that no value is present for NFilings, not even an explicit nil
### GetRestated

`func (o *BrandRevenueRow) GetRestated() bool`

GetRestated returns the Restated field if non-nil, zero value otherwise.

### GetRestatedOk

`func (o *BrandRevenueRow) GetRestatedOk() (*bool, bool)`

GetRestatedOk returns a tuple with the Restated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestated

`func (o *BrandRevenueRow) SetRestated(v bool)`

SetRestated sets Restated field to given value.

### HasRestated

`func (o *BrandRevenueRow) HasRestated() bool`

HasRestated returns a boolean if a field has been set.

### SetRestatedNil

`func (o *BrandRevenueRow) SetRestatedNil(b bool)`

 SetRestatedNil sets the value for Restated to be an explicit nil

### UnsetRestated
`func (o *BrandRevenueRow) UnsetRestated()`

UnsetRestated ensures that no value is present for Restated, not even an explicit nil
### GetConfidence

`func (o *BrandRevenueRow) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *BrandRevenueRow) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *BrandRevenueRow) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *BrandRevenueRow) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *BrandRevenueRow) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *BrandRevenueRow) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetSourceUrl

`func (o *BrandRevenueRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *BrandRevenueRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *BrandRevenueRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *BrandRevenueRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *BrandRevenueRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *BrandRevenueRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetEvidenceQuote

`func (o *BrandRevenueRow) GetEvidenceQuote() string`

GetEvidenceQuote returns the EvidenceQuote field if non-nil, zero value otherwise.

### GetEvidenceQuoteOk

`func (o *BrandRevenueRow) GetEvidenceQuoteOk() (*string, bool)`

GetEvidenceQuoteOk returns a tuple with the EvidenceQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceQuote

`func (o *BrandRevenueRow) SetEvidenceQuote(v string)`

SetEvidenceQuote sets EvidenceQuote field to given value.

### HasEvidenceQuote

`func (o *BrandRevenueRow) HasEvidenceQuote() bool`

HasEvidenceQuote returns a boolean if a field has been set.

### SetEvidenceQuoteNil

`func (o *BrandRevenueRow) SetEvidenceQuoteNil(b bool)`

 SetEvidenceQuoteNil sets the value for EvidenceQuote to be an explicit nil

### UnsetEvidenceQuote
`func (o *BrandRevenueRow) UnsetEvidenceQuote()`

UnsetEvidenceQuote ensures that no value is present for EvidenceQuote, not even an explicit nil
### GetMatchMethod

`func (o *BrandRevenueRow) GetMatchMethod() string`

GetMatchMethod returns the MatchMethod field if non-nil, zero value otherwise.

### GetMatchMethodOk

`func (o *BrandRevenueRow) GetMatchMethodOk() (*string, bool)`

GetMatchMethodOk returns a tuple with the MatchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMethod

`func (o *BrandRevenueRow) SetMatchMethod(v string)`

SetMatchMethod sets MatchMethod field to given value.

### HasMatchMethod

`func (o *BrandRevenueRow) HasMatchMethod() bool`

HasMatchMethod returns a boolean if a field has been set.

### SetMatchMethodNil

`func (o *BrandRevenueRow) SetMatchMethodNil(b bool)`

 SetMatchMethodNil sets the value for MatchMethod to be an explicit nil

### UnsetMatchMethod
`func (o *BrandRevenueRow) UnsetMatchMethod()`

UnsetMatchMethod ensures that no value is present for MatchMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


