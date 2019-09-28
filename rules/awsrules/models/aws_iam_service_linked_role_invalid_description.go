// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsIAMServiceLinkedRoleInvalidDescriptionRule checks the pattern is valid
type AwsIAMServiceLinkedRoleInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsIAMServiceLinkedRoleInvalidDescriptionRule returns new rule with default attributes
func NewAwsIAMServiceLinkedRoleInvalidDescriptionRule() *AwsIAMServiceLinkedRoleInvalidDescriptionRule {
	return &AwsIAMServiceLinkedRoleInvalidDescriptionRule{
		resourceType:  "aws_iam_service_linked_role",
		attributeName: "description",
		max:           1000,
		pattern:       regexp.MustCompile(`^[\p{L}\p{M}\p{Z}\p{S}\p{N}\p{P}]*$`),
	}
}

// Name returns the rule name
func (r *AwsIAMServiceLinkedRoleInvalidDescriptionRule) Name() string {
	return "aws_iam_service_linked_role_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMServiceLinkedRoleInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMServiceLinkedRoleInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMServiceLinkedRoleInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMServiceLinkedRoleInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 1000 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`description does not match valid pattern ^[\p{L}\p{M}\p{Z}\p{S}\p{N}\p{P}]*$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
