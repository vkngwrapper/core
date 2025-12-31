# vkngwrapper/core/v3

[![Go Reference](https://pkg.go.dev/badge/github.com/vkngwrapper/core/v3.svg)](https://pkg.go.dev/github.com/vkngwrapper/core/v3)

`go get github.com/vkngwrapper/core/v3`

Vkngwrapper (proununced "Viking Wrapper") is a handwritten cgo wrapper for the Vulkan graphics and compute API.
 The goal is to produce fast, easy-to-use, low-go-allocation, and idiomatic Go code to communicate with your graphics
 card and enable games and other graphical applications. Vkngwrapper currently supports core versions 1.0-1.2,
 as well as many extensions via the https://github.com/vkngwrapper/extensions repository.

Under the hood, Vkngwrapper uses https://github.com/cannibalvox/cgoparam to avoid calling `C.Malloc` and 
 `C.Free` while still avoiding the cost of a deep cgocheck on Go memory. This allows you to save precious
 nanoseconds (or sometimes microseconds!) on your cgo overhead.

Vkngwrapper is also heavily-tested. The marshalling and unmarshalling layer has high test coverage, giving the
 core library 86.0% test coverage and the extensions library 84.9% test coverage. While this coverage is not
 perfect, Vulkan has an extremely large API surface, and these tests ensure that there is no obviously-busted
 functionality. Additionally, the entire API is mockable (and pre-generated gomocks are provided), allowing you 
 to test your own code with ease.

Lastly, vkngwrapper has a solid and still-growing base of examples, built from Go ports of existing Vulkan
 examples.  Several key samples from https://github.com/LunarG/VulkanSamples have are included in
 [our example repository](https://github.com/vkngwrapper/examples), as well as a full port of 
 [the Vulkan tutorial](https://vulkan-tutorial.com), which can be followed step by step at
 https://github.com/vkngwrapper/vulkan-tutorial

For more information about our future roadmap, see [the org page](https://github.com/vkngwrapper).

## Getting Started

Before building any Vulkan application, you will need to install [the Vulkan SDK](https://www.lunarg.com/vulkan-sdk/)
 for your operating system. Additionally, if you intend to use SDL2 to create windows, as in vkngwrapper's examples,
 it may be necessary to download SDL2 using your local package manager. For more information, 
 see [go-sdl2 requirements](https://github.com/veandco/go-sdl2#requirements).

The first step to using vkngwrapper is to create a [GlobalDriver](https://pkg.go.dev/github.com/vkngwrapper/core/v3/core1_0#GlobalDriver).
 While we offer the option to create a GlobalDriver from a ProcAddr provided by a windowing system (such as SDL2),
 the easiest way is to build a loader from the system's local Vulkan library:

```go 
globalDriver, err := core.CreateSystemLoader()
if err != nil {
 return err 
}
```
 
Once you have a Loader, you can use that Loader to create a [CoreInstanceDriver](https://pkg.go.dev/github.com/vkngwrapper/core/v3/core1_0#CoreInstanceDriver), and the CoreInstanceDriver to
 create a [PhysicalDevice](https://pkg.go.dev/github.com/vkngwrapper/core/v3/core1_0#PhysicalDevice) and a [CoreDeviceDriver](https://pkg.go.dev/github.com/vkngwrapper/core/v3/core1_0#CoreDeviceDriver).

```go
instanceOptions := core1_0.InstanceCreateInfo{
    ApplicationName:    "My Vulkan App",
    ApplicationVersion: common.CreateVersion(1, 0, 0),
    EngineName:         "No Engine",
    EngineVersion:      common.CreateVersion(1, 0, 0),
    APIVersion:         common.Vulkan1_0,
}

instanceDriver, _, err := globalDriver.CreateInstance(nil, instanceOptions)
if err != nil {
	return err 
}

physicalDevices, _, err := instanceDriver.EnumeratePhysicalDevices()
if err != nil {
    return err
}

// The real logic is more complicated than this
queueFamilies := instanceDriver.GetPhysicalDeviceQueueFamilyProperties(physicalDevices[0])
queueIndex := -1

for index, queueFamily := range queueFamilies {
	if (queueFamily.QueueFlags & core1_0.QueueGraphics) != 0 {
        graphicsIndex = index 		
    }
}

deviceOptions := core1_0.DeviceCreateInfo{
	QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
	    {
		    QueueFamilyIndex: 	graphicsIndex,
			QueuePriorities: []float32{1.0},
        },
    },
}

deviceDriver, _, err := instanceDriver.CreateDevice(physicalDevices[0], nil, deviceOptions)
if err != nil {
	return err 
}
```

Then, the world is your oyster! Be sure to destroy these (and all other) Vulkan objects when you are finished with them.
 To learn more about how to use vkngwrapper effectively, check out the
 [examples repository](https://github.com/vkngwrapper/examples) and to learn more about how to use 
 Vulkan effectively, check out [the Vulkan tutorial](https://vulkan-tutorial.com) and the excellent [Vulkan
 Discord](https://discord.com/invite/vulkan)!

## Principals of vkngwrapper

While vkngwrapper labors to follow the Vulkan specification fairly closely, there are some unusual qualities that one should
 be aware of when working with the library.

### Use Idiomatic Types

When representing integer numbers, most types in vkngwrapper are simply `int`, while the underlying Vulkan
 type may be `uint64`, `int32`, etc. The only exception is when a type represents a bitmask. Likewise,
 while a duration in Vulkan might be represented by an integer counting nanoseconds, vkngwrapper tends to
 use `time.Duration`. This library endeavors to use go-friendly types unless doing so would result in a degradation
 of quality or performance for a substantial number of users. 

### Namespace By Availability

All types, methods, and constants in vkngwrapper (both here in the core library, as well as the [extensions library](https://github.com/vkngwrapper/extensions))
 are packaged under the Vulkan version or extension that makes them available for use. For instance, SamplerYcbcrConversion objects
 were introduced in the [VK_KHR_sampler_ycbcr_conversion](https://pkg.go.dev/github.com/vkngwrapper/extensions/v3/khr_sampler_ycbcr_conversion)
 extension, and then later promoted to [core 1.1](https://pkg.go.dev/github.com/vkngwrapper/core/v3/core1_1). As a result, 
 the SamplerYcbcrConversion object can be created via [khr_sampler_ycbcr_conversion.ExtensionDriver](https://pkg.go.dev/github.com/vkngwrapper/extensions/v3/khr_sampler_ycbcr_conversion#ExtensionDriver)
 and [core1_1.CoreDeviceDriver](https://pkg.go.dev/github.com/vkngwrapper/core/v3/core1_1#CoreDeviceDriver).

All symbols that are available in the C Vulkan headers are namespaced in this manner, with the exception of 
 [driver.AllocationCallbacks](https://pkg.go.dev/github.com/vkngwrapper/core/v3/driver#AllocationCallbacks) which
 is special for silly package interdependency and cgo reasons. Arguments that accept `*driver.AllocationCallbacks` can
 usually be left nil, but if you would like to receive callbacks when Vulkan makes internal allocations and deallocations,
 do the following:

1. Create a [common.AllocationCallbackOptions](https://pkg.go.dev/github.com/vkngwrapper/core/v3/common#AllocationCallbackOptions)
   object with the callback methods you would like to be executed, and optionally, a UserData object to be passed to all
   callbacks.
2. Use [driver.CreateAllocationCallbacks](https://pkg.go.dev/github.com/vkngwrapper/core/v3/driver#CreateAllocationCallbacks)
   to create a `driver.AllocationCallbacks` object, which can be passed to Create, Destroy, and Free methods.

While `driver.AllocationCallbacks` objects are immutable, `common.AllocationCallbackOptions` structures are not. They
 can be modified and then used to create another `driver.AllocationCallbacks` object with different behaviors. 
 `driver.AllocationCallbacks` objects need to be destroyed like any other Vulkan object when you are done with them.

### Advertise Version Support

All Vulkan objects in vkngwrapper have an `APIVersion` method which returns the highest Vulkan core version the object
 supports. Generally speaking, the `GlobalDriver` will support whatever version is available via the .dll/.so/etc. the GlobalDriver
 was created from, the `Instance` will support whatever version you requested when creating it, if lower than the
 Loader version, the `PhysicalDevice` will support whatever version your hardware supports, if lower than the `Instance`
 version, and all other objects will inherit their version from the `PhysicalDevice` they exist on.

It is helpful to be able to request information about Vulkan support from any Vulkan object, but the easiest way to
 check for core version support is with type queries.

### Promote to Add Functionality

All Vulkan versions from 1.1 upward provide expanded versions of Vulkan drivers introduced in previous core versions.
 [core1_1.CoreInstanceDriver](https://pkg.go.dev/github.com/vkngwrapper/core/v3/core1_1#CoreInstanceDriver) introduces
 several new Vulkan commands introduced in core 1.1.  In environments where you are making use of core 1.1 functionality,
 you may find it easier to work with `core1_1.CoreInstanceDriver`.  You can check for core 1.1 support in the instance 
 driver by simply using a type query: `instanceDriver11, isSupported := instanceDriver.(core1_1.CoreInstanceDriver)`.  If the
 type query is successful, then core 1.1 functionality is supported.  This can also be done with the `CoreDeviceDriver`.

### Chain Options and OutData

Vulkan has the capability to allow existing structure and method behavior to be extended by chaining
 structures using a `pNext` field added to most Vulkan structures. This field is represented in vkngwrapper
 using the [NextOptions](https://pkg.go.dev/github.com/vkngwrapper/core/v3/common#NextOptions) and 
 [NextOutData](https://pkg.go.dev/github.com/vkngwrapper/core/v3/common#NextOutData) embedded structures.

Take a look at this example:

```go
_, err := device.BindBufferMemory2(
    core1_1.BindBufferMemoryInfo {
        Buffer:       buffer,
        Memory:       memory,
        MemoryOffset: 1,

        NextOptions: common.NextOptions{
            core1_1.BindBufferMemoryDeviceGroupInfo{
                DeviceIndices: []int{1, 2, 7},
            },
        },
    },
)
```

By chaining `core1_1.BindBufferMemoryDeviceGroupInfo` onto `core1_1.BindBufferMemoryInfo`, additional
 behavior related to Device groups can be applied to an existing method. `BindBufferMemoryDeviceGroupInfo`
 also has a `NextOptions` embedded struct, so further behavior can be chained to that structure as well.

Broadly speaking, any structure that passes data into a Vulkan command embeds `NextOptions` and is passed
 in by value. Any structure that retrieves data from a Vulkan command embeds `NextOutData`
 and is passed in as a pointer. Chaining Options allows you to pass in additional parameter data to a command
 and change the behavior of a command. Chaining OutData allows you to request additional data from a command,
 which will be populated into the chained OutData.

While Vulkan has specific Options types that are intended to go together (and more can be learned as
 you understand Vulkan more deeply), from a syntactical point of view, any structure with `NextOptions`
 can be chained onto any other structure with `NextOptions`.  Likewise, any structure with `NextOutData`
 can be chained onto any other structure with `NextOutData`.

Some structures (mainly Features structures) have both `NextOptions` and `NextOutData`.  When they are being
 used to pass data into Vulkan (such as in [core1_0.CoreInstanceDriver.CreateDevice](https://pkg.go.dev/github.com/vkngwrapper/core/v3/core1_0#CoreInstanceDriver),
 when it is specifying which features to activate), you must use `NextOptions` to chain further structures.
 When they are being used to retrieve data from Vulkan (such as in 
 [core1_1.CoreInstanceDriver.GetPhysicalDeviceFeatures2](https://pkg.go.dev/github.com/vkngwrapper/core/v3/core1_1#CoreInstanceDriver),
 when it is retrieving feature support from the device), you must use `NextOutData` to chain further structures.

Chained structures in the wrong field will be ignored.
