# OwnershipPercentagesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectOrgId** | Pointer to **NullableString** | Legal-entity level: the subject organization&#39;s record, a resolution root (the identity every ownership dataset keys on). One row per organization; an insider row with no organization has no row here. | [optional] 
**SubjectCompanyId** | Pointer to **NullableString** | Company level: the subject organization&#39;s company (null when the organization does not resolve to a company). | [optional] 
**SubjectCompanyName** | Pointer to **NullableString** | Display name of the subject organization&#39;s company. | [optional] 
**SubjectGroupId** | Pointer to **NullableString** | Group level: the subject organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SubjectGroupName** | Pointer to **NullableString** | Display name of the subject organization&#39;s group. | [optional] 
**AnyOwnershipSource** | Pointer to **NullableBool** | True when an insider row or a Form 13F quarter exists for the organization. | [optional] 
**SharesOutstandingState** | Pointer to **NullableString** | reported when the organization has a cover-page share count on file, not_reported when it has none (then no percentage is taken and the share figures are shown instead). | [optional] 
**SharesOutstanding** | Pointer to **NullableFloat32** | The denominator both percentages divide into: the latest cover-page share count (dei:EntityCommonStockSharesOutstanding) from an annual form. It is measured on the filing&#39;s cover date, 0 to 3 months after the fiscal period end, and the cover date is not published, so its vintage is the fiscal year in shares_outstanding_fy. | [optional] 
**SharesOutstandingFy** | Pointer to **NullableInt32** | The fiscal year of the filing the share count comes from. | [optional] 
**SharesOutstandingConcept** | Pointer to **NullableString** | The XBRL (eXtensible Business Reporting Language) concept the share count is read from. Always dei:EntityCommonStockSharesOutstanding. | [optional] 
**InsiderState** | Pointer to **NullableString** | How the insider figure was reached. reported: The sum of the current non-derivative balance each reporting owner last filed on a Form 3, 4 or 5, divided by the share count on file. no_insider_filing: No Form 3, 4 or 5 in our data names an insider of this company. That is a statement about coverage, not a statement that insiders hold nothing. all_positions_closed: Every insider position on file has been reported closed. This is a filed zero, not missing data. derivative_only: The insiders on file hold only derivative positions (options and similar), which are not outstanding shares. This is a filed zero, not missing data. no_denominator: No percentage: this company&#39;s SEC filings in our data carry no cover-page share count, so there is no denominator to divide by. The share count on file is shown instead. not_computable: No percentage is taken here. The reported share total is given instead, with the reason no ratio is taken. | [optional] 
**InsiderReason** | Pointer to **NullableString** | Why an insider ratio was not taken when insider_state is not_computable. denominator_stale: The share count on file is more than two years older than the positions it would be divided into, so the ratio would not describe this company. exceeds_shares_outstanding: The insider total on file is larger than the company&#39;s whole share count, which is arithmetically impossible for a single source and means the filings behind it have not consolidated correctly. Null otherwise. | [optional] 
**InsiderPercent** | Pointer to **NullableFloat32** | The share of shares outstanding held by Form 3, 4 and 5 insiders (insider_shares divided by shares_outstanding, times 100, unrounded); null unless insider_state is reported. | [optional] 
**InsiderShares** | Pointer to **NullableFloat32** | The sum of the current non-derivative balance each reporting insider last filed. A ceased or options-only insider adds nothing to it. Null when no insider holds standing shares. | [optional] 
**InsiderOwnerCount** | Pointer to **NullableInt32** | Insiders on file for the organization (0 when none). | [optional] 
**InsiderStandingOwnerCount** | Pointer to **NullableInt32** | Insiders with a positive standing equity balance, the only population insider_shares describes. | [optional] 
**InsiderClosedOwnerCount** | Pointer to **NullableInt32** | Insiders whose standing equity balance is zero or absent: closed positions plus options-only insiders. | [optional] 
**InsiderDerivativeOnlyOwnerCount** | Pointer to **NullableInt32** | Of the closed insiders, those whose only remaining positions are derivatives (options and similar), which are not outstanding shares. | [optional] 
**InsiderLastFilingDate** | Pointer to **NullableString** | The latest filing date among the organization&#39;s insider rows. | [optional] 
**InstitutionalState** | Pointer to **NullableString** | How the institutional figure was reached. reported: Every position reported on Form 13F for the quarter shown, summed once per filing manager, against the same share count as the insider figure. exceeds_outstanding: Reported positions exceed the shares outstanding on file. Shown as filed. implausible: No percentage: reported positions exceed 1.5 times the shares outstanding on file, which indicates a corporate action or a denominator mismatch rather than ownership. no_13f_issuer: No Form 13F issuer resolved to this company in any quarter. Positions are reported against a SECURITY, so this is a statement about the CUSIP-to-company bridge, not about institutional interest. no_reference_period: No manager reported this company for the most recent fully-reported quarter. Positions in earlier quarters are on file and are reported per manager, but a manager whose newest filing is for an older quarter is not a current holder and is not counted here. no_reference_shares: No manager reported an equity position in this company for the most recent fully-reported quarter. Only options, debt or non-equity issues were reported, and none of those is a share count. no_denominator: No percentage: this company&#39;s SEC filings in our data carry no cover-page share count, so there is no denominator to divide by. The reported share total is shown instead. not_computable: No percentage is taken here. The reported share total is given instead, with the reason no ratio is taken. | [optional] 
**InstitutionalReason** | Pointer to **NullableString** | Why an institutional ratio was not taken when institutional_state is not_computable. denominator_stale: The share count on file is more than two years older than the positions it would be divided into, so the ratio would not describe this company. Null otherwise. | [optional] 
**InstitutionalPercent** | Pointer to **NullableFloat32** | The share of shares outstanding under reported institutional discretion at the reference quarter (institutional_shares_reported divided by shares_outstanding, times 100, unrounded, as filed and never capped); null unless institutional_state is reported or exceeds_outstanding. | [optional] 
**InstitutionalSharesReported** | Pointer to **NullableFloat32** | The shares reported on Form 13F at the reference quarter, counted equity positions only, summed once per filing manager. Null when the organization has no reference-quarter row. | [optional] 
**InstitutionalManagerCount** | Pointer to **NullableInt32** | Managers reporting a counted equity position at the reference quarter (0 when there is no reference-quarter row). | [optional] 
**InstitutionalReportPeriod** | Pointer to **NullableString** | The reference quarter end the institutional figure is measured at: the most recent quarter in which at least 80 percent of the prior quarter&#39;s filing managers had filed. | [optional] 
**InstitutionalSharesOtr** | Pointer to **NullableFloat32** | Of institutional_shares_reported, the shares under other (OTR) discretion. | [optional] 
**InstitutionalSharesExcludedNonEquityIssue** | Pointer to **NullableFloat32** | Shares reported against convertible notes, debt and preferred issues at the reference quarter, excluded from institutional_shares_reported. | [optional] 
**InstitutionalEquityCusipCount** | Pointer to **NullableInt32** | Distinct equity CUSIPs (Committee on Uniform Securities Identification Procedures identifiers) counted at the reference quarter (0 when there is no reference-quarter row). | [optional] 
**InstitutionalLatestFilingDate** | Pointer to **NullableString** | The latest Form 13F filing date among the reference quarter&#39;s positions. | [optional] 

