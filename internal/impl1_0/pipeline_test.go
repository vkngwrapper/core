package impl1_0_test

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateGraphicsPipelines_EmptySuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	layout := mocks.NewDummyPipelineLayout(device)
	renderPass := mocks.NewDummyRenderPass(device)
	pipelineHandle := mocks.NewFakePipeline()
	basePipeline := mocks.NewDummyPipeline(device)

	mockLoader.EXPECT().VkCreateGraphicsPipelines(device.Handle(), loader.VkPipelineCache(loader.NullHandle), loader.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, cache loader.VkPipelineCache, createInfoCount loader.Uint32, pCreateInfos *loader.VkGraphicsPipelineCreateInfo, pAllocator *loader.VkAllocationCallbacks, pGraphicsPipelines *loader.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]loader.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]loader.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
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
			actualLayout := (loader.VkPipelineLayout)(unsafe.Pointer(createInfo.FieldByName("layout").Elem().UnsafeAddr()))
			require.Equal(t, layout.Handle(), actualLayout)
			actualRenderPass := (loader.VkRenderPass)(unsafe.Pointer(createInfo.FieldByName("renderPass").Elem().UnsafeAddr()))
			require.Equal(t, renderPass.Handle(), actualRenderPass)
			require.Equal(t, uint64(1), createInfo.FieldByName("subpass").Uint())
			actualBasePipeline := (loader.VkPipeline)(unsafe.Pointer(createInfo.FieldByName("basePipelineHandle").Elem().UnsafeAddr()))
			require.Equal(t, basePipeline.Handle(), actualBasePipeline)
			require.Equal(t, int64(3), createInfo.FieldByName("basePipelineIndex").Int())

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := driver.CreateGraphicsPipelines(nil, nil,
		core1_0.GraphicsPipelineCreateInfo{
			Flags:             core1_0.PipelineCreateAllowDerivatives,
			Layout:            layout,
			RenderPass:        renderPass,
			Subpass:           1,
			BasePipeline:      basePipeline,
			BasePipelineIndex: 3,
		},
	)
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_ShaderStagesSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	layout := mocks.NewDummyPipelineLayout(device)
	renderPass := mocks.NewDummyRenderPass(device)
	pipelineHandle := mocks.NewFakePipeline()
	shaderModule1 := mocks.NewDummyShaderModule(device)
	shaderModule2 := mocks.NewDummyShaderModule(device)

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateGraphicsPipelines(device.Handle(), loader.VkPipelineCache(loader.NullHandle), loader.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, cache loader.VkPipelineCache, createInfoCount loader.Uint32, pCreateInfos *loader.VkGraphicsPipelineCreateInfo, pAllocator *loader.VkAllocationCallbacks, pGraphicsPipelines *loader.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]loader.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]loader.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			require.Equal(t, uint64(28), createInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO
			require.True(t, createInfo.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), createInfo.FieldByName("stageCount").Uint())

			shaderPtr := (*loader.VkPipelineShaderStageCreateInfo)(unsafe.Pointer(createInfo.FieldByName("pStages").Elem().UnsafeAddr()))
			shaderSlice := reflect.ValueOf(([]loader.VkPipelineShaderStageCreateInfo)(unsafe.Slice(shaderPtr, 2)))

			shader := shaderSlice.Index(0)
			require.Equal(t, uint64(18), shader.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
			require.True(t, shader.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), shader.FieldByName("flags").Uint())
			require.Equal(t, uint64(8), shader.FieldByName("stage").Uint()) // VK_SHADER_STAGE_GEOMETRY_BIT
			module := (loader.VkShaderModule)(unsafe.Pointer(shader.FieldByName("module").Elem().UnsafeAddr()))
			require.Equal(t, shaderModule1.Handle(), module)
			namePtr := (*loader.Char)(unsafe.Pointer(shader.FieldByName("pName").Elem().UnsafeAddr()))
			nameSlice := ([]loader.Char)(unsafe.Slice(namePtr, 256))
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
			module = (loader.VkShaderModule)(unsafe.Pointer(shader.FieldByName("module").Elem().UnsafeAddr()))
			require.Equal(t, shaderModule2.Handle(), module)
			namePtr = (*loader.Char)(unsafe.Pointer(shader.FieldByName("pName").Elem().UnsafeAddr()))
			nameSlice = ([]loader.Char)(unsafe.Slice(namePtr, 256))
			expectedName = "another shader 2"
			for i, r := range expectedName {
				require.Equal(t, r, rune(nameSlice[i]))
			}
			require.Equal(t, 0, int(nameSlice[len(expectedName)]))

			specInfo := shader.FieldByName("pSpecializationInfo").Elem()
			require.Equal(t, uint64(2), specInfo.FieldByName("mapEntryCount").Uint())
			mapEntryPtr := (*loader.VkSpecializationMapEntry)(unsafe.Pointer(specInfo.FieldByName("pMapEntries").Elem().UnsafeAddr()))
			mapEntrySlice := reflect.ValueOf(([]loader.VkSpecializationMapEntry)(unsafe.Slice(mapEntryPtr, 2)))

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
			err := binary.Read(firstValueBytes, binary.LittleEndian, &boolVal)
			require.NoError(t, err)
			require.Equal(t, uint32(1), boolVal)

			var floatVal float64
			err = binary.Read(secondValueBytes, binary.LittleEndian, &floatVal)
			require.NoError(t, err)
			require.Equal(t, float64(7.6), floatVal)

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := driver.CreateGraphicsPipelines(nil, nil,
		core1_0.GraphicsPipelineCreateInfo{
			Layout:     layout,
			RenderPass: renderPass,
			Stages: []core1_0.PipelineShaderStageCreateInfo{
				{
					Flags:  0,
					Name:   "some shader 1",
					Stage:  core1_0.StageGeometry,
					Module: shaderModule1,
				},
				{
					Name:   "another shader 2",
					Stage:  core1_0.StageFragment,
					Module: shaderModule2,
					SpecializationInfo: map[uint32]any{
						1: true,
						2: float64(7.6),
					},
				},
			},
		})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_ShaderStagesFailure_InvalidSpecializationValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	layout := mocks.NewDummyPipelineLayout(device)
	renderPass := mocks.NewDummyRenderPass(device)
	shaderModule1 := mocks.NewDummyShaderModule(device)
	shaderModule2 := mocks.NewDummyShaderModule(device)

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	_, _, err := driver.CreateGraphicsPipelines(nil, nil,
		core1_0.GraphicsPipelineCreateInfo{
			Layout:     layout,
			RenderPass: renderPass,
			Stages: []core1_0.PipelineShaderStageCreateInfo{
				{
					Name:   "some shader 1",
					Stage:  core1_0.StageGeometry,
					Module: shaderModule1,
				},
				{
					Name:   "another shader 2",
					Stage:  core1_0.StageFragment,
					Module: shaderModule2,
					SpecializationInfo: map[uint32]any{
						1: "wow, this is invalid",
						2: float64(7.6),
					},
				},
			},
		})
	require.EqualError(t, err, "failed to populate shader stage with specialization values: 1 -> wow, this is invalid: binary.Write: some values are not fixed-sized in type string")
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_VertexInputSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	layout := mocks.NewDummyPipelineLayout(device)
	renderPass := mocks.NewDummyRenderPass(device)
	pipelineHandle := mocks.NewFakePipeline()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateGraphicsPipelines(device.Handle(), loader.VkPipelineCache(loader.NullHandle), loader.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, cache loader.VkPipelineCache, createInfoCount loader.Uint32, pCreateInfos *loader.VkGraphicsPipelineCreateInfo, pAllocator *loader.VkAllocationCallbacks, pGraphicsPipelines *loader.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]loader.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]loader.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
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

			bindingDescPtr := (*loader.VkVertexInputBindingDescription)(unsafe.Pointer(state.FieldByName("pVertexBindingDescriptions").Elem().UnsafeAddr()))
			bindingDescSlice := reflect.ValueOf(([]loader.VkVertexInputBindingDescription)(unsafe.Slice(bindingDescPtr, 1)))

			binding := bindingDescSlice.Index(0)
			require.Equal(t, uint64(17), binding.FieldByName("binding").Uint())
			require.Equal(t, uint64(19), binding.FieldByName("stride").Uint())
			require.Equal(t, uint64(1), binding.FieldByName("inputRate").Uint())

			attDescPtr := (*loader.VkVertexInputAttributeDescription)(unsafe.Pointer(state.FieldByName("pVertexAttributeDescriptions").Elem().UnsafeAddr()))
			attDescSlice := reflect.ValueOf(([]loader.VkVertexInputAttributeDescription)(unsafe.Slice(attDescPtr, 2)))

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

	pipelines, _, err := driver.CreateGraphicsPipelines(nil, nil,
		core1_0.GraphicsPipelineCreateInfo{
			Layout:     layout,
			RenderPass: renderPass,
			VertexInputState: &core1_0.PipelineVertexInputStateCreateInfo{
				VertexAttributeDescriptions: []core1_0.VertexInputAttributeDescription{
					{
						Location: 1,
						Binding:  3,
						Format:   core1_0.FormatA1R5G5B5UnsignedNormalizedPacked,
						Offset:   5,
					},
					{
						Location: 7,
						Binding:  11,
						Format:   core1_0.FormatA2B10G10R10UnsignedNormalizedPacked,
						Offset:   13,
					},
				},
				VertexBindingDescriptions: []core1_0.VertexInputBindingDescription{
					{
						InputRate: core1_0.VertexInputRateInstance,
						Binding:   17,
						Stride:    19,
					},
				},
			},
		})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_InputAssemblySuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	layout := mocks.NewDummyPipelineLayout(device)
	renderPass := mocks.NewDummyRenderPass(device)
	pipelineHandle := mocks.NewFakePipeline()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateGraphicsPipelines(device.Handle(), loader.VkPipelineCache(loader.NullHandle), loader.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, cache loader.VkPipelineCache, createInfoCount loader.Uint32, pCreateInfos *loader.VkGraphicsPipelineCreateInfo, pAllocator *loader.VkAllocationCallbacks, pGraphicsPipelines *loader.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]loader.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]loader.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
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

	pipelines, _, err := driver.CreateGraphicsPipelines(nil, nil,
		core1_0.GraphicsPipelineCreateInfo{
			Layout:     layout,
			RenderPass: renderPass,
			InputAssemblyState: &core1_0.PipelineInputAssemblyStateCreateInfo{
				Topology:               core1_0.PrimitiveTopologyLineList,
				PrimitiveRestartEnable: true,
			},
		})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_TessellationSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	layout := mocks.NewDummyPipelineLayout(device)
	renderPass := mocks.NewDummyRenderPass(device)
	pipelineHandle := mocks.NewFakePipeline()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateGraphicsPipelines(device.Handle(), loader.VkPipelineCache(loader.NullHandle), loader.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, cache loader.VkPipelineCache, createInfoCount loader.Uint32, pCreateInfos *loader.VkGraphicsPipelineCreateInfo, pAllocator *loader.VkAllocationCallbacks, pGraphicsPipelines *loader.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]loader.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]loader.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			state := createInfo.FieldByName("pTessellationState").Elem()

			require.Equal(t, uint64(21), state.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO
			require.True(t, state.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), state.FieldByName("flags").Uint())
			require.Equal(t, uint64(3), state.FieldByName("patchControlPoints").Uint())

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := driver.CreateGraphicsPipelines(nil, nil,
		core1_0.GraphicsPipelineCreateInfo{
			Layout:     layout,
			RenderPass: renderPass,
			TessellationState: &core1_0.PipelineTessellationStateCreateInfo{
				PatchControlPoints: 3,
			},
		})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_ViewportSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	layout := mocks.NewDummyPipelineLayout(device)
	renderPass := mocks.NewDummyRenderPass(device)
	pipelineHandle := mocks.NewFakePipeline()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateGraphicsPipelines(device.Handle(), loader.VkPipelineCache(loader.NullHandle), loader.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, cache loader.VkPipelineCache, createInfoCount loader.Uint32, pCreateInfos *loader.VkGraphicsPipelineCreateInfo, pAllocator *loader.VkAllocationCallbacks, pGraphicsPipelines *loader.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]loader.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]loader.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			state := createInfo.FieldByName("pViewportState").Elem()
			require.Equal(t, uint64(22), state.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO
			require.True(t, state.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), state.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), state.FieldByName("viewportCount").Uint())
			require.Equal(t, uint64(2), state.FieldByName("scissorCount").Uint())

			viewportPtr := (*loader.VkViewport)(unsafe.Pointer(state.FieldByName("pViewports").Elem().UnsafeAddr()))
			viewportSlice := reflect.ValueOf(([]loader.VkViewport)(unsafe.Slice(viewportPtr, 1)))

			viewport := viewportSlice.Index(0)
			require.Equal(t, float64(1), viewport.FieldByName("x").Float())
			require.Equal(t, float64(3), viewport.FieldByName("y").Float())
			require.Equal(t, float64(5), viewport.FieldByName("width").Float())
			require.Equal(t, float64(7), viewport.FieldByName("height").Float())
			require.Equal(t, float64(11), viewport.FieldByName("minDepth").Float())
			require.Equal(t, float64(13), viewport.FieldByName("maxDepth").Float())

			scissorPtr := (*loader.VkRect2D)(unsafe.Pointer(state.FieldByName("pScissors").Elem().UnsafeAddr()))
			scissorSlice := reflect.ValueOf(([]loader.VkRect2D)(unsafe.Slice(scissorPtr, 2)))

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

	pipelines, _, err := driver.CreateGraphicsPipelines(nil, nil,
		core1_0.GraphicsPipelineCreateInfo{
			Layout:     layout,
			RenderPass: renderPass,
			ViewportState: &core1_0.PipelineViewportStateCreateInfo{
				Viewports: []core1_0.Viewport{
					{
						X:        1,
						Y:        3,
						Width:    5,
						Height:   7,
						MinDepth: 11,
						MaxDepth: 13,
					},
				},
				Scissors: []core1_0.Rect2D{
					{
						Offset: core1_0.Offset2D{X: 17, Y: 19},
						Extent: core1_0.Extent2D{Width: 23, Height: 29},
					},
					{
						Offset: core1_0.Offset2D{X: 31, Y: 37},
						Extent: core1_0.Extent2D{Width: 41, Height: 43},
					},
				},
			},
		})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_RasterizationSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	layout := mocks.NewDummyPipelineLayout(device)
	renderPass := mocks.NewDummyRenderPass(device)
	pipelineHandle := mocks.NewFakePipeline()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateGraphicsPipelines(device.Handle(), loader.VkPipelineCache(loader.NullHandle), loader.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, cache loader.VkPipelineCache, createInfoCount loader.Uint32, pCreateInfos *loader.VkGraphicsPipelineCreateInfo, pAllocator *loader.VkAllocationCallbacks, pGraphicsPipelines *loader.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]loader.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]loader.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
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

	pipelines, _, err := driver.CreateGraphicsPipelines(nil, nil,
		core1_0.GraphicsPipelineCreateInfo{
			Layout:     layout,
			RenderPass: renderPass,
			RasterizationState: &core1_0.PipelineRasterizationStateCreateInfo{
				DepthClampEnable:        true,
				RasterizerDiscardEnable: true,

				PolygonMode: core1_0.PolygonModeLine,
				CullMode:    core1_0.CullModeFront | core1_0.CullModeBack,
				FrontFace:   core1_0.FrontFaceClockwise,

				DepthBiasEnable:         true,
				DepthBiasClamp:          2.3,
				DepthBiasConstantFactor: 3.4,
				DepthBiasSlopeFactor:    4.5,

				LineWidth: 5.6,
			},
		})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_MultisampleSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	layout := mocks.NewDummyPipelineLayout(device)
	renderPass := mocks.NewDummyRenderPass(device)
	pipelineHandle := mocks.NewFakePipeline()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateGraphicsPipelines(device.Handle(), loader.VkPipelineCache(loader.NullHandle), loader.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, cache loader.VkPipelineCache, createInfoCount loader.Uint32, pCreateInfos *loader.VkGraphicsPipelineCreateInfo, pAllocator *loader.VkAllocationCallbacks, pGraphicsPipelines *loader.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]loader.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]loader.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
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

			sampleMaskPtr := (*loader.VkSampleMask)(unsafe.Pointer(state.FieldByName("pSampleMask").Elem().UnsafeAddr()))
			sampleMaskSlice := reflect.ValueOf(([]loader.VkSampleMask)(unsafe.Slice(sampleMaskPtr, 2)))
			require.Equal(t, uint64(1), sampleMaskSlice.Index(0).Uint())
			require.Equal(t, uint64(3), sampleMaskSlice.Index(1).Uint())

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := driver.CreateGraphicsPipelines(nil, nil,
		core1_0.GraphicsPipelineCreateInfo{
			Layout:     layout,
			RenderPass: renderPass,
			MultisampleState: &core1_0.PipelineMultisampleStateCreateInfo{
				RasterizationSamples:  core1_0.Samples64,
				SampleShadingEnable:   true,
				MinSampleShading:      2.3,
				SampleMask:            []uint32{1, 3},
				AlphaToCoverageEnable: true,
				AlphaToOneEnable:      true,
			},
		})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_MultisampleSuccess_NoSampleMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	layout := mocks.NewDummyPipelineLayout(device)
	renderPass := mocks.NewDummyRenderPass(device)
	pipelineHandle := mocks.NewFakePipeline()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateGraphicsPipelines(device.Handle(), loader.VkPipelineCache(loader.NullHandle), loader.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, cache loader.VkPipelineCache, createInfoCount loader.Uint32, pCreateInfos *loader.VkGraphicsPipelineCreateInfo, pAllocator *loader.VkAllocationCallbacks, pGraphicsPipelines *loader.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]loader.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]loader.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
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

	pipelines, _, err := driver.CreateGraphicsPipelines(nil, nil,
		core1_0.GraphicsPipelineCreateInfo{
			Layout:     layout,
			RenderPass: renderPass,
			MultisampleState: &core1_0.PipelineMultisampleStateCreateInfo{
				RasterizationSamples:  core1_0.Samples64,
				SampleShadingEnable:   true,
				MinSampleShading:      2.3,
				AlphaToCoverageEnable: true,
				AlphaToOneEnable:      true,
			},
		})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_MultisampleFail_MismatchSampleMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	layout := mocks.NewDummyPipelineLayout(device)
	renderPass := mocks.NewDummyRenderPass(device)

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	_, _, err := driver.CreateGraphicsPipelines(nil, nil,
		core1_0.GraphicsPipelineCreateInfo{
			Layout:     layout,
			RenderPass: renderPass,
			MultisampleState: &core1_0.PipelineMultisampleStateCreateInfo{
				RasterizationSamples:  core1_0.Samples4,
				SampleShadingEnable:   true,
				MinSampleShading:      2.3,
				SampleMask:            []uint32{1, 3},
				AlphaToCoverageEnable: true,
				AlphaToOneEnable:      true,
			},
		})
	require.EqualError(t, err, "expected a sample mask size of 1, because 4 rasterization samples were specified- however, received a sample mask size of 2")
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_DepthStencilSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	layout := mocks.NewDummyPipelineLayout(device)
	renderPass := mocks.NewDummyRenderPass(device)
	pipelineHandle := mocks.NewFakePipeline()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateGraphicsPipelines(device.Handle(), loader.VkPipelineCache(loader.NullHandle), loader.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, cache loader.VkPipelineCache, createInfoCount loader.Uint32, pCreateInfos *loader.VkGraphicsPipelineCreateInfo, pAllocator *loader.VkAllocationCallbacks, pGraphicsPipelines *loader.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]loader.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]loader.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
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

	pipelines, _, err := driver.CreateGraphicsPipelines(nil, nil,
		core1_0.GraphicsPipelineCreateInfo{
			Layout:     layout,
			RenderPass: renderPass,
			DepthStencilState: &core1_0.PipelineDepthStencilStateCreateInfo{
				DepthTestEnable:       true,
				DepthWriteEnable:      true,
				DepthCompareOp:        core1_0.CompareOpEqual,
				DepthBoundsTestEnable: true,
				StencilTestEnable:     true,
				Front: core1_0.StencilOpState{
					FailOp:      core1_0.StencilInvert,
					PassOp:      core1_0.StencilDecrementAndWrap,
					DepthFailOp: core1_0.StencilReplace,
					CompareOp:   core1_0.CompareOpGreaterOrEqual,
					CompareMask: 3,
					WriteMask:   5,
					Reference:   7,
				},
				Back: core1_0.StencilOpState{
					FailOp:      core1_0.StencilIncrementAndClamp,
					PassOp:      core1_0.StencilZero,
					DepthFailOp: core1_0.StencilKeep,
					CompareOp:   core1_0.CompareOpLess,
					CompareMask: 11,
					WriteMask:   13,
					Reference:   17,
				},
				MinDepthBounds: 2.3,
				MaxDepthBounds: 3.4,
			},
		})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_ColorBlendSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	layout := mocks.NewDummyPipelineLayout(device)
	renderPass := mocks.NewDummyRenderPass(device)
	pipelineHandle := mocks.NewFakePipeline()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateGraphicsPipelines(device.Handle(), loader.VkPipelineCache(loader.NullHandle), loader.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, cache loader.VkPipelineCache, createInfoCount loader.Uint32, pCreateInfos *loader.VkGraphicsPipelineCreateInfo, pAllocator *loader.VkAllocationCallbacks, pGraphicsPipelines *loader.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]loader.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]loader.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			state := createInfo.FieldByName("pColorBlendState").Elem()
			require.Equal(t, uint64(26), state.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO
			require.True(t, state.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), state.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), state.FieldByName("logicOpEnable").Uint())
			require.Equal(t, uint64(12), state.FieldByName("logicOp").Uint()) // VK_LOGIC_OP_COPY_INVERTED
			require.Equal(t, uint64(2), state.FieldByName("attachmentCount").Uint())

			attachmentsPtr := (*loader.VkPipelineColorBlendAttachmentState)(unsafe.Pointer(state.FieldByName("pAttachments").Elem().UnsafeAddr()))
			attachmentsSlice := reflect.ValueOf(([]loader.VkPipelineColorBlendAttachmentState)(unsafe.Slice(attachmentsPtr, 2)))

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

	pipelines, _, err := driver.CreateGraphicsPipelines(nil, nil,
		core1_0.GraphicsPipelineCreateInfo{
			Layout:     layout,
			RenderPass: renderPass,
			ColorBlendState: &core1_0.PipelineColorBlendStateCreateInfo{
				LogicOpEnabled: true,
				LogicOp:        core1_0.LogicOpCopyInverted,
				BlendConstants: [4]float32{1.2, 2.3, 3.4, 4.5},
				Attachments: []core1_0.PipelineColorBlendAttachmentState{
					{
						BlendEnabled:        true,
						SrcColorBlendFactor: core1_0.BlendFactorDstColor,
						DstColorBlendFactor: core1_0.BlendFactorOneMinusSrc1Alpha,
						ColorBlendOp:        core1_0.BlendOpSubtract,
						SrcAlphaBlendFactor: core1_0.BlendFactorSrc1Alpha,
						DstAlphaBlendFactor: core1_0.BlendFactorSrcAlpha,
						AlphaBlendOp:        core1_0.BlendOpMin,
						ColorWriteMask:      core1_0.ColorComponentAlpha,
					},
					{
						BlendEnabled:        false,
						SrcColorBlendFactor: core1_0.BlendFactorOneMinusSrcAlpha,
						DstColorBlendFactor: core1_0.BlendFactorOneMinusConstantColor,
						ColorBlendOp:        core1_0.BlendOpAdd,
						SrcAlphaBlendFactor: core1_0.BlendFactorConstantAlpha,
						DstAlphaBlendFactor: core1_0.BlendFactorOne,
						AlphaBlendOp:        core1_0.BlendOpMax,
						ColorWriteMask:      core1_0.ColorComponentRed,
					},
				},
			},
		})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateGraphicsPipelines_DynamicStateSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	layout := mocks.NewDummyPipelineLayout(device)
	renderPass := mocks.NewDummyRenderPass(device)
	pipelineHandle := mocks.NewFakePipeline()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateGraphicsPipelines(device.Handle(), loader.VkPipelineCache(loader.NullHandle), loader.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, cache loader.VkPipelineCache, createInfoCount loader.Uint32, pCreateInfos *loader.VkGraphicsPipelineCreateInfo, pAllocator *loader.VkAllocationCallbacks, pGraphicsPipelines *loader.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]loader.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]loader.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			state := createInfo.FieldByName("pDynamicState").Elem()
			require.Equal(t, uint64(27), state.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO
			require.True(t, state.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), state.FieldByName("flags").Uint())
			require.Equal(t, uint64(2), state.FieldByName("dynamicStateCount").Uint())

			statesPtr := (*loader.VkDynamicState)(unsafe.Pointer(state.FieldByName("pDynamicStates").Elem().UnsafeAddr()))
			statesSlice := reflect.ValueOf(([]loader.VkDynamicState)(unsafe.Slice(statesPtr, 2)))

			require.Equal(t, uint64(5), statesSlice.Index(0).Uint()) // VK_DYNAMIC_STATE_DEPTH_BOUNDS
			require.Equal(t, uint64(7), statesSlice.Index(1).Uint()) // VK_DYNAMIC_STATE_STENCIL_WRITE_MASK

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := driver.CreateGraphicsPipelines(nil, nil,
		core1_0.GraphicsPipelineCreateInfo{
			Layout:     layout,
			RenderPass: renderPass,
			DynamicState: &core1_0.PipelineDynamicStateCreateInfo{
				DynamicStates: []core1_0.DynamicState{
					core1_0.DynamicStateDepthBounds, core1_0.DynamicStateStencilWriteMask,
				},
			},
		})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateComputePipelines_EmptySuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	layout := mocks.NewDummyPipelineLayout(device)
	shaderModule := mocks.NewDummyShaderModule(device)
	pipelineHandle := mocks.NewFakePipeline()
	basePipeline := mocks.NewDummyPipeline(device)

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateComputePipelines(device.Handle(), loader.VkPipelineCache(loader.NullHandle), loader.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, cache loader.VkPipelineCache, createInfoCount loader.Uint32, pCreateInfos *loader.VkComputePipelineCreateInfo, pAllocator *loader.VkAllocationCallbacks, pGraphicsPipelines *loader.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]loader.VkComputePipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]loader.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			require.Equal(t, uint64(29), createInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO
			require.True(t, createInfo.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000004), createInfo.FieldByName("flags").Uint()) // VK_PIPELINE_CREATE_DERIVATIVE_BIT

			shader := createInfo.FieldByName("stage")

			require.Equal(t, uint64(0), shader.FieldByName("flags").Uint())
			require.Equal(t, uint64(0x20), shader.FieldByName("stage").Uint()) // VK_SHADER_STAGE_COMPUTE_BIT
			module := (loader.VkShaderModule)(unsafe.Pointer(shader.FieldByName("module").Elem().UnsafeAddr()))
			require.Equal(t, shaderModule.Handle(), module)
			namePtr := (*loader.Char)(unsafe.Pointer(shader.FieldByName("pName").Elem().UnsafeAddr()))
			nameSlice := ([]loader.Char)(unsafe.Slice(namePtr, 256))
			expectedName := "some compute shader"
			for i, r := range expectedName {
				require.Equal(t, r, rune(nameSlice[i]))
			}
			require.Equal(t, 0, int(nameSlice[len(expectedName)]))

			specInfo := shader.FieldByName("pSpecializationInfo").Elem()
			require.Equal(t, uint64(2), specInfo.FieldByName("mapEntryCount").Uint())
			mapEntryPtr := (*loader.VkSpecializationMapEntry)(unsafe.Pointer(specInfo.FieldByName("pMapEntries").Elem().UnsafeAddr()))
			mapEntrySlice := reflect.ValueOf(([]loader.VkSpecializationMapEntry)(unsafe.Slice(mapEntryPtr, 2)))

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
			err := binary.Read(firstValueBytes, binary.LittleEndian, &boolVal)
			require.NoError(t, err)
			require.Equal(t, uint32(1), boolVal)

			var floatVal float64
			err = binary.Read(secondValueBytes, binary.LittleEndian, &floatVal)
			require.NoError(t, err)
			require.Equal(t, float64(7.6), floatVal)

			actualLayout := (loader.VkPipelineLayout)(unsafe.Pointer(createInfo.FieldByName("layout").Elem().UnsafeAddr()))
			require.Equal(t, layout.Handle(), actualLayout)
			actualBasePipeline := (loader.VkPipeline)(unsafe.Pointer(createInfo.FieldByName("basePipelineHandle").Elem().UnsafeAddr()))
			require.Equal(t, basePipeline.Handle(), actualBasePipeline)
			require.Equal(t, int64(3), createInfo.FieldByName("basePipelineIndex").Int())

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := driver.CreateComputePipelines(nil, nil,
		core1_0.ComputePipelineCreateInfo{
			Flags: core1_0.PipelineCreateDerivative,
			Stage: core1_0.PipelineShaderStageCreateInfo{
				Flags:  0,
				Name:   "some compute shader",
				Stage:  core1_0.StageCompute,
				Module: shaderModule,
				SpecializationInfo: map[uint32]any{
					1: true,
					2: float64(7.6),
				},
			},
			Layout: layout,

			BasePipeline:      basePipeline,
			BasePipelineIndex: 3,
		},
	)
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}

