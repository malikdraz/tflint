// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsMacieS3BucketAssociationInvalidPrefixRule checks the pattern is valid
type AwsMacieS3BucketAssociationInvalidPrefixRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsMacieS3BucketAssociationInvalidPrefixRule returns new rule with default attributes
func NewAwsMacieS3BucketAssociationInvalidPrefixRule() *AwsMacieS3BucketAssociationInvalidPrefixRule {
	return &AwsMacieS3BucketAssociationInvalidPrefixRule{
		resourceType:  "aws_macie_s3_bucket_association",
		attributeName: "prefix",
		max:           10000,
	}
}

// Name returns the rule name
func (r *AwsMacieS3BucketAssociationInvalidPrefixRule) Name() string {
	return "aws_macie_s3_bucket_association_invalid_prefix"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsMacieS3BucketAssociationInvalidPrefixRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsMacieS3BucketAssociationInvalidPrefixRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsMacieS3BucketAssociationInvalidPrefixRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsMacieS3BucketAssociationInvalidPrefixRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"prefix must be 10000 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
