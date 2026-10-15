# ApiV1TrialsRecordsGet200ResponseTrialsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanonicalTrialId** | **string** |  | 
**NctId** | **string** |  | 
**SourceRegistry** | **string** |  | 
**Registry** | **string** |  | 
**BriefTitle** | **string** |  | 
**OfficialTitle** | **NullableString** |  | 
**StudyType** | **string** |  | 
**OverallStatus** | **string** |  | 
**OverallStatusSummary** | **string** |  | 
**Phase** | **NullableString** |  | 
**PhaseByStatus** | **string** |  | 
**Enrollment** | **NullableInt32** |  | 
**EnrollmentType** | **NullableString** |  | 
**EnrollmentFlagReason** | **interface{}** |  | 
**StartDate** | **string** |  | 
**StartYear** | **int32** |  | 
**PrimaryCompletionDate** | **NullableString** |  | 
**PrimaryCompletionDateType** | **NullableString** |  | 
**CompletionDate** | **NullableString** |  | 
**CompletionDateType** | **NullableString** |  | 
**StudyFirstPostedDate** | **string** |  | 
**LastUpdatePostedDate** | **string** |  | 
**WhyStopped** | **NullableString** |  | 
**NumberOfArms** | **interface{}** |  | 
**NumberOfFacilities** | **int32** |  | 
**WereResultsReported** | **bool** |  | 
**HasUsFacility** | **NullableBool** |  | 
**Gender** | **string** |  | 
**MinimumAge** | **NullableString** |  | 
**MaximumAge** | **NullableString** |  | 
**HealthyVolunteers** | **NullableBool** |  | 
**Criteria** | **string** |  | 
**LeadSponsorName** | **string** |  | 
**LeadSponsorId** | **string** |  | 
**Collaborators** | **[]string** |  | 
**ParticipatingSponsors** | **[]string** |  | 
**Conditions** | **[]string** |  | 
**ConditionMesh** | **[]string** |  | 
**Interventions** | **[]string** |  | 
**InterventionTypes** | **[]string** |  | 
**InterventionMesh** | **[]string** |  | 
**FacilityCount** | **NullableInt32** |  | 
**FacilityCountries** | **[]string** |  | 
**FacilityCountryCount** | **NullableInt32** |  | 
**EnrollmentIsPlausible** | **bool** |  | 
**PhaseLabel** | **string** |  | 
**Extracted** | [**NullableApiV1TrialsRecordsGet200ResponseTrialsInnerExtracted**](ApiV1TrialsRecordsGet200ResponseTrialsInnerExtracted.md) |  | 
**LinkedDrugs** | Pointer to [**[]ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner**](ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner.md) |  | [optional] 
**LinkedDrugsBuilt** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiV1TrialsRecordsGet200ResponseTrialsInner

`func NewApiV1TrialsRecordsGet200ResponseTrialsInner(canonicalTrialId string, nctId string, sourceRegistry string, registry string, briefTitle string, officialTitle NullableString, studyType string, overallStatus string, overallStatusSummary string, phase NullableString, phaseByStatus string, enrollment NullableInt32, enrollmentType NullableString, enrollmentFlagReason interface{}, startDate string, startYear int32, primaryCompletionDate NullableString, primaryCompletionDateType NullableString, completionDate NullableString, completionDateType NullableString, studyFirstPostedDate string, lastUpdatePostedDate string, whyStopped NullableString, numberOfArms interface{}, numberOfFacilities int32, wereResultsReported bool, hasUsFacility NullableBool, gender string, minimumAge NullableString, maximumAge NullableString, healthyVolunteers NullableBool, criteria string, leadSponsorName string, leadSponsorId string, collaborators []string, participatingSponsors []string, conditions []string, conditionMesh []string, interventions []string, interventionTypes []string, interventionMesh []string, facilityCount NullableInt32, facilityCountries []string, facilityCountryCount NullableInt32, enrollmentIsPlausible bool, phaseLabel string, extracted NullableApiV1TrialsRecordsGet200ResponseTrialsInnerExtracted, ) *ApiV1TrialsRecordsGet200ResponseTrialsInner`

