package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/core1_0"

const (
	FormatB10X6G10X6R10X6G10X6HorizontalChromaComponentPacked     core1_0.Format = C.VK_FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16
	FormatB12X4G12X4R12X4G12X4HorizontalChromaComponentPacked     core1_0.Format = C.VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16
	FormatB16G16R16G16HorizontalChroma                            core1_0.Format = C.VK_FORMAT_B16G16R16G16_422_UNORM
	FormatB8G8R8G8HorizontalChroma                                core1_0.Format = C.VK_FORMAT_B8G8R8G8_422_UNORM
	FormatG10X6B10X6G10X6R10X6HorizontalChromaComponentPacked     core1_0.Format = C.VK_FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16
	FormatG10X6_B10X6R10X6_2PlaneDualChromaComponentPacked        core1_0.Format = C.VK_FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16
	FormatG10X6_B10X6R10X6_2PlaneHorizontalChromaComponentPacked  core1_0.Format = C.VK_FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16
	FormatG10X6_B10X6_R10X6_3PlaneDualChromaComponentPacked       core1_0.Format = C.VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16
	FormatG10X6_B10X6_R10X6_3PlaneHorizontalChromaComponentPacked core1_0.Format = C.VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16
	FormatG10X6_B10X6_R10X6_3PlaneNoChromaComponentPacked         core1_0.Format = C.VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16
	FormatG12X4B12X4G12X4R12X4_HorizontalChromaComponentPacked    core1_0.Format = C.VK_FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16
	FormatG12X4_B12X4R12X4_2PlaneDualChromaComponentPacked        core1_0.Format = C.VK_FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16
	FormatG12X4_B12X4R12X4_2PlaneHorizontalChromaComponentPacked  core1_0.Format = C.VK_FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16
	FormatG12X4_B12X4_R12X4_3PlaneDualChromaComponentPacked       core1_0.Format = C.VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16
	FormatG12X4_B12X4_R12X4_3PlaneHorizontalChromaComponentPacked core1_0.Format = C.VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16
	FormatG12X4_B12X4_R12X4_3PlaneNoChromaComponentPacked         core1_0.Format = C.VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16
	FormatG16B16G16R16_HorizontalChroma                           core1_0.Format = C.VK_FORMAT_G16B16G16R16_422_UNORM
	FormatG16_B16R16_2PlaneDualChroma                             core1_0.Format = C.VK_FORMAT_G16_B16R16_2PLANE_420_UNORM
	FormatG16_B16R16_2PlaneHorizontalChroma                       core1_0.Format = C.VK_FORMAT_G16_B16R16_2PLANE_422_UNORM
	FormatG16_B16_R16_3PlaneDualChroma                            core1_0.Format = C.VK_FORMAT_G16_B16_R16_3PLANE_420_UNORM
	FormatG16_B16_R16_3PlaneHorizontalChroma                      core1_0.Format = C.VK_FORMAT_G16_B16_R16_3PLANE_422_UNORM
	FormatG16_B16_R16_3PlaneNoChroma                              core1_0.Format = C.VK_FORMAT_G16_B16_R16_3PLANE_444_UNORM
	FormatG8B8G8R8_HorizontalChroma                               core1_0.Format = C.VK_FORMAT_G8B8G8R8_422_UNORM
	FormatG8_B8R8_2PlaneDualChroma                                core1_0.Format = C.VK_FORMAT_G8_B8R8_2PLANE_420_UNORM
	FormatG8_B8R8_2PlaneHorizontalChroma                          core1_0.Format = C.VK_FORMAT_G8_B8R8_2PLANE_422_UNORM
	FormatG8_B8_R8_3PlaneDualChroma                               core1_0.Format = C.VK_FORMAT_G8_B8_R8_3PLANE_420_UNORM
	FormatG8_B8_R8_3PlaneHorizontalChroma                         core1_0.Format = C.VK_FORMAT_G8_B8_R8_3PLANE_422_UNORM
	FormatG8_B8_R8_3PlaneNoChroma                                 core1_0.Format = C.VK_FORMAT_G8_B8_R8_3PLANE_444_UNORM
	FormatR10X6G10X6B10X6A10X6UnsignedNormalizedComponentPacked   core1_0.Format = C.VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16
	FormatR10X6G10X6UnsignedNormalizedComponentPacked             core1_0.Format = C.VK_FORMAT_R10X6G10X6_UNORM_2PACK16
	FormatR10X6UnsignedNormalizedComponentPacked                  core1_0.Format = C.VK_FORMAT_R10X6_UNORM_PACK16
	FormatR12X4G12X4B12X4A12X4UnsignedNormalizedComponentPacked   core1_0.Format = C.VK_FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16
	FormatR12X4G12X4UnsignedNormalizedComponentPacked             core1_0.Format = C.VK_FORMAT_R12X4G12X4_UNORM_2PACK16
	FormatR12X4UnsignedNormalizedComponentPacked                  core1_0.Format = C.VK_FORMAT_R12X4_UNORM_PACK16
)

