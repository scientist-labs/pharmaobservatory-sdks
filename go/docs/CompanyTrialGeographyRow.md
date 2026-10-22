# CompanyTrialGeographyRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | The company key (company:&lt;key&gt;). | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**IsPublicRosterRoot** | Pointer to **NullableBool** | True when the company is a public roster root: an organization the public Clinical Research Map lists. | [optional] 
**RosterRootId** | Pointer to **NullableString** | The Clinical Research Map&#39;s organization key for this company: the group when it is a public roster root, else the company when it is one, else null. | [optional] 
**RosterLeadTrials** | Pointer to **NullableInt32** | For a public roster root, the lead-trial count the public organizations list sorts on; null for a company that is not one. | [optional] 
**DomicileCountry** | Pointer to **NullableString** | ISO 3166-1 alpha-2 code of the company&#39;s legal-entity domicile; null when unknown. | [optional] 
**SiteCountry** | Pointer to **NullableString** | ISO 3166-1 alpha-2 code of the country holding the most of the company&#39;s own located led-trial sites; null when none resolves. | [optional] 
**MapCountry** | Pointer to **NullableString** | The country the map places the company under: domicile_country, else site_country; null when neither is known. | [optional] 
**MapCountryBasis** | Pointer to **NullableString** | Which signal map_country came from: domicile or trial_sites; null exactly when map_country is null. | [optional] 
**NLocatedSites** | Pointer to **NullableInt32** | The company&#39;s own located led-trial sites. | [optional] 
**NUnlocatedSites** | Pointer to **NullableInt32** | The company&#39;s led-trial sites with no coordinate, which the map does not draw. | [optional] 
**NSiteCountries** | Pointer to **NullableInt32** | Distinct resolved countries among the company&#39;s located led-trial sites. | [optional] 
**NUnresolvedSites** | Pointer to **NullableInt32** | Located led-trial sites whose country name resolves to no country code; they do not vote for site_country. | [optional] 
**RosterMapCountry** | Pointer to **NullableString** | For a public roster root, the country the public Clinical Research Map shows the organization under, as published; null for a non-root. | [optional] 
**RosterMapCountryBasis** | Pointer to **NullableString** | For a public roster root, which signal the public map&#39;s country came from: domicile or trial_sites; null for a non-root or an unknown country. | [optional] 
**RosterNLocatedSites** | Pointer to **NullableInt32** | For a public roster root, the located sites the public map draws the organization at (its group fold); null for a non-root. | [optional] 

## Methods

### NewCompanyTrialGeographyRow

`func NewCompanyTrialGeographyRow() *CompanyTrialGeographyRow`

NewCompanyTrialGeographyRow instantiates a new CompanyTrialGeographyRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyTrialGeographyRowWithDefaults

`func NewCompanyTrialGeographyRowWithDefaults() *CompanyTrialGeographyRow`

NewCompanyTrialGeographyRowWithDefaults instantiates a new CompanyTrialGeographyRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompanyTrialGeographyRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyTrialGeographyRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyTrialGeographyRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyTrialGeographyRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyTrialGeographyRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyTrialGeographyRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyTrialGeographyRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyTrialGeographyRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyTrialGeographyRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyTrialGeographyRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyTrialGeographyRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyTrialGeographyRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *CompanyTrialGeographyRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyTrialGeographyRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyTrialGeographyRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyTrialGeographyRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyTrialGeographyRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyTrialGeographyRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyTrialGeographyRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyTrialGeographyRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyTrialGeographyRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyTrialGeographyRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyTrialGeographyRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyTrialGeographyRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetIsPublicRosterRoot

`func (o *CompanyTrialGeographyRow) GetIsPublicRosterRoot() bool`

GetIsPublicRosterRoot returns the IsPublicRosterRoot field if non-nil, zero value otherwise.

### GetIsPublicRosterRootOk

`func (o *CompanyTrialGeographyRow) GetIsPublicRosterRootOk() (*bool, bool)`

GetIsPublicRosterRootOk returns a tuple with the IsPublicRosterRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublicRosterRoot

`func (o *CompanyTrialGeographyRow) SetIsPublicRosterRoot(v bool)`

