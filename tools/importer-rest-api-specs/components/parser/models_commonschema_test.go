package parser

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/models"
)

// TODO: Edge Zones
// TODO: Identity System Or User Assigned List
// TODO: Identity System Or User Assigned Map
// TODO: Identity User Assigned List
// TODO: Identity User Assigned Map
// TODO: SystemData
// TODO: Tags
// TODO: Zone
// TODO: Zones

func TestParseModel_CommonSchema_IdentityLegacySystemAndUserAssignedList(t *testing.T) {
	actual, err := ParseSwaggerFileForTesting(t, "model_commonschema_identitylegacysystemanduserassignedlist.json")
	if err != nil {
		t.Fatalf("parsing: %+v", err)
	}

	expected := models.AzureApiDefinition{
		ServiceName: "Example",
		ApiVersion:  "2020-01-01",
		Resources: map[string]models.AzureApiResource{
			"Resource": {
				Models: map[string]models.ModelDetails{
					"Model": {
						Fields: map[string]models.FieldDetails{
							"Identity": {
								JsonName:        "identity",
								CustomFieldType: pointer.To(models.CustomFieldTypeLegacySystemAndUserAssignedIdentityList),
								Required:        false,
							},
							"Name": {
								JsonName: "name",
								ObjectDefinition: &models.ObjectDefinition{
									Type: models.ObjectDefinitionString,
								},
								Required: false,
							},
						},
					},
				},
				Operations: map[string]models.OperationDetails{
					"Test": {
						ContentType:         "application/json",
						ExpectedStatusCodes: []int{200},
						Method:              "GET",
						OperationId:         "Resource_Test",
						ResponseObject: &models.ObjectDefinition{
							ReferenceName: pointer.To("Model"),
							Type:          models.ObjectDefinitionReference,
						},
						UriSuffix: pointer.To("/example"),
					},
				},
			},
		},
	}
	validateParsedSwaggerResultMatches(t, expected, actual)
}

func TestParseModel_CommonSchema_IdentityLegacySystemAndUserAssignedMap(t *testing.T) {
	actual, err := ParseSwaggerFileForTesting(t, "model_commonschema_identitylegacysystemanduserassignedmap.json")
	if err != nil {
		t.Fatalf("parsing: %+v", err)
	}

	expected := models.AzureApiDefinition{
		ServiceName: "Example",
		ApiVersion:  "2020-01-01",
		Resources: map[string]models.AzureApiResource{
			"Resource": {
				Models: map[string]models.ModelDetails{
					"Model": {
						Fields: map[string]models.FieldDetails{
							"Identity": {
								JsonName:        "identity",
								CustomFieldType: pointer.To(models.CustomFieldTypeLegacySystemAndUserAssignedIdentityMap),
								Required:        false,
							},
							"Name": {
								JsonName: "name",
								ObjectDefinition: &models.ObjectDefinition{
									Type: models.ObjectDefinitionString,
								},
								Required: false,
							},
						},
					},
				},
				Operations: map[string]models.OperationDetails{
					"Test": {
						ContentType:         "application/json",
						ExpectedStatusCodes: []int{200},
						Method:              "GET",
						OperationId:         "Resource_Test",
						ResponseObject: &models.ObjectDefinition{
							ReferenceName: pointer.To("Model"),
							Type:          models.ObjectDefinitionReference,
						},
						UriSuffix: pointer.To("/example"),
					},
				},
			},
		},
	}
	validateParsedSwaggerResultMatches(t, expected, actual)
}

func TestParseModel_CommonSchema_IdentityLegacySystemAndUserAssignedMap_ExtraFields(t *testing.T) {
	// this handles the scenario of a System Assigned & User Assigned model having extra fields
	// in the user assigned identity block (#1066) - for example an extra `appId`
	actual, err := ParseSwaggerFileForTesting(t, "model_commonschema_identitylegacysystemanduserassignedmap_extrafields.json")
	if err != nil {
		t.Fatalf("parsing: %+v", err)
	}

	expected := models.AzureApiDefinition{
		ServiceName: "Example",
		ApiVersion:  "2020-01-01",
		Resources: map[string]models.AzureApiResource{
			"Resource": {
				Models: map[string]models.ModelDetails{
					"Model": {
						Fields: map[string]models.FieldDetails{
							"Identity": {
								JsonName:        "identity",
								CustomFieldType: pointer.To(models.CustomFieldTypeLegacySystemAndUserAssignedIdentityMap),
								Required:        false,
							},
							"Name": {
								JsonName: "name",
								ObjectDefinition: &models.ObjectDefinition{
									Type: models.ObjectDefinitionString,
								},
								Required: false,
							},
						},
					},
				},
				Operations: map[string]models.OperationDetails{
					"Test": {
						ContentType:         "application/json",
						ExpectedStatusCodes: []int{200},
						Method:              "GET",
						OperationId:         "Resource_Test",
						ResponseObject: &models.ObjectDefinition{
							ReferenceName: pointer.To("Model"),
							Type:          models.ObjectDefinitionReference,
						},
						UriSuffix: pointer.To("/example"),
					},
				},
			},
		},
	}
	validateParsedSwaggerResultMatches(t, expected, actual)
}

