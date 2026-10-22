# InstitutionalQuarterTotalsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuarterTotalId** | Pointer to **NullableString** | The issuer and report period id. | [optional] 
**IssuerOrgId** | Pointer to **NullableString** | Legal-entity level: the issuer organization as the ownership identity seam resolves it, a resolution root; a listed subsidiary&#39;s securities may be keyed at its resolved parent. | [optional] 
**IssuerCompanyId** | Pointer to **NullableString** | Company level: the issuer&#39;s company (null when the organization does not resolve to a company). | [optional] 
**IssuerCompanyName** | Pointer to **NullableString** | Display name of the issuer&#39;s company. | [optional] 
**IssuerGroupId** | Pointer to **NullableString** | Group level: the issuer&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**IssuerGroupName** | Pointer to **NullableString** | Display name of the issuer&#39;s group. | [optional] 
**ReportPeriod** | Pointer to **NullableString** | The Form 13F report period (quarter end) the total is measured at. | [optional] 
**IsReferencePeriod** | Pointer to **NullableBool** | True for the one quarter the ownership percentage reads: the most recent in which at least 80 percent of the prior quarter&#39;s filing managers had filed. | [optional] 
**ManagerCount** | Pointer to **NullableInt64** | Managers reporting a counted equity position in the issuer at the period. | [optional] 
**PositionCount** | Pointer to **NullableInt64** | Counted equity positions reported against the issuer at the period. | [optional] 
**LanePositionCount** | Pointer to **NullableInt64** | Every position reported against the issuer at the period, counted or excluded. | [optional] 
**SharesReported** | Pointer to **NullableFloat32** | Shares reported on Form 13F for the issuer at the period: counted equity positions only, summed within this issuer and period. | [optional] 
**SharesOtr** | Pointer to **NullableFloat32** | Of shares_reported, the shares under other (OTR) discretion. | [optional] 
**SharesExcludedNonEquityIssue** | Pointer to **NullableFloat32** | Shares reported against convertible notes, debt and preferred issues, excluded from shares_reported. | [optional] 
**ExcludedNonEquityIssuePositionCount** | Pointer to **NullableInt64** | Positions excluded as non-equity issues. | [optional] 
**ExcludedDerivativePositionCount** | Pointer to **NullableInt64** | Positions excluded because they are puts or calls. | [optional] 
**ExcludedPrincipalPositionCount** | Pointer to **NullableInt64** | Positions excluded because they are reported as a principal amount rather than shares. | [optional] 
**ExcludedNonEquityCusipCount** | Pointer to **NullableInt64** | Distinct CUSIPs excluded as non-equity issues. | [optional] 
**EquityCusipCount** | Pointer to **NullableInt64** | Distinct equity CUSIPs counted at the period. | [optional] 
**LatestFilingDate** | Pointer to **NullableString** | The latest Form 13F filing date among the period&#39;s positions. | [optional] 
**AggregationRule** | Pointer to **NullableString** | The rule the total was built under; always sum_within_issuer_period_only. | [optional] 

## Methods

### NewInstitutionalQuarterTotalsRow

`func NewInstitutionalQuarterTotalsRow() *InstitutionalQuarterTotalsRow`

NewInstitutionalQuarterTotalsRow instantiates a new InstitutionalQuarterTotalsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionalQuarterTotalsRowWithDefaults

`func NewInstitutionalQuarterTotalsRowWithDefaults() *InstitutionalQuarterTotalsRow`

NewInstitutionalQuarterTotalsRowWithDefaults instantiates a new InstitutionalQuarterTotalsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuarterTotalId

`func (o *InstitutionalQuarterTotalsRow) GetQuarterTotalId() string`

GetQuarterTotalId returns the QuarterTotalId field if non-nil, zero value otherwise.

### GetQuarterTotalIdOk

`func (o *InstitutionalQuarterTotalsRow) GetQuarterTotalIdOk() (*string, bool)`

