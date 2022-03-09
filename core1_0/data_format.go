package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

const (
	DataFormatUndefined              common.DataFormat = C.VK_FORMAT_UNDEFINED
	DataFormatR4G4UnsignedNormalized common.DataFormat = C.VK_FORMAT_R4G4_UNORM_PACK8

	DataFormatR4G4B4A4UnsignedNormalized common.DataFormat = C.VK_FORMAT_R4G4B4A4_UNORM_PACK16
	DataFormatB4G4R4A4UnsignedNormalized common.DataFormat = C.VK_FORMAT_B4G4R4A4_UNORM_PACK16
	DataFormatR5G6B5UnsignedNormalized   common.DataFormat = C.VK_FORMAT_R5G6B5_UNORM_PACK16
	DataFormatB5G6R5UnsignedNormalized   common.DataFormat = C.VK_FORMAT_B5G6R5_UNORM_PACK16
	DataFormatR5G5B5A1UnsignedNormalized common.DataFormat = C.VK_FORMAT_R5G5B5A1_UNORM_PACK16
	DataFormatB5G5R5A1UnsignedNormalized common.DataFormat = C.VK_FORMAT_B5G5R5A1_UNORM_PACK16
	DataFormatA1R5G5B5UnsignedNormalized common.DataFormat = C.VK_FORMAT_A1R5G5B5_UNORM_PACK16

	DataFormatR8UnsignedNormalized common.DataFormat = C.VK_FORMAT_R8_UNORM
	DataFormatR8SignedNormalized   common.DataFormat = C.VK_FORMAT_R8_SNORM
	DataFormatR8UnsignedScaled     common.DataFormat = C.VK_FORMAT_R8_USCALED
	DataFormatR8SignedScaled       common.DataFormat = C.VK_FORMAT_R8_SSCALED
	DataFormatR8UnsignedInt        common.DataFormat = C.VK_FORMAT_R8_UINT
	DataFormatR8SignedInt          common.DataFormat = C.VK_FORMAT_R8_SINT
	DataFormatR8SRGB               common.DataFormat = C.VK_FORMAT_R8_SRGB

	DataFormatR8G8UnsignedNormalized common.DataFormat = C.VK_FORMAT_R8G8_UNORM
	DataFormatR8G8SignedNormalized   common.DataFormat = C.VK_FORMAT_R8G8_SNORM
	DataFormatR8G8UnsignedScaled     common.DataFormat = C.VK_FORMAT_R8G8_USCALED
	DataFormatR8G8SignedScaled       common.DataFormat = C.VK_FORMAT_R8G8_SSCALED
	DataFormatR8G8UnsignedInt        common.DataFormat = C.VK_FORMAT_R8G8_UINT
	DataFormatR8G8SignedInt          common.DataFormat = C.VK_FORMAT_R8G8_SINT
	DataFormatR8G8SRGB               common.DataFormat = C.VK_FORMAT_R8G8_SRGB

	DataFormatR8G8B8UnsignedNormalized common.DataFormat = C.VK_FORMAT_R8G8B8_UNORM
	DataFormatR8G8B8SignedNormalized   common.DataFormat = C.VK_FORMAT_R8G8B8_SNORM
	DataFormatR8G8B8UnsignedScaled     common.DataFormat = C.VK_FORMAT_R8G8B8_USCALED
	DataFormatR8G8B8SignedScaled       common.DataFormat = C.VK_FORMAT_R8G8B8_SSCALED
	DataFormatR8G8B8UnsignedInt        common.DataFormat = C.VK_FORMAT_R8G8B8_UINT
	DataFormatR8G8B8SignedInt          common.DataFormat = C.VK_FORMAT_R8G8B8_SINT
	DataFormatR8G8B8SRGB               common.DataFormat = C.VK_FORMAT_R8G8B8_SRGB

	DataFormatB8G8R8UnsignedNormalized common.DataFormat = C.VK_FORMAT_B8G8R8_UNORM
	DataFormatB8G8R8SignedNormalized   common.DataFormat = C.VK_FORMAT_B8G8R8_SNORM
	DataFormatB8G8R8UnsignedScaled     common.DataFormat = C.VK_FORMAT_B8G8R8_USCALED
	DataFormatB8G8R8SignedScaled       common.DataFormat = C.VK_FORMAT_B8G8R8_SSCALED
	DataFormatB8G8R8UnsignedInt        common.DataFormat = C.VK_FORMAT_B8G8R8_UINT
	DataFormatB8G8R8SignedInt          common.DataFormat = C.VK_FORMAT_B8G8R8_SINT
	DataFormatB8G8R8SRGB               common.DataFormat = C.VK_FORMAT_B8G8R8_SRGB

	DataFormatR8G8B8A8UnsignedNormalized common.DataFormat = C.VK_FORMAT_R8G8B8A8_UNORM
	DataFormatR8G8B8A8SignedNormalized   common.DataFormat = C.VK_FORMAT_R8G8B8A8_SNORM
	DataFormatR8G8B8A8UnsignedScaled     common.DataFormat = C.VK_FORMAT_R8G8B8A8_USCALED
	DataFormatR8G8B8A8SignedScaled       common.DataFormat = C.VK_FORMAT_R8G8B8A8_SSCALED
	DataFormatR8G8B8A8UnsignedInt        common.DataFormat = C.VK_FORMAT_R8G8B8A8_UINT
	DataFormatR8G8B8A8SignedInt          common.DataFormat = C.VK_FORMAT_R8G8B8A8_SINT
	DataFormatR8G8B8A8SRGB               common.DataFormat = C.VK_FORMAT_R8G8B8A8_SRGB

	DataFormatB8G8R8A8UnsignedNormalized common.DataFormat = C.VK_FORMAT_B8G8R8A8_UNORM
	DataFormatB8G8R8A8SignedNormalized   common.DataFormat = C.VK_FORMAT_B8G8R8A8_SNORM
	DataFormatB8G8R8A8UnsignedScaled     common.DataFormat = C.VK_FORMAT_B8G8R8A8_USCALED
	DataFormatB8G8R8A8SignedScaled       common.DataFormat = C.VK_FORMAT_B8G8R8A8_SSCALED
	DataFormatB8G8R8A8UnsignedInt        common.DataFormat = C.VK_FORMAT_B8G8R8A8_UINT
	DataFormatB8G8R8A8SignedInt          common.DataFormat = C.VK_FORMAT_B8G8R8A8_SINT
	DataFormatB8G8R8A8SRGB               common.DataFormat = C.VK_FORMAT_B8G8R8A8_SRGB

	DataFormatA8B8G8R8UnsignedNormalized common.DataFormat = C.VK_FORMAT_A8B8G8R8_UNORM_PACK32
	DataFormatA8B8G8R8SignedNormalized   common.DataFormat = C.VK_FORMAT_A8B8G8R8_SNORM_PACK32
	DataFormatA8B8G8R8UnsignedScaled     common.DataFormat = C.VK_FORMAT_A8B8G8R8_USCALED_PACK32
	DataFormatA8B8G8R8SignedScaled       common.DataFormat = C.VK_FORMAT_A8B8G8R8_SSCALED_PACK32
	DataFormatA8B8G8R8UnsignedInt        common.DataFormat = C.VK_FORMAT_A8B8G8R8_UINT_PACK32
	DataFormatA8B8G8R8SignedInt          common.DataFormat = C.VK_FORMAT_A8B8G8R8_SINT_PACK32
	DataFormatA8B8G8R8SRGB               common.DataFormat = C.VK_FORMAT_A8B8G8R8_SRGB_PACK32

	DataFormatA2R10G10B10UnsignedNormalized common.DataFormat = C.VK_FORMAT_A2R10G10B10_UNORM_PACK32
	DataFormatA2R10G10B10SignedNormalized   common.DataFormat = C.VK_FORMAT_A2R10G10B10_SNORM_PACK32
	DataFormatA2R10G10B10UnsignedScaled     common.DataFormat = C.VK_FORMAT_A2R10G10B10_USCALED_PACK32
	DataFormatA2R10G10B10SignedScaled       common.DataFormat = C.VK_FORMAT_A2R10G10B10_SSCALED_PACK32
	DataFormatA2R10G10B10UnsignedInt        common.DataFormat = C.VK_FORMAT_A2R10G10B10_UINT_PACK32
	DataFormatA2R10G10B10SignedInt          common.DataFormat = C.VK_FORMAT_A2R10G10B10_SINT_PACK32

	DataFormatA2B10G10R10UnsignedNormalized common.DataFormat = C.VK_FORMAT_A2B10G10R10_UNORM_PACK32
	DataFormatA2B10G10R10SignedNormalized   common.DataFormat = C.VK_FORMAT_A2B10G10R10_SNORM_PACK32
	DataFormatA2B10G10R10UnsignedScaled     common.DataFormat = C.VK_FORMAT_A2B10G10R10_USCALED_PACK32
	DataFormatA2B10G10R10SignedScaled       common.DataFormat = C.VK_FORMAT_A2B10G10R10_SSCALED_PACK32
	DataFormatA2B10G10R10UnsignedInt        common.DataFormat = C.VK_FORMAT_A2B10G10R10_UINT_PACK32
	DataFormatA2B10G10R10SignedInt          common.DataFormat = C.VK_FORMAT_A2B10G10R10_SINT_PACK32

	DataFormatR16UnsignedNormalized common.DataFormat = C.VK_FORMAT_R16_UNORM
	DataFormatR16SignedNormalized   common.DataFormat = C.VK_FORMAT_R16_SNORM
	DataFormatR16UnsignedScaled     common.DataFormat = C.VK_FORMAT_R16_USCALED
	DataFormatR16SignedScaled       common.DataFormat = C.VK_FORMAT_R16_SSCALED
	DataFormatR16UnsignedInt        common.DataFormat = C.VK_FORMAT_R16_UINT
	DataFormatR16SignedInt          common.DataFormat = C.VK_FORMAT_R16_SINT
	DataFormatR16SignedFloat        common.DataFormat = C.VK_FORMAT_R16_SFLOAT

	DataFormatR16G16UnsignedNormalized common.DataFormat = C.VK_FORMAT_R16G16_UNORM
	DataFormatR16G16SignedNormalized   common.DataFormat = C.VK_FORMAT_R16G16_SNORM
	DataFormatR16G16UnsignedScaled     common.DataFormat = C.VK_FORMAT_R16G16_USCALED
	DataFormatR16G16SignedScaled       common.DataFormat = C.VK_FORMAT_R16G16_SSCALED
	DataFormatR16G16UnsignedInt        common.DataFormat = C.VK_FORMAT_R16G16_UINT
	DataFormatR16G16SignedInt          common.DataFormat = C.VK_FORMAT_R16G16_SINT
	DataFormatR16G16SignedFloat        common.DataFormat = C.VK_FORMAT_R16G16_SFLOAT

	DataFormatR16G16B16UnsignedNormalized common.DataFormat = C.VK_FORMAT_R16G16B16_UNORM
	DataFormatR16G16B16SignedNormalized   common.DataFormat = C.VK_FORMAT_R16G16B16_SNORM
	DataFormatR16G16B16UnsignedScaled     common.DataFormat = C.VK_FORMAT_R16G16B16_USCALED
	DataFormatR16G16B16SignedScaled       common.DataFormat = C.VK_FORMAT_R16G16B16_SSCALED
	DataFormatR16G16B16UnsignedInt        common.DataFormat = C.VK_FORMAT_R16G16B16_UINT
	DataFormatR16G16B16SignedInt          common.DataFormat = C.VK_FORMAT_R16G16B16_SINT
	DataFormatR16G16B16SignedFloat        common.DataFormat = C.VK_FORMAT_R16G16B16_SFLOAT

	DataFormatR16G16B16A16UnsignedNormalized common.DataFormat = C.VK_FORMAT_R16G16B16A16_UNORM
	DataFormatR16G16B16A16SignedNormalized   common.DataFormat = C.VK_FORMAT_R16G16B16A16_SNORM
	DataFormatR16G16B16A16UnsignedScaled     common.DataFormat = C.VK_FORMAT_R16G16B16A16_USCALED
	DataFormatR16G16B16A16SignedScaled       common.DataFormat = C.VK_FORMAT_R16G16B16A16_SSCALED
	DataFormatR16G16B16A16UnsignedInt        common.DataFormat = C.VK_FORMAT_R16G16B16A16_UINT
	DataFormatR16G16B16A16SignedInt          common.DataFormat = C.VK_FORMAT_R16G16B16A16_SINT
	DataFormatR16G16B16A16SignedFloat        common.DataFormat = C.VK_FORMAT_R16G16B16A16_SFLOAT

	DataFormatR32UnsignedInt          common.DataFormat = C.VK_FORMAT_R32_UINT
	DataFormatR32SignedInt            common.DataFormat = C.VK_FORMAT_R32_SINT
	DataFormatR32SignedFloat          common.DataFormat = C.VK_FORMAT_R32_SFLOAT
	DataFormatR32G32UnsignedInt       common.DataFormat = C.VK_FORMAT_R32G32_UINT
	DataFormatR32G32SignedInt         common.DataFormat = C.VK_FORMAT_R32G32_SINT
	DataFormatR32G32SignedFloat       common.DataFormat = C.VK_FORMAT_R32G32_SFLOAT
	DataFormatR32G32B32UnsignedInt    common.DataFormat = C.VK_FORMAT_R32G32B32_UINT
	DataFormatR32G32B32SignedInt      common.DataFormat = C.VK_FORMAT_R32G32B32_SINT
	DataFormatR32G32B32SignedFloat    common.DataFormat = C.VK_FORMAT_R32G32B32_SFLOAT
	DataFormatR32G32B32A32UnsignedInt common.DataFormat = C.VK_FORMAT_R32G32B32A32_UINT
	DataFormatR32G32B32A32SignedInt   common.DataFormat = C.VK_FORMAT_R32G32B32A32_SINT
	DataFormatR32G32B32A32SignedFloat common.DataFormat = C.VK_FORMAT_R32G32B32A32_SFLOAT

	DataFormatR64UnsignedInt          common.DataFormat = C.VK_FORMAT_R64_UINT
	DataFormatR64SignedInt            common.DataFormat = C.VK_FORMAT_R64_SINT
	DataFormatR64SignedFloat          common.DataFormat = C.VK_FORMAT_R64_SFLOAT
	DataFormatR64G64UnsignedInt       common.DataFormat = C.VK_FORMAT_R64G64_UINT
	DataFormatR64G64SignedInt         common.DataFormat = C.VK_FORMAT_R64G64_SINT
	DataFormatR64G64SignedFloat       common.DataFormat = C.VK_FORMAT_R64G64_SFLOAT
	DataFormatR64G64B64UnsignedInt    common.DataFormat = C.VK_FORMAT_R64G64B64_UINT
	DataFormatR64G64B64SignedInt      common.DataFormat = C.VK_FORMAT_R64G64B64_SINT
	DataFormatR64G64B64SignedFloat    common.DataFormat = C.VK_FORMAT_R64G64B64_SFLOAT
	DataFormatR64G64B64A64UnsignedInt common.DataFormat = C.VK_FORMAT_R64G64B64A64_UINT
	DataFormatR64G64B64A64SignedInt   common.DataFormat = C.VK_FORMAT_R64G64B64A64_SINT
	DataFormatR64G64B64A64SignedFloat common.DataFormat = C.VK_FORMAT_R64G64B64A64_SFLOAT

	DataFormatB10G11R11UnsignedFloat  common.DataFormat = C.VK_FORMAT_B10G11R11_UFLOAT_PACK32
	DataFormatE5B9G9R9UnsignedFloat   common.DataFormat = C.VK_FORMAT_E5B9G9R9_UFLOAT_PACK32
	DataFormatD16UnsignedNormalized   common.DataFormat = C.VK_FORMAT_D16_UNORM
	DataFormatD24X8UnsignedNormalized common.DataFormat = C.VK_FORMAT_X8_D24_UNORM_PACK32
	DataFormatD32SignedFloat          common.DataFormat = C.VK_FORMAT_D32_SFLOAT
	DataFormatS8UnsignedInt           common.DataFormat = C.VK_FORMAT_S8_UINT

	DataFormatD16UnsignedNormalizedS8UnsignedInt common.DataFormat = C.VK_FORMAT_D16_UNORM_S8_UINT
	DataFormatD24UnsignedNormalizedS8UnsignedInt common.DataFormat = C.VK_FORMAT_D24_UNORM_S8_UINT
	DataFormatD32SignedFloatS8UnsignedInt        common.DataFormat = C.VK_FORMAT_D32_SFLOAT_S8_UINT

	DataFormatBC1_RGBUnsignedNormalized  common.DataFormat = C.VK_FORMAT_BC1_RGB_UNORM_BLOCK
	DataFormatBC1_RGBsRGB                common.DataFormat = C.VK_FORMAT_BC1_RGB_SRGB_BLOCK
	DataFormatBC1_RGBAUnsignedNormalized common.DataFormat = C.VK_FORMAT_BC1_RGBA_UNORM_BLOCK
	DataFormatBC1_RGBAsRGB               common.DataFormat = C.VK_FORMAT_BC1_RGBA_SRGB_BLOCK

	DataFormatBC2_UnsignedNormalized common.DataFormat = C.VK_FORMAT_BC2_UNORM_BLOCK
	DataFormatBC2_sRGB               common.DataFormat = C.VK_FORMAT_BC2_SRGB_BLOCK

	DataFormatBC3_UnsignedNormalized common.DataFormat = C.VK_FORMAT_BC3_UNORM_BLOCK
	DataFormatBC3_sRGB               common.DataFormat = C.VK_FORMAT_BC3_SRGB_BLOCK

	DataFormatBC4_UnsignedNormalized common.DataFormat = C.VK_FORMAT_BC4_UNORM_BLOCK
	DataFormatBC4_SignedNormalized   common.DataFormat = C.VK_FORMAT_BC4_SNORM_BLOCK

	DataFormatBC5_UnsignedNormalized common.DataFormat = C.VK_FORMAT_BC5_UNORM_BLOCK
	DataFormatBC5_SignedNormalized   common.DataFormat = C.VK_FORMAT_BC5_SNORM_BLOCK

	DataFormatBC6_UnsignedFloat common.DataFormat = C.VK_FORMAT_BC6H_UFLOAT_BLOCK
	DataFormatBC6_SignedFloat   common.DataFormat = C.VK_FORMAT_BC6H_SFLOAT_BLOCK

	DataFormatBC7_UnsignedNormalized common.DataFormat = C.VK_FORMAT_BC7_UNORM_BLOCK
	DataFormatBC7_sRGB               common.DataFormat = C.VK_FORMAT_BC7_SRGB_BLOCK

	DataFormatETC2_R8G8B8UnsignedNormalized   common.DataFormat = C.VK_FORMAT_ETC2_R8G8B8_UNORM_BLOCK
	DataFormatETC2_R8G8B8sRGB                 common.DataFormat = C.VK_FORMAT_ETC2_R8G8B8_SRGB_BLOCK
	DataFormatETC2_R8G8B8A1UnsignedNormalized common.DataFormat = C.VK_FORMAT_ETC2_R8G8B8A1_UNORM_BLOCK
	DataFormatETC2_R8G8B8A1sRGB               common.DataFormat = C.VK_FORMAT_ETC2_R8G8B8A1_SRGB_BLOCK
	DataFormatETC2_R8G8B8A8UnsignedNormalized common.DataFormat = C.VK_FORMAT_ETC2_R8G8B8A8_UNORM_BLOCK
	DataFormatETC2_R8G8B8A8sRGB               common.DataFormat = C.VK_FORMAT_ETC2_R8G8B8A8_SRGB_BLOCK

	DataFormatEAC_R11UnsignedNormalized    common.DataFormat = C.VK_FORMAT_EAC_R11_UNORM_BLOCK
	DataFormatEAC_R11SignedNormalized      common.DataFormat = C.VK_FORMAT_EAC_R11_SNORM_BLOCK
	DataFormatEAC_R11G11UnsignedNormalized common.DataFormat = C.VK_FORMAT_EAC_R11G11_UNORM_BLOCK
	DataFormatEAC_R11G11SignedNormalized   common.DataFormat = C.VK_FORMAT_EAC_R11G11_SNORM_BLOCK

	DataFormatASTC4x4_UnsignedNormalized   common.DataFormat = C.VK_FORMAT_ASTC_4x4_UNORM_BLOCK
	DataFormatASTC4x4_sRGB                 common.DataFormat = C.VK_FORMAT_ASTC_4x4_SRGB_BLOCK
	DataFormatASTC5x4_UnsignedNormalized   common.DataFormat = C.VK_FORMAT_ASTC_5x4_UNORM_BLOCK
	DataFormatASTC5x4_sRGB                 common.DataFormat = C.VK_FORMAT_ASTC_5x4_SRGB_BLOCK
	DataFormatASTC5x5_UnsignedNormalized   common.DataFormat = C.VK_FORMAT_ASTC_5x5_UNORM_BLOCK
	DataFormatASTC5x5_sRGB                 common.DataFormat = C.VK_FORMAT_ASTC_5x5_SRGB_BLOCK
	DataFormatASTC6x5_UnsignedNormalized   common.DataFormat = C.VK_FORMAT_ASTC_6x5_UNORM_BLOCK
	DataFormatASTC6x5_sRGB                 common.DataFormat = C.VK_FORMAT_ASTC_6x5_SRGB_BLOCK
	DataFormatASTC6x6_UnsignedNormalized   common.DataFormat = C.VK_FORMAT_ASTC_6x6_UNORM_BLOCK
	DataFormatASTC6x6_sRGB                 common.DataFormat = C.VK_FORMAT_ASTC_6x6_SRGB_BLOCK
	DataFormatASTC8x5_UnsignedNormalized   common.DataFormat = C.VK_FORMAT_ASTC_8x5_UNORM_BLOCK
	DataFormatASTC8x5_sRGB                 common.DataFormat = C.VK_FORMAT_ASTC_8x5_SRGB_BLOCK
	DataFormatASTC8x6_UnsignedNormalized   common.DataFormat = C.VK_FORMAT_ASTC_8x6_UNORM_BLOCK
	DataFormatASTC8x6_sRGB                 common.DataFormat = C.VK_FORMAT_ASTC_8x6_SRGB_BLOCK
	DataFormatASTC8x8_UnsignedNormalized   common.DataFormat = C.VK_FORMAT_ASTC_8x8_UNORM_BLOCK
	DataFormatASTC8x8_sRGB                 common.DataFormat = C.VK_FORMAT_ASTC_8x8_SRGB_BLOCK
	DataFormatASTC10x5_UnsignedNormalized  common.DataFormat = C.VK_FORMAT_ASTC_10x5_UNORM_BLOCK
	DataFormatASTC10x5_sRGB                common.DataFormat = C.VK_FORMAT_ASTC_10x5_SRGB_BLOCK
	DataFormatASTC10x6_UnsignedNormalized  common.DataFormat = C.VK_FORMAT_ASTC_10x6_UNORM_BLOCK
	DataFormatASTC10x6_sRGB                common.DataFormat = C.VK_FORMAT_ASTC_10x6_SRGB_BLOCK
	DataFormatASTC10x8_UnsignedNormalized  common.DataFormat = C.VK_FORMAT_ASTC_10x8_UNORM_BLOCK
	DataFormatASTC10x8_sRGB                common.DataFormat = C.VK_FORMAT_ASTC_10x8_SRGB_BLOCK
	DataFormatASTC10x10_UnsignedNormalized common.DataFormat = C.VK_FORMAT_ASTC_10x10_UNORM_BLOCK
	DataFormatASTC10x10_sRGB               common.DataFormat = C.VK_FORMAT_ASTC_10x10_SRGB_BLOCK
	DataFormatASTC12x10_UnsignedNormalized common.DataFormat = C.VK_FORMAT_ASTC_12x10_UNORM_BLOCK
	DataFormatASTC12x10_sRGB               common.DataFormat = C.VK_FORMAT_ASTC_12x10_SRGB_BLOCK
	DataFormatASTC12x12_UnsignedNormalized common.DataFormat = C.VK_FORMAT_ASTC_12x12_UNORM_BLOCK
	DataFormatASTC12x12_sRGB               common.DataFormat = C.VK_FORMAT_ASTC_12x12_SRGB_BLOCK
)

