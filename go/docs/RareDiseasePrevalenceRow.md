# RareDiseasePrevalenceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrevalenceAssertionId** | Pointer to **NullableString** | The identifier of the prevalence observation. | [optional] 
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] 
**PrevalenceType** | Pointer to **NullableString** | The kind of observation: point prevalence, prevalence at birth, annual incidence, lifetime prevalence or cases and families. | [optional] 
**PrevalenceClass** | Pointer to **NullableString** | The Orphanet prevalence class of the observation, for example 1-9 / 100 000; null when the source gives none. | [optional] 
**PrevalenceValueSource** | Pointer to **NullableString** | The prevalence value as the source writes it. | [optional] 
**NormalizedPrevalenceRatio** | Pointer to **NullableFloat32** | The observation as a ratio (cases per person) where one can be derived. | [optional] 
**PrevalenceLowerBound** | Pointer to **NullableFloat32** | The lower bound of the prevalence class as a ratio. | [optional] 
**PrevalenceUpperBound** | Pointer to **NullableFloat32** | The upper bound of the prevalence class as a ratio. | [optional] 
**PopulationDenominator** | Pointer to **NullableFloat32** | The denominator of the figure (for example 100000). | [optional] 
**PopulationUnit** | Pointer to **NullableString** | The unit of the denominator. | [optional] 
**PopulationScope** | Pointer to **NullableString** | The population the observation covers. | [optional] 
**AgeScope** | Pointer to **NullableString** | The age group the observation covers. | [optional] 
**Qualification** | Pointer to **NullableString** | How the figure is qualified at the source (value and class, class only, cases, families). | [optional] 
**ValidationStatus** | Pointer to **NullableString** | validated when Orphanet marks the observation validated, not_yet_validated otherwise. | [optional] 
**Geography** | Pointer to **NullableString** | The geography the observation covers (Worldwide, a country or a region). | [optional] 
**ObservationDate** | Pointer to **NullableString** | The date of the observation where the source gives one. | [optional] 
**UncertaintyLower** | Pointer to **NullableFloat32** | The lower end of the stated uncertainty range. | [optional] 
**UncertaintyUpper** | Pointer to **NullableFloat32** | The upper end of the stated uncertainty range. | [optional] 
**Source** | Pointer to **NullableString** | The evidence source Orphanet cites for the observation. | [optional] 
**SourceSnapshotDate** | Pointer to **NullableString** | The date of the Orphadata Science snapshot the row was read from. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The URL of the source file the row was read from. | [optional] 
**LicenseName** | Pointer to **NullableString** | The licence of the Orphadata Science file. | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRareDiseasePrevalenceRow

`func NewRareDiseasePrevalenceRow() *RareDiseasePrevalenceRow`

NewRareDiseasePrevalenceRow instantiates a new RareDiseasePrevalenceRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseasePrevalenceRowWithDefaults

`func NewRareDiseasePrevalenceRowWithDefaults() *RareDiseasePrevalenceRow`

NewRareDiseasePrevalenceRowWithDefaults instantiates a new RareDiseasePrevalenceRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrevalenceAssertionId

`func (o *RareDiseasePrevalenceRow) GetPrevalenceAssertionId() string`

GetPrevalenceAssertionId returns the PrevalenceAssertionId field if non-nil, zero value otherwise.

### GetPrevalenceAssertionIdOk

`func (o *RareDiseasePrevalenceRow) GetPrevalenceAssertionIdOk() (*string, bool)`

GetPrevalenceAssertionIdOk returns a tuple with the PrevalenceAssertionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevalenceAssertionId

`func (o *RareDiseasePrevalenceRow) SetPrevalenceAssertionId(v string)`

SetPrevalenceAssertionId sets PrevalenceAssertionId field to given value.

### HasPrevalenceAssertionId

`func (o *RareDiseasePrevalenceRow) HasPrevalenceAssertionId() bool`

HasPrevalenceAssertionId returns a boolean if a field has been set.

### SetPrevalenceAssertionIdNil

`func (o *RareDiseasePrevalenceRow) SetPrevalenceAssertionIdNil(b bool)`

 SetPrevalenceAssertionIdNil sets the value for PrevalenceAssertionId to be an explicit nil

### UnsetPrevalenceAssertionId
`func (o *RareDiseasePrevalenceRow) UnsetPrevalenceAssertionId()`

