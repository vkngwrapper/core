package driver

import (
	"fmt"
	"sync"
)

const (
	Core1_0              string = "core1_0"
	Core1_1              string = "core1_1"
	Core1_2              string = "core1_2"
	Core1_3              string = "core1_3"
	Core1_1InstanceScope string = "core1_1instance"
	Core1_2InstanceScope string = "core1_2instance"
	Core1_3InstanceScope string = "core1_3instance"
)

type VulkanObjectStore struct {
	objStore     map[VulkanHandle]map[string]any
	objChildren  map[VulkanHandle]map[VulkanHandle]struct{}
	objParents   map[VulkanHandle]VulkanHandle
	controlMutex sync.Mutex
}

func NewObjectStore() *VulkanObjectStore {
	return &VulkanObjectStore{
		objStore:    make(map[VulkanHandle]map[string]any),
		objChildren: make(map[VulkanHandle]map[VulkanHandle]struct{}),
		objParents:  make(map[VulkanHandle]VulkanHandle),
	}
}

func (s *VulkanObjectStore) GetOrCreate(handle VulkanHandle, key string, create func() any) any {
	s.controlMutex.Lock()
	defer s.controlMutex.Unlock()

	var obj any
	objMap, hasObj := s.objStore[handle]
	if !hasObj {
		objMap = make(map[string]any)
		s.objStore[handle] = objMap
	} else {
		obj, hasObj = objMap[key]
		if hasObj {
			return obj
		}
	}

	obj = create()
	objMap[key] = obj

	return obj
}

func (s *VulkanObjectStore) SetParent(parent VulkanHandle, child VulkanHandle) {
	s.controlMutex.Lock()
	defer s.controlMutex.Unlock()

	_, parentExists := s.objStore[parent]
	_, childExists := s.objStore[child]
	if !parentExists || !childExists {
		return
	}

	children := s.objChildren[parent]
	if children == nil {
		children = make(map[VulkanHandle]struct{})
		s.objChildren[parent] = children
	}

	children[child] = struct{}{}
	s.objParents[child] = parent
}

func (s *VulkanObjectStore) deleteSingle(handle VulkanHandle) {
	delete(s.objStore, handle)

	parent := s.objParents[handle]
	parentChildren := s.objChildren[parent]
	delete(parentChildren, handle)
	delete(s.objParents, handle)

	thisChildren := s.objChildren[handle]
	var childrenToDelete = make([]VulkanHandle, 0, len(thisChildren))
	for key := range thisChildren {
		childrenToDelete = append(childrenToDelete, key)
	}

	for _, child := range childrenToDelete {
		s.deleteSingle(child)
	}
}

func (s *VulkanObjectStore) Delete(handle VulkanHandle) {
	s.controlMutex.Lock()
	defer s.controlMutex.Unlock()

	_, hasObj := s.objStore[handle]
	if !hasObj {
		return
	}

	s.deleteSingle(handle)
}

func (s *VulkanObjectStore) PrintDebug() {
	s.controlMutex.Lock()
	defer s.controlMutex.Unlock()

	if len(s.objStore) == 0 {
		return
	}

	fmt.Println("THE FOLLOWING VULKAN OBJECTS REMAIN LIVE:")

	for key, value := range s.objStore {
		fmt.Printf("%T - %v: %v+\n", value, key, value)
	}
}
