#include "textflag.h"

#define R_TMP R0
#define R_S_PTR R1
#define R_S_LEN R2

#define V_TBL0 V0
#define V_TBL1 V1
#define V_TBL2 V2
#define V_TBL3 V3
#define V_TBL4 V4
#define V_TBL5 V5
#define V_TBL6 V6
#define V_TBL7 V7
#define V_S_LO V8
#define V_S_HI V9
#define V_64 V10
#define V_RES_LO V11
#define V_RES_HI V12
#define V_RES V13

TEXT ·IsAlphaNumericSIMD(SB),NOSPLIT,$0-17
	// init helper registers
	MOVD s_base+0(FP), R_S_PTR
	MOVD s_len+8(FP), R_S_LEN
	VMOVI $64, V_64.B16
	// load 128b lookupTable (see isalphanumeric.go) into 8 16b NEON registers
	MOVD $·lookupTable(SB), R_TMP
	VLD1.P 64(R_TMP), [V_TBL0.B16, V_TBL1.B16, V_TBL2.B16, V_TBL3.B16]
	VLD1 (R_TMP), [V_TBL4.B16, V_TBL5.B16, V_TBL6.B16, V_TBL7.B16]
loop:
	// load 16b of input s and perform two table lookups (lo: 0-63 and hi: 64-127)
	VLD1.P 16(R_S_PTR), [V_S_LO.B16]
	VSUB V_64.B16, V_S_LO.B16, V_S_HI.B16
	VTBL V_S_LO.B16, [V_TBL0.B16, V_TBL1.B16, V_TBL2.B16, V_TBL3.B16], V_RES_LO.B16
	VTBL V_S_HI.B16, [V_TBL4.B16, V_TBL5.B16, V_TBL6.B16, V_TBL7.B16], V_RES_HI.B16
	// combine lo and hi table lookup (because Go doesn't implement VTBX : /)
	VORR V_RES_LO.B16, V_RES_HI.B16, V_RES.B16
	// return false if any invalid chars were detected
	VMOV V_RES.D[0], R_TMP
	CBNZ R_TMP, return_false
	VMOV V_RES.D[1], R_TMP
	CBNZ R_TMP, return_false
	// decrement remaining s len and keep looping if needed
	SUBS $16, R_S_LEN, R_S_LEN
	BNE loop
return_true:
	MOVW $1, R_TMP
	MOVD R_TMP, ret+16(FP)
	RET
return_false:
	MOVW $0, R_TMP
	MOVD R_TMP, ret+16(FP)
	RET