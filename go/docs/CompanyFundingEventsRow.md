# CompanyFundingEventsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FundingEventId** | Pointer to **NullableString** | The funding event key. | [optional] 
**IssuerName** | Pointer to **NullableString** | The issuer&#39;s name as filed. | [optional] 
**IssuerOrgId** | Pointer to **NullableString** | Legal-entity level: the issuer&#39;s organization record. | [optional] 
**IssuerCompanyId** | Pointer to **NullableString** | Company level: the issuer&#39;s company (null when the issuer did not resolve to a company). | [optional] 
**IssuerCompanyName** | Pointer to **NullableString** | Display name of the issuer&#39;s company. | [optional] 
**IssuerGroupId** | Pointer to **NullableString** | Group level: the issuer&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**IssuerGroupName** | Pointer to **NullableString** | Display name of the issuer&#39;s group. | [optional] 
**EventDate** | Pointer to **NullableString** | The date the source gives for the event; event_stage says what the date means. | [optional] 
**FiscalYear** | Pointer to **NullableInt32** | The fiscal year of the US National Institutes of Health (NIH), for a grant; null for the other sources. | [optional] 
**Source** | Pointer to **NullableString** | The source lane: reporter (NIH RePORTER grants), form_d (SEC Form D Regulation D offerings), edgar_offerings (SEC registered offerings) or sec_extracted (financings extracted from SEC filings). | [optional] 
**RoundType** | Pointer to **NullableString** | The kind of round (for example grant_nih, sbir, reg_d_private, ipo, follow_on, public_offering, atm, registered_direct, debt). | [optional] 
**RoundSubtype** | Pointer to **NullableString** | A finer type, such as the grant activity code, the Regulation D exemption or the SEC form. | [optional] 
**IsDilutive** | Pointer to **NullableBool** | Whether the round dilutes existing equity; null where the instrument is not classified. | [optional] 
**AmountUsd** | Pointer to **NullableFloat32** | The disclosed amount in whole US dollars; null where the amount is not disclosed. | [optional] 
**Currency** | Pointer to **NullableString** | The currency the source reports the amount in. | [optional] 
**AmountIsDisclosed** | Pointer to **NullableBool** | Whether the source states an amount for the event. | [optional] 
**CapitalTreatment** | Pointer to **NullableString** | Whether the event&#39;s dollars count as capital raised, and if not why not (for example structured_source, actual_inflow, capacity, committed). | [optional] 
**IsActualIssuerInflow** | Pointer to **NullableBool** | True where the event counts as capital the issuer raised. | [optional] 
**EventStage** | Pointer to **NullableString** | The lifecycle stage the event&#39;s date refers to (for example sold_to_date, funded, closed, registered, announced). | [optional] 
**AdministeringIc** | Pointer to **NullableString** | The National Institutes of Health institute or center that administers the grant (for example CA for the National Cancer Institute); null for the other sources. | [optional] 
**NAwards** | Pointer to **NullableInt64** | Grant applications folded into the event; 1 for the other sources. | [optional] 
**CumulativeRaisedUsd** | Pointer to **NullableFloat32** | The running total of disclosed amounts over the company&#39;s own events up to and including this one, in date order; null where the event has no company. | [optional] 
**SourceUrl** | Pointer to **NullableString** | Link to the filing the event was read from, where the source has one. | [optional] 

## Methods

### NewCompanyFundingEventsRow

`func NewCompanyFundingEventsRow() *CompanyFundingEventsRow`

NewCompanyFundingEventsRow instantiates a new CompanyFundingEventsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyFundingEventsRowWithDefaults

`func NewCompanyFundingEventsRowWithDefaults() *CompanyFundingEventsRow`

NewCompanyFundingEventsRowWithDefaults instantiates a new CompanyFundingEventsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFundingEventId

`func (o *CompanyFundingEventsRow) GetFundingEventId() string`

GetFundingEventId returns the FundingEventId field if non-nil, zero value otherwise.

### GetFundingEventIdOk

`func (o *CompanyFundingEventsRow) GetFundingEventIdOk() (*string, bool)`

GetFundingEventIdOk returns a tuple with the FundingEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingEventId

`func (o *CompanyFundingEventsRow) SetFundingEventId(v string)`

SetFundingEventId sets FundingEventId field to given value.

### HasFundingEventId

`func (o *CompanyFundingEventsRow) HasFundingEventId() bool`

HasFundingEventId returns a boolean if a field has been set.

### SetFundingEventIdNil

