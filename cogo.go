// Cooperative GC for tight loop Goroutine - prompt GC whether it needs a preempt in busy looping code.

package cogo

func WantsPreempt() bool

/*
type g struct {
0
	stack       stack   // offset known to runtime/cgo
16
	stackguard0 uintptr // offset known to liblink
	stackguard1 uintptr // offset known to liblink
32
	_panic         *_panic // innermost panic - offset known to liblink
	_defer         *_defer // innermost defer
48
	m              *m      // current m; offset known to arm liblink
56
	sched          {
		sp   uintptr
64
		pc   uintptr
		g    guintptr
80
		ctxt unsafe.Pointer
		ret  sys.Uintreg
96
		lr   uintptr
		bp   uintptr // for GOEXPERIMENT=framepointer
112
	}
	syscallsp      uintptr        // if status==Gsyscall, syscallsp = sched.sp to use during gc
	syscallpc      uintptr        // if status==Gsyscall, syscallpc = sched.pc to use during gc
128
	stktopsp       uintptr        // expected sp at top of stack, to check in traceback
	param          unsafe.Pointer // passed parameter on wakeup
144
	atomicstatus   uint32
	stackLock      uint32 // sigprof/scang lock; TODO: fold in to atomicstatus
	goid           int64
160
	schedlink      guintptr
	waitsince      int64      // approx time when the g become blocked
176
	waitreason     waitReason // if status==Gwaiting
177	preempt        bool       // preemption signal, duplicates stackguard0 = stackpreempt

 */