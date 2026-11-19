# InvestmentGraphEdgesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeKey** | Pointer to **NullableString** | The stable handle of one edge row: md5 of from_node_id, to_node_id, edge_type, source and source_ref. | [optional] 
**FromNodeId** | Pointer to **NullableString** | The node the edge starts at: an investor id or a company id (the resolved parent organization, or investor_site: and the normalized name for a website-only company the organization spine has not admitted). | [optional] 
**ToNodeId** | Pointer to **NullableString** | The node the edge ends at. Same id forms as from_node_id. invested_in goes from an investor to a company; co_invested is stored in both directions between investors. | [optional] 
**FromInvestorId** | Pointer to **NullableString** | The investor key when the from node is an investor (every edge today), null otherwise. | [optional] 
**ToInvestorId** | Pointer to **NullableString** | The investor key when the to node is an investor (a co_invested edge), null otherwise. | [optional] 
**EdgeType** | Pointer to **NullableString** | invested_in, co_invested, acquired, merged_with or licensed. acquired, merged_with and licensed have no rows until the M&amp;A events dataset lands. There is no exited type: a listing that disappears from a website is never an end. | [optional] 
**RelationshipType** | Pointer to **NullableString** | The lane&#39;s own word for the relationship, so a public-market holding is never read as a venture round: portfolio_listed (investor website), holder_5pct and selling_stockholder (offering holder tables), institutional_position (13F), beneficial_owner (13D/G), co_appearance (co_invested). | [optional] 
**Source** | Pointer to **NullableString** | The lane that states the edge: investor_site, sec_edgar_holder_table (also the lane of the derived co_invested edges), sec_13f, sec_13dg, sec_form_d, sec_8k, company_press_release or reconciled. sec_form_d, sec_8k and company_press_release have no rows yet; reconciled carries the press-release and SEC deal metadata for the pairs it enriches and is the preferred lane of such a pair. | [optional] 
**SourceRef** | Pointer to **NullableString** | The evidence key within the lane: the portfolio company key for the website and holder-table lanes, the ownership observation id for 13F and 13D/G, and the literal co_invested_pairs for a derived co-investment. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The page the edge was read from. Null for a derived co_invested edge, which has no page of its own. | [optional] 
**EventDate** | Pointer to **NullableString** | The date the lane states, of the kind in event_date_kind. Null exactly when event_date_kind is unknown. | [optional] 
**EventDateKind** | Pointer to **NullableString** | What event_date is: transaction (a date the lane states for the deal itself), announcement, filing (the filing date), as_of (a report period end, as 13F states), form_d_filing (the filing date of a Form D offering by the company, never an investment date), first_observed (the first crawl that saw a website listing, never an investment date) or unknown (no date, event_date null). | [optional] 
**ObservedThrough** | Pointer to **NullableString** | The latest date the relationship was still observed in the lane (the last crawl, the last filing). Null when the lane gives none. Not an end date. | [optional] 
**EventStage** | Pointer to **NullableString** | Where a deal stands (announced, signed, closed, completed or terminated), so an announced acquisition is never drawn as a completed one. Null for every lane today. | [optional] 
**RoundLabel** | Pointer to **NullableString** | The funding round named by the company&#39;s own press release (reconciled lane); null for every lane that does not state one. | [optional] 
**AmountValue** | Pointer to **NullableFloat32** | The ROUND amount in amount_currency, never the investor&#39;s own cheque (reconciled lane): the press-release round total, else the amount sold in a Form D offering. Null for every lane that does not state one. | [optional] 
**AmountCurrency** | Pointer to **NullableString** | The currency of amount_value as the source states it, never defaulted to USD. Null while amount_value is null. | [optional] 
**IsLead** | Pointer to **NullableBool** | Whether the investor led or co-led the round per the company&#39;s press release (reconciled lane): true for lead or co-lead, false for a named participant, null where the role is unstated or the lane does not state it. | [optional] 
**SharedCompanyCount** | Pointer to **NullableInt32** | For a derived co_invested edge, the companies the two investors are both named against in offering filings. Null for every other edge. | [optional] 
**IsDerived** | Pointer to **NullableBool** | True for an edge computed from other rows (co_invested), false for an edge a source states. | [optional] 
**PairKey** | Pointer to **NullableString** | from_node_id, to_node_id and edge_type joined with a pipe: every lane that states the same relationship shares one pair_key. | [optional] 
**LaneRank** | Pointer to **NullableInt32** | The lane&#39;s precedence when several lanes state one pair (lower wins): 13D/G 1, 13F 2, offering filings and derived co-investment 3, investor website 7. | [optional] 
**IsPreferred** | Pointer to **NullableBool** | True on exactly one row per pair_key, the lowest lane_rank (then the latest event_date): the row a picture draws. The other rows stay as evidence. | [optional] 
**SourceLicense** | Pointer to **NullableString** | The license class of the lane: public_domain_us_government or first_party_website_facts. | [optional] 
**Attribution** | Pointer to **NullableString** | The plain credit line for the lane. Show it with the edge for every reader. | [optional] 
**FromName** | Pointer to **NullableString** | Display name of the from node. | [optional] 
**FromOrgId** | Pointer to **NullableString** | Legal-entity level: the from node&#39;s organization record (null for an investor with none and for an unadmitted website-only company). | [optional] 
**FromCompanyId** | Pointer to **NullableString** | Company level: the from node&#39;s company record (null where unresolved). | [optional] 
**FromCompanyName** | Pointer to **NullableString** | Display name of the from node&#39;s company. | [optional] 
**FromGroupId** | Pointer to **NullableString** | Group level: the from node&#39;s company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**FromGroupName** | Pointer to **NullableString** | Display name of the from node&#39;s group. | [optional] 
**ToName** | Pointer to **NullableString** | Display name of the to node. | [optional] 
**ToOrgId** | Pointer to **NullableString** | Legal-entity level: the to node&#39;s organization record (null for an investor with none and for an unadmitted website-only company). | [optional] 
**ToCompanyId** | Pointer to **NullableString** | Company level: the to node&#39;s company record (null where unresolved). | [optional] 
**ToCompanyName** | Pointer to **NullableString** | Display name of the to node&#39;s company. | [optional] 
**ToGroupId** | Pointer to **NullableString** | Group level: the to node&#39;s company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**ToGroupName** | Pointer to **NullableString** | Display name of the to node&#39;s group. | [optional] 

