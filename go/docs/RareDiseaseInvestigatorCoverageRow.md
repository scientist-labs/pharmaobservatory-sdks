# RareDiseaseInvestigatorCoverageRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] 
**CoverageState** | Pointer to **NullableString** | no_exact_mapping, no_direct_trials, no_publishable_investigator or computed. | [optional] 
**NDirectTrials** | Pointer to **NullableInt64** | Exact direct trials of the disease. | [optional] 
**NDirectTrialsWithNamedPi** | Pointer to **NullableInt64** | Those trials that name a principal investigator. | [optional] 
**NPiTrialEdges** | Pointer to **NullableInt64** | Trial and principal investigator pairs among them. | [optional] 
**NPiPeople** | Pointer to **NullableInt64** | Distinct people named as principal investigator. | [optional] 
**NInvestigators** | Pointer to **NullableInt64** | Rows of rare_disease_investigators for the disease. | [optional] 
**NNameAmbiguous** | Pointer to **NullableInt64** | Investigators whose name another investigator of the disease shares. | [optional] 
**RegistryReferenceDate** | Pointer to **NullableString** | The registry&#39;s latest update date the snapshot reflects. | [optional] 
**LicenseName** | Pointer to **NullableString** | The licence of the Orphadata Science release the disease list rests on. | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRareDiseaseInvestigatorCoverageRow

`func NewRareDiseaseInvestigatorCoverageRow() *RareDiseaseInvestigatorCoverageRow`

NewRareDiseaseInvestigatorCoverageRow instantiates a new RareDiseaseInvestigatorCoverageRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseaseInvestigatorCoverageRowWithDefaults

`func NewRareDiseaseInvestigatorCoverageRowWithDefaults() *RareDiseaseInvestigatorCoverageRow`

NewRareDiseaseInvestigatorCoverageRowWithDefaults instantiates a new RareDiseaseInvestigatorCoverageRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrphaCode

`func (o *RareDiseaseInvestigatorCoverageRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseaseInvestigatorCoverageRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseaseInvestigatorCoverageRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseaseInvestigatorCoverageRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseaseInvestigatorCoverageRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseaseInvestigatorCoverageRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetCoverageState

`func (o *RareDiseaseInvestigatorCoverageRow) GetCoverageState() string`

GetCoverageState returns the CoverageState field if non-nil, zero value otherwise.

### GetCoverageStateOk

`func (o *RareDiseaseInvestigatorCoverageRow) GetCoverageStateOk() (*string, bool)`

GetCoverageStateOk returns a tuple with the CoverageState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageState

`func (o *RareDiseaseInvestigatorCoverageRow) SetCoverageState(v string)`

SetCoverageState sets CoverageState field to given value.

### HasCoverageState

`func (o *RareDiseaseInvestigatorCoverageRow) HasCoverageState() bool`

HasCoverageState returns a boolean if a field has been set.

### SetCoverageStateNil

`func (o *RareDiseaseInvestigatorCoverageRow) SetCoverageStateNil(b bool)`

 SetCoverageStateNil sets the value for CoverageState to be an explicit nil

### UnsetCoverageState
`func (o *RareDiseaseInvestigatorCoverageRow) UnsetCoverageState()`

UnsetCoverageState ensures that no value is present for CoverageState, not even an explicit nil
### GetNDirectTrials

`func (o *RareDiseaseInvestigatorCoverageRow) GetNDirectTrials() int64`

GetNDirectTrials returns the NDirectTrials field if non-nil, zero value otherwise.

### GetNDirectTrialsOk

`func (o *RareDiseaseInvestigatorCoverageRow) GetNDirectTrialsOk() (*int64, bool)`

GetNDirectTrialsOk returns a tuple with the NDirectTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDirectTrials

`func (o *RareDiseaseInvestigatorCoverageRow) SetNDirectTrials(v int64)`

SetNDirectTrials sets NDirectTrials field to given value.

### HasNDirectTrials

`func (o *RareDiseaseInvestigatorCoverageRow) HasNDirectTrials() bool`

HasNDirectTrials returns a boolean if a field has been set.

### SetNDirectTrialsNil

`func (o *RareDiseaseInvestigatorCoverageRow) SetNDirectTrialsNil(b bool)`

 SetNDirectTrialsNil sets the value for NDirectTrials to be an explicit nil

### UnsetNDirectTrials
`func (o *RareDiseaseInvestigatorCoverageRow) UnsetNDirectTrials()`

UnsetNDirectTrials ensures that no value is present for NDirectTrials, not even an explicit nil
### GetNDirectTrialsWithNamedPi

`func (o *RareDiseaseInvestigatorCoverageRow) GetNDirectTrialsWithNamedPi() int64`

GetNDirectTrialsWithNamedPi returns the NDirectTrialsWithNamedPi field if non-nil, zero value otherwise.

### GetNDirectTrialsWithNamedPiOk

`func (o *RareDiseaseInvestigatorCoverageRow) GetNDirectTrialsWithNamedPiOk() (*int64, bool)`

GetNDirectTrialsWithNamedPiOk returns a tuple with the NDirectTrialsWithNamedPi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDirectTrialsWithNamedPi

`func (o *RareDiseaseInvestigatorCoverageRow) SetNDirectTrialsWithNamedPi(v int64)`

SetNDirectTrialsWithNamedPi sets NDirectTrialsWithNamedPi field to given value.

### HasNDirectTrialsWithNamedPi

`func (o *RareDiseaseInvestigatorCoverageRow) HasNDirectTrialsWithNamedPi() bool`

HasNDirectTrialsWithNamedPi returns a boolean if a field has been set.

### SetNDirectTrialsWithNamedPiNil

`func (o *RareDiseaseInvestigatorCoverageRow) SetNDirectTrialsWithNamedPiNil(b bool)`

 SetNDirectTrialsWithNamedPiNil sets the value for NDirectTrialsWithNamedPi to be an explicit nil

### UnsetNDirectTrialsWithNamedPi
`func (o *RareDiseaseInvestigatorCoverageRow) UnsetNDirectTrialsWithNamedPi()`

UnsetNDirectTrialsWithNamedPi ensures that no value is present for NDirectTrialsWithNamedPi, not even an explicit nil
### GetNPiTrialEdges

`func (o *RareDiseaseInvestigatorCoverageRow) GetNPiTrialEdges() int64`

GetNPiTrialEdges returns the NPiTrialEdges field if non-nil, zero value otherwise.

### GetNPiTrialEdgesOk

`func (o *RareDiseaseInvestigatorCoverageRow) GetNPiTrialEdgesOk() (*int64, bool)`

GetNPiTrialEdgesOk returns a tuple with the NPiTrialEdges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPiTrialEdges

`func (o *RareDiseaseInvestigatorCoverageRow) SetNPiTrialEdges(v int64)`

SetNPiTrialEdges sets NPiTrialEdges field to given value.

### HasNPiTrialEdges

`func (o *RareDiseaseInvestigatorCoverageRow) HasNPiTrialEdges() bool`

HasNPiTrialEdges returns a boolean if a field has been set.

### SetNPiTrialEdgesNil

`func (o *RareDiseaseInvestigatorCoverageRow) SetNPiTrialEdgesNil(b bool)`

 SetNPiTrialEdgesNil sets the value for NPiTrialEdges to be an explicit nil

### UnsetNPiTrialEdges
`func (o *RareDiseaseInvestigatorCoverageRow) UnsetNPiTrialEdges()`

UnsetNPiTrialEdges ensures that no value is present for NPiTrialEdges, not even an explicit nil
### GetNPiPeople

`func (o *RareDiseaseInvestigatorCoverageRow) GetNPiPeople() int64`

GetNPiPeople returns the NPiPeople field if non-nil, zero value otherwise.

### GetNPiPeopleOk

`func (o *RareDiseaseInvestigatorCoverageRow) GetNPiPeopleOk() (*int64, bool)`

GetNPiPeopleOk returns a tuple with the NPiPeople field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPiPeople

`func (o *RareDiseaseInvestigatorCoverageRow) SetNPiPeople(v int64)`

SetNPiPeople sets NPiPeople field to given value.

### HasNPiPeople

`func (o *RareDiseaseInvestigatorCoverageRow) HasNPiPeople() bool`

HasNPiPeople returns a boolean if a field has been set.

### SetNPiPeopleNil

`func (o *RareDiseaseInvestigatorCoverageRow) SetNPiPeopleNil(b bool)`

 SetNPiPeopleNil sets the value for NPiPeople to be an explicit nil

### UnsetNPiPeople
`func (o *RareDiseaseInvestigatorCoverageRow) UnsetNPiPeople()`

UnsetNPiPeople ensures that no value is present for NPiPeople, not even an explicit nil
### GetNInvestigators

`func (o *RareDiseaseInvestigatorCoverageRow) GetNInvestigators() int64`

GetNInvestigators returns the NInvestigators field if non-nil, zero value otherwise.

### GetNInvestigatorsOk

`func (o *RareDiseaseInvestigatorCoverageRow) GetNInvestigatorsOk() (*int64, bool)`

GetNInvestigatorsOk returns a tuple with the NInvestigators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNInvestigators

`func (o *RareDiseaseInvestigatorCoverageRow) SetNInvestigators(v int64)`

SetNInvestigators sets NInvestigators field to given value.

### HasNInvestigators

`func (o *RareDiseaseInvestigatorCoverageRow) HasNInvestigators() bool`

HasNInvestigators returns a boolean if a field has been set.

### SetNInvestigatorsNil

`func (o *RareDiseaseInvestigatorCoverageRow) SetNInvestigatorsNil(b bool)`

 SetNInvestigatorsNil sets the value for NInvestigators to be an explicit nil

### UnsetNInvestigators
`func (o *RareDiseaseInvestigatorCoverageRow) UnsetNInvestigators()`

UnsetNInvestigators ensures that no value is present for NInvestigators, not even an explicit nil
### GetNNameAmbiguous

`func (o *RareDiseaseInvestigatorCoverageRow) GetNNameAmbiguous() int64`

GetNNameAmbiguous returns the NNameAmbiguous field if non-nil, zero value otherwise.

### GetNNameAmbiguousOk

`func (o *RareDiseaseInvestigatorCoverageRow) GetNNameAmbiguousOk() (*int64, bool)`

GetNNameAmbiguousOk returns a tuple with the NNameAmbiguous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNNameAmbiguous

`func (o *RareDiseaseInvestigatorCoverageRow) SetNNameAmbiguous(v int64)`

SetNNameAmbiguous sets NNameAmbiguous field to given value.

### HasNNameAmbiguous

`func (o *RareDiseaseInvestigatorCoverageRow) HasNNameAmbiguous() bool`

HasNNameAmbiguous returns a boolean if a field has been set.

### SetNNameAmbiguousNil

`func (o *RareDiseaseInvestigatorCoverageRow) SetNNameAmbiguousNil(b bool)`

 SetNNameAmbiguousNil sets the value for NNameAmbiguous to be an explicit nil

### UnsetNNameAmbiguous
`func (o *RareDiseaseInvestigatorCoverageRow) UnsetNNameAmbiguous()`

UnsetNNameAmbiguous ensures that no value is present for NNameAmbiguous, not even an explicit nil
### GetRegistryReferenceDate

`func (o *RareDiseaseInvestigatorCoverageRow) GetRegistryReferenceDate() string`

GetRegistryReferenceDate returns the RegistryReferenceDate field if non-nil, zero value otherwise.

### GetRegistryReferenceDateOk

`func (o *RareDiseaseInvestigatorCoverageRow) GetRegistryReferenceDateOk() (*string, bool)`

GetRegistryReferenceDateOk returns a tuple with the RegistryReferenceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryReferenceDate

`func (o *RareDiseaseInvestigatorCoverageRow) SetRegistryReferenceDate(v string)`

SetRegistryReferenceDate sets RegistryReferenceDate field to given value.

### HasRegistryReferenceDate

`func (o *RareDiseaseInvestigatorCoverageRow) HasRegistryReferenceDate() bool`

HasRegistryReferenceDate returns a boolean if a field has been set.

### SetRegistryReferenceDateNil

`func (o *RareDiseaseInvestigatorCoverageRow) SetRegistryReferenceDateNil(b bool)`

 SetRegistryReferenceDateNil sets the value for RegistryReferenceDate to be an explicit nil

### UnsetRegistryReferenceDate
`func (o *RareDiseaseInvestigatorCoverageRow) UnsetRegistryReferenceDate()`

UnsetRegistryReferenceDate ensures that no value is present for RegistryReferenceDate, not even an explicit nil
### GetLicenseName

`func (o *RareDiseaseInvestigatorCoverageRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *RareDiseaseInvestigatorCoverageRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *RareDiseaseInvestigatorCoverageRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *RareDiseaseInvestigatorCoverageRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *RareDiseaseInvestigatorCoverageRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *RareDiseaseInvestigatorCoverageRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *RareDiseaseInvestigatorCoverageRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RareDiseaseInvestigatorCoverageRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RareDiseaseInvestigatorCoverageRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RareDiseaseInvestigatorCoverageRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RareDiseaseInvestigatorCoverageRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RareDiseaseInvestigatorCoverageRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