GetQuarterTotalIdOk returns a tuple with the QuarterTotalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarterTotalId

`func (o *InstitutionalQuarterTotalsRow) SetQuarterTotalId(v string)`

SetQuarterTotalId sets QuarterTotalId field to given value.

### HasQuarterTotalId

`func (o *InstitutionalQuarterTotalsRow) HasQuarterTotalId() bool`

HasQuarterTotalId returns a boolean if a field has been set.

### SetQuarterTotalIdNil

`func (o *InstitutionalQuarterTotalsRow) SetQuarterTotalIdNil(b bool)`

 SetQuarterTotalIdNil sets the value for QuarterTotalId to be an explicit nil

### UnsetQuarterTotalId
`func (o *InstitutionalQuarterTotalsRow) UnsetQuarterTotalId()`

UnsetQuarterTotalId ensures that no value is present for QuarterTotalId, not even an explicit nil
### GetIssuerOrgId

`func (o *InstitutionalQuarterTotalsRow) GetIssuerOrgId() string`

GetIssuerOrgId returns the IssuerOrgId field if non-nil, zero value otherwise.

### GetIssuerOrgIdOk

`func (o *InstitutionalQuarterTotalsRow) GetIssuerOrgIdOk() (*string, bool)`

GetIssuerOrgIdOk returns a tuple with the IssuerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerOrgId

`func (o *InstitutionalQuarterTotalsRow) SetIssuerOrgId(v string)`

SetIssuerOrgId sets IssuerOrgId field to given value.

### HasIssuerOrgId

`func (o *InstitutionalQuarterTotalsRow) HasIssuerOrgId() bool`

HasIssuerOrgId returns a boolean if a field has been set.

### SetIssuerOrgIdNil

`func (o *InstitutionalQuarterTotalsRow) SetIssuerOrgIdNil(b bool)`

 SetIssuerOrgIdNil sets the value for IssuerOrgId to be an explicit nil

### UnsetIssuerOrgId
`func (o *InstitutionalQuarterTotalsRow) UnsetIssuerOrgId()`

UnsetIssuerOrgId ensures that no value is present for IssuerOrgId, not even an explicit nil
### GetIssuerCompanyId

`func (o *InstitutionalQuarterTotalsRow) GetIssuerCompanyId() string`

GetIssuerCompanyId returns the IssuerCompanyId field if non-nil, zero value otherwise.

### GetIssuerCompanyIdOk

`func (o *InstitutionalQuarterTotalsRow) GetIssuerCompanyIdOk() (*string, bool)`

GetIssuerCompanyIdOk returns a tuple with the IssuerCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCompanyId

`func (o *InstitutionalQuarterTotalsRow) SetIssuerCompanyId(v string)`

SetIssuerCompanyId sets IssuerCompanyId field to given value.

### HasIssuerCompanyId

`func (o *InstitutionalQuarterTotalsRow) HasIssuerCompanyId() bool`

HasIssuerCompanyId returns a boolean if a field has been set.

### SetIssuerCompanyIdNil

`func (o *InstitutionalQuarterTotalsRow) SetIssuerCompanyIdNil(b bool)`

 SetIssuerCompanyIdNil sets the value for IssuerCompanyId to be an explicit nil

### UnsetIssuerCompanyId
`func (o *InstitutionalQuarterTotalsRow) UnsetIssuerCompanyId()`

UnsetIssuerCompanyId ensures that no value is present for IssuerCompanyId, not even an explicit nil
### GetIssuerCompanyName

`func (o *InstitutionalQuarterTotalsRow) GetIssuerCompanyName() string`

GetIssuerCompanyName returns the IssuerCompanyName field if non-nil, zero value otherwise.

### GetIssuerCompanyNameOk

`func (o *InstitutionalQuarterTotalsRow) GetIssuerCompanyNameOk() (*string, bool)`

