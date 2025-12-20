package impl1_1_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_CreateDescriptorUpdateTemplate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)

	builder := &impl1_1.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_1, []string{}).(core1_1.Device)
	descriptorLayout := mocks.EasyMockDescriptorSetLayout(ctrl)
	pipelineLayout := mocks.EasyMockPipelineLayout(ctrl)

	handle := mocks.NewFakeDescriptorUpdateTemplate()

	coreDriver.EXPECT().VkCreateDescriptorUpdateTemplate(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		pCreateInfo *driver.VkDescriptorUpdateTemplateCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pDescriptorTemplate *driver.VkDescriptorUpdateTemplate,
	) (common.VkResult, error) {
		*pDescriptorTemplate = handle

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(1000085000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
		require.Equal(t, uint64(2), val.FieldByName("descriptorUpdateEntryCount").Uint())
		require.Equal(t, uint64(0), val.FieldByName("templateType").Uint())
		require.Equal(t, descriptorLayout.Handle(), driver.VkDescriptorSetLayout(val.FieldByName("descriptorSetLayout").UnsafePointer()))
		require.Equal(t, pipelineLayout.Handle(), driver.VkPipelineLayout(val.FieldByName("pipelineLayout").UnsafePointer()))
		require.Equal(t, uint64(0), val.FieldByName("pipelineBindPoint").Uint())
		require.Equal(t, uint64(31), val.FieldByName("set").Uint())

		entriesPtr := (*driver.VkDescriptorUpdateTemplateEntry)(val.FieldByName("pDescriptorUpdateEntries").UnsafePointer())
		entriesSlice := ([]driver.VkDescriptorUpdateTemplateEntry)(unsafe.Slice(entriesPtr, 2))
		entries := reflect.ValueOf(entriesSlice)

		entry := entries.Index(0)
		require.Equal(t, uint64(1), entry.FieldByName("dstBinding").Uint())
		require.Equal(t, uint64(3), entry.FieldByName("dstArrayElement").Uint())
		require.Equal(t, uint64(5), entry.FieldByName("descriptorCount").Uint())
		require.Equal(t, uint64(1), entry.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER
		require.Equal(t, uint64(7), entry.FieldByName("offset").Uint())
		require.Equal(t, uint64(11), entry.FieldByName("stride").Uint())

		entry = entries.Index(1)
		require.Equal(t, uint64(13), entry.FieldByName("dstBinding").Uint())
		require.Equal(t, uint64(17), entry.FieldByName("dstArrayElement").Uint())
		require.Equal(t, uint64(19), entry.FieldByName("descriptorCount").Uint())
		require.Equal(t, uint64(7), entry.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_STORAGE_BUFFER
		require.Equal(t, uint64(23), entry.FieldByName("offset").Uint())
		require.Equal(t, uint64(29), entry.FieldByName("stride").Uint())

		return core1_0.VKSuccess, nil
	})
	coreDriver.EXPECT().VkDestroyDescriptorUpdateTemplate(
		device.Handle(),
		handle,
		gomock.Nil(),
	)

	template, _, err := device.CreateDescriptorUpdateTemplate(core1_1.DescriptorUpdateTemplateCreateInfo{
		DescriptorUpdateEntries: []core1_1.DescriptorUpdateTemplateEntry{
			{
				DstBinding:      1,
				DstArrayElement: 3,
				DescriptorCount: 5,
				DescriptorType:  core1_0.DescriptorTypeCombinedImageSampler,
				Offset:          7,
				Stride:          11,
			},
			{
				DstBinding:      13,
				DstArrayElement: 17,
				DescriptorCount: 19,
				DescriptorType:  core1_0.DescriptorTypeStorageBuffer,
				Offset:          23,
				Stride:          29,
			},
		},
		TemplateType:        core1_1.DescriptorUpdateTemplateTypeDescriptorSet,
		DescriptorSetLayout: descriptorLayout,
		PipelineBindPoint:   core1_0.PipelineBindPointGraphics,
		PipelineLayout:      pipelineLayout,
		Set:                 31,
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, template)
	require.Equal(t, handle, template.Handle())

	template.Destroy(nil)
}

func TestVulkanDescriptorTemplate_UpdateDescriptorSetFromBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)

	builder := &impl1_1.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_1, []string{}).(core1_1.Device)
	descriptorSet := mocks.EasyMockDescriptorSet(ctrl)
	buffer := mocks.EasyMockBuffer(ctrl)

	handle := mocks.NewFakeDescriptorUpdateTemplate()

	coreDriver.EXPECT().VkCreateDescriptorUpdateTemplate(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		pCreateInfo *driver.VkDescriptorUpdateTemplateCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pDescriptorTemplate *driver.VkDescriptorUpdateTemplate,
	) (common.VkResult, error) {
		*pDescriptorTemplate = handle

		return core1_0.VKSuccess, nil
	})

	coreDriver.EXPECT().VkUpdateDescriptorSetWithTemplate(
		device.Handle(),
		descriptorSet.Handle(),
		handle,
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		descriptorSet driver.VkDescriptorSet,
		template driver.VkDescriptorUpdateTemplate,
		pData unsafe.Pointer,
	) {
		infoPtr := (*driver.VkDescriptorBufferInfo)(pData)
		info := reflect.ValueOf(infoPtr).Elem()
		require.Equal(t, buffer.Handle(), (driver.VkBuffer)(info.FieldByName("buffer").UnsafePointer()))
		require.Equal(t, uint64(1), info.FieldByName("offset").Uint())
		require.Equal(t, uint64(3), info.FieldByName("_range").Uint())
	})

	template, _, err := device.CreateDescriptorUpdateTemplate(core1_1.DescriptorUpdateTemplateCreateInfo{}, nil)
	require.NoError(t, err)
	require.NotNil(t, template)

	template.UpdateDescriptorSetFromBuffer(descriptorSet, core1_0.DescriptorBufferInfo{
		Buffer: buffer,
		Offset: 1,
		Range:  3,
	})
}