`func (o *CompanyFundingEventsRow) SetFundingEventIdNil(b bool)`

 SetFundingEventIdNil sets the value for FundingEventId to be an explicit nil

### UnsetFundingEventId
`func (o *CompanyFundingEventsRow) UnsetFundingEventId()`

UnsetFundingEventId ensures that no value is present for FundingEventId, not even an explicit nil
### GetIssuerName

`func (o *CompanyFundingEventsRow) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *CompanyFundingEventsRow) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *CompanyFundingEventsRow) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *CompanyFundingEventsRow) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### SetIssuerNameNil

`func (o *CompanyFundingEventsRow) SetIssuerNameNil(b bool)`

 SetIssuerNameNil sets the value for IssuerName to be an explicit nil

### UnsetIssuerName
`func (o *CompanyFundingEventsRow) UnsetIssuerName()`

UnsetIssuerName ensures that no value is present for IssuerName, not even an explicit nil
### GetIssuerOrgId

`func (o *CompanyFundingEventsRow) GetIssuerOrgId() string`

GetIssuerOrgId returns the IssuerOrgId field if non-nil, zero value otherwise.

### GetIssuerOrgIdOk

`func (o *CompanyFundingEventsRow) GetIssuerOrgIdOk() (*string, bool)`

GetIssuerOrgIdOk returns a tuple with the IssuerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerOrgId

`func (o *CompanyFundingEventsRow) SetIssuerOrgId(v string)`

SetIssuerOrgId sets IssuerOrgId field to given value.

### HasIssuerOrgId

`func (o *CompanyFundingEventsRow) HasIssuerOrgId() bool`

HasIssuerOrgId returns a boolean if a field has been set.

### SetIssuerOrgIdNil

`func (o *CompanyFundingEventsRow) SetIssuerOrgIdNil(b bool)`

 SetIssuerOrgIdNil sets the value for IssuerOrgId to be an explicit nil

### UnsetIssuerOrgId
`func (o *CompanyFundingEventsRow) UnsetIssuerOrgId()`

UnsetIssuerOrgId ensures that no value is present for IssuerOrgId, not even an explicit nil
### GetIssuerCompanyId

`func (o *CompanyFundingEventsRow) GetIssuerCompanyId() string`

GetIssuerCompanyId returns the IssuerCompanyId field if non-nil, zero value otherwise.

### GetIssuerCompanyIdOk

`func (o *CompanyFundingEventsRow) GetIssuerCompanyIdOk() (*string, bool)`

GetIssuerCompanyIdOk returns a tuple with the IssuerCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCompanyId

`func (o *CompanyFundingEventsRow) SetIssuerCompanyId(v string)`

SetIssuerCompanyId sets IssuerCompanyId field to given value.

### HasIssuerCompanyId

`func (o *CompanyFundingEventsRow) HasIssuerCompanyId() bool`

HasIssuerCompanyId returns a boolean if a field has been set.

### SetIssuerCompanyIdNil

`func (o *CompanyFundingEventsRow) SetIssuerCompanyIdNil(b bool)`

 SetIssuerCompanyIdNil sets the value for IssuerCompanyId to be an explicit nil

### UnsetIssuerCompanyId
`func (o *CompanyFundingEventsRow) UnsetIssuerCompanyId()`

UnsetIssuerCompanyId ensures that no value is present for IssuerCompanyId, not even an explicit nil
### GetIssuerCompanyName

`func (o *CompanyFundingEventsRow) GetIssuerCompanyName() string`

GetIssuerCompanyName returns the IssuerCompanyName field if non-nil, zero value otherwise.

### GetIssuerCompanyNameOk

`func (o *CompanyFundingEventsRow) GetIssuerCompanyNameOk() (*string, bool)`

GetIssuerCompanyNameOk returns a tuple with the IssuerCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCompanyName

`func (o *CompanyFundingEventsRow) SetIssuerCompanyName(v string)`

SetIssuerCompanyName sets IssuerCompanyName field to given value.

### HasIssuerCompanyName

`func (o *CompanyFundingEventsRow) HasIssuerCompanyName() bool`

HasIssuerCompanyName returns a boolean if a field has been set.

### SetIssuerCompanyNameNil

`func (o *CompanyFundingEventsRow) SetIssuerCompanyNameNil(b bool)`

 SetIssuerCompanyNameNil sets the value for IssuerCompanyName to be an explicit nil

### UnsetIssuerCompanyName
`func (o *CompanyFundingEventsRow) UnsetIssuerCompanyName()`

UnsetIssuerCompanyName ensures that no value is present for IssuerCompanyName, not even an explicit nil
### GetIssuerGroupId

`func (o *CompanyFundingEventsRow) GetIssuerGroupId() string`

GetIssuerGroupId returns the IssuerGroupId field if non-nil, zero value otherwise.

### GetIssuerGroupIdOk

`func (o *CompanyFundingEventsRow) GetIssuerGroupIdOk() (*string, bool)`

GetIssuerGroupIdOk returns a tuple with the IssuerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerGroupId

`func (o *CompanyFundingEventsRow) SetIssuerGroupId(v string)`

SetIssuerGroupId sets IssuerGroupId field to given value.

### HasIssuerGroupId

`func (o *CompanyFundingEventsRow) HasIssuerGroupId() bool`

HasIssuerGroupId returns a boolean if a field has been set.

### SetIssuerGroupIdNil

`func (o *CompanyFundingEventsRow) SetIssuerGroupIdNil(b bool)`

 SetIssuerGroupIdNil sets the value for IssuerGroupId to be an explicit nil

### UnsetIssuerGroupId
`func (o *CompanyFundingEventsRow) UnsetIssuerGroupId()`

UnsetIssuerGroupId ensures that no value is present for IssuerGroupId, not even an explicit nil
### GetIssuerGroupName

`func (o *CompanyFundingEventsRow) GetIssuerGroupName() string`

GetIssuerGroupName returns the IssuerGroupName field if non-nil, zero value otherwise.

### GetIssuerGroupNameOk

`func (o *CompanyFundingEventsRow) GetIssuerGroupNameOk() (*string, bool)`

GetIssuerGroupNameOk returns a tuple with the IssuerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerGroupName

`func (o *CompanyFundingEventsRow) SetIssuerGroupName(v string)`

SetIssuerGroupName sets IssuerGroupName field to given value.

### HasIssuerGroupName

`func (o *CompanyFundingEventsRow) HasIssuerGroupName() bool`

HasIssuerGroupName returns a boolean if a field has been set.

### SetIssuerGroupNameNil

`func (o *CompanyFundingEventsRow) SetIssuerGroupNameNil(b bool)`

 SetIssuerGroupNameNil sets the value for IssuerGroupName to be an explicit nil

### UnsetIssuerGroupName
`func (o *CompanyFundingEventsRow) UnsetIssuerGroupName()`

UnsetIssuerGroupName ensures that no value is present for IssuerGroupName, not even an explicit nil
### GetEventDate

`func (o *CompanyFundingEventsRow) GetEventDate() string`

GetEventDate returns the EventDate field if non-nil, zero value otherwise.

### GetEventDateOk

`func (o *CompanyFundingEventsRow) GetEventDateOk() (*string, bool)`

GetEventDateOk returns a tuple with the EventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDate

`func (o *CompanyFundingEventsRow) SetEventDate(v string)`

SetEventDate sets EventDate field to given value.

### HasEventDate

`func (o *CompanyFundingEventsRow) HasEventDate() bool`

HasEventDate returns a boolean if a field has been set.

### SetEventDateNil

`func (o *CompanyFundingEventsRow) SetEventDateNil(b bool)`

 SetEventDateNil sets the value for EventDate to be an explicit nil

### UnsetEventDate
`func (o *CompanyFundingEventsRow) UnsetEventDate()`

UnsetEventDate ensures that no value is present for EventDate, not even an explicit nil
### GetFiscalYear

`func (o *CompanyFundingEventsRow) GetFiscalYear() int32`

GetFiscalYear returns the FiscalYear field if non-nil, zero value otherwise.

### GetFiscalYearOk

`func (o *CompanyFundingEventsRow) GetFiscalYearOk() (*int32, bool)`

GetFiscalYearOk returns a tuple with the FiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYear

`func (o *CompanyFundingEventsRow) SetFiscalYear(v int32)`

SetFiscalYear sets FiscalYear field to given value.

### HasFiscalYear

`func (o *CompanyFundingEventsRow) HasFiscalYear() bool`

HasFiscalYear returns a boolean if a field has been set.

### SetFiscalYearNil

`func (o *CompanyFundingEventsRow) SetFiscalYearNil(b bool)`

 SetFiscalYearNil sets the value for FiscalYear to be an explicit nil

### UnsetFiscalYear
`func (o *CompanyFundingEventsRow) UnsetFiscalYear()`

UnsetFiscalYear ensures that no value is present for FiscalYear, not even an explicit nil
### GetSource

`func (o *CompanyFundingEventsRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CompanyFundingEventsRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CompanyFundingEventsRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CompanyFundingEventsRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *CompanyFundingEventsRow) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *CompanyFundingEventsRow) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetRoundType

`func (o *CompanyFundingEventsRow) GetRoundType() string`

GetRoundType returns the RoundType field if non-nil, zero value otherwise.

### GetRoundTypeOk

`func (o *CompanyFundingEventsRow) GetRoundTypeOk() (*string, bool)`

GetRoundTypeOk returns a tuple with the RoundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundType

`func (o *CompanyFundingEventsRow) SetRoundType(v string)`

SetRoundType sets RoundType field to given value.

### HasRoundType

`func (o *CompanyFundingEventsRow) HasRoundType() bool`

HasRoundType returns a boolean if a field has been set.

### SetRoundTypeNil

`func (o *CompanyFundingEventsRow) SetRoundTypeNil(b bool)`

 SetRoundTypeNil sets the value for RoundType to be an explicit nil

### UnsetRoundType
`func (o *CompanyFundingEventsRow) UnsetRoundType()`

UnsetRoundType ensures that no value is present for RoundType, not even an explicit nil
### GetRoundSubtype

`func (o *CompanyFundingEventsRow) GetRoundSubtype() string`

GetRoundSubtype returns the RoundSubtype field if non-nil, zero value otherwise.

### GetRoundSubtypeOk

`func (o *CompanyFundingEventsRow) GetRoundSubtypeOk() (*string, bool)`