## Methods

### NewOwnershipPercentagesRow

`func NewOwnershipPercentagesRow() *OwnershipPercentagesRow`

NewOwnershipPercentagesRow instantiates a new OwnershipPercentagesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnershipPercentagesRowWithDefaults

`func NewOwnershipPercentagesRowWithDefaults() *OwnershipPercentagesRow`

NewOwnershipPercentagesRowWithDefaults instantiates a new OwnershipPercentagesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectOrgId

`func (o *OwnershipPercentagesRow) GetSubjectOrgId() string`

GetSubjectOrgId returns the SubjectOrgId field if non-nil, zero value otherwise.

### GetSubjectOrgIdOk

`func (o *OwnershipPercentagesRow) GetSubjectOrgIdOk() (*string, bool)`

GetSubjectOrgIdOk returns a tuple with the SubjectOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectOrgId

`func (o *OwnershipPercentagesRow) SetSubjectOrgId(v string)`

SetSubjectOrgId sets SubjectOrgId field to given value.

### HasSubjectOrgId

`func (o *OwnershipPercentagesRow) HasSubjectOrgId() bool`

HasSubjectOrgId returns a boolean if a field has been set.

### SetSubjectOrgIdNil

`func (o *OwnershipPercentagesRow) SetSubjectOrgIdNil(b bool)`

 SetSubjectOrgIdNil sets the value for SubjectOrgId to be an explicit nil

### UnsetSubjectOrgId
`func (o *OwnershipPercentagesRow) UnsetSubjectOrgId()`

UnsetSubjectOrgId ensures that no value is present for SubjectOrgId, not even an explicit nil
### GetSubjectCompanyId

`func (o *OwnershipPercentagesRow) GetSubjectCompanyId() string`

GetSubjectCompanyId returns the SubjectCompanyId field if non-nil, zero value otherwise.

### GetSubjectCompanyIdOk

`func (o *OwnershipPercentagesRow) GetSubjectCompanyIdOk() (*string, bool)`

GetSubjectCompanyIdOk returns a tuple with the SubjectCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCompanyId

`func (o *OwnershipPercentagesRow) SetSubjectCompanyId(v string)`

SetSubjectCompanyId sets SubjectCompanyId field to given value.

### HasSubjectCompanyId

`func (o *OwnershipPercentagesRow) HasSubjectCompanyId() bool`

HasSubjectCompanyId returns a boolean if a field has been set.

### SetSubjectCompanyIdNil

`func (o *OwnershipPercentagesRow) SetSubjectCompanyIdNil(b bool)`

 SetSubjectCompanyIdNil sets the value for SubjectCompanyId to be an explicit nil

### UnsetSubjectCompanyId
`func (o *OwnershipPercentagesRow) UnsetSubjectCompanyId()`

UnsetSubjectCompanyId ensures that no value is present for SubjectCompanyId, not even an explicit nil
### GetSubjectCompanyName

`func (o *OwnershipPercentagesRow) GetSubjectCompanyName() string`

GetSubjectCompanyName returns the SubjectCompanyName field if non-nil, zero value otherwise.

### GetSubjectCompanyNameOk

`func (o *OwnershipPercentagesRow) GetSubjectCompanyNameOk() (*string, bool)`

GetSubjectCompanyNameOk returns a tuple with the SubjectCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCompanyName

`func (o *OwnershipPercentagesRow) SetSubjectCompanyName(v string)`

SetSubjectCompanyName sets SubjectCompanyName field to given value.

### HasSubjectCompanyName

`func (o *OwnershipPercentagesRow) HasSubjectCompanyName() bool`

HasSubjectCompanyName returns a boolean if a field has been set.

### SetSubjectCompanyNameNil

`func (o *OwnershipPercentagesRow) SetSubjectCompanyNameNil(b bool)`

 SetSubjectCompanyNameNil sets the value for SubjectCompanyName to be an explicit nil

### UnsetSubjectCompanyName
`func (o *OwnershipPercentagesRow) UnsetSubjectCompanyName()`

UnsetSubjectCompanyName ensures that no value is present for SubjectCompanyName, not even an explicit nil
### GetSubjectGroupId

