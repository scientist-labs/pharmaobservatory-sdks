# RareDiseaseSponsorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] 
**SponsorName** | Pointer to **NullableString** | The sponsor&#39;s name as filed in the registry. | [optional] 
**SponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the sponsor resolved to. | [optional] 
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the sponsor&#39;s company (null when unresolved). | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the sponsor&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the sponsor&#39;s group. | [optional] 
**RegistrySponsorName** | Pointer to **NullableString** | The sponsor&#39;s most common name as filed across the disease&#39;s trials. | [optional] 
**OrgType** | Pointer to **NullableString** | The organization&#39;s type. | [optional] 
**SponsorAgencyClass** | Pointer to **NullableString** | The funder class the registry declared for the sponsor. | [optional] 
**SponsorResolutionConfidence** | Pointer to **NullableString** | How firmly the sponsor resolved to an organization record (high, medium, low or unresolved). | [optional] 
**NTrials** | Pointer to **NullableInt64** | Distinct exact direct trials of the disease the sponsor leads or collaborates on. | [optional] 
**NTrialsLead** | Pointer to **NullableInt64** | Those trials the sponsor leads. | [optional] 
**NTrialsCollaborator** | Pointer to **NullableInt64** | Those trials the sponsor collaborates on. | [optional] 
**NCompleted** | Pointer to **NullableInt64** | Those trials the registry marks completed. | [optional] 
**NTerminated** | Pointer to **NullableInt64** | Those trials marked terminated. | [optional] 
**NWithdrawn** | Pointer to **NullableInt64** | Those trials marked withdrawn. | [optional] 
**NSuspended** | Pointer to **NullableInt64** | Those trials marked suspended. | [optional] 
**NOpen** | Pointer to **NullableInt64** | Those trials recruiting or not yet recruiting. | [optional] 
**NOngoing** | Pointer to **NullableInt64** | Those trials active but not recruiting, or enrolling by invitation. | [optional] 
**NOther** | Pointer to **NullableInt64** | Those trials in any other status. | [optional] 
**NUnverified3y** | Pointer to **NullableInt64** | Those trials whose record was last updated three or more years before the snapshot. | [optional] 
**FirstTrialStart** | Pointer to **NullableString** | The earliest start date among those trials. | [optional] 
**LastTrialStart** | Pointer to **NullableString** | The latest start date among those trials. | [optional] 
**FirstTrialStartYear** | Pointer to **NullableInt32** | The year of the earliest start. | [optional] 
**LastTrialStartYear** | Pointer to **NullableInt32** | The year of the latest start. | [optional] 
**LastUpdatePostedDate** | Pointer to **NullableString** | The latest registry update among those trials. | [optional] 
**IsDormant** | Pointer to **NullableBool** | True when the sponsor&#39;s latest trial start is more than five years before the registry snapshot. | [optional] 
**MappingArm** | Pointer to **NullableString** | The mapping arm the dataset is built at (always exact). | [optional] 
**EdgeScope** | Pointer to **NullableString** | The edge scope the dataset is built at (always direct). | [optional] 
**FreshnessAsOf** | Pointer to **NullableString** | The registry&#39;s latest update date the snapshot reflects. | [optional] 
**LicenseName** | Pointer to **NullableString** | The licence of the trial registry data. | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRareDiseaseSponsorsRow

`func NewRareDiseaseSponsorsRow() *RareDiseaseSponsorsRow`

NewRareDiseaseSponsorsRow instantiates a new RareDiseaseSponsorsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseaseSponsorsRowWithDefaults

`func NewRareDiseaseSponsorsRowWithDefaults() *RareDiseaseSponsorsRow`

NewRareDiseaseSponsorsRowWithDefaults instantiates a new RareDiseaseSponsorsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrphaCode

`func (o *RareDiseaseSponsorsRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseaseSponsorsRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseaseSponsorsRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseaseSponsorsRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseaseSponsorsRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseaseSponsorsRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetSponsorName

`func (o *RareDiseaseSponsorsRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *RareDiseaseSponsorsRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *RareDiseaseSponsorsRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *RareDiseaseSponsorsRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *RareDiseaseSponsorsRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *RareDiseaseSponsorsRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorOrgId

`func (o *RareDiseaseSponsorsRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *RareDiseaseSponsorsRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *RareDiseaseSponsorsRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *RareDiseaseSponsorsRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *RareDiseaseSponsorsRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *RareDiseaseSponsorsRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorCompanyId

`func (o *RareDiseaseSponsorsRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *RareDiseaseSponsorsRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *RareDiseaseSponsorsRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *RareDiseaseSponsorsRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *RareDiseaseSponsorsRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *RareDiseaseSponsorsRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *RareDiseaseSponsorsRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *RareDiseaseSponsorsRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *RareDiseaseSponsorsRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *RareDiseaseSponsorsRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *RareDiseaseSponsorsRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *RareDiseaseSponsorsRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *RareDiseaseSponsorsRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *RareDiseaseSponsorsRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *RareDiseaseSponsorsRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *RareDiseaseSponsorsRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *RareDiseaseSponsorsRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *RareDiseaseSponsorsRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *RareDiseaseSponsorsRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *RareDiseaseSponsorsRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *RareDiseaseSponsorsRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *RareDiseaseSponsorsRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *RareDiseaseSponsorsRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *RareDiseaseSponsorsRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil
### GetRegistrySponsorName

`func (o *RareDiseaseSponsorsRow) GetRegistrySponsorName() string`

GetRegistrySponsorName returns the RegistrySponsorName field if non-nil, zero value otherwise.

### GetRegistrySponsorNameOk

`func (o *RareDiseaseSponsorsRow) GetRegistrySponsorNameOk() (*string, bool)`

GetRegistrySponsorNameOk returns a tuple with the RegistrySponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrySponsorName

`func (o *RareDiseaseSponsorsRow) SetRegistrySponsorName(v string)`

SetRegistrySponsorName sets RegistrySponsorName field to given value.

### HasRegistrySponsorName

`func (o *RareDiseaseSponsorsRow) HasRegistrySponsorName() bool`

HasRegistrySponsorName returns a boolean if a field has been set.

### SetRegistrySponsorNameNil

`func (o *RareDiseaseSponsorsRow) SetRegistrySponsorNameNil(b bool)`

 SetRegistrySponsorNameNil sets the value for RegistrySponsorName to be an explicit nil

### UnsetRegistrySponsorName
`func (o *RareDiseaseSponsorsRow) UnsetRegistrySponsorName()`

UnsetRegistrySponsorName ensures that no value is present for RegistrySponsorName, not even an explicit nil
### GetOrgType

`func (o *RareDiseaseSponsorsRow) GetOrgType() string`

GetOrgType returns the OrgType field if non-nil, zero value otherwise.

### GetOrgTypeOk

`func (o *RareDiseaseSponsorsRow) GetOrgTypeOk() (*string, bool)`

GetOrgTypeOk returns a tuple with the OrgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgType

`func (o *RareDiseaseSponsorsRow) SetOrgType(v string)`

SetOrgType sets OrgType field to given value.

### HasOrgType

`func (o *RareDiseaseSponsorsRow) HasOrgType() bool`

HasOrgType returns a boolean if a field has been set.

### SetOrgTypeNil

`func (o *RareDiseaseSponsorsRow) SetOrgTypeNil(b bool)`

 SetOrgTypeNil sets the value for OrgType to be an explicit nil

### UnsetOrgType
`func (o *RareDiseaseSponsorsRow) UnsetOrgType()`

UnsetOrgType ensures that no value is present for OrgType, not even an explicit nil
### GetSponsorAgencyClass

`func (o *RareDiseaseSponsorsRow) GetSponsorAgencyClass() string`

GetSponsorAgencyClass returns the SponsorAgencyClass field if non-nil, zero value otherwise.

### GetSponsorAgencyClassOk

`func (o *RareDiseaseSponsorsRow) GetSponsorAgencyClassOk() (*string, bool)`

GetSponsorAgencyClassOk returns a tuple with the SponsorAgencyClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorAgencyClass

`func (o *RareDiseaseSponsorsRow) SetSponsorAgencyClass(v string)`

SetSponsorAgencyClass sets SponsorAgencyClass field to given value.

### HasSponsorAgencyClass

`func (o *RareDiseaseSponsorsRow) HasSponsorAgencyClass() bool`

HasSponsorAgencyClass returns a boolean if a field has been set.

### SetSponsorAgencyClassNil

`func (o *RareDiseaseSponsorsRow) SetSponsorAgencyClassNil(b bool)`

 SetSponsorAgencyClassNil sets the value for SponsorAgencyClass to be an explicit nil

### UnsetSponsorAgencyClass
`func (o *RareDiseaseSponsorsRow) UnsetSponsorAgencyClass()`

UnsetSponsorAgencyClass ensures that no value is present for SponsorAgencyClass, not even an explicit nil
### GetSponsorResolutionConfidence

`func (o *RareDiseaseSponsorsRow) GetSponsorResolutionConfidence() string`

GetSponsorResolutionConfidence returns the SponsorResolutionConfidence field if non-nil, zero value otherwise.

### GetSponsorResolutionConfidenceOk

`func (o *RareDiseaseSponsorsRow) GetSponsorResolutionConfidenceOk() (*string, bool)`

GetSponsorResolutionConfidenceOk returns a tuple with the SponsorResolutionConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorResolutionConfidence

`func (o *RareDiseaseSponsorsRow) SetSponsorResolutionConfidence(v string)`

SetSponsorResolutionConfidence sets SponsorResolutionConfidence field to given value.

### HasSponsorResolutionConfidence

`func (o *RareDiseaseSponsorsRow) HasSponsorResolutionConfidence() bool`

HasSponsorResolutionConfidence returns a boolean if a field has been set.

### SetSponsorResolutionConfidenceNil

`func (o *RareDiseaseSponsorsRow) SetSponsorResolutionConfidenceNil(b bool)`

 SetSponsorResolutionConfidenceNil sets the value for SponsorResolutionConfidence to be an explicit nil

### UnsetSponsorResolutionConfidence
`func (o *RareDiseaseSponsorsRow) UnsetSponsorResolutionConfidence()`

UnsetSponsorResolutionConfidence ensures that no value is present for SponsorResolutionConfidence, not even an explicit nil
### GetNTrials

`func (o *RareDiseaseSponsorsRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *RareDiseaseSponsorsRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *RareDiseaseSponsorsRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *RareDiseaseSponsorsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *RareDiseaseSponsorsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *RareDiseaseSponsorsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNTrialsLead

`func (o *RareDiseaseSponsorsRow) GetNTrialsLead() int64`

GetNTrialsLead returns the NTrialsLead field if non-nil, zero value otherwise.

### GetNTrialsLeadOk

`func (o *RareDiseaseSponsorsRow) GetNTrialsLeadOk() (*int64, bool)`

GetNTrialsLeadOk returns a tuple with the NTrialsLead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsLead

`func (o *RareDiseaseSponsorsRow) SetNTrialsLead(v int64)`

SetNTrialsLead sets NTrialsLead field to given value.

### HasNTrialsLead

`func (o *RareDiseaseSponsorsRow) HasNTrialsLead() bool`

HasNTrialsLead returns a boolean if a field has been set.

### SetNTrialsLeadNil

`func (o *RareDiseaseSponsorsRow) SetNTrialsLeadNil(b bool)`

 SetNTrialsLeadNil sets the value for NTrialsLead to be an explicit nil

### UnsetNTrialsLead
`func (o *RareDiseaseSponsorsRow) UnsetNTrialsLead()`

UnsetNTrialsLead ensures that no value is present for NTrialsLead, not even an explicit nil
### GetNTrialsCollaborator

`func (o *RareDiseaseSponsorsRow) GetNTrialsCollaborator() int64`

GetNTrialsCollaborator returns the NTrialsCollaborator field if non-nil, zero value otherwise.

### GetNTrialsCollaboratorOk

`func (o *RareDiseaseSponsorsRow) GetNTrialsCollaboratorOk() (*int64, bool)`

GetNTrialsCollaboratorOk returns a tuple with the NTrialsCollaborator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsCollaborator

`func (o *RareDiseaseSponsorsRow) SetNTrialsCollaborator(v int64)`

SetNTrialsCollaborator sets NTrialsCollaborator field to given value.

### HasNTrialsCollaborator

`func (o *RareDiseaseSponsorsRow) HasNTrialsCollaborator() bool`

HasNTrialsCollaborator returns a boolean if a field has been set.

### SetNTrialsCollaboratorNil

`func (o *RareDiseaseSponsorsRow) SetNTrialsCollaboratorNil(b bool)`

 SetNTrialsCollaboratorNil sets the value for NTrialsCollaborator to be an explicit nil

### UnsetNTrialsCollaborator
`func (o *RareDiseaseSponsorsRow) UnsetNTrialsCollaborator()`

UnsetNTrialsCollaborator ensures that no value is present for NTrialsCollaborator, not even an explicit nil
### GetNCompleted

`func (o *RareDiseaseSponsorsRow) GetNCompleted() int64`

GetNCompleted returns the NCompleted field if non-nil, zero value otherwise.

### GetNCompletedOk

`func (o *RareDiseaseSponsorsRow) GetNCompletedOk() (*int64, bool)`

GetNCompletedOk returns a tuple with the NCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCompleted

`func (o *RareDiseaseSponsorsRow) SetNCompleted(v int64)`

SetNCompleted sets NCompleted field to given value.

### HasNCompleted

`func (o *RareDiseaseSponsorsRow) HasNCompleted() bool`

HasNCompleted returns a boolean if a field has been set.

### SetNCompletedNil

`func (o *RareDiseaseSponsorsRow) SetNCompletedNil(b bool)`

 SetNCompletedNil sets the value for NCompleted to be an explicit nil

### UnsetNCompleted
`func (o *RareDiseaseSponsorsRow) UnsetNCompleted()`

UnsetNCompleted ensures that no value is present for NCompleted, not even an explicit nil
### GetNTerminated

`func (o *RareDiseaseSponsorsRow) GetNTerminated() int64`

GetNTerminated returns the NTerminated field if non-nil, zero value otherwise.

### GetNTerminatedOk

`func (o *RareDiseaseSponsorsRow) GetNTerminatedOk() (*int64, bool)`

GetNTerminatedOk returns a tuple with the NTerminated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTerminated

`func (o *RareDiseaseSponsorsRow) SetNTerminated(v int64)`

SetNTerminated sets NTerminated field to given value.

### HasNTerminated

`func (o *RareDiseaseSponsorsRow) HasNTerminated() bool`

HasNTerminated returns a boolean if a field has been set.

### SetNTerminatedNil

`func (o *RareDiseaseSponsorsRow) SetNTerminatedNil(b bool)`

 SetNTerminatedNil sets the value for NTerminated to be an explicit nil

### UnsetNTerminated
`func (o *RareDiseaseSponsorsRow) UnsetNTerminated()`

UnsetNTerminated ensures that no value is present for NTerminated, not even an explicit nil
### GetNWithdrawn

`func (o *RareDiseaseSponsorsRow) GetNWithdrawn() int64`

GetNWithdrawn returns the NWithdrawn field if non-nil, zero value otherwise.

### GetNWithdrawnOk

`func (o *RareDiseaseSponsorsRow) GetNWithdrawnOk() (*int64, bool)`

GetNWithdrawnOk returns a tuple with the NWithdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNWithdrawn

`func (o *RareDiseaseSponsorsRow) SetNWithdrawn(v int64)`

SetNWithdrawn sets NWithdrawn field to given value.

### HasNWithdrawn

`func (o *RareDiseaseSponsorsRow) HasNWithdrawn() bool`

HasNWithdrawn returns a boolean if a field has been set.

### SetNWithdrawnNil

`func (o *RareDiseaseSponsorsRow) SetNWithdrawnNil(b bool)`

 SetNWithdrawnNil sets the value for NWithdrawn to be an explicit nil

### UnsetNWithdrawn
`func (o *RareDiseaseSponsorsRow) UnsetNWithdrawn()`

UnsetNWithdrawn ensures that no value is present for NWithdrawn, not even an explicit nil
### GetNSuspended

`func (o *RareDiseaseSponsorsRow) GetNSuspended() int64`

GetNSuspended returns the NSuspended field if non-nil, zero value otherwise.

### GetNSuspendedOk

`func (o *RareDiseaseSponsorsRow) GetNSuspendedOk() (*int64, bool)`

GetNSuspendedOk returns a tuple with the NSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSuspended

`func (o *RareDiseaseSponsorsRow) SetNSuspended(v int64)`

SetNSuspended sets NSuspended field to given value.

### HasNSuspended

`func (o *RareDiseaseSponsorsRow) HasNSuspended() bool`

HasNSuspended returns a boolean if a field has been set.

### SetNSuspendedNil

`func (o *RareDiseaseSponsorsRow) SetNSuspendedNil(b bool)`

 SetNSuspendedNil sets the value for NSuspended to be an explicit nil

### UnsetNSuspended
`func (o *RareDiseaseSponsorsRow) UnsetNSuspended()`

UnsetNSuspended ensures that no value is present for NSuspended, not even an explicit nil
### GetNOpen

`func (o *RareDiseaseSponsorsRow) GetNOpen() int64`

GetNOpen returns the NOpen field if non-nil, zero value otherwise.

### GetNOpenOk

`func (o *RareDiseaseSponsorsRow) GetNOpenOk() (*int64, bool)`

GetNOpenOk returns a tuple with the NOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOpen

`func (o *RareDiseaseSponsorsRow) SetNOpen(v int64)`

SetNOpen sets NOpen field to given value.

### HasNOpen

`func (o *RareDiseaseSponsorsRow) HasNOpen() bool`

HasNOpen returns a boolean if a field has been set.

### SetNOpenNil

`func (o *RareDiseaseSponsorsRow) SetNOpenNil(b bool)`

 SetNOpenNil sets the value for NOpen to be an explicit nil

### UnsetNOpen
`func (o *RareDiseaseSponsorsRow) UnsetNOpen()`

UnsetNOpen ensures that no value is present for NOpen, not even an explicit nil
### GetNOngoing

`func (o *RareDiseaseSponsorsRow) GetNOngoing() int64`

GetNOngoing returns the NOngoing field if non-nil, zero value otherwise.

### GetNOngoingOk

`func (o *RareDiseaseSponsorsRow) GetNOngoingOk() (*int64, bool)`

GetNOngoingOk returns a tuple with the NOngoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOngoing

`func (o *RareDiseaseSponsorsRow) SetNOngoing(v int64)`

SetNOngoing sets NOngoing field to given value.

### HasNOngoing

`func (o *RareDiseaseSponsorsRow) HasNOngoing() bool`

HasNOngoing returns a boolean if a field has been set.

### SetNOngoingNil

`func (o *RareDiseaseSponsorsRow) SetNOngoingNil(b bool)`

 SetNOngoingNil sets the value for NOngoing to be an explicit nil

### UnsetNOngoing
`func (o *RareDiseaseSponsorsRow) UnsetNOngoing()`

UnsetNOngoing ensures that no value is present for NOngoing, not even an explicit nil
### GetNOther

`func (o *RareDiseaseSponsorsRow) GetNOther() int64`

GetNOther returns the NOther field if non-nil, zero value otherwise.

### GetNOtherOk

`func (o *RareDiseaseSponsorsRow) GetNOtherOk() (*int64, bool)`

GetNOtherOk returns a tuple with the NOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOther

`func (o *RareDiseaseSponsorsRow) SetNOther(v int64)`

SetNOther sets NOther field to given value.

### HasNOther

`func (o *RareDiseaseSponsorsRow) HasNOther() bool`

HasNOther returns a boolean if a field has been set.

### SetNOtherNil

`func (o *RareDiseaseSponsorsRow) SetNOtherNil(b bool)`

 SetNOtherNil sets the value for NOther to be an explicit nil

### UnsetNOther
`func (o *RareDiseaseSponsorsRow) UnsetNOther()`

UnsetNOther ensures that no value is present for NOther, not even an explicit nil
### GetNUnverified3y

`func (o *RareDiseaseSponsorsRow) GetNUnverified3y() int64`

GetNUnverified3y returns the NUnverified3y field if non-nil, zero value otherwise.

### GetNUnverified3yOk

`func (o *RareDiseaseSponsorsRow) GetNUnverified3yOk() (*int64, bool)`

GetNUnverified3yOk returns a tuple with the NUnverified3y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNUnverified3y

`func (o *RareDiseaseSponsorsRow) SetNUnverified3y(v int64)`

SetNUnverified3y sets NUnverified3y field to given value.

### HasNUnverified3y

`func (o *RareDiseaseSponsorsRow) HasNUnverified3y() bool`

HasNUnverified3y returns a boolean if a field has been set.

### SetNUnverified3yNil

`func (o *RareDiseaseSponsorsRow) SetNUnverified3yNil(b bool)`

 SetNUnverified3yNil sets the value for NUnverified3y to be an explicit nil

### UnsetNUnverified3y
`func (o *RareDiseaseSponsorsRow) UnsetNUnverified3y()`

UnsetNUnverified3y ensures that no value is present for NUnverified3y, not even an explicit nil
### GetFirstTrialStart

`func (o *RareDiseaseSponsorsRow) GetFirstTrialStart() string`

GetFirstTrialStart returns the FirstTrialStart field if non-nil, zero value otherwise.

### GetFirstTrialStartOk

`func (o *RareDiseaseSponsorsRow) GetFirstTrialStartOk() (*string, bool)`

GetFirstTrialStartOk returns a tuple with the FirstTrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTrialStart

`func (o *RareDiseaseSponsorsRow) SetFirstTrialStart(v string)`

SetFirstTrialStart sets FirstTrialStart field to given value.

### HasFirstTrialStart

`func (o *RareDiseaseSponsorsRow) HasFirstTrialStart() bool`

HasFirstTrialStart returns a boolean if a field has been set.

### SetFirstTrialStartNil

`func (o *RareDiseaseSponsorsRow) SetFirstTrialStartNil(b bool)`

 SetFirstTrialStartNil sets the value for FirstTrialStart to be an explicit nil

### UnsetFirstTrialStart
`func (o *RareDiseaseSponsorsRow) UnsetFirstTrialStart()`

UnsetFirstTrialStart ensures that no value is present for FirstTrialStart, not even an explicit nil
### GetLastTrialStart

`func (o *RareDiseaseSponsorsRow) GetLastTrialStart() string`

GetLastTrialStart returns the LastTrialStart field if non-nil, zero value otherwise.

### GetLastTrialStartOk

`func (o *RareDiseaseSponsorsRow) GetLastTrialStartOk() (*string, bool)`

GetLastTrialStartOk returns a tuple with the LastTrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrialStart

`func (o *RareDiseaseSponsorsRow) SetLastTrialStart(v string)`

SetLastTrialStart sets LastTrialStart field to given value.

### HasLastTrialStart

`func (o *RareDiseaseSponsorsRow) HasLastTrialStart() bool`

HasLastTrialStart returns a boolean if a field has been set.

### SetLastTrialStartNil

`func (o *RareDiseaseSponsorsRow) SetLastTrialStartNil(b bool)`

 SetLastTrialStartNil sets the value for LastTrialStart to be an explicit nil

### UnsetLastTrialStart
`func (o *RareDiseaseSponsorsRow) UnsetLastTrialStart()`

UnsetLastTrialStart ensures that no value is present for LastTrialStart, not even an explicit nil
### GetFirstTrialStartYear

`func (o *RareDiseaseSponsorsRow) GetFirstTrialStartYear() int32`

GetFirstTrialStartYear returns the FirstTrialStartYear field if non-nil, zero value otherwise.

### GetFirstTrialStartYearOk

`func (o *RareDiseaseSponsorsRow) GetFirstTrialStartYearOk() (*int32, bool)`

GetFirstTrialStartYearOk returns a tuple with the FirstTrialStartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTrialStartYear

`func (o *RareDiseaseSponsorsRow) SetFirstTrialStartYear(v int32)`

SetFirstTrialStartYear sets FirstTrialStartYear field to given value.

### HasFirstTrialStartYear

`func (o *RareDiseaseSponsorsRow) HasFirstTrialStartYear() bool`

HasFirstTrialStartYear returns a boolean if a field has been set.

### SetFirstTrialStartYearNil

`func (o *RareDiseaseSponsorsRow) SetFirstTrialStartYearNil(b bool)`

 SetFirstTrialStartYearNil sets the value for FirstTrialStartYear to be an explicit nil

### UnsetFirstTrialStartYear
`func (o *RareDiseaseSponsorsRow) UnsetFirstTrialStartYear()`

UnsetFirstTrialStartYear ensures that no value is present for FirstTrialStartYear, not even an explicit nil
### GetLastTrialStartYear

`func (o *RareDiseaseSponsorsRow) GetLastTrialStartYear() int32`

GetLastTrialStartYear returns the LastTrialStartYear field if non-nil, zero value otherwise.

### GetLastTrialStartYearOk

`func (o *RareDiseaseSponsorsRow) GetLastTrialStartYearOk() (*int32, bool)`

GetLastTrialStartYearOk returns a tuple with the LastTrialStartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrialStartYear

`func (o *RareDiseaseSponsorsRow) SetLastTrialStartYear(v int32)`

SetLastTrialStartYear sets LastTrialStartYear field to given value.

### HasLastTrialStartYear

`func (o *RareDiseaseSponsorsRow) HasLastTrialStartYear() bool`

HasLastTrialStartYear returns a boolean if a field has been set.

### SetLastTrialStartYearNil

`func (o *RareDiseaseSponsorsRow) SetLastTrialStartYearNil(b bool)`

 SetLastTrialStartYearNil sets the value for LastTrialStartYear to be an explicit nil

### UnsetLastTrialStartYear
`func (o *RareDiseaseSponsorsRow) UnsetLastTrialStartYear()`

UnsetLastTrialStartYear ensures that no value is present for LastTrialStartYear, not even an explicit nil
### GetLastUpdatePostedDate

`func (o *RareDiseaseSponsorsRow) GetLastUpdatePostedDate() string`

GetLastUpdatePostedDate returns the LastUpdatePostedDate field if non-nil, zero value otherwise.

### GetLastUpdatePostedDateOk

`func (o *RareDiseaseSponsorsRow) GetLastUpdatePostedDateOk() (*string, bool)`

GetLastUpdatePostedDateOk returns a tuple with the LastUpdatePostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatePostedDate

`func (o *RareDiseaseSponsorsRow) SetLastUpdatePostedDate(v string)`

SetLastUpdatePostedDate sets LastUpdatePostedDate field to given value.

### HasLastUpdatePostedDate

`func (o *RareDiseaseSponsorsRow) HasLastUpdatePostedDate() bool`

HasLastUpdatePostedDate returns a boolean if a field has been set.

### SetLastUpdatePostedDateNil

`func (o *RareDiseaseSponsorsRow) SetLastUpdatePostedDateNil(b bool)`

 SetLastUpdatePostedDateNil sets the value for LastUpdatePostedDate to be an explicit nil

### UnsetLastUpdatePostedDate
`func (o *RareDiseaseSponsorsRow) UnsetLastUpdatePostedDate()`

UnsetLastUpdatePostedDate ensures that no value is present for LastUpdatePostedDate, not even an explicit nil
### GetIsDormant

`func (o *RareDiseaseSponsorsRow) GetIsDormant() bool`

GetIsDormant returns the IsDormant field if non-nil, zero value otherwise.

### GetIsDormantOk

`func (o *RareDiseaseSponsorsRow) GetIsDormantOk() (*bool, bool)`

GetIsDormantOk returns a tuple with the IsDormant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDormant

`func (o *RareDiseaseSponsorsRow) SetIsDormant(v bool)`

SetIsDormant sets IsDormant field to given value.

### HasIsDormant

`func (o *RareDiseaseSponsorsRow) HasIsDormant() bool`

HasIsDormant returns a boolean if a field has been set.

### SetIsDormantNil

`func (o *RareDiseaseSponsorsRow) SetIsDormantNil(b bool)`

 SetIsDormantNil sets the value for IsDormant to be an explicit nil

### UnsetIsDormant
`func (o *RareDiseaseSponsorsRow) UnsetIsDormant()`

UnsetIsDormant ensures that no value is present for IsDormant, not even an explicit nil
### GetMappingArm

`func (o *RareDiseaseSponsorsRow) GetMappingArm() string`

GetMappingArm returns the MappingArm field if non-nil, zero value otherwise.

### GetMappingArmOk

`func (o *RareDiseaseSponsorsRow) GetMappingArmOk() (*string, bool)`

GetMappingArmOk returns a tuple with the MappingArm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingArm

`func (o *RareDiseaseSponsorsRow) SetMappingArm(v string)`

SetMappingArm sets MappingArm field to given value.

### HasMappingArm

`func (o *RareDiseaseSponsorsRow) HasMappingArm() bool`

HasMappingArm returns a boolean if a field has been set.

### SetMappingArmNil

`func (o *RareDiseaseSponsorsRow) SetMappingArmNil(b bool)`

 SetMappingArmNil sets the value for MappingArm to be an explicit nil

### UnsetMappingArm
`func (o *RareDiseaseSponsorsRow) UnsetMappingArm()`

UnsetMappingArm ensures that no value is present for MappingArm, not even an explicit nil
### GetEdgeScope

`func (o *RareDiseaseSponsorsRow) GetEdgeScope() string`

GetEdgeScope returns the EdgeScope field if non-nil, zero value otherwise.

### GetEdgeScopeOk

`func (o *RareDiseaseSponsorsRow) GetEdgeScopeOk() (*string, bool)`

GetEdgeScopeOk returns a tuple with the EdgeScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeScope

`func (o *RareDiseaseSponsorsRow) SetEdgeScope(v string)`

SetEdgeScope sets EdgeScope field to given value.

### HasEdgeScope

`func (o *RareDiseaseSponsorsRow) HasEdgeScope() bool`

HasEdgeScope returns a boolean if a field has been set.

### SetEdgeScopeNil

`func (o *RareDiseaseSponsorsRow) SetEdgeScopeNil(b bool)`

 SetEdgeScopeNil sets the value for EdgeScope to be an explicit nil

### UnsetEdgeScope
`func (o *RareDiseaseSponsorsRow) UnsetEdgeScope()`

UnsetEdgeScope ensures that no value is present for EdgeScope, not even an explicit nil
### GetFreshnessAsOf

`func (o *RareDiseaseSponsorsRow) GetFreshnessAsOf() string`

GetFreshnessAsOf returns the FreshnessAsOf field if non-nil, zero value otherwise.

### GetFreshnessAsOfOk

`func (o *RareDiseaseSponsorsRow) GetFreshnessAsOfOk() (*string, bool)`

GetFreshnessAsOfOk returns a tuple with the FreshnessAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshnessAsOf

`func (o *RareDiseaseSponsorsRow) SetFreshnessAsOf(v string)`

SetFreshnessAsOf sets FreshnessAsOf field to given value.

### HasFreshnessAsOf

`func (o *RareDiseaseSponsorsRow) HasFreshnessAsOf() bool`

HasFreshnessAsOf returns a boolean if a field has been set.

### SetFreshnessAsOfNil

`func (o *RareDiseaseSponsorsRow) SetFreshnessAsOfNil(b bool)`

 SetFreshnessAsOfNil sets the value for FreshnessAsOf to be an explicit nil

### UnsetFreshnessAsOf
`func (o *RareDiseaseSponsorsRow) UnsetFreshnessAsOf()`

UnsetFreshnessAsOf ensures that no value is present for FreshnessAsOf, not even an explicit nil
### GetLicenseName

`func (o *RareDiseaseSponsorsRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *RareDiseaseSponsorsRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *RareDiseaseSponsorsRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *RareDiseaseSponsorsRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *RareDiseaseSponsorsRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *RareDiseaseSponsorsRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *RareDiseaseSponsorsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RareDiseaseSponsorsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RareDiseaseSponsorsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RareDiseaseSponsorsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RareDiseaseSponsorsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RareDiseaseSponsorsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


