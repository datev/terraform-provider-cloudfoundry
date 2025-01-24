package provider

import (
	"bytes"
	"html/template"
	"testing"

	cfv3resource "github.com/cloudfoundry/go-cfclient/v3/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"
)

type ServiceInstanceSharingResourceModelPtr struct {
	HclType           string
	HclObjectName     string
	ServiceInstanceId *string
	SpaceId           *string
}

func hclResourceServiceInstanceSharing(model *ServiceInstanceSharingResourceModelPtr) string {
	if model != nil {
		s := `
		{{.HclType}} "cloudfoundry_service_instance_sharing" {{.HclObjectName}} {
		{{- if .ServiceInstanceId}}
			service_instance_id = "{{.ServiceInstanceId}}"
		{{- end -}}
		{{ if .SpaceId}}
			space_id = "{{.SpaceId}}"
		{{- end }}
		}
		`
		tmpl, err := template.New("resource_service_instance_sharing").Parse(s)
		if err != nil {
			panic(err)
		}
		buf := new(bytes.Buffer)
		err = tmpl.Execute(buf, model)
		if err != nil {
			panic(err)
		}
		return buf.String()
	}
	return model.HclType + ` "cloudfoundry_service_instance_sharing" ` + model.HclObjectName + ` {}`
}

func TestServiceInstanceSharingResource_Configure(t *testing.T) {
	t.Run("happy path - create service instance sharing", func(t *testing.T) {
		// setup
		resourceName := "cloudfoundry_service_instance_sharing.rs"
		cfg := getCFHomeConf()
		rec := cfg.SetupVCR(t, "fixtures/resource_service_instance_sharing")
		defer stopQuietly(rec)

		testUserProvidedServiceInstanceGUID := "5e2976bb-332e-41e1-8be3-53baafea9296"
		testSpaceGUID := "02c0cc92-6ecc-44b1-b7b2-096ca19ee143"

		// actual test
		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(rec.GetDefaultClient()),
			Steps: []resource.TestStep{
				{
					Config: hclProvider(nil) + hclResourceServiceInstanceSharing(&ServiceInstanceSharingResourceModelPtr{
						HclType:           hclObjectResource,
						HclObjectName:     "rs",
						ServiceInstanceId: strtostrptr(testUserProvidedServiceInstanceGUID),
						SpaceId:           strtostrptr(testSpaceGUID),
					}),
					Check: resource.ComposeAggregateTestCheckFunc(
						// resource.TestMatchResourceAttr(resourceName, "id", regexp.MustCompile(`asdf`)),
						resource.TestMatchResourceAttr(resourceName, "service_instance_id", regexpValidUUID),
						resource.TestMatchResourceAttr(resourceName, "space_id", regexpValidUUID),
					),
					ExpectNonEmptyPlan: true,
				},
			},
		})
	})
}

func TestMapRelationShipToType(t *testing.T) {
	relationship := &cfv3resource.ServiceInstanceSharedSpaceRelationships{
		Data: []cfv3resource.Relationship{
			{GUID: "space-guid-1"},
		},
	}

	serviceInstanceId := "service-instance-guid-1"
	expected := ServiceInstanceSharingType{
		Id:                types.StringValue("service-instance-guid-1/space-guid-1"),
		ServiceInstanceId: types.StringValue(serviceInstanceId),
		SpaceId:           types.StringValue("space-guid-1"),
	}

	result := mapRelationShipToType(relationship, serviceInstanceId)

	assert.Equal(t, expected, result)
}
