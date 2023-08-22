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

// KubernetesNodePoolProperties struct for KubernetesNodePoolProperties
type KubernetesNodePoolProperties struct {
	// A Kubernetes node pool name. Valid Kubernetes node pool name must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Name *string `json:"name"`
	// The unique identifier of the VDC where the worker nodes of the node pool are provisioned.Note that the data center is located in the exact place where the parent cluster of the node pool is located.
	DatacenterId *string `json:"datacenterId"`
	// The number of worker nodes of the node pool.
	NodeCount *int32 `json:"nodeCount"`
	// The CPU type for the nodes.
	CpuFamily *string `json:"cpuFamily"`
	// The total number of cores for the nodes.
	CoresCount *int32 `json:"coresCount"`
	// The RAM size for the nodes. Must be specified in multiples of 1024 MB, with a minimum size of 2048 MB.
	RamSize *int32 `json:"ramSize"`
	// The availability zone in which the target VM should be provisioned.
	AvailabilityZone *string `json:"availabilityZone"`
	// The storage type for the nodes.
	StorageType *string `json:"storageType"`
	// The allocated volume size in GB. The allocated volume size in GB. To achieve good performance, we recommend a size greater than 100GB for SSD.
	StorageSize *int32 `json:"storageSize"`
	// The Kubernetes version running in the node pool. Note that this imposes restrictions on which Kubernetes versions can run in the node pools of a cluster. Also, not all Kubernetes versions are suitable upgrade targets for all earlier versions.
	K8sVersion        *string                      `json:"k8sVersion,omitempty"`
	MaintenanceWindow *KubernetesMaintenanceWindow `json:"maintenanceWindow,omitempty"`
	AutoScaling       *KubernetesAutoScaling       `json:"autoScaling,omitempty"`
	// The array of existing private LANs to attach to worker nodes.
	Lans *[]KubernetesNodePoolLan `json:"lans,omitempty"`
	// The labels attached to the node pool.
	Labels *map[string]string `json:"labels,omitempty"`
	// The annotations attached to the node pool.
	Annotations *map[string]string `json:"annotations,omitempty"`
	// Optional array of reserved public IP addresses to be used by the nodes. The IPs must be from the exact location of the node pool's data center. If autoscaling is used, the array must contain one more IP than the maximum possible number of nodes (nodeCount+1 for a fixed number of nodes or maxNodeCount+1). The extra IP is used when the nodes are rebuilt.
	PublicIps *[]string `json:"publicIps,omitempty"`
	// The list of available versions for upgrading the node pool.
	AvailableUpgradeVersions *[]string `json:"availableUpgradeVersions,omitempty"`
}

// NewKubernetesNodePoolProperties instantiates a new KubernetesNodePoolProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodePoolProperties(name string, datacenterId string, nodeCount int32, cpuFamily string, coresCount int32, ramSize int32, availabilityZone string, storageType string, storageSize int32) *KubernetesNodePoolProperties {
	this := KubernetesNodePoolProperties{}

	this.Name = &name
	this.DatacenterId = &datacenterId
	this.NodeCount = &nodeCount
	this.CpuFamily = &cpuFamily
	this.CoresCount = &coresCount
	this.RamSize = &ramSize
	this.AvailabilityZone = &availabilityZone
	this.StorageType = &storageType
	this.StorageSize = &storageSize

	return &this
}

// NewKubernetesNodePoolPropertiesWithDefaults instantiates a new KubernetesNodePoolProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodePoolPropertiesWithDefaults() *KubernetesNodePoolProperties {
	this := KubernetesNodePoolProperties{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *KubernetesNodePoolProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetDatacenterId returns the DatacenterId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolProperties) GetDatacenterId() *string {
	if o == nil {
		return nil
	}

	return o.DatacenterId

}

// GetDatacenterIdOk returns a tuple with the DatacenterId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetDatacenterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.DatacenterId, true
}

