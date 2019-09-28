// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsIAMAccessKeyInvalidStatusRule checks the pattern is valid
type AwsIAMAccessKeyInvalidStatusRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsIAMAccessKeyInvalidStatusRule returns new rule with default attributes
func NewAwsIAMAccessKeyInvalidStatusRule() *AwsIAMAccessKeyInvalidStatusRule {
	return &AwsIAMAccessKeyInvalidStatusRule{
		resourceType:  "aws_iam_access_key",
		attributeName: "status",
		enum: []string{
			"Active",
			"Inactive",
		},
	}
}

// Name returns the rule name
func (r *AwsIAMAccessKeyInvalidStatusRule) Name() string {
	return "aws_iam_access_key_invalid_status"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMAccessKeyInvalidStatusRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMAccessKeyInvalidStatusRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMAccessKeyInvalidStatusRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMAccessKeyInvalidStatusRule) Check(runner *tflint.Runner) error {
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
					`status is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