SetIsPublicRosterRoot sets IsPublicRosterRoot field to given value.

### HasIsPublicRosterRoot

`func (o *CompanyTrialGeographyRow) HasIsPublicRosterRoot() bool`

HasIsPublicRosterRoot returns a boolean if a field has been set.

### SetIsPublicRosterRootNil

`func (o *CompanyTrialGeographyRow) SetIsPublicRosterRootNil(b bool)`

 SetIsPublicRosterRootNil sets the value for IsPublicRosterRoot to be an explicit nil

### UnsetIsPublicRosterRoot
`func (o *CompanyTrialGeographyRow) UnsetIsPublicRosterRoot()`

UnsetIsPublicRosterRoot ensures that no value is present for IsPublicRosterRoot, not even an explicit nil
### GetRosterRootId

`func (o *CompanyTrialGeographyRow) GetRosterRootId() string`

GetRosterRootId returns the RosterRootId field if non-nil, zero value otherwise.

### GetRosterRootIdOk

`func (o *CompanyTrialGeographyRow) GetRosterRootIdOk() (*string, bool)`

GetRosterRootIdOk returns a tuple with the RosterRootId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRosterRootId

`func (o *CompanyTrialGeographyRow) SetRosterRootId(v string)`

SetRosterRootId sets RosterRootId field to given value.

### HasRosterRootId

`func (o *CompanyTrialGeographyRow) HasRosterRootId() bool`

HasRosterRootId returns a boolean if a field has been set.

### SetRosterRootIdNil

`func (o *CompanyTrialGeographyRow) SetRosterRootIdNil(b bool)`

 SetRosterRootIdNil sets the value for RosterRootId to be an explicit nil

### UnsetRosterRootId
`func (o *CompanyTrialGeographyRow) UnsetRosterRootId()`

UnsetRosterRootId ensures that no value is present for RosterRootId, not even an explicit nil
### GetRosterLeadTrials

`func (o *CompanyTrialGeographyRow) GetRosterLeadTrials() int32`

GetRosterLeadTrials returns the RosterLeadTrials field if non-nil, zero value otherwise.

### GetRosterLeadTrialsOk

`func (o *CompanyTrialGeographyRow) GetRosterLeadTrialsOk() (*int32, bool)`

GetRosterLeadTrialsOk returns a tuple with the RosterLeadTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRosterLeadTrials

`func (o *CompanyTrialGeographyRow) SetRosterLeadTrials(v int32)`

SetRosterLeadTrials sets RosterLeadTrials field to given value.

### HasRosterLeadTrials

`func (o *CompanyTrialGeographyRow) HasRosterLeadTrials() bool`

HasRosterLeadTrials returns a boolean if a field has been set.

### SetRosterLeadTrialsNil

`func (o *CompanyTrialGeographyRow) SetRosterLeadTrialsNil(b bool)`

 SetRosterLeadTrialsNil sets the value for RosterLeadTrials to be an explicit nil

### UnsetRosterLeadTrials
`func (o *CompanyTrialGeographyRow) UnsetRosterLeadTrials()`

UnsetRosterLeadTrials ensures that no value is present for RosterLeadTrials, not even an explicit nil
### GetDomicileCountry

`func (o *CompanyTrialGeographyRow) GetDomicileCountry() string`

GetDomicileCountry returns the DomicileCountry field if non-nil, zero value otherwise.

### GetDomicileCountryOk

`func (o *CompanyTrialGeographyRow) GetDomicileCountryOk() (*string, bool)`

GetDomicileCountryOk returns a tuple with the DomicileCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomicileCountry

`func (o *CompanyTrialGeographyRow) SetDomicileCountry(v string)`

SetDomicileCountry sets DomicileCountry field to given value.

### HasDomicileCountry

`func (o *CompanyTrialGeographyRow) HasDomicileCountry() bool`

HasDomicileCountry returns a boolean if a field has been set.

### SetDomicileCountryNil

`func (o *CompanyTrialGeographyRow) SetDomicileCountryNil(b bool)`

 SetDomicileCountryNil sets the value for DomicileCountry to be an explicit nil

