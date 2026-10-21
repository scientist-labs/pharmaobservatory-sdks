# TargetsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] 
**Symbol** | Pointer to **NullableString** | The target&#39;s HGNC gene symbol; the target page address accepts it. | [optional] 
**Name** | Pointer to **NullableString** | The gene&#39;s name. | [optional] 
**Biotype** | Pointer to **NullableString** | Open Targets biotype of the gene (protein_coding and so on); null where the source gives none. | [optional] 
**UniprotAccession** | Pointer to **NullableString** | The UniProt accession of the gene&#39;s canonical protein; null where there is none. | [optional] 
**ProteinName** | Pointer to **NullableString** | The protein&#39;s recommended name from UniProt; null where the gene has no protein record. | [optional] 
**HgncId** | Pointer to **NullableString** | The HGNC gene id (HGNC:nnnn) from UniProt; null where the gene has no protein record. | [optional] 
**NcbiGeneId** | Pointer to **NullableString** | The NCBI Gene id from UniProt; null where the gene has no protein record. | [optional] 
**HasBiomarker** | Pointer to **NullableBool** | Whether the gene has a biomarker record; the biomarkers datasets enter on the gene symbol. | [optional] 
**NDrugs** | Pointer to **NullableInt64** | Distinct drugs that act on the target, any evidence tier. | [optional] 
**NDrugsCurated** | Pointer to **NullableInt64** | Of n_drugs, those with a curated Open Targets mechanism edge (the floor). | [optional] 
**NDrugsInvestigational** | Pointer to **NullableInt64** | Of n_drugs, those found only through ChEMBL or the curated investigational assets. | [optional] 
**MoaFloorOnly** | Pointer to **NullableBool** | True when no drug beyond the curated floor is known for the target. | [optional] 
**NTrials** | Pointer to **NullableInt64** | Distinct trials whose drug interventions act on the target. | [optional] 
**NTrialsCurated** | Pointer to **NullableInt64** | Of n_trials, those reached through a curated mechanism edge. | [optional] 
**NTrialsInvestigational** | Pointer to **NullableInt64** | Of n_trials, those reached only through ChEMBL or curated investigational assets. | [optional] 
**NPrograms** | Pointer to **NullableFloat32** | Distinct drug and indication programs against the target (a specific indication with at least two trials); the crowding headline. | [optional] 
**NActivePrograms** | Pointer to **NullableFloat32** | Of n_programs, those with no terminated, withdrawn or suspended trial. | [optional] 
**LeadPhase** | Pointer to **NullableInt32** | The furthest phase band (0 to 4) that carries any program; 0 when the target has none. | [optional] 
**ProgramsRank** | Pointer to **NullableInt64** | Rank of n_programs among targets that have programs (1 is the most crowded); null when the target has none. | [optional] 
**ProgramsPctile** | Pointer to **NullableFloat64** | Share of programmed targets with fewer programs than this one, 0 to 1; null when the target has none. | [optional] 
**NApprovedDrugs** | Pointer to **NullableInt64** | Distinct approved drugs with a curated mechanism edge to the target. | [optional] 
**NPublicationMentionPapers** | Pointer to **NullableInt64** | Distinct papers with a gene mention in the trial-cited PubMed corpus. | [optional] 
**HasOtAssociation** | Pointer to **NullableBool** | True when any Open Targets datasource has a target-disease association for the gene. | [optional] 
**NPdb** | Pointer to **NullableInt32** | Protein Data Bank structures mapped to the gene&#39;s protein. | [optional] 
**HasPdb** | Pointer to **NullableBool** | True when at least one Protein Data Bank structure maps to the gene&#39;s protein. | [optional] 

## Methods

### NewTargetsRow

`func NewTargetsRow() *TargetsRow`

NewTargetsRow instantiates a new TargetsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetsRowWithDefaults

`func NewTargetsRowWithDefaults() *TargetsRow`

NewTargetsRowWithDefaults instantiates a new TargetsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *TargetsRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TargetsRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TargetsRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TargetsRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *TargetsRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *TargetsRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetSymbol

`func (o *TargetsRow) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TargetsRow) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TargetsRow) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TargetsRow) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### SetSymbolNil

`func (o *TargetsRow) SetSymbolNil(b bool)`

 SetSymbolNil sets the value for Symbol to be an explicit nil

### UnsetSymbol
`func (o *TargetsRow) UnsetSymbol()`

UnsetSymbol ensures that no value is present for Symbol, not even an explicit nil
### GetName

`func (o *TargetsRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetsRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetsRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TargetsRow) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TargetsRow) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TargetsRow) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetBiotype

`func (o *TargetsRow) GetBiotype() string`

GetBiotype returns the Biotype field if non-nil, zero value otherwise.

### GetBiotypeOk

`func (o *TargetsRow) GetBiotypeOk() (*string, bool)`

GetBiotypeOk returns a tuple with the Biotype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiotype

`func (o *TargetsRow) SetBiotype(v string)`

SetBiotype sets Biotype field to given value.

### HasBiotype

`func (o *TargetsRow) HasBiotype() bool`

HasBiotype returns a boolean if a field has been set.

### SetBiotypeNil

`func (o *TargetsRow) SetBiotypeNil(b bool)`

 SetBiotypeNil sets the value for Biotype to be an explicit nil

### UnsetBiotype
`func (o *TargetsRow) UnsetBiotype()`

UnsetBiotype ensures that no value is present for Biotype, not even an explicit nil
### GetUniprotAccession

`func (o *TargetsRow) GetUniprotAccession() string`

GetUniprotAccession returns the UniprotAccession field if non-nil, zero value otherwise.

### GetUniprotAccessionOk

`func (o *TargetsRow) GetUniprotAccessionOk() (*string, bool)`

GetUniprotAccessionOk returns a tuple with the UniprotAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniprotAccession

`func (o *TargetsRow) SetUniprotAccession(v string)`

SetUniprotAccession sets UniprotAccession field to given value.

### HasUniprotAccession

`func (o *TargetsRow) HasUniprotAccession() bool`

HasUniprotAccession returns a boolean if a field has been set.

### SetUniprotAccessionNil

`func (o *TargetsRow) SetUniprotAccessionNil(b bool)`

 SetUniprotAccessionNil sets the value for UniprotAccession to be an explicit nil

### UnsetUniprotAccession
`func (o *TargetsRow) UnsetUniprotAccession()`

UnsetUniprotAccession ensures that no value is present for UniprotAccession, not even an explicit nil
### GetProteinName

`func (o *TargetsRow) GetProteinName() string`

GetProteinName returns the ProteinName field if non-nil, zero value otherwise.

### GetProteinNameOk

`func (o *TargetsRow) GetProteinNameOk() (*string, bool)`

GetProteinNameOk returns a tuple with the ProteinName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinName

`func (o *TargetsRow) SetProteinName(v string)`

SetProteinName sets ProteinName field to given value.

### HasProteinName

`func (o *TargetsRow) HasProteinName() bool`

HasProteinName returns a boolean if a field has been set.

### SetProteinNameNil

`func (o *TargetsRow) SetProteinNameNil(b bool)`

 SetProteinNameNil sets the value for ProteinName to be an explicit nil

### UnsetProteinName
`func (o *TargetsRow) UnsetProteinName()`

UnsetProteinName ensures that no value is present for ProteinName, not even an explicit nil
### GetHgncId

`func (o *TargetsRow) GetHgncId() string`

GetHgncId returns the HgncId field if non-nil, zero value otherwise.

### GetHgncIdOk

`func (o *TargetsRow) GetHgncIdOk() (*string, bool)`

GetHgncIdOk returns a tuple with the HgncId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgncId

`func (o *TargetsRow) SetHgncId(v string)`

SetHgncId sets HgncId field to given value.

### HasHgncId

`func (o *TargetsRow) HasHgncId() bool`