func init() {
	FormatB10X6G10X6R10X6G10X6HorizontalChromaComponentPacked.Register("B10(X6)G10(X6)R10(X6)G10(X6) Horizontal Chroma (Component-Packed)")
	FormatB12X4G12X4R12X4G12X4HorizontalChromaComponentPacked.Register("B12(X4)G12(X4)R12(X4)G12(X4) Horizontal Chroma (Component-Packed)")
	FormatB16G16R16G16HorizontalChroma.Register("B16G16R16G16 Horizontal Chroma")
	FormatB8G8R8G8HorizontalChroma.Register("B8G8R8G8 Horizontal Chroma")
	FormatG10X6B10X6G10X6R10X6HorizontalChromaComponentPacked.Register("G10(X6)B10(X6)G10(X6)R10(X6) Horizontal Chroma (Component-Packed)")
	FormatG10X6_B10X6R10X6_2PlaneDualChromaComponentPacked.Register("2-Plane G10(X6) B10(X6)R10(X6) Dual Chroma (Component-Packed)")
	FormatG10X6_B10X6R10X6_2PlaneHorizontalChromaComponentPacked.Register("2-Plane G10(X6) B10(X6)R10(X6) Horizontal Chroma (Component-Packed)")
	FormatG10X6_B10X6_R10X6_3PlaneDualChromaComponentPacked.Register("3-Plane G10(X6) B10(X6) R10(X6) Dual Chroma (Component-Packed)")
	FormatG10X6_B10X6_R10X6_3PlaneHorizontalChromaComponentPacked.Register("3-Plane G10(X6) B10(X6) R10(X6) Horizontal Chroma (Component-Packed)")
	FormatG10X6_B10X6_R10X6_3PlaneNoChromaComponentPacked.Register("3-Plane G10(X6) B10(X6) R10(X6) No Chroma (Component-Packed)")
	FormatG12X4B12X4G12X4R12X4_HorizontalChromaComponentPacked.Register("G12(X4)B12(X4)G12(X4)R12(X4) Horizontal Chroma (Component-Packed)")
	FormatG12X4_B12X4R12X4_2PlaneDualChromaComponentPacked.Register("2-Plane G12(X4) B12(X4)R12(X4) Dual Chroma (Component-Packed)")
	FormatG12X4_B12X4R12X4_2PlaneHorizontalChromaComponentPacked.Register("2-Plane G12(X4) B12(X4)R12(X4) Horizontal Chroma (Component-Packed)")
	FormatG12X4_B12X4_R12X4_3PlaneDualChromaComponentPacked.Register("3-Plane G12(X4) B12(X4) R12(X4) Dual Chroma (Component-Packed)")
	FormatG12X4_B12X4_R12X4_3PlaneHorizontalChromaComponentPacked.Register("3-Plane G12(X4) B12(X4) R12(X4) Horizontal Chroma (Component-Packed)")
	FormatG12X4_B12X4_R12X4_3PlaneNoChromaComponentPacked.Register("3-Plane G12(X4) B12(X4) R12(X4) No Chroma (Component-Packed)")
	FormatG16B16G16R16_HorizontalChroma.Register("G16B16G16R16 Horizontal Chroma")
	FormatG16_B16R16_2PlaneDualChroma.Register("2-Plane G16 B16R16 Dual Chroma")
	FormatG16_B16R16_2PlaneHorizontalChroma.Register("2-Plane G16 B16R16 Horizontal Chroma")
	FormatG16_B16_R16_3PlaneDualChroma.Register("3-Plane G16 B16 R16 Dual Chroma")
	FormatG16_B16_R16_3PlaneHorizontalChroma.Register("3-Plane G16 B16 R16 Horizontal Chroma")
	FormatG16_B16_R16_3PlaneNoChroma.Register("3-Plane G16 B16 R16 No Chroma")
	FormatG8B8G8R8_HorizontalChroma.Register("G8B8G8R8 Horizontal Chroma")
	FormatG8_B8R8_2PlaneDualChroma.Register("2-Plane G8 B8R8 Dual Chroma")
	FormatG8_B8R8_2PlaneHorizontalChroma.Register("2-Plane G8 B8R8 Horizontal Chroma")
	FormatG8_B8_R8_3PlaneDualChroma.Register("3-Plane G8 B8 R8 Dual Chroma")
	FormatG8_B8_R8_3PlaneHorizontalChroma.Register("3-Plane G8 B8 R8 Horizontal Chroma")
	FormatG8_B8_R8_3PlaneNoChroma.Register("3-Plane G8 B8 R8 No Chroma")
	FormatR10X6G10X6B10X6A10X6UnsignedNormalizedComponentPacked.Register("R10(X6)G10(X6)B10(X6)A10(X6) Unsigned Normalized (Component-Packed)")
	FormatR10X6G10X6UnsignedNormalizedComponentPacked.Register("R10(X6)G10(X6) Unsigned Normalized (Component-Packed)")
	FormatR10X6UnsignedNormalizedComponentPacked.Register("R10(X6) Unsigned Normalized (Component-Packed)")
	FormatR12X4G12X4B12X4A12X4UnsignedNormalizedComponentPacked.Register("R12(X4)G12(X4)B12(X4)A12(X4) Unsigned Normalized (Component-Packed)")
	FormatR12X4G12X4UnsignedNormalizedComponentPacked.Register("R12(X4)G12(X4) Unsigned Normalized (Component-Packed)")
	FormatR12X4UnsignedNormalizedComponentPacked.Register("R12(X4) Unsigned Normalized (Component-Packed)")
}
