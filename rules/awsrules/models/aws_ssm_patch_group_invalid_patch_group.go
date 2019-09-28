// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsSsmPatchGroupInvalidPatchGroupRule checks the pattern is valid
type AwsSsmPatchGroupInvalidPatchGroupRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSsmPatchGroupInvalidPatchGroupRule returns new rule with default attributes
func NewAwsSsmPatchGroupInvalidPatchGroupRule() *AwsSsmPatchGroupInvalidPatchGroupRule {
	return &AwsSsmPatchGroupInvalidPatchGroupRule{
		resourceType:  "aws_ssm_patch_group",
		attributeName: "patch_group",
		max:           256,
		min:           1,
		pattern:       regexp.MustCompile(`^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`),
	}
}

// Name returns the rule name
func (r *AwsSsmPatchGroupInvalidPatchGroupRule) Name() string {
	return "aws_ssm_patch_group_invalid_patch_group"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmPatchGroupInvalidPatchGroupRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmPatchGroupInvalidPatchGroupRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmPatchGroupInvalidPatchGroupRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmPatchGroupInvalidPatchGroupRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"patch_group must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"patch_group must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`patch_group does not match valid pattern ^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
