# OrgDisclosureTrendRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **NullableString** | Legal-entity level: the organization record of the SEC registrant that filed. | [optional] 
**OrgName** | Pointer to **NullableString** | Display name of the filing organization record. | [optional] 
**CompanyId** | Pointer to **NullableString** | Company level: the company the filing organization belongs to (null when unresolved). | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the filing organization&#39;s company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the filing organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the filing organization&#39;s group. | [optional] 
**ReportingYear** | Pointer to **NullableInt32** | The fiscal year the filing reports. | [optional] 
**ReportingCompany** | Pointer to **NullableString** | The company name as the filing states it. | [optional] 
**Cik** | Pointer to **NullableString** | The registrant&#39;s SEC Central Index Key (CIK). | [optional] 
**Ticker** | Pointer to **NullableString** | The stock ticker the filing states. | [optional] 
**BoardPctWomen** | Pointer to **NullableFloat32** | Women&#39;s share of the board, as the filing states it. | [optional] 
**WorkforcePctWomen** | Pointer to **NullableFloat32** | Women&#39;s share of the workforce. | [optional] 
**LeadershipPctWomen** | Pointer to **NullableFloat32** | Women&#39;s share of leadership roles. | [optional] 
**WorkforcePctUnderrepresented** | Pointer to **NullableFloat32** | Share of the workforce from underrepresented groups. | [optional] 
**GenderPayGapPct** | Pointer to **NullableFloat32** | The stated gender pay gap. | [optional] 
**HasDeiTargets** | Pointer to **NullableBool** | Whether the filing states diversity, equity and inclusion targets. | [optional] 
**HasNetZeroTarget** | Pointer to **NullableBool** | Whether the filing states a net-zero target. | [optional] 
**NetZeroTargetYear** | Pointer to **NullableInt32** | The year of the net-zero target. | [optional] 
**ReductionPct** | Pointer to **NullableFloat32** | The emissions reduction the filing commits to. | [optional] 
**SbtiStatusClaimed** | Pointer to **NullableString** | The Science Based Targets initiative (SBTi) status the filing claims for itself (a claim in the filing, not the SBTi list). | [optional] 
**TargetTemperatureClaimed** | Pointer to **NullableString** | The warming limit the filing says its targets align with. | [optional] 
**ClimateDocumentId** | Pointer to **NullableString** | The filing document the climate measures were extracted from. | [optional] 
**DiversityDocumentId** | Pointer to **NullableString** | The filing document the diversity measures were extracted from. | [optional] 

## Methods

### NewOrgDisclosureTrendRow

`func NewOrgDisclosureTrendRow() *OrgDisclosureTrendRow`

NewOrgDisclosureTrendRow instantiates a new OrgDisclosureTrendRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgDisclosureTrendRowWithDefaults

`func NewOrgDisclosureTrendRowWithDefaults() *OrgDisclosureTrendRow`

NewOrgDisclosureTrendRowWithDefaults instantiates a new OrgDisclosureTrendRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *OrgDisclosureTrendRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrgDisclosureTrendRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrgDisclosureTrendRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *OrgDisclosureTrendRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *OrgDisclosureTrendRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *OrgDisclosureTrendRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetOrgName

