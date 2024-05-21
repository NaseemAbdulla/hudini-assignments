[10:49 AM] Vishwanatha H M
package main
 
import (
    "fmt"
    "sync"
)
 
// SafeCounter is safe to use concurrently.
type SafeCounter struct {
    mu      sync.Mutex
    counter int
}
 
// Increment increases the counter by 1.
func (sc *SafeCounter) Increment() {
    // sc.mu.Lock()
    sc.counter++
    // sc.mu.Unlock()
}
 
// Value returns the current value of the counter.
func (sc *SafeCounter) Value() int {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    return sc.counter
}
 
func main() {
    sc := SafeCounter{}
    var wg sync.WaitGroup
 
    // Number of goroutines to increment the counter
    numGoroutines := 1000
 
    // Increment the counter in multiple goroutines
    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            sc.Increment()
        }()
    }
 
    // Wait for all goroutines to finish
    wg.Wait()
 
    // Print the final counter value
    fmt.Println("Final Counter Value:", sc.Value())
}
 