GetRoundSubtypeOk returns a tuple with the RoundSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundSubtype

`func (o *CompanyFundingEventsRow) SetRoundSubtype(v string)`

SetRoundSubtype sets RoundSubtype field to given value.

### HasRoundSubtype

`func (o *CompanyFundingEventsRow) HasRoundSubtype() bool`

HasRoundSubtype returns a boolean if a field has been set.

### SetRoundSubtypeNil

`func (o *CompanyFundingEventsRow) SetRoundSubtypeNil(b bool)`

 SetRoundSubtypeNil sets the value for RoundSubtype to be an explicit nil

### UnsetRoundSubtype
`func (o *CompanyFundingEventsRow) UnsetRoundSubtype()`

UnsetRoundSubtype ensures that no value is present for RoundSubtype, not even an explicit nil
### GetIsDilutive

`func (o *CompanyFundingEventsRow) GetIsDilutive() bool`

GetIsDilutive returns the IsDilutive field if non-nil, zero value otherwise.

### GetIsDilutiveOk

`func (o *CompanyFundingEventsRow) GetIsDilutiveOk() (*bool, bool)`

GetIsDilutiveOk returns a tuple with the IsDilutive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDilutive

`func (o *CompanyFundingEventsRow) SetIsDilutive(v bool)`

SetIsDilutive sets IsDilutive field to given value.

### HasIsDilutive

`func (o *CompanyFundingEventsRow) HasIsDilutive() bool`

HasIsDilutive returns a boolean if a field has been set.

### SetIsDilutiveNil

`func (o *CompanyFundingEventsRow) SetIsDilutiveNil(b bool)`

 SetIsDilutiveNil sets the value for IsDilutive to be an explicit nil

### UnsetIsDilutive
`func (o *CompanyFundingEventsRow) UnsetIsDilutive()`

UnsetIsDilutive ensures that no value is present for IsDilutive, not even an explicit nil
### GetAmountUsd

`func (o *CompanyFundingEventsRow) GetAmountUsd() float32`

GetAmountUsd returns the AmountUsd field if non-nil, zero value otherwise.

### GetAmountUsdOk

`func (o *CompanyFundingEventsRow) GetAmountUsdOk() (*float32, bool)`

GetAmountUsdOk returns a tuple with the AmountUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountUsd

`func (o *CompanyFundingEventsRow) SetAmountUsd(v float32)`

SetAmountUsd sets AmountUsd field to given value.

### HasAmountUsd

`func (o *CompanyFundingEventsRow) HasAmountUsd() bool`

HasAmountUsd returns a boolean if a field has been set.

### SetAmountUsdNil

`func (o *CompanyFundingEventsRow) SetAmountUsdNil(b bool)`

 SetAmountUsdNil sets the value for AmountUsd to be an explicit nil

### UnsetAmountUsd
`func (o *CompanyFundingEventsRow) UnsetAmountUsd()`

