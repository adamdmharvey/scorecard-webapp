// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2021 OpenSSF Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ScorecardResult scorecard result
//
// swagger:model ScorecardResult
type ScorecardResult struct {

	// date
	Date string `json:"date,omitempty"`

	// repo
	Repo *Repo `json:"repo,omitempty"`

	// scorecard
	Scorecard *ScorecardVersion `json:"scorecard,omitempty"`

	// Aggregate score of the repository
	Score float64 `json:"score"`

	// checks
	Checks []*ScorecardCheck `json:"checks"`

	// metadata
	Metadata string `json:"metadata,omitempty"`
}

// Validate validates this scorecard result
func (m *ScorecardResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScorecard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChecks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScorecardResult) validateRepo(formats strfmt.Registry) error {
	if swag.IsZero(m.Repo) { // not required
		return nil
	}

	if m.Repo != nil {
		if err := m.Repo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repo")
			}
			return err
		}
	}

	return nil
}

func (m *ScorecardResult) validateScorecard(formats strfmt.Registry) error {
	if swag.IsZero(m.Scorecard) { // not required
		return nil
	}

	if m.Scorecard != nil {
		if err := m.Scorecard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scorecard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scorecard")
			}
			return err
		}
	}

	return nil
}

func (m *ScorecardResult) validateChecks(formats strfmt.Registry) error {
	if swag.IsZero(m.Checks) { // not required
		return nil
	}

	for i := 0; i < len(m.Checks); i++ {
		if swag.IsZero(m.Checks[i]) { // not required
			continue
		}

		if m.Checks[i] != nil {
			if err := m.Checks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("checks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("checks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this scorecard result based on the context it is used
func (m *ScorecardResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRepo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScorecard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChecks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScorecardResult) contextValidateRepo(ctx context.Context, formats strfmt.Registry) error {

	if m.Repo != nil {
		if err := m.Repo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repo")
			}
			return err
		}
	}

	return nil
}

func (m *ScorecardResult) contextValidateScorecard(ctx context.Context, formats strfmt.Registry) error {

	if m.Scorecard != nil {
		if err := m.Scorecard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scorecard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scorecard")
			}
			return err
		}
	}

	return nil
}

func (m *ScorecardResult) contextValidateChecks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Checks); i++ {

		if m.Checks[i] != nil {
			if err := m.Checks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("checks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("checks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScorecardResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScorecardResult) UnmarshalBinary(b []byte) error {
	var res ScorecardResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
