package trace_test

import (
	"github.com/stripe/veneur/v14/trace"
)

// This demonstrates how to switch out an existing
// DefaultClient, closing the existing connection correctly.
func ExampleNewClient() {
	// Create the new client first (we'll create one that can send
	// 20 span packets in parallel):
	cl, err := trace.NewClient(trace.DefaultVeneurAddress, trace.Capacity(20))
	if err != nil {
		panic(err)
	}

	// Replace the old client:
	trace.SetDefaultClient(cl)
	// Output:
}