`func (o *OwnershipPercentagesRow) GetSubjectGroupId() string`

GetSubjectGroupId returns the SubjectGroupId field if non-nil, zero value otherwise.

### GetSubjectGroupIdOk

`func (o *OwnershipPercentagesRow) GetSubjectGroupIdOk() (*string, bool)`

GetSubjectGroupIdOk returns a tuple with the SubjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectGroupId

`func (o *OwnershipPercentagesRow) SetSubjectGroupId(v string)`

SetSubjectGroupId sets SubjectGroupId field to given value.

### HasSubjectGroupId

`func (o *OwnershipPercentagesRow) HasSubjectGroupId() bool`

HasSubjectGroupId returns a boolean if a field has been set.

### SetSubjectGroupIdNil

`func (o *OwnershipPercentagesRow) SetSubjectGroupIdNil(b bool)`

 SetSubjectGroupIdNil sets the value for SubjectGroupId to be an explicit nil

### UnsetSubjectGroupId
`func (o *OwnershipPercentagesRow) UnsetSubjectGroupId()`

UnsetSubjectGroupId ensures that no value is present for SubjectGroupId, not even an explicit nil
### GetSubjectGroupName

`func (o *OwnershipPercentagesRow) GetSubjectGroupName() string`

GetSubjectGroupName returns the SubjectGroupName field if non-nil, zero value otherwise.

### GetSubjectGroupNameOk

`func (o *OwnershipPercentagesRow) GetSubjectGroupNameOk() (*string, bool)`

GetSubjectGroupNameOk returns a tuple with the SubjectGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectGroupName

`func (o *OwnershipPercentagesRow) SetSubjectGroupName(v string)`

SetSubjectGroupName sets SubjectGroupName field to given value.

### HasSubjectGroupName

`func (o *OwnershipPercentagesRow) HasSubjectGroupName() bool`

HasSubjectGroupName returns a boolean if a field has been set.

### SetSubjectGroupNameNil

`func (o *OwnershipPercentagesRow) SetSubjectGroupNameNil(b bool)`

 SetSubjectGroupNameNil sets the value for SubjectGroupName to be an explicit nil

### UnsetSubjectGroupName
`func (o *OwnershipPercentagesRow) UnsetSubjectGroupName()`

UnsetSubjectGroupName ensures that no value is present for SubjectGroupName, not even an explicit nil
### GetAnyOwnershipSource

`func (o *OwnershipPercentagesRow) GetAnyOwnershipSource() bool`

GetAnyOwnershipSource returns the AnyOwnershipSource field if non-nil, zero value otherwise.

### GetAnyOwnershipSourceOk

`func (o *OwnershipPercentagesRow) GetAnyOwnershipSourceOk() (*bool, bool)`

GetAnyOwnershipSourceOk returns a tuple with the AnyOwnershipSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyOwnershipSource

`func (o *OwnershipPercentagesRow) SetAnyOwnershipSource(v bool)`

SetAnyOwnershipSource sets AnyOwnershipSource field to given value.

### HasAnyOwnershipSource

`func (o *OwnershipPercentagesRow) HasAnyOwnershipSource() bool`

HasAnyOwnershipSource returns a boolean if a field has been set.

### SetAnyOwnershipSourceNil

`func (o *OwnershipPercentagesRow) SetAnyOwnershipSourceNil(b bool)`

 SetAnyOwnershipSourceNil sets the value for AnyOwnershipSource to be an explicit nil

### UnsetAnyOwnershipSource
`func (o *OwnershipPercentagesRow) UnsetAnyOwnershipSource()`

UnsetAnyOwnershipSource ensures that no value is present for AnyOwnershipSource, not even an explicit nil
### GetSharesOutstandingState

`func (o *OwnershipPercentagesRow) GetSharesOutstandingState() string`

GetSharesOutstandingState returns the SharesOutstandingState field if non-nil, zero value otherwise.

### GetSharesOutstandingStateOk

`func (o *OwnershipPercentagesRow) GetSharesOutstandingStateOk() (*string, bool)`

GetSharesOutstandingStateOk returns a tuple with the SharesOutstandingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesOutstandingState

`func (o *OwnershipPercentagesRow) SetSharesOutstandingState(v string)`

SetSharesOutstandingState sets SharesOutstandingState field to given value.

### HasSharesOutstandingState

`func (o *OwnershipPercentagesRow) HasSharesOutstandingState() bool`

HasSharesOutstandingState returns a boolean if a field has been set.

### SetSharesOutstandingStateNil

`func (o *OwnershipPercentagesRow) SetSharesOutstandingStateNil(b bool)`

 SetSharesOutstandingStateNil sets the value for SharesOutstandingState to be an explicit nil

### UnsetSharesOutstandingState
`func (o *OwnershipPercentagesRow) UnsetSharesOutstandingState()`

UnsetSharesOutstandingState ensures that no value is present for SharesOutstandingState, not even an explicit nil
### GetSharesOutstanding

`func (o *OwnershipPercentagesRow) GetSharesOutstanding() float32`

GetSharesOutstanding returns the SharesOutstanding field if non-nil, zero value otherwise.

### GetSharesOutstandingOk

`func (o *OwnershipPercentagesRow) GetSharesOutstandingOk() (*float32, bool)`

GetSharesOutstandingOk returns a tuple with the SharesOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesOutstanding

`func (o *OwnershipPercentagesRow) SetSharesOutstanding(v float32)`

SetSharesOutstanding sets SharesOutstanding field to given value.

### HasSharesOutstanding

`func (o *OwnershipPercentagesRow) HasSharesOutstanding() bool`

HasSharesOutstanding returns a boolean if a field has been set.

### SetSharesOutstandingNil

