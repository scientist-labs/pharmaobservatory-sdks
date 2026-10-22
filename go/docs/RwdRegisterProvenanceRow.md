# RwdRegisterProvenanceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegisterKey** | Pointer to **NullableString** | Which of the five exports the row concerns: study, data_source_human, data_source_veterinary, institution or network. The entry key of the rwd_registers object type. | [optional] 
**RegisterLabel** | Pointer to **NullableString** | The register&#39;s plain-words label. | [optional] 
**RegisterOrdinal** | Pointer to **NullableInt32** | The register&#39;s position in the lane&#39;s order. | [optional] 
**ProvenanceSource** | Pointer to **NullableString** | The landed source name of the export. | [optional] 
**ExportIsEmptyAtSource** | Pointer to **NullableBool** | True when the export held no rows at the source. | [optional] 
**LandedRows** | Pointer to **NullableInt64** | How many rows the export landed. | [optional] 
**BuildStatus** | Pointer to **NullableString** | The status of the latest landing run. | [optional] 
**BuildSnapshotId** | Pointer to **NullableString** | The snapshot id of the latest landing run. | [optional] 
**BuildRowCount** | Pointer to **NullableInt64** | The row count of the latest landing run. | [optional] 
**BuildStartedAt** | Pointer to **NullableTime** | When the latest landing run started (UTC). | [optional] 
**BuildFinishedAt** | Pointer to **NullableTime** | When the latest landing run finished (UTC). | [optional] 
**LastCompletedSnapshotId** | Pointer to **NullableString** | The snapshot id of the last completed landing. | [optional] 
**LastCompletedRowCount** | Pointer to **NullableInt64** | The row count of the last completed landing. | [optional] 
**LastCompletedAt** | Pointer to **NullableTime** | When the last completed landing finished (UTC). | [optional] 
**RegistrantClockAsof** | Pointer to **NullableString** | The cutoff of the registrant clock: the snapshot date less twelve months; a record last updated before it counts as older than twelve months. | [optional] 
**RegisterPublishesUpdateDate** | Pointer to **NullableBool** | Whether the export carries an update date column. | [optional] 
**RecordsWithAnUpdateDate** | Pointer to **NullableInt64** | How many records carry an update date. | [optional] 
**RecordsUpdatedOverTwelveMonths** | Pointer to **NullableInt64** | How many records were last updated more than twelve months before the snapshot. | [optional] 
**ShareUpdatedOverTwelveMonths** | Pointer to **NullableFloat32** | The share of records with an update date that are older than twelve months. | [optional] 
**RegisterPublishesOutdatedFlag** | Pointer to **NullableBool** | Whether the export carries EMA&#39;s Outdated flag. | [optional] 
**RecordsEmaFlagsOutdated** | Pointer to **NullableInt64** | How many records EMA flags Outdated. | [optional] 
**RecordsOutdatedFlagBlank** | Pointer to **NullableInt64** | How many records have no Outdated flag value. | [optional] 
**ColumnsLanded** | Pointer to **NullableInt64** | How many columns of the export this lane lands. | [optional] 
**ColumnsWithheld** | Pointer to **NullableInt64** | How many columns of the export the lane withholds (contact details of named individuals). | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**SnapshotDate** | Pointer to **NullableString** | The date of the catalogue snapshot. | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdRegisterProvenanceRow

`func NewRwdRegisterProvenanceRow() *RwdRegisterProvenanceRow`

NewRwdRegisterProvenanceRow instantiates a new RwdRegisterProvenanceRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdRegisterProvenanceRowWithDefaults

`func NewRwdRegisterProvenanceRowWithDefaults() *RwdRegisterProvenanceRow`

NewRwdRegisterProvenanceRowWithDefaults instantiates a new RwdRegisterProvenanceRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegisterKey

`func (o *RwdRegisterProvenanceRow) GetRegisterKey() string`

GetRegisterKey returns the RegisterKey field if non-nil, zero value otherwise.

### GetRegisterKeyOk

`func (o *RwdRegisterProvenanceRow) GetRegisterKeyOk() (*string, bool)`

GetRegisterKeyOk returns a tuple with the RegisterKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterKey

`func (o *RwdRegisterProvenanceRow) SetRegisterKey(v string)`

SetRegisterKey sets RegisterKey field to given value.

### HasRegisterKey

`func (o *RwdRegisterProvenanceRow) HasRegisterKey() bool`

HasRegisterKey returns a boolean if a field has been set.

### SetRegisterKeyNil

`func (o *RwdRegisterProvenanceRow) SetRegisterKeyNil(b bool)`

 SetRegisterKeyNil sets the value for RegisterKey to be an explicit nil

### UnsetRegisterKey
`func (o *RwdRegisterProvenanceRow) UnsetRegisterKey()`

