package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type DataFormat int32

const (
	FormatUndefined              DataFormat = C.VK_FORMAT_UNDEFINED
	FormatR4G4UnsignedNormalized DataFormat = C.VK_FORMAT_R4G4_UNORM_PACK8

	FormatR4G4B4A4UnsignedNormalized DataFormat = C.VK_FORMAT_R4G4B4A4_UNORM_PACK16
	FormatB4G4R4A4UnsignedNormalized DataFormat = C.VK_FORMAT_B4G4R4A4_UNORM_PACK16
	FormatR5G6B5UnsignedNormalized   DataFormat = C.VK_FORMAT_R5G6B5_UNORM_PACK16
	FormatB5G6R5UnsignedNormalized   DataFormat = C.VK_FORMAT_B5G6R5_UNORM_PACK16
	FormatR5G5B5A1UnsignedNormalized DataFormat = C.VK_FORMAT_R5G5B5A1_UNORM_PACK16
	FormatB5G5R5A1UnsignedNormalized DataFormat = C.VK_FORMAT_B5G5R5A1_UNORM_PACK16
	FormatA1R5G5B5UnsignedNormalized DataFormat = C.VK_FORMAT_A1R5G5B5_UNORM_PACK16

	FormatR8UnsignedNormalized DataFormat = C.VK_FORMAT_R8_UNORM
	FormatR8SignedNormalized   DataFormat = C.VK_FORMAT_R8_SNORM
	FormatR8UnsignedScaled     DataFormat = C.VK_FORMAT_R8_USCALED
	FormatR8SignedScaled       DataFormat = C.VK_FORMAT_R8_SSCALED
	FormatR8UnsignedInt        DataFormat = C.VK_FORMAT_R8_UINT
	FormatR8SignedInt          DataFormat = C.VK_FORMAT_R8_SINT
	FormatR8SRGB               DataFormat = C.VK_FORMAT_R8_SRGB

	FormatR8G8UnsignedNormalized DataFormat = C.VK_FORMAT_R8G8_UNORM
	FormatR8G8SignedNormalized   DataFormat = C.VK_FORMAT_R8G8_SNORM
	FormatR8G8UnsignedScaled     DataFormat = C.VK_FORMAT_R8G8_USCALED
	FormatR8G8SignedScaled       DataFormat = C.VK_FORMAT_R8G8_SSCALED
	FormatR8G8UnsignedInt        DataFormat = C.VK_FORMAT_R8G8_UINT
	FormatR8G8SignedInt          DataFormat = C.VK_FORMAT_R8G8_SINT
	FormatR8G8SRGB               DataFormat = C.VK_FORMAT_R8G8_SRGB

	FormatR8G8B8UnsignedNormalized DataFormat = C.VK_FORMAT_R8G8B8_UNORM
	FormatR8G8B8SignedNormalized   DataFormat = C.VK_FORMAT_R8G8B8_SNORM
	FormatR8G8B8UnsignedScaled     DataFormat = C.VK_FORMAT_R8G8B8_USCALED
	FormatR8G8B8SignedScaled       DataFormat = C.VK_FORMAT_R8G8B8_SSCALED
	FormatR8G8B8UnsignedInt        DataFormat = C.VK_FORMAT_R8G8B8_UINT
	FormatR8G8B8SignedInt          DataFormat = C.VK_FORMAT_R8G8B8_SINT
	FormatR8G8B8SRGB               DataFormat = C.VK_FORMAT_R8G8B8_SRGB

	FormatB8G8R8UnsignedNormalized DataFormat = C.VK_FORMAT_B8G8R8_UNORM
	FormatB8G8R8SignedNormalized   DataFormat = C.VK_FORMAT_B8G8R8_SNORM
	FormatB8G8R8UnsignedScaled     DataFormat = C.VK_FORMAT_B8G8R8_USCALED
	FormatB8G8R8SignedScaled       DataFormat = C.VK_FORMAT_B8G8R8_SSCALED
	FormatB8G8R8UnsignedInt        DataFormat = C.VK_FORMAT_B8G8R8_UINT
	FormatB8G8R8SignedInt          DataFormat = C.VK_FORMAT_B8G8R8_SINT
	FormatB8G8R8SRGB               DataFormat = C.VK_FORMAT_B8G8R8_SRGB

	FormatR8G8B8A8UnsignedNormalized DataFormat = C.VK_FORMAT_R8G8B8A8_UNORM
	FormatR8G8B8A8SignedNormalized   DataFormat = C.VK_FORMAT_R8G8B8A8_SNORM
	FormatR8G8B8A8UnsignedScaled     DataFormat = C.VK_FORMAT_R8G8B8A8_USCALED
	FormatR8G8B8A8SignedScaled       DataFormat = C.VK_FORMAT_R8G8B8A8_SSCALED
	FormatR8G8B8A8UnsignedInt        DataFormat = C.VK_FORMAT_R8G8B8A8_UINT
	FormatR8G8B8A8SignedInt          DataFormat = C.VK_FORMAT_R8G8B8A8_SINT
	FormatR8G8B8A8SRGB               DataFormat = C.VK_FORMAT_R8G8B8A8_SRGB

	FormatB8G8R8A8UnsignedNormalized DataFormat = C.VK_FORMAT_B8G8R8A8_UNORM
	FormatB8G8R8A8SignedNormalized   DataFormat = C.VK_FORMAT_B8G8R8A8_SNORM
	FormatB8G8R8A8UnsignedScaled     DataFormat = C.VK_FORMAT_B8G8R8A8_USCALED
	FormatB8G8R8A8SignedScaled       DataFormat = C.VK_FORMAT_B8G8R8A8_SSCALED
	FormatB8G8R8A8UnsignedInt        DataFormat = C.VK_FORMAT_B8G8R8A8_UINT
	FormatB8G8R8A8SignedInt          DataFormat = C.VK_FORMAT_B8G8R8A8_SINT
	FormatB8G8R8A8SRGB               DataFormat = C.VK_FORMAT_B8G8R8A8_SRGB

	FormatA8B8G8R8UnsignedNormalized DataFormat = C.VK_FORMAT_A8B8G8R8_UNORM_PACK32
	FormatA8B8G8R8SignedNormalized   DataFormat = C.VK_FORMAT_A8B8G8R8_SNORM_PACK32
	FormatA8B8G8R8UnsignedScaled     DataFormat = C.VK_FORMAT_A8B8G8R8_USCALED_PACK32
	FormatA8B8G8R8SignedScaled       DataFormat = C.VK_FORMAT_A8B8G8R8_SSCALED_PACK32
	FormatA8B8G8R8UnsignedInt        DataFormat = C.VK_FORMAT_A8B8G8R8_UINT_PACK32
	FormatA8B8G8R8SignedInt          DataFormat = C.VK_FORMAT_A8B8G8R8_SINT_PACK32
	FormatA8B8G8R8SRGB               DataFormat = C.VK_FORMAT_A8B8G8R8_SRGB_PACK32

	FormatA2R10G10B10UnsignedNormalized DataFormat = C.VK_FORMAT_A2R10G10B10_UNORM_PACK32
	FormatA2R10G10B10SignedNormalized   DataFormat = C.VK_FORMAT_A2R10G10B10_SNORM_PACK32
	FormatA2R10G10B10UnsignedScaled     DataFormat = C.VK_FORMAT_A2R10G10B10_USCALED_PACK32
	FormatA2R10G10B10SignedScaled       DataFormat = C.VK_FORMAT_A2R10G10B10_SSCALED_PACK32
	FormatA2R10G10B10UnsignedInt        DataFormat = C.VK_FORMAT_A2R10G10B10_UINT_PACK32
	FormatA2R10G10B10SignedInt          DataFormat = C.VK_FORMAT_A2R10G10B10_SINT_PACK32

	FormatA2B10G10R10UnsignedNormalized DataFormat = C.VK_FORMAT_A2B10G10R10_UNORM_PACK32
	FormatA2B10G10R10SignedNormalized   DataFormat = C.VK_FORMAT_A2B10G10R10_SNORM_PACK32
	FormatA2B10G10R10UnsignedScaled     DataFormat = C.VK_FORMAT_A2B10G10R10_USCALED_PACK32
	FormatA2B10G10R10SignedScaled       DataFormat = C.VK_FORMAT_A2B10G10R10_SSCALED_PACK32
	FormatA2B10G10R10UnsignedInt        DataFormat = C.VK_FORMAT_A2B10G10R10_UINT_PACK32
	FormatA2B10G10R10SignedInt          DataFormat = C.VK_FORMAT_A2B10G10R10_SINT_PACK32

	FormatR16UnsignedNormalized DataFormat = C.VK_FORMAT_R16_UNORM
	FormatR16SignedNormalized   DataFormat = C.VK_FORMAT_R16_SNORM
	FormatR16UnsignedScaled     DataFormat = C.VK_FORMAT_R16_USCALED
	FormatR16SignedScaled       DataFormat = C.VK_FORMAT_R16_SSCALED
	FormatR16UnsignedInt        DataFormat = C.VK_FORMAT_R16_UINT
	FormatR16SignedInt          DataFormat = C.VK_FORMAT_R16_SINT
	FormatR16SignedFloat        DataFormat = C.VK_FORMAT_R16_SFLOAT

	FormatR16G16UnsignedNormalized DataFormat = C.VK_FORMAT_R16G16_UNORM
	FormatR16G16SignedNormalized   DataFormat = C.VK_FORMAT_R16G16_SNORM
	FormatR16G16UnsignedScaled     DataFormat = C.VK_FORMAT_R16G16_USCALED
	FormatR16G16SignedScaled       DataFormat = C.VK_FORMAT_R16G16_SSCALED
	FormatR16G16UnsignedInt        DataFormat = C.VK_FORMAT_R16G16_UINT
	FormatR16G16SignedInt          DataFormat = C.VK_FORMAT_R16G16_SINT
	FormatR16G16SignedFloat        DataFormat = C.VK_FORMAT_R16G16_SFLOAT

	FormatR16G16B16UnsignedNormalized DataFormat = C.VK_FORMAT_R16G16B16_UNORM
	FormatR16G16B16SignedNormalized   DataFormat = C.VK_FORMAT_R16G16B16_SNORM
	FormatR16G16B16UnsignedScaled     DataFormat = C.VK_FORMAT_R16G16B16_USCALED
	FormatR16G16B16SignedScaled       DataFormat = C.VK_FORMAT_R16G16B16_SSCALED
	FormatR16G16B16UnsignedInt        DataFormat = C.VK_FORMAT_R16G16B16_UINT
	FormatR16G16B16SignedInt          DataFormat = C.VK_FORMAT_R16G16B16_SINT
	FormatR16G16B16SignedFloat        DataFormat = C.VK_FORMAT_R16G16B16_SFLOAT

	FormatR16G16B16A16UnsignedNormalized DataFormat = C.VK_FORMAT_R16G16B16A16_UNORM
	FormatR16G16B16A16SignedNormalized   DataFormat = C.VK_FORMAT_R16G16B16A16_SNORM
	FormatR16G16B16A16UnsignedScaled     DataFormat = C.VK_FORMAT_R16G16B16A16_USCALED
	FormatR16G16B16A16SignedScaled       DataFormat = C.VK_FORMAT_R16G16B16A16_SSCALED
	FormatR16G16B16A16UnsignedInt        DataFormat = C.VK_FORMAT_R16G16B16A16_UINT
	FormatR16G16B16A16SignedInt          DataFormat = C.VK_FORMAT_R16G16B16A16_SINT
	FormatR16G16B16A16SignedFloat        DataFormat = C.VK_FORMAT_R16G16B16A16_SFLOAT

	FormatR32UnsignedInt          DataFormat = C.VK_FORMAT_R32_UINT
	FormatR32SignedInt            DataFormat = C.VK_FORMAT_R32_SINT
	FormatR32SignedFloat          DataFormat = C.VK_FORMAT_R32_SFLOAT
	FormatR32G32UnsignedInt       DataFormat = C.VK_FORMAT_R32G32_UINT
	FormatR32G32SignedInt         DataFormat = C.VK_FORMAT_R32G32_SINT
	FormatR32G32SignedFloat       DataFormat = C.VK_FORMAT_R32G32_SFLOAT
	FormatR32G32B32UnsignedInt    DataFormat = C.VK_FORMAT_R32G32B32_UINT
	FormatR32G32B32SignedInt      DataFormat = C.VK_FORMAT_R32G32B32_SINT
	FormatR32G32B32SignedFloat    DataFormat = C.VK_FORMAT_R32G32B32_SFLOAT
	FormatR32G32B32A32UnsignedInt DataFormat = C.VK_FORMAT_R32G32B32A32_UINT
	FormatR32G32B32A32SignedInt   DataFormat = C.VK_FORMAT_R32G32B32A32_SINT
	FormatR32G32B32A32SignedFloat DataFormat = C.VK_FORMAT_R32G32B32A32_SFLOAT

	FormatR64UnsignedInt          DataFormat = C.VK_FORMAT_R64_UINT
	FormatR64SignedInt            DataFormat = C.VK_FORMAT_R64_SINT
	FormatR64SignedFloat          DataFormat = C.VK_FORMAT_R64_SFLOAT
	FormatR64G64UnsignedInt       DataFormat = C.VK_FORMAT_R64G64_UINT
	FormatR64G64SignedInt         DataFormat = C.VK_FORMAT_R64G64_SINT
	FormatR64G64SignedFloat       DataFormat = C.VK_FORMAT_R64G64_SFLOAT
	FormatR64G64B64UnsignedInt    DataFormat = C.VK_FORMAT_R64G64B64_UINT
	FormatR64G64B64SignedInt      DataFormat = C.VK_FORMAT_R64G64B64_SINT
	FormatR64G64B64SignedFloat    DataFormat = C.VK_FORMAT_R64G64B64_SFLOAT
	FormatR64G64B64A64UnsignedInt DataFormat = C.VK_FORMAT_R64G64B64A64_UINT
	FormatR64G64B64A64SignedInt   DataFormat = C.VK_FORMAT_R64G64B64A64_SINT
	FormatR64G64B64A64SignedFloat DataFormat = C.VK_FORMAT_R64G64B64A64_SFLOAT

	FormatB10G11R11UnsignedFloat  DataFormat = C.VK_FORMAT_B10G11R11_UFLOAT_PACK32
	FormatE5B9G9R9UnsignedFloat   DataFormat = C.VK_FORMAT_E5B9G9R9_UFLOAT_PACK32
	FormatD16UnsignedNormalized   DataFormat = C.VK_FORMAT_D16_UNORM
	FormatD24X8UnsignedNormalized DataFormat = C.VK_FORMAT_X8_D24_UNORM_PACK32
	FormatD32SignedFloat          DataFormat = C.VK_FORMAT_D32_SFLOAT
	FormatS8UnsignedInt           DataFormat = C.VK_FORMAT_S8_UINT

	FormatD16UnsignedNormalizedS8UnsignedInt DataFormat = C.VK_FORMAT_D16_UNORM_S8_UINT
	FormatD24UnsignedNormalizedS8UnsignedInt DataFormat = C.VK_FORMAT_D24_UNORM_S8_UINT
	FormatD32SignedFloatS8UnsignedInt        DataFormat = C.VK_FORMAT_D32_SFLOAT_S8_UINT

	FormatBC1_RGBUnsignedNormalized  DataFormat = C.VK_FORMAT_BC1_RGB_UNORM_BLOCK
	FormatBC1_RGBsRGB                DataFormat = C.VK_FORMAT_BC1_RGB_SRGB_BLOCK
	FormatBC1_RGBAUnsignedNormalized DataFormat = C.VK_FORMAT_BC1_RGBA_UNORM_BLOCK
	FormatBC1_RGBAsRGB               DataFormat = C.VK_FORMAT_BC1_RGBA_SRGB_BLOCK

	FormatBC2_UnsignedNormalized DataFormat = C.VK_FORMAT_BC2_UNORM_BLOCK
	FormatBC2_sRGB               DataFormat = C.VK_FORMAT_BC2_SRGB_BLOCK

	FormatBC3_UnsignedNormalized DataFormat = C.VK_FORMAT_BC3_UNORM_BLOCK
	FormatBC3_sRGB               DataFormat = C.VK_FORMAT_BC3_SRGB_BLOCK

	FormatBC4_UnsignedNormalized DataFormat = C.VK_FORMAT_BC4_UNORM_BLOCK
	FormatBC4_SignedNormalized   DataFormat = C.VK_FORMAT_BC4_SNORM_BLOCK

	FormatBC5_UnsignedNormalized DataFormat = C.VK_FORMAT_BC5_UNORM_BLOCK
	FormatBC5_SignedNormalized   DataFormat = C.VK_FORMAT_BC5_SNORM_BLOCK

	FormatBC6_UnsignedFloat DataFormat = C.VK_FORMAT_BC6H_UFLOAT_BLOCK
	FormatBC6_SignedFloat   DataFormat = C.VK_FORMAT_BC6H_SFLOAT_BLOCK

	FormatBC7_UnsignedNormalized DataFormat = C.VK_FORMAT_BC7_UNORM_BLOCK
	FormatBC7_sRGB               DataFormat = C.VK_FORMAT_BC7_SRGB_BLOCK

	FormatETC2_R8G8B8UnsignedNormalized   DataFormat = C.VK_FORMAT_ETC2_R8G8B8_UNORM_BLOCK
	FormatETC2_R8G8B8sRGB                 DataFormat = C.VK_FORMAT_ETC2_R8G8B8_SRGB_BLOCK
	FormatETC2_R8G8B8A1UnsignedNormalized DataFormat = C.VK_FORMAT_ETC2_R8G8B8A1_UNORM_BLOCK
	FormatETC2_R8G8B8A1sRGB               DataFormat = C.VK_FORMAT_ETC2_R8G8B8A1_SRGB_BLOCK
	FormatETC2_R8G8B8A8UnsignedNormalized DataFormat = C.VK_FORMAT_ETC2_R8G8B8A8_UNORM_BLOCK
	FormatETC2_R8G8B8A8sRGB               DataFormat = C.VK_FORMAT_ETC2_R8G8B8A8_SRGB_BLOCK

	FormatEAC_R11UnsignedNormalized    DataFormat = C.VK_FORMAT_EAC_R11_UNORM_BLOCK
	FormatEAC_R11SignedNormalized      DataFormat = C.VK_FORMAT_EAC_R11_SNORM_BLOCK
	FormatEAC_R11G11UnsignedNormalized DataFormat = C.VK_FORMAT_EAC_R11G11_UNORM_BLOCK
	FormatEAC_R11G11SignedNormalized   DataFormat = C.VK_FORMAT_EAC_R11G11_SNORM_BLOCK

	FormatASTC4x4_UnsignedNormalized   DataFormat = C.VK_FORMAT_ASTC_4x4_UNORM_BLOCK
	FormatASTC4x4_sRGB                 DataFormat = C.VK_FORMAT_ASTC_4x4_SRGB_BLOCK
	FormatASTC5x4_UnsignedNormalized   DataFormat = C.VK_FORMAT_ASTC_5x4_UNORM_BLOCK
	FormatASTC5x4_sRGB                 DataFormat = C.VK_FORMAT_ASTC_5x4_SRGB_BLOCK
	FormatASTC5x5_UnsignedNormalized   DataFormat = C.VK_FORMAT_ASTC_5x5_UNORM_BLOCK
	FormatASTC5x5_sRGB                 DataFormat = C.VK_FORMAT_ASTC_5x5_SRGB_BLOCK
	FormatASTC6x5_UnsignedNormalized   DataFormat = C.VK_FORMAT_ASTC_6x5_UNORM_BLOCK
	FormatASTC6x5_sRGB                 DataFormat = C.VK_FORMAT_ASTC_6x5_SRGB_BLOCK
	FormatASTC6x6_UnsignedNormalized   DataFormat = C.VK_FORMAT_ASTC_6x6_UNORM_BLOCK
	FormatASTC6x6_sRGB                 DataFormat = C.VK_FORMAT_ASTC_6x6_SRGB_BLOCK
	FormatASTC8x5_UnsignedNormalized   DataFormat = C.VK_FORMAT_ASTC_8x5_UNORM_BLOCK
	FormatASTC8x5_sRGB                 DataFormat = C.VK_FORMAT_ASTC_8x5_SRGB_BLOCK
	FormatASTC8x6_UnsignedNormalized   DataFormat = C.VK_FORMAT_ASTC_8x6_UNORM_BLOCK
	FormatASTC8x6_sRGB                 DataFormat = C.VK_FORMAT_ASTC_8x6_SRGB_BLOCK
	FormatASTC8x8_UnsignedNormalized   DataFormat = C.VK_FORMAT_ASTC_8x8_UNORM_BLOCK
	FormatASTC8x8_sRGB                 DataFormat = C.VK_FORMAT_ASTC_8x8_SRGB_BLOCK
	FormatASTC10x5_UnsignedNormalized  DataFormat = C.VK_FORMAT_ASTC_10x5_UNORM_BLOCK
	FormatASTC10x5_sRGB                DataFormat = C.VK_FORMAT_ASTC_10x5_SRGB_BLOCK
	FormatASTC10x6_UnsignedNormalized  DataFormat = C.VK_FORMAT_ASTC_10x6_UNORM_BLOCK
	FormatASTC10x6_sRGB                DataFormat = C.VK_FORMAT_ASTC_10x6_SRGB_BLOCK
	FormatASTC10x8_UnsignedNormalized  DataFormat = C.VK_FORMAT_ASTC_10x8_UNORM_BLOCK
	FormatASTC10x8_sRGB                DataFormat = C.VK_FORMAT_ASTC_10x8_SRGB_BLOCK
	FormatASTC10x10_UnsignedNormalized DataFormat = C.VK_FORMAT_ASTC_10x10_UNORM_BLOCK
	FormatASTC10x10_sRGB               DataFormat = C.VK_FORMAT_ASTC_10x10_SRGB_BLOCK
	FormatASTC12x10_UnsignedNormalized DataFormat = C.VK_FORMAT_ASTC_12x10_UNORM_BLOCK
	FormatASTC12x10_sRGB               DataFormat = C.VK_FORMAT_ASTC_12x10_SRGB_BLOCK
	FormatASTC12x12_UnsignedNormalized DataFormat = C.VK_FORMAT_ASTC_12x12_UNORM_BLOCK
	FormatASTC12x12_sRGB               DataFormat = C.VK_FORMAT_ASTC_12x12_SRGB_BLOCK

	FormatG8B8G8R8_4x2x2UnsignedNormalized DataFormat = C.VK_FORMAT_G8B8G8R8_422_UNORM
	FormatB8G8R8G8_4x2x2UnsignedNormalized DataFormat = C.VK_FORMAT_B8G8R8G8_422_UNORM

	FormatG8B8R8_3Plane_4x2x0UnsignedNormalized DataFormat = C.VK_FORMAT_G8_B8_R8_3PLANE_420_UNORM
	FormatG8B8R8_2Plane_4x2x0UnsignedNormalized DataFormat = C.VK_FORMAT_G8_B8R8_2PLANE_420_UNORM
	FormatG8B8R8_3Plane_4x2x2UnsignedNormalized DataFormat = C.VK_FORMAT_G8_B8_R8_3PLANE_422_UNORM
	FormatG8B8R8_2Plane_4x2x2UnsignedNormalized DataFormat = C.VK_FORMAT_G8_B8R8_2PLANE_422_UNORM
	FormatG8B8R8_3Plane_4x4x4UnsignedNormalized DataFormat = C.VK_FORMAT_G8_B8_R8_3PLANE_444_UNORM

	FormatR10X6UnsignedNormalized                        DataFormat = C.VK_FORMAT_R10X6_UNORM_PACK16
	FormatR10X6G10X6UnsignedNormalized                   DataFormat = C.VK_FORMAT_R10X6G10X6_UNORM_2PACK16
	FormatR10X6G10X6B10X6A10X6UnsignedNormalized         DataFormat = C.VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16
	FormatG10X6B10X6G10X6R10X6_4x2x2UnsignedNormalized   DataFormat = C.VK_FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16
	FormatB10X6G10X6R10X6G10X6_4x2x2UnsignedNormalized   DataFormat = C.VK_FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16
	FormatG10X6G10X6R10X6_3Plane_4x2x0UnsignedNormalized DataFormat = C.VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16
	FormatG10X6B10X6R10X6_2Plane_4x2x0UnsignedNormalized DataFormat = C.VK_FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16
	FormatG10X6B10X6R10X6_3Plane_4x2x2UnsignedNormalized DataFormat = C.VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16
	FormatG10X6B10X6R10X6_2Plane_4x2x2UnsignedNormalized DataFormat = C.VK_FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16
	FormatG10X6B10X6R10X6_3Plane_4x4x4UnsignedNormalized DataFormat = C.VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16

	FormatR12X4UnsignedNormalized                        DataFormat = C.VK_FORMAT_R12X4_UNORM_PACK16
	FormatR12X4G12X4UnsignedNormalized                   DataFormat = C.VK_FORMAT_R12X4G12X4_UNORM_2PACK16
	FormatR12X4G12X4B12X4A12X4UnsignedNormalized         DataFormat = C.VK_FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16
	FormatG12X4B12X4G12X4R12X4_4x2x2UnsignedNormalized   DataFormat = C.VK_FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16
	FormatB12X4G12X4R12X4G12X4_4x2x2UnsignedNormalized   DataFormat = C.VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16
	FormatG12X4B12X4R12X4_3Plane_4x2x0UnsignedNormalized DataFormat = C.VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16
	FormatG12X4B12X4R12X4_2Plane_4x2x0UnsignedNormalized DataFormat = C.VK_FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16
	FormatG12X4B12X4R12X4_3Plane_4x2x2UnsignedNormalized DataFormat = C.VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16
	FormatG12X4B12X4R12X4_2Plane_4x2x2UnsignedNormalized DataFormat = C.VK_FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16
	FormatG12X4B12X4R12X4_3Plane_4x4x4UnsignedNormalized DataFormat = C.VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16

	FormatG16B16G16R16_4x2x2UnsignedNormalized     DataFormat = C.VK_FORMAT_G16B16G16R16_422_UNORM
	FormatB16G16R16G16_4x2x2UnsignedNormalized     DataFormat = C.VK_FORMAT_B16G16R16G16_422_UNORM
	FormatG16B16R16_3Plane_4x2x0UnsignedNormalized DataFormat = C.VK_FORMAT_G16_B16_R16_3PLANE_420_UNORM
	FormatG16B16R16_2Plane_4x2x0UnsignedNormalized DataFormat = C.VK_FORMAT_G16_B16R16_2PLANE_420_UNORM
	FormatG16B16R16_3Plane_4x2x2UnsignedNormalized DataFormat = C.VK_FORMAT_G16_B16_R16_3PLANE_422_UNORM
	FormatG16B16R16_2Plane_4x2x2UnsignedNormalized DataFormat = C.VK_FORMAT_G16_B16R16_2PLANE_422_UNORM
	FormatG16B16R16_3Plane_4x4x4UnsignedNormalized DataFormat = C.VK_FORMAT_G16_B16_R16_3PLANE_444_UNORM

	FormatPVRTC1_2BPP_UnsignedNormalized DataFormat = C.VK_FORMAT_PVRTC1_2BPP_UNORM_BLOCK_IMG
	FormatPVRTC1_4BPP_UnsignedNormalized DataFormat = C.VK_FORMAT_PVRTC1_4BPP_UNORM_BLOCK_IMG
	FormatPVRTC1_2BPP_sRGB               DataFormat = C.VK_FORMAT_PVRTC1_2BPP_SRGB_BLOCK_IMG
	FormatPVRTC1_4BPP_sRGB               DataFormat = C.VK_FORMAT_PVRTC1_4BPP_SRGB_BLOCK_IMG

	FormatPVRTC2_2BPP_UnsignedNormalized DataFormat = C.VK_FORMAT_PVRTC2_2BPP_UNORM_BLOCK_IMG
	FormatPVRTC2_4BPP_UnsignedNormalized DataFormat = C.VK_FORMAT_PVRTC2_4BPP_UNORM_BLOCK_IMG
	FormatPVRTC2_2BPP_sRGB               DataFormat = C.VK_FORMAT_PVRTC2_2BPP_SRGB_BLOCK_IMG
	FormatPVRTC2_4BPP_sRGB               DataFormat = C.VK_FORMAT_PVRTC2_4BPP_SRGB_BLOCK_IMG

	FormatASTC4x4_SignedFloat   DataFormat = C.VK_FORMAT_ASTC_4x4_SFLOAT_BLOCK_EXT
	FormatASTC5x4_SignedFloat   DataFormat = C.VK_FORMAT_ASTC_5x4_SFLOAT_BLOCK_EXT
	FormatASTC5x5_SignedFloat   DataFormat = C.VK_FORMAT_ASTC_5x5_SFLOAT_BLOCK_EXT
	FormatASTC6x5_SignedFloat   DataFormat = C.VK_FORMAT_ASTC_6x5_SFLOAT_BLOCK_EXT
	FormatASTC6x6_SignedFloat   DataFormat = C.VK_FORMAT_ASTC_6x6_SFLOAT_BLOCK_EXT
	FormatASTC8x5_SignedFloat   DataFormat = C.VK_FORMAT_ASTC_8x5_SFLOAT_BLOCK_EXT
	FormatASTC8x6_SignedFloat   DataFormat = C.VK_FORMAT_ASTC_8x6_SFLOAT_BLOCK_EXT
	FormatASTC8x8_SignedFloat   DataFormat = C.VK_FORMAT_ASTC_8x8_SFLOAT_BLOCK_EXT
	FormatASTC10x5_SignedFloat  DataFormat = C.VK_FORMAT_ASTC_10x5_SFLOAT_BLOCK_EXT
	FormatASTC10x6_SignedFloat  DataFormat = C.VK_FORMAT_ASTC_10x6_SFLOAT_BLOCK_EXT
	FormatASTC10x8_SignedFloat  DataFormat = C.VK_FORMAT_ASTC_10x8_SFLOAT_BLOCK_EXT
	FormatASTC10x10_SignedFloat DataFormat = C.VK_FORMAT_ASTC_10x10_SFLOAT_BLOCK_EXT
	FormatASTC12x10_SignedFloat DataFormat = C.VK_FORMAT_ASTC_12x10_SFLOAT_BLOCK_EXT
	FormatASTC12x12_SignedFloat DataFormat = C.VK_FORMAT_ASTC_12x12_SFLOAT_BLOCK_EXT

	FormatG8B8R8_2Plane_4x4x4UnsignedNormalized          DataFormat = C.VK_FORMAT_G8_B8R8_2PLANE_444_UNORM_EXT
	FormatG10X6B10X6R10X6_2Plane_4x4x4UnsignedNormalized DataFormat = C.VK_FORMAT_G10X6_B10X6R10X6_2PLANE_444_UNORM_3PACK16_EXT
	FormatG12X4B12X4R12X4_2Plane_4x4x4UnsignedNormalized DataFormat = C.VK_FORMAT_G12X4_B12X4R12X4_2PLANE_444_UNORM_3PACK16_EXT
	FormatG16B16R16_2Plane_4x4x4UnsignedNormalized       DataFormat = C.VK_FORMAT_G16_B16R16_2PLANE_444_UNORM_EXT
	FormatA4R4G4B4UnsignedNormalized                     DataFormat = C.VK_FORMAT_A4R4G4B4_UNORM_PACK16_EXT
	FormatA4B4G4R4UnsignedNormalized                     DataFormat = C.VK_FORMAT_A4B4G4R4_UNORM_PACK16_EXT
)

