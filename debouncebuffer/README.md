# Debounce Buffer

## Introduction

A debounce buffer is a data structure used to manage and collect incoming messages or events over a specified period before processing them. The main goal of a debounce buffer is to prevent a function from being called too frequently, thereby reducing the overhead of processing each individual message or event.

This is particularly useful in scenarios where messages or events may arrive in bursts, and it is more efficient to process them in batches rather than individually. For example, when handling rapid successive inputs from a user interface or processing messages from a high-traffic message queue.

## How It Works

The debounce buffer operates by collecting incoming messages over a defined duration (buffer time). Once the buffer time has elapsed, all collected messages are processed together. If new messages continue to arrive within the buffer time, the buffer will reset, and the countdown will start again. This ensures that the function is only called after a period of inactivity or after a specified maximum delay.

### Key Components

1. **Buffer Duration**: The time period during which incoming messages are collected before processing. If a new message arrives within this period, the timer resets.

2. **Message Buffer**: A container (e.g., an array or list) that holds the messages until they are processed.

3. **Function Handler**: The function that is called to process the collected messages after the buffer duration has elapsed.

### Example usage
```go
func main() {
	bufferTime := 5 * time.Second
	messageBufferCache := NewMessageBufferCache(func(msg string) {
		fmt.Println("Handler received:", msg)
	}, bufferTime)

	messageBufferCache.Handle("user1", "Hello")
	messageBufferCache.Handle("user1", "How are you?")

	time.Sleep(3 * time.Second)
	messageBufferCache.Handle("user1", "Still there?")

	time.Sleep(2 * time.Second)
	messageBufferCache.Handle("user2", "Hi there!")

	// Prevent the program from exiting immediately
	time.Sleep(10 * time.Second)
}
```