`func (o *OwnershipPercentagesRow) SetSharesOutstandingNil(b bool)`

 SetSharesOutstandingNil sets the value for SharesOutstanding to be an explicit nil

### UnsetSharesOutstanding
`func (o *OwnershipPercentagesRow) UnsetSharesOutstanding()`

UnsetSharesOutstanding ensures that no value is present for SharesOutstanding, not even an explicit nil
### GetSharesOutstandingFy

`func (o *OwnershipPercentagesRow) GetSharesOutstandingFy() int32`

GetSharesOutstandingFy returns the SharesOutstandingFy field if non-nil, zero value otherwise.

### GetSharesOutstandingFyOk

`func (o *OwnershipPercentagesRow) GetSharesOutstandingFyOk() (*int32, bool)`

GetSharesOutstandingFyOk returns a tuple with the SharesOutstandingFy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesOutstandingFy

`func (o *OwnershipPercentagesRow) SetSharesOutstandingFy(v int32)`

SetSharesOutstandingFy sets SharesOutstandingFy field to given value.

### HasSharesOutstandingFy

`func (o *OwnershipPercentagesRow) HasSharesOutstandingFy() bool`

HasSharesOutstandingFy returns a boolean if a field has been set.

### SetSharesOutstandingFyNil

`func (o *OwnershipPercentagesRow) SetSharesOutstandingFyNil(b bool)`

 SetSharesOutstandingFyNil sets the value for SharesOutstandingFy to be an explicit nil

### UnsetSharesOutstandingFy
`func (o *OwnershipPercentagesRow) UnsetSharesOutstandingFy()`

UnsetSharesOutstandingFy ensures that no value is present for SharesOutstandingFy, not even an explicit nil
### GetSharesOutstandingConcept

`func (o *OwnershipPercentagesRow) GetSharesOutstandingConcept() string`

GetSharesOutstandingConcept returns the SharesOutstandingConcept field if non-nil, zero value otherwise.

### GetSharesOutstandingConceptOk

`func (o *OwnershipPercentagesRow) GetSharesOutstandingConceptOk() (*string, bool)`

GetSharesOutstandingConceptOk returns a tuple with the SharesOutstandingConcept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesOutstandingConcept

`func (o *OwnershipPercentagesRow) SetSharesOutstandingConcept(v string)`

SetSharesOutstandingConcept sets SharesOutstandingConcept field to given value.

### HasSharesOutstandingConcept

`func (o *OwnershipPercentagesRow) HasSharesOutstandingConcept() bool`

HasSharesOutstandingConcept returns a boolean if a field has been set.

### SetSharesOutstandingConceptNil

`func (o *OwnershipPercentagesRow) SetSharesOutstandingConceptNil(b bool)`

 SetSharesOutstandingConceptNil sets the value for SharesOutstandingConcept to be an explicit nil

### UnsetSharesOutstandingConcept
`func (o *OwnershipPercentagesRow) UnsetSharesOutstandingConcept()`

UnsetSharesOutstandingConcept ensures that no value is present for SharesOutstandingConcept, not even an explicit nil
### GetInsiderState

`func (o *OwnershipPercentagesRow) GetInsiderState() string`

GetInsiderState returns the InsiderState field if non-nil, zero value otherwise.

### GetInsiderStateOk

`func (o *OwnershipPercentagesRow) GetInsiderStateOk() (*string, bool)`

GetInsiderStateOk returns a tuple with the InsiderState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderState

`func (o *OwnershipPercentagesRow) SetInsiderState(v string)`

SetInsiderState sets InsiderState field to given value.

### HasInsiderState

`func (o *OwnershipPercentagesRow) HasInsiderState() bool`

HasInsiderState returns a boolean if a field has been set.

### SetInsiderStateNil

`func (o *OwnershipPercentagesRow) SetInsiderStateNil(b bool)`

 SetInsiderStateNil sets the value for InsiderState to be an explicit nil

### UnsetInsiderState
`func (o *OwnershipPercentagesRow) UnsetInsiderState()`

UnsetInsiderState ensures that no value is present for InsiderState, not even an explicit nil
### GetInsiderReason

`func (o *OwnershipPercentagesRow) GetInsiderReason() string`

GetInsiderReason returns the InsiderReason field if non-nil, zero value otherwise.

### GetInsiderReasonOk

`func (o *OwnershipPercentagesRow) GetInsiderReasonOk() (*string, bool)`

GetInsiderReasonOk returns a tuple with the InsiderReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderReason

`func (o *OwnershipPercentagesRow) SetInsiderReason(v string)`

SetInsiderReason sets InsiderReason field to given value.

### HasInsiderReason

`func (o *OwnershipPercentagesRow) HasInsiderReason() bool`

HasInsiderReason returns a boolean if a field has been set.

### SetInsiderReasonNil

`func (o *OwnershipPercentagesRow) SetInsiderReasonNil(b bool)`

 SetInsiderReasonNil sets the value for InsiderReason to be an explicit nil

### UnsetInsiderReason
`func (o *OwnershipPercentagesRow) UnsetInsiderReason()`

UnsetInsiderReason ensures that no value is present for InsiderReason, not even an explicit nil
### GetInsiderPercent

`func (o *OwnershipPercentagesRow) GetInsiderPercent() float32`

GetInsiderPercent returns the InsiderPercent field if non-nil, zero value otherwise.

### GetInsiderPercentOk

`func (o *OwnershipPercentagesRow) GetInsiderPercentOk() (*float32, bool)`

GetInsiderPercentOk returns a tuple with the InsiderPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderPercent

`func (o *OwnershipPercentagesRow) SetInsiderPercent(v float32)`

SetInsiderPercent sets InsiderPercent field to given value.

### HasInsiderPercent

`func (o *OwnershipPercentagesRow) HasInsiderPercent() bool`