### UnsetDomicileCountry
`func (o *CompanyTrialGeographyRow) UnsetDomicileCountry()`

UnsetDomicileCountry ensures that no value is present for DomicileCountry, not even an explicit nil
### GetSiteCountry

`func (o *CompanyTrialGeographyRow) GetSiteCountry() string`

GetSiteCountry returns the SiteCountry field if non-nil, zero value otherwise.

### GetSiteCountryOk

`func (o *CompanyTrialGeographyRow) GetSiteCountryOk() (*string, bool)`

GetSiteCountryOk returns a tuple with the SiteCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCountry

`func (o *CompanyTrialGeographyRow) SetSiteCountry(v string)`

SetSiteCountry sets SiteCountry field to given value.

### HasSiteCountry

`func (o *CompanyTrialGeographyRow) HasSiteCountry() bool`

HasSiteCountry returns a boolean if a field has been set.

### SetSiteCountryNil

`func (o *CompanyTrialGeographyRow) SetSiteCountryNil(b bool)`

 SetSiteCountryNil sets the value for SiteCountry to be an explicit nil

### UnsetSiteCountry
`func (o *CompanyTrialGeographyRow) UnsetSiteCountry()`

UnsetSiteCountry ensures that no value is present for SiteCountry, not even an explicit nil
### GetMapCountry

`func (o *CompanyTrialGeographyRow) GetMapCountry() string`

GetMapCountry returns the MapCountry field if non-nil, zero value otherwise.

### GetMapCountryOk

`func (o *CompanyTrialGeographyRow) GetMapCountryOk() (*string, bool)`

GetMapCountryOk returns a tuple with the MapCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapCountry

`func (o *CompanyTrialGeographyRow) SetMapCountry(v string)`

SetMapCountry sets MapCountry field to given value.

### HasMapCountry

`func (o *CompanyTrialGeographyRow) HasMapCountry() bool`

HasMapCountry returns a boolean if a field has been set.

### SetMapCountryNil

`func (o *CompanyTrialGeographyRow) SetMapCountryNil(b bool)`

 SetMapCountryNil sets the value for MapCountry to be an explicit nil

### UnsetMapCountry
`func (o *CompanyTrialGeographyRow) UnsetMapCountry()`

UnsetMapCountry ensures that no value is present for MapCountry, not even an explicit nil
### GetMapCountryBasis

`func (o *CompanyTrialGeographyRow) GetMapCountryBasis() string`

GetMapCountryBasis returns the MapCountryBasis field if non-nil, zero value otherwise.

### GetMapCountryBasisOk

`func (o *CompanyTrialGeographyRow) GetMapCountryBasisOk() (*string, bool)`

GetMapCountryBasisOk returns a tuple with the MapCountryBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapCountryBasis

`func (o *CompanyTrialGeographyRow) SetMapCountryBasis(v string)`

SetMapCountryBasis sets MapCountryBasis field to given value.

### HasMapCountryBasis

`func (o *CompanyTrialGeographyRow) HasMapCountryBasis() bool`

HasMapCountryBasis returns a boolean if a field has been set.

### SetMapCountryBasisNil

`func (o *CompanyTrialGeographyRow) SetMapCountryBasisNil(b bool)`

 SetMapCountryBasisNil sets the value for MapCountryBasis to be an explicit nil

### UnsetMapCountryBasis
`func (o *CompanyTrialGeographyRow) UnsetMapCountryBasis()`

UnsetMapCountryBasis ensures that no value is present for MapCountryBasis, not even an explicit nil
### GetNLocatedSites

`func (o *CompanyTrialGeographyRow) GetNLocatedSites() int32`

GetNLocatedSites returns the NLocatedSites field if non-nil, zero value otherwise.

### GetNLocatedSitesOk

`func (o *CompanyTrialGeographyRow) GetNLocatedSitesOk() (*int32, bool)`

GetNLocatedSitesOk returns a tuple with the NLocatedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNLocatedSites

`func (o *CompanyTrialGeographyRow) SetNLocatedSites(v int32)`

SetNLocatedSites sets NLocatedSites field to given value.