UnsetRegisterKey ensures that no value is present for RegisterKey, not even an explicit nil
### GetRegisterLabel

`func (o *RwdRegisterProvenanceRow) GetRegisterLabel() string`

GetRegisterLabel returns the RegisterLabel field if non-nil, zero value otherwise.

### GetRegisterLabelOk

`func (o *RwdRegisterProvenanceRow) GetRegisterLabelOk() (*string, bool)`

GetRegisterLabelOk returns a tuple with the RegisterLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterLabel

`func (o *RwdRegisterProvenanceRow) SetRegisterLabel(v string)`

SetRegisterLabel sets RegisterLabel field to given value.

### HasRegisterLabel

`func (o *RwdRegisterProvenanceRow) HasRegisterLabel() bool`

HasRegisterLabel returns a boolean if a field has been set.

### SetRegisterLabelNil

`func (o *RwdRegisterProvenanceRow) SetRegisterLabelNil(b bool)`

 SetRegisterLabelNil sets the value for RegisterLabel to be an explicit nil

### UnsetRegisterLabel
`func (o *RwdRegisterProvenanceRow) UnsetRegisterLabel()`

UnsetRegisterLabel ensures that no value is present for RegisterLabel, not even an explicit nil
### GetRegisterOrdinal

`func (o *RwdRegisterProvenanceRow) GetRegisterOrdinal() int32`

GetRegisterOrdinal returns the RegisterOrdinal field if non-nil, zero value otherwise.

### GetRegisterOrdinalOk

`func (o *RwdRegisterProvenanceRow) GetRegisterOrdinalOk() (*int32, bool)`

GetRegisterOrdinalOk returns a tuple with the RegisterOrdinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterOrdinal

`func (o *RwdRegisterProvenanceRow) SetRegisterOrdinal(v int32)`

SetRegisterOrdinal sets RegisterOrdinal field to given value.

### HasRegisterOrdinal

`func (o *RwdRegisterProvenanceRow) HasRegisterOrdinal() bool`

HasRegisterOrdinal returns a boolean if a field has been set.

### SetRegisterOrdinalNil

`func (o *RwdRegisterProvenanceRow) SetRegisterOrdinalNil(b bool)`

 SetRegisterOrdinalNil sets the value for RegisterOrdinal to be an explicit nil

### UnsetRegisterOrdinal
`func (o *RwdRegisterProvenanceRow) UnsetRegisterOrdinal()`

UnsetRegisterOrdinal ensures that no value is present for RegisterOrdinal, not even an explicit nil
### GetProvenanceSource

`func (o *RwdRegisterProvenanceRow) GetProvenanceSource() string`

GetProvenanceSource returns the ProvenanceSource field if non-nil, zero value otherwise.

### GetProvenanceSourceOk

`func (o *RwdRegisterProvenanceRow) GetProvenanceSourceOk() (*string, bool)`

GetProvenanceSourceOk returns a tuple with the ProvenanceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenanceSource

`func (o *RwdRegisterProvenanceRow) SetProvenanceSource(v string)`

SetProvenanceSource sets ProvenanceSource field to given value.

### HasProvenanceSource

`func (o *RwdRegisterProvenanceRow) HasProvenanceSource() bool`

HasProvenanceSource returns a boolean if a field has been set.

### SetProvenanceSourceNil

`func (o *RwdRegisterProvenanceRow) SetProvenanceSourceNil(b bool)`

 SetProvenanceSourceNil sets the value for ProvenanceSource to be an explicit nil

### UnsetProvenanceSource
`func (o *RwdRegisterProvenanceRow) UnsetProvenanceSource()`

UnsetProvenanceSource ensures that no value is present for ProvenanceSource, not even an explicit nil
### GetExportIsEmptyAtSource

`func (o *RwdRegisterProvenanceRow) GetExportIsEmptyAtSource() bool`

GetExportIsEmptyAtSource returns the ExportIsEmptyAtSource field if non-nil, zero value otherwise.

### GetExportIsEmptyAtSourceOk

`func (o *RwdRegisterProvenanceRow) GetExportIsEmptyAtSourceOk() (*bool, bool)`

GetExportIsEmptyAtSourceOk returns a tuple with the ExportIsEmptyAtSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportIsEmptyAtSource

`func (o *RwdRegisterProvenanceRow) SetExportIsEmptyAtSource(v bool)`

SetExportIsEmptyAtSource sets ExportIsEmptyAtSource field to given value.

### HasExportIsEmptyAtSource

`func (o *RwdRegisterProvenanceRow) HasExportIsEmptyAtSource() bool`

HasExportIsEmptyAtSource returns a boolean if a field has been set.

### SetExportIsEmptyAtSourceNil

`func (o *RwdRegisterProvenanceRow) SetExportIsEmptyAtSourceNil(b bool)`

 SetExportIsEmptyAtSourceNil sets the value for ExportIsEmptyAtSource to be an explicit nil