var dataFormatToString = map[DataFormat]string{
	FormatUndefined:              "Undefined",
	FormatR4G4UnsignedNormalized: "R4G4 Unsigned Normalized",

	FormatR4G4B4A4UnsignedNormalized: "R4G4B4A4 Unsigned Normalized",
	FormatB4G4R4A4UnsignedNormalized: "B4G4R4A4 Unsigned Normalized",
	FormatR5G6B5UnsignedNormalized:   "R5G6B5 Unsigned Normalized",
	FormatB5G6R5UnsignedNormalized:   "B5G6R5 Unsigned Normalized",
	FormatR5G5B5A1UnsignedNormalized: "R5G5B5A1 Unsigned Normalized",
	FormatB5G5R5A1UnsignedNormalized: "B5G5R5A1 Unsigned Normalized",
	FormatA1R5G5B5UnsignedNormalized: "A1R5G5B5 Unsigned Normalized",

	FormatR8UnsignedNormalized: "R8 Unsigned Normalized",
	FormatR8SignedNormalized:   "R8 Signed Normalized",
	FormatR8UnsignedScaled:     "R8 Unsigned Scaled",
	FormatR8SignedScaled:       "R8 Signed Scaled",
	FormatR8UnsignedInt:        "R8 Unsigned Int",
	FormatR8SignedInt:          "R8 Signed Int",
	FormatR8SRGB:               "R8 sRGB",

	FormatR8G8UnsignedNormalized: "R8G8 Unsigned Normalized",
	FormatR8G8SignedNormalized:   "R8G8 Signed Normalized",
	FormatR8G8UnsignedScaled:     "R8G8 Unsigned Scaled",
	FormatR8G8SignedScaled:       "R8G8 Signed Scaled",
	FormatR8G8UnsignedInt:        "R8G8 Unsigned Int",
	FormatR8G8SignedInt:          "R8G8 Signed Int",
	FormatR8G8SRGB:               "R8G8 sRGB",

	FormatR8G8B8UnsignedNormalized: "R8G8B8 Unsigned Normalized",
	FormatR8G8B8SignedNormalized:   "R8G8B8 Signed Normalized",
	FormatR8G8B8UnsignedScaled:     "R8G8B8 Unsigned Scaled",
	FormatR8G8B8SignedScaled:       "R8G8B8 Signed Scaled",
	FormatR8G8B8UnsignedInt:        "R8G8B8 Unsigned Int",
	FormatR8G8B8SignedInt:          "R8G8B8 Signed Int",
	FormatR8G8B8SRGB:               "R8G8B8 sRGB",

	FormatB8G8R8UnsignedNormalized: "B8G8R8 Unsigned Normalized",
	FormatB8G8R8SignedNormalized:   "B8G8R8 Signed Normalized",
	FormatB8G8R8UnsignedScaled:     "B8G8R8 Unsigned Scaled",
	FormatB8G8R8SignedScaled:       "B8G8R8 Signed Scaled",
	FormatB8G8R8UnsignedInt:        "B8G8R8 Unsigned Int",
	FormatB8G8R8SignedInt:          "B8G8R8 Signed Int",
	FormatB8G8R8SRGB:               "B8G8R8 sRGB",

	FormatR8G8B8A8UnsignedNormalized: "R8G8B8A8 Unsigned Normalized",
	FormatR8G8B8A8SignedNormalized:   "R8G8B8A8 Signed Normalized",
	FormatR8G8B8A8UnsignedScaled:     "R8G8B8A8 Unsigned Scaled",
	FormatR8G8B8A8SignedScaled:       "R8G8B8A8 Signed Scaled",
	FormatR8G8B8A8UnsignedInt:        "R8G8B8A8 Unsigned Int",
	FormatR8G8B8A8SignedInt:          "R8G8B8A8 Signed Int",
	FormatR8G8B8A8SRGB:               "R8G8B8A8 sRGB",

	FormatB8G8R8A8UnsignedNormalized: "B8G8R8A8 Unsigned Normalized",
	FormatB8G8R8A8SignedNormalized:   "B8G8R8A8 Signed Normalized",
	FormatB8G8R8A8UnsignedScaled:     "B8G8R8A8 Unsigned Scaled",
	FormatB8G8R8A8SignedScaled:       "B8G8R8A8 Signed Scaled",
	FormatB8G8R8A8UnsignedInt:        "B8G8R8A8 Unsigned Int",
	FormatB8G8R8A8SignedInt:          "B8G8R8A8 Signed Int",
	FormatB8G8R8A8SRGB:               "B8G8R8A8 sRGB",

	FormatA8B8G8R8UnsignedNormalized: "A8B8G8R8 Unsigned Normalized",
	FormatA8B8G8R8SignedNormalized:   "A8B8G8R8 Signed Normalized",
	FormatA8B8G8R8UnsignedScaled:     "A8B8G8R8 Unsigned Scaled",
	FormatA8B8G8R8SignedScaled:       "A8B8G8R8 Signed Scaled",
	FormatA8B8G8R8UnsignedInt:        "A8B8G8R8 Unsigned Int",
	FormatA8B8G8R8SignedInt:          "A8B8G8R8 Signed Int",
	FormatA8B8G8R8SRGB:               "A8B8G8R8 sRGB",

	FormatA2R10G10B10UnsignedNormalized: "A2R10G10B10 Unsigned Normalized",
	FormatA2R10G10B10SignedNormalized:   "A2R10G10B10 Signed Normalized",
	FormatA2R10G10B10UnsignedScaled:     "A2R10G10B10 Unsigned Scaled",
	FormatA2R10G10B10SignedScaled:       "A2R10G10B10 Signed Scaled",
	FormatA2R10G10B10UnsignedInt:        "A2R10G10B10 Unsigned Int",
	FormatA2R10G10B10SignedInt:          "A2R10G10B10 Signed Int",

	FormatA2B10G10R10UnsignedNormalized: "A2B10G10R10 Unsigned Normalized",
	FormatA2B10G10R10SignedNormalized:   "A2B10G10R10 Signed Normalized",
	FormatA2B10G10R10UnsignedScaled:     "A2B10G10R10 Unsigned Scaled",
	FormatA2B10G10R10SignedScaled:       "A2B10G10R10 Signed Scaled",
	FormatA2B10G10R10UnsignedInt:        "A2B10G10R10 Unsigned Int",
	FormatA2B10G10R10SignedInt:          "A2B10G10R10 Signed Int",

	FormatR16UnsignedNormalized: "R16 Unsigned Normalized",
	FormatR16SignedNormalized:   "R16 Signed Normalized",
	FormatR16UnsignedScaled:     "R16 Unsigned Scaled",
	FormatR16SignedScaled:       "R16 Signed Scaled",
	FormatR16UnsignedInt:        "R16 Unsigned Int",
	FormatR16SignedInt:          "R16 Signed Int",
	FormatR16SignedFloat:        "R16 Signed Float",

	FormatR16G16UnsignedNormalized: "R16G16 Unsigned Normalized",
	FormatR16G16SignedNormalized:   "R16G16 Signed Normalized",
	FormatR16G16UnsignedScaled:     "R16G16 Unsigned Scaled",
	FormatR16G16SignedScaled:       "R16G16 Signed Scaled",
	FormatR16G16UnsignedInt:        "R16G16 Unsigned Int",
	FormatR16G16SignedInt:          "R16G16 Signed Int",
	FormatR16G16SignedFloat:        "R16G16 Signed Float",

	FormatR16G16B16UnsignedNormalized: "R16G16B16 Unsigned Normalized",
	FormatR16G16B16SignedNormalized:   "R16G16B16 Signed Normalized",
	FormatR16G16B16UnsignedScaled:     "R16G16B16 Unsigned Scaled",
	FormatR16G16B16SignedScaled:       "R16G16B16 Signed Scaled",
	FormatR16G16B16UnsignedInt:        "R16G16B16 Unsigned Int",
	FormatR16G16B16SignedInt:          "R16G16B16 Signed Int",
	FormatR16G16B16SignedFloat:        "R16G16B16 Signed Float",

	FormatR16G16B16A16UnsignedNormalized: "R16G16B16A16 Unsigned Normalized",
	FormatR16G16B16A16SignedNormalized:   "R16G16B16A16 Signed Normalized",
	FormatR16G16B16A16UnsignedScaled:     "R16G16B16A16 Unsigned Scaled",
	FormatR16G16B16A16SignedScaled:       "R16G16B16A16 Signed Scaled",
	FormatR16G16B16A16UnsignedInt:        "R16G16B16A16 Unsigned Int",
	FormatR16G16B16A16SignedInt:          "R16G16B16A16 Signed Int",
	FormatR16G16B16A16SignedFloat:        "R16G16B16A16 Signed Float",

	FormatR32UnsignedInt:          "R32 Unsigned Int",
	FormatR32SignedInt:            "R32 Signed Int",
	FormatR32SignedFloat:          "R32 Signed Float",
	FormatR32G32UnsignedInt:       "R32G32 Unsigned Int",
	FormatR32G32SignedInt:         "R32G32 Signed Int",
	FormatR32G32SignedFloat:       "R32G32 Signed Float",
	FormatR32G32B32UnsignedInt:    "R32G32B32 Unsigned Int",
	FormatR32G32B32SignedInt:      "R32G32B32 Signed Int",
	FormatR32G32B32SignedFloat:    "R32G32B32 Signed Float",
	FormatR32G32B32A32UnsignedInt: "R32G32B32A32 Unsigned Int",
	FormatR32G32B32A32SignedInt:   "R32G32B32A32 Signed Int",
	FormatR32G32B32A32SignedFloat: "R32G32B32A32 Signed Float",

	FormatR64UnsignedInt:          "R64 Unsigned Int",
	FormatR64SignedInt:            "R64 Signed Int",
	FormatR64SignedFloat:          "R64 Signed Float",
	FormatR64G64UnsignedInt:       "R64G64 Unsigned Int",
	FormatR64G64SignedInt:         "R64G64 Signed Int",
	FormatR64G64SignedFloat:       "R64G64 Signed Float",
	FormatR64G64B64UnsignedInt:    "R64G64B64 Unsigned Int",
	FormatR64G64B64SignedInt:      "R64G64B64 Signed Int",
	FormatR64G64B64SignedFloat:    "R64G64B64 Signed Float",
	FormatR64G64B64A64UnsignedInt: "R64G64B64A64 Unsigned Int",
	FormatR64G64B64A64SignedInt:   "R64G64B64A64 Signed Int",
	FormatR64G64B64A64SignedFloat: "R64G64B64A64 Signed Float",

	FormatB10G11R11UnsignedFloat:  "B10G11R11 Unsigned Float",
	FormatE5B9G9R9UnsignedFloat:   "E5B9G9R9 Unsigned Float",
	FormatD16UnsignedNormalized:   "D16 Unsigned Normalized",
	FormatD24X8UnsignedNormalized: "D24X8 Unsigned Normalized",
	FormatD32SignedFloat:          "D32 Signed Float",
	FormatS8UnsignedInt:           "S8 Unsigned Int",

	FormatD16UnsignedNormalizedS8UnsignedInt: "D16 Unsigned Normalized S8 Unsigned Int",
	FormatD24UnsignedNormalizedS8UnsignedInt: "D24 Unsigned Normalized S8 Unsigned Int",
	FormatD32SignedFloatS8UnsignedInt:        "D32 Signed Float S8 Unsigned Int",

	FormatBC1_RGBUnsignedNormalized:  "BC1-Compressed -Compressed RGB Unsigned Normalized",
	FormatBC1_RGBsRGB:                "BC1-Compressed -Compressed RGB sRGB",
	FormatBC1_RGBAUnsignedNormalized: "BC1-Compressed -Compressed RGBA Unsigned Normalized",
	FormatBC1_RGBAsRGB:               "BC1-Compressed RGBA sRGB",

	FormatBC2_UnsignedNormalized: "BC2-Compressed Unsigned Normalized",
	FormatBC2_sRGB:               "BC2-Compressed sRGB",

	FormatBC3_UnsignedNormalized: "BC3-Compressed Unsigned Normalized",
	FormatBC3_sRGB:               "BC3-Compressed sRGB",

	FormatBC4_UnsignedNormalized: "BC4-Compressed Unsigned Normalized",
	FormatBC4_SignedNormalized:   "BC4-Compressed Signed Normalized",

	FormatBC5_UnsignedNormalized: "BC5-Compressed Unsigned Normalized",
	FormatBC5_SignedNormalized:   "BC5-Compressed Signed Normalized",

	FormatBC6_UnsignedFloat: "BC6-Compressed Unsigned Float",
	FormatBC6_SignedFloat:   "BC6-Compressed Signed Float",

	FormatBC7_UnsignedNormalized: "BC7-Compressed Unsigned Normalized",
	FormatBC7_sRGB:               "BC7-Compressed sRGB",

	FormatETC2_R8G8B8UnsignedNormalized:   "ETC2-Compressed R8G8B8 Unsigned Normalized",
	FormatETC2_R8G8B8sRGB:                 "ETC2-Compressed R8G8B8 sRGB",
	FormatETC2_R8G8B8A1UnsignedNormalized: "ETC2-Compressed R8G8B8A1 Unsigned Normalized",
	FormatETC2_R8G8B8A1sRGB:               "ETC2-Compressed R8G8B8A1 sRGB",
	FormatETC2_R8G8B8A8UnsignedNormalized: "ETC2-Compressed R8G8B8A8 Unsigned Normalized",
	FormatETC2_R8G8B8A8sRGB:               "ETC2-Compressed R8G8B8A8 sRGB",

	FormatEAC_R11UnsignedNormalized:    "EAC-Compressed R11 Unsigned Normalized",
	FormatEAC_R11SignedNormalized:      "EAC-Compressed R11 Signed Normalized",
	FormatEAC_R11G11UnsignedNormalized: "EAC-Compressed R11G11 Unsigned Normalized",
	FormatEAC_R11G11SignedNormalized:   "EAC-Compressed R11G11 Signed Normalized",

	FormatASTC4x4_UnsignedNormalized:   "ASTC-Compressed (4x4) Unsigned Normalized",
	FormatASTC4x4_sRGB:                 "ASTC-Compressed (4x4) sRGB",
	FormatASTC5x4_UnsignedNormalized:   "ASTC-Compressed (5x4) Unsigned Normalized",
	FormatASTC5x4_sRGB:                 "ASTC-Compressed (5x4) sRGB",
	FormatASTC5x5_UnsignedNormalized:   "ASTC-Compressed (5x5) Unsigned Normalized",
	FormatASTC5x5_sRGB:                 "ASTC-Compressed (5x5) sRGB",
	FormatASTC6x5_UnsignedNormalized:   "ASTC-Compressed (6x5) Unsigned Normalized",
	FormatASTC6x5_sRGB:                 "ASTC-Compressed (6x5) sRGB",
	FormatASTC6x6_UnsignedNormalized:   "ASTC-Compressed (6x6) Unsigned Normalized",
	FormatASTC6x6_sRGB:                 "ASTC-Compressed (6x6) sRGB",
	FormatASTC8x5_UnsignedNormalized:   "ASTC-Compressed (8x5) Unsigned Normalized",
	FormatASTC8x5_sRGB:                 "ASTC-Compressed (8x5) sRGB",
	FormatASTC8x6_UnsignedNormalized:   "ASTC-Compressed (8x6) Unsigned Normalized",
	FormatASTC8x6_sRGB:                 "ASTC-Compressed (8x6) sRGB",
	FormatASTC8x8_UnsignedNormalized:   "ASTC-Compressed (8x8) Unsigned Normalized",
	FormatASTC8x8_sRGB:                 "ASTC-Compressed (8x8) sRGB",
	FormatASTC10x5_UnsignedNormalized:  "ASTC-Compressed (10x5) Unsigned Normalized",
	FormatASTC10x5_sRGB:                "ASTC-Compressed (10x5) sRGB",
	FormatASTC10x6_UnsignedNormalized:  "ASTC-Compressed (10x6) Unsigned Normalized",
	FormatASTC10x6_sRGB:                "ASTC-Compressed (10x6) sRGB",
	FormatASTC10x8_UnsignedNormalized:  "ASTC-Compressed (10x8) Unsigned Normalized",
	FormatASTC10x8_sRGB:                "ASTC-Compressed (10x8) sRGB",
	FormatASTC10x10_UnsignedNormalized: "ASTC-Compressed (10x10) Unsigned Normalized",
	FormatASTC10x10_sRGB:               "ASTC-Compressed (10x10) sRGB",
	FormatASTC12x10_UnsignedNormalized: "ASTC-Compressed (12x10) Unsigned Normalized",
	FormatASTC12x10_sRGB:               "ASTC-Compressed (12x10) sRGB",
	FormatASTC12x12_UnsignedNormalized: "ASTC-Compressed (12x12) Unsigned Normalized",
	FormatASTC12x12_sRGB:               "ASTC-Compressed (12x12) sRGB",

	FormatG8B8G8R8_4x2x2UnsignedNormalized: "G8B8G8R8 (4:2:2) Unsigned Normalized",
	FormatB8G8R8G8_4x2x2UnsignedNormalized: "B8G8R8G8 (4:2:2) Unsigned Normalized",

	FormatG8B8R8_3Plane_4x2x0UnsignedNormalized: "G8B8R8 3-Plane (4:2:0) Unsigned Normalized",
	FormatG8B8R8_2Plane_4x2x0UnsignedNormalized: "G8B8R8 2-Plane (4:2:0) Unsigned Normalized",
	FormatG8B8R8_3Plane_4x2x2UnsignedNormalized: "G8B8R8 3-Plane (4:2:2) Unsigned Normalized",
	FormatG8B8R8_2Plane_4x2x2UnsignedNormalized: "G8B8R8 2-Plane (4:2:2) Unsigned Normalized",
	FormatG8B8R8_3Plane_4x4x4UnsignedNormalized: "G8B8R8 3-Plane (4:4:4) Unsigned Normalized",

	FormatR10X6UnsignedNormalized:                        "R10X6 Unsigned Normalized",
	FormatR10X6G10X6UnsignedNormalized:                   "R10X6G10X6 Unsigned Normalized",
	FormatR10X6G10X6B10X6A10X6UnsignedNormalized:         "R10X6G10X6B10X6A10X6 Unsigned Normalized",
	FormatG10X6B10X6G10X6R10X6_4x2x2UnsignedNormalized:   "G10X6B10X6G10X6R10X6 (4:2:2) Unsigned Normalized",
	FormatB10X6G10X6R10X6G10X6_4x2x2UnsignedNormalized:   "B10X6G10X6R10X6G10X6 (4:2:2) Unsigned Normalized",
	FormatG10X6G10X6R10X6_3Plane_4x2x0UnsignedNormalized: "G10X6G10X6R10X6 3-Plane (4:2:0) Unsigned Normalized",
	FormatG10X6B10X6R10X6_2Plane_4x2x0UnsignedNormalized: "G10X6B10X6R10X6 2-Plane (4:2:0) Unsigned Normalized",
	FormatG10X6B10X6R10X6_3Plane_4x2x2UnsignedNormalized: "G10X6B10X6R10X6 3-Plane (4:2:2) Unsigned Normalized",
	FormatG10X6B10X6R10X6_2Plane_4x2x2UnsignedNormalized: "G10X6B10X6R10X6 2-Plane (4:2:2) Unsigned Normalized",
	FormatG10X6B10X6R10X6_3Plane_4x4x4UnsignedNormalized: "G10X6B10X6R10X6 3-Plane (4:4:4) Unsigned Normalized",

	FormatR12X4UnsignedNormalized:                        "R12X4 Unsigned Normalized",
	FormatR12X4G12X4UnsignedNormalized:                   "R12X4G12X4 Unsigned Normalized",
	FormatR12X4G12X4B12X4A12X4UnsignedNormalized:         "R12X4G12X4B12X4A12X4 Unsigned Normalized",
	FormatG12X4B12X4G12X4R12X4_4x2x2UnsignedNormalized:   "G12X4B12X4G12X4R12X4 (4:2:2) Unsigned Normalized",
	FormatB12X4G12X4R12X4G12X4_4x2x2UnsignedNormalized:   "B12X4G12X4R12X4G12X4 (4:2:2) Unsigned Normalized",
	FormatG12X4B12X4R12X4_3Plane_4x2x0UnsignedNormalized: "G12X4B12X4R12X4 3-Plane (4:2:0) Unsigned Normalized",
	FormatG12X4B12X4R12X4_2Plane_4x2x0UnsignedNormalized: "G12X4B12X4R12X4 2-Plane (4:2:0) Unsigned Normalized",
	FormatG12X4B12X4R12X4_3Plane_4x2x2UnsignedNormalized: "G12X4B12X4R12X4 3-Plane (4:2:2) Unsigned Normalized",
	FormatG12X4B12X4R12X4_2Plane_4x2x2UnsignedNormalized: "G12X4B12X4R12X4 2-Plane (4:2:2) Unsigned Normalized",
	FormatG12X4B12X4R12X4_3Plane_4x4x4UnsignedNormalized: "G12X4B12X4R12X4 3-Plane (4:4:4) Unsigned Normalized",

	FormatG16B16G16R16_4x2x2UnsignedNormalized:     "G16B16G16R16 (4:2:2) Unsigned Normalized",
	FormatB16G16R16G16_4x2x2UnsignedNormalized:     "B16G16R16G16 (4:2:2) Unsigned Normalized",
	FormatG16B16R16_3Plane_4x2x0UnsignedNormalized: "G16B16R16 3-Plane (4:2:0) Unsigned Normalized",
	FormatG16B16R16_2Plane_4x2x0UnsignedNormalized: "G16B16R16 2-Plane (4:2:0) Unsigned Normalized",
	FormatG16B16R16_3Plane_4x2x2UnsignedNormalized: "G16B16R16 3-Plane (4:2:2) Unsigned Normalized",
	FormatG16B16R16_2Plane_4x2x2UnsignedNormalized: "G16B16R16 2-Plane (4:2:2) Unsigned Normalized",
	FormatG16B16R16_3Plane_4x4x4UnsignedNormalized: "G16B16R16 3-Plane (4:4:4) Unsigned Normalized",

	FormatPVRTC1_2BPP_UnsignedNormalized: "PVRTC1-Compressed (2 BPP) Unsigned Normalized",
	FormatPVRTC1_4BPP_UnsignedNormalized: "PVRTC1-Compressed (4 BPP) Unsigned Normalized",
	FormatPVRTC1_2BPP_sRGB:               "PVRTC1-Compressed (2 BPP) sRGB",
	FormatPVRTC1_4BPP_sRGB:               "PVRTC1-Compressed (4 BPP) sRGB",

	FormatPVRTC2_2BPP_UnsignedNormalized: "PVRTC2-Compressed (2 BPP) Unsigned Normalized",
	FormatPVRTC2_4BPP_UnsignedNormalized: "PVRTC2-Compressed (4 BPP) Unsigned Normalized",
	FormatPVRTC2_2BPP_sRGB:               "PVRTC2-Compressed (2 BPP) sRGB",
	FormatPVRTC2_4BPP_sRGB:               "PVRTC2-Compressed (4 BPP) sRGB",

	FormatASTC4x4_SignedFloat:   "ASTC-Compressed (4x4) Signed Float",
	FormatASTC5x4_SignedFloat:   "ASTC-Compressed (5x4) Signed Float",
	FormatASTC5x5_SignedFloat:   "ASTC-Compressed (5x5) Signed Float",
	FormatASTC6x5_SignedFloat:   "ASTC-Compressed (6x5) Signed Float",
	FormatASTC6x6_SignedFloat:   "ASTC-Compressed (6x6) Signed Float",
	FormatASTC8x5_SignedFloat:   "ASTC-Compressed (8x5) Signed Float",
	FormatASTC8x6_SignedFloat:   "ASTC-Compressed (8x6) Signed Float",
	FormatASTC8x8_SignedFloat:   "ASTC-Compressed (8x8) Signed Float",
	FormatASTC10x5_SignedFloat:  "ASTC-Compressed (10x5) Signed Float",
	FormatASTC10x6_SignedFloat:  "ASTC-Compressed (10x6) Signed Float",
	FormatASTC10x8_SignedFloat:  "ASTC-Compressed (10x8) Signed Float",
	FormatASTC10x10_SignedFloat: "ASTC-Compressed (10x10) Signed Float",
	FormatASTC12x10_SignedFloat: "ASTC-Compressed (12x10) Signed Float",
	FormatASTC12x12_SignedFloat: "ASTC-Compressed (12x12) Signed Float",

	FormatG8B8R8_2Plane_4x4x4UnsignedNormalized:          "G8B8R8 2-Plane (4:4:4) Unsigned Normalized",
	FormatG10X6B10X6R10X6_2Plane_4x4x4UnsignedNormalized: "G10X6B10X6R10X6 2-Plane (4:4:4) Unsigned Normalized",
	FormatG12X4B12X4R12X4_2Plane_4x4x4UnsignedNormalized: "G12X4B12X4R12X4 2-Plane (4:4:4) Unsigned Normalized",
	FormatG16B16R16_2Plane_4x4x4UnsignedNormalized:       "G16B16R16 2-Plane (4:4:4) Unsigned Normalized",
	FormatA4R4G4B4UnsignedNormalized:                     "A4R4G4B4 Unsigned Normalized",
	FormatA4B4G4R4UnsignedNormalized:                     "A4B4G4R4 Unsigned Normalized",
}

func (f DataFormat) String() string {
	return dataFormatToString[f]
}