GetIssuerCompanyNameOk returns a tuple with the IssuerCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCompanyName

`func (o *InstitutionalQuarterTotalsRow) SetIssuerCompanyName(v string)`

SetIssuerCompanyName sets IssuerCompanyName field to given value.

### HasIssuerCompanyName

`func (o *InstitutionalQuarterTotalsRow) HasIssuerCompanyName() bool`

HasIssuerCompanyName returns a boolean if a field has been set.

### SetIssuerCompanyNameNil

`func (o *InstitutionalQuarterTotalsRow) SetIssuerCompanyNameNil(b bool)`

 SetIssuerCompanyNameNil sets the value for IssuerCompanyName to be an explicit nil

### UnsetIssuerCompanyName
`func (o *InstitutionalQuarterTotalsRow) UnsetIssuerCompanyName()`

UnsetIssuerCompanyName ensures that no value is present for IssuerCompanyName, not even an explicit nil
### GetIssuerGroupId

`func (o *InstitutionalQuarterTotalsRow) GetIssuerGroupId() string`

GetIssuerGroupId returns the IssuerGroupId field if non-nil, zero value otherwise.

### GetIssuerGroupIdOk

`func (o *InstitutionalQuarterTotalsRow) GetIssuerGroupIdOk() (*string, bool)`

GetIssuerGroupIdOk returns a tuple with the IssuerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerGroupId

`func (o *InstitutionalQuarterTotalsRow) SetIssuerGroupId(v string)`

SetIssuerGroupId sets IssuerGroupId field to given value.

### HasIssuerGroupId

`func (o *InstitutionalQuarterTotalsRow) HasIssuerGroupId() bool`

HasIssuerGroupId returns a boolean if a field has been set.

### SetIssuerGroupIdNil

`func (o *InstitutionalQuarterTotalsRow) SetIssuerGroupIdNil(b bool)`

 SetIssuerGroupIdNil sets the value for IssuerGroupId to be an explicit nil

### UnsetIssuerGroupId
`func (o *InstitutionalQuarterTotalsRow) UnsetIssuerGroupId()`

UnsetIssuerGroupId ensures that no value is present for IssuerGroupId, not even an explicit nil
### GetIssuerGroupName

`func (o *InstitutionalQuarterTotalsRow) GetIssuerGroupName() string`

GetIssuerGroupName returns the IssuerGroupName field if non-nil, zero value otherwise.

### GetIssuerGroupNameOk

`func (o *InstitutionalQuarterTotalsRow) GetIssuerGroupNameOk() (*string, bool)`

GetIssuerGroupNameOk returns a tuple with the IssuerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerGroupName

`func (o *InstitutionalQuarterTotalsRow) SetIssuerGroupName(v string)`

SetIssuerGroupName sets IssuerGroupName field to given value.

### HasIssuerGroupName

`func (o *InstitutionalQuarterTotalsRow) HasIssuerGroupName() bool`

HasIssuerGroupName returns a boolean if a field has been set.

### SetIssuerGroupNameNil

`func (o *InstitutionalQuarterTotalsRow) SetIssuerGroupNameNil(b bool)`

 SetIssuerGroupNameNil sets the value for IssuerGroupName to be an explicit nil

### UnsetIssuerGroupName
`func (o *InstitutionalQuarterTotalsRow) UnsetIssuerGroupName()`

UnsetIssuerGroupName ensures that no value is present for IssuerGroupName, not even an explicit nil
### GetReportPeriod

`func (o *InstitutionalQuarterTotalsRow) GetReportPeriod() string`

GetReportPeriod returns the ReportPeriod field if non-nil, zero value otherwise.

### GetReportPeriodOk

`func (o *InstitutionalQuarterTotalsRow) GetReportPeriodOk() (*string, bool)`

GetReportPeriodOk returns a tuple with the ReportPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportPeriod

`func (o *InstitutionalQuarterTotalsRow) SetReportPeriod(v string)`

