// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsIAMUserPolicyInvalidPolicyRule checks the pattern is valid
type AwsIAMUserPolicyInvalidPolicyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMUserPolicyInvalidPolicyRule returns new rule with default attributes
func NewAwsIAMUserPolicyInvalidPolicyRule() *AwsIAMUserPolicyInvalidPolicyRule {
	return &AwsIAMUserPolicyInvalidPolicyRule{
		resourceType:  "aws_iam_user_policy",
		attributeName: "policy",
		max:           131072,
		min:           1,
		pattern:       regexp.MustCompile(`^[\x{0009}\x{000A}\x{000D}\x{0020}-\x{00FF}]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMUserPolicyInvalidPolicyRule) Name() string {
	return "aws_iam_user_policy_invalid_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMUserPolicyInvalidPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMUserPolicyInvalidPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMUserPolicyInvalidPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMUserPolicyInvalidPolicyRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"policy must be 131072 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"policy must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`policy does not match valid pattern ^[\x{0009}\x{000A}\x{000D}\x{0020}-\x{00FF}]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