HasInsiderPercent returns a boolean if a field has been set.

### SetInsiderPercentNil

`func (o *OwnershipPercentagesRow) SetInsiderPercentNil(b bool)`

 SetInsiderPercentNil sets the value for InsiderPercent to be an explicit nil

### UnsetInsiderPercent
`func (o *OwnershipPercentagesRow) UnsetInsiderPercent()`

UnsetInsiderPercent ensures that no value is present for InsiderPercent, not even an explicit nil
### GetInsiderShares

`func (o *OwnershipPercentagesRow) GetInsiderShares() float32`

GetInsiderShares returns the InsiderShares field if non-nil, zero value otherwise.

### GetInsiderSharesOk

`func (o *OwnershipPercentagesRow) GetInsiderSharesOk() (*float32, bool)`

GetInsiderSharesOk returns a tuple with the InsiderShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderShares

`func (o *OwnershipPercentagesRow) SetInsiderShares(v float32)`

SetInsiderShares sets InsiderShares field to given value.

### HasInsiderShares

`func (o *OwnershipPercentagesRow) HasInsiderShares() bool`

HasInsiderShares returns a boolean if a field has been set.

### SetInsiderSharesNil

`func (o *OwnershipPercentagesRow) SetInsiderSharesNil(b bool)`

 SetInsiderSharesNil sets the value for InsiderShares to be an explicit nil

### UnsetInsiderShares
`func (o *OwnershipPercentagesRow) UnsetInsiderShares()`

UnsetInsiderShares ensures that no value is present for InsiderShares, not even an explicit nil
### GetInsiderOwnerCount

`func (o *OwnershipPercentagesRow) GetInsiderOwnerCount() int32`

GetInsiderOwnerCount returns the InsiderOwnerCount field if non-nil, zero value otherwise.

### GetInsiderOwnerCountOk

`func (o *OwnershipPercentagesRow) GetInsiderOwnerCountOk() (*int32, bool)`

GetInsiderOwnerCountOk returns a tuple with the InsiderOwnerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderOwnerCount

`func (o *OwnershipPercentagesRow) SetInsiderOwnerCount(v int32)`

SetInsiderOwnerCount sets InsiderOwnerCount field to given value.

### HasInsiderOwnerCount

`func (o *OwnershipPercentagesRow) HasInsiderOwnerCount() bool`

HasInsiderOwnerCount returns a boolean if a field has been set.

### SetInsiderOwnerCountNil

`func (o *OwnershipPercentagesRow) SetInsiderOwnerCountNil(b bool)`

 SetInsiderOwnerCountNil sets the value for InsiderOwnerCount to be an explicit nil

### UnsetInsiderOwnerCount
`func (o *OwnershipPercentagesRow) UnsetInsiderOwnerCount()`

UnsetInsiderOwnerCount ensures that no value is present for InsiderOwnerCount, not even an explicit nil
### GetInsiderStandingOwnerCount

`func (o *OwnershipPercentagesRow) GetInsiderStandingOwnerCount() int32`

GetInsiderStandingOwnerCount returns the InsiderStandingOwnerCount field if non-nil, zero value otherwise.

### GetInsiderStandingOwnerCountOk

`func (o *OwnershipPercentagesRow) GetInsiderStandingOwnerCountOk() (*int32, bool)`

GetInsiderStandingOwnerCountOk returns a tuple with the InsiderStandingOwnerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderStandingOwnerCount

`func (o *OwnershipPercentagesRow) SetInsiderStandingOwnerCount(v int32)`

SetInsiderStandingOwnerCount sets InsiderStandingOwnerCount field to given value.

### HasInsiderStandingOwnerCount

`func (o *OwnershipPercentagesRow) HasInsiderStandingOwnerCount() bool`

HasInsiderStandingOwnerCount returns a boolean if a field has been set.

### SetInsiderStandingOwnerCountNil

`func (o *OwnershipPercentagesRow) SetInsiderStandingOwnerCountNil(b bool)`

 SetInsiderStandingOwnerCountNil sets the value for InsiderStandingOwnerCount to be an explicit nil

### UnsetInsiderStandingOwnerCount
`func (o *OwnershipPercentagesRow) UnsetInsiderStandingOwnerCount()`

UnsetInsiderStandingOwnerCount ensures that no value is present for InsiderStandingOwnerCount, not even an explicit nil
### GetInsiderClosedOwnerCount

`func (o *OwnershipPercentagesRow) GetInsiderClosedOwnerCount() int32`

GetInsiderClosedOwnerCount returns the InsiderClosedOwnerCount field if non-nil, zero value otherwise.

### GetInsiderClosedOwnerCountOk

`func (o *OwnershipPercentagesRow) GetInsiderClosedOwnerCountOk() (*int32, bool)`

GetInsiderClosedOwnerCountOk returns a tuple with the InsiderClosedOwnerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderClosedOwnerCount

`func (o *OwnershipPercentagesRow) SetInsiderClosedOwnerCount(v int32)`

SetInsiderClosedOwnerCount sets InsiderClosedOwnerCount field to given value.

### HasInsiderClosedOwnerCount

`func (o *OwnershipPercentagesRow) HasInsiderClosedOwnerCount() bool`

HasInsiderClosedOwnerCount returns a boolean if a field has been set.

### SetInsiderClosedOwnerCountNil

`func (o *OwnershipPercentagesRow) SetInsiderClosedOwnerCountNil(b bool)`

 SetInsiderClosedOwnerCountNil sets the value for InsiderClosedOwnerCount to be an explicit nil

### UnsetInsiderClosedOwnerCount
`func (o *OwnershipPercentagesRow) UnsetInsiderClosedOwnerCount()`

