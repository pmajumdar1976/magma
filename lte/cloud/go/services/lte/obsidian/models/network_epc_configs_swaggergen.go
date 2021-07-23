// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkEpcConfigs EPC (evolved packet core) cellular configuration for a network
// swagger:model network_epc_configs
type NetworkEpcConfigs struct {

	// cloud subscriberdb enabled
	CloudSubscriberdbEnabled bool `json:"cloud_subscriberdb_enabled,omitempty"`

	// Network configuration flag for congestion control on EPC
	CongestionControlEnabled *bool `json:"congestion_control_enabled,omitempty"`

	// default rule id
	DefaultRuleID string `json:"default_rule_id,omitempty"`

	// gx gy relay enabled
	// Required: true
	GxGyRelayEnabled *bool `json:"gx_gy_relay_enabled"`

	// hss relay enabled
	// Required: true
	HssRelayEnabled *bool `json:"hss_relay_enabled"`

	// lte auth amf
	// Required: true
	// Format: byte
	LteAuthAmf strfmt.Base64 `json:"lte_auth_amf"`

	// lte auth op
	// Required: true
	// Max Length: 16
	// Min Length: 15
	// Format: byte
	LteAuthOp strfmt.Base64 `json:"lte_auth_op"`

	// mcc
	// Required: true
	// Pattern: ^(\d{3})$
	Mcc string `json:"mcc"`

	// mnc
	// Required: true
	// Pattern: ^(\d{2,3})$
	Mnc string `json:"mnc"`

	// mobility
	Mobility *NetworkEpcConfigsMobility `json:"mobility,omitempty"`

	// Configuration for network services. Services will be instantiated in the listed order.
	NetworkServices []string `json:"network_services,omitempty"`

	// List of IMEIs restricted in the network
	RestrictedImeis []*Imei `json:"restricted_imeis,omitempty"`

	// List of PLMN IDs restricted in the network
	RestrictedPlmns []*PlmnConfig `json:"restricted_plmns,omitempty"`

	// Mapping service areas to tacs in the network
	ServiceAreaMaps map[string]TacList `json:"service_area_maps,omitempty"`

	// sub profiles
	SubProfiles map[string]NetworkEpcConfigsSubProfilesAnon `json:"sub_profiles,omitempty"`

	// subscriberdb sync interval
	SubscriberdbSyncInterval SubscriberdbSyncInterval `json:"subscriberdb_sync_interval,omitempty"`

	// tac
	// Required: true
	// Maximum: 65535
	// Minimum: 1
	Tac uint32 `json:"tac"`
}

// Validate validates this network epc configs
func (m *NetworkEpcConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGxGyRelayEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHssRelayEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLteAuthAmf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLteAuthOp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMcc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMnc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMobility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestrictedImeis(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestrictedPlmns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceAreaMaps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberdbSyncInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTac(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkEpcConfigs) validateGxGyRelayEnabled(formats strfmt.Registry) error {

	if err := validate.Required("gx_gy_relay_enabled", "body", m.GxGyRelayEnabled); err != nil {
		return err
	}

	return nil
}

func (m *NetworkEpcConfigs) validateHssRelayEnabled(formats strfmt.Registry) error {

	if err := validate.Required("hss_relay_enabled", "body", m.HssRelayEnabled); err != nil {
		return err
	}

	return nil
}

func (m *NetworkEpcConfigs) validateLteAuthAmf(formats strfmt.Registry) error {

	if err := validate.Required("lte_auth_amf", "body", strfmt.Base64(m.LteAuthAmf)); err != nil {
		return err
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *NetworkEpcConfigs) validateLteAuthOp(formats strfmt.Registry) error {

	if err := validate.Required("lte_auth_op", "body", strfmt.Base64(m.LteAuthOp)); err != nil {
		return err
	}

	if err := validate.MinLength("lte_auth_op", "body", string(m.LteAuthOp), 15); err != nil {
		return err
	}

	if err := validate.MaxLength("lte_auth_op", "body", string(m.LteAuthOp), 16); err != nil {
		return err
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *NetworkEpcConfigs) validateMcc(formats strfmt.Registry) error {

	if err := validate.RequiredString("mcc", "body", string(m.Mcc)); err != nil {
		return err
	}

	if err := validate.Pattern("mcc", "body", string(m.Mcc), `^(\d{3})$`); err != nil {
		return err
	}

	return nil
}

func (m *NetworkEpcConfigs) validateMnc(formats strfmt.Registry) error {

	if err := validate.RequiredString("mnc", "body", string(m.Mnc)); err != nil {
		return err
	}

	if err := validate.Pattern("mnc", "body", string(m.Mnc), `^(\d{2,3})$`); err != nil {
		return err
	}

	return nil
}

func (m *NetworkEpcConfigs) validateMobility(formats strfmt.Registry) error {

	if swag.IsZero(m.Mobility) { // not required
		return nil
	}

	if m.Mobility != nil {
		if err := m.Mobility.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mobility")
			}
			return err
		}
	}

	return nil
}

var networkEpcConfigsNetworkServicesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dpi","policy_enforcement"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkEpcConfigsNetworkServicesItemsEnum = append(networkEpcConfigsNetworkServicesItemsEnum, v)
	}
}