### HasNLocatedSites

`func (o *CompanyTrialGeographyRow) HasNLocatedSites() bool`

HasNLocatedSites returns a boolean if a field has been set.

### SetNLocatedSitesNil

`func (o *CompanyTrialGeographyRow) SetNLocatedSitesNil(b bool)`

 SetNLocatedSitesNil sets the value for NLocatedSites to be an explicit nil

### UnsetNLocatedSites
`func (o *CompanyTrialGeographyRow) UnsetNLocatedSites()`

UnsetNLocatedSites ensures that no value is present for NLocatedSites, not even an explicit nil
### GetNUnlocatedSites

`func (o *CompanyTrialGeographyRow) GetNUnlocatedSites() int32`

GetNUnlocatedSites returns the NUnlocatedSites field if non-nil, zero value otherwise.

### GetNUnlocatedSitesOk

`func (o *CompanyTrialGeographyRow) GetNUnlocatedSitesOk() (*int32, bool)`

GetNUnlocatedSitesOk returns a tuple with the NUnlocatedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNUnlocatedSites

`func (o *CompanyTrialGeographyRow) SetNUnlocatedSites(v int32)`

SetNUnlocatedSites sets NUnlocatedSites field to given value.

### HasNUnlocatedSites

`func (o *CompanyTrialGeographyRow) HasNUnlocatedSites() bool`

HasNUnlocatedSites returns a boolean if a field has been set.

### SetNUnlocatedSitesNil

`func (o *CompanyTrialGeographyRow) SetNUnlocatedSitesNil(b bool)`

 SetNUnlocatedSitesNil sets the value for NUnlocatedSites to be an explicit nil

### UnsetNUnlocatedSites
`func (o *CompanyTrialGeographyRow) UnsetNUnlocatedSites()`

UnsetNUnlocatedSites ensures that no value is present for NUnlocatedSites, not even an explicit nil
### GetNSiteCountries

`func (o *CompanyTrialGeographyRow) GetNSiteCountries() int32`

GetNSiteCountries returns the NSiteCountries field if non-nil, zero value otherwise.

### GetNSiteCountriesOk

`func (o *CompanyTrialGeographyRow) GetNSiteCountriesOk() (*int32, bool)`

GetNSiteCountriesOk returns a tuple with the NSiteCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSiteCountries

`func (o *CompanyTrialGeographyRow) SetNSiteCountries(v int32)`

SetNSiteCountries sets NSiteCountries field to given value.

### HasNSiteCountries

`func (o *CompanyTrialGeographyRow) HasNSiteCountries() bool`

HasNSiteCountries returns a boolean if a field has been set.

### SetNSiteCountriesNil

`func (o *CompanyTrialGeographyRow) SetNSiteCountriesNil(b bool)`

 SetNSiteCountriesNil sets the value for NSiteCountries to be an explicit nil

### UnsetNSiteCountries
`func (o *CompanyTrialGeographyRow) UnsetNSiteCountries()`

UnsetNSiteCountries ensures that no value is present for NSiteCountries, not even an explicit nil
### GetNUnresolvedSites

`func (o *CompanyTrialGeographyRow) GetNUnresolvedSites() int32`

GetNUnresolvedSites returns the NUnresolvedSites field if non-nil, zero value otherwise.

### GetNUnresolvedSitesOk

`func (o *CompanyTrialGeographyRow) GetNUnresolvedSitesOk() (*int32, bool)`

GetNUnresolvedSitesOk returns a tuple with the NUnresolvedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNUnresolvedSites

`func (o *CompanyTrialGeographyRow) SetNUnresolvedSites(v int32)`

SetNUnresolvedSites sets NUnresolvedSites field to given value.

### HasNUnresolvedSites

`func (o *CompanyTrialGeographyRow) HasNUnresolvedSites() bool`

HasNUnresolvedSites returns a boolean if a field has been set.

### SetNUnresolvedSitesNil

`func (o *CompanyTrialGeographyRow) SetNUnresolvedSitesNil(b bool)`

 SetNUnresolvedSitesNil sets the value for NUnresolvedSites to be an explicit nil