## Methods

### NewInvestmentGraphEdgesRow

`func NewInvestmentGraphEdgesRow() *InvestmentGraphEdgesRow`

NewInvestmentGraphEdgesRow instantiates a new InvestmentGraphEdgesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestmentGraphEdgesRowWithDefaults

`func NewInvestmentGraphEdgesRowWithDefaults() *InvestmentGraphEdgesRow`

NewInvestmentGraphEdgesRowWithDefaults instantiates a new InvestmentGraphEdgesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeKey

`func (o *InvestmentGraphEdgesRow) GetEdgeKey() string`

GetEdgeKey returns the EdgeKey field if non-nil, zero value otherwise.

### GetEdgeKeyOk

`func (o *InvestmentGraphEdgesRow) GetEdgeKeyOk() (*string, bool)`

GetEdgeKeyOk returns a tuple with the EdgeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeKey

`func (o *InvestmentGraphEdgesRow) SetEdgeKey(v string)`

SetEdgeKey sets EdgeKey field to given value.

### HasEdgeKey

`func (o *InvestmentGraphEdgesRow) HasEdgeKey() bool`

HasEdgeKey returns a boolean if a field has been set.

### SetEdgeKeyNil

`func (o *InvestmentGraphEdgesRow) SetEdgeKeyNil(b bool)`

 SetEdgeKeyNil sets the value for EdgeKey to be an explicit nil

### UnsetEdgeKey
`func (o *InvestmentGraphEdgesRow) UnsetEdgeKey()`

UnsetEdgeKey ensures that no value is present for EdgeKey, not even an explicit nil
### GetFromNodeId

`func (o *InvestmentGraphEdgesRow) GetFromNodeId() string`

GetFromNodeId returns the FromNodeId field if non-nil, zero value otherwise.

### GetFromNodeIdOk

`func (o *InvestmentGraphEdgesRow) GetFromNodeIdOk() (*string, bool)`

GetFromNodeIdOk returns a tuple with the FromNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromNodeId

`func (o *InvestmentGraphEdgesRow) SetFromNodeId(v string)`

SetFromNodeId sets FromNodeId field to given value.

### HasFromNodeId

`func (o *InvestmentGraphEdgesRow) HasFromNodeId() bool`

HasFromNodeId returns a boolean if a field has been set.

### SetFromNodeIdNil

`func (o *InvestmentGraphEdgesRow) SetFromNodeIdNil(b bool)`

 SetFromNodeIdNil sets the value for FromNodeId to be an explicit nil

### UnsetFromNodeId
`func (o *InvestmentGraphEdgesRow) UnsetFromNodeId()`

UnsetFromNodeId ensures that no value is present for FromNodeId, not even an explicit nil
### GetToNodeId

`func (o *InvestmentGraphEdgesRow) GetToNodeId() string`

GetToNodeId returns the ToNodeId field if non-nil, zero value otherwise.

### GetToNodeIdOk

`func (o *InvestmentGraphEdgesRow) GetToNodeIdOk() (*string, bool)`

GetToNodeIdOk returns a tuple with the ToNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToNodeId

`func (o *InvestmentGraphEdgesRow) SetToNodeId(v string)`

SetToNodeId sets ToNodeId field to given value.

### HasToNodeId

`func (o *InvestmentGraphEdgesRow) HasToNodeId() bool`

HasToNodeId returns a boolean if a field has been set.

### SetToNodeIdNil

`func (o *InvestmentGraphEdgesRow) SetToNodeIdNil(b bool)`

 SetToNodeIdNil sets the value for ToNodeId to be an explicit nil

### UnsetToNodeId
`func (o *InvestmentGraphEdgesRow) UnsetToNodeId()`

UnsetToNodeId ensures that no value is present for ToNodeId, not even an explicit nil
### GetFromInvestorId

`func (o *InvestmentGraphEdgesRow) GetFromInvestorId() string`

GetFromInvestorId returns the FromInvestorId field if non-nil, zero value otherwise.

### GetFromInvestorIdOk

`func (o *InvestmentGraphEdgesRow) GetFromInvestorIdOk() (*string, bool)`

GetFromInvestorIdOk returns a tuple with the FromInvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromInvestorId

`func (o *InvestmentGraphEdgesRow) SetFromInvestorId(v string)`

SetFromInvestorId sets FromInvestorId field to given value.

