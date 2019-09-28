// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudwatchEventRuleInvalidNameRule checks the pattern is valid
type AwsCloudwatchEventRuleInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchEventRuleInvalidNameRule returns new rule with default attributes
func NewAwsCloudwatchEventRuleInvalidNameRule() *AwsCloudwatchEventRuleInvalidNameRule {
	return &AwsCloudwatchEventRuleInvalidNameRule{
		resourceType:  "aws_cloudwatch_event_rule",
		attributeName: "name",
		max:           64,
		min:           1,
		pattern:       regexp.MustCompile(`^[\.\-_A-Za-z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventRuleInvalidNameRule) Name() string {
	return "aws_cloudwatch_event_rule_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventRuleInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventRuleInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventRuleInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventRuleInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`name does not match valid pattern ^[\.\-_A-Za-z0-9]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
