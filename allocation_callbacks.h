extern void *goAllocationCallback(
	void *pUserData,
	size_t size,
	size_t alignment,
	VkSystemAllocationScope allocationScope);

extern void *goReallocationCallback(
    void *pUserData,
    void *pOriginal,
    size_t size,
    size_t alignment,
    VkSystemAllocationScope allocationScope);

extern void goFreeCallback(
    void *pUserData,
    void *pMemory);

extern void goInternalAllocationCallback(
    void *pUserData,
    size_t size,
    VkInternalAllocationType allocationType,
    VkSystemAllocationScope allocationScope);

extern void goInternalFreeCallback(
    void *pUserData,
    size_t size,
    VkInternalAllocationType allocationType,
    VkSystemAllocationScope allocationScope);