NewApiV1TrialsRecordsGet200ResponseTrialsInner instantiates a new ApiV1TrialsRecordsGet200ResponseTrialsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1TrialsRecordsGet200ResponseTrialsInnerWithDefaults

`func NewApiV1TrialsRecordsGet200ResponseTrialsInnerWithDefaults() *ApiV1TrialsRecordsGet200ResponseTrialsInner`

NewApiV1TrialsRecordsGet200ResponseTrialsInnerWithDefaults instantiates a new ApiV1TrialsRecordsGet200ResponseTrialsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonicalTrialId

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetCanonicalTrialId() string`

GetCanonicalTrialId returns the CanonicalTrialId field if non-nil, zero value otherwise.

### GetCanonicalTrialIdOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetCanonicalTrialIdOk() (*string, bool)`

GetCanonicalTrialIdOk returns a tuple with the CanonicalTrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalTrialId

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetCanonicalTrialId(v string)`

SetCanonicalTrialId sets CanonicalTrialId field to given value.


### GetNctId

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetNctId() string`

GetNctId returns the NctId field if non-nil, zero value otherwise.

### GetNctIdOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetNctIdOk() (*string, bool)`

GetNctIdOk returns a tuple with the NctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNctId

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetNctId(v string)`

SetNctId sets NctId field to given value.


### GetSourceRegistry

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetSourceRegistry() string`

GetSourceRegistry returns the SourceRegistry field if non-nil, zero value otherwise.

### GetSourceRegistryOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetSourceRegistryOk() (*string, bool)`

GetSourceRegistryOk returns a tuple with the SourceRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRegistry

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetSourceRegistry(v string)`

SetSourceRegistry sets SourceRegistry field to given value.


### GetRegistry

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### GetBriefTitle

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetBriefTitle() string`

GetBriefTitle returns the BriefTitle field if non-nil, zero value otherwise.

### GetBriefTitleOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetBriefTitleOk() (*string, bool)`

GetBriefTitleOk returns a tuple with the BriefTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBriefTitle

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetBriefTitle(v string)`

SetBriefTitle sets BriefTitle field to given value.


### GetOfficialTitle

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetOfficialTitle() string`

GetOfficialTitle returns the OfficialTitle field if non-nil, zero value otherwise.

### GetOfficialTitleOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetOfficialTitleOk() (*string, bool)`

GetOfficialTitleOk returns a tuple with the OfficialTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialTitle

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetOfficialTitle(v string)`

SetOfficialTitle sets OfficialTitle field to given value.


### SetOfficialTitleNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetOfficialTitleNil(b bool)`

 SetOfficialTitleNil sets the value for OfficialTitle to be an explicit nil

### UnsetOfficialTitle
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetOfficialTitle()`

UnsetOfficialTitle ensures that no value is present for OfficialTitle, not even an explicit nil
### GetStudyType

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetStudyType() string`

GetStudyType returns the StudyType field if non-nil, zero value otherwise.

### GetStudyTypeOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetStudyTypeOk() (*string, bool)`

GetStudyTypeOk returns a tuple with the StudyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyType

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetStudyType(v string)`

SetStudyType sets StudyType field to given value.


### GetOverallStatus

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.


### GetOverallStatusSummary

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetOverallStatusSummary() string`

GetOverallStatusSummary returns the OverallStatusSummary field if non-nil, zero value otherwise.

### GetOverallStatusSummaryOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetOverallStatusSummaryOk() (*string, bool)`

GetOverallStatusSummaryOk returns a tuple with the OverallStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatusSummary

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetOverallStatusSummary(v string)`

SetOverallStatusSummary sets OverallStatusSummary field to given value.


### GetPhase

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetPhase(v string)`

SetPhase sets Phase field to given value.


### SetPhaseNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetPhaseByStatus

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetPhaseByStatus() string`

GetPhaseByStatus returns the PhaseByStatus field if non-nil, zero value otherwise.

### GetPhaseByStatusOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetPhaseByStatusOk() (*string, bool)`