UnsetInsiderClosedOwnerCount ensures that no value is present for InsiderClosedOwnerCount, not even an explicit nil
### GetInsiderDerivativeOnlyOwnerCount

`func (o *OwnershipPercentagesRow) GetInsiderDerivativeOnlyOwnerCount() int32`

GetInsiderDerivativeOnlyOwnerCount returns the InsiderDerivativeOnlyOwnerCount field if non-nil, zero value otherwise.

### GetInsiderDerivativeOnlyOwnerCountOk

`func (o *OwnershipPercentagesRow) GetInsiderDerivativeOnlyOwnerCountOk() (*int32, bool)`

GetInsiderDerivativeOnlyOwnerCountOk returns a tuple with the InsiderDerivativeOnlyOwnerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderDerivativeOnlyOwnerCount

`func (o *OwnershipPercentagesRow) SetInsiderDerivativeOnlyOwnerCount(v int32)`

SetInsiderDerivativeOnlyOwnerCount sets InsiderDerivativeOnlyOwnerCount field to given value.

### HasInsiderDerivativeOnlyOwnerCount

`func (o *OwnershipPercentagesRow) HasInsiderDerivativeOnlyOwnerCount() bool`

HasInsiderDerivativeOnlyOwnerCount returns a boolean if a field has been set.

### SetInsiderDerivativeOnlyOwnerCountNil

`func (o *OwnershipPercentagesRow) SetInsiderDerivativeOnlyOwnerCountNil(b bool)`

 SetInsiderDerivativeOnlyOwnerCountNil sets the value for InsiderDerivativeOnlyOwnerCount to be an explicit nil

### UnsetInsiderDerivativeOnlyOwnerCount
`func (o *OwnershipPercentagesRow) UnsetInsiderDerivativeOnlyOwnerCount()`

UnsetInsiderDerivativeOnlyOwnerCount ensures that no value is present for InsiderDerivativeOnlyOwnerCount, not even an explicit nil
### GetInsiderLastFilingDate

`func (o *OwnershipPercentagesRow) GetInsiderLastFilingDate() string`

GetInsiderLastFilingDate returns the InsiderLastFilingDate field if non-nil, zero value otherwise.

### GetInsiderLastFilingDateOk

`func (o *OwnershipPercentagesRow) GetInsiderLastFilingDateOk() (*string, bool)`

GetInsiderLastFilingDateOk returns a tuple with the InsiderLastFilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderLastFilingDate

`func (o *OwnershipPercentagesRow) SetInsiderLastFilingDate(v string)`

SetInsiderLastFilingDate sets InsiderLastFilingDate field to given value.

### HasInsiderLastFilingDate

`func (o *OwnershipPercentagesRow) HasInsiderLastFilingDate() bool`

HasInsiderLastFilingDate returns a boolean if a field has been set.

### SetInsiderLastFilingDateNil

`func (o *OwnershipPercentagesRow) SetInsiderLastFilingDateNil(b bool)`

 SetInsiderLastFilingDateNil sets the value for InsiderLastFilingDate to be an explicit nil

### UnsetInsiderLastFilingDate
`func (o *OwnershipPercentagesRow) UnsetInsiderLastFilingDate()`

UnsetInsiderLastFilingDate ensures that no value is present for InsiderLastFilingDate, not even an explicit nil
### GetInstitutionalState

`func (o *OwnershipPercentagesRow) GetInstitutionalState() string`

GetInstitutionalState returns the InstitutionalState field if non-nil, zero value otherwise.

### GetInstitutionalStateOk

`func (o *OwnershipPercentagesRow) GetInstitutionalStateOk() (*string, bool)`

GetInstitutionalStateOk returns a tuple with the InstitutionalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionalState

`func (o *OwnershipPercentagesRow) SetInstitutionalState(v string)`

SetInstitutionalState sets InstitutionalState field to given value.

### HasInstitutionalState

`func (o *OwnershipPercentagesRow) HasInstitutionalState() bool`

HasInstitutionalState returns a boolean if a field has been set.

### SetInstitutionalStateNil

`func (o *OwnershipPercentagesRow) SetInstitutionalStateNil(b bool)`

 SetInstitutionalStateNil sets the value for InstitutionalState to be an explicit nil

### UnsetInstitutionalState
`func (o *OwnershipPercentagesRow) UnsetInstitutionalState()`

UnsetInstitutionalState ensures that no value is present for InstitutionalState, not even an explicit nil
### GetInstitutionalReason

`func (o *OwnershipPercentagesRow) GetInstitutionalReason() string`

GetInstitutionalReason returns the InstitutionalReason field if non-nil, zero value otherwise.

### GetInstitutionalReasonOk

`func (o *OwnershipPercentagesRow) GetInstitutionalReasonOk() (*string, bool)`

GetInstitutionalReasonOk returns a tuple with the InstitutionalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionalReason

`func (o *OwnershipPercentagesRow) SetInstitutionalReason(v string)`

SetInstitutionalReason sets InstitutionalReason field to given value.

### HasInstitutionalReason

`func (o *OwnershipPercentagesRow) HasInstitutionalReason() bool`

HasInstitutionalReason returns a boolean if a field has been set.

### SetInstitutionalReasonNil

`func (o *OwnershipPercentagesRow) SetInstitutionalReasonNil(b bool)`

 SetInstitutionalReasonNil sets the value for InstitutionalReason to be an explicit nil

### UnsetInstitutionalReason
`func (o *OwnershipPercentagesRow) UnsetInstitutionalReason()`

UnsetInstitutionalReason ensures that no value is present for InstitutionalReason, not even an explicit nil
### GetInstitutionalPercent

`func (o *OwnershipPercentagesRow) GetInstitutionalPercent() float32`

GetInstitutionalPercent returns the InstitutionalPercent field if non-nil, zero value otherwise.

### GetInstitutionalPercentOk

