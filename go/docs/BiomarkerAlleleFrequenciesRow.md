# BiomarkerAlleleFrequenciesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteKey** | Pointer to **NullableString** | The route key of the gene the variant belongs to, the entry key for the biomarkers object type. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol. | [optional] 
**VariantRouteKey** | Pointer to **NullableString** | The variant&#39;s own route key (the id its web page uses). | [optional] 
**VariantBiomarkerId** | Pointer to **NullableString** | The variant&#39;s own biomarker key. | [optional] 
**AminoAcidChange** | Pointer to **NullableString** | The protein change, for example p.V600E. | [optional] 
**BiomarkerVariantCoordinateId** | Pointer to **NullableString** | The key of the DNA change (the coordinate path) the row describes. | [optional] 
**SnvOrdinal** | Pointer to **NullableInt32** | Which single-base DNA change of the protein change this is, from one; most changes have one. | [optional] 
**IsAmbiguousPath** | Pointer to **NullableBool** | Whether the protein change can be produced by more than one single-base DNA change, so this row is one candidate rather than the only one. | [optional] 
**SnvPathCount** | Pointer to **NullableInt32** | How many single-base DNA changes produce the protein change. | [optional] 
**RefseqTranscriptAccession** | Pointer to **NullableString** | The reference transcript the protein change was mapped on (a RefSeq accession). | [optional] 
**Strand** | Pointer to **NullableString** | The strand of the gene on the genome, plus or minus. | [optional] 
**QueriedContig** | Pointer to **NullableString** | The chromosome as gnomAD spells it. | [optional] 
**QueriedPosition** | Pointer to **NullableInt64** | The genomic position looked up in gnomAD (GRCh38). | [optional] 
**ReferenceAllele** | Pointer to **NullableString** | The reference base on the genomic plus strand. | [optional] 
**AlternateAllele** | Pointer to **NullableString** | The alternate base on the genomic plus strand. | [optional] 
**GnomadRelease** | Pointer to **NullableString** | The gnomAD release the lookup used. | [optional] 
**CoverageStatus** | Pointer to **NullableString** | What the lookup found. observed: the DNA change is in gnomAD. absent_from_gnomad: the position was queried and the change is not there. no_gnomad_release_for_contig: gnomAD has no file for the chromosome. query_failed: the lookup failed. | [optional] 
**FilterStatus** | Pointer to **NullableString** | gnomAD&#39;s own quality filter for the site, verbatim (PASS or a semicolon list). Null unless observed. | [optional] 
**PopulationCode** | Pointer to **NullableString** | The population. overall is the unstratified total; the others are gnomAD genetic ancestry groups: afr African/African American, amr Admixed American, asj Ashkenazi Jewish, eas East Asian, fin Finnish, mid Middle Eastern, nfe Non-Finnish European, remaining the remaining individuals, sas South Asian. | [optional] 
**AlleleCount** | Pointer to **NullableInt64** | Alternate alleles observed in the population. Null unless observed. | [optional] 
**AlleleNumber** | Pointer to **NullableInt64** | Alleles called in the population, the denominator of allele_frequency. Null unless observed. | [optional] 
**AlleleFrequency** | Pointer to **NullableFloat64** | The share of called alleles that are the alternate allele (allele_count over allele_number), as a fraction from zero to one. Null unless observed. | [optional] 
**HomozygoteCount** | Pointer to **NullableInt64** | People homozygous for the alternate allele. Null unless observed. | [optional] 

## Methods

### NewBiomarkerAlleleFrequenciesRow

`func NewBiomarkerAlleleFrequenciesRow() *BiomarkerAlleleFrequenciesRow`

NewBiomarkerAlleleFrequenciesRow instantiates a new BiomarkerAlleleFrequenciesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerAlleleFrequenciesRowWithDefaults

`func NewBiomarkerAlleleFrequenciesRowWithDefaults() *BiomarkerAlleleFrequenciesRow`

