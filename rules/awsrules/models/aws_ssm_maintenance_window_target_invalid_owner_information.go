// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsSsmMaintenanceWindowTargetInvalidOwnerInformationRule checks the pattern is valid
type AwsSsmMaintenanceWindowTargetInvalidOwnerInformationRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSsmMaintenanceWindowTargetInvalidOwnerInformationRule returns new rule with default attributes
func NewAwsSsmMaintenanceWindowTargetInvalidOwnerInformationRule() *AwsSsmMaintenanceWindowTargetInvalidOwnerInformationRule {
	return &AwsSsmMaintenanceWindowTargetInvalidOwnerInformationRule{
		resourceType:  "aws_ssm_maintenance_window_target",
		attributeName: "owner_information",
		max:           128,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSsmMaintenanceWindowTargetInvalidOwnerInformationRule) Name() string {
	return "aws_ssm_maintenance_window_target_invalid_owner_information"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmMaintenanceWindowTargetInvalidOwnerInformationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmMaintenanceWindowTargetInvalidOwnerInformationRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmMaintenanceWindowTargetInvalidOwnerInformationRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmMaintenanceWindowTargetInvalidOwnerInformationRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"owner_information must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"owner_information must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