### HasFromInvestorId

`func (o *InvestmentGraphEdgesRow) HasFromInvestorId() bool`

HasFromInvestorId returns a boolean if a field has been set.

### SetFromInvestorIdNil

`func (o *InvestmentGraphEdgesRow) SetFromInvestorIdNil(b bool)`

 SetFromInvestorIdNil sets the value for FromInvestorId to be an explicit nil

### UnsetFromInvestorId
`func (o *InvestmentGraphEdgesRow) UnsetFromInvestorId()`

UnsetFromInvestorId ensures that no value is present for FromInvestorId, not even an explicit nil
### GetToInvestorId

`func (o *InvestmentGraphEdgesRow) GetToInvestorId() string`

GetToInvestorId returns the ToInvestorId field if non-nil, zero value otherwise.

### GetToInvestorIdOk

`func (o *InvestmentGraphEdgesRow) GetToInvestorIdOk() (*string, bool)`

GetToInvestorIdOk returns a tuple with the ToInvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToInvestorId

`func (o *InvestmentGraphEdgesRow) SetToInvestorId(v string)`

SetToInvestorId sets ToInvestorId field to given value.

### HasToInvestorId

`func (o *InvestmentGraphEdgesRow) HasToInvestorId() bool`

HasToInvestorId returns a boolean if a field has been set.

### SetToInvestorIdNil

`func (o *InvestmentGraphEdgesRow) SetToInvestorIdNil(b bool)`

 SetToInvestorIdNil sets the value for ToInvestorId to be an explicit nil

### UnsetToInvestorId
`func (o *InvestmentGraphEdgesRow) UnsetToInvestorId()`

UnsetToInvestorId ensures that no value is present for ToInvestorId, not even an explicit nil
### GetEdgeType

`func (o *InvestmentGraphEdgesRow) GetEdgeType() string`

GetEdgeType returns the EdgeType field if non-nil, zero value otherwise.

### GetEdgeTypeOk

`func (o *InvestmentGraphEdgesRow) GetEdgeTypeOk() (*string, bool)`

GetEdgeTypeOk returns a tuple with the EdgeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeType

`func (o *InvestmentGraphEdgesRow) SetEdgeType(v string)`

SetEdgeType sets EdgeType field to given value.

### HasEdgeType

`func (o *InvestmentGraphEdgesRow) HasEdgeType() bool`

HasEdgeType returns a boolean if a field has been set.

### SetEdgeTypeNil

`func (o *InvestmentGraphEdgesRow) SetEdgeTypeNil(b bool)`

 SetEdgeTypeNil sets the value for EdgeType to be an explicit nil

### UnsetEdgeType
`func (o *InvestmentGraphEdgesRow) UnsetEdgeType()`

UnsetEdgeType ensures that no value is present for EdgeType, not even an explicit nil
### GetRelationshipType

`func (o *InvestmentGraphEdgesRow) GetRelationshipType() string`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *InvestmentGraphEdgesRow) GetRelationshipTypeOk() (*string, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *InvestmentGraphEdgesRow) SetRelationshipType(v string)`

SetRelationshipType sets RelationshipType field to given value.

### HasRelationshipType

`func (o *InvestmentGraphEdgesRow) HasRelationshipType() bool`

HasRelationshipType returns a boolean if a field has been set.

### SetRelationshipTypeNil

`func (o *InvestmentGraphEdgesRow) SetRelationshipTypeNil(b bool)`

 SetRelationshipTypeNil sets the value for RelationshipType to be an explicit nil

### UnsetRelationshipType
`func (o *InvestmentGraphEdgesRow) UnsetRelationshipType()`

UnsetRelationshipType ensures that no value is present for RelationshipType, not even an explicit nil
### GetSource

`func (o *InvestmentGraphEdgesRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InvestmentGraphEdgesRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InvestmentGraphEdgesRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InvestmentGraphEdgesRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *InvestmentGraphEdgesRow) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *InvestmentGraphEdgesRow) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceRef

`func (o *InvestmentGraphEdgesRow) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *InvestmentGraphEdgesRow) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *InvestmentGraphEdgesRow) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *InvestmentGraphEdgesRow) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### SetSourceRefNil

`func (o *InvestmentGraphEdgesRow) SetSourceRefNil(b bool)`

 SetSourceRefNil sets the value for SourceRef to be an explicit nil

### UnsetSourceRef
`func (o *InvestmentGraphEdgesRow) UnsetSourceRef()`

UnsetSourceRef ensures that no value is present for SourceRef, not even an explicit nil
### GetSourceUrl

`func (o *InvestmentGraphEdgesRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *InvestmentGraphEdgesRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *InvestmentGraphEdgesRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *InvestmentGraphEdgesRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *InvestmentGraphEdgesRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *InvestmentGraphEdgesRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetEventDate

`func (o *InvestmentGraphEdgesRow) GetEventDate() string`

GetEventDate returns the EventDate field if non-nil, zero value otherwise.

### GetEventDateOk

`func (o *InvestmentGraphEdgesRow) GetEventDateOk() (*string, bool)`

GetEventDateOk returns a tuple with the EventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDate

`func (o *InvestmentGraphEdgesRow) SetEventDate(v string)`

SetEventDate sets EventDate field to given value.

### HasEventDate

`func (o *InvestmentGraphEdgesRow) HasEventDate() bool`

HasEventDate returns a boolean if a field has been set.

### SetEventDateNil

`func (o *InvestmentGraphEdgesRow) SetEventDateNil(b bool)`

 SetEventDateNil sets the value for EventDate to be an explicit nil

### UnsetEventDate
`func (o *InvestmentGraphEdgesRow) UnsetEventDate()`

UnsetEventDate ensures that no value is present for EventDate, not even an explicit nil
### GetEventDateKind

`func (o *InvestmentGraphEdgesRow) GetEventDateKind() string`

GetEventDateKind returns the EventDateKind field if non-nil, zero value otherwise.

### GetEventDateKindOk

`func (o *InvestmentGraphEdgesRow) GetEventDateKindOk() (*string, bool)`

GetEventDateKindOk returns a tuple with the EventDateKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateKind

`func (o *InvestmentGraphEdgesRow) SetEventDateKind(v string)`

SetEventDateKind sets EventDateKind field to given value.

### HasEventDateKind

`func (o *InvestmentGraphEdgesRow) HasEventDateKind() bool`

HasEventDateKind returns a boolean if a field has been set.

### SetEventDateKindNil

`func (o *InvestmentGraphEdgesRow) SetEventDateKindNil(b bool)`

 SetEventDateKindNil sets the value for EventDateKind to be an explicit nil

### UnsetEventDateKind
`func (o *InvestmentGraphEdgesRow) UnsetEventDateKind()`

UnsetEventDateKind ensures that no value is present for EventDateKind, not even an explicit nil
### GetObservedThrough

`func (o *InvestmentGraphEdgesRow) GetObservedThrough() string`

GetObservedThrough returns the ObservedThrough field if non-nil, zero value otherwise.

### GetObservedThroughOk

`func (o *InvestmentGraphEdgesRow) GetObservedThroughOk() (*string, bool)`

GetObservedThroughOk returns a tuple with the ObservedThrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedThrough

`func (o *InvestmentGraphEdgesRow) SetObservedThrough(v string)`

SetObservedThrough sets ObservedThrough field to given value.

### HasObservedThrough

`func (o *InvestmentGraphEdgesRow) HasObservedThrough() bool`

HasObservedThrough returns a boolean if a field has been set.

### SetObservedThroughNil

`func (o *InvestmentGraphEdgesRow) SetObservedThroughNil(b bool)`

 SetObservedThroughNil sets the value for ObservedThrough to be an explicit nil

### UnsetObservedThrough
`func (o *InvestmentGraphEdgesRow) UnsetObservedThrough()`

UnsetObservedThrough ensures that no value is present for ObservedThrough, not even an explicit nil
### GetEventStage

`func (o *InvestmentGraphEdgesRow) GetEventStage() string`

GetEventStage returns the EventStage field if non-nil, zero value otherwise.

### GetEventStageOk

`func (o *InvestmentGraphEdgesRow) GetEventStageOk() (*string, bool)`

GetEventStageOk returns a tuple with the EventStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStage

`func (o *InvestmentGraphEdgesRow) SetEventStage(v string)`

SetEventStage sets EventStage field to given value.

### HasEventStage

`func (o *InvestmentGraphEdgesRow) HasEventStage() bool`

HasEventStage returns a boolean if a field has been set.

### SetEventStageNil

`func (o *InvestmentGraphEdgesRow) SetEventStageNil(b bool)`

 SetEventStageNil sets the value for EventStage to be an explicit nil

### UnsetEventStage
`func (o *InvestmentGraphEdgesRow) UnsetEventStage()`

UnsetEventStage ensures that no value is present for EventStage, not even an explicit nil
### GetRoundLabel

`func (o *InvestmentGraphEdgesRow) GetRoundLabel() string`

GetRoundLabel returns the RoundLabel field if non-nil, zero value otherwise.

### GetRoundLabelOk

`func (o *InvestmentGraphEdgesRow) GetRoundLabelOk() (*string, bool)`

GetRoundLabelOk returns a tuple with the RoundLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundLabel

`func (o *InvestmentGraphEdgesRow) SetRoundLabel(v string)`

SetRoundLabel sets RoundLabel field to given value.

### HasRoundLabel

`func (o *InvestmentGraphEdgesRow) HasRoundLabel() bool`

HasRoundLabel returns a boolean if a field has been set.

### SetRoundLabelNil

`func (o *InvestmentGraphEdgesRow) SetRoundLabelNil(b bool)`

 SetRoundLabelNil sets the value for RoundLabel to be an explicit nil

### UnsetRoundLabel
`func (o *InvestmentGraphEdgesRow) UnsetRoundLabel()`

UnsetRoundLabel ensures that no value is present for RoundLabel, not even an explicit nil
### GetAmountValue

`func (o *InvestmentGraphEdgesRow) GetAmountValue() float32`

GetAmountValue returns the AmountValue field if non-nil, zero value otherwise.

### GetAmountValueOk

`func (o *InvestmentGraphEdgesRow) GetAmountValueOk() (*float32, bool)`

GetAmountValueOk returns a tuple with the AmountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountValue

`func (o *InvestmentGraphEdgesRow) SetAmountValue(v float32)`

SetAmountValue sets AmountValue field to given value.

### HasAmountValue

`func (o *InvestmentGraphEdgesRow) HasAmountValue() bool`

HasAmountValue returns a boolean if a field has been set.

### SetAmountValueNil

`func (o *InvestmentGraphEdgesRow) SetAmountValueNil(b bool)`

 SetAmountValueNil sets the value for AmountValue to be an explicit nil

### UnsetAmountValue
`func (o *InvestmentGraphEdgesRow) UnsetAmountValue()`

UnsetAmountValue ensures that no value is present for AmountValue, not even an explicit nil
### GetAmountCurrency

`func (o *InvestmentGraphEdgesRow) GetAmountCurrency() string`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *InvestmentGraphEdgesRow) GetAmountCurrencyOk() (*string, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *InvestmentGraphEdgesRow) SetAmountCurrency(v string)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *InvestmentGraphEdgesRow) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.