SetReportPeriod sets ReportPeriod field to given value.

### HasReportPeriod

`func (o *InstitutionalQuarterTotalsRow) HasReportPeriod() bool`

HasReportPeriod returns a boolean if a field has been set.

### SetReportPeriodNil

`func (o *InstitutionalQuarterTotalsRow) SetReportPeriodNil(b bool)`

 SetReportPeriodNil sets the value for ReportPeriod to be an explicit nil

### UnsetReportPeriod
`func (o *InstitutionalQuarterTotalsRow) UnsetReportPeriod()`

UnsetReportPeriod ensures that no value is present for ReportPeriod, not even an explicit nil
### GetIsReferencePeriod

`func (o *InstitutionalQuarterTotalsRow) GetIsReferencePeriod() bool`

GetIsReferencePeriod returns the IsReferencePeriod field if non-nil, zero value otherwise.

### GetIsReferencePeriodOk

`func (o *InstitutionalQuarterTotalsRow) GetIsReferencePeriodOk() (*bool, bool)`

GetIsReferencePeriodOk returns a tuple with the IsReferencePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReferencePeriod

`func (o *InstitutionalQuarterTotalsRow) SetIsReferencePeriod(v bool)`

SetIsReferencePeriod sets IsReferencePeriod field to given value.

### HasIsReferencePeriod

`func (o *InstitutionalQuarterTotalsRow) HasIsReferencePeriod() bool`

HasIsReferencePeriod returns a boolean if a field has been set.

### SetIsReferencePeriodNil

`func (o *InstitutionalQuarterTotalsRow) SetIsReferencePeriodNil(b bool)`

 SetIsReferencePeriodNil sets the value for IsReferencePeriod to be an explicit nil

### UnsetIsReferencePeriod
`func (o *InstitutionalQuarterTotalsRow) UnsetIsReferencePeriod()`

UnsetIsReferencePeriod ensures that no value is present for IsReferencePeriod, not even an explicit nil
### GetManagerCount

`func (o *InstitutionalQuarterTotalsRow) GetManagerCount() int64`

GetManagerCount returns the ManagerCount field if non-nil, zero value otherwise.

### GetManagerCountOk

`func (o *InstitutionalQuarterTotalsRow) GetManagerCountOk() (*int64, bool)`

GetManagerCountOk returns a tuple with the ManagerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerCount

`func (o *InstitutionalQuarterTotalsRow) SetManagerCount(v int64)`

SetManagerCount sets ManagerCount field to given value.

### HasManagerCount

`func (o *InstitutionalQuarterTotalsRow) HasManagerCount() bool`

HasManagerCount returns a boolean if a field has been set.

### SetManagerCountNil

`func (o *InstitutionalQuarterTotalsRow) SetManagerCountNil(b bool)`

 SetManagerCountNil sets the value for ManagerCount to be an explicit nil

### UnsetManagerCount
`func (o *InstitutionalQuarterTotalsRow) UnsetManagerCount()`

UnsetManagerCount ensures that no value is present for ManagerCount, not even an explicit nil
### GetPositionCount

`func (o *InstitutionalQuarterTotalsRow) GetPositionCount() int64`

GetPositionCount returns the PositionCount field if non-nil, zero value otherwise.

### GetPositionCountOk

`func (o *InstitutionalQuarterTotalsRow) GetPositionCountOk() (*int64, bool)`

GetPositionCountOk returns a tuple with the PositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionCount

`func (o *InstitutionalQuarterTotalsRow) SetPositionCount(v int64)`

SetPositionCount sets PositionCount field to given value.

### HasPositionCount

`func (o *InstitutionalQuarterTotalsRow) HasPositionCount() bool`

HasPositionCount returns a boolean if a field has been set.

### SetPositionCountNil

`func (o *InstitutionalQuarterTotalsRow) SetPositionCountNil(b bool)`

 SetPositionCountNil sets the value for PositionCount to be an explicit nil