NewBiomarkerAlleleFrequenciesRowWithDefaults instantiates a new BiomarkerAlleleFrequenciesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteKey

`func (o *BiomarkerAlleleFrequenciesRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerAlleleFrequenciesRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerAlleleFrequenciesRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerAlleleFrequenciesRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerAlleleFrequenciesRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerAlleleFrequenciesRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetGeneSymbol

`func (o *BiomarkerAlleleFrequenciesRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *BiomarkerAlleleFrequenciesRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *BiomarkerAlleleFrequenciesRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *BiomarkerAlleleFrequenciesRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *BiomarkerAlleleFrequenciesRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *BiomarkerAlleleFrequenciesRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetVariantRouteKey

`func (o *BiomarkerAlleleFrequenciesRow) GetVariantRouteKey() string`

GetVariantRouteKey returns the VariantRouteKey field if non-nil, zero value otherwise.

### GetVariantRouteKeyOk

`func (o *BiomarkerAlleleFrequenciesRow) GetVariantRouteKeyOk() (*string, bool)`

GetVariantRouteKeyOk returns a tuple with the VariantRouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantRouteKey

`func (o *BiomarkerAlleleFrequenciesRow) SetVariantRouteKey(v string)`

SetVariantRouteKey sets VariantRouteKey field to given value.

### HasVariantRouteKey

`func (o *BiomarkerAlleleFrequenciesRow) HasVariantRouteKey() bool`

HasVariantRouteKey returns a boolean if a field has been set.

### SetVariantRouteKeyNil

`func (o *BiomarkerAlleleFrequenciesRow) SetVariantRouteKeyNil(b bool)`

 SetVariantRouteKeyNil sets the value for VariantRouteKey to be an explicit nil

### UnsetVariantRouteKey
`func (o *BiomarkerAlleleFrequenciesRow) UnsetVariantRouteKey()`

UnsetVariantRouteKey ensures that no value is present for VariantRouteKey, not even an explicit nil
### GetVariantBiomarkerId

`func (o *BiomarkerAlleleFrequenciesRow) GetVariantBiomarkerId() string`

GetVariantBiomarkerId returns the VariantBiomarkerId field if non-nil, zero value otherwise.

### GetVariantBiomarkerIdOk

`func (o *BiomarkerAlleleFrequenciesRow) GetVariantBiomarkerIdOk() (*string, bool)`

GetVariantBiomarkerIdOk returns a tuple with the VariantBiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantBiomarkerId

`func (o *BiomarkerAlleleFrequenciesRow) SetVariantBiomarkerId(v string)`

SetVariantBiomarkerId sets VariantBiomarkerId field to given value.

### HasVariantBiomarkerId

`func (o *BiomarkerAlleleFrequenciesRow) HasVariantBiomarkerId() bool`

HasVariantBiomarkerId returns a boolean if a field has been set.

### SetVariantBiomarkerIdNil

`func (o *BiomarkerAlleleFrequenciesRow) SetVariantBiomarkerIdNil(b bool)`

 SetVariantBiomarkerIdNil sets the value for VariantBiomarkerId to be an explicit nil

### UnsetVariantBiomarkerId
`func (o *BiomarkerAlleleFrequenciesRow) UnsetVariantBiomarkerId()`

UnsetVariantBiomarkerId ensures that no value is present for VariantBiomarkerId, not even an explicit nil
### GetAminoAcidChange

`func (o *BiomarkerAlleleFrequenciesRow) GetAminoAcidChange() string`

GetAminoAcidChange returns the AminoAcidChange field if non-nil, zero value otherwise.

### GetAminoAcidChangeOk

`func (o *BiomarkerAlleleFrequenciesRow) GetAminoAcidChangeOk() (*string, bool)`

GetAminoAcidChangeOk returns a tuple with the AminoAcidChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcidChange

`func (o *BiomarkerAlleleFrequenciesRow) SetAminoAcidChange(v string)`

SetAminoAcidChange sets AminoAcidChange field to given value.

### HasAminoAcidChange

`func (o *BiomarkerAlleleFrequenciesRow) HasAminoAcidChange() bool`

HasAminoAcidChange returns a boolean if a field has been set.

### SetAminoAcidChangeNil

`func (o *BiomarkerAlleleFrequenciesRow) SetAminoAcidChangeNil(b bool)`

 SetAminoAcidChangeNil sets the value for AminoAcidChange to be an explicit nil

### UnsetAminoAcidChange
`func (o *BiomarkerAlleleFrequenciesRow) UnsetAminoAcidChange()`

UnsetAminoAcidChange ensures that no value is present for AminoAcidChange, not even an explicit nil
### GetBiomarkerVariantCoordinateId

`func (o *BiomarkerAlleleFrequenciesRow) GetBiomarkerVariantCoordinateId() string`

GetBiomarkerVariantCoordinateId returns the BiomarkerVariantCoordinateId field if non-nil, zero value otherwise.

### GetBiomarkerVariantCoordinateIdOk

`func (o *BiomarkerAlleleFrequenciesRow) GetBiomarkerVariantCoordinateIdOk() (*string, bool)`

GetBiomarkerVariantCoordinateIdOk returns a tuple with the BiomarkerVariantCoordinateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerVariantCoordinateId

`func (o *BiomarkerAlleleFrequenciesRow) SetBiomarkerVariantCoordinateId(v string)`

SetBiomarkerVariantCoordinateId sets BiomarkerVariantCoordinateId field to given value.

### HasBiomarkerVariantCoordinateId

`func (o *BiomarkerAlleleFrequenciesRow) HasBiomarkerVariantCoordinateId() bool`

HasBiomarkerVariantCoordinateId returns a boolean if a field has been set.

### SetBiomarkerVariantCoordinateIdNil

`func (o *BiomarkerAlleleFrequenciesRow) SetBiomarkerVariantCoordinateIdNil(b bool)`

 SetBiomarkerVariantCoordinateIdNil sets the value for BiomarkerVariantCoordinateId to be an explicit nil

### UnsetBiomarkerVariantCoordinateId
`func (o *BiomarkerAlleleFrequenciesRow) UnsetBiomarkerVariantCoordinateId()`

UnsetBiomarkerVariantCoordinateId ensures that no value is present for BiomarkerVariantCoordinateId, not even an explicit nil
### GetSnvOrdinal

`func (o *BiomarkerAlleleFrequenciesRow) GetSnvOrdinal() int32`

GetSnvOrdinal returns the SnvOrdinal field if non-nil, zero value otherwise.

### GetSnvOrdinalOk

`func (o *BiomarkerAlleleFrequenciesRow) GetSnvOrdinalOk() (*int32, bool)`

GetSnvOrdinalOk returns a tuple with the SnvOrdinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnvOrdinal

`func (o *BiomarkerAlleleFrequenciesRow) SetSnvOrdinal(v int32)`

SetSnvOrdinal sets SnvOrdinal field to given value.

### HasSnvOrdinal

`func (o *BiomarkerAlleleFrequenciesRow) HasSnvOrdinal() bool`

HasSnvOrdinal returns a boolean if a field has been set.

### SetSnvOrdinalNil

`func (o *BiomarkerAlleleFrequenciesRow) SetSnvOrdinalNil(b bool)`

 SetSnvOrdinalNil sets the value for SnvOrdinal to be an explicit nil

### UnsetSnvOrdinal
`func (o *BiomarkerAlleleFrequenciesRow) UnsetSnvOrdinal()`

UnsetSnvOrdinal ensures that no value is present for SnvOrdinal, not even an explicit nil
### GetIsAmbiguousPath

`func (o *BiomarkerAlleleFrequenciesRow) GetIsAmbiguousPath() bool`

GetIsAmbiguousPath returns the IsAmbiguousPath field if non-nil, zero value otherwise.

### GetIsAmbiguousPathOk

`func (o *BiomarkerAlleleFrequenciesRow) GetIsAmbiguousPathOk() (*bool, bool)`

GetIsAmbiguousPathOk returns a tuple with the IsAmbiguousPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAmbiguousPath

`func (o *BiomarkerAlleleFrequenciesRow) SetIsAmbiguousPath(v bool)`

SetIsAmbiguousPath sets IsAmbiguousPath field to given value.

### HasIsAmbiguousPath

`func (o *BiomarkerAlleleFrequenciesRow) HasIsAmbiguousPath() bool`

HasIsAmbiguousPath returns a boolean if a field has been set.

### SetIsAmbiguousPathNil

`func (o *BiomarkerAlleleFrequenciesRow) SetIsAmbiguousPathNil(b bool)`

 SetIsAmbiguousPathNil sets the value for IsAmbiguousPath to be an explicit nil

### UnsetIsAmbiguousPath
`func (o *BiomarkerAlleleFrequenciesRow) UnsetIsAmbiguousPath()`

UnsetIsAmbiguousPath ensures that no value is present for IsAmbiguousPath, not even an explicit nil
### GetSnvPathCount

`func (o *BiomarkerAlleleFrequenciesRow) GetSnvPathCount() int32`

GetSnvPathCount returns the SnvPathCount field if non-nil, zero value otherwise.

### GetSnvPathCountOk

`func (o *BiomarkerAlleleFrequenciesRow) GetSnvPathCountOk() (*int32, bool)`

GetSnvPathCountOk returns a tuple with the SnvPathCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnvPathCount

`func (o *BiomarkerAlleleFrequenciesRow) SetSnvPathCount(v int32)`

SetSnvPathCount sets SnvPathCount field to given value.

### HasSnvPathCount

`func (o *BiomarkerAlleleFrequenciesRow) HasSnvPathCount() bool`

HasSnvPathCount returns a boolean if a field has been set.

### SetSnvPathCountNil

`func (o *BiomarkerAlleleFrequenciesRow) SetSnvPathCountNil(b bool)`

 SetSnvPathCountNil sets the value for SnvPathCount to be an explicit nil

### UnsetSnvPathCount
`func (o *BiomarkerAlleleFrequenciesRow) UnsetSnvPathCount()`

UnsetSnvPathCount ensures that no value is present for SnvPathCount, not even an explicit nil
### GetRefseqTranscriptAccession

`func (o *BiomarkerAlleleFrequenciesRow) GetRefseqTranscriptAccession() string`

GetRefseqTranscriptAccession returns the RefseqTranscriptAccession field if non-nil, zero value otherwise.

### GetRefseqTranscriptAccessionOk

`func (o *BiomarkerAlleleFrequenciesRow) GetRefseqTranscriptAccessionOk() (*string, bool)`

GetRefseqTranscriptAccessionOk returns a tuple with the RefseqTranscriptAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqTranscriptAccession

`func (o *BiomarkerAlleleFrequenciesRow) SetRefseqTranscriptAccession(v string)`

SetRefseqTranscriptAccession sets RefseqTranscriptAccession field to given value.

### HasRefseqTranscriptAccession

`func (o *BiomarkerAlleleFrequenciesRow) HasRefseqTranscriptAccession() bool`

HasRefseqTranscriptAccession returns a boolean if a field has been set.

### SetRefseqTranscriptAccessionNil

`func (o *BiomarkerAlleleFrequenciesRow) SetRefseqTranscriptAccessionNil(b bool)`

 SetRefseqTranscriptAccessionNil sets the value for RefseqTranscriptAccession to be an explicit nil

### UnsetRefseqTranscriptAccession
`func (o *BiomarkerAlleleFrequenciesRow) UnsetRefseqTranscriptAccession()`

UnsetRefseqTranscriptAccession ensures that no value is present for RefseqTranscriptAccession, not even an explicit nil
### GetStrand

`func (o *BiomarkerAlleleFrequenciesRow) GetStrand() string`

GetStrand returns the Strand field if non-nil, zero value otherwise.

### GetStrandOk

`func (o *BiomarkerAlleleFrequenciesRow) GetStrandOk() (*string, bool)`

GetStrandOk returns a tuple with the Strand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrand

`func (o *BiomarkerAlleleFrequenciesRow) SetStrand(v string)`

SetStrand sets Strand field to given value.

### HasStrand

`func (o *BiomarkerAlleleFrequenciesRow) HasStrand() bool`

HasStrand returns a boolean if a field has been set.

### SetStrandNil

`func (o *BiomarkerAlleleFrequenciesRow) SetStrandNil(b bool)`

 SetStrandNil sets the value for Strand to be an explicit nil

### UnsetStrand
`func (o *BiomarkerAlleleFrequenciesRow) UnsetStrand()`

UnsetStrand ensures that no value is present for Strand, not even an explicit nil
### GetQueriedContig

`func (o *BiomarkerAlleleFrequenciesRow) GetQueriedContig() string`

GetQueriedContig returns the QueriedContig field if non-nil, zero value otherwise.

### GetQueriedContigOk

`func (o *BiomarkerAlleleFrequenciesRow) GetQueriedContigOk() (*string, bool)`

GetQueriedContigOk returns a tuple with the QueriedContig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueriedContig

`func (o *BiomarkerAlleleFrequenciesRow) SetQueriedContig(v string)`

SetQueriedContig sets QueriedContig field to given value.

### HasQueriedContig

`func (o *BiomarkerAlleleFrequenciesRow) HasQueriedContig() bool`

HasQueriedContig returns a boolean if a field has been set.

### SetQueriedContigNil

`func (o *BiomarkerAlleleFrequenciesRow) SetQueriedContigNil(b bool)`

 SetQueriedContigNil sets the value for QueriedContig to be an explicit nil

### UnsetQueriedContig
`func (o *BiomarkerAlleleFrequenciesRow) UnsetQueriedContig()`

UnsetQueriedContig ensures that no value is present for QueriedContig, not even an explicit nil
### GetQueriedPosition

`func (o *BiomarkerAlleleFrequenciesRow) GetQueriedPosition() int64`

GetQueriedPosition returns the QueriedPosition field if non-nil, zero value otherwise.

### GetQueriedPositionOk

`func (o *BiomarkerAlleleFrequenciesRow) GetQueriedPositionOk() (*int64, bool)`

GetQueriedPositionOk returns a tuple with the QueriedPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueriedPosition

`func (o *BiomarkerAlleleFrequenciesRow) SetQueriedPosition(v int64)`

SetQueriedPosition sets QueriedPosition field to given value.

### HasQueriedPosition

`func (o *BiomarkerAlleleFrequenciesRow) HasQueriedPosition() bool`

HasQueriedPosition returns a boolean if a field has been set.

### SetQueriedPositionNil

`func (o *BiomarkerAlleleFrequenciesRow) SetQueriedPositionNil(b bool)`

 SetQueriedPositionNil sets the value for QueriedPosition to be an explicit nil

### UnsetQueriedPosition
`func (o *BiomarkerAlleleFrequenciesRow) UnsetQueriedPosition()`

UnsetQueriedPosition ensures that no value is present for QueriedPosition, not even an explicit nil
### GetReferenceAllele

`func (o *BiomarkerAlleleFrequenciesRow) GetReferenceAllele() string`

GetReferenceAllele returns the ReferenceAllele field if non-nil, zero value otherwise.

### GetReferenceAlleleOk

`func (o *BiomarkerAlleleFrequenciesRow) GetReferenceAlleleOk() (*string, bool)`

GetReferenceAlleleOk returns a tuple with the ReferenceAllele field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceAllele

`func (o *BiomarkerAlleleFrequenciesRow) SetReferenceAllele(v string)`

SetReferenceAllele sets ReferenceAllele field to given value.

### HasReferenceAllele

`func (o *BiomarkerAlleleFrequenciesRow) HasReferenceAllele() bool`

HasReferenceAllele returns a boolean if a field has been set.

### SetReferenceAlleleNil

`func (o *BiomarkerAlleleFrequenciesRow) SetReferenceAlleleNil(b bool)`

 SetReferenceAlleleNil sets the value for ReferenceAllele to be an explicit nil

### UnsetReferenceAllele
`func (o *BiomarkerAlleleFrequenciesRow) UnsetReferenceAllele()`

UnsetReferenceAllele ensures that no value is present for ReferenceAllele, not even an explicit nil
### GetAlternateAllele

`func (o *BiomarkerAlleleFrequenciesRow) GetAlternateAllele() string`

GetAlternateAllele returns the AlternateAllele field if non-nil, zero value otherwise.

### GetAlternateAlleleOk

`func (o *BiomarkerAlleleFrequenciesRow) GetAlternateAlleleOk() (*string, bool)`

GetAlternateAlleleOk returns a tuple with the AlternateAllele field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateAllele

`func (o *BiomarkerAlleleFrequenciesRow) SetAlternateAllele(v string)`

SetAlternateAllele sets AlternateAllele field to given value.

### HasAlternateAllele

`func (o *BiomarkerAlleleFrequenciesRow) HasAlternateAllele() bool`

HasAlternateAllele returns a boolean if a field has been set.

### SetAlternateAlleleNil

`func (o *BiomarkerAlleleFrequenciesRow) SetAlternateAlleleNil(b bool)`

 SetAlternateAlleleNil sets the value for AlternateAllele to be an explicit nil

### UnsetAlternateAllele
`func (o *BiomarkerAlleleFrequenciesRow) UnsetAlternateAllele()`

UnsetAlternateAllele ensures that no value is present for AlternateAllele, not even an explicit nil
### GetGnomadRelease

`func (o *BiomarkerAlleleFrequenciesRow) GetGnomadRelease() string`

GetGnomadRelease returns the GnomadRelease field if non-nil, zero value otherwise.

### GetGnomadReleaseOk

`func (o *BiomarkerAlleleFrequenciesRow) GetGnomadReleaseOk() (*string, bool)`

GetGnomadReleaseOk returns a tuple with the GnomadRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnomadRelease

`func (o *BiomarkerAlleleFrequenciesRow) SetGnomadRelease(v string)`

SetGnomadRelease sets GnomadRelease field to given value.

### HasGnomadRelease

`func (o *BiomarkerAlleleFrequenciesRow) HasGnomadRelease() bool`

HasGnomadRelease returns a boolean if a field has been set.

### SetGnomadReleaseNil

`func (o *BiomarkerAlleleFrequenciesRow) SetGnomadReleaseNil(b bool)`

 SetGnomadReleaseNil sets the value for GnomadRelease to be an explicit nil

### UnsetGnomadRelease
`func (o *BiomarkerAlleleFrequenciesRow) UnsetGnomadRelease()`

UnsetGnomadRelease ensures that no value is present for GnomadRelease, not even an explicit nil
### GetCoverageStatus

`func (o *BiomarkerAlleleFrequenciesRow) GetCoverageStatus() string`

GetCoverageStatus returns the CoverageStatus field if non-nil, zero value otherwise.

### GetCoverageStatusOk

`func (o *BiomarkerAlleleFrequenciesRow) GetCoverageStatusOk() (*string, bool)`

GetCoverageStatusOk returns a tuple with the CoverageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageStatus

`func (o *BiomarkerAlleleFrequenciesRow) SetCoverageStatus(v string)`

SetCoverageStatus sets CoverageStatus field to given value.

### HasCoverageStatus

`func (o *BiomarkerAlleleFrequenciesRow) HasCoverageStatus() bool`

HasCoverageStatus returns a boolean if a field has been set.

### SetCoverageStatusNil

`func (o *BiomarkerAlleleFrequenciesRow) SetCoverageStatusNil(b bool)`

 SetCoverageStatusNil sets the value for CoverageStatus to be an explicit nil

### UnsetCoverageStatus
`func (o *BiomarkerAlleleFrequenciesRow) UnsetCoverageStatus()`

UnsetCoverageStatus ensures that no value is present for CoverageStatus, not even an explicit nil
### GetFilterStatus

`func (o *BiomarkerAlleleFrequenciesRow) GetFilterStatus() string`

GetFilterStatus returns the FilterStatus field if non-nil, zero value otherwise.

### GetFilterStatusOk

`func (o *BiomarkerAlleleFrequenciesRow) GetFilterStatusOk() (*string, bool)`

GetFilterStatusOk returns a tuple with the FilterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterStatus

`func (o *BiomarkerAlleleFrequenciesRow) SetFilterStatus(v string)`

SetFilterStatus sets FilterStatus field to given value.

### HasFilterStatus

`func (o *BiomarkerAlleleFrequenciesRow) HasFilterStatus() bool`

HasFilterStatus returns a boolean if a field has been set.

### SetFilterStatusNil

`func (o *BiomarkerAlleleFrequenciesRow) SetFilterStatusNil(b bool)`

 SetFilterStatusNil sets the value for FilterStatus to be an explicit nil

### UnsetFilterStatus
`func (o *BiomarkerAlleleFrequenciesRow) UnsetFilterStatus()`

UnsetFilterStatus ensures that no value is present for FilterStatus, not even an explicit nil
### GetPopulationCode

`func (o *BiomarkerAlleleFrequenciesRow) GetPopulationCode() string`

GetPopulationCode returns the PopulationCode field if non-nil, zero value otherwise.

### GetPopulationCodeOk

`func (o *BiomarkerAlleleFrequenciesRow) GetPopulationCodeOk() (*string, bool)`

GetPopulationCodeOk returns a tuple with the PopulationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationCode

`func (o *BiomarkerAlleleFrequenciesRow) SetPopulationCode(v string)`

SetPopulationCode sets PopulationCode field to given value.

### HasPopulationCode

`func (o *BiomarkerAlleleFrequenciesRow) HasPopulationCode() bool`

HasPopulationCode returns a boolean if a field has been set.

### SetPopulationCodeNil

`func (o *BiomarkerAlleleFrequenciesRow) SetPopulationCodeNil(b bool)`

 SetPopulationCodeNil sets the value for PopulationCode to be an explicit nil

### UnsetPopulationCode
`func (o *BiomarkerAlleleFrequenciesRow) UnsetPopulationCode()`

UnsetPopulationCode ensures that no value is present for PopulationCode, not even an explicit nil
### GetAlleleCount

`func (o *BiomarkerAlleleFrequenciesRow) GetAlleleCount() int64`

GetAlleleCount returns the AlleleCount field if non-nil, zero value otherwise.

### GetAlleleCountOk

`func (o *BiomarkerAlleleFrequenciesRow) GetAlleleCountOk() (*int64, bool)`

GetAlleleCountOk returns a tuple with the AlleleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlleleCount

`func (o *BiomarkerAlleleFrequenciesRow) SetAlleleCount(v int64)`

SetAlleleCount sets AlleleCount field to given value.

### HasAlleleCount

`func (o *BiomarkerAlleleFrequenciesRow) HasAlleleCount() bool`

HasAlleleCount returns a boolean if a field has been set.

### SetAlleleCountNil

`func (o *BiomarkerAlleleFrequenciesRow) SetAlleleCountNil(b bool)`

 SetAlleleCountNil sets the value for AlleleCount to be an explicit nil

### UnsetAlleleCount
`func (o *BiomarkerAlleleFrequenciesRow) UnsetAlleleCount()`

UnsetAlleleCount ensures that no value is present for AlleleCount, not even an explicit nil
### GetAlleleNumber

`func (o *BiomarkerAlleleFrequenciesRow) GetAlleleNumber() int64`

GetAlleleNumber returns the AlleleNumber field if non-nil, zero value otherwise.

### GetAlleleNumberOk

`func (o *BiomarkerAlleleFrequenciesRow) GetAlleleNumberOk() (*int64, bool)`

GetAlleleNumberOk returns a tuple with the AlleleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlleleNumber

`func (o *BiomarkerAlleleFrequenciesRow) SetAlleleNumber(v int64)`

SetAlleleNumber sets AlleleNumber field to given value.

### HasAlleleNumber

`func (o *BiomarkerAlleleFrequenciesRow) HasAlleleNumber() bool`

HasAlleleNumber returns a boolean if a field has been set.

### SetAlleleNumberNil

`func (o *BiomarkerAlleleFrequenciesRow) SetAlleleNumberNil(b bool)`

 SetAlleleNumberNil sets the value for AlleleNumber to be an explicit nil

### UnsetAlleleNumber
`func (o *BiomarkerAlleleFrequenciesRow) UnsetAlleleNumber()`

UnsetAlleleNumber ensures that no value is present for AlleleNumber, not even an explicit nil
### GetAlleleFrequency

`func (o *BiomarkerAlleleFrequenciesRow) GetAlleleFrequency() float64`

GetAlleleFrequency returns the AlleleFrequency field if non-nil, zero value otherwise.

### GetAlleleFrequencyOk

`func (o *BiomarkerAlleleFrequenciesRow) GetAlleleFrequencyOk() (*float64, bool)`

GetAlleleFrequencyOk returns a tuple with the AlleleFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlleleFrequency

`func (o *BiomarkerAlleleFrequenciesRow) SetAlleleFrequency(v float64)`

SetAlleleFrequency sets AlleleFrequency field to given value.

### HasAlleleFrequency

`func (o *BiomarkerAlleleFrequenciesRow) HasAlleleFrequency() bool`

HasAlleleFrequency returns a boolean if a field has been set.

### SetAlleleFrequencyNil

`func (o *BiomarkerAlleleFrequenciesRow) SetAlleleFrequencyNil(b bool)`

 SetAlleleFrequencyNil sets the value for AlleleFrequency to be an explicit nil

### UnsetAlleleFrequency
`func (o *BiomarkerAlleleFrequenciesRow) UnsetAlleleFrequency()`

UnsetAlleleFrequency ensures that no value is present for AlleleFrequency, not even an explicit nil
### GetHomozygoteCount

`func (o *BiomarkerAlleleFrequenciesRow) GetHomozygoteCount() int64`

GetHomozygoteCount returns the HomozygoteCount field if non-nil, zero value otherwise.

### GetHomozygoteCountOk

`func (o *BiomarkerAlleleFrequenciesRow) GetHomozygoteCountOk() (*int64, bool)`

GetHomozygoteCountOk returns a tuple with the HomozygoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomozygoteCount

`func (o *BiomarkerAlleleFrequenciesRow) SetHomozygoteCount(v int64)`

SetHomozygoteCount sets HomozygoteCount field to given value.

### HasHomozygoteCount

`func (o *BiomarkerAlleleFrequenciesRow) HasHomozygoteCount() bool`

HasHomozygoteCount returns a boolean if a field has been set.

### SetHomozygoteCountNil

`func (o *BiomarkerAlleleFrequenciesRow) SetHomozygoteCountNil(b bool)`

 SetHomozygoteCountNil sets the value for HomozygoteCount to be an explicit nil

### UnsetHomozygoteCount
`func (o *BiomarkerAlleleFrequenciesRow) UnsetHomozygoteCount()`

UnsetHomozygoteCount ensures that no value is present for HomozygoteCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