UnsetAmountUsd ensures that no value is present for AmountUsd, not even an explicit nil
### GetCurrency

`func (o *CompanyFundingEventsRow) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CompanyFundingEventsRow) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CompanyFundingEventsRow) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CompanyFundingEventsRow) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *CompanyFundingEventsRow) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CompanyFundingEventsRow) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetAmountIsDisclosed

`func (o *CompanyFundingEventsRow) GetAmountIsDisclosed() bool`

GetAmountIsDisclosed returns the AmountIsDisclosed field if non-nil, zero value otherwise.

### GetAmountIsDisclosedOk

`func (o *CompanyFundingEventsRow) GetAmountIsDisclosedOk() (*bool, bool)`

GetAmountIsDisclosedOk returns a tuple with the AmountIsDisclosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountIsDisclosed

`func (o *CompanyFundingEventsRow) SetAmountIsDisclosed(v bool)`

SetAmountIsDisclosed sets AmountIsDisclosed field to given value.

### HasAmountIsDisclosed

`func (o *CompanyFundingEventsRow) HasAmountIsDisclosed() bool`

HasAmountIsDisclosed returns a boolean if a field has been set.

### SetAmountIsDisclosedNil

`func (o *CompanyFundingEventsRow) SetAmountIsDisclosedNil(b bool)`

 SetAmountIsDisclosedNil sets the value for AmountIsDisclosed to be an explicit nil

### UnsetAmountIsDisclosed
`func (o *CompanyFundingEventsRow) UnsetAmountIsDisclosed()`

UnsetAmountIsDisclosed ensures that no value is present for AmountIsDisclosed, not even an explicit nil
### GetCapitalTreatment

`func (o *CompanyFundingEventsRow) GetCapitalTreatment() string`

GetCapitalTreatment returns the CapitalTreatment field if non-nil, zero value otherwise.

### GetCapitalTreatmentOk

`func (o *CompanyFundingEventsRow) GetCapitalTreatmentOk() (*string, bool)`

GetCapitalTreatmentOk returns a tuple with the CapitalTreatment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalTreatment

`func (o *CompanyFundingEventsRow) SetCapitalTreatment(v string)`

SetCapitalTreatment sets CapitalTreatment field to given value.

### HasCapitalTreatment

`func (o *CompanyFundingEventsRow) HasCapitalTreatment() bool`

HasCapitalTreatment returns a boolean if a field has been set.

### SetCapitalTreatmentNil

`func (o *CompanyFundingEventsRow) SetCapitalTreatmentNil(b bool)`

 SetCapitalTreatmentNil sets the value for CapitalTreatment to be an explicit nil

### UnsetCapitalTreatment
`func (o *CompanyFundingEventsRow) UnsetCapitalTreatment()`

UnsetCapitalTreatment ensures that no value is present for CapitalTreatment, not even an explicit nil
### GetIsActualIssuerInflow

`func (o *CompanyFundingEventsRow) GetIsActualIssuerInflow() bool`

GetIsActualIssuerInflow returns the IsActualIssuerInflow field if non-nil, zero value otherwise.

### GetIsActualIssuerInflowOk

`func (o *CompanyFundingEventsRow) GetIsActualIssuerInflowOk() (*bool, bool)`

GetIsActualIssuerInflowOk returns a tuple with the IsActualIssuerInflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActualIssuerInflow

`func (o *CompanyFundingEventsRow) SetIsActualIssuerInflow(v bool)`

SetIsActualIssuerInflow sets IsActualIssuerInflow field to given value.

### HasIsActualIssuerInflow

`func (o *CompanyFundingEventsRow) HasIsActualIssuerInflow() bool`

HasIsActualIssuerInflow returns a boolean if a field has been set.

### SetIsActualIssuerInflowNil

`func (o *CompanyFundingEventsRow) SetIsActualIssuerInflowNil(b bool)`

 SetIsActualIssuerInflowNil sets the value for IsActualIssuerInflow to be an explicit nil

### UnsetIsActualIssuerInflow
`func (o *CompanyFundingEventsRow) UnsetIsActualIssuerInflow()`

UnsetIsActualIssuerInflow ensures that no value is present for IsActualIssuerInflow, not even an explicit nil
### GetEventStage

`func (o *CompanyFundingEventsRow) GetEventStage() string`

GetEventStage returns the EventStage field if non-nil, zero value otherwise.

### GetEventStageOk

`func (o *CompanyFundingEventsRow) GetEventStageOk() (*string, bool)`

GetEventStageOk returns a tuple with the EventStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStage

`func (o *CompanyFundingEventsRow) SetEventStage(v string)`

SetEventStage sets EventStage field to given value.

### HasEventStage

`func (o *CompanyFundingEventsRow) HasEventStage() bool`

HasEventStage returns a boolean if a field has been set.

### SetEventStageNil

`func (o *CompanyFundingEventsRow) SetEventStageNil(b bool)`

 SetEventStageNil sets the value for EventStage to be an explicit nil

### UnsetEventStage
`func (o *CompanyFundingEventsRow) UnsetEventStage()`

UnsetEventStage ensures that no value is present for EventStage, not even an explicit nil
### GetAdministeringIc

`func (o *CompanyFundingEventsRow) GetAdministeringIc() string`

GetAdministeringIc returns the AdministeringIc field if non-nil, zero value otherwise.

### GetAdministeringIcOk

`func (o *CompanyFundingEventsRow) GetAdministeringIcOk() (*string, bool)`

GetAdministeringIcOk returns a tuple with the AdministeringIc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministeringIc

`func (o *CompanyFundingEventsRow) SetAdministeringIc(v string)`

SetAdministeringIc sets AdministeringIc field to given value.

### HasAdministeringIc

`func (o *CompanyFundingEventsRow) HasAdministeringIc() bool`

HasAdministeringIc returns a boolean if a field has been set.

### SetAdministeringIcNil

`func (o *CompanyFundingEventsRow) SetAdministeringIcNil(b bool)`

 SetAdministeringIcNil sets the value for AdministeringIc to be an explicit nil

### UnsetAdministeringIc
`func (o *CompanyFundingEventsRow) UnsetAdministeringIc()`

UnsetAdministeringIc ensures that no value is present for AdministeringIc, not even an explicit nil
### GetNAwards

`func (o *CompanyFundingEventsRow) GetNAwards() int64`

GetNAwards returns the NAwards field if non-nil, zero value otherwise.

### GetNAwardsOk

`func (o *CompanyFundingEventsRow) GetNAwardsOk() (*int64, bool)`

GetNAwardsOk returns a tuple with the NAwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAwards

`func (o *CompanyFundingEventsRow) SetNAwards(v int64)`

SetNAwards sets NAwards field to given value.

### HasNAwards

`func (o *CompanyFundingEventsRow) HasNAwards() bool`

HasNAwards returns a boolean if a field has been set.

### SetNAwardsNil

`func (o *CompanyFundingEventsRow) SetNAwardsNil(b bool)`

 SetNAwardsNil sets the value for NAwards to be an explicit nil

### UnsetNAwards
`func (o *CompanyFundingEventsRow) UnsetNAwards()`

UnsetNAwards ensures that no value is present for NAwards, not even an explicit nil
### GetCumulativeRaisedUsd

`func (o *CompanyFundingEventsRow) GetCumulativeRaisedUsd() float32`

GetCumulativeRaisedUsd returns the CumulativeRaisedUsd field if non-nil, zero value otherwise.

### GetCumulativeRaisedUsdOk

`func (o *CompanyFundingEventsRow) GetCumulativeRaisedUsdOk() (*float32, bool)`

GetCumulativeRaisedUsdOk returns a tuple with the CumulativeRaisedUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeRaisedUsd

`func (o *CompanyFundingEventsRow) SetCumulativeRaisedUsd(v float32)`

SetCumulativeRaisedUsd sets CumulativeRaisedUsd field to given value.

### HasCumulativeRaisedUsd

`func (o *CompanyFundingEventsRow) HasCumulativeRaisedUsd() bool`

HasCumulativeRaisedUsd returns a boolean if a field has been set.

### SetCumulativeRaisedUsdNil

`func (o *CompanyFundingEventsRow) SetCumulativeRaisedUsdNil(b bool)`

 SetCumulativeRaisedUsdNil sets the value for CumulativeRaisedUsd to be an explicit nil

### UnsetCumulativeRaisedUsd
`func (o *CompanyFundingEventsRow) UnsetCumulativeRaisedUsd()`

UnsetCumulativeRaisedUsd ensures that no value is present for CumulativeRaisedUsd, not even an explicit nil
### GetSourceUrl

`func (o *CompanyFundingEventsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *CompanyFundingEventsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *CompanyFundingEventsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *CompanyFundingEventsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *CompanyFundingEventsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *CompanyFundingEventsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