### SetAmountCurrencyNil

`func (o *InvestmentGraphEdgesRow) SetAmountCurrencyNil(b bool)`

 SetAmountCurrencyNil sets the value for AmountCurrency to be an explicit nil

### UnsetAmountCurrency
`func (o *InvestmentGraphEdgesRow) UnsetAmountCurrency()`

UnsetAmountCurrency ensures that no value is present for AmountCurrency, not even an explicit nil
### GetIsLead

`func (o *InvestmentGraphEdgesRow) GetIsLead() bool`

GetIsLead returns the IsLead field if non-nil, zero value otherwise.

### GetIsLeadOk

`func (o *InvestmentGraphEdgesRow) GetIsLeadOk() (*bool, bool)`

GetIsLeadOk returns a tuple with the IsLead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLead

`func (o *InvestmentGraphEdgesRow) SetIsLead(v bool)`

SetIsLead sets IsLead field to given value.

### HasIsLead

`func (o *InvestmentGraphEdgesRow) HasIsLead() bool`

HasIsLead returns a boolean if a field has been set.

### SetIsLeadNil

`func (o *InvestmentGraphEdgesRow) SetIsLeadNil(b bool)`

 SetIsLeadNil sets the value for IsLead to be an explicit nil

### UnsetIsLead
`func (o *InvestmentGraphEdgesRow) UnsetIsLead()`

UnsetIsLead ensures that no value is present for IsLead, not even an explicit nil
### GetSharedCompanyCount

`func (o *InvestmentGraphEdgesRow) GetSharedCompanyCount() int32`

GetSharedCompanyCount returns the SharedCompanyCount field if non-nil, zero value otherwise.

### GetSharedCompanyCountOk

`func (o *InvestmentGraphEdgesRow) GetSharedCompanyCountOk() (*int32, bool)`

GetSharedCompanyCountOk returns a tuple with the SharedCompanyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedCompanyCount

`func (o *InvestmentGraphEdgesRow) SetSharedCompanyCount(v int32)`

SetSharedCompanyCount sets SharedCompanyCount field to given value.

### HasSharedCompanyCount

`func (o *InvestmentGraphEdgesRow) HasSharedCompanyCount() bool`

HasSharedCompanyCount returns a boolean if a field has been set.

### SetSharedCompanyCountNil

`func (o *InvestmentGraphEdgesRow) SetSharedCompanyCountNil(b bool)`

 SetSharedCompanyCountNil sets the value for SharedCompanyCount to be an explicit nil

### UnsetSharedCompanyCount
`func (o *InvestmentGraphEdgesRow) UnsetSharedCompanyCount()`

UnsetSharedCompanyCount ensures that no value is present for SharedCompanyCount, not even an explicit nil
### GetIsDerived

`func (o *InvestmentGraphEdgesRow) GetIsDerived() bool`

GetIsDerived returns the IsDerived field if non-nil, zero value otherwise.

### GetIsDerivedOk

`func (o *InvestmentGraphEdgesRow) GetIsDerivedOk() (*bool, bool)`

GetIsDerivedOk returns a tuple with the IsDerived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDerived

`func (o *InvestmentGraphEdgesRow) SetIsDerived(v bool)`

SetIsDerived sets IsDerived field to given value.

### HasIsDerived

`func (o *InvestmentGraphEdgesRow) HasIsDerived() bool`

HasIsDerived returns a boolean if a field has been set.

### SetIsDerivedNil

`func (o *InvestmentGraphEdgesRow) SetIsDerivedNil(b bool)`

 SetIsDerivedNil sets the value for IsDerived to be an explicit nil

### UnsetIsDerived
`func (o *InvestmentGraphEdgesRow) UnsetIsDerived()`

UnsetIsDerived ensures that no value is present for IsDerived, not even an explicit nil
### GetPairKey

`func (o *InvestmentGraphEdgesRow) GetPairKey() string`

GetPairKey returns the PairKey field if non-nil, zero value otherwise.

### GetPairKeyOk

`func (o *InvestmentGraphEdgesRow) GetPairKeyOk() (*string, bool)`

GetPairKeyOk returns a tuple with the PairKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairKey

`func (o *InvestmentGraphEdgesRow) SetPairKey(v string)`

SetPairKey sets PairKey field to given value.

### HasPairKey

`func (o *InvestmentGraphEdgesRow) HasPairKey() bool`

HasPairKey returns a boolean if a field has been set.

### SetPairKeyNil

`func (o *InvestmentGraphEdgesRow) SetPairKeyNil(b bool)`

 SetPairKeyNil sets the value for PairKey to be an explicit nil

### UnsetPairKey
`func (o *InvestmentGraphEdgesRow) UnsetPairKey()`

