// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsGuarddutyThreatintelsetInvalidFormatRule checks the pattern is valid
type AwsGuarddutyThreatintelsetInvalidFormatRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	enum          []string
}

// NewAwsGuarddutyThreatintelsetInvalidFormatRule returns new rule with default attributes
func NewAwsGuarddutyThreatintelsetInvalidFormatRule() *AwsGuarddutyThreatintelsetInvalidFormatRule {
	return &AwsGuarddutyThreatintelsetInvalidFormatRule{
		resourceType:  "aws_guardduty_threatintelset",
		attributeName: "format",
		max:           300,
		min:           1,
		enum: []string{
			"TXT",
			"STIX",
			"OTX_CSV",
			"ALIEN_VAULT",
			"PROOF_POINT",
			"FIRE_EYE",
		},
	}
}

// Name returns the rule name
func (r *AwsGuarddutyThreatintelsetInvalidFormatRule) Name() string {
	return "aws_guardduty_threatintelset_invalid_format"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGuarddutyThreatintelsetInvalidFormatRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGuarddutyThreatintelsetInvalidFormatRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGuarddutyThreatintelsetInvalidFormatRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGuarddutyThreatintelsetInvalidFormatRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"format must be 300 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"format must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`format is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