func TestVulkanDescriptorTemplate_UpdateDescriptorSetFromImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)

	builder := &impl1_1.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_1, []string{}).(core1_1.Device)
	descriptorSet := mocks.EasyMockDescriptorSet(ctrl)
	sampler := mocks.EasyMockSampler(ctrl)
	imageView := mocks.EasyMockImageView(ctrl)

	handle := mocks.NewFakeDescriptorUpdateTemplate()

	coreDriver.EXPECT().VkCreateDescriptorUpdateTemplate(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		pCreateInfo *driver.VkDescriptorUpdateTemplateCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pDescriptorTemplate *driver.VkDescriptorUpdateTemplate,
	) (common.VkResult, error) {
		*pDescriptorTemplate = handle

		return core1_0.VKSuccess, nil
	})

	coreDriver.EXPECT().VkUpdateDescriptorSetWithTemplate(
		device.Handle(),
		descriptorSet.Handle(),
		handle,
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		descriptorSet driver.VkDescriptorSet,
		template driver.VkDescriptorUpdateTemplate,
		pData unsafe.Pointer,
	) {
		infoPtr := (*driver.VkDescriptorImageInfo)(pData)
		info := reflect.ValueOf(infoPtr).Elem()
		require.Equal(t, sampler.Handle(), (driver.VkSampler)(info.FieldByName("sampler").UnsafePointer()))
		require.Equal(t, imageView.Handle(), (driver.VkImageView)(info.FieldByName("imageView").UnsafePointer()))
		require.Equal(t, uint64(7), info.FieldByName("imageLayout").Uint()) // VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL
	})

	template, _, err := device.CreateDescriptorUpdateTemplate(core1_1.DescriptorUpdateTemplateCreateInfo{}, nil)
	require.NoError(t, err)
	require.NotNil(t, template)

	template.UpdateDescriptorSetFromImage(descriptorSet, core1_0.DescriptorImageInfo{
		Sampler:     sampler,
		ImageView:   imageView,
		ImageLayout: core1_0.ImageLayoutTransferDstOptimal,
	})
}

func TestVulkanDescriptorTemplate_UpdateDescriptorSetFromObjectHandle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)

	builder := &impl1_1.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_1, []string{}).(core1_1.Device)
	descriptorSet := mocks.EasyMockDescriptorSet(ctrl)
	bufferView := mocks.EasyMockBufferView(ctrl)

	handle := mocks.NewFakeDescriptorUpdateTemplate()

	coreDriver.EXPECT().VkCreateDescriptorUpdateTemplate(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		pCreateInfo *driver.VkDescriptorUpdateTemplateCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pDescriptorTemplate *driver.VkDescriptorUpdateTemplate,
	) (common.VkResult, error) {
		*pDescriptorTemplate = handle

		return core1_0.VKSuccess, nil
	})

	coreDriver.EXPECT().VkUpdateDescriptorSetWithTemplate(
		device.Handle(),
		descriptorSet.Handle(),
		handle,
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		descriptorSet driver.VkDescriptorSet,
		template driver.VkDescriptorUpdateTemplate,
		pData unsafe.Pointer,
	) {
		info := (driver.VkBufferView)(pData)
		require.Equal(t, bufferView.Handle(), info)
	})

	template, _, err := device.CreateDescriptorUpdateTemplate(core1_1.DescriptorUpdateTemplateCreateInfo{}, nil)
	require.NoError(t, err)
	require.NotNil(t, template)

	template.UpdateDescriptorSetFromObjectHandle(descriptorSet, driver.VulkanHandle(bufferView.Handle()))
}
