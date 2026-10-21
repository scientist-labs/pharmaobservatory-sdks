# LegalEntitiesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **NullableString** | The organization record id: a registered legal entity (edgar:, ror:) or a name-resolved sponsor (sponsor:). | [optional] 
**OrgName** | Pointer to **NullableString** | The organization&#39;s name. | [optional] 
**IsNameResolvedSponsor** | Pointer to **NullableBool** | True for a sponsor resolved from a filed name alone rather than from a registration. | [optional] 
**OrgOrigin** | Pointer to **NullableString** | How the organization record was minted. | [optional] 
**Country** | Pointer to **NullableString** | Country as the source states it. | [optional] 
**OrgType** | Pointer to **NullableString** | Organization type: Academic, Government, Hospital, Industry, Nonprofit or Other (from the Research Organization Registry (ROR) record or the filing), or the lowercase tags investor and manufacturer that mark a minted adviser or manufacturer record. | [optional] 
**EdgarCik** | Pointer to **NullableString** | Central index key (CIK) in EDGAR, the US Securities and Exchange Commission (SEC) filing system, when it files. | [optional] 
**EdgarTicker** | Pointer to **NullableString** | Exchange ticker, when listed. | [optional] 
**RorId** | Pointer to **NullableString** | Research Organization Registry id, when it has one. | [optional] 
**Lei** | Pointer to **NullableString** | Legal Entity Identifier, when it has one. | [optional] 
**LifecycleStatus** | Pointer to **NullableString** | Lifecycle status from the registry sources. | [optional] 
**NumLeadTrials** | Pointer to **NullableInt64** | Trials the organization leads. | [optional] 
**NumTrials** | Pointer to **NullableInt64** | Trials the organization leads or collaborates on. | [optional] 
**TotalLeadEnrollment** | Pointer to **NullableFloat32** | Planned enrollment across led trials. | [optional] 
**NumRecentLeadTrials** | Pointer to **NullableInt64** | Led trials started in the recent window. | [optional] 
**ActivityScore** | Pointer to **NullableFloat32** | Trial activity score. | [optional] 
**CompanyId** | Pointer to **NullableString** | Company level: the company the organization belongs to (null when unresolved). | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 

## Methods

### NewLegalEntitiesRow

`func NewLegalEntitiesRow() *LegalEntitiesRow`

NewLegalEntitiesRow instantiates a new LegalEntitiesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalEntitiesRowWithDefaults

`func NewLegalEntitiesRowWithDefaults() *LegalEntitiesRow`

NewLegalEntitiesRowWithDefaults instantiates a new LegalEntitiesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *LegalEntitiesRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *LegalEntitiesRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *LegalEntitiesRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *LegalEntitiesRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *LegalEntitiesRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *LegalEntitiesRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetOrgName

`func (o *LegalEntitiesRow) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *LegalEntitiesRow) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *LegalEntitiesRow) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *LegalEntitiesRow) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### SetOrgNameNil

`func (o *LegalEntitiesRow) SetOrgNameNil(b bool)`

 SetOrgNameNil sets the value for OrgName to be an explicit nil

### UnsetOrgName
`func (o *LegalEntitiesRow) UnsetOrgName()`

UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil
### GetIsNameResolvedSponsor

`func (o *LegalEntitiesRow) GetIsNameResolvedSponsor() bool`

GetIsNameResolvedSponsor returns the IsNameResolvedSponsor field if non-nil, zero value otherwise.

### GetIsNameResolvedSponsorOk

`func (o *LegalEntitiesRow) GetIsNameResolvedSponsorOk() (*bool, bool)`

GetIsNameResolvedSponsorOk returns a tuple with the IsNameResolvedSponsor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNameResolvedSponsor

`func (o *LegalEntitiesRow) SetIsNameResolvedSponsor(v bool)`

SetIsNameResolvedSponsor sets IsNameResolvedSponsor field to given value.

### HasIsNameResolvedSponsor

`func (o *LegalEntitiesRow) HasIsNameResolvedSponsor() bool`

HasIsNameResolvedSponsor returns a boolean if a field has been set.

### SetIsNameResolvedSponsorNil

`func (o *LegalEntitiesRow) SetIsNameResolvedSponsorNil(b bool)`

 SetIsNameResolvedSponsorNil sets the value for IsNameResolvedSponsor to be an explicit nil

### UnsetIsNameResolvedSponsor
`func (o *LegalEntitiesRow) UnsetIsNameResolvedSponsor()`

UnsetIsNameResolvedSponsor ensures that no value is present for IsNameResolvedSponsor, not even an explicit nil
### GetOrgOrigin

`func (o *LegalEntitiesRow) GetOrgOrigin() string`

GetOrgOrigin returns the OrgOrigin field if non-nil, zero value otherwise.

### GetOrgOriginOk

`func (o *LegalEntitiesRow) GetOrgOriginOk() (*string, bool)`

GetOrgOriginOk returns a tuple with the OrgOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgOrigin

`func (o *LegalEntitiesRow) SetOrgOrigin(v string)`

SetOrgOrigin sets OrgOrigin field to given value.

### HasOrgOrigin

`func (o *LegalEntitiesRow) HasOrgOrigin() bool`

HasOrgOrigin returns a boolean if a field has been set.

### SetOrgOriginNil

`func (o *LegalEntitiesRow) SetOrgOriginNil(b bool)`

 SetOrgOriginNil sets the value for OrgOrigin to be an explicit nil

### UnsetOrgOrigin
`func (o *LegalEntitiesRow) UnsetOrgOrigin()`

UnsetOrgOrigin ensures that no value is present for OrgOrigin, not even an explicit nil
### GetCountry

`func (o *LegalEntitiesRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LegalEntitiesRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LegalEntitiesRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *LegalEntitiesRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *LegalEntitiesRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *LegalEntitiesRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetOrgType

`func (o *LegalEntitiesRow) GetOrgType() string`

GetOrgType returns the OrgType field if non-nil, zero value otherwise.

### GetOrgTypeOk

`func (o *LegalEntitiesRow) GetOrgTypeOk() (*string, bool)`

GetOrgTypeOk returns a tuple with the OrgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgType

`func (o *LegalEntitiesRow) SetOrgType(v string)`

SetOrgType sets OrgType field to given value.

### HasOrgType

`func (o *LegalEntitiesRow) HasOrgType() bool`

HasOrgType returns a boolean if a field has been set.

### SetOrgTypeNil

`func (o *LegalEntitiesRow) SetOrgTypeNil(b bool)`

 SetOrgTypeNil sets the value for OrgType to be an explicit nil

### UnsetOrgType
`func (o *LegalEntitiesRow) UnsetOrgType()`

UnsetOrgType ensures that no value is present for OrgType, not even an explicit nil
### GetEdgarCik

`func (o *LegalEntitiesRow) GetEdgarCik() string`

GetEdgarCik returns the EdgarCik field if non-nil, zero value otherwise.

### GetEdgarCikOk

`func (o *LegalEntitiesRow) GetEdgarCikOk() (*string, bool)`

GetEdgarCikOk returns a tuple with the EdgarCik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgarCik

`func (o *LegalEntitiesRow) SetEdgarCik(v string)`

SetEdgarCik sets EdgarCik field to given value.

### HasEdgarCik

`func (o *LegalEntitiesRow) HasEdgarCik() bool`

HasEdgarCik returns a boolean if a field has been set.

### SetEdgarCikNil

`func (o *LegalEntitiesRow) SetEdgarCikNil(b bool)`

 SetEdgarCikNil sets the value for EdgarCik to be an explicit nil

### UnsetEdgarCik
`func (o *LegalEntitiesRow) UnsetEdgarCik()`

UnsetEdgarCik ensures that no value is present for EdgarCik, not even an explicit nil
### GetEdgarTicker

`func (o *LegalEntitiesRow) GetEdgarTicker() string`

GetEdgarTicker returns the EdgarTicker field if non-nil, zero value otherwise.

### GetEdgarTickerOk

`func (o *LegalEntitiesRow) GetEdgarTickerOk() (*string, bool)`

GetEdgarTickerOk returns a tuple with the EdgarTicker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgarTicker

`func (o *LegalEntitiesRow) SetEdgarTicker(v string)`

SetEdgarTicker sets EdgarTicker field to given value.

### HasEdgarTicker

`func (o *LegalEntitiesRow) HasEdgarTicker() bool`

HasEdgarTicker returns a boolean if a field has been set.

### SetEdgarTickerNil

`func (o *LegalEntitiesRow) SetEdgarTickerNil(b bool)`

 SetEdgarTickerNil sets the value for EdgarTicker to be an explicit nil

### UnsetEdgarTicker
`func (o *LegalEntitiesRow) UnsetEdgarTicker()`

UnsetEdgarTicker ensures that no value is present for EdgarTicker, not even an explicit nil
### GetRorId

`func (o *LegalEntitiesRow) GetRorId() string`

GetRorId returns the RorId field if non-nil, zero value otherwise.

### GetRorIdOk

`func (o *LegalEntitiesRow) GetRorIdOk() (*string, bool)`

GetRorIdOk returns a tuple with the RorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRorId

`func (o *LegalEntitiesRow) SetRorId(v string)`

SetRorId sets RorId field to given value.

### HasRorId

`func (o *LegalEntitiesRow) HasRorId() bool`

HasRorId returns a boolean if a field has been set.

### SetRorIdNil

`func (o *LegalEntitiesRow) SetRorIdNil(b bool)`

 SetRorIdNil sets the value for RorId to be an explicit nil

### UnsetRorId
`func (o *LegalEntitiesRow) UnsetRorId()`

UnsetRorId ensures that no value is present for RorId, not even an explicit nil
### GetLei

`func (o *LegalEntitiesRow) GetLei() string`

GetLei returns the Lei field if non-nil, zero value otherwise.

### GetLeiOk

`func (o *LegalEntitiesRow) GetLeiOk() (*string, bool)`

GetLeiOk returns a tuple with the Lei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLei

`func (o *LegalEntitiesRow) SetLei(v string)`

SetLei sets Lei field to given value.

### HasLei

`func (o *LegalEntitiesRow) HasLei() bool`

HasLei returns a boolean if a field has been set.

### SetLeiNil

`func (o *LegalEntitiesRow) SetLeiNil(b bool)`

 SetLeiNil sets the value for Lei to be an explicit nil

### UnsetLei
`func (o *LegalEntitiesRow) UnsetLei()`

UnsetLei ensures that no value is present for Lei, not even an explicit nil
### GetLifecycleStatus

`func (o *LegalEntitiesRow) GetLifecycleStatus() string`

GetLifecycleStatus returns the LifecycleStatus field if non-nil, zero value otherwise.

### GetLifecycleStatusOk

`func (o *LegalEntitiesRow) GetLifecycleStatusOk() (*string, bool)`

GetLifecycleStatusOk returns a tuple with the LifecycleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStatus

`func (o *LegalEntitiesRow) SetLifecycleStatus(v string)`

SetLifecycleStatus sets LifecycleStatus field to given value.

### HasLifecycleStatus

`func (o *LegalEntitiesRow) HasLifecycleStatus() bool`

HasLifecycleStatus returns a boolean if a field has been set.

### SetLifecycleStatusNil

`func (o *LegalEntitiesRow) SetLifecycleStatusNil(b bool)`

 SetLifecycleStatusNil sets the value for LifecycleStatus to be an explicit nil

### UnsetLifecycleStatus
`func (o *LegalEntitiesRow) UnsetLifecycleStatus()`

UnsetLifecycleStatus ensures that no value is present for LifecycleStatus, not even an explicit nil
### GetNumLeadTrials

`func (o *LegalEntitiesRow) GetNumLeadTrials() int64`

GetNumLeadTrials returns the NumLeadTrials field if non-nil, zero value otherwise.

### GetNumLeadTrialsOk

`func (o *LegalEntitiesRow) GetNumLeadTrialsOk() (*int64, bool)`

GetNumLeadTrialsOk returns a tuple with the NumLeadTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLeadTrials

`func (o *LegalEntitiesRow) SetNumLeadTrials(v int64)`

SetNumLeadTrials sets NumLeadTrials field to given value.

### HasNumLeadTrials

`func (o *LegalEntitiesRow) HasNumLeadTrials() bool`

HasNumLeadTrials returns a boolean if a field has been set.

### SetNumLeadTrialsNil

`func (o *LegalEntitiesRow) SetNumLeadTrialsNil(b bool)`

 SetNumLeadTrialsNil sets the value for NumLeadTrials to be an explicit nil

### UnsetNumLeadTrials
`func (o *LegalEntitiesRow) UnsetNumLeadTrials()`

UnsetNumLeadTrials ensures that no value is present for NumLeadTrials, not even an explicit nil
### GetNumTrials

`func (o *LegalEntitiesRow) GetNumTrials() int64`

GetNumTrials returns the NumTrials field if non-nil, zero value otherwise.

### GetNumTrialsOk

`func (o *LegalEntitiesRow) GetNumTrialsOk() (*int64, bool)`

GetNumTrialsOk returns a tuple with the NumTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTrials

`func (o *LegalEntitiesRow) SetNumTrials(v int64)`

SetNumTrials sets NumTrials field to given value.

### HasNumTrials

`func (o *LegalEntitiesRow) HasNumTrials() bool`

HasNumTrials returns a boolean if a field has been set.

### SetNumTrialsNil

`func (o *LegalEntitiesRow) SetNumTrialsNil(b bool)`

 SetNumTrialsNil sets the value for NumTrials to be an explicit nil

### UnsetNumTrials
`func (o *LegalEntitiesRow) UnsetNumTrials()`

UnsetNumTrials ensures that no value is present for NumTrials, not even an explicit nil
### GetTotalLeadEnrollment

`func (o *LegalEntitiesRow) GetTotalLeadEnrollment() float32`

GetTotalLeadEnrollment returns the TotalLeadEnrollment field if non-nil, zero value otherwise.

### GetTotalLeadEnrollmentOk

`func (o *LegalEntitiesRow) GetTotalLeadEnrollmentOk() (*float32, bool)`

GetTotalLeadEnrollmentOk returns a tuple with the TotalLeadEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLeadEnrollment

`func (o *LegalEntitiesRow) SetTotalLeadEnrollment(v float32)`

SetTotalLeadEnrollment sets TotalLeadEnrollment field to given value.

### HasTotalLeadEnrollment

`func (o *LegalEntitiesRow) HasTotalLeadEnrollment() bool`

HasTotalLeadEnrollment returns a boolean if a field has been set.

### SetTotalLeadEnrollmentNil

`func (o *LegalEntitiesRow) SetTotalLeadEnrollmentNil(b bool)`

 SetTotalLeadEnrollmentNil sets the value for TotalLeadEnrollment to be an explicit nil

### UnsetTotalLeadEnrollment
`func (o *LegalEntitiesRow) UnsetTotalLeadEnrollment()`

UnsetTotalLeadEnrollment ensures that no value is present for TotalLeadEnrollment, not even an explicit nil
### GetNumRecentLeadTrials

`func (o *LegalEntitiesRow) GetNumRecentLeadTrials() int64`

GetNumRecentLeadTrials returns the NumRecentLeadTrials field if non-nil, zero value otherwise.

### GetNumRecentLeadTrialsOk

`func (o *LegalEntitiesRow) GetNumRecentLeadTrialsOk() (*int64, bool)`

GetNumRecentLeadTrialsOk returns a tuple with the NumRecentLeadTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRecentLeadTrials

`func (o *LegalEntitiesRow) SetNumRecentLeadTrials(v int64)`

SetNumRecentLeadTrials sets NumRecentLeadTrials field to given value.

### HasNumRecentLeadTrials

`func (o *LegalEntitiesRow) HasNumRecentLeadTrials() bool`

HasNumRecentLeadTrials returns a boolean if a field has been set.

### SetNumRecentLeadTrialsNil

`func (o *LegalEntitiesRow) SetNumRecentLeadTrialsNil(b bool)`

 SetNumRecentLeadTrialsNil sets the value for NumRecentLeadTrials to be an explicit nil

### UnsetNumRecentLeadTrials
`func (o *LegalEntitiesRow) UnsetNumRecentLeadTrials()`

UnsetNumRecentLeadTrials ensures that no value is present for NumRecentLeadTrials, not even an explicit nil
### GetActivityScore

`func (o *LegalEntitiesRow) GetActivityScore() float32`

GetActivityScore returns the ActivityScore field if non-nil, zero value otherwise.

### GetActivityScoreOk

`func (o *LegalEntitiesRow) GetActivityScoreOk() (*float32, bool)`

GetActivityScoreOk returns a tuple with the ActivityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityScore

`func (o *LegalEntitiesRow) SetActivityScore(v float32)`

SetActivityScore sets ActivityScore field to given value.

### HasActivityScore

`func (o *LegalEntitiesRow) HasActivityScore() bool`

HasActivityScore returns a boolean if a field has been set.

### SetActivityScoreNil

`func (o *LegalEntitiesRow) SetActivityScoreNil(b bool)`

 SetActivityScoreNil sets the value for ActivityScore to be an explicit nil

### UnsetActivityScore
`func (o *LegalEntitiesRow) UnsetActivityScore()`

UnsetActivityScore ensures that no value is present for ActivityScore, not even an explicit nil
### GetCompanyId

`func (o *LegalEntitiesRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *LegalEntitiesRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *LegalEntitiesRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *LegalEntitiesRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *LegalEntitiesRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *LegalEntitiesRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *LegalEntitiesRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *LegalEntitiesRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *LegalEntitiesRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *LegalEntitiesRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *LegalEntitiesRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *LegalEntitiesRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *LegalEntitiesRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *LegalEntitiesRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *LegalEntitiesRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *LegalEntitiesRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *LegalEntitiesRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *LegalEntitiesRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *LegalEntitiesRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *LegalEntitiesRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *LegalEntitiesRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *LegalEntitiesRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *LegalEntitiesRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *LegalEntitiesRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


