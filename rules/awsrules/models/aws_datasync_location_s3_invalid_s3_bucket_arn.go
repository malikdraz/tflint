// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsDatasyncLocationS3InvalidS3BucketArnRule checks the pattern is valid
type AwsDatasyncLocationS3InvalidS3BucketArnRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncLocationS3InvalidS3BucketArnRule returns new rule with default attributes
func NewAwsDatasyncLocationS3InvalidS3BucketArnRule() *AwsDatasyncLocationS3InvalidS3BucketArnRule {
	return &AwsDatasyncLocationS3InvalidS3BucketArnRule{
		resourceType:  "aws_datasync_location_s3",
		attributeName: "s3_bucket_arn",
		max:           76,
		pattern:       regexp.MustCompile(`^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):s3:::([^/]*)$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncLocationS3InvalidS3BucketArnRule) Name() string {
	return "aws_datasync_location_s3_invalid_s3_bucket_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncLocationS3InvalidS3BucketArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDatasyncLocationS3InvalidS3BucketArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncLocationS3InvalidS3BucketArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncLocationS3InvalidS3BucketArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"s3_bucket_arn must be 76 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`s3_bucket_arn does not match valid pattern ^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):s3:::([^/]*)$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