// SetDatacenterId sets field value
func (o *KubernetesNodePoolProperties) SetDatacenterId(v string) {

	o.DatacenterId = &v

}

// HasDatacenterId returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasDatacenterId() bool {
	if o != nil && o.DatacenterId != nil {
		return true
	}

	return false
}

// GetNodeCount returns the NodeCount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *KubernetesNodePoolProperties) GetNodeCount() *int32 {
	if o == nil {
		return nil
	}

	return o.NodeCount

}

// GetNodeCountOk returns a tuple with the NodeCount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetNodeCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.NodeCount, true
}

// SetNodeCount sets field value
func (o *KubernetesNodePoolProperties) SetNodeCount(v int32) {

	o.NodeCount = &v

}

// HasNodeCount returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasNodeCount() bool {
	if o != nil && o.NodeCount != nil {
		return true
	}

	return false
}

// GetCpuFamily returns the CpuFamily field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolProperties) GetCpuFamily() *string {
	if o == nil {
		return nil
	}

	return o.CpuFamily

}

// GetCpuFamilyOk returns a tuple with the CpuFamily field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetCpuFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.CpuFamily, true
}

// SetCpuFamily sets field value
func (o *KubernetesNodePoolProperties) SetCpuFamily(v string) {

	o.CpuFamily = &v

}

// HasCpuFamily returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasCpuFamily() bool {
	if o != nil && o.CpuFamily != nil {
		return true
	}

	return false
}

// GetCoresCount returns the CoresCount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *KubernetesNodePoolProperties) GetCoresCount() *int32 {
	if o == nil {
		return nil
	}

	return o.CoresCount

}

// GetCoresCountOk returns a tuple with the CoresCount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetCoresCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.CoresCount, true
}

// SetCoresCount sets field value
func (o *KubernetesNodePoolProperties) SetCoresCount(v int32) {

	o.CoresCount = &v

}

// HasCoresCount returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasCoresCount() bool {
	if o != nil && o.CoresCount != nil {
		return true
	}

	return false
}

// GetRamSize returns the RamSize field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *KubernetesNodePoolProperties) GetRamSize() *int32 {
	if o == nil {
		return nil
	}

	return o.RamSize

}

// GetRamSizeOk returns a tuple with the RamSize field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetRamSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.RamSize, true
}

// SetRamSize sets field value
func (o *KubernetesNodePoolProperties) SetRamSize(v int32) {

	o.RamSize = &v

}

// HasRamSize returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasRamSize() bool {
	if o != nil && o.RamSize != nil {
		return true
	}

	return false
}

// GetAvailabilityZone returns the AvailabilityZone field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolProperties) GetAvailabilityZone() *string {
	if o == nil {
		return nil
	}

	return o.AvailabilityZone

}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value
func (o *KubernetesNodePoolProperties) SetAvailabilityZone(v string) {

	o.AvailabilityZone = &v

}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasAvailabilityZone() bool {
	if o != nil && o.AvailabilityZone != nil {
		return true
	}

	return false
}

// GetStorageType returns the StorageType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolProperties) GetStorageType() *string {
	if o == nil {
		return nil
	}

	return o.StorageType

}

// GetStorageTypeOk returns a tuple with the StorageType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetStorageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.StorageType, true
}

// SetStorageType sets field value
func (o *KubernetesNodePoolProperties) SetStorageType(v string) {

	o.StorageType = &v

}

// HasStorageType returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasStorageType() bool {
	if o != nil && o.StorageType != nil {
		return true
	}

	return false
}

// GetStorageSize returns the StorageSize field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *KubernetesNodePoolProperties) GetStorageSize() *int32 {
	if o == nil {
		return nil
	}

	return o.StorageSize

}

// GetStorageSizeOk returns a tuple with the StorageSize field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetStorageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.StorageSize, true
}

// SetStorageSize sets field value
func (o *KubernetesNodePoolProperties) SetStorageSize(v int32) {

	o.StorageSize = &v

}

