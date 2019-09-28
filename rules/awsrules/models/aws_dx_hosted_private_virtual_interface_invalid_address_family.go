// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsDxHostedPrivateVirtualInterfaceInvalidAddressFamilyRule checks the pattern is valid
type AwsDxHostedPrivateVirtualInterfaceInvalidAddressFamilyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDxHostedPrivateVirtualInterfaceInvalidAddressFamilyRule returns new rule with default attributes
func NewAwsDxHostedPrivateVirtualInterfaceInvalidAddressFamilyRule() *AwsDxHostedPrivateVirtualInterfaceInvalidAddressFamilyRule {
	return &AwsDxHostedPrivateVirtualInterfaceInvalidAddressFamilyRule{
		resourceType:  "aws_dx_hosted_private_virtual_interface",
		attributeName: "address_family",
		enum: []string{
			"ipv4",
			"ipv6",
		},
	}
}

// Name returns the rule name
func (r *AwsDxHostedPrivateVirtualInterfaceInvalidAddressFamilyRule) Name() string {
	return "aws_dx_hosted_private_virtual_interface_invalid_address_family"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDxHostedPrivateVirtualInterfaceInvalidAddressFamilyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDxHostedPrivateVirtualInterfaceInvalidAddressFamilyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDxHostedPrivateVirtualInterfaceInvalidAddressFamilyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDxHostedPrivateVirtualInterfaceInvalidAddressFamilyRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`address_family is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