HasHgncId returns a boolean if a field has been set.

### SetHgncIdNil

`func (o *TargetsRow) SetHgncIdNil(b bool)`

 SetHgncIdNil sets the value for HgncId to be an explicit nil

### UnsetHgncId
`func (o *TargetsRow) UnsetHgncId()`

UnsetHgncId ensures that no value is present for HgncId, not even an explicit nil
### GetNcbiGeneId

`func (o *TargetsRow) GetNcbiGeneId() string`

GetNcbiGeneId returns the NcbiGeneId field if non-nil, zero value otherwise.

### GetNcbiGeneIdOk

`func (o *TargetsRow) GetNcbiGeneIdOk() (*string, bool)`

GetNcbiGeneIdOk returns a tuple with the NcbiGeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcbiGeneId

`func (o *TargetsRow) SetNcbiGeneId(v string)`

SetNcbiGeneId sets NcbiGeneId field to given value.

### HasNcbiGeneId

`func (o *TargetsRow) HasNcbiGeneId() bool`

HasNcbiGeneId returns a boolean if a field has been set.

### SetNcbiGeneIdNil

`func (o *TargetsRow) SetNcbiGeneIdNil(b bool)`

 SetNcbiGeneIdNil sets the value for NcbiGeneId to be an explicit nil

### UnsetNcbiGeneId
`func (o *TargetsRow) UnsetNcbiGeneId()`

UnsetNcbiGeneId ensures that no value is present for NcbiGeneId, not even an explicit nil
### GetHasBiomarker

`func (o *TargetsRow) GetHasBiomarker() bool`

GetHasBiomarker returns the HasBiomarker field if non-nil, zero value otherwise.

### GetHasBiomarkerOk

`func (o *TargetsRow) GetHasBiomarkerOk() (*bool, bool)`

GetHasBiomarkerOk returns a tuple with the HasBiomarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBiomarker

`func (o *TargetsRow) SetHasBiomarker(v bool)`

SetHasBiomarker sets HasBiomarker field to given value.

### HasHasBiomarker

`func (o *TargetsRow) HasHasBiomarker() bool`

HasHasBiomarker returns a boolean if a field has been set.

### SetHasBiomarkerNil

`func (o *TargetsRow) SetHasBiomarkerNil(b bool)`

 SetHasBiomarkerNil sets the value for HasBiomarker to be an explicit nil

### UnsetHasBiomarker
`func (o *TargetsRow) UnsetHasBiomarker()`

UnsetHasBiomarker ensures that no value is present for HasBiomarker, not even an explicit nil
### GetNDrugs

`func (o *TargetsRow) GetNDrugs() int64`

GetNDrugs returns the NDrugs field if non-nil, zero value otherwise.

### GetNDrugsOk

`func (o *TargetsRow) GetNDrugsOk() (*int64, bool)`

GetNDrugsOk returns a tuple with the NDrugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDrugs

`func (o *TargetsRow) SetNDrugs(v int64)`

SetNDrugs sets NDrugs field to given value.

### HasNDrugs

`func (o *TargetsRow) HasNDrugs() bool`

HasNDrugs returns a boolean if a field has been set.

### SetNDrugsNil

`func (o *TargetsRow) SetNDrugsNil(b bool)`

 SetNDrugsNil sets the value for NDrugs to be an explicit nil

### UnsetNDrugs
`func (o *TargetsRow) UnsetNDrugs()`

UnsetNDrugs ensures that no value is present for NDrugs, not even an explicit nil
### GetNDrugsCurated

`func (o *TargetsRow) GetNDrugsCurated() int64`

GetNDrugsCurated returns the NDrugsCurated field if non-nil, zero value otherwise.

### GetNDrugsCuratedOk

`func (o *TargetsRow) GetNDrugsCuratedOk() (*int64, bool)`

GetNDrugsCuratedOk returns a tuple with the NDrugsCurated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDrugsCurated

`func (o *TargetsRow) SetNDrugsCurated(v int64)`

