package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyAdmin = []byte("Admin")
	// TODO: Determine the default value
	DefaultAdmin string = "admin"
)

var (
	KeyCommissionRate = []byte("CommissionRate")
	// TODO: Determine the default value
	DefaultCommissionRate string = "commission_rate"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	admin string,
	commissionRate string,
) Params {
	return Params{
		Admin:          admin,
		CommissionRate: commissionRate,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultAdmin,
		DefaultCommissionRate,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyAdmin, &p.Admin, validateAdmin),
		paramtypes.NewParamSetPair(KeyCommissionRate, &p.CommissionRate, validateCommissionRate),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateAdmin(p.Admin); err != nil {
		return err
	}

	if err := validateCommissionRate(p.CommissionRate); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateAdmin validates the Admin param
func validateAdmin(v interface{}) error {
	admin, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = admin

	return nil
}

// validateCommissionRate validates the CommissionRate param
func validateCommissionRate(v interface{}) error {
	commissionRate, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = commissionRate

	return nil
}
