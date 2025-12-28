package impl1_1_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_CreateDescriptorUpdateTemplate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	descriptorLayout := mocks.NewDummyDescriptorSetLayout(device)
	pipelineLayout := mocks.NewDummyPipelineLayout(device)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(device, coreLoader)

	handle := mocks.NewFakeDescriptorUpdateTemplate()

	coreLoader.EXPECT().VkCreateDescriptorUpdateTemplate(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		pCreateInfo *loader.VkDescriptorUpdateTemplateCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDescriptorTemplate *loader.VkDescriptorUpdateTemplate,
	) (common.VkResult, error) {
		*pDescriptorTemplate = handle

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(1000085000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
		require.Equal(t, uint64(2), val.FieldByName("descriptorUpdateEntryCount").Uint())
		require.Equal(t, uint64(0), val.FieldByName("templateType").Uint())
		require.Equal(t, descriptorLayout.Handle(), loader.VkDescriptorSetLayout(val.FieldByName("descriptorSetLayout").UnsafePointer()))
		require.Equal(t, pipelineLayout.Handle(), loader.VkPipelineLayout(val.FieldByName("pipelineLayout").UnsafePointer()))
		require.Equal(t, uint64(0), val.FieldByName("pipelineBindPoint").Uint())
		require.Equal(t, uint64(31), val.FieldByName("set").Uint())

		entriesPtr := (*loader.VkDescriptorUpdateTemplateEntry)(val.FieldByName("pDescriptorUpdateEntries").UnsafePointer())
		entriesSlice := ([]loader.VkDescriptorUpdateTemplateEntry)(unsafe.Slice(entriesPtr, 2))
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
	coreLoader.EXPECT().VkDestroyDescriptorUpdateTemplate(
		device.Handle(),
		handle,
		gomock.Nil(),
	)

	template, _, err := driver.CreateDescriptorUpdateTemplate(device, core1_1.DescriptorUpdateTemplateCreateInfo{
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

	driver.DestroyDescriptorUpdateTemplate(template, nil)
}

func TestVulkanDescriptorTemplate_UpdateDescriptorSetFromBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	pool := mocks.NewDummyDescriptorPool(device)
	descriptorSet := mocks.NewDummyDescriptorSet(pool, device)
	buffer := mocks.NewDummyBuffer(device)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(device, coreLoader)

	handle := mocks.NewFakeDescriptorUpdateTemplate()

	coreLoader.EXPECT().VkCreateDescriptorUpdateTemplate(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		pCreateInfo *loader.VkDescriptorUpdateTemplateCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDescriptorTemplate *loader.VkDescriptorUpdateTemplate,
	) (common.VkResult, error) {
		*pDescriptorTemplate = handle

		return core1_0.VKSuccess, nil
	})

	coreLoader.EXPECT().VkUpdateDescriptorSetWithTemplate(
		device.Handle(),
		descriptorSet.Handle(),
		handle,
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		descriptorSet loader.VkDescriptorSet,
		template loader.VkDescriptorUpdateTemplate,
		pData unsafe.Pointer,
	) {
		infoPtr := (*loader.VkDescriptorBufferInfo)(pData)
		info := reflect.ValueOf(infoPtr).Elem()
		require.Equal(t, buffer.Handle(), (loader.VkBuffer)(info.FieldByName("buffer").UnsafePointer()))
		require.Equal(t, uint64(1), info.FieldByName("offset").Uint())
		require.Equal(t, uint64(3), info.FieldByName("_range").Uint())
	})

	template, _, err := driver.CreateDescriptorUpdateTemplate(device, core1_1.DescriptorUpdateTemplateCreateInfo{}, nil)
	require.NoError(t, err)
	require.NotNil(t, template)

	driver.UpdateDescriptorSetWithTemplateFromBuffer(descriptorSet, template, core1_0.DescriptorBufferInfo{
		Buffer: buffer,
		Offset: 1,
		Range:  3,
	})
}

func TestVulkanDescriptorTemplate_UpdateDescriptorSetFromImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	pool := mocks.NewDummyDescriptorPool(device)
	descriptorSet := mocks.NewDummyDescriptorSet(pool, device)
	sampler := mocks.NewDummySampler(device)
	imageView := mocks.NewDummyImageView(device)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(device, coreLoader)

	handle := mocks.NewFakeDescriptorUpdateTemplate()

	coreLoader.EXPECT().VkCreateDescriptorUpdateTemplate(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		pCreateInfo *loader.VkDescriptorUpdateTemplateCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDescriptorTemplate *loader.VkDescriptorUpdateTemplate,
	) (common.VkResult, error) {
		*pDescriptorTemplate = handle

		return core1_0.VKSuccess, nil
	})

	coreLoader.EXPECT().VkUpdateDescriptorSetWithTemplate(
		device.Handle(),
		descriptorSet.Handle(),
		handle,
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		descriptorSet loader.VkDescriptorSet,
		template loader.VkDescriptorUpdateTemplate,
		pData unsafe.Pointer,
	) {
		infoPtr := (*loader.VkDescriptorImageInfo)(pData)
		info := reflect.ValueOf(infoPtr).Elem()
		require.Equal(t, sampler.Handle(), (loader.VkSampler)(info.FieldByName("sampler").UnsafePointer()))
		require.Equal(t, imageView.Handle(), (loader.VkImageView)(info.FieldByName("imageView").UnsafePointer()))
		require.Equal(t, uint64(7), info.FieldByName("imageLayout").Uint()) // VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL
	})

	template, _, err := driver.CreateDescriptorUpdateTemplate(device, core1_1.DescriptorUpdateTemplateCreateInfo{}, nil)
	require.NoError(t, err)
	require.NotNil(t, template)

	driver.UpdateDescriptorSetWithTemplateFromImage(descriptorSet, template, core1_0.DescriptorImageInfo{
		Sampler:     sampler,
		ImageView:   imageView,
		ImageLayout: core1_0.ImageLayoutTransferDstOptimal,
	})
}

func TestVulkanDescriptorTemplate_UpdateDescriptorSetFromObjectHandle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	pool := mocks.NewDummyDescriptorPool(device)
	descriptorSet := mocks.NewDummyDescriptorSet(pool, device)
	bufferView := mocks.NewDummyBufferView(device)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(device, coreLoader)

	handle := mocks.NewFakeDescriptorUpdateTemplate()

	coreLoader.EXPECT().VkCreateDescriptorUpdateTemplate(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		pCreateInfo *loader.VkDescriptorUpdateTemplateCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDescriptorTemplate *loader.VkDescriptorUpdateTemplate,
	) (common.VkResult, error) {
		*pDescriptorTemplate = handle

		return core1_0.VKSuccess, nil
	})

	coreLoader.EXPECT().VkUpdateDescriptorSetWithTemplate(
		device.Handle(),
		descriptorSet.Handle(),
		handle,
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		descriptorSet loader.VkDescriptorSet,
		template loader.VkDescriptorUpdateTemplate,
		pData unsafe.Pointer,
	) {
		info := (loader.VkBufferView)(pData)
		require.Equal(t, bufferView.Handle(), info)
	})

	template, _, err := driver.CreateDescriptorUpdateTemplate(device, core1_1.DescriptorUpdateTemplateCreateInfo{}, nil)
	require.NoError(t, err)
	require.NotNil(t, template)

	driver.UpdateDescriptorSetWithTemplateFromObjectHandle(descriptorSet, template, loader.VulkanHandle(bufferView.Handle()))
}
