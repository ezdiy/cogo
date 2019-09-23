TEXT ·WantsPreempt(SB), 7, $0
    MOVQ    TLS, CX
    MOVQ    0(CX)(TLS*1), DX
    MOVBQZX 177(DX), AX
    MOVQ    AX, ret+0(FP)
    RET
