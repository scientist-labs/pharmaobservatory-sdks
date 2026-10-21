# ConditionPrevalenceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiseasePrevalenceId** | Pointer to **NullableString** | The figure&#39;s key. | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s MeSH heading. | [optional] 
**ConditionMappingRelation** | Pointer to **NullableString** | How the figure&#39;s disease maps to the condition: exact, narrower, broader, undetermined or unspecified. | [optional] 
**DiseaseId** | Pointer to **NullableString** | The rare disease key of the figure&#39;s disease (orpha:&lt;code&gt;). | [optional] 
**DiseaseName** | Pointer to **NullableString** | The disease&#39;s name in the source. | [optional] 
**OrphaCode** | Pointer to **NullableString** | The Orphanet disease number. | [optional] 
**PrevalenceType** | Pointer to **NullableString** | What the figure measures: point_prevalence, lifetime_prevalence, annual_incidence, prevalence_at_birth or cases_families. | [optional] 
**PrevalenceClass** | Pointer to **NullableString** | The source&#39;s prevalence band, as filed. | [optional] 
**PrevalenceLowerBound** | Pointer to **NullableFloat32** | Lower bound of the prevalence band as a ratio of the population (0.0001 is 1 in 10,000). | [optional] 
**PrevalenceUpperBound** | Pointer to **NullableFloat32** | Upper bound of the prevalence band as a ratio of the population (0.0005 is 5 in 10,000). | [optional] 
**PopulationDenominator** | Pointer to **NullableFloat32** | The population size the source states the band against (for example 10000 for a band per 10,000 persons). | [optional] 
**PopulationUnit** | Pointer to **NullableString** | The unit of the population denominator. | [optional] 
**Qualification** | Pointer to **NullableString** | The source&#39;s qualification of the figure (class_only, value_and_class and so on). | [optional] 
**ValidationStatus** | Pointer to **NullableString** | The source&#39;s validation status for the figure (validated or not_yet_validated). | [optional] 
**Geography** | Pointer to **NullableString** | The geography the figure describes, as the source states it. | [optional] 
**JurisdictionCode** | Pointer to **NullableString** | US, EU (European Union) or GB (Great Britain) where the geography is that jurisdiction; null otherwise. | [optional] 
**ObservationDate** | Pointer to **NullableString** | The date the source assigns to the observation. | [optional] 
**SourceSnapshotDate** | Pointer to **NullableString** | The date of the source release the figure was read from. | [optional] 
**Source** | Pointer to **NullableString** | The publication reference the source cites for the figure. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The source release URL. | [optional] 
**LicenseName** | Pointer to **NullableString** | The figure&#39;s licence. | [optional] 
**Attribution** | Pointer to **NullableString** | The credit line the licence requires. | [optional] 

## Methods

### NewConditionPrevalenceRow

`func NewConditionPrevalenceRow() *ConditionPrevalenceRow`

NewConditionPrevalenceRow instantiates a new ConditionPrevalenceRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionPrevalenceRowWithDefaults

`func NewConditionPrevalenceRowWithDefaults() *ConditionPrevalenceRow`

NewConditionPrevalenceRowWithDefaults instantiates a new ConditionPrevalenceRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiseasePrevalenceId

`func (o *ConditionPrevalenceRow) GetDiseasePrevalenceId() string`

GetDiseasePrevalenceId returns the DiseasePrevalenceId field if non-nil, zero value otherwise.

### GetDiseasePrevalenceIdOk

`func (o *ConditionPrevalenceRow) GetDiseasePrevalenceIdOk() (*string, bool)`

GetDiseasePrevalenceIdOk returns a tuple with the DiseasePrevalenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiseasePrevalenceId

`func (o *ConditionPrevalenceRow) SetDiseasePrevalenceId(v string)`

SetDiseasePrevalenceId sets DiseasePrevalenceId field to given value.

### HasDiseasePrevalenceId

`func (o *ConditionPrevalenceRow) HasDiseasePrevalenceId() bool`

HasDiseasePrevalenceId returns a boolean if a field has been set.

### SetDiseasePrevalenceIdNil

`func (o *ConditionPrevalenceRow) SetDiseasePrevalenceIdNil(b bool)`

 SetDiseasePrevalenceIdNil sets the value for DiseasePrevalenceId to be an explicit nil

### UnsetDiseasePrevalenceId
`func (o *ConditionPrevalenceRow) UnsetDiseasePrevalenceId()`

UnsetDiseasePrevalenceId ensures that no value is present for DiseasePrevalenceId, not even an explicit nil
### GetConditionId

