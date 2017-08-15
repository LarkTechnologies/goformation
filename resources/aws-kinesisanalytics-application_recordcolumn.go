package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::Application.RecordColumn AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordcolumn.html
type AWSKinesisAnalyticsApplication_RecordColumn struct {

	// Mapping AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordcolumn.html#cfn-kinesisanalytics-application-recordcolumn-mapping

	Mapping string `json:"Mapping"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordcolumn.html#cfn-kinesisanalytics-application-recordcolumn-name

	Name string `json:"Name"`

	// SqlType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordcolumn.html#cfn-kinesisanalytics-application-recordcolumn-sqltype

	SqlType string `json:"SqlType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_RecordColumn) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.RecordColumn"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplication_RecordColumn) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplication_RecordColumnResources retrieves all AWSKinesisAnalyticsApplication_RecordColumn items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplication_RecordColumn(template *Template) map[string]*AWSKinesisAnalyticsApplication_RecordColumn {

	results := map[string]*AWSKinesisAnalyticsApplication_RecordColumn{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplication_RecordColumn{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplication_RecordColumnWithName retrieves all AWSKinesisAnalyticsApplication_RecordColumn items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplication_RecordColumn(name string, template *Template) (*AWSKinesisAnalyticsApplication_RecordColumn, error) {

	result := &AWSKinesisAnalyticsApplication_RecordColumn{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplication_RecordColumn{}, errors.New("resource not found")

}
