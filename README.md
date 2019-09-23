### Cooperative GC for tight loop Goroutine

Prompt GC whether it needs a preempt in busy looping code.

Beware that this approach is somewhat racy:
 
Two threads can request preempt by setting the flag, we grant
preempt to one, but that nukes the flag of the another request.
This does not seem to happen in practice, but never the less it's a
possibility.

The dependence on unexported field offset in `g` is a horrible hack.
The repo as is will be updated to only work with recent versions, 
now **Go 1.12.x up to git master**).

It's meant only for import into environments which can't afford
modified runtime.

A proper way to do this (and get support for all platforms, go versions
and inlining), is this patch:

```
diff --git a/src/runtime/proc.go b/src/runtime/proc.go
index c06697ef6d..00da4cdbce 100644
--- a/src/runtime/proc.go
+++ b/src/runtime/proc.go
@@ -269,6 +269,12 @@ func Gosched() {
        mcall(gosched_m)
 }

+// True if GC is stuck waiting for preemption of currently running goroutine.
+//go:nosplit
+func WantsPreempt() bool {
+       return getg().preempt
+}
+
 // goschedguarded yields the processor like gosched, but also checks
 // for forbidden states and opts out of the yield in those cases.
 //go:nosplit```
