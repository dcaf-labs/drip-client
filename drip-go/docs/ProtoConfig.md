# ProtoConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | **string** |  | 
**Granularity** | **string** |  | 
**TriggerDcaSpread** | **int32** |  | 
**BaseWithdrawalSpread** | **int32** |  | 

## Methods

### NewProtoConfig

`func NewProtoConfig(pubkey string, granularity string, triggerDcaSpread int32, baseWithdrawalSpread int32, ) *ProtoConfig`

NewProtoConfig instantiates a new ProtoConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtoConfigWithDefaults

`func NewProtoConfigWithDefaults() *ProtoConfig`

NewProtoConfigWithDefaults instantiates a new ProtoConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *ProtoConfig) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *ProtoConfig) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *ProtoConfig) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.


### GetGranularity

`func (o *ProtoConfig) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *ProtoConfig) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *ProtoConfig) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.


### GetTriggerDcaSpread

`func (o *ProtoConfig) GetTriggerDcaSpread() int32`

GetTriggerDcaSpread returns the TriggerDcaSpread field if non-nil, zero value otherwise.

### GetTriggerDcaSpreadOk

`func (o *ProtoConfig) GetTriggerDcaSpreadOk() (*int32, bool)`

GetTriggerDcaSpreadOk returns a tuple with the TriggerDcaSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerDcaSpread

`func (o *ProtoConfig) SetTriggerDcaSpread(v int32)`

SetTriggerDcaSpread sets TriggerDcaSpread field to given value.


### GetBaseWithdrawalSpread

`func (o *ProtoConfig) GetBaseWithdrawalSpread() int32`

GetBaseWithdrawalSpread returns the BaseWithdrawalSpread field if non-nil, zero value otherwise.

### GetBaseWithdrawalSpreadOk

`func (o *ProtoConfig) GetBaseWithdrawalSpreadOk() (*int32, bool)`

GetBaseWithdrawalSpreadOk returns a tuple with the BaseWithdrawalSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseWithdrawalSpread

`func (o *ProtoConfig) SetBaseWithdrawalSpread(v int32)`

SetBaseWithdrawalSpread sets BaseWithdrawalSpread field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