### UnsetExportIsEmptyAtSource
`func (o *RwdRegisterProvenanceRow) UnsetExportIsEmptyAtSource()`

UnsetExportIsEmptyAtSource ensures that no value is present for ExportIsEmptyAtSource, not even an explicit nil
### GetLandedRows

`func (o *RwdRegisterProvenanceRow) GetLandedRows() int64`

GetLandedRows returns the LandedRows field if non-nil, zero value otherwise.

### GetLandedRowsOk

`func (o *RwdRegisterProvenanceRow) GetLandedRowsOk() (*int64, bool)`

GetLandedRowsOk returns a tuple with the LandedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandedRows

`func (o *RwdRegisterProvenanceRow) SetLandedRows(v int64)`

SetLandedRows sets LandedRows field to given value.

### HasLandedRows

`func (o *RwdRegisterProvenanceRow) HasLandedRows() bool`

HasLandedRows returns a boolean if a field has been set.

### SetLandedRowsNil

`func (o *RwdRegisterProvenanceRow) SetLandedRowsNil(b bool)`

 SetLandedRowsNil sets the value for LandedRows to be an explicit nil

### UnsetLandedRows
`func (o *RwdRegisterProvenanceRow) UnsetLandedRows()`

UnsetLandedRows ensures that no value is present for LandedRows, not even an explicit nil
### GetBuildStatus

`func (o *RwdRegisterProvenanceRow) GetBuildStatus() string`

GetBuildStatus returns the BuildStatus field if non-nil, zero value otherwise.

### GetBuildStatusOk

`func (o *RwdRegisterProvenanceRow) GetBuildStatusOk() (*string, bool)`

GetBuildStatusOk returns a tuple with the BuildStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStatus

`func (o *RwdRegisterProvenanceRow) SetBuildStatus(v string)`

SetBuildStatus sets BuildStatus field to given value.

### HasBuildStatus

`func (o *RwdRegisterProvenanceRow) HasBuildStatus() bool`

HasBuildStatus returns a boolean if a field has been set.

### SetBuildStatusNil

`func (o *RwdRegisterProvenanceRow) SetBuildStatusNil(b bool)`

 SetBuildStatusNil sets the value for BuildStatus to be an explicit nil

### UnsetBuildStatus
`func (o *RwdRegisterProvenanceRow) UnsetBuildStatus()`

UnsetBuildStatus ensures that no value is present for BuildStatus, not even an explicit nil
### GetBuildSnapshotId

`func (o *RwdRegisterProvenanceRow) GetBuildSnapshotId() string`

GetBuildSnapshotId returns the BuildSnapshotId field if non-nil, zero value otherwise.

### GetBuildSnapshotIdOk

`func (o *RwdRegisterProvenanceRow) GetBuildSnapshotIdOk() (*string, bool)`

GetBuildSnapshotIdOk returns a tuple with the BuildSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSnapshotId

`func (o *RwdRegisterProvenanceRow) SetBuildSnapshotId(v string)`

SetBuildSnapshotId sets BuildSnapshotId field to given value.

### HasBuildSnapshotId

`func (o *RwdRegisterProvenanceRow) HasBuildSnapshotId() bool`

HasBuildSnapshotId returns a boolean if a field has been set.

### SetBuildSnapshotIdNil

`func (o *RwdRegisterProvenanceRow) SetBuildSnapshotIdNil(b bool)`

 SetBuildSnapshotIdNil sets the value for BuildSnapshotId to be an explicit nil

### UnsetBuildSnapshotId
`func (o *RwdRegisterProvenanceRow) UnsetBuildSnapshotId()`

UnsetBuildSnapshotId ensures that no value is present for BuildSnapshotId, not even an explicit nil
### GetBuildRowCount

`func (o *RwdRegisterProvenanceRow) GetBuildRowCount() int64`

GetBuildRowCount returns the BuildRowCount field if non-nil, zero value otherwise.

### GetBuildRowCountOk

`func (o *RwdRegisterProvenanceRow) GetBuildRowCountOk() (*int64, bool)`

GetBuildRowCountOk returns a tuple with the BuildRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildRowCount

`func (o *RwdRegisterProvenanceRow) SetBuildRowCount(v int64)`

SetBuildRowCount sets BuildRowCount field to given value.

### HasBuildRowCount

`func (o *RwdRegisterProvenanceRow) HasBuildRowCount() bool`

HasBuildRowCount returns a boolean if a field has been set.

### SetBuildRowCountNil

`func (o *RwdRegisterProvenanceRow) SetBuildRowCountNil(b bool)`

 SetBuildRowCountNil sets the value for BuildRowCount to be an explicit nil

### UnsetBuildRowCount
`func (o *RwdRegisterProvenanceRow) UnsetBuildRowCount()`

