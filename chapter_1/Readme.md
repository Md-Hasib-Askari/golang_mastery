# Go Programming Overview

## What Go Is

* Pragmatic like C
* Structured like Pascal
* Modular like Modula-2
* Opinionated like UNIX
* Strict like a disciplined team lead

---

## What Go Optimizes For

* Reliability
* Simplicity
* Concurrency
* Deployability
* Maintainability

### Go Is **Not** Optimized For

* Fancy academic features
* Complex type hierarchies
* Extreme abstraction layers

---

## Where Go Fits Best

* Backend services
* Infrastructure startups
* Cloud tools
* High-performance APIs
* DevOps tooling

---

## Go Is Strongest In

* Backend engineering
* Distributed systems
* DevOps tooling
* Infrastructure startups
* Cloud-native systems
* High-performance APIs

---

## Go Is NOT Trying To Be

* The best language for game engines
* The best for embedded OS kernels
* The best for functional programming
* The best for hardcore OOP modeling

---

# Important Coding Rules

* Imported packages **must** be used
* Declared variables **must** be used
* Curly braces have **ONE legal style**

## Additional Key Rules

* Blocks always use `{}` — even for one line
* Functions can return multiple values
* No implicit type conversions
* No unnecessary parentheses

---

# Variable Declarations

## 1️⃣ Standard Declaration

```go
var x int
var x int = 10      // With initialization
var x = 10          // Type inferred
```

---

## 2️⃣ Short Declaration Operator `:=`

```go
x := 10
```

### Critical Rule

* Can **only** be used inside functions
* Global declarations must start with a keyword:

```go
var
const
func
```

---

## 3️⃣ Zero Value Rule

| Type    | Zero Value |
| ------- | ---------- |
| int     | 0          |
| string  | ""         |
| float64 | 0.0        |
| pointer | nil        |
| bool    | false      |

---

## 4️⃣ Constants

```go
const Pi = 3.14
```

---

## 5️⃣ Global Variables

```go
var Global int = 1234
```

---

## 6️⃣ Printing

```go
fmt.Println()
fmt.Printf()
```

### Format Specifiers

| Specifier | Meaning               |
| --------- | --------------------- |
| %d        | Integer               |
| %f        | Float                 |
| %.2f      | Float with 2 decimals |
| %s        | String                |

---

# Control Flow

## 1️⃣ The `if` Statement

```go
if condition {
    // code
}
```

### The Most Important Go Pattern

```go
err := someFunction()
if err != nil {
    // handle error
}
```

---

## 2️⃣ Comments

```go
// This is a comment
```

---

## 3️⃣ The `switch` Statement

```go
switch value {
case 1:
    // code
case 2:
    // code
default:
    // code
}
```

# Loops in Go

## 1️⃣ Traditional `for` Loop

```go
for i := 0; i < 10; i++ {
    fmt.Print(i*i, " ")
}
fmt.Println()
```

### Structure

```go
for initialization; condition; post {
    // code
}
```

---

## 2️⃣ `for` Loop as a While Loop

```go
i := 0
for {
    if i == 10 {
        break
    }
    fmt.Print(i*i, " ")
    i++
}
fmt.Println()
```

---

## 3️⃣ Alternative Condition-Based Loop

```go
j := 0
for j != 10 {
    fmt.Print(j*j, " ")
    j++
}
```

### Loop Forms in Go

```go
for condition {}          // while-style loop
for {}                    // infinite loop
for init; cond; post {}   // classic loop
```

> **One keyword. Multiple behaviors.**

---

## 4️⃣ `break` and `continue`

```go
break      // exit loop completely
continue   // skip current iteration
```

---

## 5️⃣ The `range` Keyword

```go
aSlice := []int{-1, 2, 1, -1, 2, -2}

for i, v := range aSlice {
    fmt.Println("index:", i, "value:", v)
}
```

---

## 6️⃣ Ignoring Values

```go
for _, v := range aSlice {
    fmt.Println("value:", v)
}
```

---

## Important Rule Reminder

Even if your loop has only one line:

```go
for condition {
    // still required
}
```

You **must** use curly braces `{}`.

---

# Input in Go

## 1️⃣ Reading from Standard Input

```go
fmt.Printf("Please give me your name: ")

var name string
fmt.Scanln(&name)

fmt.Println("Your name is", name)
```

### Key Points

* `fmt.Scanln()` reads input from the user.
* You must pass the **address** of the variable using `&`.
* Input stops at a newline.

---

## 2️⃣ Command-Line Arguments (`os.Args`)

Command-line arguments are stored in `os.Args`.

### Structure of `os.Args`

| Index | Value              |
| ----- | ------------------ |
| 0     | Path of executable |
| 1     | First argument     |
| 2     | Second argument    |
| ...   | ...                |

### Example

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Program path:", os.Args[0])
    fmt.Println("First argument:", os.Args[1])
}
```

> Remember: Always check the length of `os.Args` before accessing elements to avoid runtime errors.

---