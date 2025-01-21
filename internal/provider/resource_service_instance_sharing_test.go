package provider

import (
	"testing"

	cfv3resource "github.com/cloudfoundry/go-cfclient/v3/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestMapRelationShipToType(t *testing.T) {
	relationship := &cfv3resource.ServiceInstanceSharedSpaceRelationships{
		Data: []cfv3resource.Relationship{
			{GUID: "space-guid-1"},
		},
	}

	serviceInstanceId := "service-instance-guid-1"
	expected := ServiceInstanceSharingType{
		Id:                types.StringValue("service-instance-guid-1/service-instance-guid-1"),
		ServiceInstanceId: types.StringValue(serviceInstanceId),
		SpaceId:           types.StringValue("space-guid-1"),
	}

	result := mapRelationShipToType(relationship, serviceInstanceId)

	assert.Equal(t, expected, result)
}