UnsetPrevalenceAssertionId ensures that no value is present for PrevalenceAssertionId, not even an explicit nil
### GetOrphaCode

`func (o *RareDiseasePrevalenceRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseasePrevalenceRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseasePrevalenceRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseasePrevalenceRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseasePrevalenceRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseasePrevalenceRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetPrevalenceType

`func (o *RareDiseasePrevalenceRow) GetPrevalenceType() string`

GetPrevalenceType returns the PrevalenceType field if non-nil, zero value otherwise.

### GetPrevalenceTypeOk

`func (o *RareDiseasePrevalenceRow) GetPrevalenceTypeOk() (*string, bool)`

GetPrevalenceTypeOk returns a tuple with the PrevalenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevalenceType

`func (o *RareDiseasePrevalenceRow) SetPrevalenceType(v string)`

SetPrevalenceType sets PrevalenceType field to given value.

### HasPrevalenceType

`func (o *RareDiseasePrevalenceRow) HasPrevalenceType() bool`

HasPrevalenceType returns a boolean if a field has been set.

### SetPrevalenceTypeNil

`func (o *RareDiseasePrevalenceRow) SetPrevalenceTypeNil(b bool)`

 SetPrevalenceTypeNil sets the value for PrevalenceType to be an explicit nil

### UnsetPrevalenceType
`func (o *RareDiseasePrevalenceRow) UnsetPrevalenceType()`

UnsetPrevalenceType ensures that no value is present for PrevalenceType, not even an explicit nil
### GetPrevalenceClass

`func (o *RareDiseasePrevalenceRow) GetPrevalenceClass() string`

GetPrevalenceClass returns the PrevalenceClass field if non-nil, zero value otherwise.

### GetPrevalenceClassOk

`func (o *RareDiseasePrevalenceRow) GetPrevalenceClassOk() (*string, bool)`

GetPrevalenceClassOk returns a tuple with the PrevalenceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevalenceClass

`func (o *RareDiseasePrevalenceRow) SetPrevalenceClass(v string)`

SetPrevalenceClass sets PrevalenceClass field to given value.

### HasPrevalenceClass

`func (o *RareDiseasePrevalenceRow) HasPrevalenceClass() bool`

HasPrevalenceClass returns a boolean if a field has been set.

### SetPrevalenceClassNil

`func (o *RareDiseasePrevalenceRow) SetPrevalenceClassNil(b bool)`

 SetPrevalenceClassNil sets the value for PrevalenceClass to be an explicit nil

### UnsetPrevalenceClass
`func (o *RareDiseasePrevalenceRow) UnsetPrevalenceClass()`

UnsetPrevalenceClass ensures that no value is present for PrevalenceClass, not even an explicit nil
### GetPrevalenceValueSource

`func (o *RareDiseasePrevalenceRow) GetPrevalenceValueSource() string`

GetPrevalenceValueSource returns the PrevalenceValueSource field if non-nil, zero value otherwise.

### GetPrevalenceValueSourceOk

`func (o *RareDiseasePrevalenceRow) GetPrevalenceValueSourceOk() (*string, bool)`

GetPrevalenceValueSourceOk returns a tuple with the PrevalenceValueSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevalenceValueSource

`func (o *RareDiseasePrevalenceRow) SetPrevalenceValueSource(v string)`

SetPrevalenceValueSource sets PrevalenceValueSource field to given value.

### HasPrevalenceValueSource

`func (o *RareDiseasePrevalenceRow) HasPrevalenceValueSource() bool`

HasPrevalenceValueSource returns a boolean if a field has been set.

### SetPrevalenceValueSourceNil

`func (o *RareDiseasePrevalenceRow) SetPrevalenceValueSourceNil(b bool)`

 SetPrevalenceValueSourceNil sets the value for PrevalenceValueSource to be an explicit nil

### UnsetPrevalenceValueSource
`func (o *RareDiseasePrevalenceRow) UnsetPrevalenceValueSource()`

UnsetPrevalenceValueSource ensures that no value is present for PrevalenceValueSource, not even an explicit nil
### GetNormalizedPrevalenceRatio

`func (o *RareDiseasePrevalenceRow) GetNormalizedPrevalenceRatio() float32`

GetNormalizedPrevalenceRatio returns the NormalizedPrevalenceRatio field if non-nil, zero value otherwise.

### GetNormalizedPrevalenceRatioOk

`func (o *RareDiseasePrevalenceRow) GetNormalizedPrevalenceRatioOk() (*float32, bool)`

GetNormalizedPrevalenceRatioOk returns a tuple with the NormalizedPrevalenceRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedPrevalenceRatio

`func (o *RareDiseasePrevalenceRow) SetNormalizedPrevalenceRatio(v float32)`

SetNormalizedPrevalenceRatio sets NormalizedPrevalenceRatio field to given value.

### HasNormalizedPrevalenceRatio

`func (o *RareDiseasePrevalenceRow) HasNormalizedPrevalenceRatio() bool`

HasNormalizedPrevalenceRatio returns a boolean if a field has been set.

### SetNormalizedPrevalenceRatioNil

`func (o *RareDiseasePrevalenceRow) SetNormalizedPrevalenceRatioNil(b bool)`

 SetNormalizedPrevalenceRatioNil sets the value for NormalizedPrevalenceRatio to be an explicit nil

### UnsetNormalizedPrevalenceRatio
`func (o *RareDiseasePrevalenceRow) UnsetNormalizedPrevalenceRatio()`

UnsetNormalizedPrevalenceRatio ensures that no value is present for NormalizedPrevalenceRatio, not even an explicit nil
### GetPrevalenceLowerBound

`func (o *RareDiseasePrevalenceRow) GetPrevalenceLowerBound() float32`

GetPrevalenceLowerBound returns the PrevalenceLowerBound field if non-nil, zero value otherwise.

### GetPrevalenceLowerBoundOk

`func (o *RareDiseasePrevalenceRow) GetPrevalenceLowerBoundOk() (*float32, bool)`

GetPrevalenceLowerBoundOk returns a tuple with the PrevalenceLowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevalenceLowerBound

`func (o *RareDiseasePrevalenceRow) SetPrevalenceLowerBound(v float32)`

SetPrevalenceLowerBound sets PrevalenceLowerBound field to given value.

### HasPrevalenceLowerBound

`func (o *RareDiseasePrevalenceRow) HasPrevalenceLowerBound() bool`

HasPrevalenceLowerBound returns a boolean if a field has been set.

### SetPrevalenceLowerBoundNil

`func (o *RareDiseasePrevalenceRow) SetPrevalenceLowerBoundNil(b bool)`

 SetPrevalenceLowerBoundNil sets the value for PrevalenceLowerBound to be an explicit nil

### UnsetPrevalenceLowerBound
`func (o *RareDiseasePrevalenceRow) UnsetPrevalenceLowerBound()`

UnsetPrevalenceLowerBound ensures that no value is present for PrevalenceLowerBound, not even an explicit nil
### GetPrevalenceUpperBound

`func (o *RareDiseasePrevalenceRow) GetPrevalenceUpperBound() float32`

GetPrevalenceUpperBound returns the PrevalenceUpperBound field if non-nil, zero value otherwise.

### GetPrevalenceUpperBoundOk

`func (o *RareDiseasePrevalenceRow) GetPrevalenceUpperBoundOk() (*float32, bool)`

GetPrevalenceUpperBoundOk returns a tuple with the PrevalenceUpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevalenceUpperBound

`func (o *RareDiseasePrevalenceRow) SetPrevalenceUpperBound(v float32)`

SetPrevalenceUpperBound sets PrevalenceUpperBound field to given value.

### HasPrevalenceUpperBound

`func (o *RareDiseasePrevalenceRow) HasPrevalenceUpperBound() bool`

HasPrevalenceUpperBound returns a boolean if a field has been set.

### SetPrevalenceUpperBoundNil

`func (o *RareDiseasePrevalenceRow) SetPrevalenceUpperBoundNil(b bool)`

 SetPrevalenceUpperBoundNil sets the value for PrevalenceUpperBound to be an explicit nil

### UnsetPrevalenceUpperBound
`func (o *RareDiseasePrevalenceRow) UnsetPrevalenceUpperBound()`

UnsetPrevalenceUpperBound ensures that no value is present for PrevalenceUpperBound, not even an explicit nil
### GetPopulationDenominator

`func (o *RareDiseasePrevalenceRow) GetPopulationDenominator() float32`

GetPopulationDenominator returns the PopulationDenominator field if non-nil, zero value otherwise.

### GetPopulationDenominatorOk

`func (o *RareDiseasePrevalenceRow) GetPopulationDenominatorOk() (*float32, bool)`

GetPopulationDenominatorOk returns a tuple with the PopulationDenominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationDenominator

`func (o *RareDiseasePrevalenceRow) SetPopulationDenominator(v float32)`

SetPopulationDenominator sets PopulationDenominator field to given value.

### HasPopulationDenominator

`func (o *RareDiseasePrevalenceRow) HasPopulationDenominator() bool`

HasPopulationDenominator returns a boolean if a field has been set.

### SetPopulationDenominatorNil

`func (o *RareDiseasePrevalenceRow) SetPopulationDenominatorNil(b bool)`

 SetPopulationDenominatorNil sets the value for PopulationDenominator to be an explicit nil

### UnsetPopulationDenominator
`func (o *RareDiseasePrevalenceRow) UnsetPopulationDenominator()`

UnsetPopulationDenominator ensures that no value is present for PopulationDenominator, not even an explicit nil
### GetPopulationUnit

`func (o *RareDiseasePrevalenceRow) GetPopulationUnit() string`

GetPopulationUnit returns the PopulationUnit field if non-nil, zero value otherwise.

### GetPopulationUnitOk

`func (o *RareDiseasePrevalenceRow) GetPopulationUnitOk() (*string, bool)`

GetPopulationUnitOk returns a tuple with the PopulationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationUnit

`func (o *RareDiseasePrevalenceRow) SetPopulationUnit(v string)`

SetPopulationUnit sets PopulationUnit field to given value.

### HasPopulationUnit

`func (o *RareDiseasePrevalenceRow) HasPopulationUnit() bool`

HasPopulationUnit returns a boolean if a field has been set.

### SetPopulationUnitNil

`func (o *RareDiseasePrevalenceRow) SetPopulationUnitNil(b bool)`

 SetPopulationUnitNil sets the value for PopulationUnit to be an explicit nil

### UnsetPopulationUnit
`func (o *RareDiseasePrevalenceRow) UnsetPopulationUnit()`

UnsetPopulationUnit ensures that no value is present for PopulationUnit, not even an explicit nil
### GetPopulationScope

`func (o *RareDiseasePrevalenceRow) GetPopulationScope() string`

GetPopulationScope returns the PopulationScope field if non-nil, zero value otherwise.

### GetPopulationScopeOk

`func (o *RareDiseasePrevalenceRow) GetPopulationScopeOk() (*string, bool)`

GetPopulationScopeOk returns a tuple with the PopulationScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationScope

`func (o *RareDiseasePrevalenceRow) SetPopulationScope(v string)`

SetPopulationScope sets PopulationScope field to given value.

### HasPopulationScope

`func (o *RareDiseasePrevalenceRow) HasPopulationScope() bool`

HasPopulationScope returns a boolean if a field has been set.

### SetPopulationScopeNil

`func (o *RareDiseasePrevalenceRow) SetPopulationScopeNil(b bool)`

 SetPopulationScopeNil sets the value for PopulationScope to be an explicit nil

### UnsetPopulationScope
`func (o *RareDiseasePrevalenceRow) UnsetPopulationScope()`

UnsetPopulationScope ensures that no value is present for PopulationScope, not even an explicit nil
### GetAgeScope

`func (o *RareDiseasePrevalenceRow) GetAgeScope() string`

GetAgeScope returns the AgeScope field if non-nil, zero value otherwise.

### GetAgeScopeOk

`func (o *RareDiseasePrevalenceRow) GetAgeScopeOk() (*string, bool)`

GetAgeScopeOk returns a tuple with the AgeScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeScope

`func (o *RareDiseasePrevalenceRow) SetAgeScope(v string)`

SetAgeScope sets AgeScope field to given value.

### HasAgeScope

`func (o *RareDiseasePrevalenceRow) HasAgeScope() bool`

HasAgeScope returns a boolean if a field has been set.

### SetAgeScopeNil

`func (o *RareDiseasePrevalenceRow) SetAgeScopeNil(b bool)`

 SetAgeScopeNil sets the value for AgeScope to be an explicit nil

### UnsetAgeScope
`func (o *RareDiseasePrevalenceRow) UnsetAgeScope()`

UnsetAgeScope ensures that no value is present for AgeScope, not even an explicit nil
### GetQualification

`func (o *RareDiseasePrevalenceRow) GetQualification() string`

GetQualification returns the Qualification field if non-nil, zero value otherwise.

### GetQualificationOk

`func (o *RareDiseasePrevalenceRow) GetQualificationOk() (*string, bool)`

GetQualificationOk returns a tuple with the Qualification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualification

`func (o *RareDiseasePrevalenceRow) SetQualification(v string)`

SetQualification sets Qualification field to given value.

### HasQualification

`func (o *RareDiseasePrevalenceRow) HasQualification() bool`

HasQualification returns a boolean if a field has been set.

### SetQualificationNil

`func (o *RareDiseasePrevalenceRow) SetQualificationNil(b bool)`

 SetQualificationNil sets the value for Qualification to be an explicit nil

### UnsetQualification
`func (o *RareDiseasePrevalenceRow) UnsetQualification()`

UnsetQualification ensures that no value is present for Qualification, not even an explicit nil
### GetValidationStatus

`func (o *RareDiseasePrevalenceRow) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *RareDiseasePrevalenceRow) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *RareDiseasePrevalenceRow) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *RareDiseasePrevalenceRow) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.

