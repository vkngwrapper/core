package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type ColorFormat int32

const (
	Undefined              ColorFormat = C.VK_FORMAT_UNDEFINED
	R4G4UnsignedNormalized ColorFormat = C.VK_FORMAT_R4G4_UNORM_PACK8

	R4G4B4A4UnsignedNormalized ColorFormat = C.VK_FORMAT_R4G4B4A4_UNORM_PACK16
	B4G4R4A4UnsignedNormalized ColorFormat = C.VK_FORMAT_B4G4R4A4_UNORM_PACK16
	R5G6B5UnsignedNormalized   ColorFormat = C.VK_FORMAT_R5G6B5_UNORM_PACK16
	B5G6R5UnsignedNormalized   ColorFormat = C.VK_FORMAT_B5G6R5_UNORM_PACK16
	R5G5B5A1UnsignedNormalized ColorFormat = C.VK_FORMAT_R5G5B5A1_UNORM_PACK16
	B5G5R5A1UnsignedNormalized ColorFormat = C.VK_FORMAT_B5G5R5A1_UNORM_PACK16
	A1R5G5B5UnsignedNormalized ColorFormat = C.VK_FORMAT_A1R5G5B5_UNORM_PACK16

	R8UnsignedNormalized ColorFormat = C.VK_FORMAT_R8_UNORM
	R8SignedNormalized   ColorFormat = C.VK_FORMAT_R8_SNORM
	R8UnsignedScaled     ColorFormat = C.VK_FORMAT_R8_USCALED
	R8SignedScaled       ColorFormat = C.VK_FORMAT_R8_SSCALED
	R8UnsignedInt        ColorFormat = C.VK_FORMAT_R8_UINT
	R8SignedInt          ColorFormat = C.VK_FORMAT_R8_SINT
	R8SRGB               ColorFormat = C.VK_FORMAT_R8_SRGB

	R8G8UnsignedNormalized ColorFormat = C.VK_FORMAT_R8G8_UNORM
	R8G8SignedNormalized   ColorFormat = C.VK_FORMAT_R8G8_SNORM
	R8G8UnsignedScaled     ColorFormat = C.VK_FORMAT_R8G8_USCALED
	R8G8SignedScaled       ColorFormat = C.VK_FORMAT_R8G8_SSCALED
	R8G8UnsignedInt        ColorFormat = C.VK_FORMAT_R8G8_UINT
	R8G8SignedInt          ColorFormat = C.VK_FORMAT_R8G8_SINT
	R8G8SRGB               ColorFormat = C.VK_FORMAT_R8G8_SRGB

	R8G8B8UnsignedNormalized ColorFormat = C.VK_FORMAT_R8G8B8_UNORM
	R8G8B8SignedNormalized   ColorFormat = C.VK_FORMAT_R8G8B8_SNORM
	R8G8B8UnsignedScaled     ColorFormat = C.VK_FORMAT_R8G8B8_USCALED
	R8G8B8SignedScaled       ColorFormat = C.VK_FORMAT_R8G8B8_SSCALED
	R8G8B8UnsignedInt        ColorFormat = C.VK_FORMAT_R8G8B8_UINT
	R8G8B8SignedInt          ColorFormat = C.VK_FORMAT_R8G8B8_SINT
	R8G8B8SRGB               ColorFormat = C.VK_FORMAT_R8G8B8_SRGB

	B8G8R8UnsignedNormalized ColorFormat = C.VK_FORMAT_B8G8R8_UNORM
	B8G8R8SignedNormalized   ColorFormat = C.VK_FORMAT_B8G8R8_SNORM
	B8G8R8UnsignedScaled     ColorFormat = C.VK_FORMAT_B8G8R8_USCALED
	B8G8R8SignedScaled       ColorFormat = C.VK_FORMAT_B8G8R8_SSCALED
	B8G8R8UnsignedInt        ColorFormat = C.VK_FORMAT_B8G8R8_UINT
	B8G8R8SignedInt          ColorFormat = C.VK_FORMAT_B8G8R8_SINT
	B8G8R8SRGB               ColorFormat = C.VK_FORMAT_B8G8R8_SRGB

	R8G8B8A8UnsignedNormalized ColorFormat = C.VK_FORMAT_R8G8B8A8_UNORM
	R8G8B8A8SignedNormalized   ColorFormat = C.VK_FORMAT_R8G8B8A8_SNORM
	R8G8B8A8UnsignedScaled     ColorFormat = C.VK_FORMAT_R8G8B8A8_USCALED
	R8G8B8A8SignedScaled       ColorFormat = C.VK_FORMAT_R8G8B8A8_SSCALED
	R8G8B8A8UnsignedInt        ColorFormat = C.VK_FORMAT_R8G8B8A8_UINT
	R8G8B8A8SignedInt          ColorFormat = C.VK_FORMAT_R8G8B8A8_SINT
	R8G8B8A8SRGB               ColorFormat = C.VK_FORMAT_R8G8B8A8_SRGB

	B8G8R8A8UnsignedNormalized ColorFormat = C.VK_FORMAT_B8G8R8A8_UNORM
	B8G8R8A8SignedNormalized   ColorFormat = C.VK_FORMAT_B8G8R8A8_SNORM
	B8G8R8A8UnsignedScaled     ColorFormat = C.VK_FORMAT_B8G8R8A8_USCALED
	B8G8R8A8SignedScaled       ColorFormat = C.VK_FORMAT_B8G8R8A8_SSCALED
	B8G8R8A8UnsignedInt        ColorFormat = C.VK_FORMAT_B8G8R8A8_UINT
	B8G8R8A8SignedInt          ColorFormat = C.VK_FORMAT_B8G8R8A8_SINT
	B8G8R8A8SRGB               ColorFormat = C.VK_FORMAT_B8G8R8A8_SRGB

	A8B8G8R8UnsignedNormalized ColorFormat = C.VK_FORMAT_A8B8G8R8_UNORM_PACK32
	A8B8G8R8SignedNormalized   ColorFormat = C.VK_FORMAT_A8B8G8R8_SNORM_PACK32
	A8B8G8R8UnsignedScaled     ColorFormat = C.VK_FORMAT_A8B8G8R8_USCALED_PACK32
	A8B8G8R8SignedScaled       ColorFormat = C.VK_FORMAT_A8B8G8R8_SSCALED_PACK32
	A8B8G8R8UnsignedInt        ColorFormat = C.VK_FORMAT_A8B8G8R8_UINT_PACK32
	A8B8G8R8SignedInt          ColorFormat = C.VK_FORMAT_A8B8G8R8_SINT_PACK32
	A8B8G8R8SRGB               ColorFormat = C.VK_FORMAT_A8B8G8R8_SRGB_PACK32

	A2R10G10B10UnsignedNormalized ColorFormat = C.VK_FORMAT_A2R10G10B10_UNORM_PACK32
	A2R10G10B10SignedNormalized   ColorFormat = C.VK_FORMAT_A2R10G10B10_SNORM_PACK32
	A2R10G10B10UnsignedScaled     ColorFormat = C.VK_FORMAT_A2R10G10B10_USCALED_PACK32
	A2R10G10B10SignedScaled       ColorFormat = C.VK_FORMAT_A2R10G10B10_SSCALED_PACK32
	A2R10G10B10UnsignedInt        ColorFormat = C.VK_FORMAT_A2R10G10B10_UINT_PACK32
	A2R10G10B10SignedInt          ColorFormat = C.VK_FORMAT_A2R10G10B10_SINT_PACK32

	A2B10G10R10UnsignedNormalized ColorFormat = C.VK_FORMAT_A2B10G10R10_UNORM_PACK32
	A2B10G10R10SignedNormalized   ColorFormat = C.VK_FORMAT_A2B10G10R10_SNORM_PACK32
	A2B10G10R10UnsignedScaled     ColorFormat = C.VK_FORMAT_A2B10G10R10_USCALED_PACK32
	A2B10G10R10SignedScaled       ColorFormat = C.VK_FORMAT_A2B10G10R10_SSCALED_PACK32
	A2B10G10R10UnsignedInt        ColorFormat = C.VK_FORMAT_A2B10G10R10_UINT_PACK32
	A2B10G10R10SignedInt          ColorFormat = C.VK_FORMAT_A2B10G10R10_SINT_PACK32

	R16UnsignedNormalized ColorFormat = C.VK_FORMAT_R16_UNORM
	R16SignedNormalized   ColorFormat = C.VK_FORMAT_R16_SNORM
	R16UnsignedScaled     ColorFormat = C.VK_FORMAT_R16_USCALED
	R16SignedScaled       ColorFormat = C.VK_FORMAT_R16_SSCALED
	R16UnsignedInt        ColorFormat = C.VK_FORMAT_R16_UINT
	R16SignedInt          ColorFormat = C.VK_FORMAT_R16_SINT
	R16SignedFloat        ColorFormat = C.VK_FORMAT_R16_SFLOAT

	R16G16UnsignedNormalized ColorFormat = C.VK_FORMAT_R16G16_UNORM
	R16G16SignedNormalized   ColorFormat = C.VK_FORMAT_R16G16_SNORM
	R16G16UnsignedScaled     ColorFormat = C.VK_FORMAT_R16G16_USCALED
	R16G16SignedScaled       ColorFormat = C.VK_FORMAT_R16G16_SSCALED
	R16G16UnsignedInt        ColorFormat = C.VK_FORMAT_R16G16_UINT
	R16G16SignedInt          ColorFormat = C.VK_FORMAT_R16G16_SINT
	R16G16SignedFloat        ColorFormat = C.VK_FORMAT_R16G16_SFLOAT

	R16G16B16UnsignedNormalized ColorFormat = C.VK_FORMAT_R16G16B16_UNORM
	R16G16B16SignedNormalized   ColorFormat = C.VK_FORMAT_R16G16B16_SNORM
	R16G16B16UnsignedScaled     ColorFormat = C.VK_FORMAT_R16G16B16_USCALED
	R16G16B16SignedScaled       ColorFormat = C.VK_FORMAT_R16G16B16_SSCALED
	R16G16B16UnsignedInt        ColorFormat = C.VK_FORMAT_R16G16B16_UINT
	R16G16B16SignedInt          ColorFormat = C.VK_FORMAT_R16G16B16_SINT
	R16G16B16SignedFloat        ColorFormat = C.VK_FORMAT_R16G16B16_SFLOAT

	R16G16B16A16UnsignedNormalized ColorFormat = C.VK_FORMAT_R16G16B16A16_UNORM
	R16G16B16A16SignedNormalized   ColorFormat = C.VK_FORMAT_R16G16B16A16_SNORM
	R16G16B16A16UnsignedScaled     ColorFormat = C.VK_FORMAT_R16G16B16A16_USCALED
	R16G16B16A16SignedScaled       ColorFormat = C.VK_FORMAT_R16G16B16A16_SSCALED
	R16G16B16A16UnsignedInt        ColorFormat = C.VK_FORMAT_R16G16B16A16_UINT
	R16G16B16A16SignedInt          ColorFormat = C.VK_FORMAT_R16G16B16A16_SINT
	R16G16B16A16SignedFloat        ColorFormat = C.VK_FORMAT_R16G16B16A16_SFLOAT

	R32UnsignedInt          ColorFormat = C.VK_FORMAT_R32_UINT
	R32SignedInt            ColorFormat = C.VK_FORMAT_R32_SINT
	R32SignedFloat          ColorFormat = C.VK_FORMAT_R32_SFLOAT
	R32G32UnsignedInt       ColorFormat = C.VK_FORMAT_R32G32_UINT
	R32G32SignedInt         ColorFormat = C.VK_FORMAT_R32G32_SINT
	R32G32SignedFloat       ColorFormat = C.VK_FORMAT_R32G32_SFLOAT
	R32G32B32UnsignedInt    ColorFormat = C.VK_FORMAT_R32G32B32_UINT
	R32G32B32SignedInt      ColorFormat = C.VK_FORMAT_R32G32B32_SINT
	R32G32B32SignedFloat    ColorFormat = C.VK_FORMAT_R32G32B32_SFLOAT
	R32G32B32A32UnsignedInt ColorFormat = C.VK_FORMAT_R32G32B32A32_UINT
	R32G32B32A32SignedInt   ColorFormat = C.VK_FORMAT_R32G32B32A32_SINT
	R32G32B32A32SignedFloat ColorFormat = C.VK_FORMAT_R32G32B32A32_SFLOAT

	R64UnsignedInt          ColorFormat = C.VK_FORMAT_R64_UINT
	R64SignedInt            ColorFormat = C.VK_FORMAT_R64_SINT
	R64SignedFloat          ColorFormat = C.VK_FORMAT_R64_SFLOAT
	R64G64UnsignedInt       ColorFormat = C.VK_FORMAT_R64G64_UINT
	R64G64SignedInt         ColorFormat = C.VK_FORMAT_R64G64_SINT
	R64G64SignedFloat       ColorFormat = C.VK_FORMAT_R64G64_SFLOAT
	R64G64B64UnsignedInt    ColorFormat = C.VK_FORMAT_R64G64B64_UINT
	R64G64B64SignedInt      ColorFormat = C.VK_FORMAT_R64G64B64_SINT
	R64G64B64SignedFloat    ColorFormat = C.VK_FORMAT_R64G64B64_SFLOAT
	R64G64B64A64UnsignedInt ColorFormat = C.VK_FORMAT_R64G64B64A64_UINT
	R64G64B64A64SignedInt   ColorFormat = C.VK_FORMAT_R64G64B64A64_SINT
	R64G64B64A64SignedFloat ColorFormat = C.VK_FORMAT_R64G64B64A64_SFLOAT

	B10G11R11UnsignedFloat  ColorFormat = C.VK_FORMAT_B10G11R11_UFLOAT_PACK32
	E5B9G9R9UnsignedFloat   ColorFormat = C.VK_FORMAT_E5B9G9R9_UFLOAT_PACK32
	D16UnsignedNormalized   ColorFormat = C.VK_FORMAT_D16_UNORM
	D24X8UnsignedNormalized ColorFormat = C.VK_FORMAT_X8_D24_UNORM_PACK32
	D32SignedFloat          ColorFormat = C.VK_FORMAT_D32_SFLOAT
	S8UnsignedInt           ColorFormat = C.VK_FORMAT_S8_UINT

	D16UnsignedNormalizedS8UnsignedInt ColorFormat = C.VK_FORMAT_D16_UNORM_S8_UINT
	D24UnsignedNormalizedS8UnsignedInt ColorFormat = C.VK_FORMAT_D24_UNORM_S8_UINT
	D32SignedFloatS8UnsignedInt        ColorFormat = C.VK_FORMAT_D32_SFLOAT_S8_UINT

	BC1_RGBUnsignedNormalized  ColorFormat = C.VK_FORMAT_BC1_RGB_UNORM_BLOCK
	BC1_RGBsRGB                ColorFormat = C.VK_FORMAT_BC1_RGB_SRGB_BLOCK
	BC1_RGBAUnsignedNormalized ColorFormat = C.VK_FORMAT_BC1_RGBA_UNORM_BLOCK
	BC1_RGBAsRGB               ColorFormat = C.VK_FORMAT_BC1_RGBA_SRGB_BLOCK

	BC2_UnsignedNormalized ColorFormat = C.VK_FORMAT_BC2_UNORM_BLOCK
	BC2_sRGB               ColorFormat = C.VK_FORMAT_BC2_SRGB_BLOCK

	BC3_UnsignedNormalized ColorFormat = C.VK_FORMAT_BC3_UNORM_BLOCK
	BC3_sRGB               ColorFormat = C.VK_FORMAT_BC3_SRGB_BLOCK

	BC4_UnsignedNormalized ColorFormat = C.VK_FORMAT_BC4_UNORM_BLOCK
	BC4_SignedNormalized   ColorFormat = C.VK_FORMAT_BC4_SNORM_BLOCK

	BC5_UnsignedNormalized ColorFormat = C.VK_FORMAT_BC5_UNORM_BLOCK
	BC5_SignedNormalized   ColorFormat = C.VK_FORMAT_BC5_SNORM_BLOCK

	BC6_UnsignedFloat ColorFormat = C.VK_FORMAT_BC6H_UFLOAT_BLOCK
	BC6_SignedFloat   ColorFormat = C.VK_FORMAT_BC6H_SFLOAT_BLOCK

	BC7_UnsignedNormalized ColorFormat = C.VK_FORMAT_BC7_UNORM_BLOCK
	BC7_sRGB               ColorFormat = C.VK_FORMAT_BC7_SRGB_BLOCK

	ETC2_R8G8B8UnsignedNormalized   ColorFormat = C.VK_FORMAT_ETC2_R8G8B8_UNORM_BLOCK
	ETC2_R8G8B8sRGB                 ColorFormat = C.VK_FORMAT_ETC2_R8G8B8_SRGB_BLOCK
	ETC2_R8G8B8A1UnsignedNormalized ColorFormat = C.VK_FORMAT_ETC2_R8G8B8A1_UNORM_BLOCK
	ETC2_R8G8B8A1sRGB               ColorFormat = C.VK_FORMAT_ETC2_R8G8B8A1_SRGB_BLOCK
	ETC2_R8G8B8A8UnsignedNormalized ColorFormat = C.VK_FORMAT_ETC2_R8G8B8A8_UNORM_BLOCK
	ETC2_R8G8B8A8sRGB               ColorFormat = C.VK_FORMAT_ETC2_R8G8B8A8_SRGB_BLOCK

	EAC_R11UnsignedNormalized    ColorFormat = C.VK_FORMAT_EAC_R11_UNORM_BLOCK
	EAC_R11SignedNormalized      ColorFormat = C.VK_FORMAT_EAC_R11_SNORM_BLOCK
	EAC_R11G11UnsignedNormalized ColorFormat = C.VK_FORMAT_EAC_R11G11_UNORM_BLOCK
	EAC_R11G11SignedNormalized   ColorFormat = C.VK_FORMAT_EAC_R11G11_SNORM_BLOCK

	ASTC4x4_UnsignedNormalized   ColorFormat = C.VK_FORMAT_ASTC_4x4_UNORM_BLOCK
	ASTC4x4_sRGB                 ColorFormat = C.VK_FORMAT_ASTC_4x4_SRGB_BLOCK
	ASTC5x4_UnsignedNormalized   ColorFormat = C.VK_FORMAT_ASTC_5x4_UNORM_BLOCK
	ASTC5x4_sRGB                 ColorFormat = C.VK_FORMAT_ASTC_5x4_SRGB_BLOCK
	ASTC5x5_UnsignedNormalized   ColorFormat = C.VK_FORMAT_ASTC_5x5_UNORM_BLOCK
	ASTC5x5_sRGB                 ColorFormat = C.VK_FORMAT_ASTC_5x5_SRGB_BLOCK
	ASTC6x5_UnsignedNormalized   ColorFormat = C.VK_FORMAT_ASTC_6x5_UNORM_BLOCK
	ASTC6x5_sRGB                 ColorFormat = C.VK_FORMAT_ASTC_6x5_SRGB_BLOCK
	ASTC6x6_UnsignedNormalized   ColorFormat = C.VK_FORMAT_ASTC_6x6_UNORM_BLOCK
	ASTC6x6_sRGB                 ColorFormat = C.VK_FORMAT_ASTC_6x6_SRGB_BLOCK
	ASTC8x5_UnsignedNormalized   ColorFormat = C.VK_FORMAT_ASTC_8x5_UNORM_BLOCK
	ASTC8x5_sRGB                 ColorFormat = C.VK_FORMAT_ASTC_8x5_SRGB_BLOCK
	ASTC8x6_UnsignedNormalized   ColorFormat = C.VK_FORMAT_ASTC_8x6_UNORM_BLOCK
	ASTC8x6_sRGB                 ColorFormat = C.VK_FORMAT_ASTC_8x6_SRGB_BLOCK
	ASTC8x8_UnsignedNormalized   ColorFormat = C.VK_FORMAT_ASTC_8x8_UNORM_BLOCK
	ASTC8x8_sRGB                 ColorFormat = C.VK_FORMAT_ASTC_8x8_SRGB_BLOCK
	ASTC10x5_UnsignedNormalized  ColorFormat = C.VK_FORMAT_ASTC_10x5_UNORM_BLOCK
	ASTC10x5_sRGB                ColorFormat = C.VK_FORMAT_ASTC_10x5_SRGB_BLOCK
	ASTC10x6_UnsignedNormalized  ColorFormat = C.VK_FORMAT_ASTC_10x6_UNORM_BLOCK
	ASTC10x6_sRGB                ColorFormat = C.VK_FORMAT_ASTC_10x6_SRGB_BLOCK
	ASTC10x8_UnsignedNormalized  ColorFormat = C.VK_FORMAT_ASTC_10x8_UNORM_BLOCK
	ASTC10x8_sRGB                ColorFormat = C.VK_FORMAT_ASTC_10x8_SRGB_BLOCK
	ASTC10x10_UnsignedNormalized ColorFormat = C.VK_FORMAT_ASTC_10x10_UNORM_BLOCK
	ASTC10x10_sRGB               ColorFormat = C.VK_FORMAT_ASTC_10x10_SRGB_BLOCK
	ASTC12x10_UnsignedNormalized ColorFormat = C.VK_FORMAT_ASTC_12x10_UNORM_BLOCK
	ASTC12x10_sRGB               ColorFormat = C.VK_FORMAT_ASTC_12x10_SRGB_BLOCK
	ASTC12x12_UnsignedNormalized ColorFormat = C.VK_FORMAT_ASTC_12x12_UNORM_BLOCK
	ASTC12x12_sRGB               ColorFormat = C.VK_FORMAT_ASTC_12x12_SRGB_BLOCK

	G8B8G8R8_4x2x2UnsignedNormalized ColorFormat = C.VK_FORMAT_G8B8G8R8_422_UNORM
	B8G8R8G8_4x2x2UnsignedNormalized ColorFormat = C.VK_FORMAT_B8G8R8G8_422_UNORM

	G8B8R8_3Plane_4x2x0UnsignedNormalized ColorFormat = C.VK_FORMAT_G8_B8_R8_3PLANE_420_UNORM
	G8B8R8_2Plane_4x2x0UnsignedNormalized ColorFormat = C.VK_FORMAT_G8_B8R8_2PLANE_420_UNORM
	G8B8R8_3Plane_4x2x2UnsignedNormalized ColorFormat = C.VK_FORMAT_G8_B8_R8_3PLANE_422_UNORM
	G8B8R8_2Plane_4x2x2UnsignedNormalized ColorFormat = C.VK_FORMAT_G8_B8R8_2PLANE_422_UNORM
	G8B8R8_3Plane_4x4x4UnsignedNormalized ColorFormat = C.VK_FORMAT_G8_B8_R8_3PLANE_444_UNORM

	R10X6UnsignedNormalized                        ColorFormat = C.VK_FORMAT_R10X6_UNORM_PACK16
	R10X6G10X6UnsignedNormalized                   ColorFormat = C.VK_FORMAT_R10X6G10X6_UNORM_2PACK16
	R10X6G10X6B10X6A10X6UnsignedNormalized         ColorFormat = C.VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16
	G10X6B10X6G10X6R10X6_4x2x2UnsignedNormalized   ColorFormat = C.VK_FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16
	B10X6G10X6R10X6G10X6_4x2x2UnsignedNormalized   ColorFormat = C.VK_FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16
	G10X6G10X6R10X6_3Plane_4x2x0UnsignedNormalized ColorFormat = C.VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16
	G10X6B10X6R10X6_2Plane_4x2x0UnsignedNormalized ColorFormat = C.VK_FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16
	G10X6B10X6R10X6_3Plane_4x2x2UnsignedNormalized ColorFormat = C.VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16
	G10X6B10X6R10X6_2Plane_4x2x2UnsignedNormalized ColorFormat = C.VK_FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16
	G10X6B10X6R10X6_3Plane_4x4x4UnsignedNormalized ColorFormat = C.VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16

	R12X4UnsignedNormalized                        ColorFormat = C.VK_FORMAT_R12X4_UNORM_PACK16
	R12X4G12X4UnsignedNormalized                   ColorFormat = C.VK_FORMAT_R12X4G12X4_UNORM_2PACK16
	R12X4G12X4B12X4A12X4UnsignedNormalized         ColorFormat = C.VK_FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16
	G12X4B12X4G12X4R12X4_4x2x2UnsignedNormalized   ColorFormat = C.VK_FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16
	B12X4G12X4R12X4G12X4_4x2x2UnsignedNormalized   ColorFormat = C.VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16
	G12X4B12X4R12X4_3Plane_4x2x0UnsignedNormalized ColorFormat = C.VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16
	G12X4B12X4R12X4_2Plane_4x2x0UnsignedNormalized ColorFormat = C.VK_FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16
	G12X4B12X4R12X4_3Plane_4x2x2UnsignedNormalized ColorFormat = C.VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16
	G12X4B12X4R12X4_2Plane_4x2x2UnsignedNormalized ColorFormat = C.VK_FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16
	G12X4B12X4R12X4_3Plane_4x4x4UnsignedNormalized ColorFormat = C.VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16

	G16B16G16R16_4x2x2UnsignedNormalized     ColorFormat = C.VK_FORMAT_G16B16G16R16_422_UNORM
	B16G16R16G16_4x2x2UnsignedNormalized     ColorFormat = C.VK_FORMAT_B16G16R16G16_422_UNORM
	G16B16R16_3Plane_4x2x0UnsignedNormalized ColorFormat = C.VK_FORMAT_G16_B16_R16_3PLANE_420_UNORM
	G16B16R16_2Plane_4x2x0UnsignedNormalized ColorFormat = C.VK_FORMAT_G16_B16R16_2PLANE_420_UNORM
	G16B16R16_3Plane_4x2x2UnsignedNormalized ColorFormat = C.VK_FORMAT_G16_B16_R16_3PLANE_422_UNORM
	G16B16R16_2Plane_4x2x2UnsignedNormalized ColorFormat = C.VK_FORMAT_G16_B16R16_2PLANE_422_UNORM
	G16B16R16_3Plane_4x4x4UnsignedNormalized ColorFormat = C.VK_FORMAT_G16_B16_R16_3PLANE_444_UNORM

	PVRTC1_2BPP_UnsignedNormalized ColorFormat = C.VK_FORMAT_PVRTC1_2BPP_UNORM_BLOCK_IMG
	PVRTC1_4BPP_UnsignedNormalized ColorFormat = C.VK_FORMAT_PVRTC1_4BPP_UNORM_BLOCK_IMG
	PVRTC1_2BPP_sRGB               ColorFormat = C.VK_FORMAT_PVRTC1_2BPP_SRGB_BLOCK_IMG
	PVRTC1_4BPP_sRGB               ColorFormat = C.VK_FORMAT_PVRTC1_4BPP_SRGB_BLOCK_IMG

	PVRTC2_2BPP_UnsignedNormalized ColorFormat = C.VK_FORMAT_PVRTC2_2BPP_UNORM_BLOCK_IMG
	PVRTC2_4BPP_UnsignedNormalized ColorFormat = C.VK_FORMAT_PVRTC2_4BPP_UNORM_BLOCK_IMG
	PVRTC2_2BPP_sRGB               ColorFormat = C.VK_FORMAT_PVRTC2_2BPP_SRGB_BLOCK_IMG
	PVRTC2_4BPP_sRGB               ColorFormat = C.VK_FORMAT_PVRTC2_4BPP_SRGB_BLOCK_IMG

	ASTC4x4_SignedFloat   ColorFormat = C.VK_FORMAT_ASTC_4x4_SFLOAT_BLOCK_EXT
	ASTC5x4_SignedFloat   ColorFormat = C.VK_FORMAT_ASTC_5x4_SFLOAT_BLOCK_EXT
	ASTC5x5_SignedFloat   ColorFormat = C.VK_FORMAT_ASTC_5x5_SFLOAT_BLOCK_EXT
	ASTC6x5_SignedFloat   ColorFormat = C.VK_FORMAT_ASTC_6x5_SFLOAT_BLOCK_EXT
	ASTC6x6_SignedFloat   ColorFormat = C.VK_FORMAT_ASTC_6x6_SFLOAT_BLOCK_EXT
	ASTC8x5_SignedFloat   ColorFormat = C.VK_FORMAT_ASTC_8x5_SFLOAT_BLOCK_EXT
	ASTC8x6_SignedFloat   ColorFormat = C.VK_FORMAT_ASTC_8x6_SFLOAT_BLOCK_EXT
	ASTC8x8_SignedFloat   ColorFormat = C.VK_FORMAT_ASTC_8x8_SFLOAT_BLOCK_EXT
	ASTC10x5_SignedFloat  ColorFormat = C.VK_FORMAT_ASTC_10x5_SFLOAT_BLOCK_EXT
	ASTC10x6_SignedFloat  ColorFormat = C.VK_FORMAT_ASTC_10x6_SFLOAT_BLOCK_EXT
	ASTC10x8_SignedFloat  ColorFormat = C.VK_FORMAT_ASTC_10x8_SFLOAT_BLOCK_EXT
	ASTC10x10_SignedFloat ColorFormat = C.VK_FORMAT_ASTC_10x10_SFLOAT_BLOCK_EXT
	ASTC12x10_SignedFloat ColorFormat = C.VK_FORMAT_ASTC_12x10_SFLOAT_BLOCK_EXT
	ASTC12x12_SignedFloat ColorFormat = C.VK_FORMAT_ASTC_12x12_SFLOAT_BLOCK_EXT

	G8B8R8_2Plane_4x4x4UnsignedNormalized          ColorFormat = C.VK_FORMAT_G8_B8R8_2PLANE_444_UNORM_EXT
	G10X6B10X6R10X6_2Plane_4x4x4UnsignedNormalized ColorFormat = C.VK_FORMAT_G10X6_B10X6R10X6_2PLANE_444_UNORM_3PACK16_EXT
	G12X4B12X4R12X4_2Plane_4x4x4UnsignedNormalized ColorFormat = C.VK_FORMAT_G12X4_B12X4R12X4_2PLANE_444_UNORM_3PACK16_EXT
	G16B16R16_2Plane_4x4x4UnsignedNormalized       ColorFormat = C.VK_FORMAT_G16_B16R16_2PLANE_444_UNORM_EXT
	A4R4G4B4UnsignedNormalized                     ColorFormat = C.VK_FORMAT_A4R4G4B4_UNORM_PACK16_EXT
	A4B4G4R4UnsignedNormalized                     ColorFormat = C.VK_FORMAT_A4B4G4R4_UNORM_PACK16_EXT
)