`func (o *OwnershipPercentagesRow) GetInstitutionalPercentOk() (*float32, bool)`

GetInstitutionalPercentOk returns a tuple with the InstitutionalPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionalPercent

`func (o *OwnershipPercentagesRow) SetInstitutionalPercent(v float32)`

SetInstitutionalPercent sets InstitutionalPercent field to given value.

### HasInstitutionalPercent

`func (o *OwnershipPercentagesRow) HasInstitutionalPercent() bool`

HasInstitutionalPercent returns a boolean if a field has been set.

### SetInstitutionalPercentNil

`func (o *OwnershipPercentagesRow) SetInstitutionalPercentNil(b bool)`

 SetInstitutionalPercentNil sets the value for InstitutionalPercent to be an explicit nil

### UnsetInstitutionalPercent
`func (o *OwnershipPercentagesRow) UnsetInstitutionalPercent()`

UnsetInstitutionalPercent ensures that no value is present for InstitutionalPercent, not even an explicit nil
### GetInstitutionalSharesReported

`func (o *OwnershipPercentagesRow) GetInstitutionalSharesReported() float32`

GetInstitutionalSharesReported returns the InstitutionalSharesReported field if non-nil, zero value otherwise.

### GetInstitutionalSharesReportedOk

`func (o *OwnershipPercentagesRow) GetInstitutionalSharesReportedOk() (*float32, bool)`

GetInstitutionalSharesReportedOk returns a tuple with the InstitutionalSharesReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionalSharesReported

`func (o *OwnershipPercentagesRow) SetInstitutionalSharesReported(v float32)`

SetInstitutionalSharesReported sets InstitutionalSharesReported field to given value.

### HasInstitutionalSharesReported

`func (o *OwnershipPercentagesRow) HasInstitutionalSharesReported() bool`

HasInstitutionalSharesReported returns a boolean if a field has been set.

### SetInstitutionalSharesReportedNil

`func (o *OwnershipPercentagesRow) SetInstitutionalSharesReportedNil(b bool)`

 SetInstitutionalSharesReportedNil sets the value for InstitutionalSharesReported to be an explicit nil

### UnsetInstitutionalSharesReported
`func (o *OwnershipPercentagesRow) UnsetInstitutionalSharesReported()`

UnsetInstitutionalSharesReported ensures that no value is present for InstitutionalSharesReported, not even an explicit nil
### GetInstitutionalManagerCount

`func (o *OwnershipPercentagesRow) GetInstitutionalManagerCount() int32`

GetInstitutionalManagerCount returns the InstitutionalManagerCount field if non-nil, zero value otherwise.

### GetInstitutionalManagerCountOk

`func (o *OwnershipPercentagesRow) GetInstitutionalManagerCountOk() (*int32, bool)`

GetInstitutionalManagerCountOk returns a tuple with the InstitutionalManagerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionalManagerCount

`func (o *OwnershipPercentagesRow) SetInstitutionalManagerCount(v int32)`

SetInstitutionalManagerCount sets InstitutionalManagerCount field to given value.

### HasInstitutionalManagerCount

`func (o *OwnershipPercentagesRow) HasInstitutionalManagerCount() bool`

HasInstitutionalManagerCount returns a boolean if a field has been set.

### SetInstitutionalManagerCountNil

`func (o *OwnershipPercentagesRow) SetInstitutionalManagerCountNil(b bool)`

 SetInstitutionalManagerCountNil sets the value for InstitutionalManagerCount to be an explicit nil

### UnsetInstitutionalManagerCount
`func (o *OwnershipPercentagesRow) UnsetInstitutionalManagerCount()`

UnsetInstitutionalManagerCount ensures that no value is present for InstitutionalManagerCount, not even an explicit nil
### GetInstitutionalReportPeriod

`func (o *OwnershipPercentagesRow) GetInstitutionalReportPeriod() string`

GetInstitutionalReportPeriod returns the InstitutionalReportPeriod field if non-nil, zero value otherwise.

### GetInstitutionalReportPeriodOk

`func (o *OwnershipPercentagesRow) GetInstitutionalReportPeriodOk() (*string, bool)`

GetInstitutionalReportPeriodOk returns a tuple with the InstitutionalReportPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionalReportPeriod

`func (o *OwnershipPercentagesRow) SetInstitutionalReportPeriod(v string)`

SetInstitutionalReportPeriod sets InstitutionalReportPeriod field to given value.

### HasInstitutionalReportPeriod

`func (o *OwnershipPercentagesRow) HasInstitutionalReportPeriod() bool`

HasInstitutionalReportPeriod returns a boolean if a field has been set.

### SetInstitutionalReportPeriodNil

`func (o *OwnershipPercentagesRow) SetInstitutionalReportPeriodNil(b bool)`

 SetInstitutionalReportPeriodNil sets the value for InstitutionalReportPeriod to be an explicit nil

### UnsetInstitutionalReportPeriod
`func (o *OwnershipPercentagesRow) UnsetInstitutionalReportPeriod()`

UnsetInstitutionalReportPeriod ensures that no value is present for InstitutionalReportPeriod, not even an explicit nil
### GetInstitutionalSharesOtr

`func (o *OwnershipPercentagesRow) GetInstitutionalSharesOtr() float32`

GetInstitutionalSharesOtr returns the InstitutionalSharesOtr field if non-nil, zero value otherwise.

### GetInstitutionalSharesOtrOk

`func (o *OwnershipPercentagesRow) GetInstitutionalSharesOtrOk() (*float32, bool)`

GetInstitutionalSharesOtrOk returns a tuple with the InstitutionalSharesOtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionalSharesOtr

`func (o *OwnershipPercentagesRow) SetInstitutionalSharesOtr(v float32)`

