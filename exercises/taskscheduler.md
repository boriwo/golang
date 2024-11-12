Building a Concurrent Task Scheduler

Build a task scheduler in Go that can execute multiple tasks concurrently. Each task can be represented as a function that takes an integer parameter and returns an integer. Your goal is to create a program that can efficiently execute these tasks concurrently and aggregate their results.

Requirements:

Implement a function func ExecuteTask(task func(int) int, input int) int that takes a task function and an integer input. This function should execute the task function concurrently and return its result.

Implement a function func ConcurrentTaskScheduler(tasks []func(int) int, inputs []int) []int that takes a list of task functions and a list of integers as inputs. This function should execute all the tasks concurrently and return a slice of integers containing their results. The order of results in the output slice should match the order of tasks in the input list.

Ensure that your solution handles errors gracefully, especially in scenarios where tasks may return errors. Implement proper error handling and make sure to convey any errors back to the caller.

Write tests for your functions to ensure they work correctly in different scenarios. Consider testing with various task functions and input values.

Bonus:

Implement a timeout mechanism that allows tasks to execute for a maximum duration. If a task takes longer than a specified timeout, it should be terminated, and the program should proceed with other tasks. Handle timeouts gracefully and ensure that the scheduler doesn't hang indefinitely.