func (m *NetworkEpcConfigs) validateNetworkServicesItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, networkEpcConfigsNetworkServicesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *NetworkEpcConfigs) validateNetworkServices(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkServices) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkServices); i++ {

		// value enum
		if err := m.validateNetworkServicesItemsEnum("network_services"+"."+strconv.Itoa(i), "body", m.NetworkServices[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *NetworkEpcConfigs) validateRestrictedImeis(formats strfmt.Registry) error {

	if swag.IsZero(m.RestrictedImeis) { // not required
		return nil
	}

	for i := 0; i < len(m.RestrictedImeis); i++ {
		if swag.IsZero(m.RestrictedImeis[i]) { // not required
			continue
		}

		if m.RestrictedImeis[i] != nil {
			if err := m.RestrictedImeis[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("restricted_imeis" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkEpcConfigs) validateRestrictedPlmns(formats strfmt.Registry) error {

	if swag.IsZero(m.RestrictedPlmns) { // not required
		return nil
	}

	for i := 0; i < len(m.RestrictedPlmns); i++ {
		if swag.IsZero(m.RestrictedPlmns[i]) { // not required
			continue
		}

		if m.RestrictedPlmns[i] != nil {
			if err := m.RestrictedPlmns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("restricted_plmns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkEpcConfigs) validateServiceAreaMaps(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceAreaMaps) { // not required
		return nil
	}

	for k := range m.ServiceAreaMaps {

		if err := m.ServiceAreaMaps[k].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_area_maps" + "." + k)
			}
			return err
		}

	}

	return nil
}

func (m *NetworkEpcConfigs) validateSubProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.SubProfiles) { // not required
		return nil
	}

	for k := range m.SubProfiles {

		if swag.IsZero(m.SubProfiles[k]) { // not required
			continue
		}
		if val, ok := m.SubProfiles[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *NetworkEpcConfigs) validateSubscriberdbSyncInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriberdbSyncInterval) { // not required
		return nil
	}

	if err := m.SubscriberdbSyncInterval.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subscriberdb_sync_interval")
		}
		return err
	}

	return nil
}

func (m *NetworkEpcConfigs) validateTac(formats strfmt.Registry) error {

	if err := validate.Required("tac", "body", uint32(m.Tac)); err != nil {
		return err
	}

	if err := validate.MinimumInt("tac", "body", int64(m.Tac), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("tac", "body", int64(m.Tac), 65535, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkEpcConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkEpcConfigs) UnmarshalBinary(b []byte) error {
	var res NetworkEpcConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkEpcConfigsMobility Configuration for IP Allocation (Mobility).
// swagger:model NetworkEpcConfigsMobility
type NetworkEpcConfigsMobility struct {

	// enable multi apn ip allocation
	EnableMultiApnIPAllocation bool `json:"enable_multi_apn_ip_allocation,omitempty"`

	// enable static ip assignments
	EnableStaticIPAssignments bool `json:"enable_static_ip_assignments,omitempty"`

	// ip allocation mode
	// Required: true
	// Enum: [NAT STATIC DHCP_PASSTHROUGH DHCP_BROADCAST]
	IPAllocationMode string `json:"ip_allocation_mode"`

	// nat
	Nat *NetworkEpcConfigsMobilityNat `json:"nat,omitempty"`

	// reserved addresses
	ReservedAddresses []strfmt.IPv4 `json:"reserved_addresses"`

	// static
	Static *NetworkEpcConfigsMobilityStatic `json:"static,omitempty"`
}

// Validate validates this network epc configs mobility
func (m *NetworkEpcConfigsMobility) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAllocationMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReservedAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var networkEpcConfigsMobilityTypeIPAllocationModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NAT","STATIC","DHCP_PASSTHROUGH","DHCP_BROADCAST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkEpcConfigsMobilityTypeIPAllocationModePropEnum = append(networkEpcConfigsMobilityTypeIPAllocationModePropEnum, v)
	}
}

const (

	// NetworkEpcConfigsMobilityIPAllocationModeNAT captures enum value "NAT"
	NetworkEpcConfigsMobilityIPAllocationModeNAT string = "NAT"

	// NetworkEpcConfigsMobilityIPAllocationModeSTATIC captures enum value "STATIC"
	NetworkEpcConfigsMobilityIPAllocationModeSTATIC string = "STATIC"

	// NetworkEpcConfigsMobilityIPAllocationModeDHCPPASSTHROUGH captures enum value "DHCP_PASSTHROUGH"
	NetworkEpcConfigsMobilityIPAllocationModeDHCPPASSTHROUGH string = "DHCP_PASSTHROUGH"

	// NetworkEpcConfigsMobilityIPAllocationModeDHCPBROADCAST captures enum value "DHCP_BROADCAST"
	NetworkEpcConfigsMobilityIPAllocationModeDHCPBROADCAST string = "DHCP_BROADCAST"
)

// prop value enum
func (m *NetworkEpcConfigsMobility) validateIPAllocationModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, networkEpcConfigsMobilityTypeIPAllocationModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NetworkEpcConfigsMobility) validateIPAllocationMode(formats strfmt.Registry) error {

	if err := validate.RequiredString("mobility"+"."+"ip_allocation_mode", "body", string(m.IPAllocationMode)); err != nil {
		return err
	}

	// value enum
	if err := m.validateIPAllocationModeEnum("mobility"+"."+"ip_allocation_mode", "body", m.IPAllocationMode); err != nil {
		return err
	}

	return nil
}

func (m *NetworkEpcConfigsMobility) validateNat(formats strfmt.Registry) error {

	if swag.IsZero(m.Nat) { // not required
		return nil
	}

	if m.Nat != nil {
		if err := m.Nat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mobility" + "." + "nat")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkEpcConfigsMobility) validateReservedAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.ReservedAddresses) { // not required
		return nil
	}

	for i := 0; i < len(m.ReservedAddresses); i++ {

		if err := validate.FormatOf("mobility"+"."+"reserved_addresses"+"."+strconv.Itoa(i), "body", "ipv4", m.ReservedAddresses[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *NetworkEpcConfigsMobility) validateStatic(formats strfmt.Registry) error {

	if swag.IsZero(m.Static) { // not required
		return nil
	}

	if m.Static != nil {
		if err := m.Static.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mobility" + "." + "static")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkEpcConfigsMobility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkEpcConfigsMobility) UnmarshalBinary(b []byte) error {
	var res NetworkEpcConfigsMobility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkEpcConfigsMobilityNat network epc configs mobility nat
// swagger:model NetworkEpcConfigsMobilityNat
type NetworkEpcConfigsMobilityNat struct {

	// ip blocks
	IPBlocks []string `json:"ip_blocks"`
}

// Validate validates this network epc configs mobility nat
func (m *NetworkEpcConfigsMobilityNat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPBlocks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkEpcConfigsMobilityNat) validateIPBlocks(formats strfmt.Registry) error {

	if swag.IsZero(m.IPBlocks) { // not required
		return nil
	}

	for i := 0; i < len(m.IPBlocks); i++ {

		if err := validate.MinLength("mobility"+"."+"nat"+"."+"ip_blocks"+"."+strconv.Itoa(i), "body", string(m.IPBlocks[i]), 5); err != nil {
			return err
		}

		if err := validate.MaxLength("mobility"+"."+"nat"+"."+"ip_blocks"+"."+strconv.Itoa(i), "body", string(m.IPBlocks[i]), 49); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkEpcConfigsMobilityNat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkEpcConfigsMobilityNat) UnmarshalBinary(b []byte) error {
	var res NetworkEpcConfigsMobilityNat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkEpcConfigsMobilityStatic network epc configs mobility static
// swagger:model NetworkEpcConfigsMobilityStatic
type NetworkEpcConfigsMobilityStatic struct {

	// ip blocks by tac
	IPBlocksByTac map[string][]string `json:"ip_blocks_by_tac,omitempty"`
}

// Validate validates this network epc configs mobility static
func (m *NetworkEpcConfigsMobilityStatic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPBlocksByTac(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkEpcConfigsMobilityStatic) validateIPBlocksByTac(formats strfmt.Registry) error {

	if swag.IsZero(m.IPBlocksByTac) { // not required
		return nil
	}

	for k := range m.IPBlocksByTac {

		for i := 0; i < len(m.IPBlocksByTac[k]); i++ {

			if err := validate.MinLength("mobility"+"."+"static"+"."+"ip_blocks_by_tac"+"."+k+"."+strconv.Itoa(i), "body", string(m.IPBlocksByTac[k][i]), 5); err != nil {
				return err
			}

			if err := validate.MaxLength("mobility"+"."+"static"+"."+"ip_blocks_by_tac"+"."+k+"."+strconv.Itoa(i), "body", string(m.IPBlocksByTac[k][i]), 49); err != nil {
				return err
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkEpcConfigsMobilityStatic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkEpcConfigsMobilityStatic) UnmarshalBinary(b []byte) error {
	var res NetworkEpcConfigsMobilityStatic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkEpcConfigsSubProfilesAnon network epc configs sub profiles anon
// swagger:model NetworkEpcConfigsSubProfilesAnon
type NetworkEpcConfigsSubProfilesAnon struct {

	// max dl bit rate
	// Required: true
	// Minimum: > 0
	MaxDlBitRate uint64 `json:"max_dl_bit_rate"`

	// max ul bit rate
	// Required: true
	// Minimum: > 0
	MaxUlBitRate uint64 `json:"max_ul_bit_rate"`
}

// Validate validates this network epc configs sub profiles anon
func (m *NetworkEpcConfigsSubProfilesAnon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaxDlBitRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxUlBitRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkEpcConfigsSubProfilesAnon) validateMaxDlBitRate(formats strfmt.Registry) error {

	if err := validate.Required("max_dl_bit_rate", "body", uint64(m.MaxDlBitRate)); err != nil {
		return err
	}

	if err := validate.MinimumInt("max_dl_bit_rate", "body", int64(m.MaxDlBitRate), 0, true); err != nil {
		return err
	}

	return nil
}

func (m *NetworkEpcConfigsSubProfilesAnon) validateMaxUlBitRate(formats strfmt.Registry) error {

	if err := validate.Required("max_ul_bit_rate", "body", uint64(m.MaxUlBitRate)); err != nil {
		return err
	}

	if err := validate.MinimumInt("max_ul_bit_rate", "body", int64(m.MaxUlBitRate), 0, true); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkEpcConfigsSubProfilesAnon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkEpcConfigsSubProfilesAnon) UnmarshalBinary(b []byte) error {
	var res NetworkEpcConfigsSubProfilesAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}