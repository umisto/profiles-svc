# UpdateProfileDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pseudonym** | Pointer to **string** | pseudonym | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**DeleteAvatar** | **bool** | delete avatar | 

## Methods

### NewUpdateProfileDataAttributes

`func NewUpdateProfileDataAttributes(deleteAvatar bool, ) *UpdateProfileDataAttributes`

NewUpdateProfileDataAttributes instantiates a new UpdateProfileDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProfileDataAttributesWithDefaults

`func NewUpdateProfileDataAttributesWithDefaults() *UpdateProfileDataAttributes`

NewUpdateProfileDataAttributesWithDefaults instantiates a new UpdateProfileDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPseudonym

`func (o *UpdateProfileDataAttributes) GetPseudonym() string`

GetPseudonym returns the Pseudonym field if non-nil, zero value otherwise.

### GetPseudonymOk

`func (o *UpdateProfileDataAttributes) GetPseudonymOk() (*string, bool)`

GetPseudonymOk returns a tuple with the Pseudonym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudonym

`func (o *UpdateProfileDataAttributes) SetPseudonym(v string)`

SetPseudonym sets Pseudonym field to given value.

### HasPseudonym

`func (o *UpdateProfileDataAttributes) HasPseudonym() bool`

HasPseudonym returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateProfileDataAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProfileDataAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProfileDataAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProfileDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeleteAvatar

`func (o *UpdateProfileDataAttributes) GetDeleteAvatar() bool`

GetDeleteAvatar returns the DeleteAvatar field if non-nil, zero value otherwise.

### GetDeleteAvatarOk

`func (o *UpdateProfileDataAttributes) GetDeleteAvatarOk() (*bool, bool)`

GetDeleteAvatarOk returns a tuple with the DeleteAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAvatar

`func (o *UpdateProfileDataAttributes) SetDeleteAvatar(v bool)`

SetDeleteAvatar sets DeleteAvatar field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


