# BiomarkerAssayEvidenceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BiomarkerAssayEvidenceId** | Pointer to **NullableString** | The extracted mention&#39;s id. | [optional] 
**RouteKey** | Pointer to **NullableString** | The route key of the gene the mention resolved to, the entry key for the biomarkers object type. Null when the gene did not resolve. | [optional] 
**BiomarkerId** | Pointer to **NullableString** | The resolved gene&#39;s biomarker key. Null when the gene did not resolve. | [optional] 
**DocumentId** | Pointer to **NullableString** | The id of the source document the facts were extracted from. | [optional] 
**Pmid** | Pointer to **NullableString** | The PubMed id of the paper. | [optional] 
**BiomarkerOrdinal** | Pointer to **NullableInt32** | The position of the mention in the paper&#39;s extracted biomarker list, from 1. | [optional] 
**BiomarkerNameAsExtracted** | Pointer to **NullableString** | The biomarker name as extracted. | [optional] 
**GeneSymbolAsExtracted** | Pointer to **NullableString** | The gene symbol as extracted, before resolution. Null when the mention names no gene. | [optional] 
**BiomarkerSynonyms** | Pointer to **[]string** | The synonyms the paper gives, a document-level list repeated on each of its rows. | [optional] 
**GeneResolutionStatus** | Pointer to **NullableString** | How the mention resolved to a gene. resolved_direct: exact symbol match. resolved_alias: an unambiguous dictionary alias. unresolved: a symbol was extracted that matches nothing or an ambiguous alias; not guessed. not_applicable: no gene symbol was extracted. | [optional] 
**GeneResolutionMethod** | Pointer to **NullableString** | exact_gene_symbol or dictionary_alias; null when the gene did not resolve. | [optional] 
**AssayName** | Pointer to **NullableString** | The assay, trade name or platform the paper names. A document-level value. | [optional] 
**SpecimenType** | Pointer to **NullableString** | The specimen type (blood, plasma, tissue). A document-level value. | [optional] 
**ComparatorMethod** | Pointer to **NullableString** | The reference method the performance was measured against, when reported. | [optional] 
**ReportedPerformance** | Pointer to **NullableString** | The performance as the paper states it (sensitivity, specificity, area under the curve), as text and not computed. | [optional] 
**ValidationTypeMentioned** | Pointer to **NullableString** | What the paper claims about the validation type: a fact about the text, not a tier. | [optional] 
**AssayProviderCompany** | Pointer to **[]string** | Organization names as extracted that run or provide the assay. Not resolved to a company. | [optional] 
**FundingOrganizations** | Pointer to **[]string** | Organization names as extracted that funded the study. Never merged with the assay provider. | [optional] 
**EvidenceType** | Pointer to **NullableString** | Always literature_assay_validation, which tells these rows from the other evidence types. | [optional] 
**EvidenceQuote** | Pointer to **[]string** | The sentences of the paper the facts were read from. Null when the extraction gave none. | [optional] 
**Confidence** | Pointer to **NullableString** | The extraction model&#39;s self-reported confidence. Not a reliability signal: it does not track whether the extraction is correct; never use it as a quality gate. | [optional] 
**Title** | Pointer to **NullableString** | The paper&#39;s title. | [optional] 
**PublishedAt** | Pointer to **NullableTime** | The paper&#39;s publication timestamp, in UTC. Null for about one paper in seven, whose source gives a bare year. | [optional] 
**ExtractionModel** | Pointer to **NullableString** | The model that extracted the facts. | [optional] 
**ExtractedAt** | Pointer to **NullableTime** | When the facts were extracted, in UTC. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The address of the paper the facts were extracted from. | [optional] 

## Methods

### NewBiomarkerAssayEvidenceRow

`func NewBiomarkerAssayEvidenceRow() *BiomarkerAssayEvidenceRow`

NewBiomarkerAssayEvidenceRow instantiates a new BiomarkerAssayEvidenceRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerAssayEvidenceRowWithDefaults

`func NewBiomarkerAssayEvidenceRowWithDefaults() *BiomarkerAssayEvidenceRow`

NewBiomarkerAssayEvidenceRowWithDefaults instantiates a new BiomarkerAssayEvidenceRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiomarkerAssayEvidenceId

`func (o *BiomarkerAssayEvidenceRow) GetBiomarkerAssayEvidenceId() string`