UnsetBuildRowCount ensures that no value is present for BuildRowCount, not even an explicit nil
### GetBuildStartedAt

`func (o *RwdRegisterProvenanceRow) GetBuildStartedAt() time.Time`

GetBuildStartedAt returns the BuildStartedAt field if non-nil, zero value otherwise.

### GetBuildStartedAtOk

`func (o *RwdRegisterProvenanceRow) GetBuildStartedAtOk() (*time.Time, bool)`

GetBuildStartedAtOk returns a tuple with the BuildStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStartedAt

`func (o *RwdRegisterProvenanceRow) SetBuildStartedAt(v time.Time)`

SetBuildStartedAt sets BuildStartedAt field to given value.

### HasBuildStartedAt

`func (o *RwdRegisterProvenanceRow) HasBuildStartedAt() bool`

HasBuildStartedAt returns a boolean if a field has been set.

### SetBuildStartedAtNil

`func (o *RwdRegisterProvenanceRow) SetBuildStartedAtNil(b bool)`

 SetBuildStartedAtNil sets the value for BuildStartedAt to be an explicit nil

### UnsetBuildStartedAt
`func (o *RwdRegisterProvenanceRow) UnsetBuildStartedAt()`

UnsetBuildStartedAt ensures that no value is present for BuildStartedAt, not even an explicit nil
### GetBuildFinishedAt

`func (o *RwdRegisterProvenanceRow) GetBuildFinishedAt() time.Time`

GetBuildFinishedAt returns the BuildFinishedAt field if non-nil, zero value otherwise.

### GetBuildFinishedAtOk

`func (o *RwdRegisterProvenanceRow) GetBuildFinishedAtOk() (*time.Time, bool)`

GetBuildFinishedAtOk returns a tuple with the BuildFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildFinishedAt

`func (o *RwdRegisterProvenanceRow) SetBuildFinishedAt(v time.Time)`

SetBuildFinishedAt sets BuildFinishedAt field to given value.

### HasBuildFinishedAt

`func (o *RwdRegisterProvenanceRow) HasBuildFinishedAt() bool`

HasBuildFinishedAt returns a boolean if a field has been set.

### SetBuildFinishedAtNil

`func (o *RwdRegisterProvenanceRow) SetBuildFinishedAtNil(b bool)`

 SetBuildFinishedAtNil sets the value for BuildFinishedAt to be an explicit nil

### UnsetBuildFinishedAt
`func (o *RwdRegisterProvenanceRow) UnsetBuildFinishedAt()`

UnsetBuildFinishedAt ensures that no value is present for BuildFinishedAt, not even an explicit nil
### GetLastCompletedSnapshotId

`func (o *RwdRegisterProvenanceRow) GetLastCompletedSnapshotId() string`

GetLastCompletedSnapshotId returns the LastCompletedSnapshotId field if non-nil, zero value otherwise.

### GetLastCompletedSnapshotIdOk

`func (o *RwdRegisterProvenanceRow) GetLastCompletedSnapshotIdOk() (*string, bool)`

GetLastCompletedSnapshotIdOk returns a tuple with the LastCompletedSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletedSnapshotId

`func (o *RwdRegisterProvenanceRow) SetLastCompletedSnapshotId(v string)`

SetLastCompletedSnapshotId sets LastCompletedSnapshotId field to given value.

### HasLastCompletedSnapshotId

`func (o *RwdRegisterProvenanceRow) HasLastCompletedSnapshotId() bool`

HasLastCompletedSnapshotId returns a boolean if a field has been set.

### SetLastCompletedSnapshotIdNil

`func (o *RwdRegisterProvenanceRow) SetLastCompletedSnapshotIdNil(b bool)`

 SetLastCompletedSnapshotIdNil sets the value for LastCompletedSnapshotId to be an explicit nil

### UnsetLastCompletedSnapshotId
`func (o *RwdRegisterProvenanceRow) UnsetLastCompletedSnapshotId()`

UnsetLastCompletedSnapshotId ensures that no value is present for LastCompletedSnapshotId, not even an explicit nil
### GetLastCompletedRowCount

`func (o *RwdRegisterProvenanceRow) GetLastCompletedRowCount() int64`

GetLastCompletedRowCount returns the LastCompletedRowCount field if non-nil, zero value otherwise.

### GetLastCompletedRowCountOk

`func (o *RwdRegisterProvenanceRow) GetLastCompletedRowCountOk() (*int64, bool)`

GetLastCompletedRowCountOk returns a tuple with the LastCompletedRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletedRowCount

`func (o *RwdRegisterProvenanceRow) SetLastCompletedRowCount(v int64)`

SetLastCompletedRowCount sets LastCompletedRowCount field to given value.

### HasLastCompletedRowCount