### UnsetPositionCount
`func (o *InstitutionalQuarterTotalsRow) UnsetPositionCount()`

UnsetPositionCount ensures that no value is present for PositionCount, not even an explicit nil
### GetLanePositionCount

`func (o *InstitutionalQuarterTotalsRow) GetLanePositionCount() int64`

GetLanePositionCount returns the LanePositionCount field if non-nil, zero value otherwise.

### GetLanePositionCountOk

`func (o *InstitutionalQuarterTotalsRow) GetLanePositionCountOk() (*int64, bool)`

GetLanePositionCountOk returns a tuple with the LanePositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanePositionCount

`func (o *InstitutionalQuarterTotalsRow) SetLanePositionCount(v int64)`

SetLanePositionCount sets LanePositionCount field to given value.

### HasLanePositionCount

`func (o *InstitutionalQuarterTotalsRow) HasLanePositionCount() bool`

HasLanePositionCount returns a boolean if a field has been set.

### SetLanePositionCountNil

`func (o *InstitutionalQuarterTotalsRow) SetLanePositionCountNil(b bool)`

 SetLanePositionCountNil sets the value for LanePositionCount to be an explicit nil

### UnsetLanePositionCount
`func (o *InstitutionalQuarterTotalsRow) UnsetLanePositionCount()`

UnsetLanePositionCount ensures that no value is present for LanePositionCount, not even an explicit nil
### GetSharesReported

`func (o *InstitutionalQuarterTotalsRow) GetSharesReported() float32`

GetSharesReported returns the SharesReported field if non-nil, zero value otherwise.

### GetSharesReportedOk

`func (o *InstitutionalQuarterTotalsRow) GetSharesReportedOk() (*float32, bool)`

GetSharesReportedOk returns a tuple with the SharesReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesReported

`func (o *InstitutionalQuarterTotalsRow) SetSharesReported(v float32)`

SetSharesReported sets SharesReported field to given value.

### HasSharesReported

`func (o *InstitutionalQuarterTotalsRow) HasSharesReported() bool`

HasSharesReported returns a boolean if a field has been set.

### SetSharesReportedNil

`func (o *InstitutionalQuarterTotalsRow) SetSharesReportedNil(b bool)`

 SetSharesReportedNil sets the value for SharesReported to be an explicit nil

### UnsetSharesReported
`func (o *InstitutionalQuarterTotalsRow) UnsetSharesReported()`

UnsetSharesReported ensures that no value is present for SharesReported, not even an explicit nil
### GetSharesOtr

`func (o *InstitutionalQuarterTotalsRow) GetSharesOtr() float32`

GetSharesOtr returns the SharesOtr field if non-nil, zero value otherwise.

### GetSharesOtrOk

`func (o *InstitutionalQuarterTotalsRow) GetSharesOtrOk() (*float32, bool)`

GetSharesOtrOk returns a tuple with the SharesOtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesOtr

`func (o *InstitutionalQuarterTotalsRow) SetSharesOtr(v float32)`

SetSharesOtr sets SharesOtr field to given value.

### HasSharesOtr

`func (o *InstitutionalQuarterTotalsRow) HasSharesOtr() bool`

HasSharesOtr returns a boolean if a field has been set.

### SetSharesOtrNil

`func (o *InstitutionalQuarterTotalsRow) SetSharesOtrNil(b bool)`

 SetSharesOtrNil sets the value for SharesOtr to be an explicit nil

### UnsetSharesOtr
`func (o *InstitutionalQuarterTotalsRow) UnsetSharesOtr()`

UnsetSharesOtr ensures that no value is present for SharesOtr, not even an explicit nil
### GetSharesExcludedNonEquityIssue

`func (o *InstitutionalQuarterTotalsRow) GetSharesExcludedNonEquityIssue() float32`

