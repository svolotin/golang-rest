// Copyright (c) 2023 ....

package models

// just an example
type Entity struct {
	Name        string `json:"name" binding:"required" dynamodbav:"EntityName"`
	Code        string `json:"code" binding:"required" dynamodbav:"EntityCode"`
	Description string `json:"description" binding:"required" dynamodbav:"SiteDescription, omitempty" `
	Timezone    string `json:"timezone" binding:"required" dynamodbav:"SiteTimeZone"`
}
