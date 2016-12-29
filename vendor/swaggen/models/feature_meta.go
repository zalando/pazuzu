package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// FeatureMeta feature meta
// swagger:model FeatureMeta
type FeatureMeta struct {

	// Name of the feature author
	Author string `json:"author,omitempty"`

	// array of feature names
	Dependencies []string `json:"dependencies"`

	// Description of the feature
	Description string `json:"description,omitempty"`

	// Unique identifier representing a specific feature.
	Name string `json:"name,omitempty"`

	// Last update date in ISO 8601 format
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this feature meta
func (m *FeatureMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDependencies(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeatureMeta) validateDependencies(formats strfmt.Registry) error {

	if swag.IsZero(m.Dependencies) { // not required
		return nil
	}

	return nil
}
