// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package segment holds the schemas of the segment resource and data sources.
// A segment is a group of entities, defined by rules on their attributes, that features and properties can target.
package segment

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Field names of the API that only this package uses.
const (
	attributeName = "attribute_name"
	operator      = "operator"
	values        = "values"
)

const (
	resourceName  = "segment"
	resourcesName = "segments"
)

var operators = []string{
	"contains", "endsWith", "greaterThan", "greaterThanEquals", "is", "isNot",
	"lesserThan", "lesserThanEquals", "notContains", "notEndsWith", "notStartsWith", "startsWith",
}

// getFields returns the fields of a segment, as the user writes them (Input) or as the API returns them (Output).
func getFields(mode utils.Mode) map[string]*schema.Schema {
	return utils.Merge(
		utils.GetIdentityFields(mode, resourceName),
		utils.GetAuditFields(resourceName),
		map[string]*schema.Schema{
			utils.SegmentID: mode.RequiredForceNew(utils.NewString(
				"Segment id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.",
			)),
			utils.Rules: mode.Required(utils.NewList(
				"Rules that determine if an entity belongs to the segment during feature and property evaluation. "+
					"An entity belongs to the segment when it matches any of the rules.",
				map[string]*schema.Schema{
					attributeName: mode.Required(utils.NewString("Attribute name.")),
					operator: mode.Required(utils.NewEnum(
						"Operator to be used for the evaluation if the entity belongs to the segment.", operators...,
					)),
					values: mode.Required(utils.NewStringList(
						"List of values. Entities matching any of the given values belong to the segment.",
					)),
				},
			)),
			utils.Features:         utils.GetAssociatedSchema("feature", utils.FeatureID, resourceName),
			utils.Properties:       utils.GetAssociatedSchema("property", utils.PropertyID, resourceName),
			utils.WorkflowApproval: utils.GetWorkflowApprovalSchema(),
		},
	)
}

var idKeys = []string{utils.SegmentID}

// DataSourceIBMAppConfigSegment returns the `ibm_app_config_segment` data source.
func DataSourceIBMAppConfigSegment() *schema.Resource {
	return utils.NewDataSource("Segment", idKeys, getFields(utils.Output), readSegment,
		map[string]*schema.Schema{
			utils.Include: utils.GetIncludeValuesSchema(utils.Features, utils.Properties, utils.WorkflowApproval),
		},
	)
}

// DataSourceIBMAppConfigSegments returns the `ibm_app_config_segments` data source.
func DataSourceIBMAppConfigSegments() *schema.Resource {
	return utils.NewListDataSource(resourcesName, "Array of segments.", getFields(utils.Output), readSegments,
		utils.GetPageQuery(),
		map[string]*schema.Schema{
			utils.Expand:  utils.GetExpandSchema(),
			utils.Sort:    utils.GetSortValuesSchema(resourceName),
			utils.Tags:    utils.GetOptionalTagsSchema(),
			utils.Include: utils.GetIncludeValuesSchema(utils.Rules, utils.WorkflowApproval),
			utils.Search:  utils.GetSearchSchema("Name OR Tag"),
		},
	)
}

// ResourceIBMAppConfigSegment returns the `ibm_app_config_segment` resource.
func ResourceIBMAppConfigSegment() *schema.Resource {
	return utils.NewResource("Segment", idKeys, getFields(utils.Input), utils.Calls{
		Create: createSegment,
		Get:    getSegment,
		Update: updateSegment,
		Delete: deleteSegment,
	})
}
