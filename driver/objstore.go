package driver

import (
	"fmt"
	"sync"
)

type VulkanObjectStore struct {
	objStore     map[VulkanHandle]interface{}
	objChildren  map[VulkanHandle]map[VulkanHandle]struct{}
	objParents   map[VulkanHandle]VulkanHandle
	controlMutex sync.Mutex
}

func NewObjectStore() *VulkanObjectStore {
	return &VulkanObjectStore{
		objStore:    make(map[VulkanHandle]interface{}),
		objChildren: make(map[VulkanHandle]map[VulkanHandle]struct{}),
		objParents:  make(map[VulkanHandle]VulkanHandle),
	}
}

func (s *VulkanObjectStore) GetOrCreate(handle VulkanHandle, create func() interface{}) interface{} {
	s.controlMutex.Lock()
	defer s.controlMutex.Unlock()

	obj, hasObj := s.objStore[handle]
	if hasObj {
		return obj
	}

	obj = create()
	s.objStore[handle] = obj

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

func (s *VulkanObjectStore) Delete(handle VulkanHandle, obj interface{}) {
	s.controlMutex.Lock()
	defer s.controlMutex.Unlock()

	actualObj, hasObj := s.objStore[handle]
	if !hasObj || actualObj != obj {
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