### SetValidationStatusNil

`func (o *RareDiseasePrevalenceRow) SetValidationStatusNil(b bool)`

 SetValidationStatusNil sets the value for ValidationStatus to be an explicit nil

### UnsetValidationStatus
`func (o *RareDiseasePrevalenceRow) UnsetValidationStatus()`

UnsetValidationStatus ensures that no value is present for ValidationStatus, not even an explicit nil
### GetGeography

`func (o *RareDiseasePrevalenceRow) GetGeography() string`

GetGeography returns the Geography field if non-nil, zero value otherwise.

### GetGeographyOk

`func (o *RareDiseasePrevalenceRow) GetGeographyOk() (*string, bool)`

GetGeographyOk returns a tuple with the Geography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeography

`func (o *RareDiseasePrevalenceRow) SetGeography(v string)`

SetGeography sets Geography field to given value.

### HasGeography

`func (o *RareDiseasePrevalenceRow) HasGeography() bool`

HasGeography returns a boolean if a field has been set.

### SetGeographyNil

`func (o *RareDiseasePrevalenceRow) SetGeographyNil(b bool)`

 SetGeographyNil sets the value for Geography to be an explicit nil

### UnsetGeography
`func (o *RareDiseasePrevalenceRow) UnsetGeography()`

UnsetGeography ensures that no value is present for Geography, not even an explicit nil
### GetObservationDate