func TestParseModel_CommonSchema_IdentityLegacySystemAndUserAssignedMap_GenericDictionary(t *testing.T) {
	actual, err := ParseSwaggerFileForTesting(t, "model_commonschema_identitylegacysystemanduserassignedmap_genericdictionary.json")
	if err != nil {
		t.Fatalf("parsing: %+v", err)
	}

	expected := models.AzureApiDefinition{
		ServiceName: "Example",
		ApiVersion:  "2020-01-01",
		Resources: map[string]models.AzureApiResource{
			"Resource": {
				Models: map[string]models.ModelDetails{
					"Model": {
						Fields: map[string]models.FieldDetails{
							"Identity": {
								JsonName:        "identity",
								CustomFieldType: pointer.To(models.CustomFieldTypeLegacySystemAndUserAssignedIdentityMap),
								Required:        false,
							},
							"Name": {
								JsonName: "name",
								ObjectDefinition: &models.ObjectDefinition{
									Type: models.ObjectDefinitionString,
								},
								Required: false,
							},
						},
					},
				},
				Operations: map[string]models.OperationDetails{
					"Test": {
						ContentType:         "application/json",
						ExpectedStatusCodes: []int{200},
						Method:              "GET",
						OperationId:         "Resource_Test",
						ResponseObject: &models.ObjectDefinition{
							ReferenceName: pointer.To("Model"),
							Type:          models.ObjectDefinitionReference,
						},
						UriSuffix: pointer.To("/example"),
					},
				},
			},
		},
	}
	validateParsedSwaggerResultMatches(t, expected, actual)
}

func TestParseModel_CommonSchema_IdentitySystemAssigned(t *testing.T) {
	actual, err := ParseSwaggerFileForTesting(t, "model_commonschema_identitysystemassigned.json")
	if err != nil {
		t.Fatalf("parsing: %+v", err)
	}

	expected := models.AzureApiDefinition{
		ServiceName: "Example",
		ApiVersion:  "2020-01-01",
		Resources: map[string]models.AzureApiResource{
			"Resource": {
				Models: map[string]models.ModelDetails{
					"Model": {
						Fields: map[string]models.FieldDetails{
							"Identity": {
								JsonName:        "identity",
								CustomFieldType: pointer.To(models.CustomFieldTypeSystemAssignedIdentity),
								Required:        false,
							},
							"Name": {
								JsonName: "name",
								ObjectDefinition: &models.ObjectDefinition{
									Type: models.ObjectDefinitionString,
								},
								Required: false,
							},
						},
					},
				},
				Operations: map[string]models.OperationDetails{
					"Test": {
						ContentType:         "application/json",
						ExpectedStatusCodes: []int{200},
						Method:              "GET",
						OperationId:         "Resource_Test",
						ResponseObject: &models.ObjectDefinition{
							ReferenceName: pointer.To("Model"),
							Type:          models.ObjectDefinitionReference,
						},
						UriSuffix: pointer.To("/example"),
					},
				},
			},
		},
	}
	validateParsedSwaggerResultMatches(t, expected, actual)
}

func TestParseModel_CommonSchema_IdentitySystemAndUserAssignedList(t *testing.T) {
	actual, err := ParseSwaggerFileForTesting(t, "model_commonschema_identitysystemanduserassignedlist.json")
	if err != nil {
		t.Fatalf("parsing: %+v", err)
	}

	expected := models.AzureApiDefinition{
		ServiceName: "Example",
		ApiVersion:  "2020-01-01",
		Resources: map[string]models.AzureApiResource{
			"Resource": {
				Models: map[string]models.ModelDetails{
					"Model": {
						Fields: map[string]models.FieldDetails{
							"Identity": {
								JsonName:        "identity",
								CustomFieldType: pointer.To(models.CustomFieldTypeSystemAndUserAssignedIdentityList),
								Required:        false,
							},
							"Name": {
								JsonName: "name",
								ObjectDefinition: &models.ObjectDefinition{
									Type: models.ObjectDefinitionString,
								},
								Required: false,
							},
						},
					},
				},
				Operations: map[string]models.OperationDetails{
					"Test": {
						ContentType:         "application/json",
						ExpectedStatusCodes: []int{200},
						Method:              "GET",
						OperationId:         "Resource_Test",
						ResponseObject: &models.ObjectDefinition{
							ReferenceName: pointer.To("Model"),
							Type:          models.ObjectDefinitionReference,
						},
						UriSuffix: pointer.To("/example"),
					},
				},
			},
		},
	}
	validateParsedSwaggerResultMatches(t, expected, actual)
}