GetBiomarkerAssayEvidenceId returns the BiomarkerAssayEvidenceId field if non-nil, zero value otherwise.

### GetBiomarkerAssayEvidenceIdOk

`func (o *BiomarkerAssayEvidenceRow) GetBiomarkerAssayEvidenceIdOk() (*string, bool)`

GetBiomarkerAssayEvidenceIdOk returns a tuple with the BiomarkerAssayEvidenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerAssayEvidenceId

`func (o *BiomarkerAssayEvidenceRow) SetBiomarkerAssayEvidenceId(v string)`

SetBiomarkerAssayEvidenceId sets BiomarkerAssayEvidenceId field to given value.

### HasBiomarkerAssayEvidenceId

`func (o *BiomarkerAssayEvidenceRow) HasBiomarkerAssayEvidenceId() bool`

HasBiomarkerAssayEvidenceId returns a boolean if a field has been set.

### SetBiomarkerAssayEvidenceIdNil

`func (o *BiomarkerAssayEvidenceRow) SetBiomarkerAssayEvidenceIdNil(b bool)`

 SetBiomarkerAssayEvidenceIdNil sets the value for BiomarkerAssayEvidenceId to be an explicit nil

### UnsetBiomarkerAssayEvidenceId
`func (o *BiomarkerAssayEvidenceRow) UnsetBiomarkerAssayEvidenceId()`

UnsetBiomarkerAssayEvidenceId ensures that no value is present for BiomarkerAssayEvidenceId, not even an explicit nil
### GetRouteKey

`func (o *BiomarkerAssayEvidenceRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerAssayEvidenceRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerAssayEvidenceRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerAssayEvidenceRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerAssayEvidenceRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerAssayEvidenceRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetBiomarkerId

`func (o *BiomarkerAssayEvidenceRow) GetBiomarkerId() string`

GetBiomarkerId returns the BiomarkerId field if non-nil, zero value otherwise.

### GetBiomarkerIdOk

`func (o *BiomarkerAssayEvidenceRow) GetBiomarkerIdOk() (*string, bool)`

GetBiomarkerIdOk returns a tuple with the BiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerId

`func (o *BiomarkerAssayEvidenceRow) SetBiomarkerId(v string)`

SetBiomarkerId sets BiomarkerId field to given value.

### HasBiomarkerId

`func (o *BiomarkerAssayEvidenceRow) HasBiomarkerId() bool`

HasBiomarkerId returns a boolean if a field has been set.

### SetBiomarkerIdNil

`func (o *BiomarkerAssayEvidenceRow) SetBiomarkerIdNil(b bool)`

 SetBiomarkerIdNil sets the value for BiomarkerId to be an explicit nil

### UnsetBiomarkerId
`func (o *BiomarkerAssayEvidenceRow) UnsetBiomarkerId()`

UnsetBiomarkerId ensures that no value is present for BiomarkerId, not even an explicit nil
### GetDocumentId

`func (o *BiomarkerAssayEvidenceRow) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BiomarkerAssayEvidenceRow) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BiomarkerAssayEvidenceRow) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BiomarkerAssayEvidenceRow) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### SetDocumentIdNil

`func (o *BiomarkerAssayEvidenceRow) SetDocumentIdNil(b bool)`

 SetDocumentIdNil sets the value for DocumentId to be an explicit nil

### UnsetDocumentId
`func (o *BiomarkerAssayEvidenceRow) UnsetDocumentId()`

UnsetDocumentId ensures that no value is present for DocumentId, not even an explicit nil
### GetPmid

`func (o *BiomarkerAssayEvidenceRow) GetPmid() string`

GetPmid returns the Pmid field if non-nil, zero value otherwise.

### GetPmidOk

`func (o *BiomarkerAssayEvidenceRow) GetPmidOk() (*string, bool)`

GetPmidOk returns a tuple with the Pmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmid

`func (o *BiomarkerAssayEvidenceRow) SetPmid(v string)`

SetPmid sets Pmid field to given value.

### HasPmid

`func (o *BiomarkerAssayEvidenceRow) HasPmid() bool`

HasPmid returns a boolean if a field has been set.

### SetPmidNil

`func (o *BiomarkerAssayEvidenceRow) SetPmidNil(b bool)`

 SetPmidNil sets the value for Pmid to be an explicit nil

### UnsetPmid
`func (o *BiomarkerAssayEvidenceRow) UnsetPmid()`