`func (o *RwdRegisterProvenanceRow) HasLastCompletedRowCount() bool`

HasLastCompletedRowCount returns a boolean if a field has been set.

### SetLastCompletedRowCountNil

`func (o *RwdRegisterProvenanceRow) SetLastCompletedRowCountNil(b bool)`

 SetLastCompletedRowCountNil sets the value for LastCompletedRowCount to be an explicit nil

### UnsetLastCompletedRowCount
`func (o *RwdRegisterProvenanceRow) UnsetLastCompletedRowCount()`

UnsetLastCompletedRowCount ensures that no value is present for LastCompletedRowCount, not even an explicit nil
### GetLastCompletedAt

`func (o *RwdRegisterProvenanceRow) GetLastCompletedAt() time.Time`

GetLastCompletedAt returns the LastCompletedAt field if non-nil, zero value otherwise.

### GetLastCompletedAtOk

`func (o *RwdRegisterProvenanceRow) GetLastCompletedAtOk() (*time.Time, bool)`

GetLastCompletedAtOk returns a tuple with the LastCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletedAt

`func (o *RwdRegisterProvenanceRow) SetLastCompletedAt(v time.Time)`

SetLastCompletedAt sets LastCompletedAt field to given value.

### HasLastCompletedAt

`func (o *RwdRegisterProvenanceRow) HasLastCompletedAt() bool`

HasLastCompletedAt returns a boolean if a field has been set.

### SetLastCompletedAtNil

`func (o *RwdRegisterProvenanceRow) SetLastCompletedAtNil(b bool)`

 SetLastCompletedAtNil sets the value for LastCompletedAt to be an explicit nil

### UnsetLastCompletedAt
`func (o *RwdRegisterProvenanceRow) UnsetLastCompletedAt()`

UnsetLastCompletedAt ensures that no value is present for LastCompletedAt, not even an explicit nil
### GetRegistrantClockAsof

`func (o *RwdRegisterProvenanceRow) GetRegistrantClockAsof() string`

GetRegistrantClockAsof returns the RegistrantClockAsof field if non-nil, zero value otherwise.

### GetRegistrantClockAsofOk

`func (o *RwdRegisterProvenanceRow) GetRegistrantClockAsofOk() (*string, bool)`

GetRegistrantClockAsofOk returns a tuple with the RegistrantClockAsof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrantClockAsof

`func (o *RwdRegisterProvenanceRow) SetRegistrantClockAsof(v string)`

SetRegistrantClockAsof sets RegistrantClockAsof field to given value.

### HasRegistrantClockAsof

`func (o *RwdRegisterProvenanceRow) HasRegistrantClockAsof() bool`

HasRegistrantClockAsof returns a boolean if a field has been set.

### SetRegistrantClockAsofNil

`func (o *RwdRegisterProvenanceRow) SetRegistrantClockAsofNil(b bool)`

 SetRegistrantClockAsofNil sets the value for RegistrantClockAsof to be an explicit nil

### UnsetRegistrantClockAsof
`func (o *RwdRegisterProvenanceRow) UnsetRegistrantClockAsof()`

UnsetRegistrantClockAsof ensures that no value is present for RegistrantClockAsof, not even an explicit nil
### GetRegisterPublishesUpdateDate

`func (o *RwdRegisterProvenanceRow) GetRegisterPublishesUpdateDate() bool`

GetRegisterPublishesUpdateDate returns the RegisterPublishesUpdateDate field if non-nil, zero value otherwise.

### GetRegisterPublishesUpdateDateOk

`func (o *RwdRegisterProvenanceRow) GetRegisterPublishesUpdateDateOk() (*bool, bool)`

GetRegisterPublishesUpdateDateOk returns a tuple with the RegisterPublishesUpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterPublishesUpdateDate

`func (o *RwdRegisterProvenanceRow) SetRegisterPublishesUpdateDate(v bool)`

SetRegisterPublishesUpdateDate sets RegisterPublishesUpdateDate field to given value.

### HasRegisterPublishesUpdateDate

`func (o *RwdRegisterProvenanceRow) HasRegisterPublishesUpdateDate() bool`

HasRegisterPublishesUpdateDate returns a boolean if a field has been set.

### SetRegisterPublishesUpdateDateNil

`func (o *RwdRegisterProvenanceRow) SetRegisterPublishesUpdateDateNil(b bool)`

 SetRegisterPublishesUpdateDateNil sets the value for RegisterPublishesUpdateDate to be an explicit nil

### UnsetRegisterPublishesUpdateDate
`func (o *RwdRegisterProvenanceRow) UnsetRegisterPublishesUpdateDate()`

UnsetRegisterPublishesUpdateDate ensures that no value is present for RegisterPublishesUpdateDate, not even an explicit nil
### GetRecordsWithAnUpdateDate

