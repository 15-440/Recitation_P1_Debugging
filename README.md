# Examples for the Debugging Recitation
This repository contains the demo code for the debugging recitation of 15-440/15-640. Feel free to play around with it!

## Example 1: Spin Loop
Timeout is a common issue, where there are many potential bugs that give rise to it.

One of the possible reasons is a "spin loop" in the code, which causes the function to consume CPU resources without doing actual work most of the time.

Try running the below in `spinloop/`,
```sh
go run spinloop.go
```
and see the elasped time for the 100000 elements to be sent and received through the channel. It's surprisingly slow!

Now, let's comment line 20, and run the program again. The elasped time should much shorter now.

Why is it the case? With the empty `default`, the `spinLoopConsumer` will continually enter the empty `default` case most of the times, which severely harms the performance of the program.

## Example 2: Deadlock
Deadlock is also a common reason for timeouts.

Sometimes, it's easy to notice deadlocks, because Go shows a fatal error when all goroutines are stuck. Let's run the `go run deadlock.go` and see the error caught by Go.

Now, let's uncomment `go do()` and `fmt.Println(<-ch)`. If you take a look at the code, there's still deadlock inside it. But running the program this time will not show the fatal error, because *not all* goroutines are blocked, as `do()` isn't. The program just hangs and doesn't output anything.

If you type `Ctrl + \` in the terminal now, you will get a printout of all current goroutines and their running/blocked status.

## Example 3: Using the `log` package
The `log` package (https://pkg.go.dev/log) is handy for debugging.

In this example, we show examples of how you can redirect the `Logger` outputs either to nowhere, to a file, to stdout, or to stderr. We created 4 loggers with prefixes `TRACE: `, `INFO: `, `WARNING: `, `ERROR: `, and wrote them to different output writers.

Note that the file `trace.log` is created with the `os.O_APPEND` flag, so if you run the program multiple times, you will see multiple lines in the file. You can explore the flags and use the ones that suit you the most when you are debugging your P1 code.