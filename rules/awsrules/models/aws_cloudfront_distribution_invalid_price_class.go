// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudfrontDistributionInvalidPriceClassRule checks the pattern is valid
type AwsCloudfrontDistributionInvalidPriceClassRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCloudfrontDistributionInvalidPriceClassRule returns new rule with default attributes
func NewAwsCloudfrontDistributionInvalidPriceClassRule() *AwsCloudfrontDistributionInvalidPriceClassRule {
	return &AwsCloudfrontDistributionInvalidPriceClassRule{
		resourceType:  "aws_cloudfront_distribution",
		attributeName: "price_class",
		enum: []string{
			"PriceClass_100",
			"PriceClass_200",
			"PriceClass_All",
		},
	}
}

// Name returns the rule name
func (r *AwsCloudfrontDistributionInvalidPriceClassRule) Name() string {
	return "aws_cloudfront_distribution_invalid_price_class"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudfrontDistributionInvalidPriceClassRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudfrontDistributionInvalidPriceClassRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudfrontDistributionInvalidPriceClassRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudfrontDistributionInvalidPriceClassRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`price_class is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
