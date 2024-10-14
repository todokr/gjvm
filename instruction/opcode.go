package instruction

// https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-6.html#jvms-6.5
const (
	OpNop             = 0x00
	OpAConstNull      = 0x01
	OpIConstM1        = 0x02
	OpIConst0         = 0x03
	OpIConst1         = 0x04
	OpIConst2         = 0x05
	OpIConst3         = 0x06
	OpIConst4         = 0x07
	OpIConst5         = 0x08
	OpLConst0         = 0x09
	OpLConst1         = 0x0a
	OpFConst0         = 0x0b
	OpFConst1         = 0x0c
	OpFConst2         = 0x0d
	OpDConst0         = 0x0e
	OpDConst1         = 0x0f
	OpBIPush          = 0x10
	OpSIPush          = 0x11
	OpLDC             = 0x12
	OpLDCw            = 0x13
	OpLDC2w           = 0x14
	OpILoad           = 0x15
	OpLLoad           = 0x16
	OpFLoad           = 0x17
	OpDLoad           = 0x18
	OpALoad           = 0x19
	OpILoad0          = 0x1a
	OpILoad1          = 0x1b
	OpILoad2          = 0x1c
	OpILoad3          = 0x1d
	OpLLoad0          = 0x1e
	OpLLoad1          = 0x1f
	OpLLoad2          = 0x20
	OpLLoad3          = 0x21
	OpFLoad0          = 0x22
	OpFLoad1          = 0x23
	OpFLoad2          = 0x24
	OpFLoad3          = 0x25
	OpDLoad0          = 0x26
	OpDLoad1          = 0x27
	OpDLoad2          = 0x28
	OpDLoad3          = 0x29
	OpALoad0          = 0x2a
	OpALoad1          = 0x2b
	OpALoad2          = 0x2c
	OpALoad3          = 0x2d
	OpIALoad          = 0x2e
	OpLALoad          = 0x2f
	OpFALoad          = 0x30
	OpDALoad          = 0x31
	OpAALoad          = 0x32
	OpBALoad          = 0x33
	OpCALoad          = 0x34
	OpSALoad          = 0x35
	OpIStore          = 0x36
	OpLStore          = 0x37
	OpFStore          = 0x38
	OpDStore          = 0x39
	OpAStore          = 0x3a
	OpIStore0         = 0x3b
	OpIStore1         = 0x3c
	OpIStore2         = 0x3d
	OpIStore3         = 0x3e
	OpLStore0         = 0x3f
	OpLStore1         = 0x40
	OpLStore2         = 0x41
	OpLStore3         = 0x42
	OpFStore0         = 0x43
	OpFStore1         = 0x44
	OpFStore2         = 0x45
	OpFStore3         = 0x46
	OpDStore0         = 0x47
	OpDStore1         = 0x48
	OpDStore2         = 0x49
	OpDStore3         = 0x4a
	OpAStore0         = 0x4b
	OpAStore1         = 0x4c
	OpAStore2         = 0x4d
	OpAStore3         = 0x4e
	OpIAStore         = 0x4f
	OpLAStore         = 0x50
	OpFAStore         = 0x51
	OpDAStore         = 0x52
	OpAAStore         = 0x53
	OpBAStore         = 0x54
	OpCAStore         = 0x55
	OpSAStore         = 0x56
	OpPop             = 0x57
	OpPop2            = 0x58
	OpDup             = 0x59
	OpDupX1           = 0x5a
	OpDupX2           = 0x5b
	OpDup2            = 0x5c
	OpDup2X1          = 0x5d
	OpDup2X2          = 0x5e
	OpSwap            = 0x5f
	OpIAdd            = 0x60
	OpLAdd            = 0x61
	OpFAdd            = 0x62
	OpDAdd            = 0x63
	OpISub            = 0x64
	OpLSub            = 0x65
	OpFSub            = 0x66
	OpDSub            = 0x67
	OpIMul            = 0x68
	OpLMul            = 0x69
	OpFMul            = 0x6a
	OpDMul            = 0x6b
	OpIDiv            = 0x6c
	OpLDiv            = 0x6d
	OpFDiv            = 0x6e
	OpDDiv            = 0x6f
	OpIRem            = 0x70
	OpLRem            = 0x71
	OpFRem            = 0x72
	OpDRem            = 0x73
	OpINeg            = 0x74
	OpLNeg            = 0x75
	OpFNeg            = 0x76
	OpDNeg            = 0x77
	OpIShl            = 0x78
	OpLShl            = 0x79
	OpIShr            = 0x7a
	OpLShr            = 0x7b
	OpIUshr           = 0x7c
	OpLUshr           = 0x7d
	OpIAnd            = 0x7e
	OpLAnd            = 0x7f
	OpIOr             = 0x80
	OpLOr             = 0x81
	OpIXor            = 0x82
	OpLXor            = 0x83
	OpIInc            = 0x84
	OpI2L             = 0x85
	OpI2F             = 0x86
	OpI2D             = 0x87
	OpL2I             = 0x88
	OpL2F             = 0x89
	OpL2D             = 0x8a
	OpF2I             = 0x8b
	OpF2L             = 0x8c
	OpF2D             = 0x8d
	OpD2I             = 0x8e
	OpD2L             = 0x8f
	OpD2F             = 0x90
	OpI2B             = 0x91
	OpI2C             = 0x92
	OpI2S             = 0x93
	OpLCmp            = 0x94
	OpFCmpL           = 0x95
	OpFCmpG           = 0x96
	OpDCmpL           = 0x97
	OpDCmpG           = 0x98
	OpIfEQ            = 0x99
	OpIfNE            = 0x9a
	OpIfLT            = 0x9b
	OpIfGE            = 0x9c
	OpIfGT            = 0x9d
	OpIfLE            = 0x9e
	OpIfICmpEQ        = 0x9f
	OpIfICmpNE        = 0xa0
	OpIfICmpLT        = 0xa1
	OpIfICmpGE        = 0xa2
	OpIfICmpGT        = 0xa3
	OpIfICmpLE        = 0xa4
	OpIfACmpEQ        = 0xa5
	OpIfACmpNE        = 0xa6
	OpGoto            = 0xa7
	OpJSR             = 0xa8
	OpRET             = 0xa9
	OpTableSwitch     = 0xaa
	OpLookupSwitch    = 0xab
	OpIReturn         = 0xac
	OpLReturn         = 0xad
	OpFReturn         = 0xae
	OpDReturn         = 0xaf
	OpAReturn         = 0xb0
	OpReturn          = 0xb1
	OpGetStatic       = 0xb2
	OpPupStatic       = 0xb3
	OpGetField        = 0xb4
	OpPutField        = 0xb5
	OpInvokeVirtual   = 0xb6
	OpInvokeSpecial   = 0xb7
	OpInvokeStatic    = 0xb8
	OpInvokeInterface = 0xb9
	OpInvokeDynamic   = 0xba
	OpNew             = 0xbb
	OpNewArray        = 0xbc
	OpANewArray       = 0xbd
	OpArrayLength     = 0xbe
	OpAThrow          = 0xbf
	OpCheckCast       = 0xc0
	OpInstanceOf      = 0xc1
	OpMonitorEnter    = 0xc2
	OpMonitorExit     = 0xc3
	OpWide            = 0xc4
	OpMultiANewArray  = 0xc5
	OpIfNull          = 0xc6
	OpIfNonNull       = 0xc7
	OpGotoW           = 0xc8
	OpJSRw            = 0xc9
	OpBreakpoint      = 0xca
	OpInvokeNative    = 0xfe // impdep1
	OpBootstrap       = 0xff // impdep2
)
