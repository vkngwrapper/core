package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/cockroachdb/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func TestDevice_GetQueue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	queueHandle := mocks.NewFakeQueue()

	mockDriver.EXPECT().VkGetDeviceQueue(device.Handle(), core.Uint32(1), core.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, queueFamilyIndex, queueIndex core.Uint32, pQueue *core.VkQueue) error {
			*pQueue = queueHandle
			return nil
		})

	queue, err := device.GetQueue(1, 2)
	require.NoError(t, err)
	require.NotNil(t, queue)
	require.Equal(t, queueHandle, queue.Handle())
}

func TestDevice_WaitForFences_Timeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	fence1 := mocks.EasyDummyFence(t, loader, device)
	fence2 := mocks.EasyDummyFence(t, loader, device)

	mockDriver.EXPECT().VkWaitForFences(device.Handle(), core.Uint32(2), gomock.Not(nil), core.VkBool32(1), core.Uint64(1)).DoAndReturn(
		func(device core.VkDevice, fenceCount core.Uint32, pFences *core.VkFence, waitAll core.VkBool32, timeout core.Uint64) (core.VkResult, error) {
			fenceSlice := ([]core.VkFence)(unsafe.Slice(pFences, 2))
			require.ElementsMatch(t, []core.VkFence{fence1.Handle(), fence2.Handle()}, fenceSlice)

			return core.VKSuccess, nil
		})

	_, err = device.WaitForFences(true, time.Nanosecond, []core.Fence{
		fence1, fence2,
	})
	require.NoError(t, err)
}

func TestDevice_WaitForFences_NoTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	fence1 := mocks.EasyDummyFence(t, loader, device)

	mockDriver.EXPECT().VkWaitForFences(device.Handle(), core.Uint32(1), gomock.Not(nil), core.VkBool32(0), core.Uint64(0xffffffffffffffff)).DoAndReturn(
		func(device core.VkDevice, fenceCount core.Uint32, pFences *core.VkFence, waitAll core.VkBool32, timeout core.Uint64) (core.VkResult, error) {
			fenceSlice := ([]core.VkFence)(unsafe.Slice(pFences, 1))
			require.ElementsMatch(t, []core.VkFence{fence1.Handle()}, fenceSlice)

			return core.VKSuccess, nil
		})

	_, err = device.WaitForFences(false, common.NoTimeout, []core.Fence{
		fence1,
	})
	require.NoError(t, err)
}

func TestDevice_WaitForIdle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)

	mockDriver.EXPECT().VkDeviceWaitIdle(device.Handle()).Return(core.VKSuccess, nil)
	_, err = device.WaitForIdle()
	require.NoError(t, err)
}

func TestVulkanDevice_ResetFences(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	fence1 := mocks.EasyDummyFence(t, loader, device)
	fence2 := mocks.EasyDummyFence(t, loader, device)

	mockDriver.EXPECT().VkResetFences(device.Handle(), core.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, fenceCount core.Uint32, pFence *core.VkFence) (core.VkResult, error) {
			fences := ([]core.VkFence)(unsafe.Slice(pFence, 2))

			require.ElementsMatch(t, []core.VkFence{fence1.Handle(), fence2.Handle()}, fences)
			return core.VKSuccess, nil
		})

	_, err = device.ResetFences([]core.Fence{fence1, fence2})
	require.NoError(t, err)
}

func TestVulkanDevice_AllocateAndFreeMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	memoryHandle := mocks.NewFakeDeviceMemoryHandle()

	mockDriver.EXPECT().VkAllocateMemory(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkMemoryAllocateInfo, pAllocator *core.VkAllocationCallbacks, pMemory *core.VkDeviceMemory) (core.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(5), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(7), val.FieldByName("allocationSize").Uint())
			require.Equal(t, uint64(3), val.FieldByName("memoryTypeIndex").Uint())

			*pMemory = memoryHandle
			return core.VKSuccess, nil
		})
	mockDriver.EXPECT().VkFreeMemory(device.Handle(), memoryHandle, nil).Return(nil)

	memory, _, err := device.AllocateMemory(&core.DeviceMemoryOptions{
		AllocationSize:  7,
		MemoryTypeIndex: 3,
	})
	require.NoError(t, err)
	require.NotNil(t, memory)
	require.Equal(t, memoryHandle, memory.Handle())

	err = device.FreeMemory(memory)
	require.NoError(t, err)
}

