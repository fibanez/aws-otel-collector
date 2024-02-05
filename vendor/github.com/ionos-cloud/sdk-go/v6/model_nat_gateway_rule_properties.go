/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// NatGatewayRuleProperties struct for NatGatewayRuleProperties
type NatGatewayRuleProperties struct {
	// The name of the NAT Gateway rule.
	Name *string `json:"name"`
	// Protocol of the NAT Gateway rule. Defaults to ALL. If protocol is 'ICMP' then targetPortRange start and end cannot be set.
	Protocol *NatGatewayRuleProtocol `json:"protocol,omitempty"`
	// Public IP address of the NAT Gateway rule. Specifies the address used for masking outgoing packets source address field. Should be one of the customer reserved IP address already configured on the NAT Gateway resource
	PublicIp *string `json:"publicIp"`
	// Source subnet of the NAT Gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets source IP address.
	SourceSubnet    *string          `json:"sourceSubnet"`
	TargetPortRange *TargetPortRange `json:"targetPortRange,omitempty"`
	// Target or destination subnet of the NAT Gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets destination IP address. If none is provided, rule will match any address.
	TargetSubnet *string `json:"targetSubnet,omitempty"`
	// Type of the NAT Gateway rule.
	Type *NatGatewayRuleType `json:"type,omitempty"`
}

// NewNatGatewayRuleProperties instantiates a new NatGatewayRuleProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNatGatewayRuleProperties(name string, publicIp string, sourceSubnet string) *NatGatewayRuleProperties {
	this := NatGatewayRuleProperties{}

	this.Name = &name
	this.PublicIp = &publicIp
	this.SourceSubnet = &sourceSubnet

	return &this
}

// NewNatGatewayRulePropertiesWithDefaults instantiates a new NatGatewayRuleProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNatGatewayRulePropertiesWithDefaults() *NatGatewayRuleProperties {
	this := NatGatewayRuleProperties{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, nil is returned
func (o *NatGatewayRuleProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NatGatewayRuleProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *NatGatewayRuleProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *NatGatewayRuleProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetProtocol returns the Protocol field value
// If the value is explicit nil, nil is returned
func (o *NatGatewayRuleProperties) GetProtocol() *NatGatewayRuleProtocol {
	if o == nil {
		return nil
	}

	return o.Protocol

}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NatGatewayRuleProperties) GetProtocolOk() (*NatGatewayRuleProtocol, bool) {
	if o == nil {
		return nil, false
	}

	return o.Protocol, true
}

// SetProtocol sets field value
func (o *NatGatewayRuleProperties) SetProtocol(v NatGatewayRuleProtocol) {

	o.Protocol = &v

}

// HasProtocol returns a boolean if a field has been set.
func (o *NatGatewayRuleProperties) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// GetPublicIp returns the PublicIp field value
// If the value is explicit nil, nil is returned
func (o *NatGatewayRuleProperties) GetPublicIp() *string {
	if o == nil {
		return nil
	}

	return o.PublicIp

}

// GetPublicIpOk returns a tuple with the PublicIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NatGatewayRuleProperties) GetPublicIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.PublicIp, true
}

// SetPublicIp sets field value
func (o *NatGatewayRuleProperties) SetPublicIp(v string) {

	o.PublicIp = &v

}

// HasPublicIp returns a boolean if a field has been set.
func (o *NatGatewayRuleProperties) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// GetSourceSubnet returns the SourceSubnet field value
// If the value is explicit nil, nil is returned
func (o *NatGatewayRuleProperties) GetSourceSubnet() *string {
	if o == nil {
		return nil
	}

	return o.SourceSubnet

}

// GetSourceSubnetOk returns a tuple with the SourceSubnet field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NatGatewayRuleProperties) GetSourceSubnetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.SourceSubnet, true
}

// SetSourceSubnet sets field value
func (o *NatGatewayRuleProperties) SetSourceSubnet(v string) {

	o.SourceSubnet = &v

}

// HasSourceSubnet returns a boolean if a field has been set.
func (o *NatGatewayRuleProperties) HasSourceSubnet() bool {
	if o != nil && o.SourceSubnet != nil {
		return true
	}

	return false
}

// GetTargetPortRange returns the TargetPortRange field value
// If the value is explicit nil, nil is returned
func (o *NatGatewayRuleProperties) GetTargetPortRange() *TargetPortRange {
	if o == nil {
		return nil
	}

	return o.TargetPortRange

}

// GetTargetPortRangeOk returns a tuple with the TargetPortRange field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NatGatewayRuleProperties) GetTargetPortRangeOk() (*TargetPortRange, bool) {
	if o == nil {
		return nil, false
	}

	return o.TargetPortRange, true
}

// SetTargetPortRange sets field value
func (o *NatGatewayRuleProperties) SetTargetPortRange(v TargetPortRange) {

	o.TargetPortRange = &v

}

// HasTargetPortRange returns a boolean if a field has been set.
func (o *NatGatewayRuleProperties) HasTargetPortRange() bool {
	if o != nil && o.TargetPortRange != nil {
		return true
	}

	return false
}

// GetTargetSubnet returns the TargetSubnet field value
// If the value is explicit nil, nil is returned
func (o *NatGatewayRuleProperties) GetTargetSubnet() *string {
	if o == nil {
		return nil
	}

	return o.TargetSubnet

}

// GetTargetSubnetOk returns a tuple with the TargetSubnet field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NatGatewayRuleProperties) GetTargetSubnetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.TargetSubnet, true
}

// SetTargetSubnet sets field value
func (o *NatGatewayRuleProperties) SetTargetSubnet(v string) {

	o.TargetSubnet = &v

}

// HasTargetSubnet returns a boolean if a field has been set.
func (o *NatGatewayRuleProperties) HasTargetSubnet() bool {
	if o != nil && o.TargetSubnet != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, nil is returned
func (o *NatGatewayRuleProperties) GetType() *NatGatewayRuleType {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NatGatewayRuleProperties) GetTypeOk() (*NatGatewayRuleType, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *NatGatewayRuleProperties) SetType(v NatGatewayRuleType) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *NatGatewayRuleProperties) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

func (o NatGatewayRuleProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}

	if o.PublicIp != nil {
		toSerialize["publicIp"] = o.PublicIp
	}

	if o.SourceSubnet != nil {
		toSerialize["sourceSubnet"] = o.SourceSubnet
	}

	if o.TargetPortRange != nil {
		toSerialize["targetPortRange"] = o.TargetPortRange
	}

	if o.TargetSubnet != nil {
		toSerialize["targetSubnet"] = o.TargetSubnet
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	return json.Marshal(toSerialize)
}

type NullableNatGatewayRuleProperties struct {
	value *NatGatewayRuleProperties
	isSet bool
}

func (v NullableNatGatewayRuleProperties) Get() *NatGatewayRuleProperties {
	return v.value
}

func (v *NullableNatGatewayRuleProperties) Set(val *NatGatewayRuleProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableNatGatewayRuleProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableNatGatewayRuleProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNatGatewayRuleProperties(val *NatGatewayRuleProperties) *NullableNatGatewayRuleProperties {
	return &NullableNatGatewayRuleProperties{value: val, isSet: true}
}

func (v NullableNatGatewayRuleProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNatGatewayRuleProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
