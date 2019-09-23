### Cooperative GC for tight loop Goroutine

Prompt GC whether it needs a preempt in busy looping code.

Beware that this approach is somewhat racy:
 
Two threads can request preempt by setting the flag, we grant
preempt to one, but that nukes the flag of the another request.
This does not seem to happen in practice, but never the less it's a
possibility.

Furthermore, the dependence on unexported field offset in `g` is
horrible a hack, never the less if you want seamless import thats
what this package is for.

A proper way to do this (and get support for all platforms and
inlining), is this simple patch to your Go runtime:

```
diff --git a/src/runtime/proc.go b/src/runtime/proc.go
index c06697ef6d..00da4cdbce 100644
--- a/src/runtime/proc.go
+++ b/src/runtime/proc.go
@@ -269,6 +269,12 @@ func Gosched() {
        mcall(gosched_m)
 }

+// True if a GC is stuck waiting for preempt of currently running goroutine.
+//go:nosplit
+func WantsPreempt() bool {
+       return getg().preempt
+}
+
 // goschedguarded yields the processor like gosched, but also checks
 // for forbidden states and opts out of the yield in those cases.
 //go:nosplit

```