# CompanyFundingSummaryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | The company key (company:&lt;key&gt;). | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**NRounds** | Pointer to **NullableInt64** | Funding events of the company, disclosed or not. | [optional] 
**TotalNonDilutiveUsd** | Pointer to **NullableFloat32** | Disclosed non-dilutive amounts (grants, debt, royalty and similar), in whole US dollars. | [optional] 
**TotalDilutiveUsd** | Pointer to **NullableFloat32** | Disclosed dilutive amounts (equity-issuing rounds), in whole US dollars. | [optional] 
**TotalUnclassifiedUsd** | Pointer to **NullableFloat32** | Disclosed amounts of events whose instrument is not classified, kept apart from the other two totals. | [optional] 
**NUnclassifiedRounds** | Pointer to **NullableInt64** | Events whose instrument is not classified. | [optional] 
**TotalRaisedUsd** | Pointer to **NullableFloat32** | All disclosed amounts, the sum of the three totals, in whole US dollars. | [optional] 
**FirstFundedDate** | Pointer to **NullableString** | Date of the earliest event. | [optional] 
**LastFundedDate** | Pointer to **NullableString** | Date of the latest event. | [optional] 
**LatestRoundType** | Pointer to **NullableString** | The round type of the latest event. | [optional] 
**LatestRoundDate** | Pointer to **NullableString** | The date of the latest event. | [optional] 
**StageLabel** | Pointer to **NullableString** | A coarse reading of the round types: public, growth_private, seed_private, unclassified or non_dilutive_only. | [optional] 

## Methods

### NewCompanyFundingSummaryRow

`func NewCompanyFundingSummaryRow() *CompanyFundingSummaryRow`

NewCompanyFundingSummaryRow instantiates a new CompanyFundingSummaryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyFundingSummaryRowWithDefaults

`func NewCompanyFundingSummaryRowWithDefaults() *CompanyFundingSummaryRow`

NewCompanyFundingSummaryRowWithDefaults instantiates a new CompanyFundingSummaryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompanyFundingSummaryRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyFundingSummaryRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyFundingSummaryRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyFundingSummaryRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyFundingSummaryRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyFundingSummaryRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyFundingSummaryRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyFundingSummaryRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyFundingSummaryRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyFundingSummaryRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyFundingSummaryRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyFundingSummaryRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *CompanyFundingSummaryRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyFundingSummaryRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyFundingSummaryRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyFundingSummaryRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyFundingSummaryRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyFundingSummaryRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyFundingSummaryRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyFundingSummaryRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyFundingSummaryRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyFundingSummaryRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyFundingSummaryRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyFundingSummaryRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetNRounds

`func (o *CompanyFundingSummaryRow) GetNRounds() int64`

GetNRounds returns the NRounds field if non-nil, zero value otherwise.

### GetNRoundsOk

`func (o *CompanyFundingSummaryRow) GetNRoundsOk() (*int64, bool)`

GetNRoundsOk returns a tuple with the NRounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRounds

`func (o *CompanyFundingSummaryRow) SetNRounds(v int64)`

SetNRounds sets NRounds field to given value.

### HasNRounds

`func (o *CompanyFundingSummaryRow) HasNRounds() bool`

HasNRounds returns a boolean if a field has been set.

### SetNRoundsNil

`func (o *CompanyFundingSummaryRow) SetNRoundsNil(b bool)`

 SetNRoundsNil sets the value for NRounds to be an explicit nil

### UnsetNRounds
`func (o *CompanyFundingSummaryRow) UnsetNRounds()`

UnsetNRounds ensures that no value is present for NRounds, not even an explicit nil
### GetTotalNonDilutiveUsd

`func (o *CompanyFundingSummaryRow) GetTotalNonDilutiveUsd() float32`

GetTotalNonDilutiveUsd returns the TotalNonDilutiveUsd field if non-nil, zero value otherwise.

### GetTotalNonDilutiveUsdOk

`func (o *CompanyFundingSummaryRow) GetTotalNonDilutiveUsdOk() (*float32, bool)`

GetTotalNonDilutiveUsdOk returns a tuple with the TotalNonDilutiveUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNonDilutiveUsd

`func (o *CompanyFundingSummaryRow) SetTotalNonDilutiveUsd(v float32)`

SetTotalNonDilutiveUsd sets TotalNonDilutiveUsd field to given value.

### HasTotalNonDilutiveUsd

`func (o *CompanyFundingSummaryRow) HasTotalNonDilutiveUsd() bool`

HasTotalNonDilutiveUsd returns a boolean if a field has been set.

### SetTotalNonDilutiveUsdNil

`func (o *CompanyFundingSummaryRow) SetTotalNonDilutiveUsdNil(b bool)`

 SetTotalNonDilutiveUsdNil sets the value for TotalNonDilutiveUsd to be an explicit nil

### UnsetTotalNonDilutiveUsd
`func (o *CompanyFundingSummaryRow) UnsetTotalNonDilutiveUsd()`

UnsetTotalNonDilutiveUsd ensures that no value is present for TotalNonDilutiveUsd, not even an explicit nil
### GetTotalDilutiveUsd

`func (o *CompanyFundingSummaryRow) GetTotalDilutiveUsd() float32`

GetTotalDilutiveUsd returns the TotalDilutiveUsd field if non-nil, zero value otherwise.

### GetTotalDilutiveUsdOk

`func (o *CompanyFundingSummaryRow) GetTotalDilutiveUsdOk() (*float32, bool)`

GetTotalDilutiveUsdOk returns a tuple with the TotalDilutiveUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDilutiveUsd

`func (o *CompanyFundingSummaryRow) SetTotalDilutiveUsd(v float32)`

SetTotalDilutiveUsd sets TotalDilutiveUsd field to given value.

### HasTotalDilutiveUsd

`func (o *CompanyFundingSummaryRow) HasTotalDilutiveUsd() bool`

HasTotalDilutiveUsd returns a boolean if a field has been set.

### SetTotalDilutiveUsdNil

`func (o *CompanyFundingSummaryRow) SetTotalDilutiveUsdNil(b bool)`

 SetTotalDilutiveUsdNil sets the value for TotalDilutiveUsd to be an explicit nil

### UnsetTotalDilutiveUsd
`func (o *CompanyFundingSummaryRow) UnsetTotalDilutiveUsd()`

UnsetTotalDilutiveUsd ensures that no value is present for TotalDilutiveUsd, not even an explicit nil
### GetTotalUnclassifiedUsd

`func (o *CompanyFundingSummaryRow) GetTotalUnclassifiedUsd() float32`

GetTotalUnclassifiedUsd returns the TotalUnclassifiedUsd field if non-nil, zero value otherwise.

### GetTotalUnclassifiedUsdOk

`func (o *CompanyFundingSummaryRow) GetTotalUnclassifiedUsdOk() (*float32, bool)`

GetTotalUnclassifiedUsdOk returns a tuple with the TotalUnclassifiedUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnclassifiedUsd

`func (o *CompanyFundingSummaryRow) SetTotalUnclassifiedUsd(v float32)`

SetTotalUnclassifiedUsd sets TotalUnclassifiedUsd field to given value.

### HasTotalUnclassifiedUsd

`func (o *CompanyFundingSummaryRow) HasTotalUnclassifiedUsd() bool`

HasTotalUnclassifiedUsd returns a boolean if a field has been set.

### SetTotalUnclassifiedUsdNil

`func (o *CompanyFundingSummaryRow) SetTotalUnclassifiedUsdNil(b bool)`

 SetTotalUnclassifiedUsdNil sets the value for TotalUnclassifiedUsd to be an explicit nil

### UnsetTotalUnclassifiedUsd
`func (o *CompanyFundingSummaryRow) UnsetTotalUnclassifiedUsd()`

UnsetTotalUnclassifiedUsd ensures that no value is present for TotalUnclassifiedUsd, not even an explicit nil
### GetNUnclassifiedRounds

`func (o *CompanyFundingSummaryRow) GetNUnclassifiedRounds() int64`

GetNUnclassifiedRounds returns the NUnclassifiedRounds field if non-nil, zero value otherwise.