UnsetPmid ensures that no value is present for Pmid, not even an explicit nil
### GetBiomarkerOrdinal

`func (o *BiomarkerAssayEvidenceRow) GetBiomarkerOrdinal() int32`

GetBiomarkerOrdinal returns the BiomarkerOrdinal field if non-nil, zero value otherwise.

### GetBiomarkerOrdinalOk

`func (o *BiomarkerAssayEvidenceRow) GetBiomarkerOrdinalOk() (*int32, bool)`

GetBiomarkerOrdinalOk returns a tuple with the BiomarkerOrdinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerOrdinal

`func (o *BiomarkerAssayEvidenceRow) SetBiomarkerOrdinal(v int32)`

SetBiomarkerOrdinal sets BiomarkerOrdinal field to given value.

### HasBiomarkerOrdinal

`func (o *BiomarkerAssayEvidenceRow) HasBiomarkerOrdinal() bool`

HasBiomarkerOrdinal returns a boolean if a field has been set.

### SetBiomarkerOrdinalNil

`func (o *BiomarkerAssayEvidenceRow) SetBiomarkerOrdinalNil(b bool)`

 SetBiomarkerOrdinalNil sets the value for BiomarkerOrdinal to be an explicit nil

### UnsetBiomarkerOrdinal
`func (o *BiomarkerAssayEvidenceRow) UnsetBiomarkerOrdinal()`

UnsetBiomarkerOrdinal ensures that no value is present for BiomarkerOrdinal, not even an explicit nil
### GetBiomarkerNameAsExtracted

`func (o *BiomarkerAssayEvidenceRow) GetBiomarkerNameAsExtracted() string`

GetBiomarkerNameAsExtracted returns the BiomarkerNameAsExtracted field if non-nil, zero value otherwise.

### GetBiomarkerNameAsExtractedOk

`func (o *BiomarkerAssayEvidenceRow) GetBiomarkerNameAsExtractedOk() (*string, bool)`

GetBiomarkerNameAsExtractedOk returns a tuple with the BiomarkerNameAsExtracted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerNameAsExtracted

`func (o *BiomarkerAssayEvidenceRow) SetBiomarkerNameAsExtracted(v string)`

SetBiomarkerNameAsExtracted sets BiomarkerNameAsExtracted field to given value.

### HasBiomarkerNameAsExtracted

`func (o *BiomarkerAssayEvidenceRow) HasBiomarkerNameAsExtracted() bool`

HasBiomarkerNameAsExtracted returns a boolean if a field has been set.

### SetBiomarkerNameAsExtractedNil

`func (o *BiomarkerAssayEvidenceRow) SetBiomarkerNameAsExtractedNil(b bool)`

 SetBiomarkerNameAsExtractedNil sets the value for BiomarkerNameAsExtracted to be an explicit nil

### UnsetBiomarkerNameAsExtracted
`func (o *BiomarkerAssayEvidenceRow) UnsetBiomarkerNameAsExtracted()`

UnsetBiomarkerNameAsExtracted ensures that no value is present for BiomarkerNameAsExtracted, not even an explicit nil
### GetGeneSymbolAsExtracted

`func (o *BiomarkerAssayEvidenceRow) GetGeneSymbolAsExtracted() string`

GetGeneSymbolAsExtracted returns the GeneSymbolAsExtracted field if non-nil, zero value otherwise.

### GetGeneSymbolAsExtractedOk

`func (o *BiomarkerAssayEvidenceRow) GetGeneSymbolAsExtractedOk() (*string, bool)`

GetGeneSymbolAsExtractedOk returns a tuple with the GeneSymbolAsExtracted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbolAsExtracted

`func (o *BiomarkerAssayEvidenceRow) SetGeneSymbolAsExtracted(v string)`

SetGeneSymbolAsExtracted sets GeneSymbolAsExtracted field to given value.

### HasGeneSymbolAsExtracted

`func (o *BiomarkerAssayEvidenceRow) HasGeneSymbolAsExtracted() bool`

HasGeneSymbolAsExtracted returns a boolean if a field has been set.

### SetGeneSymbolAsExtractedNil

`func (o *BiomarkerAssayEvidenceRow) SetGeneSymbolAsExtractedNil(b bool)`

 SetGeneSymbolAsExtractedNil sets the value for GeneSymbolAsExtracted to be an explicit nil

### UnsetGeneSymbolAsExtracted
`func (o *BiomarkerAssayEvidenceRow) UnsetGeneSymbolAsExtracted()`