### UnsetNUnresolvedSites
`func (o *CompanyTrialGeographyRow) UnsetNUnresolvedSites()`

UnsetNUnresolvedSites ensures that no value is present for NUnresolvedSites, not even an explicit nil
### GetRosterMapCountry

`func (o *CompanyTrialGeographyRow) GetRosterMapCountry() string`

GetRosterMapCountry returns the RosterMapCountry field if non-nil, zero value otherwise.

### GetRosterMapCountryOk

`func (o *CompanyTrialGeographyRow) GetRosterMapCountryOk() (*string, bool)`

GetRosterMapCountryOk returns a tuple with the RosterMapCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRosterMapCountry

`func (o *CompanyTrialGeographyRow) SetRosterMapCountry(v string)`

SetRosterMapCountry sets RosterMapCountry field to given value.

### HasRosterMapCountry

`func (o *CompanyTrialGeographyRow) HasRosterMapCountry() bool`

HasRosterMapCountry returns a boolean if a field has been set.

### SetRosterMapCountryNil

`func (o *CompanyTrialGeographyRow) SetRosterMapCountryNil(b bool)`

 SetRosterMapCountryNil sets the value for RosterMapCountry to be an explicit nil

### UnsetRosterMapCountry
`func (o *CompanyTrialGeographyRow) UnsetRosterMapCountry()`

UnsetRosterMapCountry ensures that no value is present for RosterMapCountry, not even an explicit nil
### GetRosterMapCountryBasis

`func (o *CompanyTrialGeographyRow) GetRosterMapCountryBasis() string`

GetRosterMapCountryBasis returns the RosterMapCountryBasis field if non-nil, zero value otherwise.

### GetRosterMapCountryBasisOk

`func (o *CompanyTrialGeographyRow) GetRosterMapCountryBasisOk() (*string, bool)`

GetRosterMapCountryBasisOk returns a tuple with the RosterMapCountryBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRosterMapCountryBasis

`func (o *CompanyTrialGeographyRow) SetRosterMapCountryBasis(v string)`

SetRosterMapCountryBasis sets RosterMapCountryBasis field to given value.

### HasRosterMapCountryBasis

`func (o *CompanyTrialGeographyRow) HasRosterMapCountryBasis() bool`

HasRosterMapCountryBasis returns a boolean if a field has been set.

### SetRosterMapCountryBasisNil

`func (o *CompanyTrialGeographyRow) SetRosterMapCountryBasisNil(b bool)`

 SetRosterMapCountryBasisNil sets the value for RosterMapCountryBasis to be an explicit nil

### UnsetRosterMapCountryBasis
`func (o *CompanyTrialGeographyRow) UnsetRosterMapCountryBasis()`

UnsetRosterMapCountryBasis ensures that no value is present for RosterMapCountryBasis, not even an explicit nil
### GetRosterNLocatedSites

`func (o *CompanyTrialGeographyRow) GetRosterNLocatedSites() int32`

GetRosterNLocatedSites returns the RosterNLocatedSites field if non-nil, zero value otherwise.

### GetRosterNLocatedSitesOk

`func (o *CompanyTrialGeographyRow) GetRosterNLocatedSitesOk() (*int32, bool)`

GetRosterNLocatedSitesOk returns a tuple with the RosterNLocatedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRosterNLocatedSites

`func (o *CompanyTrialGeographyRow) SetRosterNLocatedSites(v int32)`

SetRosterNLocatedSites sets RosterNLocatedSites field to given value.

### HasRosterNLocatedSites

`func (o *CompanyTrialGeographyRow) HasRosterNLocatedSites() bool`

HasRosterNLocatedSites returns a boolean if a field has been set.

### SetRosterNLocatedSitesNil

`func (o *CompanyTrialGeographyRow) SetRosterNLocatedSitesNil(b bool)`

 SetRosterNLocatedSitesNil sets the value for RosterNLocatedSites to be an explicit nil

### UnsetRosterNLocatedSites
`func (o *CompanyTrialGeographyRow) UnsetRosterNLocatedSites()`

UnsetRosterNLocatedSites ensures that no value is present for RosterNLocatedSites, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