SetNDrugsCurated sets NDrugsCurated field to given value.

### HasNDrugsCurated

`func (o *TargetsRow) HasNDrugsCurated() bool`

HasNDrugsCurated returns a boolean if a field has been set.

### SetNDrugsCuratedNil

`func (o *TargetsRow) SetNDrugsCuratedNil(b bool)`

 SetNDrugsCuratedNil sets the value for NDrugsCurated to be an explicit nil

### UnsetNDrugsCurated
`func (o *TargetsRow) UnsetNDrugsCurated()`

UnsetNDrugsCurated ensures that no value is present for NDrugsCurated, not even an explicit nil
### GetNDrugsInvestigational

`func (o *TargetsRow) GetNDrugsInvestigational() int64`

GetNDrugsInvestigational returns the NDrugsInvestigational field if non-nil, zero value otherwise.

### GetNDrugsInvestigationalOk

`func (o *TargetsRow) GetNDrugsInvestigationalOk() (*int64, bool)`

GetNDrugsInvestigationalOk returns a tuple with the NDrugsInvestigational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDrugsInvestigational

`func (o *TargetsRow) SetNDrugsInvestigational(v int64)`

SetNDrugsInvestigational sets NDrugsInvestigational field to given value.

### HasNDrugsInvestigational

`func (o *TargetsRow) HasNDrugsInvestigational() bool`

HasNDrugsInvestigational returns a boolean if a field has been set.

### SetNDrugsInvestigationalNil

`func (o *TargetsRow) SetNDrugsInvestigationalNil(b bool)`

 SetNDrugsInvestigationalNil sets the value for NDrugsInvestigational to be an explicit nil

### UnsetNDrugsInvestigational
`func (o *TargetsRow) UnsetNDrugsInvestigational()`

UnsetNDrugsInvestigational ensures that no value is present for NDrugsInvestigational, not even an explicit nil
### GetMoaFloorOnly

`func (o *TargetsRow) GetMoaFloorOnly() bool`

GetMoaFloorOnly returns the MoaFloorOnly field if non-nil, zero value otherwise.

### GetMoaFloorOnlyOk

`func (o *TargetsRow) GetMoaFloorOnlyOk() (*bool, bool)`

GetMoaFloorOnlyOk returns a tuple with the MoaFloorOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoaFloorOnly

`func (o *TargetsRow) SetMoaFloorOnly(v bool)`

SetMoaFloorOnly sets MoaFloorOnly field to given value.

### HasMoaFloorOnly

`func (o *TargetsRow) HasMoaFloorOnly() bool`

HasMoaFloorOnly returns a boolean if a field has been set.

### SetMoaFloorOnlyNil

`func (o *TargetsRow) SetMoaFloorOnlyNil(b bool)`

 SetMoaFloorOnlyNil sets the value for MoaFloorOnly to be an explicit nil

### UnsetMoaFloorOnly
`func (o *TargetsRow) UnsetMoaFloorOnly()`

UnsetMoaFloorOnly ensures that no value is present for MoaFloorOnly, not even an explicit nil
### GetNTrials

