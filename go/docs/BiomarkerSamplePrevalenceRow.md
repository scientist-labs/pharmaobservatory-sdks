# BiomarkerSamplePrevalenceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteKey** | Pointer to **NullableString** | The route key of the gene the variants belong to, the entry key for the biomarkers object type. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol. | [optional] 
**VariantBiomarkerId** | Pointer to **NullableString** | The variant&#39;s own biomarker key. | [optional] 
**VariantRouteKey** | Pointer to **NullableString** | The variant&#39;s own route key (the id its web page uses). | [optional] 
**AminoAcidChange** | Pointer to **NullableString** | The protein change, for example p.V600E. | [optional] 
**GdcSsmId** | Pointer to **NullableString** | The identifier of the somatic mutation in the Genomic Data Commons. | [optional] 
**GdcProjectId** | Pointer to **NullableString** | The Genomic Data Commons research project (cohort) the counts come from. | [optional] 
**GdcProjectName** | Pointer to **NullableString** | The name of the research project. | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition key (a MeSH, Medical Subject Headings, descriptor id) the project&#39;s disease maps to, where one is mapped. | [optional] 
**ConditionName** | Pointer to **NullableString** | The mapped condition&#39;s MeSH heading. | [optional] 
**SourceDiseaseTypes** | Pointer to **[]string** | The disease types the source states for the project. | [optional] 
**SourcePrimarySites** | Pointer to **[]string** | The primary sites the source states for the project. | [optional] 
**AffectedCaseCount** | Pointer to **NullableInt64** | Cases in the cohort with the exact variant. | [optional] 
**SsmProfiledCaseCount** | Pointer to **NullableInt64** | Cases in the cohort profiled for somatic mutations, the denominator of the observed frequency. | [optional] 
**AllCaseCount** | Pointer to **NullableInt64** | All cases in the project. | [optional] 
**ObservedFrequency** | Pointer to **NullableFloat32** | Affected cases over mutation-profiled cases in this cohort. | [optional] 
**FrequencyCi95Low** | Pointer to **NullableFloat32** | Lower bound of the 95 percent confidence interval of the observed frequency. | [optional] 
**FrequencyCi95High** | Pointer to **NullableFloat32** | Upper bound of the 95 percent confidence interval of the observed frequency. | [optional] 
**GdcDataRelease** | Pointer to **NullableString** | The Genomic Data Commons data release the counts come from. | [optional] 
**ClaimType** | Pointer to **NullableString** | What the figure states; an observed frequency in a named cohort. | [optional] 

## Methods

### NewBiomarkerSamplePrevalenceRow

`func NewBiomarkerSamplePrevalenceRow() *BiomarkerSamplePrevalenceRow`

NewBiomarkerSamplePrevalenceRow instantiates a new BiomarkerSamplePrevalenceRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerSamplePrevalenceRowWithDefaults

`func NewBiomarkerSamplePrevalenceRowWithDefaults() *BiomarkerSamplePrevalenceRow`

NewBiomarkerSamplePrevalenceRowWithDefaults instantiates a new BiomarkerSamplePrevalenceRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteKey

`func (o *BiomarkerSamplePrevalenceRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerSamplePrevalenceRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerSamplePrevalenceRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerSamplePrevalenceRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerSamplePrevalenceRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerSamplePrevalenceRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetGeneSymbol

`func (o *BiomarkerSamplePrevalenceRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *BiomarkerSamplePrevalenceRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *BiomarkerSamplePrevalenceRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *BiomarkerSamplePrevalenceRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *BiomarkerSamplePrevalenceRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *BiomarkerSamplePrevalenceRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetVariantBiomarkerId

`func (o *BiomarkerSamplePrevalenceRow) GetVariantBiomarkerId() string`

GetVariantBiomarkerId returns the VariantBiomarkerId field if non-nil, zero value otherwise.

### GetVariantBiomarkerIdOk

`func (o *BiomarkerSamplePrevalenceRow) GetVariantBiomarkerIdOk() (*string, bool)`

GetVariantBiomarkerIdOk returns a tuple with the VariantBiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantBiomarkerId

`func (o *BiomarkerSamplePrevalenceRow) SetVariantBiomarkerId(v string)`

SetVariantBiomarkerId sets VariantBiomarkerId field to given value.

### HasVariantBiomarkerId

`func (o *BiomarkerSamplePrevalenceRow) HasVariantBiomarkerId() bool`

HasVariantBiomarkerId returns a boolean if a field has been set.

### SetVariantBiomarkerIdNil

`func (o *BiomarkerSamplePrevalenceRow) SetVariantBiomarkerIdNil(b bool)`

 SetVariantBiomarkerIdNil sets the value for VariantBiomarkerId to be an explicit nil

### UnsetVariantBiomarkerId
`func (o *BiomarkerSamplePrevalenceRow) UnsetVariantBiomarkerId()`

UnsetVariantBiomarkerId ensures that no value is present for VariantBiomarkerId, not even an explicit nil
### GetVariantRouteKey

`func (o *BiomarkerSamplePrevalenceRow) GetVariantRouteKey() string`

GetVariantRouteKey returns the VariantRouteKey field if non-nil, zero value otherwise.

### GetVariantRouteKeyOk

`func (o *BiomarkerSamplePrevalenceRow) GetVariantRouteKeyOk() (*string, bool)`

GetVariantRouteKeyOk returns a tuple with the VariantRouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantRouteKey

`func (o *BiomarkerSamplePrevalenceRow) SetVariantRouteKey(v string)`

SetVariantRouteKey sets VariantRouteKey field to given value.

### HasVariantRouteKey

`func (o *BiomarkerSamplePrevalenceRow) HasVariantRouteKey() bool`

HasVariantRouteKey returns a boolean if a field has been set.

### SetVariantRouteKeyNil

`func (o *BiomarkerSamplePrevalenceRow) SetVariantRouteKeyNil(b bool)`

 SetVariantRouteKeyNil sets the value for VariantRouteKey to be an explicit nil

### UnsetVariantRouteKey
`func (o *BiomarkerSamplePrevalenceRow) UnsetVariantRouteKey()`

UnsetVariantRouteKey ensures that no value is present for VariantRouteKey, not even an explicit nil
### GetAminoAcidChange

`func (o *BiomarkerSamplePrevalenceRow) GetAminoAcidChange() string`

GetAminoAcidChange returns the AminoAcidChange field if non-nil, zero value otherwise.

### GetAminoAcidChangeOk

`func (o *BiomarkerSamplePrevalenceRow) GetAminoAcidChangeOk() (*string, bool)`

GetAminoAcidChangeOk returns a tuple with the AminoAcidChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcidChange

`func (o *BiomarkerSamplePrevalenceRow) SetAminoAcidChange(v string)`

SetAminoAcidChange sets AminoAcidChange field to given value.

### HasAminoAcidChange

`func (o *BiomarkerSamplePrevalenceRow) HasAminoAcidChange() bool`

HasAminoAcidChange returns a boolean if a field has been set.

### SetAminoAcidChangeNil

`func (o *BiomarkerSamplePrevalenceRow) SetAminoAcidChangeNil(b bool)`

 SetAminoAcidChangeNil sets the value for AminoAcidChange to be an explicit nil

### UnsetAminoAcidChange
`func (o *BiomarkerSamplePrevalenceRow) UnsetAminoAcidChange()`

UnsetAminoAcidChange ensures that no value is present for AminoAcidChange, not even an explicit nil
### GetGdcSsmId

`func (o *BiomarkerSamplePrevalenceRow) GetGdcSsmId() string`

GetGdcSsmId returns the GdcSsmId field if non-nil, zero value otherwise.

### GetGdcSsmIdOk

`func (o *BiomarkerSamplePrevalenceRow) GetGdcSsmIdOk() (*string, bool)`

GetGdcSsmIdOk returns a tuple with the GdcSsmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdcSsmId

`func (o *BiomarkerSamplePrevalenceRow) SetGdcSsmId(v string)`

SetGdcSsmId sets GdcSsmId field to given value.

### HasGdcSsmId

`func (o *BiomarkerSamplePrevalenceRow) HasGdcSsmId() bool`

HasGdcSsmId returns a boolean if a field has been set.

### SetGdcSsmIdNil

`func (o *BiomarkerSamplePrevalenceRow) SetGdcSsmIdNil(b bool)`

 SetGdcSsmIdNil sets the value for GdcSsmId to be an explicit nil

### UnsetGdcSsmId
`func (o *BiomarkerSamplePrevalenceRow) UnsetGdcSsmId()`

UnsetGdcSsmId ensures that no value is present for GdcSsmId, not even an explicit nil
### GetGdcProjectId

`func (o *BiomarkerSamplePrevalenceRow) GetGdcProjectId() string`

GetGdcProjectId returns the GdcProjectId field if non-nil, zero value otherwise.

### GetGdcProjectIdOk

`func (o *BiomarkerSamplePrevalenceRow) GetGdcProjectIdOk() (*string, bool)`

GetGdcProjectIdOk returns a tuple with the GdcProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdcProjectId

`func (o *BiomarkerSamplePrevalenceRow) SetGdcProjectId(v string)`

SetGdcProjectId sets GdcProjectId field to given value.

### HasGdcProjectId

`func (o *BiomarkerSamplePrevalenceRow) HasGdcProjectId() bool`

HasGdcProjectId returns a boolean if a field has been set.

### SetGdcProjectIdNil

`func (o *BiomarkerSamplePrevalenceRow) SetGdcProjectIdNil(b bool)`

 SetGdcProjectIdNil sets the value for GdcProjectId to be an explicit nil

### UnsetGdcProjectId
`func (o *BiomarkerSamplePrevalenceRow) UnsetGdcProjectId()`

UnsetGdcProjectId ensures that no value is present for GdcProjectId, not even an explicit nil
### GetGdcProjectName

`func (o *BiomarkerSamplePrevalenceRow) GetGdcProjectName() string`

GetGdcProjectName returns the GdcProjectName field if non-nil, zero value otherwise.

### GetGdcProjectNameOk

`func (o *BiomarkerSamplePrevalenceRow) GetGdcProjectNameOk() (*string, bool)`

GetGdcProjectNameOk returns a tuple with the GdcProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdcProjectName

`func (o *BiomarkerSamplePrevalenceRow) SetGdcProjectName(v string)`

SetGdcProjectName sets GdcProjectName field to given value.

### HasGdcProjectName

`func (o *BiomarkerSamplePrevalenceRow) HasGdcProjectName() bool`

HasGdcProjectName returns a boolean if a field has been set.

### SetGdcProjectNameNil

`func (o *BiomarkerSamplePrevalenceRow) SetGdcProjectNameNil(b bool)`

 SetGdcProjectNameNil sets the value for GdcProjectName to be an explicit nil

### UnsetGdcProjectName
`func (o *BiomarkerSamplePrevalenceRow) UnsetGdcProjectName()`

UnsetGdcProjectName ensures that no value is present for GdcProjectName, not even an explicit nil
### GetConditionId

`func (o *BiomarkerSamplePrevalenceRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *BiomarkerSamplePrevalenceRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *BiomarkerSamplePrevalenceRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *BiomarkerSamplePrevalenceRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *BiomarkerSamplePrevalenceRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *BiomarkerSamplePrevalenceRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *BiomarkerSamplePrevalenceRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *BiomarkerSamplePrevalenceRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *BiomarkerSamplePrevalenceRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *BiomarkerSamplePrevalenceRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *BiomarkerSamplePrevalenceRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *BiomarkerSamplePrevalenceRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetSourceDiseaseTypes

`func (o *BiomarkerSamplePrevalenceRow) GetSourceDiseaseTypes() []string`

GetSourceDiseaseTypes returns the SourceDiseaseTypes field if non-nil, zero value otherwise.

### GetSourceDiseaseTypesOk

`func (o *BiomarkerSamplePrevalenceRow) GetSourceDiseaseTypesOk() (*[]string, bool)`

GetSourceDiseaseTypesOk returns a tuple with the SourceDiseaseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDiseaseTypes

`func (o *BiomarkerSamplePrevalenceRow) SetSourceDiseaseTypes(v []string)`

SetSourceDiseaseTypes sets SourceDiseaseTypes field to given value.

### HasSourceDiseaseTypes

`func (o *BiomarkerSamplePrevalenceRow) HasSourceDiseaseTypes() bool`

HasSourceDiseaseTypes returns a boolean if a field has been set.

### SetSourceDiseaseTypesNil

`func (o *BiomarkerSamplePrevalenceRow) SetSourceDiseaseTypesNil(b bool)`

 SetSourceDiseaseTypesNil sets the value for SourceDiseaseTypes to be an explicit nil

### UnsetSourceDiseaseTypes
`func (o *BiomarkerSamplePrevalenceRow) UnsetSourceDiseaseTypes()`

UnsetSourceDiseaseTypes ensures that no value is present for SourceDiseaseTypes, not even an explicit nil
### GetSourcePrimarySites

`func (o *BiomarkerSamplePrevalenceRow) GetSourcePrimarySites() []string`

GetSourcePrimarySites returns the SourcePrimarySites field if non-nil, zero value otherwise.

### GetSourcePrimarySitesOk

`func (o *BiomarkerSamplePrevalenceRow) GetSourcePrimarySitesOk() (*[]string, bool)`

GetSourcePrimarySitesOk returns a tuple with the SourcePrimarySites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePrimarySites

`func (o *BiomarkerSamplePrevalenceRow) SetSourcePrimarySites(v []string)`

SetSourcePrimarySites sets SourcePrimarySites field to given value.

### HasSourcePrimarySites

`func (o *BiomarkerSamplePrevalenceRow) HasSourcePrimarySites() bool`

HasSourcePrimarySites returns a boolean if a field has been set.

### SetSourcePrimarySitesNil

`func (o *BiomarkerSamplePrevalenceRow) SetSourcePrimarySitesNil(b bool)`

 SetSourcePrimarySitesNil sets the value for SourcePrimarySites to be an explicit nil

### UnsetSourcePrimarySites
`func (o *BiomarkerSamplePrevalenceRow) UnsetSourcePrimarySites()`

UnsetSourcePrimarySites ensures that no value is present for SourcePrimarySites, not even an explicit nil
### GetAffectedCaseCount

`func (o *BiomarkerSamplePrevalenceRow) GetAffectedCaseCount() int64`

GetAffectedCaseCount returns the AffectedCaseCount field if non-nil, zero value otherwise.

### GetAffectedCaseCountOk

`func (o *BiomarkerSamplePrevalenceRow) GetAffectedCaseCountOk() (*int64, bool)`

GetAffectedCaseCountOk returns a tuple with the AffectedCaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedCaseCount

`func (o *BiomarkerSamplePrevalenceRow) SetAffectedCaseCount(v int64)`

SetAffectedCaseCount sets AffectedCaseCount field to given value.

### HasAffectedCaseCount

`func (o *BiomarkerSamplePrevalenceRow) HasAffectedCaseCount() bool`

HasAffectedCaseCount returns a boolean if a field has been set.

### SetAffectedCaseCountNil

`func (o *BiomarkerSamplePrevalenceRow) SetAffectedCaseCountNil(b bool)`

 SetAffectedCaseCountNil sets the value for AffectedCaseCount to be an explicit nil

### UnsetAffectedCaseCount
`func (o *BiomarkerSamplePrevalenceRow) UnsetAffectedCaseCount()`

UnsetAffectedCaseCount ensures that no value is present for AffectedCaseCount, not even an explicit nil
### GetSsmProfiledCaseCount

`func (o *BiomarkerSamplePrevalenceRow) GetSsmProfiledCaseCount() int64`

GetSsmProfiledCaseCount returns the SsmProfiledCaseCount field if non-nil, zero value otherwise.

### GetSsmProfiledCaseCountOk

`func (o *BiomarkerSamplePrevalenceRow) GetSsmProfiledCaseCountOk() (*int64, bool)`

GetSsmProfiledCaseCountOk returns a tuple with the SsmProfiledCaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsmProfiledCaseCount

`func (o *BiomarkerSamplePrevalenceRow) SetSsmProfiledCaseCount(v int64)`

SetSsmProfiledCaseCount sets SsmProfiledCaseCount field to given value.

### HasSsmProfiledCaseCount

`func (o *BiomarkerSamplePrevalenceRow) HasSsmProfiledCaseCount() bool`

HasSsmProfiledCaseCount returns a boolean if a field has been set.

### SetSsmProfiledCaseCountNil

`func (o *BiomarkerSamplePrevalenceRow) SetSsmProfiledCaseCountNil(b bool)`

 SetSsmProfiledCaseCountNil sets the value for SsmProfiledCaseCount to be an explicit nil

### UnsetSsmProfiledCaseCount
`func (o *BiomarkerSamplePrevalenceRow) UnsetSsmProfiledCaseCount()`

UnsetSsmProfiledCaseCount ensures that no value is present for SsmProfiledCaseCount, not even an explicit nil
### GetAllCaseCount

`func (o *BiomarkerSamplePrevalenceRow) GetAllCaseCount() int64`

GetAllCaseCount returns the AllCaseCount field if non-nil, zero value otherwise.

### GetAllCaseCountOk

`func (o *BiomarkerSamplePrevalenceRow) GetAllCaseCountOk() (*int64, bool)`

GetAllCaseCountOk returns a tuple with the AllCaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllCaseCount

`func (o *BiomarkerSamplePrevalenceRow) SetAllCaseCount(v int64)`

SetAllCaseCount sets AllCaseCount field to given value.

### HasAllCaseCount

`func (o *BiomarkerSamplePrevalenceRow) HasAllCaseCount() bool`

HasAllCaseCount returns a boolean if a field has been set.

### SetAllCaseCountNil

`func (o *BiomarkerSamplePrevalenceRow) SetAllCaseCountNil(b bool)`

 SetAllCaseCountNil sets the value for AllCaseCount to be an explicit nil

### UnsetAllCaseCount
`func (o *BiomarkerSamplePrevalenceRow) UnsetAllCaseCount()`

UnsetAllCaseCount ensures that no value is present for AllCaseCount, not even an explicit nil
### GetObservedFrequency

`func (o *BiomarkerSamplePrevalenceRow) GetObservedFrequency() float32`

GetObservedFrequency returns the ObservedFrequency field if non-nil, zero value otherwise.

### GetObservedFrequencyOk

`func (o *BiomarkerSamplePrevalenceRow) GetObservedFrequencyOk() (*float32, bool)`

GetObservedFrequencyOk returns a tuple with the ObservedFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedFrequency

`func (o *BiomarkerSamplePrevalenceRow) SetObservedFrequency(v float32)`

SetObservedFrequency sets ObservedFrequency field to given value.

### HasObservedFrequency

`func (o *BiomarkerSamplePrevalenceRow) HasObservedFrequency() bool`

HasObservedFrequency returns a boolean if a field has been set.

### SetObservedFrequencyNil

`func (o *BiomarkerSamplePrevalenceRow) SetObservedFrequencyNil(b bool)`

 SetObservedFrequencyNil sets the value for ObservedFrequency to be an explicit nil

### UnsetObservedFrequency
`func (o *BiomarkerSamplePrevalenceRow) UnsetObservedFrequency()`

UnsetObservedFrequency ensures that no value is present for ObservedFrequency, not even an explicit nil
### GetFrequencyCi95Low

`func (o *BiomarkerSamplePrevalenceRow) GetFrequencyCi95Low() float32`

GetFrequencyCi95Low returns the FrequencyCi95Low field if non-nil, zero value otherwise.

### GetFrequencyCi95LowOk

`func (o *BiomarkerSamplePrevalenceRow) GetFrequencyCi95LowOk() (*float32, bool)`

GetFrequencyCi95LowOk returns a tuple with the FrequencyCi95Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyCi95Low

`func (o *BiomarkerSamplePrevalenceRow) SetFrequencyCi95Low(v float32)`

SetFrequencyCi95Low sets FrequencyCi95Low field to given value.

### HasFrequencyCi95Low

`func (o *BiomarkerSamplePrevalenceRow) HasFrequencyCi95Low() bool`

HasFrequencyCi95Low returns a boolean if a field has been set.

### SetFrequencyCi95LowNil

`func (o *BiomarkerSamplePrevalenceRow) SetFrequencyCi95LowNil(b bool)`

 SetFrequencyCi95LowNil sets the value for FrequencyCi95Low to be an explicit nil

### UnsetFrequencyCi95Low
`func (o *BiomarkerSamplePrevalenceRow) UnsetFrequencyCi95Low()`

UnsetFrequencyCi95Low ensures that no value is present for FrequencyCi95Low, not even an explicit nil
### GetFrequencyCi95High

`func (o *BiomarkerSamplePrevalenceRow) GetFrequencyCi95High() float32`

GetFrequencyCi95High returns the FrequencyCi95High field if non-nil, zero value otherwise.

### GetFrequencyCi95HighOk

`func (o *BiomarkerSamplePrevalenceRow) GetFrequencyCi95HighOk() (*float32, bool)`

GetFrequencyCi95HighOk returns a tuple with the FrequencyCi95High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyCi95High

`func (o *BiomarkerSamplePrevalenceRow) SetFrequencyCi95High(v float32)`

SetFrequencyCi95High sets FrequencyCi95High field to given value.

### HasFrequencyCi95High

`func (o *BiomarkerSamplePrevalenceRow) HasFrequencyCi95High() bool`

HasFrequencyCi95High returns a boolean if a field has been set.

### SetFrequencyCi95HighNil

`func (o *BiomarkerSamplePrevalenceRow) SetFrequencyCi95HighNil(b bool)`

 SetFrequencyCi95HighNil sets the value for FrequencyCi95High to be an explicit nil

### UnsetFrequencyCi95High
`func (o *BiomarkerSamplePrevalenceRow) UnsetFrequencyCi95High()`

UnsetFrequencyCi95High ensures that no value is present for FrequencyCi95High, not even an explicit nil
### GetGdcDataRelease

`func (o *BiomarkerSamplePrevalenceRow) GetGdcDataRelease() string`

GetGdcDataRelease returns the GdcDataRelease field if non-nil, zero value otherwise.

### GetGdcDataReleaseOk

`func (o *BiomarkerSamplePrevalenceRow) GetGdcDataReleaseOk() (*string, bool)`

GetGdcDataReleaseOk returns a tuple with the GdcDataRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdcDataRelease

`func (o *BiomarkerSamplePrevalenceRow) SetGdcDataRelease(v string)`

SetGdcDataRelease sets GdcDataRelease field to given value.

### HasGdcDataRelease

`func (o *BiomarkerSamplePrevalenceRow) HasGdcDataRelease() bool`

HasGdcDataRelease returns a boolean if a field has been set.

### SetGdcDataReleaseNil

`func (o *BiomarkerSamplePrevalenceRow) SetGdcDataReleaseNil(b bool)`

 SetGdcDataReleaseNil sets the value for GdcDataRelease to be an explicit nil

### UnsetGdcDataRelease
`func (o *BiomarkerSamplePrevalenceRow) UnsetGdcDataRelease()`

UnsetGdcDataRelease ensures that no value is present for GdcDataRelease, not even an explicit nil
### GetClaimType

`func (o *BiomarkerSamplePrevalenceRow) GetClaimType() string`

GetClaimType returns the ClaimType field if non-nil, zero value otherwise.

### GetClaimTypeOk

`func (o *BiomarkerSamplePrevalenceRow) GetClaimTypeOk() (*string, bool)`

GetClaimTypeOk returns a tuple with the ClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimType

`func (o *BiomarkerSamplePrevalenceRow) SetClaimType(v string)`

SetClaimType sets ClaimType field to given value.

### HasClaimType

`func (o *BiomarkerSamplePrevalenceRow) HasClaimType() bool`

HasClaimType returns a boolean if a field has been set.

### SetClaimTypeNil

`func (o *BiomarkerSamplePrevalenceRow) SetClaimTypeNil(b bool)`

 SetClaimTypeNil sets the value for ClaimType to be an explicit nil

### UnsetClaimType
`func (o *BiomarkerSamplePrevalenceRow) UnsetClaimType()`

UnsetClaimType ensures that no value is present for ClaimType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