`func (o *RareDiseasePrevalenceRow) GetObservationDate() string`

GetObservationDate returns the ObservationDate field if non-nil, zero value otherwise.

### GetObservationDateOk

`func (o *RareDiseasePrevalenceRow) GetObservationDateOk() (*string, bool)`

GetObservationDateOk returns a tuple with the ObservationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationDate

`func (o *RareDiseasePrevalenceRow) SetObservationDate(v string)`

SetObservationDate sets ObservationDate field to given value.

### HasObservationDate

`func (o *RareDiseasePrevalenceRow) HasObservationDate() bool`

HasObservationDate returns a boolean if a field has been set.

### SetObservationDateNil

`func (o *RareDiseasePrevalenceRow) SetObservationDateNil(b bool)`

 SetObservationDateNil sets the value for ObservationDate to be an explicit nil

### UnsetObservationDate
`func (o *RareDiseasePrevalenceRow) UnsetObservationDate()`

UnsetObservationDate ensures that no value is present for ObservationDate, not even an explicit nil
### GetUncertaintyLower

`func (o *RareDiseasePrevalenceRow) GetUncertaintyLower() float32`

GetUncertaintyLower returns the UncertaintyLower field if non-nil, zero value otherwise.

### GetUncertaintyLowerOk

`func (o *RareDiseasePrevalenceRow) GetUncertaintyLowerOk() (*float32, bool)`