UnsetPairKey ensures that no value is present for PairKey, not even an explicit nil
### GetLaneRank

`func (o *InvestmentGraphEdgesRow) GetLaneRank() int32`

GetLaneRank returns the LaneRank field if non-nil, zero value otherwise.

### GetLaneRankOk

`func (o *InvestmentGraphEdgesRow) GetLaneRankOk() (*int32, bool)`

GetLaneRankOk returns a tuple with the LaneRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaneRank

`func (o *InvestmentGraphEdgesRow) SetLaneRank(v int32)`

SetLaneRank sets LaneRank field to given value.

### HasLaneRank

`func (o *InvestmentGraphEdgesRow) HasLaneRank() bool`

HasLaneRank returns a boolean if a field has been set.

### SetLaneRankNil

`func (o *InvestmentGraphEdgesRow) SetLaneRankNil(b bool)`

 SetLaneRankNil sets the value for LaneRank to be an explicit nil

### UnsetLaneRank
`func (o *InvestmentGraphEdgesRow) UnsetLaneRank()`

UnsetLaneRank ensures that no value is present for LaneRank, not even an explicit nil
### GetIsPreferred

`func (o *InvestmentGraphEdgesRow) GetIsPreferred() bool`

GetIsPreferred returns the IsPreferred field if non-nil, zero value otherwise.

### GetIsPreferredOk

`func (o *InvestmentGraphEdgesRow) GetIsPreferredOk() (*bool, bool)`

GetIsPreferredOk returns a tuple with the IsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferred

`func (o *InvestmentGraphEdgesRow) SetIsPreferred(v bool)`

SetIsPreferred sets IsPreferred field to given value.

### HasIsPreferred

`func (o *InvestmentGraphEdgesRow) HasIsPreferred() bool`

HasIsPreferred returns a boolean if a field has been set.

### SetIsPreferredNil

`func (o *InvestmentGraphEdgesRow) SetIsPreferredNil(b bool)`

 SetIsPreferredNil sets the value for IsPreferred to be an explicit nil

### UnsetIsPreferred
`func (o *InvestmentGraphEdgesRow) UnsetIsPreferred()`

UnsetIsPreferred ensures that no value is present for IsPreferred, not even an explicit nil
### GetSourceLicense

`func (o *InvestmentGraphEdgesRow) GetSourceLicense() string`

GetSourceLicense returns the SourceLicense field if non-nil, zero value otherwise.

### GetSourceLicenseOk

`func (o *InvestmentGraphEdgesRow) GetSourceLicenseOk() (*string, bool)`

GetSourceLicenseOk returns a tuple with the SourceLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLicense

`func (o *InvestmentGraphEdgesRow) SetSourceLicense(v string)`

SetSourceLicense sets SourceLicense field to given value.

### HasSourceLicense

`func (o *InvestmentGraphEdgesRow) HasSourceLicense() bool`

HasSourceLicense returns a boolean if a field has been set.

### SetSourceLicenseNil

`func (o *InvestmentGraphEdgesRow) SetSourceLicenseNil(b bool)`

 SetSourceLicenseNil sets the value for SourceLicense to be an explicit nil

### UnsetSourceLicense
`func (o *InvestmentGraphEdgesRow) UnsetSourceLicense()`

UnsetSourceLicense ensures that no value is present for SourceLicense, not even an explicit nil
### GetAttribution

`func (o *InvestmentGraphEdgesRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *InvestmentGraphEdgesRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *InvestmentGraphEdgesRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *InvestmentGraphEdgesRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *InvestmentGraphEdgesRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *InvestmentGraphEdgesRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil
### GetFromName

`func (o *InvestmentGraphEdgesRow) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *InvestmentGraphEdgesRow) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *InvestmentGraphEdgesRow) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *InvestmentGraphEdgesRow) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### SetFromNameNil

`func (o *InvestmentGraphEdgesRow) SetFromNameNil(b bool)`

 SetFromNameNil sets the value for FromName to be an explicit nil

### UnsetFromName
`func (o *InvestmentGraphEdgesRow) UnsetFromName()`

UnsetFromName ensures that no value is present for FromName, not even an explicit nil
### GetFromOrgId

`func (o *InvestmentGraphEdgesRow) GetFromOrgId() string`

GetFromOrgId returns the FromOrgId field if non-nil, zero value otherwise.

### GetFromOrgIdOk

`func (o *InvestmentGraphEdgesRow) GetFromOrgIdOk() (*string, bool)`

GetFromOrgIdOk returns a tuple with the FromOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromOrgId

`func (o *InvestmentGraphEdgesRow) SetFromOrgId(v string)`

SetFromOrgId sets FromOrgId field to given value.

### HasFromOrgId

`func (o *InvestmentGraphEdgesRow) HasFromOrgId() bool`

HasFromOrgId returns a boolean if a field has been set.

### SetFromOrgIdNil

`func (o *InvestmentGraphEdgesRow) SetFromOrgIdNil(b bool)`

 SetFromOrgIdNil sets the value for FromOrgId to be an explicit nil

### UnsetFromOrgId
`func (o *InvestmentGraphEdgesRow) UnsetFromOrgId()`

UnsetFromOrgId ensures that no value is present for FromOrgId, not even an explicit nil
### GetFromCompanyId

`func (o *InvestmentGraphEdgesRow) GetFromCompanyId() string`

