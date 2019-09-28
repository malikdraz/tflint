// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsDatasyncTaskInvalidSourceLocationArnRule checks the pattern is valid
type AwsDatasyncTaskInvalidSourceLocationArnRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncTaskInvalidSourceLocationArnRule returns new rule with default attributes
func NewAwsDatasyncTaskInvalidSourceLocationArnRule() *AwsDatasyncTaskInvalidSourceLocationArnRule {
	return &AwsDatasyncTaskInvalidSourceLocationArnRule{
		resourceType:  "aws_datasync_task",
		attributeName: "source_location_arn",
		max:           128,
		pattern:       regexp.MustCompile(`^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncTaskInvalidSourceLocationArnRule) Name() string {
	return "aws_datasync_task_invalid_source_location_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncTaskInvalidSourceLocationArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDatasyncTaskInvalidSourceLocationArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncTaskInvalidSourceLocationArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncTaskInvalidSourceLocationArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"source_location_arn must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`source_location_arn does not match valid pattern ^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