GetSharesExcludedNonEquityIssue returns the SharesExcludedNonEquityIssue field if non-nil, zero value otherwise.

### GetSharesExcludedNonEquityIssueOk

`func (o *InstitutionalQuarterTotalsRow) GetSharesExcludedNonEquityIssueOk() (*float32, bool)`

GetSharesExcludedNonEquityIssueOk returns a tuple with the SharesExcludedNonEquityIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesExcludedNonEquityIssue

`func (o *InstitutionalQuarterTotalsRow) SetSharesExcludedNonEquityIssue(v float32)`

SetSharesExcludedNonEquityIssue sets SharesExcludedNonEquityIssue field to given value.

### HasSharesExcludedNonEquityIssue

`func (o *InstitutionalQuarterTotalsRow) HasSharesExcludedNonEquityIssue() bool`

HasSharesExcludedNonEquityIssue returns a boolean if a field has been set.

### SetSharesExcludedNonEquityIssueNil

`func (o *InstitutionalQuarterTotalsRow) SetSharesExcludedNonEquityIssueNil(b bool)`

 SetSharesExcludedNonEquityIssueNil sets the value for SharesExcludedNonEquityIssue to be an explicit nil

### UnsetSharesExcludedNonEquityIssue
`func (o *InstitutionalQuarterTotalsRow) UnsetSharesExcludedNonEquityIssue()`

UnsetSharesExcludedNonEquityIssue ensures that no value is present for SharesExcludedNonEquityIssue, not even an explicit nil
### GetExcludedNonEquityIssuePositionCount

`func (o *InstitutionalQuarterTotalsRow) GetExcludedNonEquityIssuePositionCount() int64`

GetExcludedNonEquityIssuePositionCount returns the ExcludedNonEquityIssuePositionCount field if non-nil, zero value otherwise.

### GetExcludedNonEquityIssuePositionCountOk

`func (o *InstitutionalQuarterTotalsRow) GetExcludedNonEquityIssuePositionCountOk() (*int64, bool)`

GetExcludedNonEquityIssuePositionCountOk returns a tuple with the ExcludedNonEquityIssuePositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedNonEquityIssuePositionCount

`func (o *InstitutionalQuarterTotalsRow) SetExcludedNonEquityIssuePositionCount(v int64)`

SetExcludedNonEquityIssuePositionCount sets ExcludedNonEquityIssuePositionCount field to given value.

### HasExcludedNonEquityIssuePositionCount

`func (o *InstitutionalQuarterTotalsRow) HasExcludedNonEquityIssuePositionCount() bool`

HasExcludedNonEquityIssuePositionCount returns a boolean if a field has been set.

### SetExcludedNonEquityIssuePositionCountNil

`func (o *InstitutionalQuarterTotalsRow) SetExcludedNonEquityIssuePositionCountNil(b bool)`

 SetExcludedNonEquityIssuePositionCountNil sets the value for ExcludedNonEquityIssuePositionCount to be an explicit nil

### UnsetExcludedNonEquityIssuePositionCount
`func (o *InstitutionalQuarterTotalsRow) UnsetExcludedNonEquityIssuePositionCount()`

UnsetExcludedNonEquityIssuePositionCount ensures that no value is present for ExcludedNonEquityIssuePositionCount, not even an explicit nil
### GetExcludedDerivativePositionCount

`func (o *InstitutionalQuarterTotalsRow) GetExcludedDerivativePositionCount() int64`

GetExcludedDerivativePositionCount returns the ExcludedDerivativePositionCount field if non-nil, zero value otherwise.

### GetExcludedDerivativePositionCountOk

`func (o *InstitutionalQuarterTotalsRow) GetExcludedDerivativePositionCountOk() (*int64, bool)`

GetExcludedDerivativePositionCountOk returns a tuple with the ExcludedDerivativePositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedDerivativePositionCount

`func (o *InstitutionalQuarterTotalsRow) SetExcludedDerivativePositionCount(v int64)`