`func (o *OrgDisclosureTrendRow) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *OrgDisclosureTrendRow) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *OrgDisclosureTrendRow) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *OrgDisclosureTrendRow) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### SetOrgNameNil

`func (o *OrgDisclosureTrendRow) SetOrgNameNil(b bool)`

 SetOrgNameNil sets the value for OrgName to be an explicit nil

### UnsetOrgName
`func (o *OrgDisclosureTrendRow) UnsetOrgName()`

UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil
### GetCompanyId

`func (o *OrgDisclosureTrendRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *OrgDisclosureTrendRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *OrgDisclosureTrendRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *OrgDisclosureTrendRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *OrgDisclosureTrendRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *OrgDisclosureTrendRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *OrgDisclosureTrendRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrgDisclosureTrendRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrgDisclosureTrendRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrgDisclosureTrendRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *OrgDisclosureTrendRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *OrgDisclosureTrendRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *OrgDisclosureTrendRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *OrgDisclosureTrendRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *OrgDisclosureTrendRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *OrgDisclosureTrendRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *OrgDisclosureTrendRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *OrgDisclosureTrendRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *OrgDisclosureTrendRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *OrgDisclosureTrendRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *OrgDisclosureTrendRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *OrgDisclosureTrendRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *OrgDisclosureTrendRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *OrgDisclosureTrendRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetReportingYear

`func (o *OrgDisclosureTrendRow) GetReportingYear() int32`

GetReportingYear returns the ReportingYear field if non-nil, zero value otherwise.

### GetReportingYearOk

`func (o *OrgDisclosureTrendRow) GetReportingYearOk() (*int32, bool)`

GetReportingYearOk returns a tuple with the ReportingYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingYear

`func (o *OrgDisclosureTrendRow) SetReportingYear(v int32)`

SetReportingYear sets ReportingYear field to given value.

### HasReportingYear

`func (o *OrgDisclosureTrendRow) HasReportingYear() bool`

HasReportingYear returns a boolean if a field has been set.

### SetReportingYearNil

`func (o *OrgDisclosureTrendRow) SetReportingYearNil(b bool)`

 SetReportingYearNil sets the value for ReportingYear to be an explicit nil

### UnsetReportingYear
`func (o *OrgDisclosureTrendRow) UnsetReportingYear()`

UnsetReportingYear ensures that no value is present for ReportingYear, not even an explicit nil
### GetReportingCompany

`func (o *OrgDisclosureTrendRow) GetReportingCompany() string`

GetReportingCompany returns the ReportingCompany field if non-nil, zero value otherwise.

### GetReportingCompanyOk

`func (o *OrgDisclosureTrendRow) GetReportingCompanyOk() (*string, bool)`

GetReportingCompanyOk returns a tuple with the ReportingCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingCompany

`func (o *OrgDisclosureTrendRow) SetReportingCompany(v string)`

SetReportingCompany sets ReportingCompany field to given value.

### HasReportingCompany

`func (o *OrgDisclosureTrendRow) HasReportingCompany() bool`

HasReportingCompany returns a boolean if a field has been set.

### SetReportingCompanyNil

`func (o *OrgDisclosureTrendRow) SetReportingCompanyNil(b bool)`

 SetReportingCompanyNil sets the value for ReportingCompany to be an explicit nil

### UnsetReportingCompany
`func (o *OrgDisclosureTrendRow) UnsetReportingCompany()`

UnsetReportingCompany ensures that no value is present for ReportingCompany, not even an explicit nil
### GetCik

`func (o *OrgDisclosureTrendRow) GetCik() string`

GetCik returns the Cik field if non-nil, zero value otherwise.

### GetCikOk

`func (o *OrgDisclosureTrendRow) GetCikOk() (*string, bool)`

GetCikOk returns a tuple with the Cik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCik

`func (o *OrgDisclosureTrendRow) SetCik(v string)`

SetCik sets Cik field to given value.

### HasCik

`func (o *OrgDisclosureTrendRow) HasCik() bool`

HasCik returns a boolean if a field has been set.

### SetCikNil

`func (o *OrgDisclosureTrendRow) SetCikNil(b bool)`

 SetCikNil sets the value for Cik to be an explicit nil

### UnsetCik
`func (o *OrgDisclosureTrendRow) UnsetCik()`

UnsetCik ensures that no value is present for Cik, not even an explicit nil
### GetTicker

`func (o *OrgDisclosureTrendRow) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *OrgDisclosureTrendRow) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *OrgDisclosureTrendRow) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *OrgDisclosureTrendRow) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### SetTickerNil

`func (o *OrgDisclosureTrendRow) SetTickerNil(b bool)`

 SetTickerNil sets the value for Ticker to be an explicit nil

### UnsetTicker
`func (o *OrgDisclosureTrendRow) UnsetTicker()`

UnsetTicker ensures that no value is present for Ticker, not even an explicit nil
### GetBoardPctWomen

`func (o *OrgDisclosureTrendRow) GetBoardPctWomen() float32`

GetBoardPctWomen returns the BoardPctWomen field if non-nil, zero value otherwise.

### GetBoardPctWomenOk

`func (o *OrgDisclosureTrendRow) GetBoardPctWomenOk() (*float32, bool)`

GetBoardPctWomenOk returns a tuple with the BoardPctWomen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardPctWomen

`func (o *OrgDisclosureTrendRow) SetBoardPctWomen(v float32)`

SetBoardPctWomen sets BoardPctWomen field to given value.

### HasBoardPctWomen

`func (o *OrgDisclosureTrendRow) HasBoardPctWomen() bool`

HasBoardPctWomen returns a boolean if a field has been set.

### SetBoardPctWomenNil

`func (o *OrgDisclosureTrendRow) SetBoardPctWomenNil(b bool)`

 SetBoardPctWomenNil sets the value for BoardPctWomen to be an explicit nil

### UnsetBoardPctWomen
`func (o *OrgDisclosureTrendRow) UnsetBoardPctWomen()`

UnsetBoardPctWomen ensures that no value is present for BoardPctWomen, not even an explicit nil
### GetWorkforcePctWomen

`func (o *OrgDisclosureTrendRow) GetWorkforcePctWomen() float32`

GetWorkforcePctWomen returns the WorkforcePctWomen field if non-nil, zero value otherwise.

### GetWorkforcePctWomenOk

`func (o *OrgDisclosureTrendRow) GetWorkforcePctWomenOk() (*float32, bool)`

GetWorkforcePctWomenOk returns a tuple with the WorkforcePctWomen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkforcePctWomen

`func (o *OrgDisclosureTrendRow) SetWorkforcePctWomen(v float32)`

SetWorkforcePctWomen sets WorkforcePctWomen field to given value.

### HasWorkforcePctWomen

`func (o *OrgDisclosureTrendRow) HasWorkforcePctWomen() bool`

HasWorkforcePctWomen returns a boolean if a field has been set.

### SetWorkforcePctWomenNil

`func (o *OrgDisclosureTrendRow) SetWorkforcePctWomenNil(b bool)`

 SetWorkforcePctWomenNil sets the value for WorkforcePctWomen to be an explicit nil

### UnsetWorkforcePctWomen
`func (o *OrgDisclosureTrendRow) UnsetWorkforcePctWomen()`

UnsetWorkforcePctWomen ensures that no value is present for WorkforcePctWomen, not even an explicit nil
### GetLeadershipPctWomen

`func (o *OrgDisclosureTrendRow) GetLeadershipPctWomen() float32`

GetLeadershipPctWomen returns the LeadershipPctWomen field if non-nil, zero value otherwise.

### GetLeadershipPctWomenOk

`func (o *OrgDisclosureTrendRow) GetLeadershipPctWomenOk() (*float32, bool)`

GetLeadershipPctWomenOk returns a tuple with the LeadershipPctWomen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadershipPctWomen

`func (o *OrgDisclosureTrendRow) SetLeadershipPctWomen(v float32)`

SetLeadershipPctWomen sets LeadershipPctWomen field to given value.

### HasLeadershipPctWomen

`func (o *OrgDisclosureTrendRow) HasLeadershipPctWomen() bool`

HasLeadershipPctWomen returns a boolean if a field has been set.

### SetLeadershipPctWomenNil

`func (o *OrgDisclosureTrendRow) SetLeadershipPctWomenNil(b bool)`

 SetLeadershipPctWomenNil sets the value for LeadershipPctWomen to be an explicit nil

### UnsetLeadershipPctWomen
`func (o *OrgDisclosureTrendRow) UnsetLeadershipPctWomen()`

UnsetLeadershipPctWomen ensures that no value is present for LeadershipPctWomen, not even an explicit nil
### GetWorkforcePctUnderrepresented

`func (o *OrgDisclosureTrendRow) GetWorkforcePctUnderrepresented() float32`

GetWorkforcePctUnderrepresented returns the WorkforcePctUnderrepresented field if non-nil, zero value otherwise.

### GetWorkforcePctUnderrepresentedOk

`func (o *OrgDisclosureTrendRow) GetWorkforcePctUnderrepresentedOk() (*float32, bool)`

GetWorkforcePctUnderrepresentedOk returns a tuple with the WorkforcePctUnderrepresented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkforcePctUnderrepresented

`func (o *OrgDisclosureTrendRow) SetWorkforcePctUnderrepresented(v float32)`

SetWorkforcePctUnderrepresented sets WorkforcePctUnderrepresented field to given value.

### HasWorkforcePctUnderrepresented

`func (o *OrgDisclosureTrendRow) HasWorkforcePctUnderrepresented() bool`

HasWorkforcePctUnderrepresented returns a boolean if a field has been set.

### SetWorkforcePctUnderrepresentedNil

`func (o *OrgDisclosureTrendRow) SetWorkforcePctUnderrepresentedNil(b bool)`

 SetWorkforcePctUnderrepresentedNil sets the value for WorkforcePctUnderrepresented to be an explicit nil

### UnsetWorkforcePctUnderrepresented
`func (o *OrgDisclosureTrendRow) UnsetWorkforcePctUnderrepresented()`

UnsetWorkforcePctUnderrepresented ensures that no value is present for WorkforcePctUnderrepresented, not even an explicit nil
### GetGenderPayGapPct

`func (o *OrgDisclosureTrendRow) GetGenderPayGapPct() float32`

GetGenderPayGapPct returns the GenderPayGapPct field if non-nil, zero value otherwise.

### GetGenderPayGapPctOk

`func (o *OrgDisclosureTrendRow) GetGenderPayGapPctOk() (*float32, bool)`

GetGenderPayGapPctOk returns a tuple with the GenderPayGapPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenderPayGapPct

`func (o *OrgDisclosureTrendRow) SetGenderPayGapPct(v float32)`

SetGenderPayGapPct sets GenderPayGapPct field to given value.

### HasGenderPayGapPct

`func (o *OrgDisclosureTrendRow) HasGenderPayGapPct() bool`

HasGenderPayGapPct returns a boolean if a field has been set.

### SetGenderPayGapPctNil

`func (o *OrgDisclosureTrendRow) SetGenderPayGapPctNil(b bool)`

 SetGenderPayGapPctNil sets the value for GenderPayGapPct to be an explicit nil

### UnsetGenderPayGapPct
`func (o *OrgDisclosureTrendRow) UnsetGenderPayGapPct()`

UnsetGenderPayGapPct ensures that no value is present for GenderPayGapPct, not even an explicit nil
### GetHasDeiTargets

`func (o *OrgDisclosureTrendRow) GetHasDeiTargets() bool`

GetHasDeiTargets returns the HasDeiTargets field if non-nil, zero value otherwise.

### GetHasDeiTargetsOk

`func (o *OrgDisclosureTrendRow) GetHasDeiTargetsOk() (*bool, bool)`

GetHasDeiTargetsOk returns a tuple with the HasDeiTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDeiTargets

`func (o *OrgDisclosureTrendRow) SetHasDeiTargets(v bool)`

SetHasDeiTargets sets HasDeiTargets field to given value.

### HasHasDeiTargets

`func (o *OrgDisclosureTrendRow) HasHasDeiTargets() bool`

HasHasDeiTargets returns a boolean if a field has been set.

### SetHasDeiTargetsNil

`func (o *OrgDisclosureTrendRow) SetHasDeiTargetsNil(b bool)`

 SetHasDeiTargetsNil sets the value for HasDeiTargets to be an explicit nil

### UnsetHasDeiTargets
`func (o *OrgDisclosureTrendRow) UnsetHasDeiTargets()`

UnsetHasDeiTargets ensures that no value is present for HasDeiTargets, not even an explicit nil
### GetHasNetZeroTarget

`func (o *OrgDisclosureTrendRow) GetHasNetZeroTarget() bool`

GetHasNetZeroTarget returns the HasNetZeroTarget field if non-nil, zero value otherwise.

### GetHasNetZeroTargetOk

`func (o *OrgDisclosureTrendRow) GetHasNetZeroTargetOk() (*bool, bool)`

GetHasNetZeroTargetOk returns a tuple with the HasNetZeroTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetZeroTarget

`func (o *OrgDisclosureTrendRow) SetHasNetZeroTarget(v bool)`

SetHasNetZeroTarget sets HasNetZeroTarget field to given value.

### HasHasNetZeroTarget

`func (o *OrgDisclosureTrendRow) HasHasNetZeroTarget() bool`

HasHasNetZeroTarget returns a boolean if a field has been set.

### SetHasNetZeroTargetNil

`func (o *OrgDisclosureTrendRow) SetHasNetZeroTargetNil(b bool)`

 SetHasNetZeroTargetNil sets the value for HasNetZeroTarget to be an explicit nil

### UnsetHasNetZeroTarget
`func (o *OrgDisclosureTrendRow) UnsetHasNetZeroTarget()`

UnsetHasNetZeroTarget ensures that no value is present for HasNetZeroTarget, not even an explicit nil
### GetNetZeroTargetYear

`func (o *OrgDisclosureTrendRow) GetNetZeroTargetYear() int32`

GetNetZeroTargetYear returns the NetZeroTargetYear field if non-nil, zero value otherwise.

### GetNetZeroTargetYearOk

`func (o *OrgDisclosureTrendRow) GetNetZeroTargetYearOk() (*int32, bool)`

GetNetZeroTargetYearOk returns a tuple with the NetZeroTargetYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetZeroTargetYear

`func (o *OrgDisclosureTrendRow) SetNetZeroTargetYear(v int32)`

SetNetZeroTargetYear sets NetZeroTargetYear field to given value.

### HasNetZeroTargetYear

`func (o *OrgDisclosureTrendRow) HasNetZeroTargetYear() bool`

HasNetZeroTargetYear returns a boolean if a field has been set.

### SetNetZeroTargetYearNil

`func (o *OrgDisclosureTrendRow) SetNetZeroTargetYearNil(b bool)`

 SetNetZeroTargetYearNil sets the value for NetZeroTargetYear to be an explicit nil

### UnsetNetZeroTargetYear
`func (o *OrgDisclosureTrendRow) UnsetNetZeroTargetYear()`

UnsetNetZeroTargetYear ensures that no value is present for NetZeroTargetYear, not even an explicit nil
### GetReductionPct

`func (o *OrgDisclosureTrendRow) GetReductionPct() float32`

GetReductionPct returns the ReductionPct field if non-nil, zero value otherwise.

### GetReductionPctOk

`func (o *OrgDisclosureTrendRow) GetReductionPctOk() (*float32, bool)`

GetReductionPctOk returns a tuple with the ReductionPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReductionPct

`func (o *OrgDisclosureTrendRow) SetReductionPct(v float32)`

SetReductionPct sets ReductionPct field to given value.

### HasReductionPct

`func (o *OrgDisclosureTrendRow) HasReductionPct() bool`

HasReductionPct returns a boolean if a field has been set.

### SetReductionPctNil

`func (o *OrgDisclosureTrendRow) SetReductionPctNil(b bool)`

 SetReductionPctNil sets the value for ReductionPct to be an explicit nil

### UnsetReductionPct
`func (o *OrgDisclosureTrendRow) UnsetReductionPct()`

UnsetReductionPct ensures that no value is present for ReductionPct, not even an explicit nil
### GetSbtiStatusClaimed

`func (o *OrgDisclosureTrendRow) GetSbtiStatusClaimed() string`

GetSbtiStatusClaimed returns the SbtiStatusClaimed field if non-nil, zero value otherwise.

### GetSbtiStatusClaimedOk

`func (o *OrgDisclosureTrendRow) GetSbtiStatusClaimedOk() (*string, bool)`

GetSbtiStatusClaimedOk returns a tuple with the SbtiStatusClaimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbtiStatusClaimed

`func (o *OrgDisclosureTrendRow) SetSbtiStatusClaimed(v string)`

SetSbtiStatusClaimed sets SbtiStatusClaimed field to given value.

### HasSbtiStatusClaimed

`func (o *OrgDisclosureTrendRow) HasSbtiStatusClaimed() bool`

HasSbtiStatusClaimed returns a boolean if a field has been set.

### SetSbtiStatusClaimedNil

`func (o *OrgDisclosureTrendRow) SetSbtiStatusClaimedNil(b bool)`

 SetSbtiStatusClaimedNil sets the value for SbtiStatusClaimed to be an explicit nil

### UnsetSbtiStatusClaimed
`func (o *OrgDisclosureTrendRow) UnsetSbtiStatusClaimed()`

UnsetSbtiStatusClaimed ensures that no value is present for SbtiStatusClaimed, not even an explicit nil
### GetTargetTemperatureClaimed

`func (o *OrgDisclosureTrendRow) GetTargetTemperatureClaimed() string`

GetTargetTemperatureClaimed returns the TargetTemperatureClaimed field if non-nil, zero value otherwise.

### GetTargetTemperatureClaimedOk

`func (o *OrgDisclosureTrendRow) GetTargetTemperatureClaimedOk() (*string, bool)`

GetTargetTemperatureClaimedOk returns a tuple with the TargetTemperatureClaimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTemperatureClaimed

`func (o *OrgDisclosureTrendRow) SetTargetTemperatureClaimed(v string)`

SetTargetTemperatureClaimed sets TargetTemperatureClaimed field to given value.

### HasTargetTemperatureClaimed

`func (o *OrgDisclosureTrendRow) HasTargetTemperatureClaimed() bool`

HasTargetTemperatureClaimed returns a boolean if a field has been set.

### SetTargetTemperatureClaimedNil

`func (o *OrgDisclosureTrendRow) SetTargetTemperatureClaimedNil(b bool)`

 SetTargetTemperatureClaimedNil sets the value for TargetTemperatureClaimed to be an explicit nil

### UnsetTargetTemperatureClaimed
`func (o *OrgDisclosureTrendRow) UnsetTargetTemperatureClaimed()`

UnsetTargetTemperatureClaimed ensures that no value is present for TargetTemperatureClaimed, not even an explicit nil
### GetClimateDocumentId

`func (o *OrgDisclosureTrendRow) GetClimateDocumentId() string`

GetClimateDocumentId returns the ClimateDocumentId field if non-nil, zero value otherwise.

### GetClimateDocumentIdOk

`func (o *OrgDisclosureTrendRow) GetClimateDocumentIdOk() (*string, bool)`

GetClimateDocumentIdOk returns a tuple with the ClimateDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClimateDocumentId

`func (o *OrgDisclosureTrendRow) SetClimateDocumentId(v string)`

SetClimateDocumentId sets ClimateDocumentId field to given value.

### HasClimateDocumentId

`func (o *OrgDisclosureTrendRow) HasClimateDocumentId() bool`

HasClimateDocumentId returns a boolean if a field has been set.

### SetClimateDocumentIdNil

`func (o *OrgDisclosureTrendRow) SetClimateDocumentIdNil(b bool)`

 SetClimateDocumentIdNil sets the value for ClimateDocumentId to be an explicit nil

### UnsetClimateDocumentId
`func (o *OrgDisclosureTrendRow) UnsetClimateDocumentId()`

UnsetClimateDocumentId ensures that no value is present for ClimateDocumentId, not even an explicit nil
### GetDiversityDocumentId

`func (o *OrgDisclosureTrendRow) GetDiversityDocumentId() string`

GetDiversityDocumentId returns the DiversityDocumentId field if non-nil, zero value otherwise.

### GetDiversityDocumentIdOk

`func (o *OrgDisclosureTrendRow) GetDiversityDocumentIdOk() (*string, bool)`

GetDiversityDocumentIdOk returns a tuple with the DiversityDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiversityDocumentId

`func (o *OrgDisclosureTrendRow) SetDiversityDocumentId(v string)`

SetDiversityDocumentId sets DiversityDocumentId field to given value.

### HasDiversityDocumentId

`func (o *OrgDisclosureTrendRow) HasDiversityDocumentId() bool`

HasDiversityDocumentId returns a boolean if a field has been set.

### SetDiversityDocumentIdNil

`func (o *OrgDisclosureTrendRow) SetDiversityDocumentIdNil(b bool)`

 SetDiversityDocumentIdNil sets the value for DiversityDocumentId to be an explicit nil

### UnsetDiversityDocumentId
`func (o *OrgDisclosureTrendRow) UnsetDiversityDocumentId()`

UnsetDiversityDocumentId ensures that no value is present for DiversityDocumentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