func TestParseModel_CommonSchema_IdentitySystemAndUserAssignedMap(t *testing.T) {
	actual, err := ParseSwaggerFileForTesting(t, "model_commonschema_identitysystemanduserassignedmap.json")
	if err != nil {
		t.Fatalf("parsing: %+v", err)
	}

	expected := models.AzureApiDefinition{
		ServiceName: "Example",
		ApiVersion:  "2020-01-01",
		Resources: map[string]models.AzureApiResource{
			"Resource": {
				Models: map[string]models.ModelDetails{
					"Model": {
						Fields: map[string]models.FieldDetails{
							"Identity": {
								JsonName:        "identity",
								CustomFieldType: pointer.To(models.CustomFieldTypeSystemAndUserAssignedIdentityMap),
								Required:        false,
							},
							"Name": {
								JsonName: "name",
								ObjectDefinition: &models.ObjectDefinition{
									Type: models.ObjectDefinitionString,
								},
								Required: false,
							},
						},
					},
				},
				Operations: map[string]models.OperationDetails{
					"Test": {
						ContentType:         "application/json",
						ExpectedStatusCodes: []int{200},
						Method:              "GET",
						OperationId:         "Resource_Test",
						ResponseObject: &models.ObjectDefinition{
							ReferenceName: pointer.To("Model"),
							Type:          models.ObjectDefinitionReference,
						},
						UriSuffix: pointer.To("/example"),
					},
				},
			},
		},
	}
	validateParsedSwaggerResultMatches(t, expected, actual)
}

func TestParseModel_CommonSchema_IdentitySystemAndUserAssignedMap_ExtraFields(t *testing.T) {
	// this handles the scenario of a System Assigned & User Assigned model having extra fields
	// in the user assigned identity block (#1066) - for example an extra `appId`
	actual, err := ParseSwaggerFileForTesting(t, "model_commonschema_identitysystemanduserassignedmap.json")
	if err != nil {
		t.Fatalf("parsing: %+v", err)
	}

	expected := models.AzureApiDefinition{
		ServiceName: "Example",
		ApiVersion:  "2020-01-01",
		Resources: map[string]models.AzureApiResource{
			"Resource": {
				Models: map[string]models.ModelDetails{
					"Model": {
						Fields: map[string]models.FieldDetails{
							"Identity": {
								JsonName:        "identity",
								CustomFieldType: pointer.To(models.CustomFieldTypeSystemAndUserAssignedIdentityMap),
								Required:        false,
							},
							"Name": {
								JsonName: "name",
								ObjectDefinition: &models.ObjectDefinition{
									Type: models.ObjectDefinitionString,
								},
								Required: false,
							},
						},
					},
				},
				Operations: map[string]models.OperationDetails{
					"Test": {
						ContentType:         "application/json",
						ExpectedStatusCodes: []int{200},
						Method:              "GET",
						OperationId:         "Resource_Test",
						ResponseObject: &models.ObjectDefinition{
							ReferenceName: pointer.To("Model"),
							Type:          models.ObjectDefinitionReference,
						},
						UriSuffix: pointer.To("/example"),
					},
				},
			},
		},
	}
	validateParsedSwaggerResultMatches(t, expected, actual)
}

func TestParseModel_CommonSchema_Location(t *testing.T) {
	actual, err := ParseSwaggerFileForTesting(t, "model_commonschema_location.json")
	if err != nil {
		t.Fatalf("parsing: %+v", err)
	}

	expected := models.AzureApiDefinition{
		ServiceName: "Example",
		ApiVersion:  "2020-01-01",
		Resources: map[string]models.AzureApiResource{
			"Resource": {
				Models: map[string]models.ModelDetails{
					"Model": {
						Fields: map[string]models.FieldDetails{
							"Location": {
								JsonName:        "location",
								CustomFieldType: pointer.To(models.CustomFieldTypeLocation),
								Required:        false,
							},
							"Name": {
								JsonName: "name",
								ObjectDefinition: &models.ObjectDefinition{
									Type: models.ObjectDefinitionString,
								},
								Required: false,
							},
						},
					},
				},
				Operations: map[string]models.OperationDetails{
					"Test": {
						ContentType:         "application/json",
						ExpectedStatusCodes: []int{200},
						Method:              "GET",
						OperationId:         "Resource_Test",
						ResponseObject: &models.ObjectDefinition{
							ReferenceName: pointer.To("Model"),
							Type:          models.ObjectDefinitionReference,
						},
						UriSuffix: pointer.To("/example"),
					},
				},
			},
		},
	}
	validateParsedSwaggerResultMatches(t, expected, actual)
}