### GetNUnclassifiedRoundsOk

`func (o *CompanyFundingSummaryRow) GetNUnclassifiedRoundsOk() (*int64, bool)`

GetNUnclassifiedRoundsOk returns a tuple with the NUnclassifiedRounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNUnclassifiedRounds

`func (o *CompanyFundingSummaryRow) SetNUnclassifiedRounds(v int64)`

SetNUnclassifiedRounds sets NUnclassifiedRounds field to given value.

### HasNUnclassifiedRounds

`func (o *CompanyFundingSummaryRow) HasNUnclassifiedRounds() bool`

HasNUnclassifiedRounds returns a boolean if a field has been set.

### SetNUnclassifiedRoundsNil

`func (o *CompanyFundingSummaryRow) SetNUnclassifiedRoundsNil(b bool)`

 SetNUnclassifiedRoundsNil sets the value for NUnclassifiedRounds to be an explicit nil

### UnsetNUnclassifiedRounds
`func (o *CompanyFundingSummaryRow) UnsetNUnclassifiedRounds()`

UnsetNUnclassifiedRounds ensures that no value is present for NUnclassifiedRounds, not even an explicit nil
### GetTotalRaisedUsd

`func (o *CompanyFundingSummaryRow) GetTotalRaisedUsd() float32`

GetTotalRaisedUsd returns the TotalRaisedUsd field if non-nil, zero value otherwise.

### GetTotalRaisedUsdOk

`func (o *CompanyFundingSummaryRow) GetTotalRaisedUsdOk() (*float32, bool)`

GetTotalRaisedUsdOk returns a tuple with the TotalRaisedUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRaisedUsd

`func (o *CompanyFundingSummaryRow) SetTotalRaisedUsd(v float32)`

SetTotalRaisedUsd sets TotalRaisedUsd field to given value.

### HasTotalRaisedUsd

`func (o *CompanyFundingSummaryRow) HasTotalRaisedUsd() bool`

HasTotalRaisedUsd returns a boolean if a field has been set.

### SetTotalRaisedUsdNil

`func (o *CompanyFundingSummaryRow) SetTotalRaisedUsdNil(b bool)`

 SetTotalRaisedUsdNil sets the value for TotalRaisedUsd to be an explicit nil

### UnsetTotalRaisedUsd
`func (o *CompanyFundingSummaryRow) UnsetTotalRaisedUsd()`

UnsetTotalRaisedUsd ensures that no value is present for TotalRaisedUsd, not even an explicit nil
### GetFirstFundedDate

`func (o *CompanyFundingSummaryRow) GetFirstFundedDate() string`

GetFirstFundedDate returns the FirstFundedDate field if non-nil, zero value otherwise.

### GetFirstFundedDateOk

`func (o *CompanyFundingSummaryRow) GetFirstFundedDateOk() (*string, bool)`

GetFirstFundedDateOk returns a tuple with the FirstFundedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstFundedDate

`func (o *CompanyFundingSummaryRow) SetFirstFundedDate(v string)`

SetFirstFundedDate sets FirstFundedDate field to given value.

### HasFirstFundedDate

`func (o *CompanyFundingSummaryRow) HasFirstFundedDate() bool`

HasFirstFundedDate returns a boolean if a field has been set.

### SetFirstFundedDateNil

`func (o *CompanyFundingSummaryRow) SetFirstFundedDateNil(b bool)`

 SetFirstFundedDateNil sets the value for FirstFundedDate to be an explicit nil

### UnsetFirstFundedDate
`func (o *CompanyFundingSummaryRow) UnsetFirstFundedDate()`

UnsetFirstFundedDate ensures that no value is present for FirstFundedDate, not even an explicit nil
### GetLastFundedDate

`func (o *CompanyFundingSummaryRow) GetLastFundedDate() string`

GetLastFundedDate returns the LastFundedDate field if non-nil, zero value otherwise.

### GetLastFundedDateOk

`func (o *CompanyFundingSummaryRow) GetLastFundedDateOk() (*string, bool)`

GetLastFundedDateOk returns a tuple with the LastFundedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFundedDate

`func (o *CompanyFundingSummaryRow) SetLastFundedDate(v string)`

SetLastFundedDate sets LastFundedDate field to given value.

### HasLastFundedDate

`func (o *CompanyFundingSummaryRow) HasLastFundedDate() bool`

HasLastFundedDate returns a boolean if a field has been set.

### SetLastFundedDateNil

`func (o *CompanyFundingSummaryRow) SetLastFundedDateNil(b bool)`

 SetLastFundedDateNil sets the value for LastFundedDate to be an explicit nil

### UnsetLastFundedDate
`func (o *CompanyFundingSummaryRow) UnsetLastFundedDate()`

UnsetLastFundedDate ensures that no value is present for LastFundedDate, not even an explicit nil
### GetLatestRoundType

`func (o *CompanyFundingSummaryRow) GetLatestRoundType() string`

GetLatestRoundType returns the LatestRoundType field if non-nil, zero value otherwise.

### GetLatestRoundTypeOk

`func (o *CompanyFundingSummaryRow) GetLatestRoundTypeOk() (*string, bool)`

GetLatestRoundTypeOk returns a tuple with the LatestRoundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRoundType

`func (o *CompanyFundingSummaryRow) SetLatestRoundType(v string)`

SetLatestRoundType sets LatestRoundType field to given value.

### HasLatestRoundType

`func (o *CompanyFundingSummaryRow) HasLatestRoundType() bool`

HasLatestRoundType returns a boolean if a field has been set.

### SetLatestRoundTypeNil

`func (o *CompanyFundingSummaryRow) SetLatestRoundTypeNil(b bool)`

 SetLatestRoundTypeNil sets the value for LatestRoundType to be an explicit nil

### UnsetLatestRoundType
`func (o *CompanyFundingSummaryRow) UnsetLatestRoundType()`

UnsetLatestRoundType ensures that no value is present for LatestRoundType, not even an explicit nil
### GetLatestRoundDate

`func (o *CompanyFundingSummaryRow) GetLatestRoundDate() string`

GetLatestRoundDate returns the LatestRoundDate field if non-nil, zero value otherwise.

### GetLatestRoundDateOk

`func (o *CompanyFundingSummaryRow) GetLatestRoundDateOk() (*string, bool)`

GetLatestRoundDateOk returns a tuple with the LatestRoundDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRoundDate

`func (o *CompanyFundingSummaryRow) SetLatestRoundDate(v string)`

SetLatestRoundDate sets LatestRoundDate field to given value.

### HasLatestRoundDate

`func (o *CompanyFundingSummaryRow) HasLatestRoundDate() bool`

HasLatestRoundDate returns a boolean if a field has been set.

### SetLatestRoundDateNil

`func (o *CompanyFundingSummaryRow) SetLatestRoundDateNil(b bool)`

 SetLatestRoundDateNil sets the value for LatestRoundDate to be an explicit nil

### UnsetLatestRoundDate
`func (o *CompanyFundingSummaryRow) UnsetLatestRoundDate()`

UnsetLatestRoundDate ensures that no value is present for LatestRoundDate, not even an explicit nil
### GetStageLabel

`func (o *CompanyFundingSummaryRow) GetStageLabel() string`

GetStageLabel returns the StageLabel field if non-nil, zero value otherwise.

### GetStageLabelOk

`func (o *CompanyFundingSummaryRow) GetStageLabelOk() (*string, bool)`

GetStageLabelOk returns a tuple with the StageLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageLabel

`func (o *CompanyFundingSummaryRow) SetStageLabel(v string)`

SetStageLabel sets StageLabel field to given value.

### HasStageLabel

`func (o *CompanyFundingSummaryRow) HasStageLabel() bool`

HasStageLabel returns a boolean if a field has been set.

### SetStageLabelNil

`func (o *CompanyFundingSummaryRow) SetStageLabelNil(b bool)`

 SetStageLabelNil sets the value for StageLabel to be an explicit nil

### UnsetStageLabel
`func (o *CompanyFundingSummaryRow) UnsetStageLabel()`

UnsetStageLabel ensures that no value is present for StageLabel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