`func (o *RwdRegisterProvenanceRow) GetRecordsWithAnUpdateDate() int64`

GetRecordsWithAnUpdateDate returns the RecordsWithAnUpdateDate field if non-nil, zero value otherwise.

### GetRecordsWithAnUpdateDateOk

`func (o *RwdRegisterProvenanceRow) GetRecordsWithAnUpdateDateOk() (*int64, bool)`

GetRecordsWithAnUpdateDateOk returns a tuple with the RecordsWithAnUpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsWithAnUpdateDate

`func (o *RwdRegisterProvenanceRow) SetRecordsWithAnUpdateDate(v int64)`

SetRecordsWithAnUpdateDate sets RecordsWithAnUpdateDate field to given value.

### HasRecordsWithAnUpdateDate

`func (o *RwdRegisterProvenanceRow) HasRecordsWithAnUpdateDate() bool`

HasRecordsWithAnUpdateDate returns a boolean if a field has been set.

### SetRecordsWithAnUpdateDateNil

`func (o *RwdRegisterProvenanceRow) SetRecordsWithAnUpdateDateNil(b bool)`

 SetRecordsWithAnUpdateDateNil sets the value for RecordsWithAnUpdateDate to be an explicit nil

### UnsetRecordsWithAnUpdateDate
`func (o *RwdRegisterProvenanceRow) UnsetRecordsWithAnUpdateDate()`

UnsetRecordsWithAnUpdateDate ensures that no value is present for RecordsWithAnUpdateDate, not even an explicit nil
### GetRecordsUpdatedOverTwelveMonths

`func (o *RwdRegisterProvenanceRow) GetRecordsUpdatedOverTwelveMonths() int64`

GetRecordsUpdatedOverTwelveMonths returns the RecordsUpdatedOverTwelveMonths field if non-nil, zero value otherwise.

### GetRecordsUpdatedOverTwelveMonthsOk

`func (o *RwdRegisterProvenanceRow) GetRecordsUpdatedOverTwelveMonthsOk() (*int64, bool)`

GetRecordsUpdatedOverTwelveMonthsOk returns a tuple with the RecordsUpdatedOverTwelveMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsUpdatedOverTwelveMonths

`func (o *RwdRegisterProvenanceRow) SetRecordsUpdatedOverTwelveMonths(v int64)`

SetRecordsUpdatedOverTwelveMonths sets RecordsUpdatedOverTwelveMonths field to given value.

### HasRecordsUpdatedOverTwelveMonths

`func (o *RwdRegisterProvenanceRow) HasRecordsUpdatedOverTwelveMonths() bool`

HasRecordsUpdatedOverTwelveMonths returns a boolean if a field has been set.

### SetRecordsUpdatedOverTwelveMonthsNil

`func (o *RwdRegisterProvenanceRow) SetRecordsUpdatedOverTwelveMonthsNil(b bool)`

 SetRecordsUpdatedOverTwelveMonthsNil sets the value for RecordsUpdatedOverTwelveMonths to be an explicit nil

### UnsetRecordsUpdatedOverTwelveMonths
`func (o *RwdRegisterProvenanceRow) UnsetRecordsUpdatedOverTwelveMonths()`

UnsetRecordsUpdatedOverTwelveMonths ensures that no value is present for RecordsUpdatedOverTwelveMonths, not even an explicit nil
### GetShareUpdatedOverTwelveMonths

`func (o *RwdRegisterProvenanceRow) GetShareUpdatedOverTwelveMonths() float32`

GetShareUpdatedOverTwelveMonths returns the ShareUpdatedOverTwelveMonths field if non-nil, zero value otherwise.

### GetShareUpdatedOverTwelveMonthsOk

`func (o *RwdRegisterProvenanceRow) GetShareUpdatedOverTwelveMonthsOk() (*float32, bool)`

GetShareUpdatedOverTwelveMonthsOk returns a tuple with the ShareUpdatedOverTwelveMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareUpdatedOverTwelveMonths

`func (o *RwdRegisterProvenanceRow) SetShareUpdatedOverTwelveMonths(v float32)`

SetShareUpdatedOverTwelveMonths sets ShareUpdatedOverTwelveMonths field to given value.

### HasShareUpdatedOverTwelveMonths

`func (o *RwdRegisterProvenanceRow) HasShareUpdatedOverTwelveMonths() bool`

HasShareUpdatedOverTwelveMonths returns a boolean if a field has been set.

### SetShareUpdatedOverTwelveMonthsNil

`func (o *RwdRegisterProvenanceRow) SetShareUpdatedOverTwelveMonthsNil(b bool)`

 SetShareUpdatedOverTwelveMonthsNil sets the value for ShareUpdatedOverTwelveMonths to be an explicit nil

### UnsetShareUpdatedOverTwelveMonths
`func (o *RwdRegisterProvenanceRow) UnsetShareUpdatedOverTwelveMonths()`