UnsetGeneSymbolAsExtracted ensures that no value is present for GeneSymbolAsExtracted, not even an explicit nil
### GetBiomarkerSynonyms

`func (o *BiomarkerAssayEvidenceRow) GetBiomarkerSynonyms() []string`

GetBiomarkerSynonyms returns the BiomarkerSynonyms field if non-nil, zero value otherwise.

### GetBiomarkerSynonymsOk

`func (o *BiomarkerAssayEvidenceRow) GetBiomarkerSynonymsOk() (*[]string, bool)`

GetBiomarkerSynonymsOk returns a tuple with the BiomarkerSynonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerSynonyms

`func (o *BiomarkerAssayEvidenceRow) SetBiomarkerSynonyms(v []string)`

SetBiomarkerSynonyms sets BiomarkerSynonyms field to given value.

### HasBiomarkerSynonyms

`func (o *BiomarkerAssayEvidenceRow) HasBiomarkerSynonyms() bool`

HasBiomarkerSynonyms returns a boolean if a field has been set.

### SetBiomarkerSynonymsNil

`func (o *BiomarkerAssayEvidenceRow) SetBiomarkerSynonymsNil(b bool)`

 SetBiomarkerSynonymsNil sets the value for BiomarkerSynonyms to be an explicit nil

### UnsetBiomarkerSynonyms
`func (o *BiomarkerAssayEvidenceRow) UnsetBiomarkerSynonyms()`

UnsetBiomarkerSynonyms ensures that no value is present for BiomarkerSynonyms, not even an explicit nil
### GetGeneResolutionStatus

`func (o *BiomarkerAssayEvidenceRow) GetGeneResolutionStatus() string`

GetGeneResolutionStatus returns the GeneResolutionStatus field if non-nil, zero value otherwise.

### GetGeneResolutionStatusOk

`func (o *BiomarkerAssayEvidenceRow) GetGeneResolutionStatusOk() (*string, bool)`

GetGeneResolutionStatusOk returns a tuple with the GeneResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneResolutionStatus

`func (o *BiomarkerAssayEvidenceRow) SetGeneResolutionStatus(v string)`

SetGeneResolutionStatus sets GeneResolutionStatus field to given value.

### HasGeneResolutionStatus

`func (o *BiomarkerAssayEvidenceRow) HasGeneResolutionStatus() bool`

HasGeneResolutionStatus returns a boolean if a field has been set.

### SetGeneResolutionStatusNil

`func (o *BiomarkerAssayEvidenceRow) SetGeneResolutionStatusNil(b bool)`

 SetGeneResolutionStatusNil sets the value for GeneResolutionStatus to be an explicit nil

### UnsetGeneResolutionStatus
`func (o *BiomarkerAssayEvidenceRow) UnsetGeneResolutionStatus()`

UnsetGeneResolutionStatus ensures that no value is present for GeneResolutionStatus, not even an explicit nil
### GetGeneResolutionMethod

`func (o *BiomarkerAssayEvidenceRow) GetGeneResolutionMethod() string`

GetGeneResolutionMethod returns the GeneResolutionMethod field if non-nil, zero value otherwise.

### GetGeneResolutionMethodOk

`func (o *BiomarkerAssayEvidenceRow) GetGeneResolutionMethodOk() (*string, bool)`

GetGeneResolutionMethodOk returns a tuple with the GeneResolutionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneResolutionMethod

`func (o *BiomarkerAssayEvidenceRow) SetGeneResolutionMethod(v string)`

SetGeneResolutionMethod sets GeneResolutionMethod field to given value.

### HasGeneResolutionMethod

`func (o *BiomarkerAssayEvidenceRow) HasGeneResolutionMethod() bool`

HasGeneResolutionMethod returns a boolean if a field has been set.

### SetGeneResolutionMethodNil

`func (o *BiomarkerAssayEvidenceRow) SetGeneResolutionMethodNil(b bool)`

 SetGeneResolutionMethodNil sets the value for GeneResolutionMethod to be an explicit nil

### UnsetGeneResolutionMethod
`func (o *BiomarkerAssayEvidenceRow) UnsetGeneResolutionMethod()`

UnsetGeneResolutionMethod ensures that no value is present for GeneResolutionMethod, not even an explicit nil
### GetAssayName

`func (o *BiomarkerAssayEvidenceRow) GetAssayName() string`

