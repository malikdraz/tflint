// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsSsmAssociationInvalidAssociationNameRule checks the pattern is valid
type AwsSsmAssociationInvalidAssociationNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsSsmAssociationInvalidAssociationNameRule returns new rule with default attributes
func NewAwsSsmAssociationInvalidAssociationNameRule() *AwsSsmAssociationInvalidAssociationNameRule {
	return &AwsSsmAssociationInvalidAssociationNameRule{
		resourceType:  "aws_ssm_association",
		attributeName: "association_name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_\-.]{3,128}$`),
	}
}

// Name returns the rule name
func (r *AwsSsmAssociationInvalidAssociationNameRule) Name() string {
	return "aws_ssm_association_invalid_association_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmAssociationInvalidAssociationNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmAssociationInvalidAssociationNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmAssociationInvalidAssociationNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmAssociationInvalidAssociationNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`association_name does not match valid pattern ^[a-zA-Z0-9_\-.]{3,128}$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