GetUncertaintyLowerOk returns a tuple with the UncertaintyLower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertaintyLower

`func (o *RareDiseasePrevalenceRow) SetUncertaintyLower(v float32)`

SetUncertaintyLower sets UncertaintyLower field to given value.

### HasUncertaintyLower

`func (o *RareDiseasePrevalenceRow) HasUncertaintyLower() bool`

HasUncertaintyLower returns a boolean if a field has been set.

### SetUncertaintyLowerNil

`func (o *RareDiseasePrevalenceRow) SetUncertaintyLowerNil(b bool)`

 SetUncertaintyLowerNil sets the value for UncertaintyLower to be an explicit nil

### UnsetUncertaintyLower
`func (o *RareDiseasePrevalenceRow) UnsetUncertaintyLower()`

UnsetUncertaintyLower ensures that no value is present for UncertaintyLower, not even an explicit nil
### GetUncertaintyUpper

`func (o *RareDiseasePrevalenceRow) GetUncertaintyUpper() float32`

GetUncertaintyUpper returns the UncertaintyUpper field if non-nil, zero value otherwise.

### GetUncertaintyUpperOk

`func (o *RareDiseasePrevalenceRow) GetUncertaintyUpperOk() (*float32, bool)`

GetUncertaintyUpperOk returns a tuple with the UncertaintyUpper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertaintyUpper

