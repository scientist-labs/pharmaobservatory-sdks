# BiomarkerSampleCoverageRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteKey** | Pointer to **NullableString** | The route key of the gene the variants belong to, the entry key for the biomarkers object type. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol. | [optional] 
**VariantBiomarkerId** | Pointer to **NullableString** | The variant&#39;s own biomarker key. | [optional] 
**VariantRouteKey** | Pointer to **NullableString** | The variant&#39;s own route key (the id its web page uses). | [optional] 
**AminoAcidChange** | Pointer to **NullableString** | The protein change, for example p.V600E. | [optional] 
**GdcDataRelease** | Pointer to **NullableString** | The Genomic Data Commons data release the match was made against. | [optional] 
**NormalizationStatus** | Pointer to **NullableString** | How the variant matched: resolved_single_ssm, no_exact_ssm, multiple_exact_ssms or source_error. | [optional] 
**MatchCount** | Pointer to **NullableInt32** | Somatic mutations in the catalogue that encode the same protein change. | [optional] 
**ResolutionReason** | Pointer to **NullableString** | The reason for the status, in words. | [optional] 
**MeasuredProjectCount** | Pointer to **NullableInt64** | Research projects with an observed frequency for the variant in biomarker_sample_prevalence. | [optional] 

## Methods

### NewBiomarkerSampleCoverageRow

`func NewBiomarkerSampleCoverageRow() *BiomarkerSampleCoverageRow`

NewBiomarkerSampleCoverageRow instantiates a new BiomarkerSampleCoverageRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerSampleCoverageRowWithDefaults

`func NewBiomarkerSampleCoverageRowWithDefaults() *BiomarkerSampleCoverageRow`

NewBiomarkerSampleCoverageRowWithDefaults instantiates a new BiomarkerSampleCoverageRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteKey

`func (o *BiomarkerSampleCoverageRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerSampleCoverageRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerSampleCoverageRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerSampleCoverageRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerSampleCoverageRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerSampleCoverageRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetGeneSymbol

`func (o *BiomarkerSampleCoverageRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *BiomarkerSampleCoverageRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *BiomarkerSampleCoverageRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *BiomarkerSampleCoverageRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *BiomarkerSampleCoverageRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *BiomarkerSampleCoverageRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetVariantBiomarkerId

`func (o *BiomarkerSampleCoverageRow) GetVariantBiomarkerId() string`

GetVariantBiomarkerId returns the VariantBiomarkerId field if non-nil, zero value otherwise.

### GetVariantBiomarkerIdOk

`func (o *BiomarkerSampleCoverageRow) GetVariantBiomarkerIdOk() (*string, bool)`

GetVariantBiomarkerIdOk returns a tuple with the VariantBiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantBiomarkerId

`func (o *BiomarkerSampleCoverageRow) SetVariantBiomarkerId(v string)`

SetVariantBiomarkerId sets VariantBiomarkerId field to given value.

### HasVariantBiomarkerId

`func (o *BiomarkerSampleCoverageRow) HasVariantBiomarkerId() bool`

HasVariantBiomarkerId returns a boolean if a field has been set.

### SetVariantBiomarkerIdNil

`func (o *BiomarkerSampleCoverageRow) SetVariantBiomarkerIdNil(b bool)`

 SetVariantBiomarkerIdNil sets the value for VariantBiomarkerId to be an explicit nil

### UnsetVariantBiomarkerId
`func (o *BiomarkerSampleCoverageRow) UnsetVariantBiomarkerId()`

UnsetVariantBiomarkerId ensures that no value is present for VariantBiomarkerId, not even an explicit nil
### GetVariantRouteKey

`func (o *BiomarkerSampleCoverageRow) GetVariantRouteKey() string`

GetVariantRouteKey returns the VariantRouteKey field if non-nil, zero value otherwise.

### GetVariantRouteKeyOk

`func (o *BiomarkerSampleCoverageRow) GetVariantRouteKeyOk() (*string, bool)`

GetVariantRouteKeyOk returns a tuple with the VariantRouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantRouteKey

`func (o *BiomarkerSampleCoverageRow) SetVariantRouteKey(v string)`

SetVariantRouteKey sets VariantRouteKey field to given value.

### HasVariantRouteKey

`func (o *BiomarkerSampleCoverageRow) HasVariantRouteKey() bool`

HasVariantRouteKey returns a boolean if a field has been set.

### SetVariantRouteKeyNil

`func (o *BiomarkerSampleCoverageRow) SetVariantRouteKeyNil(b bool)`

 SetVariantRouteKeyNil sets the value for VariantRouteKey to be an explicit nil

### UnsetVariantRouteKey
`func (o *BiomarkerSampleCoverageRow) UnsetVariantRouteKey()`

UnsetVariantRouteKey ensures that no value is present for VariantRouteKey, not even an explicit nil
### GetAminoAcidChange

`func (o *BiomarkerSampleCoverageRow) GetAminoAcidChange() string`

GetAminoAcidChange returns the AminoAcidChange field if non-nil, zero value otherwise.

### GetAminoAcidChangeOk

`func (o *BiomarkerSampleCoverageRow) GetAminoAcidChangeOk() (*string, bool)`

GetAminoAcidChangeOk returns a tuple with the AminoAcidChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcidChange

`func (o *BiomarkerSampleCoverageRow) SetAminoAcidChange(v string)`

SetAminoAcidChange sets AminoAcidChange field to given value.

### HasAminoAcidChange

`func (o *BiomarkerSampleCoverageRow) HasAminoAcidChange() bool`

HasAminoAcidChange returns a boolean if a field has been set.

### SetAminoAcidChangeNil

`func (o *BiomarkerSampleCoverageRow) SetAminoAcidChangeNil(b bool)`

 SetAminoAcidChangeNil sets the value for AminoAcidChange to be an explicit nil

### UnsetAminoAcidChange
`func (o *BiomarkerSampleCoverageRow) UnsetAminoAcidChange()`

UnsetAminoAcidChange ensures that no value is present for AminoAcidChange, not even an explicit nil
### GetGdcDataRelease

`func (o *BiomarkerSampleCoverageRow) GetGdcDataRelease() string`

GetGdcDataRelease returns the GdcDataRelease field if non-nil, zero value otherwise.

### GetGdcDataReleaseOk

`func (o *BiomarkerSampleCoverageRow) GetGdcDataReleaseOk() (*string, bool)`

GetGdcDataReleaseOk returns a tuple with the GdcDataRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdcDataRelease

`func (o *BiomarkerSampleCoverageRow) SetGdcDataRelease(v string)`

SetGdcDataRelease sets GdcDataRelease field to given value.

### HasGdcDataRelease

`func (o *BiomarkerSampleCoverageRow) HasGdcDataRelease() bool`

HasGdcDataRelease returns a boolean if a field has been set.

### SetGdcDataReleaseNil

`func (o *BiomarkerSampleCoverageRow) SetGdcDataReleaseNil(b bool)`

 SetGdcDataReleaseNil sets the value for GdcDataRelease to be an explicit nil

### UnsetGdcDataRelease
`func (o *BiomarkerSampleCoverageRow) UnsetGdcDataRelease()`

UnsetGdcDataRelease ensures that no value is present for GdcDataRelease, not even an explicit nil
### GetNormalizationStatus

`func (o *BiomarkerSampleCoverageRow) GetNormalizationStatus() string`

GetNormalizationStatus returns the NormalizationStatus field if non-nil, zero value otherwise.

### GetNormalizationStatusOk

`func (o *BiomarkerSampleCoverageRow) GetNormalizationStatusOk() (*string, bool)`

GetNormalizationStatusOk returns a tuple with the NormalizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizationStatus

`func (o *BiomarkerSampleCoverageRow) SetNormalizationStatus(v string)`

SetNormalizationStatus sets NormalizationStatus field to given value.

### HasNormalizationStatus

`func (o *BiomarkerSampleCoverageRow) HasNormalizationStatus() bool`

HasNormalizationStatus returns a boolean if a field has been set.

### SetNormalizationStatusNil

`func (o *BiomarkerSampleCoverageRow) SetNormalizationStatusNil(b bool)`

 SetNormalizationStatusNil sets the value for NormalizationStatus to be an explicit nil

### UnsetNormalizationStatus
`func (o *BiomarkerSampleCoverageRow) UnsetNormalizationStatus()`

UnsetNormalizationStatus ensures that no value is present for NormalizationStatus, not even an explicit nil
### GetMatchCount

`func (o *BiomarkerSampleCoverageRow) GetMatchCount() int32`

GetMatchCount returns the MatchCount field if non-nil, zero value otherwise.

### GetMatchCountOk

`func (o *BiomarkerSampleCoverageRow) GetMatchCountOk() (*int32, bool)`

GetMatchCountOk returns a tuple with the MatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCount

`func (o *BiomarkerSampleCoverageRow) SetMatchCount(v int32)`

SetMatchCount sets MatchCount field to given value.

### HasMatchCount

`func (o *BiomarkerSampleCoverageRow) HasMatchCount() bool`

HasMatchCount returns a boolean if a field has been set.

### SetMatchCountNil

`func (o *BiomarkerSampleCoverageRow) SetMatchCountNil(b bool)`

 SetMatchCountNil sets the value for MatchCount to be an explicit nil

### UnsetMatchCount
`func (o *BiomarkerSampleCoverageRow) UnsetMatchCount()`

UnsetMatchCount ensures that no value is present for MatchCount, not even an explicit nil
### GetResolutionReason

`func (o *BiomarkerSampleCoverageRow) GetResolutionReason() string`

GetResolutionReason returns the ResolutionReason field if non-nil, zero value otherwise.

### GetResolutionReasonOk

`func (o *BiomarkerSampleCoverageRow) GetResolutionReasonOk() (*string, bool)`

GetResolutionReasonOk returns a tuple with the ResolutionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionReason

`func (o *BiomarkerSampleCoverageRow) SetResolutionReason(v string)`

SetResolutionReason sets ResolutionReason field to given value.

### HasResolutionReason

`func (o *BiomarkerSampleCoverageRow) HasResolutionReason() bool`

HasResolutionReason returns a boolean if a field has been set.

### SetResolutionReasonNil

`func (o *BiomarkerSampleCoverageRow) SetResolutionReasonNil(b bool)`

 SetResolutionReasonNil sets the value for ResolutionReason to be an explicit nil

### UnsetResolutionReason
`func (o *BiomarkerSampleCoverageRow) UnsetResolutionReason()`

UnsetResolutionReason ensures that no value is present for ResolutionReason, not even an explicit nil
### GetMeasuredProjectCount

`func (o *BiomarkerSampleCoverageRow) GetMeasuredProjectCount() int64`

GetMeasuredProjectCount returns the MeasuredProjectCount field if non-nil, zero value otherwise.

### GetMeasuredProjectCountOk

`func (o *BiomarkerSampleCoverageRow) GetMeasuredProjectCountOk() (*int64, bool)`

GetMeasuredProjectCountOk returns a tuple with the MeasuredProjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredProjectCount

`func (o *BiomarkerSampleCoverageRow) SetMeasuredProjectCount(v int64)`

SetMeasuredProjectCount sets MeasuredProjectCount field to given value.

### HasMeasuredProjectCount

`func (o *BiomarkerSampleCoverageRow) HasMeasuredProjectCount() bool`

HasMeasuredProjectCount returns a boolean if a field has been set.

### SetMeasuredProjectCountNil

`func (o *BiomarkerSampleCoverageRow) SetMeasuredProjectCountNil(b bool)`

 SetMeasuredProjectCountNil sets the value for MeasuredProjectCount to be an explicit nil

### UnsetMeasuredProjectCount
`func (o *BiomarkerSampleCoverageRow) UnsetMeasuredProjectCount()`

UnsetMeasuredProjectCount ensures that no value is present for MeasuredProjectCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


