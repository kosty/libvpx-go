// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 22 Sep 2016 02:16:17 MSK.
// By https://git.io/cgogen. DO NOT EDIT.

package vpx

/*
#cgo pkg-config: vpx
#include <vpx/vpx_encoder.h>
#include <vpx/vpx_decoder.h>
#include <vpx/vp8.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// TsMaxPeriodicity as defined in vpx-1.5.0/vpx_encoder.h:37
	TsMaxPeriodicity = 16
	// TsMaxLayers as defined in vpx-1.5.0/vpx_encoder.h:40
	TsMaxLayers = 5
	// MaxLayers as defined in vpx-1.5.0/vpx_encoder.h:46
	MaxLayers = 12
	// SsMaxLayers as defined in vpx-1.5.0/vpx_encoder.h:52
	SsMaxLayers = 5
	// SsDefaultLayers as defined in vpx-1.5.0/vpx_encoder.h:55
	SsDefaultLayers = 1
	// EncoderABIVersion as defined in vpx-1.5.0/vpx_encoder.h:65
	EncoderABIVersion = 11
	// CodecCapPsnr as defined in vpx-1.5.0/vpx_encoder.h:77
	CodecCapPsnr = 65536
	// CodecCapOutputPartition as defined in vpx-1.5.0/vpx_encoder.h:84
	CodecCapOutputPartition = 131072
	// CodecCapHighbitdepth as defined in vpx-1.5.0/vpx_encoder.h:88
	CodecCapHighbitdepth = 262144
	// CodecUsePsnr as defined in vpx-1.5.0/vpx_encoder.h:97
	CodecUsePsnr = 65536
	// CodecUseOutputPartition as defined in vpx-1.5.0/vpx_encoder.h:98
	CodecUseOutputPartition = 131072
	// CodecUseHighbitdepth as defined in vpx-1.5.0/vpx_encoder.h:100
	CodecUseHighbitdepth = 262144
	// FrameIsKey as defined in vpx-1.5.0/vpx_encoder.h:129
	FrameIsKey = 1
	// FrameIsDroppable as defined in vpx-1.5.0/vpx_encoder.h:130
	FrameIsDroppable = 2
	// FrameIsInvisible as defined in vpx-1.5.0/vpx_encoder.h:133
	FrameIsInvisible = 4
	// FrameIsFragment as defined in vpx-1.5.0/vpx_encoder.h:135
	FrameIsFragment = 8
	// ErrorResilientDefault as defined in vpx-1.5.0/vpx_encoder.h:145
	ErrorResilientDefault = 1
	// ErrorResilientPartitions as defined in vpx-1.5.0/vpx_encoder.h:147
	ErrorResilientPartitions = 2
	// EflagForceKf as defined in vpx-1.5.0/vpx_encoder.h:292
	EflagForceKf = 1
	// DlRealtime as defined in vpx-1.5.0/vpx_encoder.h:898
	DlRealtime = 1
	// DlGoodQuality as defined in vpx-1.5.0/vpx_encoder.h:900
	DlGoodQuality = 1000000
	// DlBestQuality as defined in vpx-1.5.0/vpx_encoder.h:902
	DlBestQuality = 0
	// CodecABIVersion as defined in vpx-1.5.0/vpx_codec.h:86
	CodecABIVersion = 6
	// CodecCapDecoder as defined in vpx-1.5.0/vpx_codec.h:154
	CodecCapDecoder = 1
	// CodecCapEncoder as defined in vpx-1.5.0/vpx_codec.h:155
	CodecCapEncoder = 2
	// ForceInline as defined in vpx-1.5.0/vpx_integer.h:22
	ForceInline = 0
	// Inline as defined in vpx-1.5.0/vpx_integer.h:24
	Inline = 0
	// ImageABIVersion as defined in vpx-1.5.0/vpx_image.h:31
	ImageABIVersion = 3
	// ImageFormatPlanar as defined in vpx-1.5.0/vpx_image.h:34
	ImageFormatPlanar = 256
	// ImageFormatUvFlip as defined in vpx-1.5.0/vpx_image.h:35
	ImageFormatUvFlip = 512
	// ImageFormatHasAlpha as defined in vpx-1.5.0/vpx_image.h:36
	ImageFormatHasAlpha = 1024
	// ImageFormatHighbitdepth as defined in vpx-1.5.0/vpx_image.h:37
	ImageFormatHighbitdepth = 2048
	// PlanePacked as defined in vpx-1.5.0/vpx_image.h:111
	PlanePacked = 0
	// PlaneY as defined in vpx-1.5.0/vpx_image.h:112
	PlaneY = 0
	// PlaneU as defined in vpx-1.5.0/vpx_image.h:113
	PlaneU = 1
	// PlaneV as defined in vpx-1.5.0/vpx_image.h:114
	PlaneV = 2
	// PlaneAlpha as defined in vpx-1.5.0/vpx_image.h:115
	PlaneAlpha = 3
	// DecoderABIVersion as defined in vpx-1.5.0/vpx_decoder.h:43
	DecoderABIVersion = 9
	// CodecCapPutSlice as defined in vpx-1.5.0/vpx_decoder.h:53
	CodecCapPutSlice = 65536
	// CodecCapPutFrame as defined in vpx-1.5.0/vpx_decoder.h:54
	CodecCapPutFrame = 131072
	// CodecCapPostproc as defined in vpx-1.5.0/vpx_decoder.h:55
	CodecCapPostproc = 262144
	// CodecCapErrorConcealment as defined in vpx-1.5.0/vpx_decoder.h:56
	CodecCapErrorConcealment = 524288
	// CodecCapInputFragments as defined in vpx-1.5.0/vpx_decoder.h:58
	CodecCapInputFragments = 1048576
	// CodecCapFrameThreading as defined in vpx-1.5.0/vpx_decoder.h:68
	CodecCapFrameThreading = 2097152
	// CodecCapExternalFrameBuffer as defined in vpx-1.5.0/vpx_decoder.h:70
	CodecCapExternalFrameBuffer = 4194304
	// CodecUsePostproc as defined in vpx-1.5.0/vpx_decoder.h:73
	CodecUsePostproc = 65536
	// CodecUseErrorConcealment as defined in vpx-1.5.0/vpx_decoder.h:74
	CodecUseErrorConcealment = 131072
	// CodecUseInputFragments as defined in vpx-1.5.0/vpx_decoder.h:76
	CodecUseInputFragments = 262144
	// CodecUseFrameThreading as defined in vpx-1.5.0/vpx_decoder.h:79
	CodecUseFrameThreading = 524288
	// MaximumWorkBuffers as defined in vpx-1.5.0/vpx_frame_buffer.h:29
	MaximumWorkBuffers = 8
	// Vp9MaximumRefBuffers as defined in vpx-1.5.0/vpx_frame_buffer.h:33
	Vp9MaximumRefBuffers = 8
)

// CodecCxPktKind as declared in vpx-1.5.0/vpx_encoder.h:162
type CodecCxPktKind int32

// CodecCxPktKind enumeration from vpx-1.5.0/vpx_encoder.h:162
const (
	CodecCxFramePkt   = C.VPX_CODEC_CX_FRAME_PKT
	CodecStatsPkt     = C.VPX_CODEC_STATS_PKT
	CodecFpmbStatsPkt = C.VPX_CODEC_FPMB_STATS_PKT
	CodecPsnrPkt      = C.VPX_CODEC_PSNR_PKT
	CodecCustomPkt    = C.VPX_CODEC_CUSTOM_PKT
)

// EncPass as declared in vpx-1.5.0/vpx_encoder.h:253
type EncPass int32

// EncPass enumeration from vpx-1.5.0/vpx_encoder.h:253
const (
	RcOnePass   = C.VPX_RC_ONE_PASS
	RcFirstPass = C.VPX_RC_FIRST_PASS
	RcLastPass  = C.VPX_RC_LAST_PASS
)

// RcMode as declared in vpx-1.5.0/vpx_encoder.h:261
type RcMode int32

// RcMode enumeration from vpx-1.5.0/vpx_encoder.h:261
const (
	Vbr = C.VPX_VBR
	Cbr = C.VPX_CBR
	Cq  = C.VPX_CQ
	Q   = C.VPX_Q
)

// KfMode as declared in vpx-1.5.0/vpx_encoder.h:277
type KfMode int32

// KfMode enumeration from vpx-1.5.0/vpx_encoder.h:277
const (
	KfFixed    = C.VPX_KF_FIXED
	KfAuto     = C.VPX_KF_AUTO
	KfDisabled = C.VPX_KF_DISABLED
)

// BitDepth as declared in vpx-1.5.0/vpx_codec.h:223
type BitDepth int32

// BitDepth enumeration from vpx-1.5.0/vpx_codec.h:223
const (
	Bits8  BitDepth = C.VPX_BITS_8
	Bits10 BitDepth = C.VPX_BITS_10
	Bits12 BitDepth = C.VPX_BITS_12
)

// ImageFormat as declared in vpx-1.5.0/vpx_image.h:67
type ImageFormat int32

// ImageFormat enumeration from vpx-1.5.0/vpx_image.h:67
const (
	ImageFormatNone     ImageFormat = C.VPX_IMG_FMT_NONE
	ImageFormatRgb24    ImageFormat = C.VPX_IMG_FMT_RGB24
	ImageFormatRgb32    ImageFormat = C.VPX_IMG_FMT_RGB32
	ImageFormatRgb565   ImageFormat = C.VPX_IMG_FMT_RGB565
	ImageFormatRgb555   ImageFormat = C.VPX_IMG_FMT_RGB555
	ImageFormatUyvy     ImageFormat = C.VPX_IMG_FMT_UYVY
	ImageFormatYuy2     ImageFormat = C.VPX_IMG_FMT_YUY2
	ImageFormatYvyu     ImageFormat = C.VPX_IMG_FMT_YVYU
	ImageFormatBgr24    ImageFormat = C.VPX_IMG_FMT_BGR24
	ImageFormatRgb32Le  ImageFormat = C.VPX_IMG_FMT_RGB32_LE
	ImageFormatArgb     ImageFormat = C.VPX_IMG_FMT_ARGB
	ImageFormatArgbLe   ImageFormat = C.VPX_IMG_FMT_ARGB_LE
	ImageFormatRgb565Le ImageFormat = C.VPX_IMG_FMT_RGB565_LE
	ImageFormatRgb555Le ImageFormat = C.VPX_IMG_FMT_RGB555_LE
	ImageFormatYv12     ImageFormat = C.VPX_IMG_FMT_YV12
	ImageFormatI420     ImageFormat = C.VPX_IMG_FMT_I420
	ImageFormatVpxyv12  ImageFormat = C.VPX_IMG_FMT_VPXYV12
	ImageFormatVpxi420  ImageFormat = C.VPX_IMG_FMT_VPXI420
	ImageFormatI422     ImageFormat = C.VPX_IMG_FMT_I422
	ImageFormatI444     ImageFormat = C.VPX_IMG_FMT_I444
	ImageFormatI440     ImageFormat = C.VPX_IMG_FMT_I440
	ImageFormat444a     ImageFormat = C.VPX_IMG_FMT_444A
	ImageFormatI42016   ImageFormat = C.VPX_IMG_FMT_I42016
	ImageFormatI42216   ImageFormat = C.VPX_IMG_FMT_I42216
	ImageFormatI44416   ImageFormat = C.VPX_IMG_FMT_I44416
	ImageFormatI44016   ImageFormat = C.VPX_IMG_FMT_I44016
)

// ColorSpace as declared in vpx-1.5.0/vpx_image.h:79
type ColorSpace int32

// ColorSpace enumeration from vpx-1.5.0/vpx_image.h:79
const (
	ColorSpaceUnknown  ColorSpace = C.VPX_CS_UNKNOWN
	ColorSpaceBt601    ColorSpace = C.VPX_CS_BT_601
	ColorSpaceBt709    ColorSpace = C.VPX_CS_BT_709
	ColorSpaceSmpte170 ColorSpace = C.VPX_CS_SMPTE_170
	ColorSpaceSmpte240 ColorSpace = C.VPX_CS_SMPTE_240
	ColorSpaceBt2020   ColorSpace = C.VPX_CS_BT_2020
	ColorSpaceReserved ColorSpace = C.VPX_CS_RESERVED
	ColorSpaceSrgb     ColorSpace = C.VPX_CS_SRGB
)

// ColorRange as declared in vpx-1.5.0/vpx_image.h:85
type ColorRange int32

// ColorRange enumeration from vpx-1.5.0/vpx_image.h:85
const (
	CrStudioRange ColorRange = C.VPX_CR_STUDIO_RANGE
	CrFullRange   ColorRange = C.VPX_CR_FULL_RANGE
)

// CodecErr as declared in vpx-1.5.0/vpx_codec.h:142
type CodecErr int32

// CodecErr enumeration from vpx-1.5.0/vpx_codec.h:142
const (
	CodecOk             CodecErr = C.VPX_CODEC_OK
	CodecError          CodecErr = C.VPX_CODEC_ERROR
	CodecMemError       CodecErr = C.VPX_CODEC_MEM_ERROR
	CodecABIMismatch    CodecErr = C.VPX_CODEC_ABI_MISMATCH
	CodecIncapable      CodecErr = C.VPX_CODEC_INCAPABLE
	CodecUnsupBitstream CodecErr = C.VPX_CODEC_UNSUP_BITSTREAM
	CodecUnsupFeature   CodecErr = C.VPX_CODEC_UNSUP_FEATURE
	CodecCorruptFrame   CodecErr = C.VPX_CODEC_CORRUPT_FRAME
	CodecInvalidParam   CodecErr = C.VPX_CODEC_INVALID_PARAM
	CodecListEnd        CodecErr = C.VPX_CODEC_LIST_END
)