SetExcludedDerivativePositionCount sets ExcludedDerivativePositionCount field to given value.

### HasExcludedDerivativePositionCount

`func (o *InstitutionalQuarterTotalsRow) HasExcludedDerivativePositionCount() bool`

HasExcludedDerivativePositionCount returns a boolean if a field has been set.

### SetExcludedDerivativePositionCountNil

`func (o *InstitutionalQuarterTotalsRow) SetExcludedDerivativePositionCountNil(b bool)`

 SetExcludedDerivativePositionCountNil sets the value for ExcludedDerivativePositionCount to be an explicit nil

### UnsetExcludedDerivativePositionCount
`func (o *InstitutionalQuarterTotalsRow) UnsetExcludedDerivativePositionCount()`

UnsetExcludedDerivativePositionCount ensures that no value is present for ExcludedDerivativePositionCount, not even an explicit nil
### GetExcludedPrincipalPositionCount

`func (o *InstitutionalQuarterTotalsRow) GetExcludedPrincipalPositionCount() int64`

GetExcludedPrincipalPositionCount returns the ExcludedPrincipalPositionCount field if non-nil, zero value otherwise.

### GetExcludedPrincipalPositionCountOk

`func (o *InstitutionalQuarterTotalsRow) GetExcludedPrincipalPositionCountOk() (*int64, bool)`

GetExcludedPrincipalPositionCountOk returns a tuple with the ExcludedPrincipalPositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedPrincipalPositionCount

`func (o *InstitutionalQuarterTotalsRow) SetExcludedPrincipalPositionCount(v int64)`

SetExcludedPrincipalPositionCount sets ExcludedPrincipalPositionCount field to given value.

### HasExcludedPrincipalPositionCount

`func (o *InstitutionalQuarterTotalsRow) HasExcludedPrincipalPositionCount() bool`

HasExcludedPrincipalPositionCount returns a boolean if a field has been set.

### SetExcludedPrincipalPositionCountNil

`func (o *InstitutionalQuarterTotalsRow) SetExcludedPrincipalPositionCountNil(b bool)`

 SetExcludedPrincipalPositionCountNil sets the value for ExcludedPrincipalPositionCount to be an explicit nil

### UnsetExcludedPrincipalPositionCount
`func (o *InstitutionalQuarterTotalsRow) UnsetExcludedPrincipalPositionCount()`

UnsetExcludedPrincipalPositionCount ensures that no value is present for ExcludedPrincipalPositionCount, not even an explicit nil
### GetExcludedNonEquityCusipCount

`func (o *InstitutionalQuarterTotalsRow) GetExcludedNonEquityCusipCount() int64`

GetExcludedNonEquityCusipCount returns the ExcludedNonEquityCusipCount field if non-nil, zero value otherwise.

### GetExcludedNonEquityCusipCountOk

`func (o *InstitutionalQuarterTotalsRow) GetExcludedNonEquityCusipCountOk() (*int64, bool)`

GetExcludedNonEquityCusipCountOk returns a tuple with the ExcludedNonEquityCusipCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedNonEquityCusipCount

`func (o *InstitutionalQuarterTotalsRow) SetExcludedNonEquityCusipCount(v int64)`

SetExcludedNonEquityCusipCount sets ExcludedNonEquityCusipCount field to given value.

### HasExcludedNonEquityCusipCount

`func (o *InstitutionalQuarterTotalsRow) HasExcludedNonEquityCusipCount() bool`

HasExcludedNonEquityCusipCount returns a boolean if a field has been set.

### SetExcludedNonEquityCusipCountNil

`func (o *InstitutionalQuarterTotalsRow) SetExcludedNonEquityCusipCountNil(b bool)`

 SetExcludedNonEquityCusipCountNil sets the value for ExcludedNonEquityCusipCount to be an explicit nil

