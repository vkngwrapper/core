package core1_0_test

import (
	"bytes"
	"encoding/binary"
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	internal_mocks "github.com/CannibalVox/VKng/core/internal/mocks"
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

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()
	basePipeline := internal_mocks.EasyDummyPipeline(t, device, loader)

	mockDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), driver.VkPipelineCache(driver.NullHandle), driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pGraphicsPipelines *driver.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]driver.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			require.Equal(t, uint64(28), createInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO
			require.True(t, createInfo.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000002), createInfo.FieldByName("flags").Uint()) // VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT
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
			require.Equal(t, layout.Handle(), actualLayout)
			actualRenderPass := (driver.VkRenderPass)(unsafe.Pointer(createInfo.FieldByName("renderPass").Elem().UnsafeAddr()))
			require.Equal(t, renderPass.Handle(), actualRenderPass)
			require.Equal(t, uint64(1), createInfo.FieldByName("subpass").Uint())
			actualBasePipeline := (driver.VkPipeline)(unsafe.Pointer(createInfo.FieldByName("basePipelineHandle").Elem().UnsafeAddr()))
			require.Equal(t, basePipeline.Handle(), actualBasePipeline)
			require.Equal(t, int64(3), createInfo.FieldByName("basePipelineIndex").Int())

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{
		{
			Flags:             core1_0.PipelineCreateAllowDerivatives,
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

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()
	shaderModule1 := mocks.EasyMockShaderModule(ctrl)
	shaderModule2 := mocks.EasyMockShaderModule(ctrl)

	mockDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), driver.VkPipelineCache(driver.NullHandle), driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
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
			require.Equal(t, uint64(0), shader.FieldByName("flags").Uint())
			require.Equal(t, uint64(8), shader.FieldByName("stage").Uint()) // VK_SHADER_STAGE_GEOMETRY_BIT
			module := (driver.VkShaderModule)(unsafe.Pointer(shader.FieldByName("module").Elem().UnsafeAddr()))
			require.Equal(t, shaderModule1.Handle(), module)
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
			require.Equal(t, shaderModule2.Handle(), module)
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

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			ShaderStages: []core1_0.ShaderStageOptions{
				{
					Flags:  0,
					Name:   "some shader 1",
					Stage:  core1_0.StageGeometry,
					Shader: shaderModule1,
				},
				{
					Name:   "another shader 2",
					Stage:  core1_0.StageFragment,
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
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_ShaderStagesFailure_InvalidSpecializationValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	shaderModule1 := mocks.EasyMockShaderModule(ctrl)
	shaderModule2 := mocks.EasyMockShaderModule(ctrl)

	_, _, err = loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			ShaderStages: []core1_0.ShaderStageOptions{
				{
					Name:   "some shader 1",
					Stage:  core1_0.StageGeometry,
					Shader: shaderModule1,
				},
				{
					Name:   "another shader 2",
					Stage:  core1_0.StageFragment,
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

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), driver.VkPipelineCache(driver.NullHandle), driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
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

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			VertexInput: &core1_0.VertexInputOptions{
				VertexAttributeDescriptions: []core1_0.VertexAttributeDescription{
					{
						Location: 1,
						Binding:  3,
						Format:   core1_0.DataFormatA1R5G5B5UnsignedNormalized,
						Offset:   5,
					},
					{
						Location: 7,
						Binding:  11,
						Format:   core1_0.DataFormatA2B10G10R10UnsignedNormalized,
						Offset:   13,
					},
				},
				VertexBindingDescriptions: []core1_0.VertexBindingDescription{
					{
						InputRate: core1_0.RateInstance,
						Binding:   17,
						Stride:    19,
					},
				},
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_InputAssemblySuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), driver.VkPipelineCache(driver.NullHandle), driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
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

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			InputAssembly: &core1_0.InputAssemblyOptions{
				Topology:               core1_0.TopologyLineList,
				EnablePrimitiveRestart: true,
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_TessellationSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), driver.VkPipelineCache(driver.NullHandle), driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
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

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			Tessellation: &core1_0.TessellationOptions{
				PatchControlPoints: 3,
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_ViewportSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), driver.VkPipelineCache(driver.NullHandle), driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
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

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			Viewport: &core1_0.ViewportOptions{
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
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_RasterizationSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), driver.VkPipelineCache(driver.NullHandle), driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
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

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			Rasterization: &core1_0.RasterizationOptions{
				DepthClamp:        true,
				RasterizerDiscard: true,

				PolygonMode: core1_0.PolygonModeLine,
				CullMode:    core1_0.CullFront | core1_0.CullBack,
				FrontFace:   core1_0.FrontFaceClockwise,

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
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_MultisampleSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), driver.VkPipelineCache(driver.NullHandle), driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
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

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			Multisample: &core1_0.MultisampleOptions{
				RasterizationSamples: core1_0.Samples64,
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
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_MultisampleSuccess_NoSampleMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), driver.VkPipelineCache(driver.NullHandle), driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
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

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			Multisample: &core1_0.MultisampleOptions{
				RasterizationSamples: core1_0.Samples64,
				SampleShading:        true,
				MinSampleShading:     2.3,
				AlphaToCoverage:      true,
				AlphaToOne:           true,
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_MultisampleFail_MismatchSampleMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)

	_, _, err = loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			Multisample: &core1_0.MultisampleOptions{
				RasterizationSamples: core1_0.Samples4,
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

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), driver.VkPipelineCache(driver.NullHandle), driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
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

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			DepthStencil: &core1_0.DepthStencilOptions{
				DepthTestEnable:       true,
				DepthWriteEnable:      true,
				DepthCompareOp:        core1_0.CompareEqual,
				DepthBoundsTestEnable: true,
				StencilTestEnable:     true,
				FrontStencilState: core1_0.StencilOpState{
					FailOp:      core1_0.StencilInvert,
					PassOp:      core1_0.StencilDecrementAndWrap,
					DepthFailOp: core1_0.StencilReplace,
					CompareOp:   core1_0.CompareGreaterOrEqual,
					CompareMask: 3,
					WriteMask:   5,
					Reference:   7,
				},
				BackStencilState: core1_0.StencilOpState{
					FailOp:      core1_0.StencilIncrementAndClamp,
					PassOp:      core1_0.StencilZero,
					DepthFailOp: core1_0.StencilKeep,
					CompareOp:   core1_0.CompareLess,
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
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_ColorBlendSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), driver.VkPipelineCache(driver.NullHandle), driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
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
			require.Equal(t, uint64(4), attachment.FieldByName("srcColorBlendFactor").Uint())  // VK_BLEND_FACTOR_DST_COLOR
			require.Equal(t, uint64(18), attachment.FieldByName("dstColorBlendFactor").Uint()) // VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA
			require.Equal(t, uint64(1), attachment.FieldByName("colorBlendOp").Uint())         // VK_BLEND_OP_SUBTRACT
			require.Equal(t, uint64(17), attachment.FieldByName("srcAlphaBlendFactor").Uint()) // VK_BLEND_FACTOR_SRC1_ALPHA
			require.Equal(t, uint64(6), attachment.FieldByName("dstAlphaBlendFactor").Uint())  // VK_BLEND_FACTOR_SRC_ALPHA
			require.Equal(t, uint64(3), attachment.FieldByName("alphaBlendOp").Uint())         // VK_BLEND_OP_MIN
			require.Equal(t, uint64(8), attachment.FieldByName("colorWriteMask").Uint())       // VK_COLOR_COMPONENT_A_BIT

			attachment = attachmentsSlice.Index(1)
			require.Equal(t, uint64(0), attachment.FieldByName("blendEnable").Uint())
			require.Equal(t, uint64(7), attachment.FieldByName("srcColorBlendFactor").Uint())  // VK_BLEND_FACTOR_ONE_MINUS_SRC_ALPHA
			require.Equal(t, uint64(11), attachment.FieldByName("dstColorBlendFactor").Uint()) // VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR
			require.Equal(t, uint64(0), attachment.FieldByName("colorBlendOp").Uint())         // VK_BLEND_OP_ADD
			require.Equal(t, uint64(12), attachment.FieldByName("srcAlphaBlendFactor").Uint()) // VK_BLEND_FACTOR_CONSTANT_ALPHA
			require.Equal(t, uint64(1), attachment.FieldByName("dstAlphaBlendFactor").Uint())  // VK_BLEND_FACTOR_ONE
			require.Equal(t, uint64(4), attachment.FieldByName("alphaBlendOp").Uint())         // VK_BLEND_OP_MAX
			require.Equal(t, uint64(1), attachment.FieldByName("colorWriteMask").Uint())       // VK_COLOR_COMPONENT_R_BIT

			constants := state.FieldByName("blendConstants")
			require.InDelta(t, 1.2, constants.Index(0).Float(), 0.0001)
			require.InDelta(t, 2.3, constants.Index(1).Float(), 0.0001)
			require.InDelta(t, 3.4, constants.Index(2).Float(), 0.0001)
			require.InDelta(t, 4.5, constants.Index(3).Float(), 0.0001)

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			ColorBlend: &core1_0.ColorBlendOptions{
				LogicOpEnabled: true,
				LogicOp:        core1_0.LogicOpCopyInverted,
				BlendConstants: [4]float32{1.2, 2.3, 3.4, 4.5},
				Attachments: []core1_0.ColorBlendAttachment{
					{
						BlendEnabled: true,
						SrcColor:     core1_0.BlendDstColor,
						DstColor:     core1_0.BlendOneMinusSrc1Alpha,
						ColorBlendOp: core1_0.BlendOpSubtract,
						SrcAlpha:     core1_0.BlendSrc1Alpha,
						DstAlpha:     core1_0.BlendSrcAlpha,
						AlphaBlendOp: core1_0.BlendOpMin,
						WriteMask:    common.ComponentAlpha,
					},
					{
						BlendEnabled: false,
						SrcColor:     core1_0.BlendOneMinusSrcAlpha,
						DstColor:     core1_0.BlendOneMinusConstantColor,
						ColorBlendOp: core1_0.BlendOpAdd,
						SrcAlpha:     core1_0.BlendConstantAlpha,
						DstAlpha:     core1_0.BlendOne,
						AlphaBlendOp: core1_0.BlendOpMax,
						WriteMask:    common.ComponentRed,
					},
				},
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_DynamicStateSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	pipelineHandle := mocks.NewFakePipeline()

	mockDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), driver.VkPipelineCache(driver.NullHandle), driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
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

			require.Equal(t, uint64(5), statesSlice.Index(0).Uint()) // VK_DYNAMIC_STATE_DEPTH_BOUNDS
			require.Equal(t, uint64(7), statesSlice.Index(1).Uint()) // VK_DYNAMIC_STATE_STENCIL_WRITE_MASK

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{
		{
			Layout:     layout,
			RenderPass: renderPass,
			DynamicState: &core1_0.DynamicStateOptions{
				DynamicStates: []common.DynamicState{
					core1_0.DynamicStateDepthBounds, core1_0.DynamicStateStencilWriteMask,
				},
			},
		}})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateComputePipelines_EmptySuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	layout := mocks.EasyMockPipelineLayout(ctrl)
	shaderModule := mocks.EasyMockShaderModule(ctrl)
	pipelineHandle := mocks.NewFakePipeline()
	basePipeline := internal_mocks.EasyDummyPipeline(t, device, loader)

	mockDriver.EXPECT().VkCreateComputePipelines(device.Handle(), driver.VkPipelineCache(driver.NullHandle), driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkComputePipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pGraphicsPipelines *driver.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]driver.VkComputePipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			require.Equal(t, uint64(29), createInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO
			require.True(t, createInfo.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000004), createInfo.FieldByName("flags").Uint()) // VK_PIPELINE_CREATE_DERIVATIVE_BIT

			shader := createInfo.FieldByName("stage")

			require.Equal(t, uint64(0), shader.FieldByName("flags").Uint())
			require.Equal(t, uint64(0x20), shader.FieldByName("stage").Uint()) // VK_SHADER_STAGE_COMPUTE_BIT
			module := (driver.VkShaderModule)(unsafe.Pointer(shader.FieldByName("module").Elem().UnsafeAddr()))
			require.Equal(t, shaderModule.Handle(), module)
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
			require.Equal(t, layout.Handle(), actualLayout)
			actualBasePipeline := (driver.VkPipeline)(unsafe.Pointer(createInfo.FieldByName("basePipelineHandle").Elem().UnsafeAddr()))
			require.Equal(t, basePipeline.Handle(), actualBasePipeline)
			require.Equal(t, int64(3), createInfo.FieldByName("basePipelineIndex").Int())

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateComputePipelines(device, nil, nil, []core1_0.ComputePipelineOptions{
		{
			Flags: core1_0.PipelineCreateDerivative,
			Shader: core1_0.ShaderStageOptions{
				Flags:  0,
				Name:   "some compute shader",
				Stage:  core1_0.StageCompute,
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
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}