`func (o *ConditionPrevalenceRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionPrevalenceRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionPrevalenceRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionPrevalenceRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionPrevalenceRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionPrevalenceRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *ConditionPrevalenceRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *ConditionPrevalenceRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *ConditionPrevalenceRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *ConditionPrevalenceRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *ConditionPrevalenceRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *ConditionPrevalenceRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetConditionMappingRelation

`func (o *ConditionPrevalenceRow) GetConditionMappingRelation() string`

GetConditionMappingRelation returns the ConditionMappingRelation field if non-nil, zero value otherwise.

### GetConditionMappingRelationOk

`func (o *ConditionPrevalenceRow) GetConditionMappingRelationOk() (*string, bool)`

GetConditionMappingRelationOk returns a tuple with the ConditionMappingRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionMappingRelation

`func (o *ConditionPrevalenceRow) SetConditionMappingRelation(v string)`

SetConditionMappingRelation sets ConditionMappingRelation field to given value.

### HasConditionMappingRelation

`func (o *ConditionPrevalenceRow) HasConditionMappingRelation() bool`

HasConditionMappingRelation returns a boolean if a field has been set.

### SetConditionMappingRelationNil

`func (o *ConditionPrevalenceRow) SetConditionMappingRelationNil(b bool)`

 SetConditionMappingRelationNil sets the value for ConditionMappingRelation to be an explicit nil

### UnsetConditionMappingRelation
`func (o *ConditionPrevalenceRow) UnsetConditionMappingRelation()`

UnsetConditionMappingRelation ensures that no value is present for ConditionMappingRelation, not even an explicit nil
### GetDiseaseId

`func (o *ConditionPrevalenceRow) GetDiseaseId() string`

GetDiseaseId returns the DiseaseId field if non-nil, zero value otherwise.

### GetDiseaseIdOk

`func (o *ConditionPrevalenceRow) GetDiseaseIdOk() (*string, bool)`

GetDiseaseIdOk returns a tuple with the DiseaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiseaseId

`func (o *ConditionPrevalenceRow) SetDiseaseId(v string)`

SetDiseaseId sets DiseaseId field to given value.

### HasDiseaseId

`func (o *ConditionPrevalenceRow) HasDiseaseId() bool`

HasDiseaseId returns a boolean if a field has been set.

### SetDiseaseIdNil

`func (o *ConditionPrevalenceRow) SetDiseaseIdNil(b bool)`

 SetDiseaseIdNil sets the value for DiseaseId to be an explicit nil

### UnsetDiseaseId
`func (o *ConditionPrevalenceRow) UnsetDiseaseId()`

UnsetDiseaseId ensures that no value is present for DiseaseId, not even an explicit nil
### GetDiseaseName

`func (o *ConditionPrevalenceRow) GetDiseaseName() string`

GetDiseaseName returns the DiseaseName field if non-nil, zero value otherwise.

### GetDiseaseNameOk

`func (o *ConditionPrevalenceRow) GetDiseaseNameOk() (*string, bool)`

GetDiseaseNameOk returns a tuple with the DiseaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiseaseName

`func (o *ConditionPrevalenceRow) SetDiseaseName(v string)`

SetDiseaseName sets DiseaseName field to given value.

### HasDiseaseName

`func (o *ConditionPrevalenceRow) HasDiseaseName() bool`

HasDiseaseName returns a boolean if a field has been set.

### SetDiseaseNameNil

`func (o *ConditionPrevalenceRow) SetDiseaseNameNil(b bool)`

 SetDiseaseNameNil sets the value for DiseaseName to be an explicit nil

### UnsetDiseaseName
`func (o *ConditionPrevalenceRow) UnsetDiseaseName()`

UnsetDiseaseName ensures that no value is present for DiseaseName, not even an explicit nil
### GetOrphaCode

`func (o *ConditionPrevalenceRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *ConditionPrevalenceRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *ConditionPrevalenceRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *ConditionPrevalenceRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *ConditionPrevalenceRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *ConditionPrevalenceRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetPrevalenceType

`func (o *ConditionPrevalenceRow) GetPrevalenceType() string`

GetPrevalenceType returns the PrevalenceType field if non-nil, zero value otherwise.

### GetPrevalenceTypeOk

`func (o *ConditionPrevalenceRow) GetPrevalenceTypeOk() (*string, bool)`

GetPrevalenceTypeOk returns a tuple with the PrevalenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevalenceType

`func (o *ConditionPrevalenceRow) SetPrevalenceType(v string)`

SetPrevalenceType sets PrevalenceType field to given value.

### HasPrevalenceType

`func (o *ConditionPrevalenceRow) HasPrevalenceType() bool`

HasPrevalenceType returns a boolean if a field has been set.

### SetPrevalenceTypeNil

`func (o *ConditionPrevalenceRow) SetPrevalenceTypeNil(b bool)`

 SetPrevalenceTypeNil sets the value for PrevalenceType to be an explicit nil

### UnsetPrevalenceType
`func (o *ConditionPrevalenceRow) UnsetPrevalenceType()`

UnsetPrevalenceType ensures that no value is present for PrevalenceType, not even an explicit nil
### GetPrevalenceClass

`func (o *ConditionPrevalenceRow) GetPrevalenceClass() string`

GetPrevalenceClass returns the PrevalenceClass field if non-nil, zero value otherwise.

### GetPrevalenceClassOk

`func (o *ConditionPrevalenceRow) GetPrevalenceClassOk() (*string, bool)`

GetPrevalenceClassOk returns a tuple with the PrevalenceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevalenceClass

`func (o *ConditionPrevalenceRow) SetPrevalenceClass(v string)`

SetPrevalenceClass sets PrevalenceClass field to given value.

### HasPrevalenceClass

`func (o *ConditionPrevalenceRow) HasPrevalenceClass() bool`

HasPrevalenceClass returns a boolean if a field has been set.

### SetPrevalenceClassNil

`func (o *ConditionPrevalenceRow) SetPrevalenceClassNil(b bool)`

 SetPrevalenceClassNil sets the value for PrevalenceClass to be an explicit nil

### UnsetPrevalenceClass
`func (o *ConditionPrevalenceRow) UnsetPrevalenceClass()`

UnsetPrevalenceClass ensures that no value is present for PrevalenceClass, not even an explicit nil
### GetPrevalenceLowerBound

`func (o *ConditionPrevalenceRow) GetPrevalenceLowerBound() float32`

GetPrevalenceLowerBound returns the PrevalenceLowerBound field if non-nil, zero value otherwise.

### GetPrevalenceLowerBoundOk

`func (o *ConditionPrevalenceRow) GetPrevalenceLowerBoundOk() (*float32, bool)`

GetPrevalenceLowerBoundOk returns a tuple with the PrevalenceLowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevalenceLowerBound

`func (o *ConditionPrevalenceRow) SetPrevalenceLowerBound(v float32)`

SetPrevalenceLowerBound sets PrevalenceLowerBound field to given value.

### HasPrevalenceLowerBound

`func (o *ConditionPrevalenceRow) HasPrevalenceLowerBound() bool`

HasPrevalenceLowerBound returns a boolean if a field has been set.

### SetPrevalenceLowerBoundNil

`func (o *ConditionPrevalenceRow) SetPrevalenceLowerBoundNil(b bool)`

 SetPrevalenceLowerBoundNil sets the value for PrevalenceLowerBound to be an explicit nil

### UnsetPrevalenceLowerBound
`func (o *ConditionPrevalenceRow) UnsetPrevalenceLowerBound()`

UnsetPrevalenceLowerBound ensures that no value is present for PrevalenceLowerBound, not even an explicit nil
### GetPrevalenceUpperBound

`func (o *ConditionPrevalenceRow) GetPrevalenceUpperBound() float32`

GetPrevalenceUpperBound returns the PrevalenceUpperBound field if non-nil, zero value otherwise.

### GetPrevalenceUpperBoundOk

`func (o *ConditionPrevalenceRow) GetPrevalenceUpperBoundOk() (*float32, bool)`

GetPrevalenceUpperBoundOk returns a tuple with the PrevalenceUpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevalenceUpperBound

`func (o *ConditionPrevalenceRow) SetPrevalenceUpperBound(v float32)`

SetPrevalenceUpperBound sets PrevalenceUpperBound field to given value.

### HasPrevalenceUpperBound

`func (o *ConditionPrevalenceRow) HasPrevalenceUpperBound() bool`

HasPrevalenceUpperBound returns a boolean if a field has been set.

### SetPrevalenceUpperBoundNil

`func (o *ConditionPrevalenceRow) SetPrevalenceUpperBoundNil(b bool)`

 SetPrevalenceUpperBoundNil sets the value for PrevalenceUpperBound to be an explicit nil

### UnsetPrevalenceUpperBound
`func (o *ConditionPrevalenceRow) UnsetPrevalenceUpperBound()`

UnsetPrevalenceUpperBound ensures that no value is present for PrevalenceUpperBound, not even an explicit nil
### GetPopulationDenominator

`func (o *ConditionPrevalenceRow) GetPopulationDenominator() float32`

GetPopulationDenominator returns the PopulationDenominator field if non-nil, zero value otherwise.

### GetPopulationDenominatorOk

`func (o *ConditionPrevalenceRow) GetPopulationDenominatorOk() (*float32, bool)`

GetPopulationDenominatorOk returns a tuple with the PopulationDenominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationDenominator

`func (o *ConditionPrevalenceRow) SetPopulationDenominator(v float32)`

SetPopulationDenominator sets PopulationDenominator field to given value.

### HasPopulationDenominator

`func (o *ConditionPrevalenceRow) HasPopulationDenominator() bool`

HasPopulationDenominator returns a boolean if a field has been set.

### SetPopulationDenominatorNil

`func (o *ConditionPrevalenceRow) SetPopulationDenominatorNil(b bool)`

 SetPopulationDenominatorNil sets the value for PopulationDenominator to be an explicit nil

### UnsetPopulationDenominator
`func (o *ConditionPrevalenceRow) UnsetPopulationDenominator()`

UnsetPopulationDenominator ensures that no value is present for PopulationDenominator, not even an explicit nil
### GetPopulationUnit

`func (o *ConditionPrevalenceRow) GetPopulationUnit() string`

GetPopulationUnit returns the PopulationUnit field if non-nil, zero value otherwise.

### GetPopulationUnitOk

`func (o *ConditionPrevalenceRow) GetPopulationUnitOk() (*string, bool)`

GetPopulationUnitOk returns a tuple with the PopulationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationUnit

`func (o *ConditionPrevalenceRow) SetPopulationUnit(v string)`

SetPopulationUnit sets PopulationUnit field to given value.

### HasPopulationUnit

`func (o *ConditionPrevalenceRow) HasPopulationUnit() bool`

HasPopulationUnit returns a boolean if a field has been set.

### SetPopulationUnitNil

`func (o *ConditionPrevalenceRow) SetPopulationUnitNil(b bool)`

 SetPopulationUnitNil sets the value for PopulationUnit to be an explicit nil

### UnsetPopulationUnit
`func (o *ConditionPrevalenceRow) UnsetPopulationUnit()`

UnsetPopulationUnit ensures that no value is present for PopulationUnit, not even an explicit nil
### GetQualification

`func (o *ConditionPrevalenceRow) GetQualification() string`

GetQualification returns the Qualification field if non-nil, zero value otherwise.

### GetQualificationOk

`func (o *ConditionPrevalenceRow) GetQualificationOk() (*string, bool)`

GetQualificationOk returns a tuple with the Qualification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualification

`func (o *ConditionPrevalenceRow) SetQualification(v string)`

SetQualification sets Qualification field to given value.

### HasQualification

`func (o *ConditionPrevalenceRow) HasQualification() bool`

HasQualification returns a boolean if a field has been set.

### SetQualificationNil

`func (o *ConditionPrevalenceRow) SetQualificationNil(b bool)`

 SetQualificationNil sets the value for Qualification to be an explicit nil

### UnsetQualification
`func (o *ConditionPrevalenceRow) UnsetQualification()`

UnsetQualification ensures that no value is present for Qualification, not even an explicit nil
### GetValidationStatus

`func (o *ConditionPrevalenceRow) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *ConditionPrevalenceRow) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *ConditionPrevalenceRow) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *ConditionPrevalenceRow) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.

### SetValidationStatusNil

`func (o *ConditionPrevalenceRow) SetValidationStatusNil(b bool)`

 SetValidationStatusNil sets the value for ValidationStatus to be an explicit nil

### UnsetValidationStatus
`func (o *ConditionPrevalenceRow) UnsetValidationStatus()`

UnsetValidationStatus ensures that no value is present for ValidationStatus, not even an explicit nil
### GetGeography

`func (o *ConditionPrevalenceRow) GetGeography() string`

GetGeography returns the Geography field if non-nil, zero value otherwise.

### GetGeographyOk

`func (o *ConditionPrevalenceRow) GetGeographyOk() (*string, bool)`

GetGeographyOk returns a tuple with the Geography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeography

`func (o *ConditionPrevalenceRow) SetGeography(v string)`

SetGeography sets Geography field to given value.

### HasGeography

`func (o *ConditionPrevalenceRow) HasGeography() bool`

HasGeography returns a boolean if a field has been set.

### SetGeographyNil

`func (o *ConditionPrevalenceRow) SetGeographyNil(b bool)`

 SetGeographyNil sets the value for Geography to be an explicit nil

### UnsetGeography
`func (o *ConditionPrevalenceRow) UnsetGeography()`

UnsetGeography ensures that no value is present for Geography, not even an explicit nil
### GetJurisdictionCode

`func (o *ConditionPrevalenceRow) GetJurisdictionCode() string`

GetJurisdictionCode returns the JurisdictionCode field if non-nil, zero value otherwise.

### GetJurisdictionCodeOk

`func (o *ConditionPrevalenceRow) GetJurisdictionCodeOk() (*string, bool)`

GetJurisdictionCodeOk returns a tuple with the JurisdictionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionCode

`func (o *ConditionPrevalenceRow) SetJurisdictionCode(v string)`

SetJurisdictionCode sets JurisdictionCode field to given value.

### HasJurisdictionCode

`func (o *ConditionPrevalenceRow) HasJurisdictionCode() bool`

HasJurisdictionCode returns a boolean if a field has been set.

### SetJurisdictionCodeNil

`func (o *ConditionPrevalenceRow) SetJurisdictionCodeNil(b bool)`

 SetJurisdictionCodeNil sets the value for JurisdictionCode to be an explicit nil

### UnsetJurisdictionCode
`func (o *ConditionPrevalenceRow) UnsetJurisdictionCode()`

UnsetJurisdictionCode ensures that no value is present for JurisdictionCode, not even an explicit nil
### GetObservationDate

`func (o *ConditionPrevalenceRow) GetObservationDate() string`

GetObservationDate returns the ObservationDate field if non-nil, zero value otherwise.

### GetObservationDateOk

`func (o *ConditionPrevalenceRow) GetObservationDateOk() (*string, bool)`

GetObservationDateOk returns a tuple with the ObservationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationDate

`func (o *ConditionPrevalenceRow) SetObservationDate(v string)`

SetObservationDate sets ObservationDate field to given value.

### HasObservationDate

`func (o *ConditionPrevalenceRow) HasObservationDate() bool`

HasObservationDate returns a boolean if a field has been set.

### SetObservationDateNil

`func (o *ConditionPrevalenceRow) SetObservationDateNil(b bool)`

 SetObservationDateNil sets the value for ObservationDate to be an explicit nil

### UnsetObservationDate
`func (o *ConditionPrevalenceRow) UnsetObservationDate()`

UnsetObservationDate ensures that no value is present for ObservationDate, not even an explicit nil
### GetSourceSnapshotDate

`func (o *ConditionPrevalenceRow) GetSourceSnapshotDate() string`

GetSourceSnapshotDate returns the SourceSnapshotDate field if non-nil, zero value otherwise.

### GetSourceSnapshotDateOk

`func (o *ConditionPrevalenceRow) GetSourceSnapshotDateOk() (*string, bool)`

GetSourceSnapshotDateOk returns a tuple with the SourceSnapshotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSnapshotDate

`func (o *ConditionPrevalenceRow) SetSourceSnapshotDate(v string)`

SetSourceSnapshotDate sets SourceSnapshotDate field to given value.

### HasSourceSnapshotDate

`func (o *ConditionPrevalenceRow) HasSourceSnapshotDate() bool`

HasSourceSnapshotDate returns a boolean if a field has been set.

### SetSourceSnapshotDateNil

`func (o *ConditionPrevalenceRow) SetSourceSnapshotDateNil(b bool)`

 SetSourceSnapshotDateNil sets the value for SourceSnapshotDate to be an explicit nil

### UnsetSourceSnapshotDate
`func (o *ConditionPrevalenceRow) UnsetSourceSnapshotDate()`

UnsetSourceSnapshotDate ensures that no value is present for SourceSnapshotDate, not even an explicit nil
### GetSource

`func (o *ConditionPrevalenceRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConditionPrevalenceRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConditionPrevalenceRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConditionPrevalenceRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *ConditionPrevalenceRow) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *ConditionPrevalenceRow) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceUrl

`func (o *ConditionPrevalenceRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *ConditionPrevalenceRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *ConditionPrevalenceRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *ConditionPrevalenceRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *ConditionPrevalenceRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *ConditionPrevalenceRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetLicenseName

`func (o *ConditionPrevalenceRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *ConditionPrevalenceRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *ConditionPrevalenceRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *ConditionPrevalenceRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *ConditionPrevalenceRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *ConditionPrevalenceRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *ConditionPrevalenceRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *ConditionPrevalenceRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *ConditionPrevalenceRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *ConditionPrevalenceRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *ConditionPrevalenceRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *ConditionPrevalenceRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


