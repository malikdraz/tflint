// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsInspectorAssessmentTargetInvalidNameRule checks the pattern is valid
type AwsInspectorAssessmentTargetInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsInspectorAssessmentTargetInvalidNameRule returns new rule with default attributes
func NewAwsInspectorAssessmentTargetInvalidNameRule() *AwsInspectorAssessmentTargetInvalidNameRule {
	return &AwsInspectorAssessmentTargetInvalidNameRule{
		resourceType:  "aws_inspector_assessment_target",
		attributeName: "name",
		max:           140,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsInspectorAssessmentTargetInvalidNameRule) Name() string {
	return "aws_inspector_assessment_target_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsInspectorAssessmentTargetInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsInspectorAssessmentTargetInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsInspectorAssessmentTargetInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsInspectorAssessmentTargetInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 140 characters or less",
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
			return nil
		})
	})
}