UnsetShareUpdatedOverTwelveMonths ensures that no value is present for ShareUpdatedOverTwelveMonths, not even an explicit nil
### GetRegisterPublishesOutdatedFlag

`func (o *RwdRegisterProvenanceRow) GetRegisterPublishesOutdatedFlag() bool`

GetRegisterPublishesOutdatedFlag returns the RegisterPublishesOutdatedFlag field if non-nil, zero value otherwise.

### GetRegisterPublishesOutdatedFlagOk

`func (o *RwdRegisterProvenanceRow) GetRegisterPublishesOutdatedFlagOk() (*bool, bool)`

GetRegisterPublishesOutdatedFlagOk returns a tuple with the RegisterPublishesOutdatedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterPublishesOutdatedFlag

`func (o *RwdRegisterProvenanceRow) SetRegisterPublishesOutdatedFlag(v bool)`

SetRegisterPublishesOutdatedFlag sets RegisterPublishesOutdatedFlag field to given value.

### HasRegisterPublishesOutdatedFlag

`func (o *RwdRegisterProvenanceRow) HasRegisterPublishesOutdatedFlag() bool`

HasRegisterPublishesOutdatedFlag returns a boolean if a field has been set.

### SetRegisterPublishesOutdatedFlagNil

`func (o *RwdRegisterProvenanceRow) SetRegisterPublishesOutdatedFlagNil(b bool)`

 SetRegisterPublishesOutdatedFlagNil sets the value for RegisterPublishesOutdatedFlag to be an explicit nil

### UnsetRegisterPublishesOutdatedFlag
`func (o *RwdRegisterProvenanceRow) UnsetRegisterPublishesOutdatedFlag()`

UnsetRegisterPublishesOutdatedFlag ensures that no value is present for RegisterPublishesOutdatedFlag, not even an explicit nil
### GetRecordsEmaFlagsOutdated

`func (o *RwdRegisterProvenanceRow) GetRecordsEmaFlagsOutdated() int64`

GetRecordsEmaFlagsOutdated returns the RecordsEmaFlagsOutdated field if non-nil, zero value otherwise.

### GetRecordsEmaFlagsOutdatedOk

`func (o *RwdRegisterProvenanceRow) GetRecordsEmaFlagsOutdatedOk() (*int64, bool)`

GetRecordsEmaFlagsOutdatedOk returns a tuple with the RecordsEmaFlagsOutdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsEmaFlagsOutdated

`func (o *RwdRegisterProvenanceRow) SetRecordsEmaFlagsOutdated(v int64)`

SetRecordsEmaFlagsOutdated sets RecordsEmaFlagsOutdated field to given value.

### HasRecordsEmaFlagsOutdated

`func (o *RwdRegisterProvenanceRow) HasRecordsEmaFlagsOutdated() bool`

HasRecordsEmaFlagsOutdated returns a boolean if a field has been set.

### SetRecordsEmaFlagsOutdatedNil

`func (o *RwdRegisterProvenanceRow) SetRecordsEmaFlagsOutdatedNil(b bool)`

 SetRecordsEmaFlagsOutdatedNil sets the value for RecordsEmaFlagsOutdated to be an explicit nil

### UnsetRecordsEmaFlagsOutdated
`func (o *RwdRegisterProvenanceRow) UnsetRecordsEmaFlagsOutdated()`

UnsetRecordsEmaFlagsOutdated ensures that no value is present for RecordsEmaFlagsOutdated, not even an explicit nil
### GetRecordsOutdatedFlagBlank

`func (o *RwdRegisterProvenanceRow) GetRecordsOutdatedFlagBlank() int64`

GetRecordsOutdatedFlagBlank returns the RecordsOutdatedFlagBlank field if non-nil, zero value otherwise.

### GetRecordsOutdatedFlagBlankOk

`func (o *RwdRegisterProvenanceRow) GetRecordsOutdatedFlagBlankOk() (*int64, bool)`

GetRecordsOutdatedFlagBlankOk returns a tuple with the RecordsOutdatedFlagBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsOutdatedFlagBlank

`func (o *RwdRegisterProvenanceRow) SetRecordsOutdatedFlagBlank(v int64)`

SetRecordsOutdatedFlagBlank sets RecordsOutdatedFlagBlank field to given value.

### HasRecordsOutdatedFlagBlank

`func (o *RwdRegisterProvenanceRow) HasRecordsOutdatedFlagBlank() bool`

HasRecordsOutdatedFlagBlank returns a boolean if a field has been set.

### SetRecordsOutdatedFlagBlankNil

`func (o *RwdRegisterProvenanceRow) SetRecordsOutdatedFlagBlankNil(b bool)`

 SetRecordsOutdatedFlagBlankNil sets the value for RecordsOutdatedFlagBlank to be an explicit nil