GetAssayName returns the AssayName field if non-nil, zero value otherwise.

### GetAssayNameOk

`func (o *BiomarkerAssayEvidenceRow) GetAssayNameOk() (*string, bool)`

GetAssayNameOk returns a tuple with the AssayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayName

`func (o *BiomarkerAssayEvidenceRow) SetAssayName(v string)`

SetAssayName sets AssayName field to given value.

### HasAssayName

`func (o *BiomarkerAssayEvidenceRow) HasAssayName() bool`

HasAssayName returns a boolean if a field has been set.

### SetAssayNameNil

`func (o *BiomarkerAssayEvidenceRow) SetAssayNameNil(b bool)`

 SetAssayNameNil sets the value for AssayName to be an explicit nil

### UnsetAssayName
`func (o *BiomarkerAssayEvidenceRow) UnsetAssayName()`

UnsetAssayName ensures that no value is present for AssayName, not even an explicit nil
### GetSpecimenType

`func (o *BiomarkerAssayEvidenceRow) GetSpecimenType() string`

GetSpecimenType returns the SpecimenType field if non-nil, zero value otherwise.

### GetSpecimenTypeOk

`func (o *BiomarkerAssayEvidenceRow) GetSpecimenTypeOk() (*string, bool)`

GetSpecimenTypeOk returns a tuple with the SpecimenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecimenType

`func (o *BiomarkerAssayEvidenceRow) SetSpecimenType(v string)`

SetSpecimenType sets SpecimenType field to given value.

### HasSpecimenType

`func (o *BiomarkerAssayEvidenceRow) HasSpecimenType() bool`

HasSpecimenType returns a boolean if a field has been set.

### SetSpecimenTypeNil

`func (o *BiomarkerAssayEvidenceRow) SetSpecimenTypeNil(b bool)`

 SetSpecimenTypeNil sets the value for SpecimenType to be an explicit nil

### UnsetSpecimenType
`func (o *BiomarkerAssayEvidenceRow) UnsetSpecimenType()`

UnsetSpecimenType ensures that no value is present for SpecimenType, not even an explicit nil
### GetComparatorMethod

`func (o *BiomarkerAssayEvidenceRow) GetComparatorMethod() string`

GetComparatorMethod returns the ComparatorMethod field if non-nil, zero value otherwise.

### GetComparatorMethodOk

`func (o *BiomarkerAssayEvidenceRow) GetComparatorMethodOk() (*string, bool)`

GetComparatorMethodOk returns a tuple with the ComparatorMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparatorMethod

`func (o *BiomarkerAssayEvidenceRow) SetComparatorMethod(v string)`

SetComparatorMethod sets ComparatorMethod field to given value.

### HasComparatorMethod

`func (o *BiomarkerAssayEvidenceRow) HasComparatorMethod() bool`

HasComparatorMethod returns a boolean if a field has been set.

### SetComparatorMethodNil

`func (o *BiomarkerAssayEvidenceRow) SetComparatorMethodNil(b bool)`

 SetComparatorMethodNil sets the value for ComparatorMethod to be an explicit nil

### UnsetComparatorMethod
`func (o *BiomarkerAssayEvidenceRow) UnsetComparatorMethod()`

UnsetComparatorMethod ensures that no value is present for ComparatorMethod, not even an explicit nil
### GetReportedPerformance

`func (o *BiomarkerAssayEvidenceRow) GetReportedPerformance() string`

GetReportedPerformance returns the ReportedPerformance field if non-nil, zero value otherwise.

### GetReportedPerformanceOk

`func (o *BiomarkerAssayEvidenceRow) GetReportedPerformanceOk() (*string, bool)`

GetReportedPerformanceOk returns a tuple with the ReportedPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedPerformance

`func (o *BiomarkerAssayEvidenceRow) SetReportedPerformance(v string)`

SetReportedPerformance sets ReportedPerformance field to given value.

### HasReportedPerformance

`func (o *BiomarkerAssayEvidenceRow) HasReportedPerformance() bool`

HasReportedPerformance returns a boolean if a field has been set.

### SetReportedPerformanceNil

`func (o *BiomarkerAssayEvidenceRow) SetReportedPerformanceNil(b bool)`

 SetReportedPerformanceNil sets the value for ReportedPerformance to be an explicit nil

### UnsetReportedPerformance
`func (o *BiomarkerAssayEvidenceRow) UnsetReportedPerformance()`

