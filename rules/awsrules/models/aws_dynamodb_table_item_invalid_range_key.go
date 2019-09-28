// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsDynamoDBTableItemInvalidRangeKeyRule checks the pattern is valid
type AwsDynamoDBTableItemInvalidRangeKeyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsDynamoDBTableItemInvalidRangeKeyRule returns new rule with default attributes
func NewAwsDynamoDBTableItemInvalidRangeKeyRule() *AwsDynamoDBTableItemInvalidRangeKeyRule {
	return &AwsDynamoDBTableItemInvalidRangeKeyRule{
		resourceType:  "aws_dynamodb_table_item",
		attributeName: "range_key",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsDynamoDBTableItemInvalidRangeKeyRule) Name() string {
	return "aws_dynamodb_table_item_invalid_range_key"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDynamoDBTableItemInvalidRangeKeyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDynamoDBTableItemInvalidRangeKeyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDynamoDBTableItemInvalidRangeKeyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDynamoDBTableItemInvalidRangeKeyRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"range_key must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"range_key must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
