package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	driver3 "github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type DescriptorPoolFlags int32

const (
	DescriptorPoolFreeDescriptorSet DescriptorPoolFlags = C.VK_DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT
	DescriptorPoolUpdateAfterBind   DescriptorPoolFlags = C.VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT
	DescriptorPoolHostOnlyValve     DescriptorPoolFlags = C.VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_VALVE
)

var descriptorPoolFlagsToString = map[DescriptorPoolFlags]string{
	DescriptorPoolFreeDescriptorSet: "Free Descriptor Set",
	DescriptorPoolUpdateAfterBind:   "Update After Bind",
	DescriptorPoolHostOnlyValve:     "Host-Only (Valve)",
}

func (f DescriptorPoolFlags) String() string {
	return common.FlagsToString(f, descriptorPoolFlagsToString)
}

type PoolSize struct {
	Type            common.DescriptorType
	DescriptorCount int
}

type DescriptorPoolOptions struct {
	Flags DescriptorPoolFlags

	MaxSets   int
	PoolSizes []PoolSize

	common.HaveNext
}

func (o *DescriptorPoolOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkDescriptorPoolCreateInfo)(allocator.Malloc(C.sizeof_struct_VkDescriptorPoolCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO
	createInfo.flags = C.VkDescriptorPoolCreateFlags(o.Flags)
	createInfo.pNext = next

	createInfo.maxSets = C.uint32_t(o.MaxSets)
	sizeCount := len(o.PoolSizes)
	createInfo.poolSizeCount = C.uint32_t(sizeCount)
	createInfo.pPoolSizes = nil

	if sizeCount > 0 {
		poolsPtr := (*C.VkDescriptorPoolSize)(allocator.Malloc(sizeCount * C.sizeof_struct_VkDescriptorPoolSize))
		poolsSlice := ([]C.VkDescriptorPoolSize)(unsafe.Slice(poolsPtr, sizeCount))

		for i := 0; i < sizeCount; i++ {
			poolsSlice[i]._type = C.VkDescriptorType(o.PoolSizes[i].Type)
			poolsSlice[i].descriptorCount = C.uint32_t(o.PoolSizes[i].DescriptorCount)
		}

		createInfo.pPoolSizes = poolsPtr
	}

	return unsafe.Pointer(createInfo), nil
}

type vulkanDescriptorPool struct {
	driver driver3.Driver
	handle driver3.VkDescriptorPool
	device driver3.VkDevice
}

func (p *vulkanDescriptorPool) Handle() driver3.VkDescriptorPool {
	return p.handle
}

func (p *vulkanDescriptorPool) Destroy(callbacks *AllocationCallbacks) {
	p.driver.VkDestroyDescriptorPool(p.device, p.handle, callbacks.Handle())
}

func (p *vulkanDescriptorPool) AllocateDescriptorSets(o *DescriptorSetOptions) ([]DescriptorSet, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	o.descriptorPool = p

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	setCount := len(o.AllocationLayouts)
	descriptorSets := (*driver3.VkDescriptorSet)(arena.Malloc(setCount * int(unsafe.Sizeof([1]C.VkDescriptorSet{}))))

	res, err := p.driver.VkAllocateDescriptorSets(p.device, (*driver3.VkDescriptorSetAllocateInfo)(createInfo), descriptorSets)
	if err != nil {
		return nil, res, err
	}

	var sets []DescriptorSet
	descriptorSetSlice := ([]driver3.VkDescriptorSet)(unsafe.Slice(descriptorSets, setCount))
	for i := 0; i < setCount; i++ {
		sets = append(sets, &vulkanDescriptorSet{handle: descriptorSetSlice[i]})
	}

	return sets, res, nil
}

func (p *vulkanDescriptorPool) FreeDescriptorSets(sets []DescriptorSet) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	setCount := len(sets)
	descriptorsPtrUnsafe := arena.Malloc(setCount * int(unsafe.Sizeof([1]C.VkDescriptorSet{})))
	descriptorSlice := ([]C.VkDescriptorSet)(unsafe.Slice((*C.VkDescriptorSet)(descriptorsPtrUnsafe), setCount))

	for i := 0; i < setCount; i++ {
		descriptorSlice[i] = (C.VkDescriptorSet)(unsafe.Pointer(sets[i].Handle()))
	}

	return p.driver.VkFreeDescriptorSets(p.device, p.handle, driver3.Uint32(setCount), (*driver3.VkDescriptorSet)(descriptorsPtrUnsafe))
}

type DescriptorPoolResetFlags int32

func (p *vulkanDescriptorPool) Reset(flags DescriptorPoolResetFlags) (common.VkResult, error) {
	return p.driver.VkResetDescriptorPool(p.device, p.handle, driver3.VkDescriptorPoolResetFlags(flags))
}
