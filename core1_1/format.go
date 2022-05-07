package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

const (
	DataFormatB10X6G10X6R10X6G10X6HorizontalChromaComponentPacked     common.DataFormat = C.VK_FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16
	DataFormatB12X4G12X4R12X4G12X4HorizontalChromaComponentPacked     common.DataFormat = C.VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16
	DataFormatB16G16R16G16HorizontalChroma                            common.DataFormat = C.VK_FORMAT_B16G16R16G16_422_UNORM
	DataFormatB8G8R8G8HorizontalChroma                                common.DataFormat = C.VK_FORMAT_B8G8R8G8_422_UNORM
	DataFormatG10X6B10X6G10X6R10X6HorizontalChromaComponentPacked     common.DataFormat = C.VK_FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16
	DataFormatG10X6_B10X6R10X6_2PlaneDualChromaComponentPacked        common.DataFormat = C.VK_FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16
	DataFormatG10X6_B10X6R10X6_2PlaneHorizontalChromaComponentPacked  common.DataFormat = C.VK_FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16
	DataFormatG10X6_B10X6_R10X6_3PlaneDualChromaComponentPacked       common.DataFormat = C.VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16
	DataFormatG10X6_B10X6_R10X6_3PlaneHorizontalChromaComponentPacked common.DataFormat = C.VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16
	DataFormatG10X6_B10X6_R10X6_3PlaneNoChromaComponentPacked         common.DataFormat = C.VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16
	DataFormatG12X4B12X4G12X4R12X4_HorizontalChromaComponentPacked    common.DataFormat = C.VK_FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16
	DataFormatG12X4_B12X4R12X4_2PlaneDualChromaComponentPacked        common.DataFormat = C.VK_FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16
	DataFormatG12X4_B12X4R12X4_2PlaneHorizontalChromaComponentPacked  common.DataFormat = C.VK_FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16
	DataFormatG12X4_B12X4_R12X4_3PlaneDualChromaComponentPacked       common.DataFormat = C.VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16
	DataFormatG12X4_B12X4_R12X4_3PlaneHorizontalChromaComponentPacked common.DataFormat = C.VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16
	DataFormatG12X4_B12X4_R12X4_3PlaneNoChromaComponentPacked         common.DataFormat = C.VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16
	DataFormatG16B16G16R16_HorizontalChroma                           common.DataFormat = C.VK_FORMAT_G16B16G16R16_422_UNORM
	DataFormatG16_B16R16_2PlaneDualChroma                             common.DataFormat = C.VK_FORMAT_G16_B16R16_2PLANE_420_UNORM
	DataFormatG16_B16R16_2PlaneHorizontalChroma                       common.DataFormat = C.VK_FORMAT_G16_B16R16_2PLANE_422_UNORM
	DataFormatG16_B16_R16_3PlaneDualChroma                            common.DataFormat = C.VK_FORMAT_G16_B16_R16_3PLANE_420_UNORM
	DataFormatG16_B16_R16_3PlaneHorizontalChroma                      common.DataFormat = C.VK_FORMAT_G16_B16_R16_3PLANE_422_UNORM
	DataFormatG16_B16_R16_3PlaneNoChroma                              common.DataFormat = C.VK_FORMAT_G16_B16_R16_3PLANE_444_UNORM
	DataFormatG8B8G8R8_HorizontalChroma                               common.DataFormat = C.VK_FORMAT_G8B8G8R8_422_UNORM
	DataFormatG8_B8R8_2PlaneDualChroma                                common.DataFormat = C.VK_FORMAT_G8_B8R8_2PLANE_420_UNORM
	DataFormatG8_B8R8_2PlaneHorizontalChroma                          common.DataFormat = C.VK_FORMAT_G8_B8R8_2PLANE_422_UNORM
	DataFormatG8_B8_R8_3PlaneDualChroma                               common.DataFormat = C.VK_FORMAT_G8_B8_R8_3PLANE_420_UNORM
	DataFormatG8_B8_R8_3PlaneHorizontalChroma                         common.DataFormat = C.VK_FORMAT_G8_B8_R8_3PLANE_422_UNORM
	DataFormatG8_B8_R8_3PlaneNoChroma                                 common.DataFormat = C.VK_FORMAT_G8_B8_R8_3PLANE_444_UNORM
	DataFormatR10X6G10X6B10X6A10X6UnsignedNormalizedComponentPacked   common.DataFormat = C.VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16
	DataFormatR10X6G10X6UnsignedNormalizedComponentPacked             common.DataFormat = C.VK_FORMAT_R10X6G10X6_UNORM_2PACK16
	DataFormatR10X6UnsignedNormalizedComponentPacked                  common.DataFormat = C.VK_FORMAT_R10X6_UNORM_PACK16
	DataFormatR12X4G12X4B12X4A12X4UnsignedNormalizedComponentPacked   common.DataFormat = C.VK_FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16
	DataFormatR12X4G12X4UnsignedNormalizedComponentPacked             common.DataFormat = C.VK_FORMAT_R12X4G12X4_UNORM_2PACK16
	DataFormatR12X4UnsignedNormalizedComponentPacked                  common.DataFormat = C.VK_FORMAT_R12X4_UNORM_PACK16
)

