# COSC 3750 — Go Programming Assignment  
## Linked Lists, Stacks, and Queues (Intro to Go)

---
## Go File Structure, Packages, and Exports

This assignment uses **multiple Go files** organized into **two packages**. Understanding how Go organizes code into directories and packages is essential for completing the assignment correctly.

Here is the Go documentation on getting started: https://go.dev/doc/tutorial/getting-started

Rember, run this on your Linux VM or the CS Lab Machines before submission. 
---

### Directory Structure for This Assignment

Your project must use the following layout:

```text
hw-linked/
├── go.mod
├── main.go
└── ds/
    ├── node.go
    ├── linkedlist.go
    ├── stack.go
    └── queue.go
```

- `main.go` belongs to the **main package**
- All data structure code (`Node`, `LinkedList`, `Stack`, `Queue`) belongs to the **ds package**
- Each directory corresponds to **exactly one package**

---

### Package Declarations

Each Go source file must declare the correct package at the top.

#### main.go
```go
package main
```

#### All files in the `ds/` directory
```go
package ds
```

All files in the `ds` directory are compiled together as a single package.

---

### How Imports Work in Go

In Go, you **import packages, not individual files**.

The `main` package imports the `ds` package using the module path:

```go
import "hw-linked/ds"
```

(The exact module name comes from the `go.mod` file you create.)
The `go.mod` is created by running the `go mod init` command
For example if you ran `go mod init hw01` your module name would be `hw01`
You should always run `go mod init` first in your directory

Once imported, `main.go` can use exported identifiers from the `ds` package:

```go

var stack ds.Stack
var queue ds.Queue
queue.Pop()
```

Files inside the `ds` package **do not import each other**.  
They automatically have access to all code in the same package.

---

### Visibility and Export Rules

Go controls visibility using capitalization.

An identifier is **exported** if and only if it starts with a capital letter.

This rule applies to:
- Types
- Struct fields
- Functions
- Methods
- Variables
- Constants

Examples:
```go
Node        // exported
LinkedList  // exported
Insert      // exported
removeAt   // unexported
size        // unexported
```

---

### What Must Be Exported for This Assignment

Because `main.go` is in a **different package** from `ds`:

- `Node`, `LinkedList`, `Stack`, and `Queue` **must be exported**
- Any methods called from `main.go` (such as `Push`, `Pop`, `Insert`) **must be exported**

Helper methods used only within the `ds` package should remain **unexported**.

Example:
```go
func (l *LinkedList) Insert(pos int, value int) bool   // exported
func (l *LinkedList) validate() bool                  // unexported
```

---

### What Files Do Not Control

Individual files do **not**:
- define visibility boundaries
- require imports between them
- act as namespaces

Only **packages** control visibility and access in Go.

## Overview

The goal of this assignment is to give you hands-on experience with **Go as a language**.  
This is **not** a systems programming assignment. Instead, you will focus on:

- Go structs
- Pointer usage (safely and idiomatically)
- Method receivers
- Composition
- Basic data-structure design

You will implement a **linked list**, then use it to build a **stack** and a **queue**.

You will create 

---

## Requirements Summary

You will implement **four structs**:

1. `Node`
2. `LinkedList`
3. `Stack`
4. `Queue`

All code should be written in Go.

---

## Part 1 — Node

### Struct Definition

```go
type Node struct {
    data string
    Next  *Node
}
```

---

## Part 2 — LinkedList

### Struct Definition

```go
type LinkedList struct {
    Head *Node
    Tail *Node
    Size int
}
```

### Required Methods

```go
func (l *LinkedList) Insert(value string)// insert at the tail
func (l *LinkedList) InsertAt(position int, value string) error //inserts at a position, returns true if success or false if not, like if pos doesn't exist
func (l *LinkedList) Remove(value string) error // remove first occurrence of the value
func (l *LinkedList) RemoveAll(value string) error //remove all occurrences of a value
func (l *LinkedList) RemoveAt(pos int) error // remove at a position, if index exists
func (l *LinkedList) IsEmpty() bool // checks if the linked list is empty
func (l *LinkedList) GetSize() int // get size of ll
func (l *LinkedList) Reverse() //reverse the list
func (l *LinkedList) PrintList() //print the list 
```

---

## Part 3 — Stack

### Struct Definition

```go
type Stack struct {
    list LinkedList
}
```

### Required Methods

```go
func (s *Stack) Push(value string) // add new head node
func (s *Stack) Pop() (string, bool) // remove the head
```

---

## Part 4 — Queue

### Struct Definition

```go
type Queue struct {
    list LinkedList
}
```

### Required Methods

```go
func (q *Queue) Push(value string) // add tail node
func (q *Queue) Pop() (string, error) // remove the head
```

---



---

## Constraints

- No slices internally
- No recursion
- No generics
- No concurrency

---

## Building/Running 
Go files can be run in two ways:
* Compiled with the `go build` command
    * has the `-o` flag to designate the output file
* run like a script with the `go run` command

Building Example:
```
go build -o hw01
./hw01
```
Running example:
```
go run main.go
```


## Submission

Submit your Go project through GitHub Classroom
