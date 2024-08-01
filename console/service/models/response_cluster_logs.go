// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseClusterLogs Logs for cluster
//
// swagger:model Response.ClusterLogs
type ResponseClusterLogs struct {

	// all available logs
	Logs string `json:"logs,omitempty"`
}

// Validate validates this response cluster logs
func (m *ResponseClusterLogs) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this response cluster logs based on context it is used
func (m *ResponseClusterLogs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseClusterLogs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseClusterLogs) UnmarshalBinary(b []byte) error {
	var res ResponseClusterLogs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