`func (o *RareDiseasePrevalenceRow) SetUncertaintyUpper(v float32)`

SetUncertaintyUpper sets UncertaintyUpper field to given value.

### HasUncertaintyUpper

`func (o *RareDiseasePrevalenceRow) HasUncertaintyUpper() bool`

HasUncertaintyUpper returns a boolean if a field has been set.

### SetUncertaintyUpperNil

`func (o *RareDiseasePrevalenceRow) SetUncertaintyUpperNil(b bool)`

 SetUncertaintyUpperNil sets the value for UncertaintyUpper to be an explicit nil

### UnsetUncertaintyUpper
`func (o *RareDiseasePrevalenceRow) UnsetUncertaintyUpper()`

UnsetUncertaintyUpper ensures that no value is present for UncertaintyUpper, not even an explicit nil
### GetSource

`func (o *RareDiseasePrevalenceRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RareDiseasePrevalenceRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RareDiseasePrevalenceRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RareDiseasePrevalenceRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *RareDiseasePrevalenceRow) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RareDiseasePrevalenceRow) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceSnapshotDate

`func (o *RareDiseasePrevalenceRow) GetSourceSnapshotDate() string`

GetSourceSnapshotDate returns the SourceSnapshotDate field if non-nil, zero value otherwise.

### GetSourceSnapshotDateOk

`func (o *RareDiseasePrevalenceRow) GetSourceSnapshotDateOk() (*string, bool)`

GetSourceSnapshotDateOk returns a tuple with the SourceSnapshotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSnapshotDate

`func (o *RareDiseasePrevalenceRow) SetSourceSnapshotDate(v string)`

SetSourceSnapshotDate sets SourceSnapshotDate field to given value.

### HasSourceSnapshotDate

`func (o *RareDiseasePrevalenceRow) HasSourceSnapshotDate() bool`

HasSourceSnapshotDate returns a boolean if a field has been set.

### SetSourceSnapshotDateNil

`func (o *RareDiseasePrevalenceRow) SetSourceSnapshotDateNil(b bool)`

 SetSourceSnapshotDateNil sets the value for SourceSnapshotDate to be an explicit nil

### UnsetSourceSnapshotDate
`func (o *RareDiseasePrevalenceRow) UnsetSourceSnapshotDate()`

UnsetSourceSnapshotDate ensures that no value is present for SourceSnapshotDate, not even an explicit nil
### GetSourceUrl

`func (o *RareDiseasePrevalenceRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *RareDiseasePrevalenceRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *RareDiseasePrevalenceRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *RareDiseasePrevalenceRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *RareDiseasePrevalenceRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *RareDiseasePrevalenceRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetLicenseName

`func (o *RareDiseasePrevalenceRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *RareDiseasePrevalenceRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *RareDiseasePrevalenceRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *RareDiseasePrevalenceRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *RareDiseasePrevalenceRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *RareDiseasePrevalenceRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *RareDiseasePrevalenceRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RareDiseasePrevalenceRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RareDiseasePrevalenceRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RareDiseasePrevalenceRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RareDiseasePrevalenceRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RareDiseasePrevalenceRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


