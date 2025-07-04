// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BGPApplyPolicyToNeighborMod b g p apply policy to neighbor mod
//
// swagger:model BGPApplyPolicyToNeighborMod
type BGPApplyPolicyToNeighborMod struct {

	// BGP Neighbor IP address
	// Required: true
	IPAddress *string `json:"ipAddress"`

	// policies
	Policies []string `json:"policies"`

	// policy type
	// Required: true
	// Enum: [import export]
	PolicyType *string `json:"policyType"`

	// route action
	// Required: true
	// Enum: [accept reject]
	RouteAction *string `json:"routeAction"`
}

// Validate validates this b g p apply policy to neighbor mod
func (m *BGPApplyPolicyToNeighborMod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BGPApplyPolicyToNeighborMod) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ipAddress", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

var bGPApplyPolicyToNeighborModTypePolicyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["import","export"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bGPApplyPolicyToNeighborModTypePolicyTypePropEnum = append(bGPApplyPolicyToNeighborModTypePolicyTypePropEnum, v)
	}
}

const (

	// BGPApplyPolicyToNeighborModPolicyTypeImport captures enum value "import"
	BGPApplyPolicyToNeighborModPolicyTypeImport string = "import"

	// BGPApplyPolicyToNeighborModPolicyTypeExport captures enum value "export"
	BGPApplyPolicyToNeighborModPolicyTypeExport string = "export"
)

// prop value enum
func (m *BGPApplyPolicyToNeighborMod) validatePolicyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bGPApplyPolicyToNeighborModTypePolicyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BGPApplyPolicyToNeighborMod) validatePolicyType(formats strfmt.Registry) error {

	if err := validate.Required("policyType", "body", m.PolicyType); err != nil {
		return err
	}

	// value enum
	if err := m.validatePolicyTypeEnum("policyType", "body", *m.PolicyType); err != nil {
		return err
	}

	return nil
}

var bGPApplyPolicyToNeighborModTypeRouteActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["accept","reject"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bGPApplyPolicyToNeighborModTypeRouteActionPropEnum = append(bGPApplyPolicyToNeighborModTypeRouteActionPropEnum, v)
	}
}

const (

	// BGPApplyPolicyToNeighborModRouteActionAccept captures enum value "accept"
	BGPApplyPolicyToNeighborModRouteActionAccept string = "accept"

	// BGPApplyPolicyToNeighborModRouteActionReject captures enum value "reject"
	BGPApplyPolicyToNeighborModRouteActionReject string = "reject"
)

// prop value enum
func (m *BGPApplyPolicyToNeighborMod) validateRouteActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bGPApplyPolicyToNeighborModTypeRouteActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BGPApplyPolicyToNeighborMod) validateRouteAction(formats strfmt.Registry) error {

	if err := validate.Required("routeAction", "body", m.RouteAction); err != nil {
		return err
	}

	// value enum
	if err := m.validateRouteActionEnum("routeAction", "body", *m.RouteAction); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this b g p apply policy to neighbor mod based on context it is used
func (m *BGPApplyPolicyToNeighborMod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BGPApplyPolicyToNeighborMod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BGPApplyPolicyToNeighborMod) UnmarshalBinary(b []byte) error {
	var res BGPApplyPolicyToNeighborMod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