func init() {
	DataFormatUndefined.Register("Undefined")
	DataFormatR4G4UnsignedNormalized.Register("R4G4 Unsigned Normalized")

	DataFormatR4G4B4A4UnsignedNormalized.Register("R4G4B4A4 Unsigned Normalized")
	DataFormatB4G4R4A4UnsignedNormalized.Register("B4G4R4A4 Unsigned Normalized")
	DataFormatR5G6B5UnsignedNormalized.Register("R5G6B5 Unsigned Normalized")
	DataFormatB5G6R5UnsignedNormalized.Register("B5G6R5 Unsigned Normalized")
	DataFormatR5G5B5A1UnsignedNormalized.Register("R5G5B5A1 Unsigned Normalized")
	DataFormatB5G5R5A1UnsignedNormalized.Register("B5G5R5A1 Unsigned Normalized")
	DataFormatA1R5G5B5UnsignedNormalized.Register("A1R5G5B5 Unsigned Normalized")

	DataFormatR8UnsignedNormalized.Register("R8 Unsigned Normalized")
	DataFormatR8SignedNormalized.Register("R8 Signed Normalized")
	DataFormatR8UnsignedScaled.Register("R8 Unsigned Scaled")
	DataFormatR8SignedScaled.Register("R8 Signed Scaled")
	DataFormatR8UnsignedInt.Register("R8 Unsigned Int")
	DataFormatR8SignedInt.Register("R8 Signed Int")
	DataFormatR8SRGB.Register("R8 sRGB")

	DataFormatR8G8UnsignedNormalized.Register("R8G8 Unsigned Normalized")
	DataFormatR8G8SignedNormalized.Register("R8G8 Signed Normalized")
	DataFormatR8G8UnsignedScaled.Register("R8G8 Unsigned Scaled")
	DataFormatR8G8SignedScaled.Register("R8G8 Signed Scaled")
	DataFormatR8G8UnsignedInt.Register("R8G8 Unsigned Int")
	DataFormatR8G8SignedInt.Register("R8G8 Signed Int")
	DataFormatR8G8SRGB.Register("R8G8 sRGB")

	DataFormatR8G8B8UnsignedNormalized.Register("R8G8B8 Unsigned Normalized")
	DataFormatR8G8B8SignedNormalized.Register("R8G8B8 Signed Normalized")
	DataFormatR8G8B8UnsignedScaled.Register("R8G8B8 Unsigned Scaled")
	DataFormatR8G8B8SignedScaled.Register("R8G8B8 Signed Scaled")
	DataFormatR8G8B8UnsignedInt.Register("R8G8B8 Unsigned Int")
	DataFormatR8G8B8SignedInt.Register("R8G8B8 Signed Int")
	DataFormatR8G8B8SRGB.Register("R8G8B8 sRGB")

	DataFormatB8G8R8UnsignedNormalized.Register("B8G8R8 Unsigned Normalized")
	DataFormatB8G8R8SignedNormalized.Register("B8G8R8 Signed Normalized")
	DataFormatB8G8R8UnsignedScaled.Register("B8G8R8 Unsigned Scaled")
	DataFormatB8G8R8SignedScaled.Register("B8G8R8 Signed Scaled")
	DataFormatB8G8R8UnsignedInt.Register("B8G8R8 Unsigned Int")
	DataFormatB8G8R8SignedInt.Register("B8G8R8 Signed Int")
	DataFormatB8G8R8SRGB.Register("B8G8R8 sRGB")

	DataFormatR8G8B8A8UnsignedNormalized.Register("R8G8B8A8 Unsigned Normalized")
	DataFormatR8G8B8A8SignedNormalized.Register("R8G8B8A8 Signed Normalized")
	DataFormatR8G8B8A8UnsignedScaled.Register("R8G8B8A8 Unsigned Scaled")
	DataFormatR8G8B8A8SignedScaled.Register("R8G8B8A8 Signed Scaled")
	DataFormatR8G8B8A8UnsignedInt.Register("R8G8B8A8 Unsigned Int")
	DataFormatR8G8B8A8SignedInt.Register("R8G8B8A8 Signed Int")
	DataFormatR8G8B8A8SRGB.Register("R8G8B8A8 sRGB")

	DataFormatB8G8R8A8UnsignedNormalized.Register("B8G8R8A8 Unsigned Normalized")
	DataFormatB8G8R8A8SignedNormalized.Register("B8G8R8A8 Signed Normalized")
	DataFormatB8G8R8A8UnsignedScaled.Register("B8G8R8A8 Unsigned Scaled")
	DataFormatB8G8R8A8SignedScaled.Register("B8G8R8A8 Signed Scaled")
	DataFormatB8G8R8A8UnsignedInt.Register("B8G8R8A8 Unsigned Int")
	DataFormatB8G8R8A8SignedInt.Register("B8G8R8A8 Signed Int")
	DataFormatB8G8R8A8SRGB.Register("B8G8R8A8 sRGB")

	DataFormatA8B8G8R8UnsignedNormalized.Register("A8B8G8R8 Unsigned Normalized")
	DataFormatA8B8G8R8SignedNormalized.Register("A8B8G8R8 Signed Normalized")
	DataFormatA8B8G8R8UnsignedScaled.Register("A8B8G8R8 Unsigned Scaled")
	DataFormatA8B8G8R8SignedScaled.Register("A8B8G8R8 Signed Scaled")
	DataFormatA8B8G8R8UnsignedInt.Register("A8B8G8R8 Unsigned Int")
	DataFormatA8B8G8R8SignedInt.Register("A8B8G8R8 Signed Int")
	DataFormatA8B8G8R8SRGB.Register("A8B8G8R8 sRGB")

	DataFormatA2R10G10B10UnsignedNormalized.Register("A2R10G10B10 Unsigned Normalized")
	DataFormatA2R10G10B10SignedNormalized.Register("A2R10G10B10 Signed Normalized")
	DataFormatA2R10G10B10UnsignedScaled.Register("A2R10G10B10 Unsigned Scaled")
	DataFormatA2R10G10B10SignedScaled.Register("A2R10G10B10 Signed Scaled")
	DataFormatA2R10G10B10UnsignedInt.Register("A2R10G10B10 Unsigned Int")
	DataFormatA2R10G10B10SignedInt.Register("A2R10G10B10 Signed Int")

	DataFormatA2B10G10R10UnsignedNormalized.Register("A2B10G10R10 Unsigned Normalized")
	DataFormatA2B10G10R10SignedNormalized.Register("A2B10G10R10 Signed Normalized")
	DataFormatA2B10G10R10UnsignedScaled.Register("A2B10G10R10 Unsigned Scaled")
	DataFormatA2B10G10R10SignedScaled.Register("A2B10G10R10 Signed Scaled")
	DataFormatA2B10G10R10UnsignedInt.Register("A2B10G10R10 Unsigned Int")
	DataFormatA2B10G10R10SignedInt.Register("A2B10G10R10 Signed Int")

	DataFormatR16UnsignedNormalized.Register("R16 Unsigned Normalized")
	DataFormatR16SignedNormalized.Register("R16 Signed Normalized")
	DataFormatR16UnsignedScaled.Register("R16 Unsigned Scaled")
	DataFormatR16SignedScaled.Register("R16 Signed Scaled")
	DataFormatR16UnsignedInt.Register("R16 Unsigned Int")
	DataFormatR16SignedInt.Register("R16 Signed Int")
	DataFormatR16SignedFloat.Register("R16 Signed Float")

	DataFormatR16G16UnsignedNormalized.Register("R16G16 Unsigned Normalized")
	DataFormatR16G16SignedNormalized.Register("R16G16 Signed Normalized")
	DataFormatR16G16UnsignedScaled.Register("R16G16 Unsigned Scaled")
	DataFormatR16G16SignedScaled.Register("R16G16 Signed Scaled")
	DataFormatR16G16UnsignedInt.Register("R16G16 Unsigned Int")
	DataFormatR16G16SignedInt.Register("R16G16 Signed Int")
	DataFormatR16G16SignedFloat.Register("R16G16 Signed Float")

	DataFormatR16G16B16UnsignedNormalized.Register("R16G16B16 Unsigned Normalized")
	DataFormatR16G16B16SignedNormalized.Register("R16G16B16 Signed Normalized")
	DataFormatR16G16B16UnsignedScaled.Register("R16G16B16 Unsigned Scaled")
	DataFormatR16G16B16SignedScaled.Register("R16G16B16 Signed Scaled")
	DataFormatR16G16B16UnsignedInt.Register("R16G16B16 Unsigned Int")
	DataFormatR16G16B16SignedInt.Register("R16G16B16 Signed Int")
	DataFormatR16G16B16SignedFloat.Register("R16G16B16 Signed Float")

	DataFormatR16G16B16A16UnsignedNormalized.Register("R16G16B16A16 Unsigned Normalized")
	DataFormatR16G16B16A16SignedNormalized.Register("R16G16B16A16 Signed Normalized")
	DataFormatR16G16B16A16UnsignedScaled.Register("R16G16B16A16 Unsigned Scaled")
	DataFormatR16G16B16A16SignedScaled.Register("R16G16B16A16 Signed Scaled")
	DataFormatR16G16B16A16UnsignedInt.Register("R16G16B16A16 Unsigned Int")
	DataFormatR16G16B16A16SignedInt.Register("R16G16B16A16 Signed Int")
	DataFormatR16G16B16A16SignedFloat.Register("R16G16B16A16 Signed Float")

	DataFormatR32UnsignedInt.Register("R32 Unsigned Int")
	DataFormatR32SignedInt.Register("R32 Signed Int")
	DataFormatR32SignedFloat.Register("R32 Signed Float")
	DataFormatR32G32UnsignedInt.Register("R32G32 Unsigned Int")
	DataFormatR32G32SignedInt.Register("R32G32 Signed Int")
	DataFormatR32G32SignedFloat.Register("R32G32 Signed Float")
	DataFormatR32G32B32UnsignedInt.Register("R32G32B32 Unsigned Int")
	DataFormatR32G32B32SignedInt.Register("R32G32B32 Signed Int")
	DataFormatR32G32B32SignedFloat.Register("R32G32B32 Signed Float")
	DataFormatR32G32B32A32UnsignedInt.Register("R32G32B32A32 Unsigned Int")
	DataFormatR32G32B32A32SignedInt.Register("R32G32B32A32 Signed Int")
	DataFormatR32G32B32A32SignedFloat.Register("R32G32B32A32 Signed Float")

	DataFormatR64UnsignedInt.Register("R64 Unsigned Int")
	DataFormatR64SignedInt.Register("R64 Signed Int")
	DataFormatR64SignedFloat.Register("R64 Signed Float")
	DataFormatR64G64UnsignedInt.Register("R64G64 Unsigned Int")
	DataFormatR64G64SignedInt.Register("R64G64 Signed Int")
	DataFormatR64G64SignedFloat.Register("R64G64 Signed Float")
	DataFormatR64G64B64UnsignedInt.Register("R64G64B64 Unsigned Int")
	DataFormatR64G64B64SignedInt.Register("R64G64B64 Signed Int")
	DataFormatR64G64B64SignedFloat.Register("R64G64B64 Signed Float")
	DataFormatR64G64B64A64UnsignedInt.Register("R64G64B64A64 Unsigned Int")
	DataFormatR64G64B64A64SignedInt.Register("R64G64B64A64 Signed Int")
	DataFormatR64G64B64A64SignedFloat.Register("R64G64B64A64 Signed Float")

	DataFormatB10G11R11UnsignedFloat.Register("B10G11R11 Unsigned Float")
	DataFormatE5B9G9R9UnsignedFloat.Register("E5B9G9R9 Unsigned Float")
	DataFormatD16UnsignedNormalized.Register("D16 Unsigned Normalized")
	DataFormatD24X8UnsignedNormalized.Register("D24X8 Unsigned Normalized")
	DataFormatD32SignedFloat.Register("D32 Signed Float")
	DataFormatS8UnsignedInt.Register("S8 Unsigned Int")

	DataFormatD16UnsignedNormalizedS8UnsignedInt.Register("D16 Unsigned Normalized S8 Unsigned Int")
	DataFormatD24UnsignedNormalizedS8UnsignedInt.Register("D24 Unsigned Normalized S8 Unsigned Int")
	DataFormatD32SignedFloatS8UnsignedInt.Register("D32 Signed Float S8 Unsigned Int")

	DataFormatBC1_RGBUnsignedNormalized.Register("BC1-Compressed -Compressed RGB Unsigned Normalized")
	DataFormatBC1_RGBsRGB.Register("BC1-Compressed -Compressed RGB sRGB")
	DataFormatBC1_RGBAUnsignedNormalized.Register("BC1-Compressed -Compressed RGBA Unsigned Normalized")
	DataFormatBC1_RGBAsRGB.Register("BC1-Compressed RGBA sRGB")

	DataFormatBC2_UnsignedNormalized.Register("BC2-Compressed Unsigned Normalized")
	DataFormatBC2_sRGB.Register("BC2-Compressed sRGB")

	DataFormatBC3_UnsignedNormalized.Register("BC3-Compressed Unsigned Normalized")
	DataFormatBC3_sRGB.Register("BC3-Compressed sRGB")

	DataFormatBC4_UnsignedNormalized.Register("BC4-Compressed Unsigned Normalized")
	DataFormatBC4_SignedNormalized.Register("BC4-Compressed Signed Normalized")

	DataFormatBC5_UnsignedNormalized.Register("BC5-Compressed Unsigned Normalized")
	DataFormatBC5_SignedNormalized.Register("BC5-Compressed Signed Normalized")

	DataFormatBC6_UnsignedFloat.Register("BC6-Compressed Unsigned Float")
	DataFormatBC6_SignedFloat.Register("BC6-Compressed Signed Float")

	DataFormatBC7_UnsignedNormalized.Register("BC7-Compressed Unsigned Normalized")
	DataFormatBC7_sRGB.Register("BC7-Compressed sRGB")

	DataFormatETC2_R8G8B8UnsignedNormalized.Register("ETC2-Compressed R8G8B8 Unsigned Normalized")
	DataFormatETC2_R8G8B8sRGB.Register("ETC2-Compressed R8G8B8 sRGB")
	DataFormatETC2_R8G8B8A1UnsignedNormalized.Register("ETC2-Compressed R8G8B8A1 Unsigned Normalized")
	DataFormatETC2_R8G8B8A1sRGB.Register("ETC2-Compressed R8G8B8A1 sRGB")
	DataFormatETC2_R8G8B8A8UnsignedNormalized.Register("ETC2-Compressed R8G8B8A8 Unsigned Normalized")
	DataFormatETC2_R8G8B8A8sRGB.Register("ETC2-Compressed R8G8B8A8 sRGB")

	DataFormatEAC_R11UnsignedNormalized.Register("EAC-Compressed R11 Unsigned Normalized")
	DataFormatEAC_R11SignedNormalized.Register("EAC-Compressed R11 Signed Normalized")
	DataFormatEAC_R11G11UnsignedNormalized.Register("EAC-Compressed R11G11 Unsigned Normalized")
	DataFormatEAC_R11G11SignedNormalized.Register("EAC-Compressed R11G11 Signed Normalized")

	DataFormatASTC4x4_UnsignedNormalized.Register("ASTC-Compressed (4x4) Unsigned Normalized")
	DataFormatASTC4x4_sRGB.Register("ASTC-Compressed (4x4) sRGB")
	DataFormatASTC5x4_UnsignedNormalized.Register("ASTC-Compressed (5x4) Unsigned Normalized")
	DataFormatASTC5x4_sRGB.Register("ASTC-Compressed (5x4) sRGB")
	DataFormatASTC5x5_UnsignedNormalized.Register("ASTC-Compressed (5x5) Unsigned Normalized")
	DataFormatASTC5x5_sRGB.Register("ASTC-Compressed (5x5) sRGB")
	DataFormatASTC6x5_UnsignedNormalized.Register("ASTC-Compressed (6x5) Unsigned Normalized")
	DataFormatASTC6x5_sRGB.Register("ASTC-Compressed (6x5) sRGB")
	DataFormatASTC6x6_UnsignedNormalized.Register("ASTC-Compressed (6x6) Unsigned Normalized")
	DataFormatASTC6x6_sRGB.Register("ASTC-Compressed (6x6) sRGB")
	DataFormatASTC8x5_UnsignedNormalized.Register("ASTC-Compressed (8x5) Unsigned Normalized")
	DataFormatASTC8x5_sRGB.Register("ASTC-Compressed (8x5) sRGB")
	DataFormatASTC8x6_UnsignedNormalized.Register("ASTC-Compressed (8x6) Unsigned Normalized")
	DataFormatASTC8x6_sRGB.Register("ASTC-Compressed (8x6) sRGB")
	DataFormatASTC8x8_UnsignedNormalized.Register("ASTC-Compressed (8x8) Unsigned Normalized")
	DataFormatASTC8x8_sRGB.Register("ASTC-Compressed (8x8) sRGB")
	DataFormatASTC10x5_UnsignedNormalized.Register("ASTC-Compressed (10x5) Unsigned Normalized")
	DataFormatASTC10x5_sRGB.Register("ASTC-Compressed (10x5) sRGB")
	DataFormatASTC10x6_UnsignedNormalized.Register("ASTC-Compressed (10x6) Unsigned Normalized")
	DataFormatASTC10x6_sRGB.Register("ASTC-Compressed (10x6) sRGB")
	DataFormatASTC10x8_UnsignedNormalized.Register("ASTC-Compressed (10x8) Unsigned Normalized")
	DataFormatASTC10x8_sRGB.Register("ASTC-Compressed (10x8) sRGB")
	DataFormatASTC10x10_UnsignedNormalized.Register("ASTC-Compressed (10x10) Unsigned Normalized")
	DataFormatASTC10x10_sRGB.Register("ASTC-Compressed (10x10) sRGB")
	DataFormatASTC12x10_UnsignedNormalized.Register("ASTC-Compressed (12x10) Unsigned Normalized")
	DataFormatASTC12x10_sRGB.Register("ASTC-Compressed (12x10) sRGB")
	DataFormatASTC12x12_UnsignedNormalized.Register("ASTC-Compressed (12x12) Unsigned Normalized")
	DataFormatASTC12x12_sRGB.Register("ASTC-Compressed (12x12) sRGB")
}
