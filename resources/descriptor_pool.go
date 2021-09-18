package resources

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/cgoparam"
	"strings"
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
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := DescriptorPoolFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := descriptorPoolFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteRune('|')
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type PoolSize struct {
	Type  core.DescriptorType
	Count int
}

type DescriptorPoolOptions struct {
	Flags DescriptorPoolFlags

	MaxSets   int
	PoolSizes []PoolSize

	core.HaveNext
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
			poolsSlice[i].descriptorCount = C.uint32_t(o.PoolSizes[i].Count)
		}

		createInfo.pPoolSizes = poolsPtr
	}

	return unsafe.Pointer(createInfo), nil
}

type vulkanDescriptorPool struct {
	loader loader.Loader
	handle loader.VkDescriptorPool
	device loader.VkDevice
}

func (p *vulkanDescriptorPool) Handle() loader.VkDescriptorPool {
	return p.handle
}

func (p *vulkanDescriptorPool) Destroy() error {
	return p.loader.VkDestroyDescriptorPool(p.device, p.handle, nil)
}

func (p *vulkanDescriptorPool) FreeDescriptorSets(sets []DescriptorSet) (loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	setCount := len(sets)
	descriptorsPtrUnsafe := arena.Malloc(setCount * int(unsafe.Sizeof([1]C.VkDescriptorSet{})))
	descriptorSlice := ([]C.VkDescriptorSet)(unsafe.Slice((*C.VkDescriptorSet)(descriptorsPtrUnsafe), setCount))

	for i := 0; i < setCount; i++ {
		descriptorSlice[i] = (C.VkDescriptorSet)(unsafe.Pointer(sets[i].Handle()))
	}

	return p.loader.VkFreeDescriptorSets(p.device, p.handle, loader.Uint32(setCount), (*loader.VkDescriptorSet)(descriptorsPtrUnsafe))
}