func init() {
	DataFormatB10X6G10X6R10X6G10X6HorizontalChromaComponentPacked.Register("B10(X6)G10(X6)R10(X6)G10(X6) Horizontal Chroma (Component-Packed)")
	DataFormatB12X4G12X4R12X4G12X4HorizontalChromaComponentPacked.Register("B12(X4)G12(X4)R12(X4)G12(X4) Horizontal Chroma (Component-Packed)")
	DataFormatB16G16R16G16HorizontalChroma.Register("B16G16R16G16 Horizontal Chroma")
	DataFormatB8G8R8G8HorizontalChroma.Register("B8G8R8G8 Horizontal Chroma")
	DataFormatG10X6B10X6G10X6R10X6HorizontalChromaComponentPacked.Register("G10(X6)B10(X6)G10(X6)R10(X6) Horizontal Chroma (Component-Packed)")
	DataFormatG10X6_B10X6R10X6_2PlaneDualChromaComponentPacked.Register("2-Plane G10(X6) B10(X6)R10(X6) Dual Chroma (Component-Packed)")
	DataFormatG10X6_B10X6R10X6_2PlaneHorizontalChromaComponentPacked.Register("2-Plane G10(X6) B10(X6)R10(X6) Horizontal Chroma (Component-Packed)")
	DataFormatG10X6_B10X6_R10X6_3PlaneDualChromaComponentPacked.Register("3-Plane G10(X6) B10(X6) R10(X6) Dual Chroma (Component-Packed)")
	DataFormatG10X6_B10X6_R10X6_3PlaneHorizontalChromaComponentPacked.Register("3-Plane G10(X6) B10(X6) R10(X6) Horizontal Chroma (Component-Packed)")
	DataFormatG10X6_B10X6_R10X6_3PlaneNoChromaComponentPacked.Register("3-Plane G10(X6) B10(X6) R10(X6) No Chroma (Component-Packed)")
	DataFormatG12X4B12X4G12X4R12X4_HorizontalChromaComponentPacked.Register("G12(X4)B12(X4)G12(X4)R12(X4) Horizontal Chroma (Component-Packed)")
	DataFormatG12X4_B12X4R12X4_2PlaneDualChromaComponentPacked.Register("2-Plane G12(X4) B12(X4)R12(X4) Dual Chroma (Component-Packed)")
	DataFormatG12X4_B12X4R12X4_2PlaneHorizontalChromaComponentPacked.Register("2-Plane G12(X4) B12(X4)R12(X4) Horizontal Chroma (Component-Packed)")
	DataFormatG12X4_B12X4_R12X4_3PlaneDualChromaComponentPacked.Register("3-Plane G12(X4) B12(X4) R12(X4) Dual Chroma (Component-Packed)")
	DataFormatG12X4_B12X4_R12X4_3PlaneHorizontalChromaComponentPacked.Register("3-Plane G12(X4) B12(X4) R12(X4) Horizontal Chroma (Component-Packed)")
	DataFormatG12X4_B12X4_R12X4_3PlaneNoChromaComponentPacked.Register("3-Plane G12(X4) B12(X4) R12(X4) No Chroma (Component-Packed)")
	DataFormatG16B16G16R16_HorizontalChroma.Register("G16B16G16R16 Horizontal Chroma")
	DataFormatG16_B16R16_2PlaneDualChroma.Register("2-Plane G16 B16R16 Dual Chroma")
	DataFormatG16_B16R16_2PlaneHorizontalChroma.Register("2-Plane G16 B16R16 Horizontal Chroma")
	DataFormatG16_B16_R16_3PlaneDualChroma.Register("3-Plane G16 B16 R16 Dual Chroma")
	DataFormatG16_B16_R16_3PlaneHorizontalChroma.Register("3-Plane G16 B16 R16 Horizontal Chroma")
	DataFormatG16_B16_R16_3PlaneNoChroma.Register("3-Plane G16 B16 R16 No Chroma")
	DataFormatG8B8G8R8_HorizontalChroma.Register("G8B8G8R8 Horizontal Chroma")
	DataFormatG8_B8R8_2PlaneDualChroma.Register("2-Plane G8 B8R8 Dual Chroma")
	DataFormatG8_B8R8_2PlaneHorizontalChroma.Register("2-Plane G8 B8R8 Horizontal Chroma")
	DataFormatG8_B8_R8_3PlaneDualChroma.Register("3-Plane G8 B8 R8 Dual Chroma")
	DataFormatG8_B8_R8_3PlaneHorizontalChroma.Register("3-Plane G8 B8 R8 Horizontal Chroma")
	DataFormatG8_B8_R8_3PlaneNoChroma.Register("3-Plane G8 B8 R8 No Chroma")
	DataFormatR10X6G10X6B10X6A10X6UnsignedNormalizedComponentPacked.Register("R10(X6)G10(X6)B10(X6)A10(X6) Unsigned Normalized (Component-Packed)")
	DataFormatR10X6G10X6UnsignedNormalizedComponentPacked.Register("R10(X6)G10(X6) Unsigned Normalized (Component-Packed)")
	DataFormatR10X6UnsignedNormalizedComponentPacked.Register("R10(X6) Unsigned Normalized (Component-Packed)")
	DataFormatR12X4G12X4B12X4A12X4UnsignedNormalizedComponentPacked.Register("R12(X4)G12(X4)B12(X4)A12(X4) Unsigned Normalized (Component-Packed)")
	DataFormatR12X4G12X4UnsignedNormalizedComponentPacked.Register("R12(X4)G12(X4) Unsigned Normalized (Component-Packed)")
	DataFormatR12X4UnsignedNormalizedComponentPacked.Register("R12(X4) Unsigned Normalized (Component-Packed)")
}