// HasStorageSize returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasStorageSize() bool {
	if o != nil && o.StorageSize != nil {
		return true
	}

	return false
}

// GetK8sVersion returns the K8sVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolProperties) GetK8sVersion() *string {
	if o == nil {
		return nil
	}

	return o.K8sVersion

}

// GetK8sVersionOk returns a tuple with the K8sVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetK8sVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.K8sVersion, true
}

// SetK8sVersion sets field value
func (o *KubernetesNodePoolProperties) SetK8sVersion(v string) {

	o.K8sVersion = &v

}

// HasK8sVersion returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasK8sVersion() bool {
	if o != nil && o.K8sVersion != nil {
		return true
	}

	return false
}

// GetMaintenanceWindow returns the MaintenanceWindow field value
// If the value is explicit nil, the zero value for KubernetesMaintenanceWindow will be returned
func (o *KubernetesNodePoolProperties) GetMaintenanceWindow() *KubernetesMaintenanceWindow {
	if o == nil {
		return nil
	}

	return o.MaintenanceWindow

}

// GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetMaintenanceWindowOk() (*KubernetesMaintenanceWindow, bool) {
	if o == nil {
		return nil, false
	}

	return o.MaintenanceWindow, true
}

// SetMaintenanceWindow sets field value
func (o *KubernetesNodePoolProperties) SetMaintenanceWindow(v KubernetesMaintenanceWindow) {

	o.MaintenanceWindow = &v

}

// HasMaintenanceWindow returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasMaintenanceWindow() bool {
	if o != nil && o.MaintenanceWindow != nil {
		return true
	}

	return false
}

// GetAutoScaling returns the AutoScaling field value
// If the value is explicit nil, the zero value for KubernetesAutoScaling will be returned
func (o *KubernetesNodePoolProperties) GetAutoScaling() *KubernetesAutoScaling {
	if o == nil {
		return nil
	}

	return o.AutoScaling

}

// GetAutoScalingOk returns a tuple with the AutoScaling field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetAutoScalingOk() (*KubernetesAutoScaling, bool) {
	if o == nil {
		return nil, false
	}

	return o.AutoScaling, true
}

// SetAutoScaling sets field value
func (o *KubernetesNodePoolProperties) SetAutoScaling(v KubernetesAutoScaling) {

	o.AutoScaling = &v

}

// HasAutoScaling returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasAutoScaling() bool {
	if o != nil && o.AutoScaling != nil {
		return true
	}

	return false
}

// GetLans returns the Lans field value
// If the value is explicit nil, the zero value for []KubernetesNodePoolLan will be returned
func (o *KubernetesNodePoolProperties) GetLans() *[]KubernetesNodePoolLan {
	if o == nil {
		return nil
	}

	return o.Lans

}

// GetLansOk returns a tuple with the Lans field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetLansOk() (*[]KubernetesNodePoolLan, bool) {
	if o == nil {
		return nil, false
	}

	return o.Lans, true
}

// SetLans sets field value
func (o *KubernetesNodePoolProperties) SetLans(v []KubernetesNodePoolLan) {

	o.Lans = &v

}

// HasLans returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasLans() bool {
	if o != nil && o.Lans != nil {
		return true
	}

	return false
}

// GetLabels returns the Labels field value
// If the value is explicit nil, the zero value for map[string]string will be returned
func (o *KubernetesNodePoolProperties) GetLabels() *map[string]string {
	if o == nil {
		return nil
	}

	return o.Labels

}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetLabelsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Labels, true
}

// SetLabels sets field value
func (o *KubernetesNodePoolProperties) SetLabels(v map[string]string) {

	o.Labels = &v

}

// HasLabels returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// GetAnnotations returns the Annotations field value
// If the value is explicit nil, the zero value for map[string]string will be returned
func (o *KubernetesNodePoolProperties) GetAnnotations() *map[string]string {
	if o == nil {
		return nil
	}

	return o.Annotations

}

