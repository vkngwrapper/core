package universal_test

import (
	"bytes"
	"encoding/binary"
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanLoader1_0_CreateGraphicsPipelines_EmptySuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()
	basePipeline := mocks.EasyDummyPipeline(t, device, loader)

	mockDriver.EXPECT().VkCreateGraphicsPipelines(mocks.Exactly(device.Handle()), nil, driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pGraphicsPipelines *driver.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]driver.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			require.Equal(t, uint64(28), createInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO
			require.True(t, createInfo.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000800), createInfo.FieldByName("flags").Uint())
			require.Equal(t, uint64(0), createInfo.FieldByName("stageCount").Uint())
			require.True(t, createInfo.FieldByName("pStages").IsNil())
			require.True(t, createInfo.FieldByName("pVertexInputState").IsNil())
			require.True(t, createInfo.FieldByName("pInputAssemblyState").IsNil())
			require.True(t, createInfo.FieldByName("pTessellationState").IsNil())
			require.True(t, createInfo.FieldByName("pViewportState").IsNil())
			require.True(t, createInfo.FieldByName("pRasterizationState").IsNil())
			require.True(t, createInfo.FieldByName("pMultisampleState").IsNil())
			require.True(t, createInfo.FieldByName("pDepthStencilState").IsNil())
			require.True(t, createInfo.FieldByName("pColorBlendState").IsNil())
			require.True(t, createInfo.FieldByName("pDynamicState").IsNil())
			actualLayout := (driver.VkPipelineLayout)(unsafe.Pointer(createInfo.FieldByName("layout").Elem().UnsafeAddr()))
			require.Same(t, layout.Handle(), actualLayout)
			actualRenderPass := (driver.VkRenderPass)(unsafe.Pointer(createInfo.FieldByName("renderPass").Elem().UnsafeAddr()))
			require.Same(t, renderPass.Handle(), actualRenderPass)
			require.Equal(t, uint64(1), createInfo.FieldByName("subpass").Uint())
			actualBasePipeline := (driver.VkPipeline)(unsafe.Pointer(createInfo.FieldByName("basePipelineHandle").Elem().UnsafeAddr()))
			require.Same(t, basePipeline.Handle(), actualBasePipeline)
			require.Equal(t, int64(3), createInfo.FieldByName("basePipelineIndex").Int())

			return common.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []*core.GraphicsPipelineOptions{
		{
			Flags:             core.PipelineLibraryKHR,
			Layout:            layout,
			RenderPass:        renderPass,
			SubPass:           1,
			BasePipeline:      basePipeline,
			BasePipelineIndex: 3,
		},
	})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_ShaderStagesSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()
	shaderModule1 := mocks.EasyMockShaderModule(ctrl)
	shaderModule2 := mocks.EasyMockShaderModule(ctrl)

	mockDriver.EXPECT().VkCreateGraphicsPipelines(mocks.Exactly(device.Handle()), nil, driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pGraphicsPipelines *driver.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]driver.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			require.Equal(t, uint64(28), createInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO
			require.True(t, createInfo.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), createInfo.FieldByName("stageCount").Uint())

			shaderPtr := (*driver.VkPipelineShaderStageCreateInfo)(unsafe.Pointer(createInfo.FieldByName("pStages").Elem().UnsafeAddr()))
			shaderSlice := reflect.ValueOf(([]driver.VkPipelineShaderStageCreateInfo)(unsafe.Slice(shaderPtr, 2)))

			shader := shaderSlice.Index(0)
			require.Equal(t, uint64(18), shader.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
			require.True(t, shader.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), shader.FieldByName("flags").Uint()) // VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT
			require.Equal(t, uint64(8), shader.FieldByName("stage").Uint()) // VK_SHADER_STAGE_GEOMETRY_BIT
			module := (driver.VkShaderModule)(unsafe.Pointer(shader.FieldByName("module").Elem().UnsafeAddr()))
			require.Same(t, shaderModule1.Handle(), module)
			namePtr := (*driver.Char)(unsafe.Pointer(shader.FieldByName("pName").Elem().UnsafeAddr()))
			nameSlice := ([]driver.Char)(unsafe.Slice(namePtr, 256))
			expectedName := "some shader 1"
			for i, r := range expectedName {
				require.Equal(t, r, rune(nameSlice[i]))
			}
			require.Equal(t, 0, int(nameSlice[len(expectedName)]))
			require.True(t, shader.FieldByName("pSpecializationInfo").IsNil())

			shader = shaderSlice.Index(1)
			require.Equal(t, uint64(18), shader.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
			require.True(t, shader.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), shader.FieldByName("flags").Uint())
			require.Equal(t, uint64(16), shader.FieldByName("stage").Uint())
			module = (driver.VkShaderModule)(unsafe.Pointer(shader.FieldByName("module").Elem().UnsafeAddr()))
			require.Same(t, shaderModule2.Handle(), module)
			namePtr = (*driver.Char)(unsafe.Pointer(shader.FieldByName("pName").Elem().UnsafeAddr()))
			nameSlice = ([]driver.Char)(unsafe.Slice(namePtr, 256))
			expectedName = "another shader 2"
			for i, r := range expectedName {
				require.Equal(t, r, rune(nameSlice[i]))
			}
			require.Equal(t, 0, int(nameSlice[len(expectedName)]))

			specInfo := shader.FieldByName("pSpecializationInfo").Elem()
			require.Equal(t, uint64(2), specInfo.FieldByName("mapEntryCount").Uint())
			mapEntryPtr := (*driver.VkSpecializationMapEntry)(unsafe.Pointer(specInfo.FieldByName("pMapEntries").Elem().UnsafeAddr()))
			mapEntrySlice := reflect.ValueOf(([]driver.VkSpecializationMapEntry)(unsafe.Slice(mapEntryPtr, 2)))

			firstEntryID := mapEntrySlice.Index(0).FieldByName("constantID").Uint()
			secondEntryID := mapEntrySlice.Index(1).FieldByName("constantID").Uint()

			// We deliver the map as a go map so it could come out in any order
			require.ElementsMatch(t, []uint64{1, 2}, []uint64{firstEntryID, secondEntryID})

			firstEntryOffset := mapEntrySlice.Index(0).FieldByName("offset").Uint()
			firstEntrySize := mapEntrySlice.Index(0).FieldByName("size").Uint()

			secondEntryOffset := mapEntrySlice.Index(1).FieldByName("offset").Uint()
			secondEntrySize := mapEntrySlice.Index(1).FieldByName("size").Uint()

			require.Equal(t, uint64(0), firstEntryOffset)
			require.Equal(t, firstEntrySize, secondEntryOffset)

			require.Equal(t, firstEntrySize+secondEntrySize, specInfo.FieldByName("dataSize").Uint())
			dataPtr := (*byte)(unsafe.Pointer(specInfo.FieldByName("pData").Pointer()))
			dataSlice := ([]byte)(unsafe.Slice(dataPtr, firstEntrySize+secondEntrySize))

			firstEntryBytes := bytes.NewBuffer(dataSlice[:firstEntrySize])
			secondEntryBytes := bytes.NewBuffer(dataSlice[firstEntrySize : firstEntrySize+secondEntrySize])

			firstValueBytes := firstEntryBytes
			secondValueBytes := secondEntryBytes

			if firstEntryID == 2 {
				firstValueBytes = secondEntryBytes
				secondValueBytes = firstEntryBytes
			}

			var boolVal uint32
			err = binary.Read(firstValueBytes, binary.LittleEndian, &boolVal)
			require.NoError(t, err)
			require.Equal(t, uint32(1), boolVal)

			var floatVal float64
			err = binary.Read(secondValueBytes, binary.LittleEndian, &floatVal)
			require.NoError(t, err)
			require.Equal(t, float64(7.6), floatVal)

			return common.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []*core.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			ShaderStages: []*core.ShaderStage{
				{
					Flags:  core.ShaderStageAllowVaryingSubgroupSizeEXT,
					Name:   "some shader 1",
					Stage:  common.StageGeometry,
					Shader: shaderModule1,
				},
				{
					Name:   "another shader 2",
					Stage:  common.StageFragment,
					Shader: shaderModule2,
					SpecializationInfo: map[uint32]interface{}{
						1: true,
						2: float64(7.6),
					},
				},
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Same(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_ShaderStagesFailure_InvalidSpecializationValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	shaderModule1 := mocks.EasyMockShaderModule(ctrl)
	shaderModule2 := mocks.EasyMockShaderModule(ctrl)

	_, _, err = loader.CreateGraphicsPipelines(device, nil, nil, []*core.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			ShaderStages: []*core.ShaderStage{
				{
					Flags:  core.ShaderStageAllowVaryingSubgroupSizeEXT,
					Name:   "some shader 1",
					Stage:  common.StageGeometry,
					Shader: shaderModule1,
				},
				{
					Name:   "another shader 2",
					Stage:  common.StageFragment,
					Shader: shaderModule2,
					SpecializationInfo: map[uint32]interface{}{
						1: "wow, this is invalid",
						2: float64(7.6),
					},
				},
			},
		}})
	require.EqualError(t, err, "failed to populate shader stage with specialization values: 1 -> wow, this is invalid: binary.Write: invalid type string")
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_VertexInputSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(mocks.Exactly(device.Handle()), nil, driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pGraphicsPipelines *driver.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]driver.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			require.Equal(t, uint64(28), createInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO
			require.True(t, createInfo.FieldByName("pNext").IsNil())

			state := createInfo.FieldByName("pVertexInputState").Elem()
			require.Equal(t, uint64(19), state.FieldByName("sType").Uint())
			require.True(t, state.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), state.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), state.FieldByName("vertexBindingDescriptionCount").Uint())
			require.Equal(t, uint64(2), state.FieldByName("vertexAttributeDescriptionCount").Uint())

			bindingDescPtr := (*driver.VkVertexInputBindingDescription)(unsafe.Pointer(state.FieldByName("pVertexBindingDescriptions").Elem().UnsafeAddr()))
			bindingDescSlice := reflect.ValueOf(([]driver.VkVertexInputBindingDescription)(unsafe.Slice(bindingDescPtr, 1)))

			binding := bindingDescSlice.Index(0)
			require.Equal(t, uint64(17), binding.FieldByName("binding").Uint())
			require.Equal(t, uint64(19), binding.FieldByName("stride").Uint())
			require.Equal(t, uint64(1), binding.FieldByName("inputRate").Uint())

			attDescPtr := (*driver.VkVertexInputAttributeDescription)(unsafe.Pointer(state.FieldByName("pVertexAttributeDescriptions").Elem().UnsafeAddr()))
			attDescSlice := reflect.ValueOf(([]driver.VkVertexInputAttributeDescription)(unsafe.Slice(attDescPtr, 2)))

			att := attDescSlice.Index(0)
			require.Equal(t, uint64(1), att.FieldByName("location").Uint())
			require.Equal(t, uint64(3), att.FieldByName("binding").Uint())
			require.Equal(t, uint64(5), att.FieldByName("offset").Uint())
			require.Equal(t, uint64(8), att.FieldByName("format").Uint())

			att = attDescSlice.Index(1)
			require.Equal(t, uint64(7), att.FieldByName("location").Uint())
			require.Equal(t, uint64(11), att.FieldByName("binding").Uint())
			require.Equal(t, uint64(13), att.FieldByName("offset").Uint())
			require.Equal(t, uint64(64), att.FieldByName("format").Uint())

			return common.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []*core.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			VertexInput: &core.VertexInputOptions{
				VertexAttributeDescriptions: []core.VertexAttributeDescription{
					{
						Location: 1,
						Binding:  3,
						Format:   common.FormatA1R5G5B5UnsignedNormalized,
						Offset:   5,
					},
					{
						Location: 7,
						Binding:  11,
						Format:   common.FormatA2B10G10R10UnsignedNormalized,
						Offset:   13,
					},
				},
				VertexBindingDescriptions: []core.VertexBindingDescription{
					{
						InputRate: core.RateInstance,
						Binding:   17,
						Stride:    19,
					},
				},
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Same(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_InputAssemblySuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(mocks.Exactly(device.Handle()), nil, driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pGraphicsPipelines *driver.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]driver.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			state := createInfo.FieldByName("pInputAssemblyState").Elem()
			require.Equal(t, uint64(20), state.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO
			require.True(t, state.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), state.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), state.FieldByName("topology").Uint()) // VK_PRIMITIVE_TOPOLOGY_LINE_LIST
			require.Equal(t, uint64(1), state.FieldByName("primitiveRestartEnable").Uint())

			return common.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []*core.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			InputAssembly: &core.InputAssemblyOptions{
				Topology:               common.TopologyLineList,
				EnablePrimitiveRestart: true,
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Same(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_TessellationSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(mocks.Exactly(device.Handle()), nil, driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pGraphicsPipelines *driver.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]driver.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			state := createInfo.FieldByName("pTessellationState").Elem()

			require.Equal(t, uint64(21), state.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO
			require.True(t, state.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), state.FieldByName("flags").Uint())
			require.Equal(t, uint64(3), state.FieldByName("patchControlPoints").Uint())

			return common.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []*core.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			Tessellation: &core.TessellationOptions{
				PatchControlPoints: 3,
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Same(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_ViewportSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(mocks.Exactly(device.Handle()), nil, driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pGraphicsPipelines *driver.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]driver.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			state := createInfo.FieldByName("pViewportState").Elem()
			require.Equal(t, uint64(22), state.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO
			require.True(t, state.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), state.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), state.FieldByName("viewportCount").Uint())
			require.Equal(t, uint64(2), state.FieldByName("scissorCount").Uint())

			viewportPtr := (*driver.VkViewport)(unsafe.Pointer(state.FieldByName("pViewports").Elem().UnsafeAddr()))
			viewportSlice := reflect.ValueOf(([]driver.VkViewport)(unsafe.Slice(viewportPtr, 1)))

			viewport := viewportSlice.Index(0)
			require.Equal(t, float64(1), viewport.FieldByName("x").Float())
			require.Equal(t, float64(3), viewport.FieldByName("y").Float())
			require.Equal(t, float64(5), viewport.FieldByName("width").Float())
			require.Equal(t, float64(7), viewport.FieldByName("height").Float())
			require.Equal(t, float64(11), viewport.FieldByName("minDepth").Float())
			require.Equal(t, float64(13), viewport.FieldByName("maxDepth").Float())

			scissorPtr := (*driver.VkRect2D)(unsafe.Pointer(state.FieldByName("pScissors").Elem().UnsafeAddr()))
			scissorSlice := reflect.ValueOf(([]driver.VkRect2D)(unsafe.Slice(scissorPtr, 2)))

			scissor := scissorSlice.Index(0)
			scissorOffset := scissor.FieldByName("offset")
			require.Equal(t, int64(17), scissorOffset.FieldByName("x").Int())
			require.Equal(t, int64(19), scissorOffset.FieldByName("y").Int())
			scissorExtent := scissor.FieldByName("extent")
			require.Equal(t, uint64(23), scissorExtent.FieldByName("width").Uint())
			require.Equal(t, uint64(29), scissorExtent.FieldByName("height").Uint())

			scissor = scissorSlice.Index(1)
			scissorOffset = scissor.FieldByName("offset")
			require.Equal(t, int64(31), scissorOffset.FieldByName("x").Int())
			require.Equal(t, int64(37), scissorOffset.FieldByName("y").Int())
			scissorExtent = scissor.FieldByName("extent")
			require.Equal(t, uint64(41), scissorExtent.FieldByName("width").Uint())
			require.Equal(t, uint64(43), scissorExtent.FieldByName("height").Uint())

			return common.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []*core.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			Viewport: &core.ViewportOptions{
				Viewports: []common.Viewport{
					{
						X:        1,
						Y:        3,
						Width:    5,
						Height:   7,
						MinDepth: 11,
						MaxDepth: 13,
					},
				},
				Scissors: []common.Rect2D{
					{
						Offset: common.Offset2D{X: 17, Y: 19},
						Extent: common.Extent2D{Width: 23, Height: 29},
					},
					{
						Offset: common.Offset2D{X: 31, Y: 37},
						Extent: common.Extent2D{Width: 41, Height: 43},
					},
				},
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Same(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_RasterizationSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(mocks.Exactly(device.Handle()), nil, driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pGraphicsPipelines *driver.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]driver.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			state := createInfo.FieldByName("pRasterizationState").Elem()
			require.Equal(t, uint64(23), state.FieldByName("sType").Uint())
			require.True(t, state.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), state.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), state.FieldByName("depthClampEnable").Uint())
			require.Equal(t, uint64(1), state.FieldByName("rasterizerDiscardEnable").Uint())
			require.Equal(t, uint64(1), state.FieldByName("polygonMode").Uint()) // VK_POLYGON_MODE_LINE
			require.Equal(t, uint64(3), state.FieldByName("cullMode").Uint())    // VK_CULL_MODE_FRONT_AND_BACK
			require.Equal(t, uint64(1), state.FieldByName("frontFace").Uint())   // VK_FRONT_FACE_CLOCKWISE
			require.Equal(t, uint64(1), state.FieldByName("depthBiasEnable").Uint())
			require.InDelta(t, float64(2.3), state.FieldByName("depthBiasClamp").Float(), 0.0001)
			require.InDelta(t, float64(3.4), state.FieldByName("depthBiasConstantFactor").Float(), 0.0001)
			require.InDelta(t, float64(4.5), state.FieldByName("depthBiasSlopeFactor").Float(), 0.0001)
			require.InDelta(t, float64(5.6), state.FieldByName("lineWidth").Float(), 0.0001)

			return common.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []*core.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			Rasterization: &core.RasterizationOptions{
				DepthClamp:        true,
				RasterizerDiscard: true,

				PolygonMode: core.PolygonModeLine,
				CullMode:    common.CullFrontAndBack,
				FrontFace:   common.FrontFaceClockwise,

				DepthBias:               true,
				DepthBiasClamp:          2.3,
				DepthBiasConstantFactor: 3.4,
				DepthBiasSlopeFactor:    4.5,

				LineWidth: 5.6,
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Same(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_MultisampleSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(mocks.Exactly(device.Handle()), nil, driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pGraphicsPipelines *driver.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]driver.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			state := createInfo.FieldByName("pMultisampleState").Elem()
			require.Equal(t, uint64(24), state.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO
			require.True(t, state.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), state.FieldByName("flags").Uint())
			require.Equal(t, uint64(64), state.FieldByName("rasterizationSamples").Uint())
			require.Equal(t, uint64(1), state.FieldByName("sampleShadingEnable").Uint())
			require.InDelta(t, float64(2.3), state.FieldByName("minSampleShading").Float(), 0.0001)
			require.Equal(t, uint64(1), state.FieldByName("alphaToCoverageEnable").Uint())
			require.Equal(t, uint64(1), state.FieldByName("alphaToOneEnable").Uint())

			sampleMaskPtr := (*driver.VkSampleMask)(unsafe.Pointer(state.FieldByName("pSampleMask").Elem().UnsafeAddr()))
			sampleMaskSlice := reflect.ValueOf(([]driver.VkSampleMask)(unsafe.Slice(sampleMaskPtr, 2)))
			require.Equal(t, uint64(1), sampleMaskSlice.Index(0).Uint())
			require.Equal(t, uint64(3), sampleMaskSlice.Index(1).Uint())

			return common.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []*core.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			Multisample: &core.MultisampleOptions{
				RasterizationSamples: common.Samples64,
				SampleShading:        true,
				MinSampleShading:     2.3,
				SampleMask:           []uint32{1, 3},
				AlphaToCoverage:      true,
				AlphaToOne:           true,
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Same(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_MultisampleSuccess_NoSampleMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(mocks.Exactly(device.Handle()), nil, driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pGraphicsPipelines *driver.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]driver.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			state := createInfo.FieldByName("pMultisampleState").Elem()
			require.Equal(t, uint64(24), state.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO
			require.True(t, state.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), state.FieldByName("flags").Uint())
			require.Equal(t, uint64(64), state.FieldByName("rasterizationSamples").Uint())
			require.Equal(t, uint64(1), state.FieldByName("sampleShadingEnable").Uint())
			require.InDelta(t, float64(2.3), state.FieldByName("minSampleShading").Float(), 0.0001)
			require.Equal(t, uint64(1), state.FieldByName("alphaToCoverageEnable").Uint())
			require.Equal(t, uint64(1), state.FieldByName("alphaToOneEnable").Uint())
			require.True(t, state.FieldByName("pSampleMask").IsNil())

			return common.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []*core.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			Multisample: &core.MultisampleOptions{
				RasterizationSamples: common.Samples64,
				SampleShading:        true,
				MinSampleShading:     2.3,
				AlphaToCoverage:      true,
				AlphaToOne:           true,
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Same(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_MultisampleFail_MismatchSampleMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)

	_, _, err = loader.CreateGraphicsPipelines(device, nil, nil, []*core.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			Multisample: &core.MultisampleOptions{
				RasterizationSamples: common.Samples4,
				SampleShading:        true,
				MinSampleShading:     2.3,
				SampleMask:           []uint32{1, 3},
				AlphaToCoverage:      true,
				AlphaToOne:           true,
			},
		}})
	require.EqualError(t, err, "expected a sample mask size of 1, because 4 rasterization samples were specified- however, received a sample mask size of 2")
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_DepthStencilSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(mocks.Exactly(device.Handle()), nil, driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pGraphicsPipelines *driver.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]driver.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			state := createInfo.FieldByName("pDepthStencilState").Elem()
			require.Equal(t, uint64(25), state.FieldByName("sType").Uint())
			require.True(t, state.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), state.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), state.FieldByName("depthTestEnable").Uint())
			require.Equal(t, uint64(1), state.FieldByName("depthWriteEnable").Uint())
			require.Equal(t, uint64(2), state.FieldByName("depthCompareOp").Uint()) // VK_COMPARE_OP_EQUAL
			require.Equal(t, uint64(1), state.FieldByName("depthBoundsTestEnable").Uint())
			require.Equal(t, uint64(1), state.FieldByName("stencilTestEnable").Uint())
			require.InDelta(t, 2.3, state.FieldByName("minDepthBounds").Float(), 0.0001)
			require.InDelta(t, 3.4, state.FieldByName("maxDepthBounds").Float(), 0.0001)

			frontState := state.FieldByName("front")
			require.Equal(t, uint64(5), frontState.FieldByName("failOp").Uint())      // VK_STENCIL_OP_INVERT
			require.Equal(t, uint64(7), frontState.FieldByName("passOp").Uint())      // VK_STENCIL_OP_DECREMENT_AND_WRAP
			require.Equal(t, uint64(2), frontState.FieldByName("depthFailOp").Uint()) // VK_STENCIL_OP_REPLACE
			require.Equal(t, uint64(6), frontState.FieldByName("compareOp").Uint())   // VK_COMPARE_OP_GREATER_OR_EQUAL
			require.Equal(t, uint64(3), frontState.FieldByName("compareMask").Uint())
			require.Equal(t, uint64(5), frontState.FieldByName("writeMask").Uint())
			require.Equal(t, uint64(7), frontState.FieldByName("reference").Uint())

			backState := state.FieldByName("back")
			require.Equal(t, uint64(3), backState.FieldByName("failOp").Uint())      // VK_STENCIL_OP_INCREMENT_AND_CLAMP
			require.Equal(t, uint64(1), backState.FieldByName("passOp").Uint())      // VK_STENCIL_OP_ZERO
			require.Equal(t, uint64(0), backState.FieldByName("depthFailOp").Uint()) // VK_STENCIL_OP_KEEP
			require.Equal(t, uint64(1), backState.FieldByName("compareOp").Uint())   // VK_COMPARE_OP_LESS
			require.Equal(t, uint64(11), backState.FieldByName("compareMask").Uint())
			require.Equal(t, uint64(13), backState.FieldByName("writeMask").Uint())
			require.Equal(t, uint64(17), backState.FieldByName("reference").Uint())

			return common.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []*core.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			DepthStencil: &core.DepthStencilOptions{
				DepthTestEnable:       true,
				DepthWriteEnable:      true,
				DepthCompareOp:        common.CompareEqual,
				DepthBoundsTestEnable: true,
				StencilTestEnable:     true,
				FrontStencilState: core.StencilOpState{
					FailOp:      common.StencilInvert,
					PassOp:      common.StencilDecrementAndWrap,
					DepthFailOp: common.StencilReplace,
					CompareOp:   common.CompareGreaterOrEqual,
					CompareMask: 3,
					WriteMask:   5,
					Reference:   7,
				},
				BackStencilState: core.StencilOpState{
					FailOp:      common.StencilIncrementAndClamp,
					PassOp:      common.StencilZero,
					DepthFailOp: common.StencilKeep,
					CompareOp:   common.CompareLess,
					CompareMask: 11,
					WriteMask:   13,
					Reference:   17,
				},
				MinDepthBounds: 2.3,
				MaxDepthBounds: 3.4,
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Same(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_ColorBlendSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(mocks.Exactly(device.Handle()), nil, driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pGraphicsPipelines *driver.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]driver.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			state := createInfo.FieldByName("pColorBlendState").Elem()
			require.Equal(t, uint64(26), state.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO
			require.True(t, state.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), state.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), state.FieldByName("logicOpEnable").Uint())
			require.Equal(t, uint64(12), state.FieldByName("logicOp").Uint()) // VK_LOGIC_OP_COPY_INVERTED
			require.Equal(t, uint64(2), state.FieldByName("attachmentCount").Uint())

			attachmentsPtr := (*driver.VkPipelineColorBlendAttachmentState)(unsafe.Pointer(state.FieldByName("pAttachments").Elem().UnsafeAddr()))
			attachmentsSlice := reflect.ValueOf(([]driver.VkPipelineColorBlendAttachmentState)(unsafe.Slice(attachmentsPtr, 2)))

			attachment := attachmentsSlice.Index(0)
			require.Equal(t, uint64(1), attachment.FieldByName("blendEnable").Uint())
			require.Equal(t, uint64(4), attachment.FieldByName("srcColorBlendFactor").Uint())   // VK_BLEND_FACTOR_DST_COLOR
			require.Equal(t, uint64(18), attachment.FieldByName("dstColorBlendFactor").Uint())  // VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA
			require.Equal(t, uint64(1000148015), attachment.FieldByName("colorBlendOp").Uint()) //VK_BLEND_OP_DARKEN_EXT
			require.Equal(t, uint64(17), attachment.FieldByName("srcAlphaBlendFactor").Uint())  // VK_BLEND_FACTOR_SRC1_ALPHA
			require.Equal(t, uint64(6), attachment.FieldByName("dstAlphaBlendFactor").Uint())   // VK_BLEND_FACTOR_SRC_ALPHA
			require.Equal(t, uint64(1000148021), attachment.FieldByName("alphaBlendOp").Uint()) // VK_BLEND_OP_DIFFERENCE_EXT
			require.Equal(t, uint64(8), attachment.FieldByName("colorWriteMask").Uint())        // VK_COLOR_COMPONENT_A_BIT

			attachment = attachmentsSlice.Index(1)
			require.Equal(t, uint64(0), attachment.FieldByName("blendEnable").Uint())
			require.Equal(t, uint64(7), attachment.FieldByName("srcColorBlendFactor").Uint())   // VK_BLEND_FACTOR_ONE_MINUS_SRC_ALPHA
			require.Equal(t, uint64(11), attachment.FieldByName("dstColorBlendFactor").Uint())  // VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR
			require.Equal(t, uint64(1000148002), attachment.FieldByName("colorBlendOp").Uint()) // VK_BLEND_OP_DST_EXT
			require.Equal(t, uint64(12), attachment.FieldByName("srcAlphaBlendFactor").Uint())  // VK_BLEND_FACTOR_CONSTANT_ALPHA
			require.Equal(t, uint64(1), attachment.FieldByName("dstAlphaBlendFactor").Uint())   // VK_BLEND_FACTOR_ONE
			require.Equal(t, uint64(1000148030), attachment.FieldByName("alphaBlendOp").Uint()) // VK_BLEND_OP_HARDMIX_EXT
			require.Equal(t, uint64(1), attachment.FieldByName("colorWriteMask").Uint())        // VK_COLOR_COMPONENT_R_BIT

			constants := state.FieldByName("blendConstants")
			require.InDelta(t, 1.2, constants.Index(0).Float(), 0.0001)
			require.InDelta(t, 2.3, constants.Index(1).Float(), 0.0001)
			require.InDelta(t, 3.4, constants.Index(2).Float(), 0.0001)
			require.InDelta(t, 4.5, constants.Index(3).Float(), 0.0001)

			return common.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []*core.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			ColorBlend: &core.ColorBlendOptions{
				LogicOpEnabled: true,
				LogicOp:        common.LogicOpCopyInverted,
				BlendConstants: [4]float32{1.2, 2.3, 3.4, 4.5},
				Attachments: []core.ColorBlendAttachment{
					{
						BlendEnabled: true,
						SrcColor:     common.BlendDstColor,
						DstColor:     common.BlendOneMinusSrc1Alpha,
						ColorBlendOp: common.BlendOpDarkenEXT,
						SrcAlpha:     common.BlendSrc1Alpha,
						DstAlpha:     common.BlendSrcAlpha,
						AlphaBlendOp: common.BlendOpDifferenceEXT,
						WriteMask:    common.ComponentAlpha,
					},
					{
						BlendEnabled: false,
						SrcColor:     common.BlendOneMinusSrcAlpha,
						DstColor:     common.BlendOneMinusConstantColor,
						ColorBlendOp: common.BlendOpDstEXT,
						SrcAlpha:     common.BlendConstantAlpha,
						DstAlpha:     common.BlendOne,
						AlphaBlendOp: common.BlendOpHardMixEXT,
						WriteMask:    common.ComponentRed,
					},
				},
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Same(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_DynamicStateSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(mocks.Exactly(device.Handle()), nil, driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pGraphicsPipelines *driver.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]driver.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			state := createInfo.FieldByName("pDynamicState").Elem()
			require.Equal(t, uint64(27), state.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO
			require.True(t, state.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), state.FieldByName("flags").Uint())
			require.Equal(t, uint64(2), state.FieldByName("dynamicStateCount").Uint())

			statesPtr := (*driver.VkDynamicState)(unsafe.Pointer(state.FieldByName("pDynamicStates").Elem().UnsafeAddr()))
			statesSlice := reflect.ValueOf(([]driver.VkDynamicState)(unsafe.Slice(statesPtr, 2)))

			require.Equal(t, uint64(5), statesSlice.Index(0).Uint())          // VK_DYNAMIC_STATE_DEPTH_BOUNDS
			require.Equal(t, uint64(1000267000), statesSlice.Index(1).Uint()) // VK_DYNAMIC_STATE_CULL_MODE_EXT

			return common.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []*core.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			DynamicState: &core.DynamicStateOptions{
				DynamicStates: []core.DynamicState{
					core.DynamicStateDepthBounds, core.DynamicStateCullModeEXT,
				},
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Same(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateComputePipelines_EmptySuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	shaderModule := mocks.EasyMockShaderModule(ctrl)
	pipelineHandle := mocks.NewFakePipeline()
	basePipeline := mocks.EasyDummyPipeline(t, device, loader)

	mockDriver.EXPECT().VkCreateComputePipelines(mocks.Exactly(device.Handle()), nil, driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkComputePipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pGraphicsPipelines *driver.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]driver.VkComputePipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			require.Equal(t, uint64(29), createInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO
			require.True(t, createInfo.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000800), createInfo.FieldByName("flags").Uint()) // VK_PIPELINE_CREATE_LIBRARY_BIT_KHR

			shader := createInfo.FieldByName("stage")

			require.Equal(t, uint64(1), shader.FieldByName("flags").Uint())    // VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT
			require.Equal(t, uint64(0x20), shader.FieldByName("stage").Uint()) // VK_SHADER_STAGE_COMPUTE_BIT
			module := (driver.VkShaderModule)(unsafe.Pointer(shader.FieldByName("module").Elem().UnsafeAddr()))
			require.Same(t, shaderModule.Handle(), module)
			namePtr := (*driver.Char)(unsafe.Pointer(shader.FieldByName("pName").Elem().UnsafeAddr()))
			nameSlice := ([]driver.Char)(unsafe.Slice(namePtr, 256))
			expectedName := "some compute shader"
			for i, r := range expectedName {
				require.Equal(t, r, rune(nameSlice[i]))
			}
			require.Equal(t, 0, int(nameSlice[len(expectedName)]))

			specInfo := shader.FieldByName("pSpecializationInfo").Elem()
			require.Equal(t, uint64(2), specInfo.FieldByName("mapEntryCount").Uint())
			mapEntryPtr := (*driver.VkSpecializationMapEntry)(unsafe.Pointer(specInfo.FieldByName("pMapEntries").Elem().UnsafeAddr()))
			mapEntrySlice := reflect.ValueOf(([]driver.VkSpecializationMapEntry)(unsafe.Slice(mapEntryPtr, 2)))

			firstEntryID := mapEntrySlice.Index(0).FieldByName("constantID").Uint()
			secondEntryID := mapEntrySlice.Index(1).FieldByName("constantID").Uint()

			// We deliver the map as a go map so it could come out in any order
			require.ElementsMatch(t, []uint64{1, 2}, []uint64{firstEntryID, secondEntryID})

			firstEntryOffset := mapEntrySlice.Index(0).FieldByName("offset").Uint()
			firstEntrySize := mapEntrySlice.Index(0).FieldByName("size").Uint()

			secondEntryOffset := mapEntrySlice.Index(1).FieldByName("offset").Uint()
			secondEntrySize := mapEntrySlice.Index(1).FieldByName("size").Uint()

			require.Equal(t, uint64(0), firstEntryOffset)
			require.Equal(t, firstEntrySize, secondEntryOffset)

			require.Equal(t, firstEntrySize+secondEntrySize, specInfo.FieldByName("dataSize").Uint())
			dataPtr := (*byte)(unsafe.Pointer(specInfo.FieldByName("pData").Pointer()))
			dataSlice := ([]byte)(unsafe.Slice(dataPtr, firstEntrySize+secondEntrySize))

			firstEntryBytes := bytes.NewBuffer(dataSlice[:firstEntrySize])
			secondEntryBytes := bytes.NewBuffer(dataSlice[firstEntrySize : firstEntrySize+secondEntrySize])

			firstValueBytes := firstEntryBytes
			secondValueBytes := secondEntryBytes

			if firstEntryID == 2 {
				firstValueBytes = secondEntryBytes
				secondValueBytes = firstEntryBytes
			}

			var boolVal uint32
			err = binary.Read(firstValueBytes, binary.LittleEndian, &boolVal)
			require.NoError(t, err)
			require.Equal(t, uint32(1), boolVal)

			var floatVal float64
			err = binary.Read(secondValueBytes, binary.LittleEndian, &floatVal)
			require.NoError(t, err)
			require.Equal(t, float64(7.6), floatVal)

			actualLayout := (driver.VkPipelineLayout)(unsafe.Pointer(createInfo.FieldByName("layout").Elem().UnsafeAddr()))
			require.Same(t, layout.Handle(), actualLayout)
			actualBasePipeline := (driver.VkPipeline)(unsafe.Pointer(createInfo.FieldByName("basePipelineHandle").Elem().UnsafeAddr()))
			require.Same(t, basePipeline.Handle(), actualBasePipeline)
			require.Equal(t, int64(3), createInfo.FieldByName("basePipelineIndex").Int())

			return common.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateComputePipelines(device, nil, nil, []*core.ComputePipelineOptions{
		{
			Flags: core.PipelineLibraryKHR,
			Shader: core.ShaderStage{
				Flags:  core.ShaderStageAllowVaryingSubgroupSizeEXT,
				Name:   "some compute shader",
				Stage:  common.StageCompute,
				Shader: shaderModule,
				SpecializationInfo: map[uint32]interface{}{
					1: true,
					2: float64(7.6),
				},
			},
			Layout: layout,

			BasePipeline:      basePipeline,
			BasePipelineIndex: 3,
		},
	})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Same(t, pipelineHandle, pipelines[0].Handle())
}