func TestVulkanLoader1_0_CreateComputePipelines_NilBasePipeline(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	layout := mocks.NewDummyPipelineLayout(device)
	shaderModule := mocks.NewDummyShaderModule(device)
	pipelineHandle := mocks.NewFakePipeline()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateComputePipelines(device.Handle(), loader.VkPipelineCache(loader.NullHandle), loader.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, cache loader.VkPipelineCache, createInfoCount loader.Uint32, pCreateInfos *loader.VkComputePipelineCreateInfo, pAllocator *loader.VkAllocationCallbacks, pGraphicsPipelines *loader.VkPipeline) (common.VkResult, error) {
			createInfos := reflect.ValueOf(([]loader.VkComputePipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1)))
			pipelines := ([]loader.VkPipeline)(unsafe.Slice(pGraphicsPipelines, 1))
			pipelines[0] = pipelineHandle

			createInfo := createInfos.Index(0)
			require.Equal(t, uint64(29), createInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO
			require.True(t, createInfo.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000004), createInfo.FieldByName("flags").Uint()) // VK_PIPELINE_CREATE_DERIVATIVE_BIT

			shader := createInfo.FieldByName("stage")

			require.Equal(t, uint64(0), shader.FieldByName("flags").Uint())
			require.Equal(t, uint64(0x20), shader.FieldByName("stage").Uint()) // VK_SHADER_STAGE_COMPUTE_BIT
			module := (loader.VkShaderModule)(unsafe.Pointer(shader.FieldByName("module").Elem().UnsafeAddr()))
			require.Equal(t, shaderModule.Handle(), module)
			namePtr := (*loader.Char)(unsafe.Pointer(shader.FieldByName("pName").Elem().UnsafeAddr()))
			nameSlice := ([]loader.Char)(unsafe.Slice(namePtr, 256))
			expectedName := "some compute shader"
			for i, r := range expectedName {
				require.Equal(t, r, rune(nameSlice[i]))
			}
			require.Equal(t, 0, int(nameSlice[len(expectedName)]))

			specInfo := shader.FieldByName("pSpecializationInfo").Elem()
			require.Equal(t, uint64(2), specInfo.FieldByName("mapEntryCount").Uint())
			mapEntryPtr := (*loader.VkSpecializationMapEntry)(unsafe.Pointer(specInfo.FieldByName("pMapEntries").Elem().UnsafeAddr()))
			mapEntrySlice := reflect.ValueOf(([]loader.VkSpecializationMapEntry)(unsafe.Slice(mapEntryPtr, 2)))

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
			err := binary.Read(firstValueBytes, binary.LittleEndian, &boolVal)
			require.NoError(t, err)
			require.Equal(t, uint32(1), boolVal)

			var floatVal float64
			err = binary.Read(secondValueBytes, binary.LittleEndian, &floatVal)
			require.NoError(t, err)
			require.Equal(t, float64(7.6), floatVal)

			actualLayout := (loader.VkPipelineLayout)(unsafe.Pointer(createInfo.FieldByName("layout").Elem().UnsafeAddr()))
			require.Equal(t, layout.Handle(), actualLayout)
			require.True(t, createInfo.FieldByName("basePipelineHandle").IsZero())

			require.Equal(t, int64(-1), createInfo.FieldByName("basePipelineIndex").Int())

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := driver.CreateComputePipelines(nil, nil,
		core1_0.ComputePipelineCreateInfo{
			Flags: core1_0.PipelineCreateDerivative,
			Stage: core1_0.PipelineShaderStageCreateInfo{
				Flags:  0,
				Name:   "some compute shader",
				Stage:  core1_0.StageCompute,
				Module: shaderModule,
				SpecializationInfo: map[uint32]any{
					1: true,
					2: float64(7.6),
				},
			},
			Layout: layout,

			BasePipelineIndex: -1,
		},
	)
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.NotNil(t, pipelines[0])
	require.Equal(t, pipelineHandle, pipelines[0].Handle())
}
