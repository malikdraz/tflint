// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsKmsExternalKeyInvalidDescriptionRule checks the pattern is valid
type AwsKmsExternalKeyInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsKmsExternalKeyInvalidDescriptionRule returns new rule with default attributes
func NewAwsKmsExternalKeyInvalidDescriptionRule() *AwsKmsExternalKeyInvalidDescriptionRule {
	return &AwsKmsExternalKeyInvalidDescriptionRule{
		resourceType:  "aws_kms_external_key",
		attributeName: "description",
		max:           8192,
	}
}

// Name returns the rule name
func (r *AwsKmsExternalKeyInvalidDescriptionRule) Name() string {
	return "aws_kms_external_key_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsKmsExternalKeyInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsKmsExternalKeyInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsKmsExternalKeyInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsKmsExternalKeyInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 8192 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