GetFromCompanyId returns the FromCompanyId field if non-nil, zero value otherwise.

### GetFromCompanyIdOk

`func (o *InvestmentGraphEdgesRow) GetFromCompanyIdOk() (*string, bool)`

GetFromCompanyIdOk returns a tuple with the FromCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCompanyId

`func (o *InvestmentGraphEdgesRow) SetFromCompanyId(v string)`

SetFromCompanyId sets FromCompanyId field to given value.

### HasFromCompanyId

`func (o *InvestmentGraphEdgesRow) HasFromCompanyId() bool`

HasFromCompanyId returns a boolean if a field has been set.

### SetFromCompanyIdNil

`func (o *InvestmentGraphEdgesRow) SetFromCompanyIdNil(b bool)`

 SetFromCompanyIdNil sets the value for FromCompanyId to be an explicit nil

### UnsetFromCompanyId
`func (o *InvestmentGraphEdgesRow) UnsetFromCompanyId()`

UnsetFromCompanyId ensures that no value is present for FromCompanyId, not even an explicit nil
### GetFromCompanyName

`func (o *InvestmentGraphEdgesRow) GetFromCompanyName() string`

GetFromCompanyName returns the FromCompanyName field if non-nil, zero value otherwise.

### GetFromCompanyNameOk

`func (o *InvestmentGraphEdgesRow) GetFromCompanyNameOk() (*string, bool)`

GetFromCompanyNameOk returns a tuple with the FromCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCompanyName

`func (o *InvestmentGraphEdgesRow) SetFromCompanyName(v string)`

SetFromCompanyName sets FromCompanyName field to given value.

### HasFromCompanyName

`func (o *InvestmentGraphEdgesRow) HasFromCompanyName() bool`

HasFromCompanyName returns a boolean if a field has been set.

### SetFromCompanyNameNil

`func (o *InvestmentGraphEdgesRow) SetFromCompanyNameNil(b bool)`

 SetFromCompanyNameNil sets the value for FromCompanyName to be an explicit nil

### UnsetFromCompanyName
`func (o *InvestmentGraphEdgesRow) UnsetFromCompanyName()`

UnsetFromCompanyName ensures that no value is present for FromCompanyName, not even an explicit nil
### GetFromGroupId

`func (o *InvestmentGraphEdgesRow) GetFromGroupId() string`

GetFromGroupId returns the FromGroupId field if non-nil, zero value otherwise.

### GetFromGroupIdOk

`func (o *InvestmentGraphEdgesRow) GetFromGroupIdOk() (*string, bool)`

GetFromGroupIdOk returns a tuple with the FromGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGroupId

`func (o *InvestmentGraphEdgesRow) SetFromGroupId(v string)`

SetFromGroupId sets FromGroupId field to given value.

### HasFromGroupId

`func (o *InvestmentGraphEdgesRow) HasFromGroupId() bool`

HasFromGroupId returns a boolean if a field has been set.

### SetFromGroupIdNil

`func (o *InvestmentGraphEdgesRow) SetFromGroupIdNil(b bool)`

 SetFromGroupIdNil sets the value for FromGroupId to be an explicit nil

### UnsetFromGroupId
`func (o *InvestmentGraphEdgesRow) UnsetFromGroupId()`

UnsetFromGroupId ensures that no value is present for FromGroupId, not even an explicit nil
### GetFromGroupName

`func (o *InvestmentGraphEdgesRow) GetFromGroupName() string`

GetFromGroupName returns the FromGroupName field if non-nil, zero value otherwise.

### GetFromGroupNameOk

`func (o *InvestmentGraphEdgesRow) GetFromGroupNameOk() (*string, bool)`

GetFromGroupNameOk returns a tuple with the FromGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGroupName

`func (o *InvestmentGraphEdgesRow) SetFromGroupName(v string)`

SetFromGroupName sets FromGroupName field to given value.

### HasFromGroupName

`func (o *InvestmentGraphEdgesRow) HasFromGroupName() bool`

HasFromGroupName returns a boolean if a field has been set.

### SetFromGroupNameNil

`func (o *InvestmentGraphEdgesRow) SetFromGroupNameNil(b bool)`

 SetFromGroupNameNil sets the value for FromGroupName to be an explicit nil

### UnsetFromGroupName
`func (o *InvestmentGraphEdgesRow) UnsetFromGroupName()`

UnsetFromGroupName ensures that no value is present for FromGroupName, not even an explicit nil
### GetToName

`func (o *InvestmentGraphEdgesRow) GetToName() string`

GetToName returns the ToName field if non-nil, zero value otherwise.

### GetToNameOk

`func (o *InvestmentGraphEdgesRow) GetToNameOk() (*string, bool)`

GetToNameOk returns a tuple with the ToName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToName

`func (o *InvestmentGraphEdgesRow) SetToName(v string)`

SetToName sets ToName field to given value.

### HasToName

`func (o *InvestmentGraphEdgesRow) HasToName() bool`

HasToName returns a boolean if a field has been set.

### SetToNameNil

`func (o *InvestmentGraphEdgesRow) SetToNameNil(b bool)`

 SetToNameNil sets the value for ToName to be an explicit nil

### UnsetToName
`func (o *InvestmentGraphEdgesRow) UnsetToName()`

UnsetToName ensures that no value is present for ToName, not even an explicit nil
### GetToOrgId