### UnsetExcludedNonEquityCusipCount
`func (o *InstitutionalQuarterTotalsRow) UnsetExcludedNonEquityCusipCount()`

UnsetExcludedNonEquityCusipCount ensures that no value is present for ExcludedNonEquityCusipCount, not even an explicit nil
### GetEquityCusipCount

`func (o *InstitutionalQuarterTotalsRow) GetEquityCusipCount() int64`

GetEquityCusipCount returns the EquityCusipCount field if non-nil, zero value otherwise.

### GetEquityCusipCountOk

`func (o *InstitutionalQuarterTotalsRow) GetEquityCusipCountOk() (*int64, bool)`

GetEquityCusipCountOk returns a tuple with the EquityCusipCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityCusipCount

`func (o *InstitutionalQuarterTotalsRow) SetEquityCusipCount(v int64)`

SetEquityCusipCount sets EquityCusipCount field to given value.

### HasEquityCusipCount

`func (o *InstitutionalQuarterTotalsRow) HasEquityCusipCount() bool`

HasEquityCusipCount returns a boolean if a field has been set.

### SetEquityCusipCountNil

`func (o *InstitutionalQuarterTotalsRow) SetEquityCusipCountNil(b bool)`

 SetEquityCusipCountNil sets the value for EquityCusipCount to be an explicit nil

### UnsetEquityCusipCount
`func (o *InstitutionalQuarterTotalsRow) UnsetEquityCusipCount()`

UnsetEquityCusipCount ensures that no value is present for EquityCusipCount, not even an explicit nil
### GetLatestFilingDate

`func (o *InstitutionalQuarterTotalsRow) GetLatestFilingDate() string`

GetLatestFilingDate returns the LatestFilingDate field if non-nil, zero value otherwise.

### GetLatestFilingDateOk

`func (o *InstitutionalQuarterTotalsRow) GetLatestFilingDateOk() (*string, bool)`

GetLatestFilingDateOk returns a tuple with the LatestFilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestFilingDate

`func (o *InstitutionalQuarterTotalsRow) SetLatestFilingDate(v string)`

SetLatestFilingDate sets LatestFilingDate field to given value.

### HasLatestFilingDate

`func (o *InstitutionalQuarterTotalsRow) HasLatestFilingDate() bool`

HasLatestFilingDate returns a boolean if a field has been set.

### SetLatestFilingDateNil

`func (o *InstitutionalQuarterTotalsRow) SetLatestFilingDateNil(b bool)`

 SetLatestFilingDateNil sets the value for LatestFilingDate to be an explicit nil

### UnsetLatestFilingDate
`func (o *InstitutionalQuarterTotalsRow) UnsetLatestFilingDate()`

UnsetLatestFilingDate ensures that no value is present for LatestFilingDate, not even an explicit nil
### GetAggregationRule

`func (o *InstitutionalQuarterTotalsRow) GetAggregationRule() string`

GetAggregationRule returns the AggregationRule field if non-nil, zero value otherwise.

### GetAggregationRuleOk

`func (o *InstitutionalQuarterTotalsRow) GetAggregationRuleOk() (*string, bool)`

GetAggregationRuleOk returns a tuple with the AggregationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationRule

`func (o *InstitutionalQuarterTotalsRow) SetAggregationRule(v string)`

SetAggregationRule sets AggregationRule field to given value.

### HasAggregationRule

`func (o *InstitutionalQuarterTotalsRow) HasAggregationRule() bool`

HasAggregationRule returns a boolean if a field has been set.

### SetAggregationRuleNil

`func (o *InstitutionalQuarterTotalsRow) SetAggregationRuleNil(b bool)`

 SetAggregationRuleNil sets the value for AggregationRule to be an explicit nil

### UnsetAggregationRule
`func (o *InstitutionalQuarterTotalsRow) UnsetAggregationRule()`

UnsetAggregationRule ensures that no value is present for AggregationRule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


