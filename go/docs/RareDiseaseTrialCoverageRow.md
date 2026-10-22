# RareDiseaseTrialCoverageRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] 
**CoverageState** | Pointer to **NullableString** | no_bridge, bridged_no_trials or trials. | [optional] 
**MappingArms** | Pointer to **[]string** | The mapping arms the disease has a condition under (exact, unspecified). | [optional] 
**MappingTargetCurie** | Pointer to **NullableString** | The MeSH condition or conditions the disease maps to (prefixed identifiers, comma separated), the exact ones where any exist. | [optional] 
**MappingTargetName** | Pointer to **NullableString** | The names of those MeSH conditions. | [optional] 
**MappingSources** | Pointer to **[]string** | The files that assert the mapping. | [optional] 
**MappingSourceReleases** | Pointer to **[]string** | The releases of those files. | [optional] 
**NTrialsExactDirect** | Pointer to **NullableInt64** | Trials linked through an exact mapping whose own condition list names the disease. Count within this column only. | [optional] 
**NTrialsExactBroader** | Pointer to **NullableInt64** | Trials linked through an exact mapping to a broader condition than the trial studies. Count within this column only. | [optional] 
**NTrialsUnspecified** | Pointer to **NullableInt64** | Trials linked through a MONDO cross-reference that states no relation. Count within this column only. | [optional] 
**NOpenDirect** | Pointer to **NullableInt64** | Exact direct trials that are recruiting or not yet recruiting. | [optional] 
**NOngoingDirect** | Pointer to **NullableInt64** | Exact direct trials that are active but not recruiting, or enrolling by invitation. | [optional] 
**NUnverified2yDirect** | Pointer to **NullableInt64** | Exact direct trials whose record was last updated two or more years before the snapshot. | [optional] 
**DirectShare** | Pointer to **NullableFloat32** | The exact direct trials as a share of the exact direct and exact broader trials; null when both are zero. | [optional] 
**NSponsors** | Pointer to **NullableInt64** | Sponsoring organizations of the exact direct trials. | [optional] 
**SnapshotDate** | Pointer to **NullableString** | The date of the registry snapshot. | [optional] 
**FreshnessAsOf** | Pointer to **NullableString** | The registry&#39;s latest update date the snapshot reflects. | [optional] 
**LicenseName** | Pointer to **NullableString** | The licence of the Orphadata Science release the mapping rests on. | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] 
**TrialsLicenseName** | Pointer to **NullableString** | The licence of the trial registry data. | [optional] 
**TrialsAttribution** | Pointer to **NullableString** | The attribution line the trial registry requires. | [optional] 

## Methods

### NewRareDiseaseTrialCoverageRow

`func NewRareDiseaseTrialCoverageRow() *RareDiseaseTrialCoverageRow`

NewRareDiseaseTrialCoverageRow instantiates a new RareDiseaseTrialCoverageRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseaseTrialCoverageRowWithDefaults

`func NewRareDiseaseTrialCoverageRowWithDefaults() *RareDiseaseTrialCoverageRow`

NewRareDiseaseTrialCoverageRowWithDefaults instantiates a new RareDiseaseTrialCoverageRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrphaCode

`func (o *RareDiseaseTrialCoverageRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseaseTrialCoverageRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseaseTrialCoverageRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseaseTrialCoverageRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseaseTrialCoverageRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseaseTrialCoverageRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetCoverageState

`func (o *RareDiseaseTrialCoverageRow) GetCoverageState() string`

GetCoverageState returns the CoverageState field if non-nil, zero value otherwise.

### GetCoverageStateOk

`func (o *RareDiseaseTrialCoverageRow) GetCoverageStateOk() (*string, bool)`

GetCoverageStateOk returns a tuple with the CoverageState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageState

`func (o *RareDiseaseTrialCoverageRow) SetCoverageState(v string)`

SetCoverageState sets CoverageState field to given value.

### HasCoverageState

`func (o *RareDiseaseTrialCoverageRow) HasCoverageState() bool`

HasCoverageState returns a boolean if a field has been set.

### SetCoverageStateNil

`func (o *RareDiseaseTrialCoverageRow) SetCoverageStateNil(b bool)`

 SetCoverageStateNil sets the value for CoverageState to be an explicit nil

### UnsetCoverageState
`func (o *RareDiseaseTrialCoverageRow) UnsetCoverageState()`

UnsetCoverageState ensures that no value is present for CoverageState, not even an explicit nil
### GetMappingArms

`func (o *RareDiseaseTrialCoverageRow) GetMappingArms() []string`

GetMappingArms returns the MappingArms field if non-nil, zero value otherwise.

### GetMappingArmsOk

`func (o *RareDiseaseTrialCoverageRow) GetMappingArmsOk() (*[]string, bool)`

GetMappingArmsOk returns a tuple with the MappingArms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingArms

`func (o *RareDiseaseTrialCoverageRow) SetMappingArms(v []string)`

SetMappingArms sets MappingArms field to given value.

### HasMappingArms

`func (o *RareDiseaseTrialCoverageRow) HasMappingArms() bool`

HasMappingArms returns a boolean if a field has been set.

### SetMappingArmsNil

`func (o *RareDiseaseTrialCoverageRow) SetMappingArmsNil(b bool)`

 SetMappingArmsNil sets the value for MappingArms to be an explicit nil

### UnsetMappingArms
`func (o *RareDiseaseTrialCoverageRow) UnsetMappingArms()`

UnsetMappingArms ensures that no value is present for MappingArms, not even an explicit nil
### GetMappingTargetCurie

`func (o *RareDiseaseTrialCoverageRow) GetMappingTargetCurie() string`

GetMappingTargetCurie returns the MappingTargetCurie field if non-nil, zero value otherwise.

### GetMappingTargetCurieOk

`func (o *RareDiseaseTrialCoverageRow) GetMappingTargetCurieOk() (*string, bool)`

GetMappingTargetCurieOk returns a tuple with the MappingTargetCurie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingTargetCurie

`func (o *RareDiseaseTrialCoverageRow) SetMappingTargetCurie(v string)`

SetMappingTargetCurie sets MappingTargetCurie field to given value.

### HasMappingTargetCurie

`func (o *RareDiseaseTrialCoverageRow) HasMappingTargetCurie() bool`

HasMappingTargetCurie returns a boolean if a field has been set.

### SetMappingTargetCurieNil

`func (o *RareDiseaseTrialCoverageRow) SetMappingTargetCurieNil(b bool)`

 SetMappingTargetCurieNil sets the value for MappingTargetCurie to be an explicit nil

### UnsetMappingTargetCurie
`func (o *RareDiseaseTrialCoverageRow) UnsetMappingTargetCurie()`

UnsetMappingTargetCurie ensures that no value is present for MappingTargetCurie, not even an explicit nil
### GetMappingTargetName

`func (o *RareDiseaseTrialCoverageRow) GetMappingTargetName() string`

GetMappingTargetName returns the MappingTargetName field if non-nil, zero value otherwise.

### GetMappingTargetNameOk

`func (o *RareDiseaseTrialCoverageRow) GetMappingTargetNameOk() (*string, bool)`

GetMappingTargetNameOk returns a tuple with the MappingTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingTargetName

`func (o *RareDiseaseTrialCoverageRow) SetMappingTargetName(v string)`

SetMappingTargetName sets MappingTargetName field to given value.

### HasMappingTargetName

`func (o *RareDiseaseTrialCoverageRow) HasMappingTargetName() bool`

HasMappingTargetName returns a boolean if a field has been set.

### SetMappingTargetNameNil

`func (o *RareDiseaseTrialCoverageRow) SetMappingTargetNameNil(b bool)`

 SetMappingTargetNameNil sets the value for MappingTargetName to be an explicit nil

### UnsetMappingTargetName
`func (o *RareDiseaseTrialCoverageRow) UnsetMappingTargetName()`

UnsetMappingTargetName ensures that no value is present for MappingTargetName, not even an explicit nil
### GetMappingSources

`func (o *RareDiseaseTrialCoverageRow) GetMappingSources() []string`

GetMappingSources returns the MappingSources field if non-nil, zero value otherwise.

### GetMappingSourcesOk

`func (o *RareDiseaseTrialCoverageRow) GetMappingSourcesOk() (*[]string, bool)`

GetMappingSourcesOk returns a tuple with the MappingSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingSources

`func (o *RareDiseaseTrialCoverageRow) SetMappingSources(v []string)`

SetMappingSources sets MappingSources field to given value.

### HasMappingSources

`func (o *RareDiseaseTrialCoverageRow) HasMappingSources() bool`

HasMappingSources returns a boolean if a field has been set.

### SetMappingSourcesNil

`func (o *RareDiseaseTrialCoverageRow) SetMappingSourcesNil(b bool)`

 SetMappingSourcesNil sets the value for MappingSources to be an explicit nil

### UnsetMappingSources
`func (o *RareDiseaseTrialCoverageRow) UnsetMappingSources()`

UnsetMappingSources ensures that no value is present for MappingSources, not even an explicit nil
### GetMappingSourceReleases

`func (o *RareDiseaseTrialCoverageRow) GetMappingSourceReleases() []string`

GetMappingSourceReleases returns the MappingSourceReleases field if non-nil, zero value otherwise.

### GetMappingSourceReleasesOk

`func (o *RareDiseaseTrialCoverageRow) GetMappingSourceReleasesOk() (*[]string, bool)`

GetMappingSourceReleasesOk returns a tuple with the MappingSourceReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingSourceReleases

`func (o *RareDiseaseTrialCoverageRow) SetMappingSourceReleases(v []string)`

SetMappingSourceReleases sets MappingSourceReleases field to given value.

### HasMappingSourceReleases

`func (o *RareDiseaseTrialCoverageRow) HasMappingSourceReleases() bool`

HasMappingSourceReleases returns a boolean if a field has been set.

### SetMappingSourceReleasesNil

`func (o *RareDiseaseTrialCoverageRow) SetMappingSourceReleasesNil(b bool)`

 SetMappingSourceReleasesNil sets the value for MappingSourceReleases to be an explicit nil

### UnsetMappingSourceReleases
`func (o *RareDiseaseTrialCoverageRow) UnsetMappingSourceReleases()`

UnsetMappingSourceReleases ensures that no value is present for MappingSourceReleases, not even an explicit nil
### GetNTrialsExactDirect

`func (o *RareDiseaseTrialCoverageRow) GetNTrialsExactDirect() int64`

GetNTrialsExactDirect returns the NTrialsExactDirect field if non-nil, zero value otherwise.

### GetNTrialsExactDirectOk

`func (o *RareDiseaseTrialCoverageRow) GetNTrialsExactDirectOk() (*int64, bool)`

GetNTrialsExactDirectOk returns a tuple with the NTrialsExactDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsExactDirect

`func (o *RareDiseaseTrialCoverageRow) SetNTrialsExactDirect(v int64)`

SetNTrialsExactDirect sets NTrialsExactDirect field to given value.

### HasNTrialsExactDirect

`func (o *RareDiseaseTrialCoverageRow) HasNTrialsExactDirect() bool`

HasNTrialsExactDirect returns a boolean if a field has been set.

### SetNTrialsExactDirectNil

`func (o *RareDiseaseTrialCoverageRow) SetNTrialsExactDirectNil(b bool)`

 SetNTrialsExactDirectNil sets the value for NTrialsExactDirect to be an explicit nil

### UnsetNTrialsExactDirect
`func (o *RareDiseaseTrialCoverageRow) UnsetNTrialsExactDirect()`

UnsetNTrialsExactDirect ensures that no value is present for NTrialsExactDirect, not even an explicit nil
### GetNTrialsExactBroader

`func (o *RareDiseaseTrialCoverageRow) GetNTrialsExactBroader() int64`

GetNTrialsExactBroader returns the NTrialsExactBroader field if non-nil, zero value otherwise.

### GetNTrialsExactBroaderOk

`func (o *RareDiseaseTrialCoverageRow) GetNTrialsExactBroaderOk() (*int64, bool)`

GetNTrialsExactBroaderOk returns a tuple with the NTrialsExactBroader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsExactBroader

`func (o *RareDiseaseTrialCoverageRow) SetNTrialsExactBroader(v int64)`

SetNTrialsExactBroader sets NTrialsExactBroader field to given value.

### HasNTrialsExactBroader

`func (o *RareDiseaseTrialCoverageRow) HasNTrialsExactBroader() bool`

HasNTrialsExactBroader returns a boolean if a field has been set.

### SetNTrialsExactBroaderNil

`func (o *RareDiseaseTrialCoverageRow) SetNTrialsExactBroaderNil(b bool)`

 SetNTrialsExactBroaderNil sets the value for NTrialsExactBroader to be an explicit nil

### UnsetNTrialsExactBroader
`func (o *RareDiseaseTrialCoverageRow) UnsetNTrialsExactBroader()`

UnsetNTrialsExactBroader ensures that no value is present for NTrialsExactBroader, not even an explicit nil
### GetNTrialsUnspecified

`func (o *RareDiseaseTrialCoverageRow) GetNTrialsUnspecified() int64`

GetNTrialsUnspecified returns the NTrialsUnspecified field if non-nil, zero value otherwise.

### GetNTrialsUnspecifiedOk

`func (o *RareDiseaseTrialCoverageRow) GetNTrialsUnspecifiedOk() (*int64, bool)`

GetNTrialsUnspecifiedOk returns a tuple with the NTrialsUnspecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsUnspecified

`func (o *RareDiseaseTrialCoverageRow) SetNTrialsUnspecified(v int64)`

SetNTrialsUnspecified sets NTrialsUnspecified field to given value.

### HasNTrialsUnspecified

`func (o *RareDiseaseTrialCoverageRow) HasNTrialsUnspecified() bool`

HasNTrialsUnspecified returns a boolean if a field has been set.

### SetNTrialsUnspecifiedNil

`func (o *RareDiseaseTrialCoverageRow) SetNTrialsUnspecifiedNil(b bool)`

 SetNTrialsUnspecifiedNil sets the value for NTrialsUnspecified to be an explicit nil

### UnsetNTrialsUnspecified
`func (o *RareDiseaseTrialCoverageRow) UnsetNTrialsUnspecified()`

UnsetNTrialsUnspecified ensures that no value is present for NTrialsUnspecified, not even an explicit nil
### GetNOpenDirect

`func (o *RareDiseaseTrialCoverageRow) GetNOpenDirect() int64`

GetNOpenDirect returns the NOpenDirect field if non-nil, zero value otherwise.

### GetNOpenDirectOk

`func (o *RareDiseaseTrialCoverageRow) GetNOpenDirectOk() (*int64, bool)`

GetNOpenDirectOk returns a tuple with the NOpenDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOpenDirect

`func (o *RareDiseaseTrialCoverageRow) SetNOpenDirect(v int64)`

SetNOpenDirect sets NOpenDirect field to given value.

### HasNOpenDirect

`func (o *RareDiseaseTrialCoverageRow) HasNOpenDirect() bool`

HasNOpenDirect returns a boolean if a field has been set.

### SetNOpenDirectNil

`func (o *RareDiseaseTrialCoverageRow) SetNOpenDirectNil(b bool)`

 SetNOpenDirectNil sets the value for NOpenDirect to be an explicit nil

### UnsetNOpenDirect
`func (o *RareDiseaseTrialCoverageRow) UnsetNOpenDirect()`

UnsetNOpenDirect ensures that no value is present for NOpenDirect, not even an explicit nil
### GetNOngoingDirect

`func (o *RareDiseaseTrialCoverageRow) GetNOngoingDirect() int64`

GetNOngoingDirect returns the NOngoingDirect field if non-nil, zero value otherwise.

### GetNOngoingDirectOk

`func (o *RareDiseaseTrialCoverageRow) GetNOngoingDirectOk() (*int64, bool)`

GetNOngoingDirectOk returns a tuple with the NOngoingDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOngoingDirect

`func (o *RareDiseaseTrialCoverageRow) SetNOngoingDirect(v int64)`

SetNOngoingDirect sets NOngoingDirect field to given value.

### HasNOngoingDirect

`func (o *RareDiseaseTrialCoverageRow) HasNOngoingDirect() bool`

HasNOngoingDirect returns a boolean if a field has been set.

### SetNOngoingDirectNil

`func (o *RareDiseaseTrialCoverageRow) SetNOngoingDirectNil(b bool)`

 SetNOngoingDirectNil sets the value for NOngoingDirect to be an explicit nil

### UnsetNOngoingDirect
`func (o *RareDiseaseTrialCoverageRow) UnsetNOngoingDirect()`

UnsetNOngoingDirect ensures that no value is present for NOngoingDirect, not even an explicit nil
### GetNUnverified2yDirect

`func (o *RareDiseaseTrialCoverageRow) GetNUnverified2yDirect() int64`

GetNUnverified2yDirect returns the NUnverified2yDirect field if non-nil, zero value otherwise.

### GetNUnverified2yDirectOk

`func (o *RareDiseaseTrialCoverageRow) GetNUnverified2yDirectOk() (*int64, bool)`

GetNUnverified2yDirectOk returns a tuple with the NUnverified2yDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNUnverified2yDirect

`func (o *RareDiseaseTrialCoverageRow) SetNUnverified2yDirect(v int64)`

SetNUnverified2yDirect sets NUnverified2yDirect field to given value.

### HasNUnverified2yDirect

`func (o *RareDiseaseTrialCoverageRow) HasNUnverified2yDirect() bool`

HasNUnverified2yDirect returns a boolean if a field has been set.

### SetNUnverified2yDirectNil

`func (o *RareDiseaseTrialCoverageRow) SetNUnverified2yDirectNil(b bool)`

 SetNUnverified2yDirectNil sets the value for NUnverified2yDirect to be an explicit nil

### UnsetNUnverified2yDirect
`func (o *RareDiseaseTrialCoverageRow) UnsetNUnverified2yDirect()`

UnsetNUnverified2yDirect ensures that no value is present for NUnverified2yDirect, not even an explicit nil
### GetDirectShare

`func (o *RareDiseaseTrialCoverageRow) GetDirectShare() float32`

GetDirectShare returns the DirectShare field if non-nil, zero value otherwise.

### GetDirectShareOk

`func (o *RareDiseaseTrialCoverageRow) GetDirectShareOk() (*float32, bool)`

GetDirectShareOk returns a tuple with the DirectShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectShare

`func (o *RareDiseaseTrialCoverageRow) SetDirectShare(v float32)`

SetDirectShare sets DirectShare field to given value.

### HasDirectShare

`func (o *RareDiseaseTrialCoverageRow) HasDirectShare() bool`

HasDirectShare returns a boolean if a field has been set.

### SetDirectShareNil

`func (o *RareDiseaseTrialCoverageRow) SetDirectShareNil(b bool)`

 SetDirectShareNil sets the value for DirectShare to be an explicit nil

### UnsetDirectShare
`func (o *RareDiseaseTrialCoverageRow) UnsetDirectShare()`

UnsetDirectShare ensures that no value is present for DirectShare, not even an explicit nil
### GetNSponsors

`func (o *RareDiseaseTrialCoverageRow) GetNSponsors() int64`

GetNSponsors returns the NSponsors field if non-nil, zero value otherwise.

### GetNSponsorsOk

`func (o *RareDiseaseTrialCoverageRow) GetNSponsorsOk() (*int64, bool)`

GetNSponsorsOk returns a tuple with the NSponsors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSponsors

`func (o *RareDiseaseTrialCoverageRow) SetNSponsors(v int64)`

SetNSponsors sets NSponsors field to given value.

### HasNSponsors

`func (o *RareDiseaseTrialCoverageRow) HasNSponsors() bool`

HasNSponsors returns a boolean if a field has been set.

### SetNSponsorsNil

`func (o *RareDiseaseTrialCoverageRow) SetNSponsorsNil(b bool)`

 SetNSponsorsNil sets the value for NSponsors to be an explicit nil

### UnsetNSponsors
`func (o *RareDiseaseTrialCoverageRow) UnsetNSponsors()`

UnsetNSponsors ensures that no value is present for NSponsors, not even an explicit nil
### GetSnapshotDate

`func (o *RareDiseaseTrialCoverageRow) GetSnapshotDate() string`

GetSnapshotDate returns the SnapshotDate field if non-nil, zero value otherwise.

### GetSnapshotDateOk

`func (o *RareDiseaseTrialCoverageRow) GetSnapshotDateOk() (*string, bool)`

GetSnapshotDateOk returns a tuple with the SnapshotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotDate

`func (o *RareDiseaseTrialCoverageRow) SetSnapshotDate(v string)`

SetSnapshotDate sets SnapshotDate field to given value.

### HasSnapshotDate

`func (o *RareDiseaseTrialCoverageRow) HasSnapshotDate() bool`

HasSnapshotDate returns a boolean if a field has been set.

### SetSnapshotDateNil

`func (o *RareDiseaseTrialCoverageRow) SetSnapshotDateNil(b bool)`

 SetSnapshotDateNil sets the value for SnapshotDate to be an explicit nil

### UnsetSnapshotDate
`func (o *RareDiseaseTrialCoverageRow) UnsetSnapshotDate()`

UnsetSnapshotDate ensures that no value is present for SnapshotDate, not even an explicit nil
### GetFreshnessAsOf

`func (o *RareDiseaseTrialCoverageRow) GetFreshnessAsOf() string`

GetFreshnessAsOf returns the FreshnessAsOf field if non-nil, zero value otherwise.

### GetFreshnessAsOfOk

`func (o *RareDiseaseTrialCoverageRow) GetFreshnessAsOfOk() (*string, bool)`

GetFreshnessAsOfOk returns a tuple with the FreshnessAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshnessAsOf

`func (o *RareDiseaseTrialCoverageRow) SetFreshnessAsOf(v string)`

SetFreshnessAsOf sets FreshnessAsOf field to given value.

### HasFreshnessAsOf

`func (o *RareDiseaseTrialCoverageRow) HasFreshnessAsOf() bool`

HasFreshnessAsOf returns a boolean if a field has been set.

### SetFreshnessAsOfNil

`func (o *RareDiseaseTrialCoverageRow) SetFreshnessAsOfNil(b bool)`

 SetFreshnessAsOfNil sets the value for FreshnessAsOf to be an explicit nil

### UnsetFreshnessAsOf
`func (o *RareDiseaseTrialCoverageRow) UnsetFreshnessAsOf()`

UnsetFreshnessAsOf ensures that no value is present for FreshnessAsOf, not even an explicit nil
### GetLicenseName

`func (o *RareDiseaseTrialCoverageRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *RareDiseaseTrialCoverageRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *RareDiseaseTrialCoverageRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *RareDiseaseTrialCoverageRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *RareDiseaseTrialCoverageRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *RareDiseaseTrialCoverageRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *RareDiseaseTrialCoverageRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RareDiseaseTrialCoverageRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RareDiseaseTrialCoverageRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RareDiseaseTrialCoverageRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RareDiseaseTrialCoverageRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RareDiseaseTrialCoverageRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil
### GetTrialsLicenseName

`func (o *RareDiseaseTrialCoverageRow) GetTrialsLicenseName() string`

GetTrialsLicenseName returns the TrialsLicenseName field if non-nil, zero value otherwise.

### GetTrialsLicenseNameOk

`func (o *RareDiseaseTrialCoverageRow) GetTrialsLicenseNameOk() (*string, bool)`

GetTrialsLicenseNameOk returns a tuple with the TrialsLicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialsLicenseName

`func (o *RareDiseaseTrialCoverageRow) SetTrialsLicenseName(v string)`

SetTrialsLicenseName sets TrialsLicenseName field to given value.

### HasTrialsLicenseName

`func (o *RareDiseaseTrialCoverageRow) HasTrialsLicenseName() bool`

HasTrialsLicenseName returns a boolean if a field has been set.

### SetTrialsLicenseNameNil

`func (o *RareDiseaseTrialCoverageRow) SetTrialsLicenseNameNil(b bool)`

 SetTrialsLicenseNameNil sets the value for TrialsLicenseName to be an explicit nil

### UnsetTrialsLicenseName
`func (o *RareDiseaseTrialCoverageRow) UnsetTrialsLicenseName()`

UnsetTrialsLicenseName ensures that no value is present for TrialsLicenseName, not even an explicit nil
### GetTrialsAttribution

`func (o *RareDiseaseTrialCoverageRow) GetTrialsAttribution() string`

GetTrialsAttribution returns the TrialsAttribution field if non-nil, zero value otherwise.

### GetTrialsAttributionOk

`func (o *RareDiseaseTrialCoverageRow) GetTrialsAttributionOk() (*string, bool)`

GetTrialsAttributionOk returns a tuple with the TrialsAttribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialsAttribution

`func (o *RareDiseaseTrialCoverageRow) SetTrialsAttribution(v string)`

SetTrialsAttribution sets TrialsAttribution field to given value.

### HasTrialsAttribution

`func (o *RareDiseaseTrialCoverageRow) HasTrialsAttribution() bool`

HasTrialsAttribution returns a boolean if a field has been set.

### SetTrialsAttributionNil

`func (o *RareDiseaseTrialCoverageRow) SetTrialsAttributionNil(b bool)`

 SetTrialsAttributionNil sets the value for TrialsAttribution to be an explicit nil

### UnsetTrialsAttribution
`func (o *RareDiseaseTrialCoverageRow) UnsetTrialsAttribution()`

UnsetTrialsAttribution ensures that no value is present for TrialsAttribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