func TestVulkanDevice_UpdateDescriptorSets_WriteImageInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)
	sampler1 := mocks.EasyMockSampler(ctrl)
	sampler2 := mocks.EasyMockSampler(ctrl)
	imageView1 := mocks.EasyMockImageView(ctrl)
	imageView2 := mocks.EasyMockImageView(ctrl)

	mockDriver.EXPECT().VkUpdateDescriptorSets(device.Handle(), core.Uint32(1), gomock.Not(nil), core.Uint32(0), nil).DoAndReturn(
		func(device core.VkDevice, descriptorWriteCount core.Uint32, pDescriptorWrites *core.VkWriteDescriptorSet, descriptorCopyCount core.Uint32, pDescriptorCopies *core.VkCopyDescriptorSet) error {
			writeSlice := ([]core.VkWriteDescriptorSet)(unsafe.Slice(pDescriptorWrites, 1))
			writeVal := reflect.ValueOf(writeSlice[0])

			require.Equal(t, uint64(35), writeVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
			require.True(t, writeVal.FieldByName("pNext").IsNil())
			require.Equal(t, destDescriptor.Handle(), core.VkDescriptorSet(unsafe.Pointer(writeVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), writeVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(6), writeVal.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
			require.True(t, writeVal.FieldByName("pBufferInfo").IsNil())
			require.True(t, writeVal.FieldByName("pTexelBufferView").IsNil())

			imageInfoPtr := (*core.VkDescriptorImageInfo)(unsafe.Pointer(writeVal.FieldByName("pImageInfo").Elem().UnsafeAddr()))
			imageInfoSlice := ([]core.VkDescriptorImageInfo)(unsafe.Slice(imageInfoPtr, 2))

			require.Len(t, imageInfoSlice, 2)

			imageInfoVal := reflect.ValueOf(imageInfoSlice[0])
			require.Equal(t, sampler1.Handle(), (core.VkSampler)(unsafe.Pointer(imageInfoVal.FieldByName("sampler").Elem().UnsafeAddr())))
			require.Equal(t, imageView1.Handle(), (core.VkImageView)(unsafe.Pointer(imageInfoVal.FieldByName("imageView").Elem().UnsafeAddr())))
			require.Equal(t, uint64(3), imageInfoVal.FieldByName("imageLayout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL

			imageInfoVal = reflect.ValueOf(imageInfoSlice[1])
			require.Equal(t, sampler2.Handle(), (core.VkSampler)(unsafe.Pointer(imageInfoVal.FieldByName("sampler").Elem().UnsafeAddr())))
			require.Equal(t, imageView2.Handle(), (core.VkImageView)(unsafe.Pointer(imageInfoVal.FieldByName("imageView").Elem().UnsafeAddr())))
			require.Equal(t, uint64(8), imageInfoVal.FieldByName("imageLayout").Uint()) // VK_IMAGE_LAYOUT_PREINITIALIZED

			return nil
		})

	err = device.UpdateDescriptorSets([]core.WriteDescriptorSetOptions{
		{
			Destination:             destDescriptor,
			DestinationBinding:      1,
			DestinationArrayElement: 2,
			DescriptorType:          common.DescriptorUniformBuffer,
			ImageInfo: []core.DescriptorImageInfo{
				{
					Sampler:     sampler1,
					ImageView:   imageView1,
					ImageLayout: common.LayoutDepthStencilAttachmentOptimal,
				},
				{
					Sampler:     sampler2,
					ImageView:   imageView2,
					ImageLayout: common.LayoutPreInitialized,
				},
			},
		},
	}, nil)

	require.NoError(t, err)
}
func TestVulkanDevice_UpdateDescriptorSets_WriteBufferInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)
	buffer1 := mocks.EasyMockBuffer(ctrl)
	buffer2 := mocks.EasyMockBuffer(ctrl)

	mockDriver.EXPECT().VkUpdateDescriptorSets(device.Handle(), core.Uint32(1), gomock.Not(nil), core.Uint32(0), nil).DoAndReturn(
		func(device core.VkDevice, descriptorWriteCount core.Uint32, pDescriptorWrites *core.VkWriteDescriptorSet, descriptorCopyCount core.Uint32, pDescriptorCopies *core.VkCopyDescriptorSet) error {
			writeSlice := ([]core.VkWriteDescriptorSet)(unsafe.Slice(pDescriptorWrites, 1))
			writeVal := reflect.ValueOf(writeSlice[0])

			require.Equal(t, uint64(35), writeVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
			require.True(t, writeVal.FieldByName("pNext").IsNil())
			require.Equal(t, destDescriptor.Handle(), core.VkDescriptorSet(unsafe.Pointer(writeVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), writeVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(3), writeVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(6), writeVal.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
			require.True(t, writeVal.FieldByName("pImageInfo").IsNil())
			require.True(t, writeVal.FieldByName("pTexelBufferView").IsNil())

			bufferInfoPtr := (*core.VkDescriptorBufferInfo)(unsafe.Pointer(writeVal.FieldByName("pBufferInfo").Elem().UnsafeAddr()))
			bufferInfoSlice := ([]core.VkDescriptorBufferInfo)(unsafe.Slice(bufferInfoPtr, 2))

			require.Len(t, bufferInfoSlice, 2)

			bufferInfoVal := reflect.ValueOf(bufferInfoSlice[0])
			require.Equal(t, buffer1.Handle(), (core.VkBuffer)(unsafe.Pointer(bufferInfoVal.FieldByName("buffer").Elem().UnsafeAddr())))
			require.Equal(t, uint64(7), bufferInfoVal.FieldByName("offset").Uint())
			require.Equal(t, uint64(11), bufferInfoVal.FieldByName("_range").Uint())

			bufferInfoVal = reflect.ValueOf(bufferInfoSlice[1])
			require.Equal(t, buffer2.Handle(), (core.VkBuffer)(unsafe.Pointer(bufferInfoVal.FieldByName("buffer").Elem().UnsafeAddr())))
			require.Equal(t, uint64(13), bufferInfoVal.FieldByName("offset").Uint())
			require.Equal(t, uint64(17), bufferInfoVal.FieldByName("_range").Uint())

			return nil
		})

	err = device.UpdateDescriptorSets([]core.WriteDescriptorSetOptions{
		{
			Destination:             destDescriptor,
			DestinationBinding:      1,
			DestinationArrayElement: 3,
			DescriptorType:          common.DescriptorUniformBuffer,
			BufferInfo: []core.DescriptorBufferInfo{
				{
					Buffer: buffer1,
					Offset: 7,
					Range:  11,
				},
				{
					Buffer: buffer2,
					Offset: 13,
					Range:  17,
				},
			},
		},
	}, nil)

	require.NoError(t, err)
}

func TestVulkanDevice_UpdateDescriptorSets_TexelBufferView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)
	bufferView1 := mocks.EasyMockBufferView(ctrl)
	bufferView2 := mocks.EasyMockBufferView(ctrl)

	mockDriver.EXPECT().VkUpdateDescriptorSets(device.Handle(), core.Uint32(1), gomock.Not(nil), core.Uint32(0), nil).DoAndReturn(
		func(device core.VkDevice, descriptorWriteCount core.Uint32, pDescriptorWrites *core.VkWriteDescriptorSet, descriptorCopyCount core.Uint32, pDescriptorCopies *core.VkCopyDescriptorSet) error {
			writeSlice := ([]core.VkWriteDescriptorSet)(unsafe.Slice(pDescriptorWrites, 1))
			writeVal := reflect.ValueOf(writeSlice[0])

			require.Equal(t, uint64(35), writeVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
			require.True(t, writeVal.FieldByName("pNext").IsNil())
			require.Equal(t, destDescriptor.Handle(), core.VkDescriptorSet(unsafe.Pointer(writeVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), writeVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(3), writeVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(6), writeVal.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
			require.True(t, writeVal.FieldByName("pImageInfo").IsNil())
			require.True(t, writeVal.FieldByName("pBufferInfo").IsNil())

			bufferInfoPtr := (*core.VkBufferView)(unsafe.Pointer(writeVal.FieldByName("pTexelBufferView").Elem().UnsafeAddr()))
			bufferInfoSlice := ([]core.VkBufferView)(unsafe.Slice(bufferInfoPtr, 2))

			require.Len(t, bufferInfoSlice, 2)

			require.Equal(t, bufferView1.Handle(), (core.VkBufferView)(unsafe.Pointer(bufferInfoSlice[0])))
			require.Equal(t, bufferView2.Handle(), (core.VkBufferView)(unsafe.Pointer(bufferInfoSlice[1])))

			return nil
		})

	err = device.UpdateDescriptorSets([]core.WriteDescriptorSetOptions{
		{
			Destination:             destDescriptor,
			DestinationBinding:      1,
			DestinationArrayElement: 3,
			DescriptorType:          common.DescriptorUniformBuffer,
			TexelBufferView: []core.BufferView{
				bufferView1, bufferView2,
			},
		},
	}, nil)

	require.NoError(t, err)
}
func TestVulkanDevice_UpdateDescriptorSets_Copy(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	srcDescriptor := mocks.EasyMockDescriptorSet(ctrl)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)

	mockDriver.EXPECT().VkUpdateDescriptorSets(device.Handle(), core.Uint32(0), nil, core.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, descriptorWriteCount core.Uint32, pDescriptorWrites *core.VkWriteDescriptorSet, descriptorCopyCount core.Uint32, pDescriptorCopies *core.VkCopyDescriptorSet) error {
			copySlice := ([]core.VkCopyDescriptorSet)(unsafe.Slice(pDescriptorCopies, 1))
			copyVal := reflect.ValueOf(copySlice[0])

			require.Equal(t, uint64(36), copyVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COPY_DESCRIPTOR_SET
			require.True(t, copyVal.FieldByName("pNext").IsNil())
			require.Equal(t, srcDescriptor.Handle(), (core.VkDescriptorSet)(unsafe.Pointer(copyVal.FieldByName("srcSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(3), copyVal.FieldByName("srcBinding").Uint())
			require.Equal(t, uint64(5), copyVal.FieldByName("srcArrayElement").Uint())
			require.Equal(t, destDescriptor.Handle(), (core.VkDescriptorSet)(unsafe.Pointer(copyVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(7), copyVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(11), copyVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(13), copyVal.FieldByName("descriptorCount").Uint())

			return nil
		})

	err = device.UpdateDescriptorSets(nil, []core.CopyDescriptorSetOptions{
		{
			Source:             srcDescriptor,
			SourceBinding:      3,
			SourceArrayElement: 5,

			Destination:             destDescriptor,
			DestinationBinding:      7,
			DestinationArrayElement: 11,

			Count: 13,
		},
	})

	require.NoError(t, err)
}

func TestVulkanDevice_UpdateDescriptorSets_FailureImageInfoAndBufferInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)
	buffer1 := mocks.EasyMockBuffer(ctrl)
	buffer2 := mocks.EasyMockBuffer(ctrl)
	sampler1 := mocks.EasyMockSampler(ctrl)
	sampler2 := mocks.EasyMockSampler(ctrl)
	imageView1 := mocks.EasyMockImageView(ctrl)
	imageView2 := mocks.EasyMockImageView(ctrl)

	err = device.UpdateDescriptorSets([]core.WriteDescriptorSetOptions{
		{
			Destination:             destDescriptor,
			DestinationBinding:      1,
			DestinationArrayElement: 3,
			DescriptorType:          common.DescriptorUniformBuffer,
			BufferInfo: []core.DescriptorBufferInfo{
				{
					Buffer: buffer1,
					Offset: 7,
					Range:  11,
				},
				{
					Buffer: buffer2,
					Offset: 13,
					Range:  17,
				},
			},
			ImageInfo: []core.DescriptorImageInfo{
				{
					Sampler:     sampler1,
					ImageView:   imageView1,
					ImageLayout: common.LayoutDepthStencilAttachmentOptimal,
				},
				{
					Sampler:     sampler2,
					ImageView:   imageView2,
					ImageLayout: common.LayoutPreInitialized,
				},
			},
		},
	}, nil)

	require.Error(t, errors.New("a WriteDescriptorSetOptions may have one or more ImageInfo sources OR one or more BufferInfo sources, but not both"), err)
}

func TestVulkanDevice_UpdateDescriptorSets_FailureImageInfoAndBufferView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)
	bufferView1 := mocks.EasyMockBufferView(ctrl)
	bufferView2 := mocks.EasyMockBufferView(ctrl)
	sampler1 := mocks.EasyMockSampler(ctrl)
	sampler2 := mocks.EasyMockSampler(ctrl)
	imageView1 := mocks.EasyMockImageView(ctrl)
	imageView2 := mocks.EasyMockImageView(ctrl)

	err = device.UpdateDescriptorSets([]core.WriteDescriptorSetOptions{
		{
			Destination:             destDescriptor,
			DestinationBinding:      1,
			DestinationArrayElement: 3,
			DescriptorType:          common.DescriptorUniformBuffer,
			ImageInfo: []core.DescriptorImageInfo{
				{
					Sampler:     sampler1,
					ImageView:   imageView1,
					ImageLayout: common.LayoutDepthStencilAttachmentOptimal,
				},
				{
					Sampler:     sampler2,
					ImageView:   imageView2,
					ImageLayout: common.LayoutPreInitialized,
				},
			},
			TexelBufferView: []core.BufferView{
				bufferView1, bufferView2,
			},
		},
	}, nil)

	require.Error(t, errors.New("a WriteDescriptorSetOptions may have one or more ImageInfo sources OR one or more TexelBufferView sources, but not both"), err)
}

func TestVulkanDevice_UpdateDescriptorSets_FailureBufferInfoAndBufferView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)
	buffer1 := mocks.EasyMockBuffer(ctrl)
	buffer2 := mocks.EasyMockBuffer(ctrl)
	bufferView1 := mocks.EasyMockBufferView(ctrl)
	bufferView2 := mocks.EasyMockBufferView(ctrl)

	err = device.UpdateDescriptorSets([]core.WriteDescriptorSetOptions{
		{
			Destination:             destDescriptor,
			DestinationBinding:      1,
			DestinationArrayElement: 3,
			DescriptorType:          common.DescriptorUniformBuffer,
			BufferInfo: []core.DescriptorBufferInfo{
				{
					Buffer: buffer1,
					Offset: 7,
					Range:  11,
				},
				{
					Buffer: buffer2,
					Offset: 13,
					Range:  17,
				},
			},
			TexelBufferView: []core.BufferView{
				bufferView1, bufferView2,
			},
		},
	}, nil)

	require.Error(t, errors.New("a WriteDescriptorSetOptions may have one or more BufferInfo sources OR one or more TexelBufferView sources, but not both"), err)
}

func TestVulkanDevice_UpdateDescriptorSets_FailureNoSource(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)

	err = device.UpdateDescriptorSets([]core.WriteDescriptorSetOptions{
		{
			Destination:             destDescriptor,
			DestinationBinding:      1,
			DestinationArrayElement: 3,
			DescriptorType:          common.DescriptorUniformBuffer,
		},
	}, nil)

	require.Error(t, errors.New("a WriteDescriptorSetOptions must have a source to write the descriptor from: ImageInfo, BufferInfo, TexelBufferView, or an extension source"), err)
}
