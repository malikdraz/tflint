// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsAppsyncResolverInvalidTypeRule checks the pattern is valid
type AwsAppsyncResolverInvalidTypeRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsAppsyncResolverInvalidTypeRule returns new rule with default attributes
func NewAwsAppsyncResolverInvalidTypeRule() *AwsAppsyncResolverInvalidTypeRule {
	return &AwsAppsyncResolverInvalidTypeRule{
		resourceType:  "aws_appsync_resolver",
		attributeName: "type",
		pattern:       regexp.MustCompile(`^[_A-Za-z][_0-9A-Za-z]*$`),
	}
}

// Name returns the rule name
func (r *AwsAppsyncResolverInvalidTypeRule) Name() string {
	return "aws_appsync_resolver_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppsyncResolverInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppsyncResolverInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppsyncResolverInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppsyncResolverInvalidTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`type does not match valid pattern ^[_A-Za-z][_0-9A-Za-z]*$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