`func (o *InvestmentGraphEdgesRow) GetToOrgId() string`

GetToOrgId returns the ToOrgId field if non-nil, zero value otherwise.

### GetToOrgIdOk

`func (o *InvestmentGraphEdgesRow) GetToOrgIdOk() (*string, bool)`

GetToOrgIdOk returns a tuple with the ToOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToOrgId

`func (o *InvestmentGraphEdgesRow) SetToOrgId(v string)`

SetToOrgId sets ToOrgId field to given value.

### HasToOrgId

`func (o *InvestmentGraphEdgesRow) HasToOrgId() bool`

HasToOrgId returns a boolean if a field has been set.

### SetToOrgIdNil

`func (o *InvestmentGraphEdgesRow) SetToOrgIdNil(b bool)`

 SetToOrgIdNil sets the value for ToOrgId to be an explicit nil

### UnsetToOrgId
`func (o *InvestmentGraphEdgesRow) UnsetToOrgId()`

UnsetToOrgId ensures that no value is present for ToOrgId, not even an explicit nil
### GetToCompanyId

`func (o *InvestmentGraphEdgesRow) GetToCompanyId() string`

GetToCompanyId returns the ToCompanyId field if non-nil, zero value otherwise.

### GetToCompanyIdOk

`func (o *InvestmentGraphEdgesRow) GetToCompanyIdOk() (*string, bool)`

GetToCompanyIdOk returns a tuple with the ToCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCompanyId

`func (o *InvestmentGraphEdgesRow) SetToCompanyId(v string)`

SetToCompanyId sets ToCompanyId field to given value.

### HasToCompanyId

`func (o *InvestmentGraphEdgesRow) HasToCompanyId() bool`

HasToCompanyId returns a boolean if a field has been set.

### SetToCompanyIdNil

`func (o *InvestmentGraphEdgesRow) SetToCompanyIdNil(b bool)`

 SetToCompanyIdNil sets the value for ToCompanyId to be an explicit nil

### UnsetToCompanyId
`func (o *InvestmentGraphEdgesRow) UnsetToCompanyId()`

UnsetToCompanyId ensures that no value is present for ToCompanyId, not even an explicit nil
### GetToCompanyName

`func (o *InvestmentGraphEdgesRow) GetToCompanyName() string`

GetToCompanyName returns the ToCompanyName field if non-nil, zero value otherwise.

### GetToCompanyNameOk

`func (o *InvestmentGraphEdgesRow) GetToCompanyNameOk() (*string, bool)`

GetToCompanyNameOk returns a tuple with the ToCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCompanyName

`func (o *InvestmentGraphEdgesRow) SetToCompanyName(v string)`

SetToCompanyName sets ToCompanyName field to given value.

### HasToCompanyName

`func (o *InvestmentGraphEdgesRow) HasToCompanyName() bool`

HasToCompanyName returns a boolean if a field has been set.

### SetToCompanyNameNil

`func (o *InvestmentGraphEdgesRow) SetToCompanyNameNil(b bool)`

 SetToCompanyNameNil sets the value for ToCompanyName to be an explicit nil

### UnsetToCompanyName
`func (o *InvestmentGraphEdgesRow) UnsetToCompanyName()`

UnsetToCompanyName ensures that no value is present for ToCompanyName, not even an explicit nil
### GetToGroupId

`func (o *InvestmentGraphEdgesRow) GetToGroupId() string`

GetToGroupId returns the ToGroupId field if non-nil, zero value otherwise.

### GetToGroupIdOk

`func (o *InvestmentGraphEdgesRow) GetToGroupIdOk() (*string, bool)`

GetToGroupIdOk returns a tuple with the ToGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroupId

`func (o *InvestmentGraphEdgesRow) SetToGroupId(v string)`

SetToGroupId sets ToGroupId field to given value.

### HasToGroupId

`func (o *InvestmentGraphEdgesRow) HasToGroupId() bool`

HasToGroupId returns a boolean if a field has been set.

### SetToGroupIdNil

`func (o *InvestmentGraphEdgesRow) SetToGroupIdNil(b bool)`

 SetToGroupIdNil sets the value for ToGroupId to be an explicit nil

### UnsetToGroupId
`func (o *InvestmentGraphEdgesRow) UnsetToGroupId()`

UnsetToGroupId ensures that no value is present for ToGroupId, not even an explicit nil
### GetToGroupName

`func (o *InvestmentGraphEdgesRow) GetToGroupName() string`

GetToGroupName returns the ToGroupName field if non-nil, zero value otherwise.

### GetToGroupNameOk

`func (o *InvestmentGraphEdgesRow) GetToGroupNameOk() (*string, bool)`

GetToGroupNameOk returns a tuple with the ToGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroupName

`func (o *InvestmentGraphEdgesRow) SetToGroupName(v string)`

SetToGroupName sets ToGroupName field to given value.

### HasToGroupName

`func (o *InvestmentGraphEdgesRow) HasToGroupName() bool`

HasToGroupName returns a boolean if a field has been set.

### SetToGroupNameNil

`func (o *InvestmentGraphEdgesRow) SetToGroupNameNil(b bool)`

 SetToGroupNameNil sets the value for ToGroupName to be an explicit nil

### UnsetToGroupName
`func (o *InvestmentGraphEdgesRow) UnsetToGroupName()`

UnsetToGroupName ensures that no value is present for ToGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