SetInstitutionalSharesOtr sets InstitutionalSharesOtr field to given value.

### HasInstitutionalSharesOtr

`func (o *OwnershipPercentagesRow) HasInstitutionalSharesOtr() bool`

HasInstitutionalSharesOtr returns a boolean if a field has been set.

### SetInstitutionalSharesOtrNil

`func (o *OwnershipPercentagesRow) SetInstitutionalSharesOtrNil(b bool)`

 SetInstitutionalSharesOtrNil sets the value for InstitutionalSharesOtr to be an explicit nil

### UnsetInstitutionalSharesOtr
`func (o *OwnershipPercentagesRow) UnsetInstitutionalSharesOtr()`

UnsetInstitutionalSharesOtr ensures that no value is present for InstitutionalSharesOtr, not even an explicit nil
### GetInstitutionalSharesExcludedNonEquityIssue

`func (o *OwnershipPercentagesRow) GetInstitutionalSharesExcludedNonEquityIssue() float32`

GetInstitutionalSharesExcludedNonEquityIssue returns the InstitutionalSharesExcludedNonEquityIssue field if non-nil, zero value otherwise.

### GetInstitutionalSharesExcludedNonEquityIssueOk

`func (o *OwnershipPercentagesRow) GetInstitutionalSharesExcludedNonEquityIssueOk() (*float32, bool)`

GetInstitutionalSharesExcludedNonEquityIssueOk returns a tuple with the InstitutionalSharesExcludedNonEquityIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionalSharesExcludedNonEquityIssue

`func (o *OwnershipPercentagesRow) SetInstitutionalSharesExcludedNonEquityIssue(v float32)`

SetInstitutionalSharesExcludedNonEquityIssue sets InstitutionalSharesExcludedNonEquityIssue field to given value.

### HasInstitutionalSharesExcludedNonEquityIssue

`func (o *OwnershipPercentagesRow) HasInstitutionalSharesExcludedNonEquityIssue() bool`

HasInstitutionalSharesExcludedNonEquityIssue returns a boolean if a field has been set.

### SetInstitutionalSharesExcludedNonEquityIssueNil

`func (o *OwnershipPercentagesRow) SetInstitutionalSharesExcludedNonEquityIssueNil(b bool)`

 SetInstitutionalSharesExcludedNonEquityIssueNil sets the value for InstitutionalSharesExcludedNonEquityIssue to be an explicit nil

### UnsetInstitutionalSharesExcludedNonEquityIssue
`func (o *OwnershipPercentagesRow) UnsetInstitutionalSharesExcludedNonEquityIssue()`

UnsetInstitutionalSharesExcludedNonEquityIssue ensures that no value is present for InstitutionalSharesExcludedNonEquityIssue, not even an explicit nil
### GetInstitutionalEquityCusipCount

`func (o *OwnershipPercentagesRow) GetInstitutionalEquityCusipCount() int32`

GetInstitutionalEquityCusipCount returns the InstitutionalEquityCusipCount field if non-nil, zero value otherwise.

### GetInstitutionalEquityCusipCountOk

`func (o *OwnershipPercentagesRow) GetInstitutionalEquityCusipCountOk() (*int32, bool)`

GetInstitutionalEquityCusipCountOk returns a tuple with the InstitutionalEquityCusipCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionalEquityCusipCount

`func (o *OwnershipPercentagesRow) SetInstitutionalEquityCusipCount(v int32)`

SetInstitutionalEquityCusipCount sets InstitutionalEquityCusipCount field to given value.

### HasInstitutionalEquityCusipCount

`func (o *OwnershipPercentagesRow) HasInstitutionalEquityCusipCount() bool`

HasInstitutionalEquityCusipCount returns a boolean if a field has been set.

### SetInstitutionalEquityCusipCountNil

`func (o *OwnershipPercentagesRow) SetInstitutionalEquityCusipCountNil(b bool)`

 SetInstitutionalEquityCusipCountNil sets the value for InstitutionalEquityCusipCount to be an explicit nil

### UnsetInstitutionalEquityCusipCount
`func (o *OwnershipPercentagesRow) UnsetInstitutionalEquityCusipCount()`

UnsetInstitutionalEquityCusipCount ensures that no value is present for InstitutionalEquityCusipCount, not even an explicit nil
### GetInstitutionalLatestFilingDate

`func (o *OwnershipPercentagesRow) GetInstitutionalLatestFilingDate() string`

GetInstitutionalLatestFilingDate returns the InstitutionalLatestFilingDate field if non-nil, zero value otherwise.

### GetInstitutionalLatestFilingDateOk

`func (o *OwnershipPercentagesRow) GetInstitutionalLatestFilingDateOk() (*string, bool)`

GetInstitutionalLatestFilingDateOk returns a tuple with the InstitutionalLatestFilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionalLatestFilingDate

`func (o *OwnershipPercentagesRow) SetInstitutionalLatestFilingDate(v string)`

SetInstitutionalLatestFilingDate sets InstitutionalLatestFilingDate field to given value.

### HasInstitutionalLatestFilingDate

`func (o *OwnershipPercentagesRow) HasInstitutionalLatestFilingDate() bool`

HasInstitutionalLatestFilingDate returns a boolean if a field has been set.

### SetInstitutionalLatestFilingDateNil

`func (o *OwnershipPercentagesRow) SetInstitutionalLatestFilingDateNil(b bool)`

 SetInstitutionalLatestFilingDateNil sets the value for InstitutionalLatestFilingDate to be an explicit nil

### UnsetInstitutionalLatestFilingDate
`func (o *OwnershipPercentagesRow) UnsetInstitutionalLatestFilingDate()`

UnsetInstitutionalLatestFilingDate ensures that no value is present for InstitutionalLatestFilingDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


