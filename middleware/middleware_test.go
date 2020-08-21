package middleware

import (
	"testing"

	"github.com/golib/assert"
)

type RR string

type regionStruct struct {
	RegionId string
	RegionID *string
	Region   RR
	Name     string
}

type testParams struct {
	Region_id string

	regionStruct
	List      []*regionStruct
	Struct    regionStruct
	StructPtr *regionStruct
}

func Test_RegionTransformHandler_ModifyRequestParams(t *testing.T) {


	tmp := "h#cn-north-2"
	params := &testParams{
		Region_id: tmp,

		regionStruct: regionStruct{
			RegionId: "h#cn-north-1",
			RegionID: &tmp,
			Region:   "h#cn-north-3",
			Name:     "h#cn-north-4",
		},

		List: []*regionStruct{
			{
				RegionId: "h#cn-north-1",
				RegionID: &tmp,
				Region:   "h#cn-north-3",
				Name:     "h#cn-north-4",
			},
			{
				RegionId: "h#cn-north-1",
			},
		},
		Struct: regionStruct{
			RegionId: "h#cn-north-1",
		},
		StructPtr: &regionStruct{
			RegionId: "h#cn-north-1",
		},
	}

	ModifyRequestParams(params)
	assert.Equal(t, "cn-north-2", params.Region_id)
	assert.Equal(t, "cn-north-1", params.RegionId)
	assert.Equal(t, "cn-north-2", *params.RegionID)
	assert.Equal(t, "cn-north-3", string(params.Region))
	assert.Equal(t, "h#cn-north-4", params.Name)
	assert.Equal(t, "cn-north-1", params.List[0].RegionId)
	assert.Equal(t, "cn-north-2", *params.List[0].RegionID)
	assert.Equal(t, "cn-north-1", params.Struct.RegionId)
	assert.Equal(t, "cn-north-1", params.StructPtr.RegionId)

}

func Test_RegionTransformHandler_ModifyResponseData(t *testing.T) {


	tmp := "cn-north-2"
	params := &testParams{
		Region_id: tmp,

		regionStruct: regionStruct{
			RegionId: "cn-north-1",
			RegionID: &tmp,
			Region:   "cn-north-3",
			Name:     "cn-north-4",
		},

		List: []*regionStruct{
			{
				RegionId: "cn-north-1",
				RegionID: &tmp,
				Region:   "cn-north-3",
				Name:     "cn-north-4",
			},
			{
				RegionId: "cn-north-1",
			},
		},
		Struct: regionStruct{
			RegionId: "cn-north-1",
		},
		StructPtr: &regionStruct{
			RegionId: "cn-north-1",
		},
	}

	ModifyResponseData(params)
	assert.Equal(t, "h#cn-north-2", params.Region_id)
	assert.Equal(t, "h#cn-north-1", params.RegionId)
	assert.Equal(t, "h#cn-north-2", *params.RegionID)
	assert.Equal(t, "h#cn-north-3", string(params.Region))
	assert.Equal(t, "cn-north-4", params.Name)
	assert.Equal(t, "h#cn-north-1", params.List[0].RegionId)
	assert.Equal(t, "h#cn-north-2", *params.List[0].RegionID)
	assert.Equal(t, "h#cn-north-1", params.Struct.RegionId)
	assert.Equal(t, "h#cn-north-1", params.StructPtr.RegionId)

}