### UnsetRecordsOutdatedFlagBlank
`func (o *RwdRegisterProvenanceRow) UnsetRecordsOutdatedFlagBlank()`

UnsetRecordsOutdatedFlagBlank ensures that no value is present for RecordsOutdatedFlagBlank, not even an explicit nil
### GetColumnsLanded

`func (o *RwdRegisterProvenanceRow) GetColumnsLanded() int64`

GetColumnsLanded returns the ColumnsLanded field if non-nil, zero value otherwise.

### GetColumnsLandedOk

`func (o *RwdRegisterProvenanceRow) GetColumnsLandedOk() (*int64, bool)`

GetColumnsLandedOk returns a tuple with the ColumnsLanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnsLanded

`func (o *RwdRegisterProvenanceRow) SetColumnsLanded(v int64)`

SetColumnsLanded sets ColumnsLanded field to given value.

### HasColumnsLanded

`func (o *RwdRegisterProvenanceRow) HasColumnsLanded() bool`

HasColumnsLanded returns a boolean if a field has been set.

### SetColumnsLandedNil

`func (o *RwdRegisterProvenanceRow) SetColumnsLandedNil(b bool)`

 SetColumnsLandedNil sets the value for ColumnsLanded to be an explicit nil

### UnsetColumnsLanded
`func (o *RwdRegisterProvenanceRow) UnsetColumnsLanded()`

UnsetColumnsLanded ensures that no value is present for ColumnsLanded, not even an explicit nil
### GetColumnsWithheld

`func (o *RwdRegisterProvenanceRow) GetColumnsWithheld() int64`

GetColumnsWithheld returns the ColumnsWithheld field if non-nil, zero value otherwise.

### GetColumnsWithheldOk

`func (o *RwdRegisterProvenanceRow) GetColumnsWithheldOk() (*int64, bool)`

GetColumnsWithheldOk returns a tuple with the ColumnsWithheld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnsWithheld

`func (o *RwdRegisterProvenanceRow) SetColumnsWithheld(v int64)`

SetColumnsWithheld sets ColumnsWithheld field to given value.

### HasColumnsWithheld

`func (o *RwdRegisterProvenanceRow) HasColumnsWithheld() bool`

HasColumnsWithheld returns a boolean if a field has been set.

### SetColumnsWithheldNil

`func (o *RwdRegisterProvenanceRow) SetColumnsWithheldNil(b bool)`

 SetColumnsWithheldNil sets the value for ColumnsWithheld to be an explicit nil

### UnsetColumnsWithheld
`func (o *RwdRegisterProvenanceRow) UnsetColumnsWithheld()`

UnsetColumnsWithheld ensures that no value is present for ColumnsWithheld, not even an explicit nil
### GetSnapshotLabel

`func (o *RwdRegisterProvenanceRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *RwdRegisterProvenanceRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *RwdRegisterProvenanceRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *RwdRegisterProvenanceRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *RwdRegisterProvenanceRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *RwdRegisterProvenanceRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetSnapshotDate

`func (o *RwdRegisterProvenanceRow) GetSnapshotDate() string`

GetSnapshotDate returns the SnapshotDate field if non-nil, zero value otherwise.

### GetSnapshotDateOk

`func (o *RwdRegisterProvenanceRow) GetSnapshotDateOk() (*string, bool)`

GetSnapshotDateOk returns a tuple with the SnapshotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotDate

`func (o *RwdRegisterProvenanceRow) SetSnapshotDate(v string)`

SetSnapshotDate sets SnapshotDate field to given value.

### HasSnapshotDate

`func (o *RwdRegisterProvenanceRow) HasSnapshotDate() bool`

HasSnapshotDate returns a boolean if a field has been set.

### SetSnapshotDateNil

`func (o *RwdRegisterProvenanceRow) SetSnapshotDateNil(b bool)`

 SetSnapshotDateNil sets the value for SnapshotDate to be an explicit nil

### UnsetSnapshotDate
`func (o *RwdRegisterProvenanceRow) UnsetSnapshotDate()`

UnsetSnapshotDate ensures that no value is present for SnapshotDate, not even an explicit nil
### GetAccessedMonth

`func (o *RwdRegisterProvenanceRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *RwdRegisterProvenanceRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *RwdRegisterProvenanceRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *RwdRegisterProvenanceRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *RwdRegisterProvenanceRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *RwdRegisterProvenanceRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *RwdRegisterProvenanceRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdRegisterProvenanceRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdRegisterProvenanceRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdRegisterProvenanceRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdRegisterProvenanceRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdRegisterProvenanceRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdRegisterProvenanceRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdRegisterProvenanceRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdRegisterProvenanceRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdRegisterProvenanceRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdRegisterProvenanceRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdRegisterProvenanceRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