var colorFormatToString = map[ColorFormat]string{
	Undefined:              "Undefined",
	R4G4UnsignedNormalized: "R4G4 Unsigned Normalized",

	R4G4B4A4UnsignedNormalized: "R4G4B4A4 Unsigned Normalized",
	B4G4R4A4UnsignedNormalized: "B4G4R4A4 Unsigned Normalized",
	R5G6B5UnsignedNormalized:   "R5G6B5 Unsigned Normalized",
	B5G6R5UnsignedNormalized:   "B5G6R5 Unsigned Normalized",
	R5G5B5A1UnsignedNormalized: "R5G5B5A1 Unsigned Normalized",
	B5G5R5A1UnsignedNormalized: "B5G5R5A1 Unsigned Normalized",
	A1R5G5B5UnsignedNormalized: "A1R5G5B5 Unsigned Normalized",

	R8UnsignedNormalized: "R8 Unsigned Normalized",
	R8SignedNormalized:   "R8 Signed Normalized",
	R8UnsignedScaled:     "R8 Unsigned Scaled",
	R8SignedScaled:       "R8 Signed Scaled",
	R8UnsignedInt:        "R8 Unsigned Int",
	R8SignedInt:          "R8 Signed Int",
	R8SRGB:               "R8 sRGB",

	R8G8UnsignedNormalized: "R8G8 Unsigned Normalized",
	R8G8SignedNormalized:   "R8G8 Signed Normalized",
	R8G8UnsignedScaled:     "R8G8 Unsigned Scaled",
	R8G8SignedScaled:       "R8G8 Signed Scaled",
	R8G8UnsignedInt:        "R8G8 Unsigned Int",
	R8G8SignedInt:          "R8G8 Signed Int",
	R8G8SRGB:               "R8G8 sRGB",

	R8G8B8UnsignedNormalized: "R8G8B8 Unsigned Normalized",
	R8G8B8SignedNormalized:   "R8G8B8 Signed Normalized",
	R8G8B8UnsignedScaled:     "R8G8B8 Unsigned Scaled",
	R8G8B8SignedScaled:       "R8G8B8 Signed Scaled",
	R8G8B8UnsignedInt:        "R8G8B8 Unsigned Int",
	R8G8B8SignedInt:          "R8G8B8 Signed Int",
	R8G8B8SRGB:               "R8G8B8 sRGB",

	B8G8R8UnsignedNormalized: "B8G8R8 Unsigned Normalized",
	B8G8R8SignedNormalized:   "B8G8R8 Signed Normalized",
	B8G8R8UnsignedScaled:     "B8G8R8 Unsigned Scaled",
	B8G8R8SignedScaled:       "B8G8R8 Signed Scaled",
	B8G8R8UnsignedInt:        "B8G8R8 Unsigned Int",
	B8G8R8SignedInt:          "B8G8R8 Signed Int",
	B8G8R8SRGB:               "B8G8R8 sRGB",

	R8G8B8A8UnsignedNormalized: "R8G8B8A8 Unsigned Normalized",
	R8G8B8A8SignedNormalized:   "R8G8B8A8 Signed Normalized",
	R8G8B8A8UnsignedScaled:     "R8G8B8A8 Unsigned Scaled",
	R8G8B8A8SignedScaled:       "R8G8B8A8 Signed Scaled",
	R8G8B8A8UnsignedInt:        "R8G8B8A8 Unsigned Int",
	R8G8B8A8SignedInt:          "R8G8B8A8 Signed Int",
	R8G8B8A8SRGB:               "R8G8B8A8 sRGB",

	B8G8R8A8UnsignedNormalized: "B8G8R8A8 Unsigned Normalized",
	B8G8R8A8SignedNormalized:   "B8G8R8A8 Signed Normalized",
	B8G8R8A8UnsignedScaled:     "B8G8R8A8 Unsigned Scaled",
	B8G8R8A8SignedScaled:       "B8G8R8A8 Signed Scaled",
	B8G8R8A8UnsignedInt:        "B8G8R8A8 Unsigned Int",
	B8G8R8A8SignedInt:          "B8G8R8A8 Signed Int",
	B8G8R8A8SRGB:               "B8G8R8A8 sRGB",

	A8B8G8R8UnsignedNormalized: "A8B8G8R8 Unsigned Normalized",
	A8B8G8R8SignedNormalized:   "A8B8G8R8 Signed Normalized",
	A8B8G8R8UnsignedScaled:     "A8B8G8R8 Unsigned Scaled",
	A8B8G8R8SignedScaled:       "A8B8G8R8 Signed Scaled",
	A8B8G8R8UnsignedInt:        "A8B8G8R8 Unsigned Int",
	A8B8G8R8SignedInt:          "A8B8G8R8 Signed Int",
	A8B8G8R8SRGB:               "A8B8G8R8 sRGB",

	A2R10G10B10UnsignedNormalized: "A2R10G10B10 Unsigned Normalized",
	A2R10G10B10SignedNormalized:   "A2R10G10B10 Signed Normalized",
	A2R10G10B10UnsignedScaled:     "A2R10G10B10 Unsigned Scaled",
	A2R10G10B10SignedScaled:       "A2R10G10B10 Signed Scaled",
	A2R10G10B10UnsignedInt:        "A2R10G10B10 Unsigned Int",
	A2R10G10B10SignedInt:          "A2R10G10B10 Signed Int",

	A2B10G10R10UnsignedNormalized: "A2B10G10R10 Unsigned Normalized",
	A2B10G10R10SignedNormalized:   "A2B10G10R10 Signed Normalized",
	A2B10G10R10UnsignedScaled:     "A2B10G10R10 Unsigned Scaled",
	A2B10G10R10SignedScaled:       "A2B10G10R10 Signed Scaled",
	A2B10G10R10UnsignedInt:        "A2B10G10R10 Unsigned Int",
	A2B10G10R10SignedInt:          "A2B10G10R10 Signed Int",

	R16UnsignedNormalized: "R16 Unsigned Normalized",
	R16SignedNormalized:   "R16 Signed Normalized",
	R16UnsignedScaled:     "R16 Unsigned Scaled",
	R16SignedScaled:       "R16 Signed Scaled",
	R16UnsignedInt:        "R16 Unsigned Int",
	R16SignedInt:          "R16 Signed Int",
	R16SignedFloat:        "R16 Signed Float",

	R16G16UnsignedNormalized: "R16G16 Unsigned Normalized",
	R16G16SignedNormalized:   "R16G16 Signed Normalized",
	R16G16UnsignedScaled:     "R16G16 Unsigned Scaled",
	R16G16SignedScaled:       "R16G16 Signed Scaled",
	R16G16UnsignedInt:        "R16G16 Unsigned Int",
	R16G16SignedInt:          "R16G16 Signed Int",
	R16G16SignedFloat:        "R16G16 Signed Float",

	R16G16B16UnsignedNormalized: "R16G16B16 Unsigned Normalized",
	R16G16B16SignedNormalized:   "R16G16B16 Signed Normalized",
	R16G16B16UnsignedScaled:     "R16G16B16 Unsigned Scaled",
	R16G16B16SignedScaled:       "R16G16B16 Signed Scaled",
	R16G16B16UnsignedInt:        "R16G16B16 Unsigned Int",
	R16G16B16SignedInt:          "R16G16B16 Signed Int",
	R16G16B16SignedFloat:        "R16G16B16 Signed Float",

	R16G16B16A16UnsignedNormalized: "R16G16B16A16 Unsigned Normalized",
	R16G16B16A16SignedNormalized:   "R16G16B16A16 Signed Normalized",
	R16G16B16A16UnsignedScaled:     "R16G16B16A16 Unsigned Scaled",
	R16G16B16A16SignedScaled:       "R16G16B16A16 Signed Scaled",
	R16G16B16A16UnsignedInt:        "R16G16B16A16 Unsigned Int",
	R16G16B16A16SignedInt:          "R16G16B16A16 Signed Int",
	R16G16B16A16SignedFloat:        "R16G16B16A16 Signed Float",

	R32UnsignedInt:          "R32 Unsigned Int",
	R32SignedInt:            "R32 Signed Int",
	R32SignedFloat:          "R32 Signed Float",
	R32G32UnsignedInt:       "R32G32 Unsigned Int",
	R32G32SignedInt:         "R32G32 Signed Int",
	R32G32SignedFloat:       "R32G32 Signed Float",
	R32G32B32UnsignedInt:    "R32G32B32 Unsigned Int",
	R32G32B32SignedInt:      "R32G32B32 Signed Int",
	R32G32B32SignedFloat:    "R32G32B32 Signed Float",
	R32G32B32A32UnsignedInt: "R32G32B32A32 Unsigned Int",
	R32G32B32A32SignedInt:   "R32G32B32A32 Signed Int",
	R32G32B32A32SignedFloat: "R32G32B32A32 Signed Float",

	R64UnsignedInt:          "R64 Unsigned Int",
	R64SignedInt:            "R64 Signed Int",
	R64SignedFloat:          "R64 Signed Float",
	R64G64UnsignedInt:       "R64G64 Unsigned Int",
	R64G64SignedInt:         "R64G64 Signed Int",
	R64G64SignedFloat:       "R64G64 Signed Float",
	R64G64B64UnsignedInt:    "R64G64B64 Unsigned Int",
	R64G64B64SignedInt:      "R64G64B64 Signed Int",
	R64G64B64SignedFloat:    "R64G64B64 Signed Float",
	R64G64B64A64UnsignedInt: "R64G64B64A64 Unsigned Int",
	R64G64B64A64SignedInt:   "R64G64B64A64 Signed Int",
	R64G64B64A64SignedFloat: "R64G64B64A64 Signed Float",

	B10G11R11UnsignedFloat:  "B10G11R11 Unsigned Float",
	E5B9G9R9UnsignedFloat:   "E5B9G9R9 Unsigned Float",
	D16UnsignedNormalized:   "D16 Unsigned Normalized",
	D24X8UnsignedNormalized: "D24X8 Unsigned Normalized",
	D32SignedFloat:          "D32 Signed Float",
	S8UnsignedInt:           "S8 Unsigned Int",

	D16UnsignedNormalizedS8UnsignedInt: "D16 Unsigned Normalized S8 Unsigned Int",
	D24UnsignedNormalizedS8UnsignedInt: "D24 Unsigned Normalized S8 Unsigned Int",
	D32SignedFloatS8UnsignedInt:        "D32 Signed Float S8 Unsigned Int",

	BC1_RGBUnsignedNormalized:  "BC1-Compressed -Compressed RGB Unsigned Normalized",
	BC1_RGBsRGB:                "BC1-Compressed -Compressed RGB sRGB",
	BC1_RGBAUnsignedNormalized: "BC1-Compressed -Compressed RGBA Unsigned Normalized",
	BC1_RGBAsRGB:               "BC1-Compressed RGBA sRGB",

	BC2_UnsignedNormalized: "BC2-Compressed Unsigned Normalized",
	BC2_sRGB:               "BC2-Compressed sRGB",

	BC3_UnsignedNormalized: "BC3-Compressed Unsigned Normalized",
	BC3_sRGB:               "BC3-Compressed sRGB",

	BC4_UnsignedNormalized: "BC4-Compressed Unsigned Normalized",
	BC4_SignedNormalized:   "BC4-Compressed Signed Normalized",

	BC5_UnsignedNormalized: "BC5-Compressed Unsigned Normalized",
	BC5_SignedNormalized:   "BC5-Compressed Signed Normalized",

	BC6_UnsignedFloat: "BC6-Compressed Unsigned Float",
	BC6_SignedFloat:   "BC6-Compressed Signed Float",

	BC7_UnsignedNormalized: "BC7-Compressed Unsigned Normalized",
	BC7_sRGB:               "BC7-Compressed sRGB",

	ETC2_R8G8B8UnsignedNormalized:   "ETC2-Compressed R8G8B8 Unsigned Normalized",
	ETC2_R8G8B8sRGB:                 "ETC2-Compressed R8G8B8 sRGB",
	ETC2_R8G8B8A1UnsignedNormalized: "ETC2-Compressed R8G8B8A1 Unsigned Normalized",
	ETC2_R8G8B8A1sRGB:               "ETC2-Compressed R8G8B8A1 sRGB",
	ETC2_R8G8B8A8UnsignedNormalized: "ETC2-Compressed R8G8B8A8 Unsigned Normalized",
	ETC2_R8G8B8A8sRGB:               "ETC2-Compressed R8G8B8A8 sRGB",

	EAC_R11UnsignedNormalized:    "EAC-Compressed R11 Unsigned Normalized",
	EAC_R11SignedNormalized:      "EAC-Compressed R11 Signed Normalized",
	EAC_R11G11UnsignedNormalized: "EAC-Compressed R11G11 Unsigned Normalized",
	EAC_R11G11SignedNormalized:   "EAC-Compressed R11G11 Signed Normalized",

	ASTC4x4_UnsignedNormalized:   "ASTC-Compressed (4x4) Unsigned Normalized",
	ASTC4x4_sRGB:                 "ASTC-Compressed (4x4) sRGB",
	ASTC5x4_UnsignedNormalized:   "ASTC-Compressed (5x4) Unsigned Normalized",
	ASTC5x4_sRGB:                 "ASTC-Compressed (5x4) sRGB",
	ASTC5x5_UnsignedNormalized:   "ASTC-Compressed (5x5) Unsigned Normalized",
	ASTC5x5_sRGB:                 "ASTC-Compressed (5x5) sRGB",
	ASTC6x5_UnsignedNormalized:   "ASTC-Compressed (6x5) Unsigned Normalized",
	ASTC6x5_sRGB:                 "ASTC-Compressed (6x5) sRGB",
	ASTC6x6_UnsignedNormalized:   "ASTC-Compressed (6x6) Unsigned Normalized",
	ASTC6x6_sRGB:                 "ASTC-Compressed (6x6) sRGB",
	ASTC8x5_UnsignedNormalized:   "ASTC-Compressed (8x5) Unsigned Normalized",
	ASTC8x5_sRGB:                 "ASTC-Compressed (8x5) sRGB",
	ASTC8x6_UnsignedNormalized:   "ASTC-Compressed (8x6) Unsigned Normalized",
	ASTC8x6_sRGB:                 "ASTC-Compressed (8x6) sRGB",
	ASTC8x8_UnsignedNormalized:   "ASTC-Compressed (8x8) Unsigned Normalized",
	ASTC8x8_sRGB:                 "ASTC-Compressed (8x8) sRGB",
	ASTC10x5_UnsignedNormalized:  "ASTC-Compressed (10x5) Unsigned Normalized",
	ASTC10x5_sRGB:                "ASTC-Compressed (10x5) sRGB",
	ASTC10x6_UnsignedNormalized:  "ASTC-Compressed (10x6) Unsigned Normalized",
	ASTC10x6_sRGB:                "ASTC-Compressed (10x6) sRGB",
	ASTC10x8_UnsignedNormalized:  "ASTC-Compressed (10x8) Unsigned Normalized",
	ASTC10x8_sRGB:                "ASTC-Compressed (10x8) sRGB",
	ASTC10x10_UnsignedNormalized: "ASTC-Compressed (10x10) Unsigned Normalized",
	ASTC10x10_sRGB:               "ASTC-Compressed (10x10) sRGB",
	ASTC12x10_UnsignedNormalized: "ASTC-Compressed (12x10) Unsigned Normalized",
	ASTC12x10_sRGB:               "ASTC-Compressed (12x10) sRGB",
	ASTC12x12_UnsignedNormalized: "ASTC-Compressed (12x12) Unsigned Normalized",
	ASTC12x12_sRGB:               "ASTC-Compressed (12x12) sRGB",

	G8B8G8R8_4x2x2UnsignedNormalized: "G8B8G8R8 (4:2:2) Unsigned Normalized",
	B8G8R8G8_4x2x2UnsignedNormalized: "B8G8R8G8 (4:2:2) Unsigned Normalized",

	G8B8R8_3Plane_4x2x0UnsignedNormalized: "G8B8R8 3-Plane (4:2:0) Unsigned Normalized",
	G8B8R8_2Plane_4x2x0UnsignedNormalized: "G8B8R8 2-Plane (4:2:0) Unsigned Normalized",
	G8B8R8_3Plane_4x2x2UnsignedNormalized: "G8B8R8 3-Plane (4:2:2) Unsigned Normalized",
	G8B8R8_2Plane_4x2x2UnsignedNormalized: "G8B8R8 2-Plane (4:2:2) Unsigned Normalized",
	G8B8R8_3Plane_4x4x4UnsignedNormalized: "G8B8R8 3-Plane (4:4:4) Unsigned Normalized",

	R10X6UnsignedNormalized:                        "R10X6 Unsigned Normalized",
	R10X6G10X6UnsignedNormalized:                   "R10X6G10X6 Unsigned Normalized",
	R10X6G10X6B10X6A10X6UnsignedNormalized:         "R10X6G10X6B10X6A10X6 Unsigned Normalized",
	G10X6B10X6G10X6R10X6_4x2x2UnsignedNormalized:   "G10X6B10X6G10X6R10X6 (4:2:2) Unsigned Normalized",
	B10X6G10X6R10X6G10X6_4x2x2UnsignedNormalized:   "B10X6G10X6R10X6G10X6 (4:2:2) Unsigned Normalized",
	G10X6G10X6R10X6_3Plane_4x2x0UnsignedNormalized: "G10X6G10X6R10X6 3-Plane (4:2:0) Unsigned Normalized",
	G10X6B10X6R10X6_2Plane_4x2x0UnsignedNormalized: "G10X6B10X6R10X6 2-Plane (4:2:0) Unsigned Normalized",
	G10X6B10X6R10X6_3Plane_4x2x2UnsignedNormalized: "G10X6B10X6R10X6 3-Plane (4:2:2) Unsigned Normalized",
	G10X6B10X6R10X6_2Plane_4x2x2UnsignedNormalized: "G10X6B10X6R10X6 2-Plane (4:2:2) Unsigned Normalized",
	G10X6B10X6R10X6_3Plane_4x4x4UnsignedNormalized: "G10X6B10X6R10X6 3-Plane (4:4:4) Unsigned Normalized",

	R12X4UnsignedNormalized:                        "R12X4 Unsigned Normalized",
	R12X4G12X4UnsignedNormalized:                   "R12X4G12X4 Unsigned Normalized",
	R12X4G12X4B12X4A12X4UnsignedNormalized:         "R12X4G12X4B12X4A12X4 Unsigned Normalized",
	G12X4B12X4G12X4R12X4_4x2x2UnsignedNormalized:   "G12X4B12X4G12X4R12X4 (4:2:2) Unsigned Normalized",
	B12X4G12X4R12X4G12X4_4x2x2UnsignedNormalized:   "B12X4G12X4R12X4G12X4 (4:2:2) Unsigned Normalized",
	G12X4B12X4R12X4_3Plane_4x2x0UnsignedNormalized: "G12X4B12X4R12X4 3-Plane (4:2:0) Unsigned Normalized",
	G12X4B12X4R12X4_2Plane_4x2x0UnsignedNormalized: "G12X4B12X4R12X4 2-Plane (4:2:0) Unsigned Normalized",
	G12X4B12X4R12X4_3Plane_4x2x2UnsignedNormalized: "G12X4B12X4R12X4 3-Plane (4:2:2) Unsigned Normalized",
	G12X4B12X4R12X4_2Plane_4x2x2UnsignedNormalized: "G12X4B12X4R12X4 2-Plane (4:2:2) Unsigned Normalized",
	G12X4B12X4R12X4_3Plane_4x4x4UnsignedNormalized: "G12X4B12X4R12X4 3-Plane (4:4:4) Unsigned Normalized",

	G16B16G16R16_4x2x2UnsignedNormalized:     "G16B16G16R16 (4:2:2) Unsigned Normalized",
	B16G16R16G16_4x2x2UnsignedNormalized:     "B16G16R16G16 (4:2:2) Unsigned Normalized",
	G16B16R16_3Plane_4x2x0UnsignedNormalized: "G16B16R16 3-Plane (4:2:0) Unsigned Normalized",
	G16B16R16_2Plane_4x2x0UnsignedNormalized: "G16B16R16 2-Plane (4:2:0) Unsigned Normalized",
	G16B16R16_3Plane_4x2x2UnsignedNormalized: "G16B16R16 3-Plane (4:2:2) Unsigned Normalized",
	G16B16R16_2Plane_4x2x2UnsignedNormalized: "G16B16R16 2-Plane (4:2:2) Unsigned Normalized",
	G16B16R16_3Plane_4x4x4UnsignedNormalized: "G16B16R16 3-Plane (4:4:4) Unsigned Normalized",

	PVRTC1_2BPP_UnsignedNormalized: "PVRTC1-Compressed (2 BPP) Unsigned Normalized",
	PVRTC1_4BPP_UnsignedNormalized: "PVRTC1-Compressed (4 BPP) Unsigned Normalized",
	PVRTC1_2BPP_sRGB:               "PVRTC1-Compressed (2 BPP) sRGB",
	PVRTC1_4BPP_sRGB:               "PVRTC1-Compressed (4 BPP) sRGB",

	PVRTC2_2BPP_UnsignedNormalized: "PVRTC2-Compressed (2 BPP) Unsigned Normalized",
	PVRTC2_4BPP_UnsignedNormalized: "PVRTC2-Compressed (4 BPP) Unsigned Normalized",
	PVRTC2_2BPP_sRGB:               "PVRTC2-Compressed (2 BPP) sRGB",
	PVRTC2_4BPP_sRGB:               "PVRTC2-Compressed (4 BPP) sRGB",

	ASTC4x4_SignedFloat:   "ASTC-Compressed (4x4) Signed Float",
	ASTC5x4_SignedFloat:   "ASTC-Compressed (5x4) Signed Float",
	ASTC5x5_SignedFloat:   "ASTC-Compressed (5x5) Signed Float",
	ASTC6x5_SignedFloat:   "ASTC-Compressed (6x5) Signed Float",
	ASTC6x6_SignedFloat:   "ASTC-Compressed (6x6) Signed Float",
	ASTC8x5_SignedFloat:   "ASTC-Compressed (8x5) Signed Float",
	ASTC8x6_SignedFloat:   "ASTC-Compressed (8x6) Signed Float",
	ASTC8x8_SignedFloat:   "ASTC-Compressed (8x8) Signed Float",
	ASTC10x5_SignedFloat:  "ASTC-Compressed (10x5) Signed Float",
	ASTC10x6_SignedFloat:  "ASTC-Compressed (10x6) Signed Float",
	ASTC10x8_SignedFloat:  "ASTC-Compressed (10x8) Signed Float",
	ASTC10x10_SignedFloat: "ASTC-Compressed (10x10) Signed Float",
	ASTC12x10_SignedFloat: "ASTC-Compressed (12x10) Signed Float",
	ASTC12x12_SignedFloat: "ASTC-Compressed (12x12) Signed Float",

	G8B8R8_2Plane_4x4x4UnsignedNormalized:          "G8B8R8 2-Plane (4:4:4) Unsigned Normalized",
	G10X6B10X6R10X6_2Plane_4x4x4UnsignedNormalized: "G10X6B10X6R10X6 2-Plane (4:4:4) Unsigned Normalized",
	G12X4B12X4R12X4_2Plane_4x4x4UnsignedNormalized: "G12X4B12X4R12X4 2-Plane (4:4:4) Unsigned Normalized",
	G16B16R16_2Plane_4x4x4UnsignedNormalized:       "G16B16R16 2-Plane (4:4:4) Unsigned Normalized",
	A4R4G4B4UnsignedNormalized:                     "A4R4G4B4 Unsigned Normalized",
	A4B4G4R4UnsignedNormalized:                     "A4B4G4R4 Unsigned Normalized",
}

func (f ColorFormat) String() string {
	return colorFormatToString[f]
}