// GetAnnotationsOk returns a tuple with the Annotations field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Annotations, true
}

// SetAnnotations sets field value
func (o *KubernetesNodePoolProperties) SetAnnotations(v map[string]string) {

	o.Annotations = &v

}

// HasAnnotations returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// GetPublicIps returns the PublicIps field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *KubernetesNodePoolProperties) GetPublicIps() *[]string {
	if o == nil {
		return nil
	}

	return o.PublicIps

}

// GetPublicIpsOk returns a tuple with the PublicIps field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetPublicIpsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.PublicIps, true
}

// SetPublicIps sets field value
func (o *KubernetesNodePoolProperties) SetPublicIps(v []string) {

	o.PublicIps = &v

}

// HasPublicIps returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasPublicIps() bool {
	if o != nil && o.PublicIps != nil {
		return true
	}

	return false
}

// GetAvailableUpgradeVersions returns the AvailableUpgradeVersions field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *KubernetesNodePoolProperties) GetAvailableUpgradeVersions() *[]string {
	if o == nil {
		return nil
	}

	return o.AvailableUpgradeVersions

}

// GetAvailableUpgradeVersionsOk returns a tuple with the AvailableUpgradeVersions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolProperties) GetAvailableUpgradeVersionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.AvailableUpgradeVersions, true
}

// SetAvailableUpgradeVersions sets field value
func (o *KubernetesNodePoolProperties) SetAvailableUpgradeVersions(v []string) {

	o.AvailableUpgradeVersions = &v

}

// HasAvailableUpgradeVersions returns a boolean if a field has been set.
func (o *KubernetesNodePoolProperties) HasAvailableUpgradeVersions() bool {
	if o != nil && o.AvailableUpgradeVersions != nil {
		return true
	}

	return false
}

func (o KubernetesNodePoolProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DatacenterId != nil {
		toSerialize["datacenterId"] = o.DatacenterId
	}
	if o.NodeCount != nil {
		toSerialize["nodeCount"] = o.NodeCount
	}
	if o.CpuFamily != nil {
		toSerialize["cpuFamily"] = o.CpuFamily
	}
	if o.CoresCount != nil {
		toSerialize["coresCount"] = o.CoresCount
	}
	if o.RamSize != nil {
		toSerialize["ramSize"] = o.RamSize
	}
	if o.AvailabilityZone != nil {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}
	if o.StorageType != nil {
		toSerialize["storageType"] = o.StorageType
	}
	if o.StorageSize != nil {
		toSerialize["storageSize"] = o.StorageSize
	}
	if o.K8sVersion != nil {
		toSerialize["k8sVersion"] = o.K8sVersion
	}
	if o.MaintenanceWindow != nil {
		toSerialize["maintenanceWindow"] = o.MaintenanceWindow
	}
	if o.AutoScaling != nil {
		toSerialize["autoScaling"] = o.AutoScaling
	}
	if o.Lans != nil {
		toSerialize["lans"] = o.Lans
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	if o.PublicIps != nil {
		toSerialize["publicIps"] = o.PublicIps
	}
	if o.AvailableUpgradeVersions != nil {
		toSerialize["availableUpgradeVersions"] = o.AvailableUpgradeVersions
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesNodePoolProperties struct {
	value *KubernetesNodePoolProperties
	isSet bool
}

func (v NullableKubernetesNodePoolProperties) Get() *KubernetesNodePoolProperties {
	return v.value
}

func (v *NullableKubernetesNodePoolProperties) Set(val *KubernetesNodePoolProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodePoolProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodePoolProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodePoolProperties(val *KubernetesNodePoolProperties) *NullableKubernetesNodePoolProperties {
	return &NullableKubernetesNodePoolProperties{value: val, isSet: true}
}

func (v NullableKubernetesNodePoolProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodePoolProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