UnsetReportedPerformance ensures that no value is present for ReportedPerformance, not even an explicit nil
### GetValidationTypeMentioned

`func (o *BiomarkerAssayEvidenceRow) GetValidationTypeMentioned() string`

GetValidationTypeMentioned returns the ValidationTypeMentioned field if non-nil, zero value otherwise.

### GetValidationTypeMentionedOk

`func (o *BiomarkerAssayEvidenceRow) GetValidationTypeMentionedOk() (*string, bool)`

GetValidationTypeMentionedOk returns a tuple with the ValidationTypeMentioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationTypeMentioned

`func (o *BiomarkerAssayEvidenceRow) SetValidationTypeMentioned(v string)`

SetValidationTypeMentioned sets ValidationTypeMentioned field to given value.

### HasValidationTypeMentioned

`func (o *BiomarkerAssayEvidenceRow) HasValidationTypeMentioned() bool`

HasValidationTypeMentioned returns a boolean if a field has been set.

### SetValidationTypeMentionedNil

`func (o *BiomarkerAssayEvidenceRow) SetValidationTypeMentionedNil(b bool)`

 SetValidationTypeMentionedNil sets the value for ValidationTypeMentioned to be an explicit nil

### UnsetValidationTypeMentioned
`func (o *BiomarkerAssayEvidenceRow) UnsetValidationTypeMentioned()`

UnsetValidationTypeMentioned ensures that no value is present for ValidationTypeMentioned, not even an explicit nil
### GetAssayProviderCompany

`func (o *BiomarkerAssayEvidenceRow) GetAssayProviderCompany() []string`

GetAssayProviderCompany returns the AssayProviderCompany field if non-nil, zero value otherwise.

### GetAssayProviderCompanyOk

`func (o *BiomarkerAssayEvidenceRow) GetAssayProviderCompanyOk() (*[]string, bool)`

GetAssayProviderCompanyOk returns a tuple with the AssayProviderCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayProviderCompany

`func (o *BiomarkerAssayEvidenceRow) SetAssayProviderCompany(v []string)`

SetAssayProviderCompany sets AssayProviderCompany field to given value.

### HasAssayProviderCompany

`func (o *BiomarkerAssayEvidenceRow) HasAssayProviderCompany() bool`

HasAssayProviderCompany returns a boolean if a field has been set.

### SetAssayProviderCompanyNil

`func (o *BiomarkerAssayEvidenceRow) SetAssayProviderCompanyNil(b bool)`

 SetAssayProviderCompanyNil sets the value for AssayProviderCompany to be an explicit nil

### UnsetAssayProviderCompany
`func (o *BiomarkerAssayEvidenceRow) UnsetAssayProviderCompany()`

UnsetAssayProviderCompany ensures that no value is present for AssayProviderCompany, not even an explicit nil
### GetFundingOrganizations

`func (o *BiomarkerAssayEvidenceRow) GetFundingOrganizations() []string`

GetFundingOrganizations returns the FundingOrganizations field if non-nil, zero value otherwise.

### GetFundingOrganizationsOk

`func (o *BiomarkerAssayEvidenceRow) GetFundingOrganizationsOk() (*[]string, bool)`

GetFundingOrganizationsOk returns a tuple with the FundingOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingOrganizations

`func (o *BiomarkerAssayEvidenceRow) SetFundingOrganizations(v []string)`

SetFundingOrganizations sets FundingOrganizations field to given value.

### HasFundingOrganizations

`func (o *BiomarkerAssayEvidenceRow) HasFundingOrganizations() bool`

HasFundingOrganizations returns a boolean if a field has been set.

### SetFundingOrganizationsNil

`func (o *BiomarkerAssayEvidenceRow) SetFundingOrganizationsNil(b bool)`

 SetFundingOrganizationsNil sets the value for FundingOrganizations to be an explicit nil

### UnsetFundingOrganizations
`func (o *BiomarkerAssayEvidenceRow) UnsetFundingOrganizations()`

UnsetFundingOrganizations ensures that no value is present for FundingOrganizations, not even an explicit nil
### GetEvidenceType

`func (o *BiomarkerAssayEvidenceRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *BiomarkerAssayEvidenceRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *BiomarkerAssayEvidenceRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *BiomarkerAssayEvidenceRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *BiomarkerAssayEvidenceRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *BiomarkerAssayEvidenceRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetEvidenceQuote

`func (o *BiomarkerAssayEvidenceRow) GetEvidenceQuote() []string`

GetEvidenceQuote returns the EvidenceQuote field if non-nil, zero value otherwise.

### GetEvidenceQuoteOk

`func (o *BiomarkerAssayEvidenceRow) GetEvidenceQuoteOk() (*[]string, bool)`

GetEvidenceQuoteOk returns a tuple with the EvidenceQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceQuote

`func (o *BiomarkerAssayEvidenceRow) SetEvidenceQuote(v []string)`

SetEvidenceQuote sets EvidenceQuote field to given value.

### HasEvidenceQuote

`func (o *BiomarkerAssayEvidenceRow) HasEvidenceQuote() bool`

HasEvidenceQuote returns a boolean if a field has been set.

### SetEvidenceQuoteNil

`func (o *BiomarkerAssayEvidenceRow) SetEvidenceQuoteNil(b bool)`

 SetEvidenceQuoteNil sets the value for EvidenceQuote to be an explicit nil

### UnsetEvidenceQuote
`func (o *BiomarkerAssayEvidenceRow) UnsetEvidenceQuote()`

UnsetEvidenceQuote ensures that no value is present for EvidenceQuote, not even an explicit nil
### GetConfidence

`func (o *BiomarkerAssayEvidenceRow) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *BiomarkerAssayEvidenceRow) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *BiomarkerAssayEvidenceRow) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *BiomarkerAssayEvidenceRow) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *BiomarkerAssayEvidenceRow) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *BiomarkerAssayEvidenceRow) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetTitle

`func (o *BiomarkerAssayEvidenceRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BiomarkerAssayEvidenceRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BiomarkerAssayEvidenceRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BiomarkerAssayEvidenceRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BiomarkerAssayEvidenceRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BiomarkerAssayEvidenceRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPublishedAt

`func (o *BiomarkerAssayEvidenceRow) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *BiomarkerAssayEvidenceRow) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *BiomarkerAssayEvidenceRow) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *BiomarkerAssayEvidenceRow) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### SetPublishedAtNil

`func (o *BiomarkerAssayEvidenceRow) SetPublishedAtNil(b bool)`

 SetPublishedAtNil sets the value for PublishedAt to be an explicit nil

### UnsetPublishedAt
`func (o *BiomarkerAssayEvidenceRow) UnsetPublishedAt()`

UnsetPublishedAt ensures that no value is present for PublishedAt, not even an explicit nil
### GetExtractionModel

`func (o *BiomarkerAssayEvidenceRow) GetExtractionModel() string`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *BiomarkerAssayEvidenceRow) GetExtractionModelOk() (*string, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *BiomarkerAssayEvidenceRow) SetExtractionModel(v string)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *BiomarkerAssayEvidenceRow) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *BiomarkerAssayEvidenceRow) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *BiomarkerAssayEvidenceRow) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil
### GetExtractedAt

`func (o *BiomarkerAssayEvidenceRow) GetExtractedAt() time.Time`

GetExtractedAt returns the ExtractedAt field if non-nil, zero value otherwise.

### GetExtractedAtOk

`func (o *BiomarkerAssayEvidenceRow) GetExtractedAtOk() (*time.Time, bool)`

GetExtractedAtOk returns a tuple with the ExtractedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedAt

`func (o *BiomarkerAssayEvidenceRow) SetExtractedAt(v time.Time)`

SetExtractedAt sets ExtractedAt field to given value.

### HasExtractedAt

`func (o *BiomarkerAssayEvidenceRow) HasExtractedAt() bool`

HasExtractedAt returns a boolean if a field has been set.

### SetExtractedAtNil

`func (o *BiomarkerAssayEvidenceRow) SetExtractedAtNil(b bool)`

 SetExtractedAtNil sets the value for ExtractedAt to be an explicit nil

### UnsetExtractedAt
`func (o *BiomarkerAssayEvidenceRow) UnsetExtractedAt()`

UnsetExtractedAt ensures that no value is present for ExtractedAt, not even an explicit nil
### GetSourceUrl

`func (o *BiomarkerAssayEvidenceRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *BiomarkerAssayEvidenceRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *BiomarkerAssayEvidenceRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *BiomarkerAssayEvidenceRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *BiomarkerAssayEvidenceRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *BiomarkerAssayEvidenceRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