`func (o *TargetsRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *TargetsRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *TargetsRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *TargetsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *TargetsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *TargetsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNTrialsCurated

`func (o *TargetsRow) GetNTrialsCurated() int64`

GetNTrialsCurated returns the NTrialsCurated field if non-nil, zero value otherwise.

### GetNTrialsCuratedOk

`func (o *TargetsRow) GetNTrialsCuratedOk() (*int64, bool)`

GetNTrialsCuratedOk returns a tuple with the NTrialsCurated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsCurated

`func (o *TargetsRow) SetNTrialsCurated(v int64)`

SetNTrialsCurated sets NTrialsCurated field to given value.

### HasNTrialsCurated

`func (o *TargetsRow) HasNTrialsCurated() bool`

HasNTrialsCurated returns a boolean if a field has been set.

### SetNTrialsCuratedNil

`func (o *TargetsRow) SetNTrialsCuratedNil(b bool)`

 SetNTrialsCuratedNil sets the value for NTrialsCurated to be an explicit nil

### UnsetNTrialsCurated
`func (o *TargetsRow) UnsetNTrialsCurated()`

UnsetNTrialsCurated ensures that no value is present for NTrialsCurated, not even an explicit nil
### GetNTrialsInvestigational

`func (o *TargetsRow) GetNTrialsInvestigational() int64`

GetNTrialsInvestigational returns the NTrialsInvestigational field if non-nil, zero value otherwise.

### GetNTrialsInvestigationalOk

`func (o *TargetsRow) GetNTrialsInvestigationalOk() (*int64, bool)`

GetNTrialsInvestigationalOk returns a tuple with the NTrialsInvestigational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsInvestigational

`func (o *TargetsRow) SetNTrialsInvestigational(v int64)`

SetNTrialsInvestigational sets NTrialsInvestigational field to given value.

### HasNTrialsInvestigational

`func (o *TargetsRow) HasNTrialsInvestigational() bool`

HasNTrialsInvestigational returns a boolean if a field has been set.

### SetNTrialsInvestigationalNil

`func (o *TargetsRow) SetNTrialsInvestigationalNil(b bool)`

 SetNTrialsInvestigationalNil sets the value for NTrialsInvestigational to be an explicit nil

### UnsetNTrialsInvestigational
`func (o *TargetsRow) UnsetNTrialsInvestigational()`

UnsetNTrialsInvestigational ensures that no value is present for NTrialsInvestigational, not even an explicit nil
### GetNPrograms

`func (o *TargetsRow) GetNPrograms() float32`

GetNPrograms returns the NPrograms field if non-nil, zero value otherwise.

### GetNProgramsOk

`func (o *TargetsRow) GetNProgramsOk() (*float32, bool)`

GetNProgramsOk returns a tuple with the NPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPrograms

`func (o *TargetsRow) SetNPrograms(v float32)`

SetNPrograms sets NPrograms field to given value.

### HasNPrograms

`func (o *TargetsRow) HasNPrograms() bool`

HasNPrograms returns a boolean if a field has been set.

### SetNProgramsNil

`func (o *TargetsRow) SetNProgramsNil(b bool)`

 SetNProgramsNil sets the value for NPrograms to be an explicit nil

### UnsetNPrograms
`func (o *TargetsRow) UnsetNPrograms()`

UnsetNPrograms ensures that no value is present for NPrograms, not even an explicit nil
### GetNActivePrograms

`func (o *TargetsRow) GetNActivePrograms() float32`

GetNActivePrograms returns the NActivePrograms field if non-nil, zero value otherwise.

### GetNActiveProgramsOk

`func (o *TargetsRow) GetNActiveProgramsOk() (*float32, bool)`

GetNActiveProgramsOk returns a tuple with the NActivePrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNActivePrograms

`func (o *TargetsRow) SetNActivePrograms(v float32)`

SetNActivePrograms sets NActivePrograms field to given value.

### HasNActivePrograms

`func (o *TargetsRow) HasNActivePrograms() bool`

HasNActivePrograms returns a boolean if a field has been set.

### SetNActiveProgramsNil

`func (o *TargetsRow) SetNActiveProgramsNil(b bool)`

 SetNActiveProgramsNil sets the value for NActivePrograms to be an explicit nil

### UnsetNActivePrograms
`func (o *TargetsRow) UnsetNActivePrograms()`

UnsetNActivePrograms ensures that no value is present for NActivePrograms, not even an explicit nil
### GetLeadPhase

`func (o *TargetsRow) GetLeadPhase() int32`

GetLeadPhase returns the LeadPhase field if non-nil, zero value otherwise.

### GetLeadPhaseOk

`func (o *TargetsRow) GetLeadPhaseOk() (*int32, bool)`

GetLeadPhaseOk returns a tuple with the LeadPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadPhase

`func (o *TargetsRow) SetLeadPhase(v int32)`

SetLeadPhase sets LeadPhase field to given value.

### HasLeadPhase

`func (o *TargetsRow) HasLeadPhase() bool`

HasLeadPhase returns a boolean if a field has been set.

### SetLeadPhaseNil

`func (o *TargetsRow) SetLeadPhaseNil(b bool)`

 SetLeadPhaseNil sets the value for LeadPhase to be an explicit nil

### UnsetLeadPhase
`func (o *TargetsRow) UnsetLeadPhase()`

UnsetLeadPhase ensures that no value is present for LeadPhase, not even an explicit nil
### GetProgramsRank

`func (o *TargetsRow) GetProgramsRank() int64`

GetProgramsRank returns the ProgramsRank field if non-nil, zero value otherwise.

### GetProgramsRankOk

`func (o *TargetsRow) GetProgramsRankOk() (*int64, bool)`

GetProgramsRankOk returns a tuple with the ProgramsRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramsRank

`func (o *TargetsRow) SetProgramsRank(v int64)`

SetProgramsRank sets ProgramsRank field to given value.

### HasProgramsRank

`func (o *TargetsRow) HasProgramsRank() bool`

HasProgramsRank returns a boolean if a field has been set.

### SetProgramsRankNil

`func (o *TargetsRow) SetProgramsRankNil(b bool)`

 SetProgramsRankNil sets the value for ProgramsRank to be an explicit nil

### UnsetProgramsRank
`func (o *TargetsRow) UnsetProgramsRank()`

UnsetProgramsRank ensures that no value is present for ProgramsRank, not even an explicit nil
### GetProgramsPctile

`func (o *TargetsRow) GetProgramsPctile() float64`

GetProgramsPctile returns the ProgramsPctile field if non-nil, zero value otherwise.

### GetProgramsPctileOk

`func (o *TargetsRow) GetProgramsPctileOk() (*float64, bool)`

GetProgramsPctileOk returns a tuple with the ProgramsPctile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramsPctile

`func (o *TargetsRow) SetProgramsPctile(v float64)`

SetProgramsPctile sets ProgramsPctile field to given value.

### HasProgramsPctile

`func (o *TargetsRow) HasProgramsPctile() bool`

HasProgramsPctile returns a boolean if a field has been set.

### SetProgramsPctileNil

`func (o *TargetsRow) SetProgramsPctileNil(b bool)`

 SetProgramsPctileNil sets the value for ProgramsPctile to be an explicit nil

### UnsetProgramsPctile
`func (o *TargetsRow) UnsetProgramsPctile()`

UnsetProgramsPctile ensures that no value is present for ProgramsPctile, not even an explicit nil
### GetNApprovedDrugs

`func (o *TargetsRow) GetNApprovedDrugs() int64`

GetNApprovedDrugs returns the NApprovedDrugs field if non-nil, zero value otherwise.

### GetNApprovedDrugsOk

`func (o *TargetsRow) GetNApprovedDrugsOk() (*int64, bool)`

GetNApprovedDrugsOk returns a tuple with the NApprovedDrugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNApprovedDrugs

`func (o *TargetsRow) SetNApprovedDrugs(v int64)`

SetNApprovedDrugs sets NApprovedDrugs field to given value.

### HasNApprovedDrugs

`func (o *TargetsRow) HasNApprovedDrugs() bool`

HasNApprovedDrugs returns a boolean if a field has been set.

### SetNApprovedDrugsNil

`func (o *TargetsRow) SetNApprovedDrugsNil(b bool)`

 SetNApprovedDrugsNil sets the value for NApprovedDrugs to be an explicit nil

### UnsetNApprovedDrugs
`func (o *TargetsRow) UnsetNApprovedDrugs()`

UnsetNApprovedDrugs ensures that no value is present for NApprovedDrugs, not even an explicit nil
### GetNPublicationMentionPapers

`func (o *TargetsRow) GetNPublicationMentionPapers() int64`

GetNPublicationMentionPapers returns the NPublicationMentionPapers field if non-nil, zero value otherwise.

### GetNPublicationMentionPapersOk

`func (o *TargetsRow) GetNPublicationMentionPapersOk() (*int64, bool)`

GetNPublicationMentionPapersOk returns a tuple with the NPublicationMentionPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPublicationMentionPapers

`func (o *TargetsRow) SetNPublicationMentionPapers(v int64)`

SetNPublicationMentionPapers sets NPublicationMentionPapers field to given value.

### HasNPublicationMentionPapers

`func (o *TargetsRow) HasNPublicationMentionPapers() bool`

HasNPublicationMentionPapers returns a boolean if a field has been set.

### SetNPublicationMentionPapersNil

`func (o *TargetsRow) SetNPublicationMentionPapersNil(b bool)`

 SetNPublicationMentionPapersNil sets the value for NPublicationMentionPapers to be an explicit nil

### UnsetNPublicationMentionPapers
`func (o *TargetsRow) UnsetNPublicationMentionPapers()`

UnsetNPublicationMentionPapers ensures that no value is present for NPublicationMentionPapers, not even an explicit nil
### GetHasOtAssociation

`func (o *TargetsRow) GetHasOtAssociation() bool`

GetHasOtAssociation returns the HasOtAssociation field if non-nil, zero value otherwise.

### GetHasOtAssociationOk

`func (o *TargetsRow) GetHasOtAssociationOk() (*bool, bool)`

GetHasOtAssociationOk returns a tuple with the HasOtAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOtAssociation

`func (o *TargetsRow) SetHasOtAssociation(v bool)`

SetHasOtAssociation sets HasOtAssociation field to given value.

### HasHasOtAssociation

`func (o *TargetsRow) HasHasOtAssociation() bool`

HasHasOtAssociation returns a boolean if a field has been set.

### SetHasOtAssociationNil

`func (o *TargetsRow) SetHasOtAssociationNil(b bool)`

 SetHasOtAssociationNil sets the value for HasOtAssociation to be an explicit nil

### UnsetHasOtAssociation
`func (o *TargetsRow) UnsetHasOtAssociation()`

UnsetHasOtAssociation ensures that no value is present for HasOtAssociation, not even an explicit nil
### GetNPdb

`func (o *TargetsRow) GetNPdb() int32`

GetNPdb returns the NPdb field if non-nil, zero value otherwise.

### GetNPdbOk

`func (o *TargetsRow) GetNPdbOk() (*int32, bool)`

GetNPdbOk returns a tuple with the NPdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPdb

`func (o *TargetsRow) SetNPdb(v int32)`

SetNPdb sets NPdb field to given value.

### HasNPdb

`func (o *TargetsRow) HasNPdb() bool`

HasNPdb returns a boolean if a field has been set.

### SetNPdbNil

`func (o *TargetsRow) SetNPdbNil(b bool)`

 SetNPdbNil sets the value for NPdb to be an explicit nil

### UnsetNPdb
`func (o *TargetsRow) UnsetNPdb()`

UnsetNPdb ensures that no value is present for NPdb, not even an explicit nil
### GetHasPdb

`func (o *TargetsRow) GetHasPdb() bool`

GetHasPdb returns the HasPdb field if non-nil, zero value otherwise.

### GetHasPdbOk

`func (o *TargetsRow) GetHasPdbOk() (*bool, bool)`

GetHasPdbOk returns a tuple with the HasPdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPdb

`func (o *TargetsRow) SetHasPdb(v bool)`

SetHasPdb sets HasPdb field to given value.

### HasHasPdb

`func (o *TargetsRow) HasHasPdb() bool`

HasHasPdb returns a boolean if a field has been set.

### SetHasPdbNil

`func (o *TargetsRow) SetHasPdbNil(b bool)`

 SetHasPdbNil sets the value for HasPdb to be an explicit nil

### UnsetHasPdb
`func (o *TargetsRow) UnsetHasPdb()`

UnsetHasPdb ensures that no value is present for HasPdb, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