GetPhaseByStatusOk returns a tuple with the PhaseByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseByStatus

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetPhaseByStatus(v string)`

SetPhaseByStatus sets PhaseByStatus field to given value.


### GetEnrollment

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetEnrollment() int32`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetEnrollmentOk() (*int32, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetEnrollment(v int32)`

SetEnrollment sets Enrollment field to given value.


### SetEnrollmentNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetEnrollmentNil(b bool)`

 SetEnrollmentNil sets the value for Enrollment to be an explicit nil

### UnsetEnrollment
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetEnrollment()`

UnsetEnrollment ensures that no value is present for Enrollment, not even an explicit nil
### GetEnrollmentType

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetEnrollmentType() string`

GetEnrollmentType returns the EnrollmentType field if non-nil, zero value otherwise.

### GetEnrollmentTypeOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetEnrollmentTypeOk() (*string, bool)`

GetEnrollmentTypeOk returns a tuple with the EnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentType

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetEnrollmentType(v string)`

SetEnrollmentType sets EnrollmentType field to given value.


### SetEnrollmentTypeNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetEnrollmentTypeNil(b bool)`

 SetEnrollmentTypeNil sets the value for EnrollmentType to be an explicit nil

### UnsetEnrollmentType
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetEnrollmentType()`

UnsetEnrollmentType ensures that no value is present for EnrollmentType, not even an explicit nil
### GetEnrollmentFlagReason

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetEnrollmentFlagReason() interface{}`

GetEnrollmentFlagReason returns the EnrollmentFlagReason field if non-nil, zero value otherwise.

### GetEnrollmentFlagReasonOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetEnrollmentFlagReasonOk() (*interface{}, bool)`

GetEnrollmentFlagReasonOk returns a tuple with the EnrollmentFlagReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlagReason

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetEnrollmentFlagReason(v interface{})`

SetEnrollmentFlagReason sets EnrollmentFlagReason field to given value.


### SetEnrollmentFlagReasonNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetEnrollmentFlagReasonNil(b bool)`

 SetEnrollmentFlagReasonNil sets the value for EnrollmentFlagReason to be an explicit nil

### UnsetEnrollmentFlagReason
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetEnrollmentFlagReason()`

UnsetEnrollmentFlagReason ensures that no value is present for EnrollmentFlagReason, not even an explicit nil
### GetStartDate

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetStartYear

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetStartYear() int32`

GetStartYear returns the StartYear field if non-nil, zero value otherwise.

### GetStartYearOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetStartYearOk() (*int32, bool)`

GetStartYearOk returns a tuple with the StartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartYear

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetStartYear(v int32)`

SetStartYear sets StartYear field to given value.


### GetPrimaryCompletionDate

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetPrimaryCompletionDate() string`

GetPrimaryCompletionDate returns the PrimaryCompletionDate field if non-nil, zero value otherwise.

### GetPrimaryCompletionDateOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetPrimaryCompletionDateOk() (*string, bool)`

GetPrimaryCompletionDateOk returns a tuple with the PrimaryCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCompletionDate

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetPrimaryCompletionDate(v string)`

SetPrimaryCompletionDate sets PrimaryCompletionDate field to given value.


### SetPrimaryCompletionDateNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetPrimaryCompletionDateNil(b bool)`

 SetPrimaryCompletionDateNil sets the value for PrimaryCompletionDate to be an explicit nil

### UnsetPrimaryCompletionDate
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetPrimaryCompletionDate()`

UnsetPrimaryCompletionDate ensures that no value is present for PrimaryCompletionDate, not even an explicit nil
### GetPrimaryCompletionDateType

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetPrimaryCompletionDateType() string`

GetPrimaryCompletionDateType returns the PrimaryCompletionDateType field if non-nil, zero value otherwise.

### GetPrimaryCompletionDateTypeOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetPrimaryCompletionDateTypeOk() (*string, bool)`

GetPrimaryCompletionDateTypeOk returns a tuple with the PrimaryCompletionDateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCompletionDateType

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetPrimaryCompletionDateType(v string)`

SetPrimaryCompletionDateType sets PrimaryCompletionDateType field to given value.


### SetPrimaryCompletionDateTypeNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetPrimaryCompletionDateTypeNil(b bool)`

 SetPrimaryCompletionDateTypeNil sets the value for PrimaryCompletionDateType to be an explicit nil

### UnsetPrimaryCompletionDateType
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetPrimaryCompletionDateType()`

UnsetPrimaryCompletionDateType ensures that no value is present for PrimaryCompletionDateType, not even an explicit nil
### GetCompletionDate

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetCompletionDate() string`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetCompletionDateOk() (*string, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetCompletionDate(v string)`

SetCompletionDate sets CompletionDate field to given value.


### SetCompletionDateNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetCompletionDateNil(b bool)`

 SetCompletionDateNil sets the value for CompletionDate to be an explicit nil

### UnsetCompletionDate
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetCompletionDate()`

UnsetCompletionDate ensures that no value is present for CompletionDate, not even an explicit nil
### GetCompletionDateType

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetCompletionDateType() string`

GetCompletionDateType returns the CompletionDateType field if non-nil, zero value otherwise.

### GetCompletionDateTypeOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetCompletionDateTypeOk() (*string, bool)`

GetCompletionDateTypeOk returns a tuple with the CompletionDateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateType

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetCompletionDateType(v string)`

SetCompletionDateType sets CompletionDateType field to given value.


### SetCompletionDateTypeNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetCompletionDateTypeNil(b bool)`

 SetCompletionDateTypeNil sets the value for CompletionDateType to be an explicit nil

### UnsetCompletionDateType
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetCompletionDateType()`

UnsetCompletionDateType ensures that no value is present for CompletionDateType, not even an explicit nil
### GetStudyFirstPostedDate

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetStudyFirstPostedDate() string`

GetStudyFirstPostedDate returns the StudyFirstPostedDate field if non-nil, zero value otherwise.

### GetStudyFirstPostedDateOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetStudyFirstPostedDateOk() (*string, bool)`

GetStudyFirstPostedDateOk returns a tuple with the StudyFirstPostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyFirstPostedDate

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetStudyFirstPostedDate(v string)`

SetStudyFirstPostedDate sets StudyFirstPostedDate field to given value.


### GetLastUpdatePostedDate

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetLastUpdatePostedDate() string`

GetLastUpdatePostedDate returns the LastUpdatePostedDate field if non-nil, zero value otherwise.

### GetLastUpdatePostedDateOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetLastUpdatePostedDateOk() (*string, bool)`

GetLastUpdatePostedDateOk returns a tuple with the LastUpdatePostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatePostedDate

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetLastUpdatePostedDate(v string)`

SetLastUpdatePostedDate sets LastUpdatePostedDate field to given value.


### GetWhyStopped

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetWhyStopped() string`

GetWhyStopped returns the WhyStopped field if non-nil, zero value otherwise.

### GetWhyStoppedOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetWhyStoppedOk() (*string, bool)`

GetWhyStoppedOk returns a tuple with the WhyStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhyStopped

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetWhyStopped(v string)`

SetWhyStopped sets WhyStopped field to given value.


### SetWhyStoppedNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetWhyStoppedNil(b bool)`

 SetWhyStoppedNil sets the value for WhyStopped to be an explicit nil

### UnsetWhyStopped
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetWhyStopped()`

UnsetWhyStopped ensures that no value is present for WhyStopped, not even an explicit nil
### GetNumberOfArms

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetNumberOfArms() interface{}`

GetNumberOfArms returns the NumberOfArms field if non-nil, zero value otherwise.

### GetNumberOfArmsOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetNumberOfArmsOk() (*interface{}, bool)`

GetNumberOfArmsOk returns a tuple with the NumberOfArms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfArms

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetNumberOfArms(v interface{})`

SetNumberOfArms sets NumberOfArms field to given value.


### SetNumberOfArmsNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetNumberOfArmsNil(b bool)`

 SetNumberOfArmsNil sets the value for NumberOfArms to be an explicit nil

### UnsetNumberOfArms
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetNumberOfArms()`

UnsetNumberOfArms ensures that no value is present for NumberOfArms, not even an explicit nil
### GetNumberOfFacilities

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetNumberOfFacilities() int32`

GetNumberOfFacilities returns the NumberOfFacilities field if non-nil, zero value otherwise.

### GetNumberOfFacilitiesOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetNumberOfFacilitiesOk() (*int32, bool)`

GetNumberOfFacilitiesOk returns a tuple with the NumberOfFacilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFacilities

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetNumberOfFacilities(v int32)`

SetNumberOfFacilities sets NumberOfFacilities field to given value.


### GetWereResultsReported

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetWereResultsReported() bool`

GetWereResultsReported returns the WereResultsReported field if non-nil, zero value otherwise.

### GetWereResultsReportedOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetWereResultsReportedOk() (*bool, bool)`

GetWereResultsReportedOk returns a tuple with the WereResultsReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWereResultsReported

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetWereResultsReported(v bool)`

SetWereResultsReported sets WereResultsReported field to given value.


### GetHasUsFacility

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetHasUsFacility() bool`

GetHasUsFacility returns the HasUsFacility field if non-nil, zero value otherwise.

### GetHasUsFacilityOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetHasUsFacilityOk() (*bool, bool)`

GetHasUsFacilityOk returns a tuple with the HasUsFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUsFacility

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetHasUsFacility(v bool)`

SetHasUsFacility sets HasUsFacility field to given value.


### SetHasUsFacilityNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetHasUsFacilityNil(b bool)`

 SetHasUsFacilityNil sets the value for HasUsFacility to be an explicit nil

### UnsetHasUsFacility
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetHasUsFacility()`

UnsetHasUsFacility ensures that no value is present for HasUsFacility, not even an explicit nil
### GetGender

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetGender(v string)`

SetGender sets Gender field to given value.


### GetMinimumAge

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetMinimumAge() string`

GetMinimumAge returns the MinimumAge field if non-nil, zero value otherwise.

### GetMinimumAgeOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetMinimumAgeOk() (*string, bool)`

GetMinimumAgeOk returns a tuple with the MinimumAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAge

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetMinimumAge(v string)`

SetMinimumAge sets MinimumAge field to given value.


### SetMinimumAgeNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetMinimumAgeNil(b bool)`

 SetMinimumAgeNil sets the value for MinimumAge to be an explicit nil

### UnsetMinimumAge
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetMinimumAge()`

UnsetMinimumAge ensures that no value is present for MinimumAge, not even an explicit nil
### GetMaximumAge

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetMaximumAge() string`

GetMaximumAge returns the MaximumAge field if non-nil, zero value otherwise.

### GetMaximumAgeOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetMaximumAgeOk() (*string, bool)`

GetMaximumAgeOk returns a tuple with the MaximumAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAge

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetMaximumAge(v string)`

SetMaximumAge sets MaximumAge field to given value.


### SetMaximumAgeNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetMaximumAgeNil(b bool)`

 SetMaximumAgeNil sets the value for MaximumAge to be an explicit nil

### UnsetMaximumAge
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetMaximumAge()`

UnsetMaximumAge ensures that no value is present for MaximumAge, not even an explicit nil
### GetHealthyVolunteers

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetHealthyVolunteers() bool`

GetHealthyVolunteers returns the HealthyVolunteers field if non-nil, zero value otherwise.

### GetHealthyVolunteersOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetHealthyVolunteersOk() (*bool, bool)`

GetHealthyVolunteersOk returns a tuple with the HealthyVolunteers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyVolunteers

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetHealthyVolunteers(v bool)`

SetHealthyVolunteers sets HealthyVolunteers field to given value.


### SetHealthyVolunteersNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetHealthyVolunteersNil(b bool)`

 SetHealthyVolunteersNil sets the value for HealthyVolunteers to be an explicit nil

### UnsetHealthyVolunteers
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetHealthyVolunteers()`

UnsetHealthyVolunteers ensures that no value is present for HealthyVolunteers, not even an explicit nil
### GetCriteria

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.


### GetLeadSponsorName

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetLeadSponsorName() string`

GetLeadSponsorName returns the LeadSponsorName field if non-nil, zero value otherwise.

### GetLeadSponsorNameOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetLeadSponsorNameOk() (*string, bool)`

GetLeadSponsorNameOk returns a tuple with the LeadSponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorName

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetLeadSponsorName(v string)`

SetLeadSponsorName sets LeadSponsorName field to given value.


### GetLeadSponsorId

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetLeadSponsorId() string`

GetLeadSponsorId returns the LeadSponsorId field if non-nil, zero value otherwise.

### GetLeadSponsorIdOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetLeadSponsorIdOk() (*string, bool)`

GetLeadSponsorIdOk returns a tuple with the LeadSponsorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSponsorId

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetLeadSponsorId(v string)`

SetLeadSponsorId sets LeadSponsorId field to given value.


### GetCollaborators

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetCollaborators() []string`

GetCollaborators returns the Collaborators field if non-nil, zero value otherwise.

### GetCollaboratorsOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetCollaboratorsOk() (*[]string, bool)`

GetCollaboratorsOk returns a tuple with the Collaborators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborators

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetCollaborators(v []string)`

SetCollaborators sets Collaborators field to given value.


### GetParticipatingSponsors

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetParticipatingSponsors() []string`

GetParticipatingSponsors returns the ParticipatingSponsors field if non-nil, zero value otherwise.

### GetParticipatingSponsorsOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetParticipatingSponsorsOk() (*[]string, bool)`

GetParticipatingSponsorsOk returns a tuple with the ParticipatingSponsors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipatingSponsors

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetParticipatingSponsors(v []string)`

SetParticipatingSponsors sets ParticipatingSponsors field to given value.


### GetConditions

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetConditions() []string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetConditionsOk() (*[]string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetConditions(v []string)`

SetConditions sets Conditions field to given value.


### GetConditionMesh

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetConditionMesh() []string`

GetConditionMesh returns the ConditionMesh field if non-nil, zero value otherwise.

### GetConditionMeshOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetConditionMeshOk() (*[]string, bool)`

GetConditionMeshOk returns a tuple with the ConditionMesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionMesh

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetConditionMesh(v []string)`

SetConditionMesh sets ConditionMesh field to given value.


### GetInterventions

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetInterventions() []string`

GetInterventions returns the Interventions field if non-nil, zero value otherwise.

### GetInterventionsOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetInterventionsOk() (*[]string, bool)`

GetInterventionsOk returns a tuple with the Interventions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterventions

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetInterventions(v []string)`

SetInterventions sets Interventions field to given value.


### GetInterventionTypes

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetInterventionTypes() []string`

GetInterventionTypes returns the InterventionTypes field if non-nil, zero value otherwise.

### GetInterventionTypesOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetInterventionTypesOk() (*[]string, bool)`

GetInterventionTypesOk returns a tuple with the InterventionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterventionTypes

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetInterventionTypes(v []string)`

SetInterventionTypes sets InterventionTypes field to given value.


### GetInterventionMesh

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetInterventionMesh() []string`

GetInterventionMesh returns the InterventionMesh field if non-nil, zero value otherwise.

### GetInterventionMeshOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetInterventionMeshOk() (*[]string, bool)`

GetInterventionMeshOk returns a tuple with the InterventionMesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterventionMesh

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetInterventionMesh(v []string)`

SetInterventionMesh sets InterventionMesh field to given value.


### SetInterventionMeshNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetInterventionMeshNil(b bool)`

 SetInterventionMeshNil sets the value for InterventionMesh to be an explicit nil

### UnsetInterventionMesh
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetInterventionMesh()`

UnsetInterventionMesh ensures that no value is present for InterventionMesh, not even an explicit nil
### GetFacilityCount

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetFacilityCount() int32`

GetFacilityCount returns the FacilityCount field if non-nil, zero value otherwise.

### GetFacilityCountOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetFacilityCountOk() (*int32, bool)`

GetFacilityCountOk returns a tuple with the FacilityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCount

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetFacilityCount(v int32)`

SetFacilityCount sets FacilityCount field to given value.


### SetFacilityCountNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetFacilityCountNil(b bool)`

 SetFacilityCountNil sets the value for FacilityCount to be an explicit nil

### UnsetFacilityCount
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetFacilityCount()`

UnsetFacilityCount ensures that no value is present for FacilityCount, not even an explicit nil
### GetFacilityCountries

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetFacilityCountries() []string`

GetFacilityCountries returns the FacilityCountries field if non-nil, zero value otherwise.

### GetFacilityCountriesOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetFacilityCountriesOk() (*[]string, bool)`

GetFacilityCountriesOk returns a tuple with the FacilityCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCountries

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetFacilityCountries(v []string)`

SetFacilityCountries sets FacilityCountries field to given value.


### SetFacilityCountriesNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetFacilityCountriesNil(b bool)`

 SetFacilityCountriesNil sets the value for FacilityCountries to be an explicit nil

### UnsetFacilityCountries
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetFacilityCountries()`

UnsetFacilityCountries ensures that no value is present for FacilityCountries, not even an explicit nil
### GetFacilityCountryCount

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetFacilityCountryCount() int32`

GetFacilityCountryCount returns the FacilityCountryCount field if non-nil, zero value otherwise.

### GetFacilityCountryCountOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetFacilityCountryCountOk() (*int32, bool)`

GetFacilityCountryCountOk returns a tuple with the FacilityCountryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCountryCount

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetFacilityCountryCount(v int32)`

SetFacilityCountryCount sets FacilityCountryCount field to given value.


### SetFacilityCountryCountNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetFacilityCountryCountNil(b bool)`

 SetFacilityCountryCountNil sets the value for FacilityCountryCount to be an explicit nil

### UnsetFacilityCountryCount
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetFacilityCountryCount()`

UnsetFacilityCountryCount ensures that no value is present for FacilityCountryCount, not even an explicit nil
### GetEnrollmentIsPlausible

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetEnrollmentIsPlausible() bool`

GetEnrollmentIsPlausible returns the EnrollmentIsPlausible field if non-nil, zero value otherwise.

### GetEnrollmentIsPlausibleOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetEnrollmentIsPlausibleOk() (*bool, bool)`

GetEnrollmentIsPlausibleOk returns a tuple with the EnrollmentIsPlausible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentIsPlausible

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetEnrollmentIsPlausible(v bool)`

SetEnrollmentIsPlausible sets EnrollmentIsPlausible field to given value.


### GetPhaseLabel

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetPhaseLabel() string`

GetPhaseLabel returns the PhaseLabel field if non-nil, zero value otherwise.

### GetPhaseLabelOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetPhaseLabelOk() (*string, bool)`

GetPhaseLabelOk returns a tuple with the PhaseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseLabel

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetPhaseLabel(v string)`

SetPhaseLabel sets PhaseLabel field to given value.


### GetExtracted

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetExtracted() ApiV1TrialsRecordsGet200ResponseTrialsInnerExtracted`

GetExtracted returns the Extracted field if non-nil, zero value otherwise.

### GetExtractedOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetExtractedOk() (*ApiV1TrialsRecordsGet200ResponseTrialsInnerExtracted, bool)`

GetExtractedOk returns a tuple with the Extracted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtracted

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetExtracted(v ApiV1TrialsRecordsGet200ResponseTrialsInnerExtracted)`

SetExtracted sets Extracted field to given value.


### SetExtractedNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetExtractedNil(b bool)`

 SetExtractedNil sets the value for Extracted to be an explicit nil

### UnsetExtracted
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) UnsetExtracted()`

UnsetExtracted ensures that no value is present for Extracted, not even an explicit nil
### GetLinkedDrugs

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetLinkedDrugs() []ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner`

GetLinkedDrugs returns the LinkedDrugs field if non-nil, zero value otherwise.

### GetLinkedDrugsOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetLinkedDrugsOk() (*[]ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner, bool)`

GetLinkedDrugsOk returns a tuple with the LinkedDrugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedDrugs

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetLinkedDrugs(v []ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner)`

SetLinkedDrugs sets LinkedDrugs field to given value.

### HasLinkedDrugs

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) HasLinkedDrugs() bool`

HasLinkedDrugs returns a boolean if a field has been set.

### GetLinkedDrugsBuilt

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetLinkedDrugsBuilt() bool`

GetLinkedDrugsBuilt returns the LinkedDrugsBuilt field if non-nil, zero value otherwise.

### GetLinkedDrugsBuiltOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) GetLinkedDrugsBuiltOk() (*bool, bool)`

GetLinkedDrugsBuiltOk returns a tuple with the LinkedDrugsBuilt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedDrugsBuilt

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) SetLinkedDrugsBuilt(v bool)`

SetLinkedDrugsBuilt sets LinkedDrugsBuilt field to given value.

### HasLinkedDrugsBuilt

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInner) HasLinkedDrugsBuilt() bool`

HasLinkedDrugsBuilt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


