## 3. STRUCTS

### Explanation
Structs are composite data types that group related data fields together. They're the foundation of Go's object-oriented programming approach.

### How It Works Internally
- Structs are laid out in memory with fields in declaration order
- Memory padding occurs for alignment optimization
- No inheritance; composition is used instead
- Methods are separate from struct definition
- Pointer receivers modify the struct; value receivers don't

### Code Examples

```go
package main

import "fmt"

// Define a struct
type Book struct {
    Title  string
    Author string
    Pages  int
    Price  float64
}

// Struct with embedded struct (composition)
type Author struct {
    Name    string
    Country string
}

type Novel struct {
    Title  string
    Author Author
    Pages  int
}

// Methods with value receiver (doesn't modify original)
func (b Book) String() string {
    return fmt.Sprintf("%s by %s (%d pages)", b.Title, b.Author, b.Pages)
}

// Methods with pointer receiver (can modify original)
func (b *Book) ApplyDiscount(percent float64) {
    b.Price = b.Price * (1 - percent/100)
}

// Constructor function
func NewBook(title, author string, pages int, price float64) *Book {
    return &Book{title, author, pages, price}
}

func main() {
    // Creating structs
    book1 := Book{"Go in Action", "William Kennedy", 450, 50.00}
    book2 := NewBook("The Go Way", "John Doe", 380, 45.00)
    
    fmt.Println(book1.String())
    
    book2.ApplyDiscount(10)
    fmt.Printf("Price after discount: %.2f\n", book2.Price)
    
    // Nested structs
    novel := Novel{
        Title: "1984",
        Author: Author{"George Orwell", "UK"},
        Pages: 328,
    }
    fmt.Printf("%s by %s from %s\n", novel.Title, novel.Author.Name, novel.Author.Country)
}
```

### Mini-Project: Library Management System
Create a simple library system with books:

```go
package main

import "fmt"

type Book struct {
    ID       int
    Title    string
    Author   string
    Available bool
}

type Library struct {
    Name  string
    Books []Book
}

func (lib *Library) AddBook(book Book) {
    lib.Books = append(lib.Books, book)
    fmt.Printf("Added: %s\n", book.Title)
}

func (lib *Library) BorrowBook(id int) error {
    for i := range lib.Books {
        if lib.Books[i].ID == id {
            if !lib.Books[i].Available {
                return fmt.Errorf("book already borrowed")
            }
            lib.Books[i].Available = false
            return nil
        }
    }
    return fmt.Errorf("book not found")
}

func (lib *Library) ReturnBook(id int) error {
    for i := range lib.Books {
        if lib.Books[i].ID == id {
            lib.Books[i].Available = true
            return nil
        }
    }
    return fmt.Errorf("book not found")
}

func (lib Library) ListBooks() {
    for _, book := range lib.Books {
        status := "Available"
        if !book.Available {
            status = "Borrowed"
        }
        fmt.Printf("[%d] %s by %s - %s\n", book.ID, book.Title, book.Author, status)
    }
}

func main() {
    lib := Library{Name: "City Library"}
    
    lib.AddBook(Book{1, "Go Programming", "John Doe", true})
    lib.AddBook(Book{2, "Clean Code", "Robert Martin", true})
    lib.AddBook(Book{3, "Design Patterns", "Gang of Four", true})
    
    lib.ListBooks()
    
    lib.BorrowBook(1)
    fmt.Println("After borrowing:")
    lib.ListBooks()
}
```

---

## 4. INTERFACES

### Explanation
Interfaces define a set of method signatures that types must implement. They enable polymorphism and loose coupling in Go.

### How It Works Internally
- Interfaces are stored as (type, value) pairs internally
- Dynamic dispatch determines which method to call at runtime
- Interface types use vtable-like mechanism for method lookup
- Type assertions check the concrete type of interface values
- Zero interface{} can hold any type

### Code Examples

```go
package main

import "fmt"

// Define an interface
type Writer interface {
    Write(data string) error
}

type Reader interface {
    Read() (string, error)
}

// ReadWriter combines two interfaces
type ReadWriter interface {
    Reader
    Writer
}

// Concrete type implementing Writer
type FileWriter struct {
    name string
}

func (fw FileWriter) Write(data string) error {
    fmt.Printf("Writing to file %s: %s\n", fw.name, data)
    return nil
}

// Another concrete type
type NetworkWriter struct {
    address string
}

func (nw NetworkWriter) Write(data string) error {
    fmt.Printf("Sending over network to %s: %s\n", nw.address, data)
    return nil
}

// Function accepting interface (polymorphism)
func SaveData(writer Writer, data string) error {
    return writer.Write(data)
}

// Type assertion
func PrintConcreteType(w Writer) {
    switch v := w.(type) {
    case FileWriter:
        fmt.Printf("FileWriter to: %s\n", v.name)
    case NetworkWriter:
        fmt.Printf("NetworkWriter to: %s\n", v.address)
    default:
        fmt.Println("Unknown writer type")
    }
}

func main() {
    writers := []Writer{
        FileWriter{"data.txt"},
        NetworkWriter{"192.168.1.1"},
    }
    
    for _, w := range writers {
        SaveData(w, "Hello, World!")
        PrintConcreteType(w)
    }
}
```

### Mini-Project: Payment Processing System
Create a system that accepts different payment methods:

```go
package main

import "fmt"

type PaymentProcessor interface {
    Process(amount float64) error
    Refund(amount float64) error
}

type CreditCard struct {
    CardNumber string
    Balance    float64
}

func (cc *CreditCard) Process(amount float64) error {
    if cc.Balance < amount {
        return fmt.Errorf("insufficient balance")
    }
    cc.Balance -= amount
    fmt.Printf("Charged %f to credit card %s\n", amount, cc.CardNumber)
    return nil
}

func (cc *CreditCard) Refund(amount float64) error {
    cc.Balance += amount
    fmt.Printf("Refunded %f to credit card\n", amount)
    return nil
}

type PayPal struct {
    Email   string
    Balance float64
}

func (p *PayPal) Process(amount float64) error {
    if p.Balance < amount {
        return fmt.Errorf("insufficient balance")
    }
    p.Balance -= amount
    fmt.Printf("Charged %f via PayPal (%s)\n", amount, p.Email)
    return nil
}

func (p *PayPal) Refund(amount float64) error {
    p.Balance += amount
    fmt.Printf("Refunded %f via PayPal\n", amount)
    return nil
}

func ProcessPayment(processor PaymentProcessor, amount float64) {
    if err := processor.Process(amount); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}

func main() {
    cc := &CreditCard{"1234-5678-9000", 1000}
    pp := &PayPal{"user@example.com", 500}
    
    ProcessPayment(cc, 100)
    ProcessPayment(pp, 50)
    ProcessPayment(cc, 2000) // This will fail
}
```

---

## 5. ERRORS

### Explanation
Errors in Go are values that implement the error interface. They represent exceptional conditions that need handling.

### How It Works Internally
- error interface has one method: Error() string
- Errors are checked explicitly, not thrown/caught
- panic and recover provide exception-like behavior for severe errors
- Stack unwinding occurs automatically on panic
- defer statements execute before returning from function

### Code Examples

```go
package main

import (
    "errors"
    "fmt"
)

// Custom error type
type ValidationError struct {
    Field string
    Msg   string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error on %s: %s", e.Field, e.Msg)
}

// Function returning error
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Function with custom error
func validateEmail(email string) error {
    if email == "" {
        return ValidationError{"email", "cannot be empty"}
    }
    if len(email) < 5 {
        return ValidationError{"email", "too short"}
    }
    return nil
}

// Wrapping errors (Go 1.13+)
func processFile(filename string) error {
    if filename == "" {
        return fmt.Errorf("failed to process: %w", errors.New("empty filename"))
    }
    return nil
}

// Using Is and As (Go 1.13+)
func handleError(err error) {
    if errors.Is(err, errors.New("test")) {
        fmt.Println("Specific error")
    }
    
    var valErr ValidationError
    if errors.As(err, &valErr) {
        fmt.Printf("Validation error on field: %s\n", valErr.Field)
    }
}

// Panic and recover
func safeDivide(a, b float64) (result float64) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
            result = 0
        }
    }()
    
    if b == 0 {
        panic("division by zero!")
    }
    return a / b
}

func main() {
    // Checking errors
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
    
    // Custom errors
    if err := validateEmail("test@example.com"); err != nil {
        fmt.Println(err)
    }
    
    // Safe operations
    safeDivide(10, 0)
}
```

### Mini-Project: Form Validator
Create a form validation system with detailed errors:

```go
package main

import (
    "errors"
    "fmt"
    "regexp"
)

type FieldError struct {
    Field   string
    Message string
}

type ValidationErrors []FieldError

func (ve ValidationErrors) Error() string {
    if len(ve) == 0 {
        return "no validation errors"
    }
    msg := "validation failed:\n"
    for _, e := range ve {
        msg += fmt.Sprintf("  %s: %s\n", e.Field, e.Message)
    }
    return msg
}

type User struct {
    Name  string
    Email string
    Age   int
}

func ValidateUser(u User) error {
    var errs ValidationErrors
    
    if u.Name == "" {
        errs = append(errs, FieldError{"name", "required"})
    }
    
    if u.Email == "" {
        errs = append(errs, FieldError{"email", "required"})
    } else {
        emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
        if !emailRegex.MatchString(u.Email) {
            errs = append(errs, FieldError{"email", "invalid format"})
        }
    }
    
    if u.Age < 18 {
        errs = append(errs, FieldError{"age", "must be 18 or older"})
    }
    
    if len(errs) > 0 {
        return errs
    }
    return nil
}

func main() {
    user := User{"John", "invalid-email", 15}
    if err := ValidateUser(user); err != nil {
        fmt.Println(err)
    }
}
```

---

## 6. LOOPS

### Explanation
Loops repeat a block of code multiple times. Go has only one loop construct: `for`.

### How It Works Internally
- Loop condition is evaluated before each iteration (or after for do-while pattern)
- Break jumps to code after loop
- Continue jumps to next iteration
- Compiler optimizes loops with inlining and bounds-check elimination
- Range loops create iterators internally

### Code Examples

```go
package main

import "fmt"

func main() {
    // Traditional for loop
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    
    // While-style loop
    counter := 0
    for counter < 5 {
        fmt.Println(counter)
        counter++
    }
    
    // Infinite loop with break
    for {
        fmt.Println("Breaking out...")
        break
    }
    
    // Range over slice
    numbers := []int{1, 2, 3, 4, 5}
    for i, val := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", i, val)
    }
    
    // Range over string
    str := "Hello"
    for i, char := range str {
        fmt.Printf("%d: %c\n", i, char)
    }
    
    // Range over map
    m := map[string]int{"a": 1, "b": 2, "c": 3}
    for key, value := range m {
        fmt.Printf("%s: %d\n", key, value)
    }
    
    // Continue and break
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            continue // Skip even numbers
        }
        if i > 5 {
            break // Exit loop
        }
        fmt.Println(i)
    }
    
    // Labeled loops
    outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                continue outer // Continue outer loop
            }
            fmt.Printf("(%d,%d) ", i, j)
        }
    }
}
```

### Mini-Project: Number Pattern Generator
Create patterns with loops:

```go
package main

import "fmt"

func pyramidPattern(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("* ")
        }
        fmt.Println()
    }
}

func diamondPattern(n int) {
    // Upper half
    for i := 1; i <= n; i++ {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        for j := 1; j <= 2*i-1; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
    
    // Lower half
    for i := n - 1; i >= 1; i-- {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        for j := 1; j <= 2*i-1; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}

func multiplicationTable(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            fmt.Printf("%3d ", i*j)
        }
        fmt.Println()
    }
}

func main() {
    fmt.Println("Pyramid:")
    pyramidPattern(5)
    
    fmt.Println("\nDiamond:")
    diamondPattern(4)
    
    fmt.Println("\nMultiplication Table:")
    multiplicationTable(5)
}
```

---

## 7. SLICES

### Explanation
Slices are dynamic arrays that can grow and shrink. They're more flexible than arrays and are the most common data structure in Go.

### How It Works Internally
- Slice has three fields: pointer to data, length, capacity
- Appending may trigger allocation if capacity exceeded
- Slicing creates new slice header pointing to same underlying array
- Copy copies actual data; assigning copies slice header
- Capacity is often 2x current length to reduce allocations

### Code Examples

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // Create slice with literal
    nums := []int{1, 2, 3, 4, 5}
    
    // Create slice with make
    slice1 := make([]int, 5)        // length 5, capacity 5
    slice2 := make([]int, 5, 10)    // length 5, capacity 10
    
    // Append to slice
    nums = append(nums, 6, 7, 8)
    fmt.Println(nums)
    
    // Slicing (start:end) - excludes end index
    portion := nums[1:4]
    fmt.Println(portion) // [2 3 4]
    
    // Slicing with capacity specification
    subslice := nums[1:4:6] // [start:end:capacity]
    
    // Copy slice
    dest := make([]int, len(nums))
    copy(dest, nums)
    
    // Slice operations
    fmt.Println("Length:", len(nums))
    fmt.Println("Capacity:", cap(nums))
    
    // Slices share underlying array
    original := []int{1, 2, 3, 4, 5}
    view1 := original[0:3]
    view1[0] = 999
    fmt.Println(original) // [999 2 3 4 5] - modified!
    
    // Sorting
    nums = []int{3, 1, 4, 1, 5, 9}
    sort.Ints(nums)
    fmt.Println(nums)
    
    // Finding element
    idx := sort.SearchInts(nums, 4)
    
    // Removing element
    nums = append(nums[:2], nums[3:]...) // Remove index 2
    
    // Clearing slice
    clear := make([]int, 5)
    
    // Strings slice
    str := "hello"
    chars := []rune(str)
    fmt.Println(chars)
}
```

### Mini-Project: Dynamic Array Utilities
Create utility functions for slices:

```go
package main

import (
    "fmt"
    "sort"
)

// RemoveDuplicates returns slice without duplicates
func RemoveDuplicates(nums []int) []int {
    sort.Ints(nums)
    result := make([]int, 0, len(nums))
    
    for i, v := range nums {
        if i == 0 || nums[i-1] != v {
            result = append(result, v)
        }
    }
    return result
}

// Reverse reverses slice in-place
func Reverse(nums []int) {
    for i := 0; i < len(nums)/2; i++ {
        j := len(nums) - 1 - i
        nums[i], nums[j] = nums[j], nums[i]
    }
}

// Merge merges two sorted slices
func Merge(a, b []int) []int {
    result := make([]int, 0, len(a)+len(b))
    i, j := 0, 0
    
    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            result = append(result, a[i])
            i++
        } else {
            result = append(result, b[j])
            j++
        }
    }
    result = append(result, a[i:]...)
    result = append(result, b[j:]...)
    return result
}

// Filter keeps elements that pass predicate
func Filter(nums []int, predicate func(int) bool) []int {
    result := make([]int, 0)
    for _, v := range nums {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}

// Map applies function to each element
func Map(nums []int, transform func(int) int) []int {
    result := make([]int, len(nums))
    for i, v := range nums {
        result[i] = transform(v)
    }
    return result
}

func main() {
    fmt.Println("Remove Duplicates:")
    fmt.Println(RemoveDuplicates([]int{3, 1, 4, 1, 5, 9, 2, 6, 5}))
    
    fmt.Println("Reverse:")
    nums := []int{1, 2, 3, 4, 5}
    Reverse(nums)
    fmt.Println(nums)
    
    fmt.Println("Merge:")
    fmt.Println(Merge([]int{1, 3, 5}, []int{2, 4, 6}))
    
    fmt.Println("Filter evens:")
    fmt.Println(Filter([]int{1, 2, 3, 4, 5, 6}, func(x int) bool { return x%2 == 0 }))
    
    fmt.Println("Map square:")
    fmt.Println(Map([]int{1, 2, 3, 4}, func(x int) int { return x * x }))
}
```

---

## 8. MAP

### Explanation
Maps are unordered collections of key-value pairs. They provide O(1) average lookup time.

### How It Works Internally
- Maps use hash tables internally
- Hash function converts key to bucket index
- Collision resolution uses chaining or open addressing
- Maps are not safe for concurrent reads/writes
- Iteration order is randomized for security
- Grow dynamically when load factor exceeds threshold

### Code Examples

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // Create map with literal
    m1 := map[string]int{"a": 1, "b": 2, "c": 3}
    
    // Create empty map
    m2 := make(map[string]string)
    
    // Add elements
    m2["key1"] = "value1"
    m2["key2"] = "value2"
    
    // Access value
    val := m1["a"]
    fmt.Println(val)
    
    // Check if key exists
    if val, ok := m1["x"]; ok {
        fmt.Println("Found:", val)
    } else {
        fmt.Println("Key not found")
    }
    
    // Delete key
    delete(m1, "b")
    
    // Get length
    fmt.Println("Length:", len(m1))
    
    // Iterate map (order is random)
    for key, value := range m1 {
        fmt.Printf("%s: %d\n", key, value)
    }
    
    // Maps of maps
    nested := map[string]map[string]int{
        "row1": {"col1": 1, "col2": 2},
        "row2": {"col1": 3, "col2": 4},
    }
    
    fmt.Println(nested["row1"]["col2"])
    
    // Map with interface{} as value
    config := map[string]interface{}{
        "name":   "app",
        "port":   8080,
        "debug":  true,
        "tags":   []string{"web", "api"},
    }
    
    for k, v := range config {
        fmt.Printf("%s: %v (%T)\n", k, v, v)
    }
    
    // Sorting map keys
    m := map[string]int{"c": 3, "a": 1, "b": 2}
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, k := range keys {
        fmt.Printf("%s: %d\n", k, m[k])
    }
}
```

### Mini-Project: Word Frequency Counter
Count word frequencies in text:

```go
package main

import (
    "fmt"
    "sort"
    "strings"
)

type WordCount struct {
    Word  string
    Count int
}

func CountWords(text string) map[string]int {
    words := strings.Fields(strings.ToLower(text))
    counts := make(map[string]int)
    
    for _, word := range words {
        // Remove punctuation
        word = strings.Trim(word, ".,!?;:")
        counts[word]++
    }
    return counts
}

func SortByFrequency(counts map[string]int) []WordCount {
    result := make([]WordCount, 0, len(counts))
    
    for word, count := range counts {
        result = append(result, WordCount{word, count})
    }
    
    sort.Slice(result, func(i, j int) bool {
        return result[i].Count > result[j].Count
    })
    
    return result
}

func main() {
    text := `Go is a programming language. Go is fast. Go is simple.
             Programming in Go is fun. Go has great tools.`
    
    counts := CountWords(text)
    sorted := SortByFrequency(counts)
    
    fmt.Println("Word Frequency:")
    for i, wc := range sorted {
        if i < 5 { // Top 5
            fmt.Printf("%s: %d\n", wc.Word, wc.Count)
        }
    }
}
```

---

## 9. ADVANCED FUNCTIONS

### Explanation
Advanced function concepts include closures, higher-order functions, and decorators that enable functional programming patterns.

### How It Works Internally
- Closures capture variables by reference from enclosing scope
- Function types enable passing functions as arguments
- Deferred functions are pushed onto defer stack during execution
- Method values create function values with receiver bound
- Inlining optimizes small function calls

### Code Examples

```go
package main

import "fmt"

// Closure - function accessing outer scope variables
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// Higher-order function - returns a function
func multiply(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

// Function decorator
func withLogging(fn func(string) string) func(string) string {
    return func(input string) string {
        fmt.Printf("Calling function with input: %s\n", input)
        result := fn(input)
        fmt.Printf("Function returned: %s\n", result)
        return result
    }
}

// Pipeline - composing functions
func pipeline(value int, funcs ...func(int) int) int {
    for _, fn := range funcs {
        value = fn(value)
    }
    return value
}

// Reduce - fold operation
func reduce(nums []int, initial int, accumulator func(int, int) int) int {
    result := initial
    for _, v := range nums {
        result = accumulator(result, v)
    }
    return result
}

// Memoization - caching function results
func memoize(fn func(int) int) func(int) int {
    cache := make(map[int]int)
    return func(n int) int {
        if val, ok := cache[n]; ok {
            return val
        }
        result := fn(n)
        cache[n] = result
        return result
    }
}

// Fibonacci with memoization
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    // Closure example
    c := counter()
    fmt.Println(c()) // 1
    fmt.Println(c()) // 2
    fmt.Println(c()) // 3
    
    // Higher-order function
    double := multiply(2)
    triple := multiply(3)
    fmt.Println(double(5)) // 10
    fmt.Println(triple(5)) // 15
    
    // Decorator
    greet := func(name string) string {
        return fmt.Sprintf("Hello, %s!", name)
    }
    decoratedGreet := withLogging(greet)
    decoratedGreet("Alice")
    
    // Pipeline
    result := pipeline(5,
        func(x int) int { return x * 2 },
        func(x int) int { return x + 10 },
        func(x int) int { return x / 3 },
    )
    fmt.Println("Pipeline result:", result)
    
    // Reduce
    nums := []int{1, 2, 3, 4, 5}
    sum := reduce(nums, 0, func(acc, val int) int { return acc + val })
    product := reduce(nums, 1, func(acc, val int) int { return acc * val })
    fmt.Println("Sum:", sum, "Product:", product)
    
    // Memoization
    fib := memoize(fibonacci)
    fmt.Println("Fibonacci(10):", fib(10))
}
```

### Mini-Project: Function Middleware Chain
Build a middleware system:

```go
package main

import (
    "fmt"
    "time"
)

type Handler func(string) string

// Middleware that times execution
func withTiming(handler Handler) Handler {
    return func(input string) string {
        start := time.Now()
        result := handler(input)
        fmt.Printf("Executed in %v\n", time.Since(start))
        return result
    }
}

// Middleware that validates input
func withValidation(handler Handler) Handler {
    return func(input string) string {
        if input == "" {
            return "Error: input cannot be empty"
        }
        return handler(input)
    }
}

// Middleware that caches results
func withCaching(handler Handler) Handler {
    cache := make(map[string]string)
    return func(input string) string {
        if result, ok := cache[input]; ok {
            fmt.Println("Cache hit")
            return result
        }
        result := handler(input)
        cache[input] = result
        return result
    }
}

// Chain middlewares
func chain(handler Handler, middlewares ...func(Handler) Handler) Handler {
    for _, mw := range middlewares {
        handler = mw(handler)
    }
    return handler
}

// Simple handler
func processString(input string) string {
    time.Sleep(100 * time.Millisecond) // Simulate work
    return fmt.Sprintf("Processed: %s", input)
}

func main() {
    handler := chain(
        processString,
        withValidation,
        withCaching,
        withTiming,
    )
    
    handler("hello")
    handler("")  // Validation will catch this
    handler("hello") // Cache hit
}
```

---

## 10. POINTERS

### Explanation
Pointers store memory addresses of variables. They enable pass-by-reference and efficient data manipulation.

### How It Works Internally
- Pointers are memory addresses (8 bytes on 64-bit systems)
- & operator gets address of variable
- * operator dereferences pointer to access value
- Go performs pointer escape analysis to determine stack vs heap allocation
- Pointer nil value is zero address
- Slice and map headers contain pointers to data

### Code Examples

```go
package main

import "fmt"

func main() {
    // Creating pointers
    x := 42
    ptr := &x           // Address of x
    fmt.Println(ptr)    // 0xc000...
    fmt.Println(*ptr)   // 42
    
    // Modifying through pointer
    *ptr = 100
    fmt.Println(x)      // 100
    
    // Pointer to pointer
    ptrToPtr := &ptr
    fmt.Println(**ptrToPtr) // 100
    
    // Array pointers
    arr := [3]int{1, 2, 3}
    arrPtr := &arr
    fmt.Println(arrPtr[0]) // 1 - can index array pointers
    
    // Nil pointer
    var nilPtr *int
    if nilPtr == nil {
        fmt.Println("Pointer is nil")
    }
    
    // Pointer receivers (methods that modify receiver)
    type Person struct {
        Name string
        Age  int
    }
    
    p := Person{"Alice", 30}
    increaseAge(&p)
    fmt.Println(p.Age) // 31
    
    // Passing by value vs reference
    original := 10
    modifyByValue(original)      // Doesn't affect original
    modifyByReference(&original) // Affects original
    fmt.Println(original)         // 20
    
    // Pointer arithmetic (rare in Go, usually use slices)
    nums := []int{1, 2, 3, 4, 5}
    firstPtr := &nums[0]
    fmt.Println(*firstPtr) // 1
}

func increaseAge(p *Person) {
    p.Age++
}

func modifyByValue(x int) {
    x = 20 // Only modifies local copy
}

func modifyByReference(x *int) {
    *x = 20 // Modifies original
}
```

### Mini-Project: Binary Tree Implementation
Build a tree using pointers:

```go
package main

import "fmt"

type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

func NewNode(val int) *TreeNode {
    return &TreeNode{Value: val}
}

func (n *TreeNode) Insert(val int) {
    if val < n.Value {
        if n.Left == nil {
            n.Left = NewNode(val)
        } else {
            n.Left.Insert(val)
        }
    } else {
        if n.Right == nil {
            n.Right = NewNode(val)
        } else {
            n.Right.Insert(val)
        }
    }
}

func (n *TreeNode) Search(val int) bool {
    if n == nil {
        return false
    }
    if n.Value == val {
        return true
    }
    if val < n.Value {
        return n.Left.Search(val)
    }
    return n.Right.Search(val)
}

func (n *TreeNode) InOrder() {
    if n == nil {
        return
    }
    n.Left.InOrder()
    fmt.Printf("%d ", n.Value)
    n.Right.InOrder()
}

func main() {
    root := NewNode(50)
    root.Insert(30)
    root.Insert(70)
    root.Insert(20)
    root.Insert(40)
    root.Insert(60)
    root.Insert(80)
    
    fmt.Println("In-order traversal:")
    root.InOrder()
    fmt.Println()
    
    fmt.Println("Search 40:", root.Search(40))
    fmt.Println("Search 25:", root.Search(25))
}
```

---

## 11. LOCAL DEVELOPMENT

### Explanation
Local development involves setting up your Go environment, managing project structure, and using Go tools effectively.

### How It Works Internally
- Go modules (go.mod) manage dependencies
- Import paths map to filesystem locations
- Go compiler builds executables from source
- go.sum ensures reproducible builds
- Package init() functions run at import time

### Code Examples

```go
// Project structure:
// myproject/
// ├── go.mod
// ├── go.sum
// ├── main.go
// ├── config/
// │   └── config.go
// └── utils/
//     └── helpers.go

// go.mod example:
// module myapp.com/myproject
// 
// go 1.21
//
// require (
//     github.com/google/uuid v1.3.0
// )

// main.go
package main

import (
    "fmt"
    "myapp.com/myproject/config"
    "myapp.com/myproject/utils"
)

func init() {
    fmt.Println("Initializing application...")
}

func main() {
    cfg := config.Load()
    result := utils.Process(cfg)
    fmt.Println(result)
}

// config/config.go
package config

type Config struct {
    AppName string
    Port    int
}

func Load() Config {
    return Config{
        AppName: "MyApp",
        Port:    8080,
    }
}

// utils/helpers.go
package utils

import "myapp.com/myproject/config"

func Process(cfg config.Config) string {
    return cfg.AppName
}
```

### Mini-Project: Multi-Package Application
Create a complete project structure:

```go
// Project: task-manager

// main.go
package main

import (
    "fmt"
    "task-manager/tasks"
)

func main() {
    manager := tasks.NewManager()
    
    manager.Add("Learn Go", 1)
    manager.Add("Build project", 2)
    manager.Add("Master concurrency", 3)
    
    manager.ListAll()
    
    manager.Complete(1)
    manager.ListPending()
}

// tasks/manager.go
package tasks

import "fmt"

type Task struct {
    ID       int
    Title    string
    Complete bool
}

type Manager struct {
    tasks []Task
    nextID int
}

func NewManager() *Manager {
    return &Manager{
        tasks:  make([]Task, 0),
        nextID: 1,
    }
}

func (m *Manager) Add(title string, priority int) {
    m.tasks = append(m.tasks, Task{
        ID:       m.nextID,
        Title:    title,
        Complete: false,
    })
    m.nextID++
}

func (m *Manager) Complete(id int) {
    for i := range m.tasks {
        if m.tasks[i].ID == id {
            m.tasks[i].Complete = true
            fmt.Printf("Completed: %s\n", m.tasks[i].Title)
            return
        }
    }
}

func (m *Manager) ListAll() {
    fmt.Println("All Tasks:")
    for _, task := range m.tasks {
        status := "[ ]"
        if task.Complete {
            status = "[x]"
        }
        fmt.Printf("%s %d. %s\n", status, task.ID, task.Title)
    }
}

func (m *Manager) ListPending() {
    fmt.Println("Pending Tasks:")
    for _, task := range m.tasks {
        if !task.Complete {
            fmt.Printf("[ ] %d. %s\n", task.ID, task.Title)
        }
    }
}
```

---

## 12. CHANNELS AND CONCURRENCY

### Explanation
Channels enable safe communication between goroutines. They're the Go way of coordinating concurrent operations.

### How It Works Internally
- Channels use internal queues to store values
- Sends block if buffer is full
- Receives block if channel is empty
- Close broadcasts signal to all receivers
- Select multiplexes multiple channel operations
- Mutexes protect shared memory access
- Atomic operations provide lock-free synchronization

### Code Examples

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    // Unbuffered channel - synchronous
    unbuffered := make(chan int)
    go func() {
        unbuffered <- 42 // Blocks until received
    }()
    val := <-unbuffered
    fmt.Println("Received:", val)
    
    // Buffered channel
    buffered := make(chan string, 2)
    buffered <- "first"
    buffered <- "second"
    
    // Receive from channel
    fmt.Println(<-buffered) // "first"
    fmt.Println(<-buffered) // "second"
    
    // Close channel (signaling)
    c := make(chan int, 3)
    c <- 1
    c <- 2
    c <- 3
    close(c)
    
    for val := range c {
        fmt.Println("Got:", val)
    }
    
    // Select - multiplexing
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(100 * time.Millisecond)
        ch1 <- "one"
    }()
    
    go func() {
        time.Sleep(200 * time.Millisecond)
        ch2 <- "two"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Got from ch1:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Got from ch2:", msg2)
        }
    }
    
    // WaitGroup - synchronization
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Goroutine %d working\n", id)
        }(i)
    }
    wg.Wait()
    fmt.Println("All goroutines done")
    
    // Mutex - protecting shared data
    var mu sync.Mutex
    var counter int
    
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            counter++
            mu.Unlock()
        }()
    }
    wg.Wait()
    fmt.Println("Counter:", counter)
}
```

### Mini-Project: Worker Pool
Implement a job queue with workers:

```go
package main

import (
    "fmt"
    "sync"
)

type Job struct {
    ID   int
    Data string
}

type Result struct {
    JobID int
    Data  string
}

type WorkerPool struct {
    jobs    chan Job
    results chan Result
    wg      sync.WaitGroup
}

func NewWorkerPool(numWorkers int) *WorkerPool {
    wp := &WorkerPool{
        jobs:    make(chan Job, 100),
        results: make(chan Result, 100),
    }
    
    for i := 0; i < numWorkers; i++ {
        wp.wg.Add(1)
        go wp.worker(i)
    }
    
    return wp
}

func (wp *WorkerPool) worker(id int) {
    defer wp.wg.Done()
    for job := range wp.jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job.ID)
        result := Result{
            JobID: job.ID,
            Data:  job.Data + " processed",
        }
        wp.results <- result
    }
}

func (wp *WorkerPool) Submit(job Job) {
    wp.jobs <- job
}

func (wp *WorkerPool) Close() {
    close(wp.jobs)
    wp.wg.Wait()
    close(wp.results)
}

func (wp *WorkerPool) Results() <-chan Result {
    return wp.results
}

func main() {
    pool := NewWorkerPool(3)
    
    // Submit jobs
    for i := 0; i < 10; i++ {
        pool.Submit(Job{ID: i, Data: fmt.Sprintf("task%d", i)})
    }
    
    pool.Close()
    
    // Collect results
    for result := range pool.Results() {
        fmt.Printf("Result: Job %d -> %s\n", result.JobID, result.Data)
    }
}
```

---

## 13. MUTEX CONCURRENCY

### Explanation
Mutexes (mutual exclusion) protect shared data from concurrent access. RWMutex allows multiple readers but exclusive writer access.

### How It Works Internally
- Mutex uses atomic operations for lock/unlock
- Lock blocks waiting threads in OS scheduler
- Goroutines yielding allows others to run
- RWMutex maintains reader/writer state counters
- Fairness ensured through OS kernel support

### Code Examples

```go
package main

import (
    "fmt"
    "sync"
)

// Using Mutex
type SafeCounter struct {
    mu    sync.Mutex
    value int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock() // Ensures unlock even if panic
    c.value++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

// Using RWMutex for read-heavy workloads
type SafeMap struct {
    mu      sync.RWMutex
    data    map[string]int
}

func (sm *SafeMap) Set(key string, value int) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
    sm.mu.RLock()
    defer sm.mu.RUnlock()
    val, ok := sm.data[key]
    return val, ok
}

// Condition variable - coordinating goroutines
func main() {
    // Basic mutex example
    counter := &SafeCounter{}
    
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    
    wg.Wait()
    fmt.Println("Final count:", counter.Value())
    
    // RWMutex example
    safeMap := &SafeMap{data: make(map[string]int)}
    
    // Writers
    go func() {
        for i := 0; i < 10; i++ {
            safeMap.Set(fmt.Sprintf("key%d", i), i)
        }
    }()
    
    // Readers
    for i := 0; i < 5; i++ {
        go func() {
            for j := 0; j < 10; j++ {
                safeMap.Get(fmt.Sprintf("key%d", j))
            }
        }()
    }
}
```

### Mini-Project: Rate Limiter
Implement concurrent rate limiting:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type RateLimiter struct {
    mu        sync.Mutex
    maxRequests int
    window    time.Duration
    requests  []time.Time
}

func NewRateLimiter(maxRequests int, window time.Duration) *RateLimiter {
    return &RateLimiter{
        maxRequests: maxRequests,
        window:      window,
        requests:    make([]time.Time, 0),
    }
}

func (rl *RateLimiter) Allow() bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    now := time.Now()
    // Remove old requests outside window
    validIdx := 0
    for i, t := range rl.requests {
        if now.Sub(t) <= rl.window {
            validIdx = i
            break
        }
    }
    rl.requests = rl.requests[validIdx:]
    
    if len(rl.requests) < rl.maxRequests {
        rl.requests = append(rl.requests, now)
        return true
    }
    return false
}

func main() {
    limiter := NewRateLimiter(5, time.Second)
    
    // Try 10 requests
    for i := 0; i < 10; i++ {
        if limiter.Allow() {
            fmt.Printf("Request %d: Allowed\n", i+1)
        } else {
            fmt.Printf("Request %d: Rate limited\n", i+1)
        }
    }
    
    fmt.Println("Waiting 1 second...")
    time.Sleep(time.Second)
    
    // More requests
    for i := 0; i < 3; i++ {
        if limiter.Allow() {
            fmt.Printf("Request %d: Allowed\n", 10+i+1)
        }
    }
}
```

---

## 14. GENERICS

### Explanation
Generics (Go 1.18+) allow writing type-safe code that works with different types without repeating logic.

### How It Works Internally
- Type parameters are substituted at compile time
- Compiler generates specialized code for each type used (monomorphization)
- Constraints define what operations are allowed on type parameters
- Empty interface{} becomes necessary in pre-1.18 Go
- Generic instances are inlined by optimizer

### Code Examples

```go
package main

import "fmt"

// Generic function with constraint
func Print[T any](value T) {
    fmt.Println(value)
}

// Generic slice utilities
func Map[T any, U any](items []T, transform func(T) U) []U {
    result := make([]U, len(items))
    for i, item := range items {
        result[i] = transform(item)
    }
    return result
}

// Constraint - types must implement Comparable
type Comparable interface {
    int | int64 | float64 | string
}

func Max[T Comparable](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Generic struct
type Pair[T, U any] struct {
    First  T
    Second U
}

func (p Pair[T, U]) String() string {
    return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}

// Receiver constraint using interface
type Number interface {
    int | int64 | float64
}

func (p Pair[T, U]) Swap() Pair[U, T] {
    return Pair[U, T]{p.Second, p.First}
}

// Generic interface
type Container[T any] interface {
    Add(T)
    Get(int) T
    Len() int
}

// Generic stack implementation
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Add(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
    if len(s.items) == 0 {
        var zero T
        return zero
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

func (s *Stack[T]) Get(i int) T {
    return s.items[i]
}

func (s *Stack[T]) Len() int {
    return len(s.items)
}

func main() {
    // Generic functions
    Print(42)
    Print("Hello")
    
    // Map function
    numbers := []int{1, 2, 3, 4, 5}
    squared := Map(numbers, func(n int) int { return n * n })
    fmt.Println(squared)
    
    // Max function
    fmt.Println(Max(10, 20))
    fmt.Println(Max(3.14, 2.71))
    
    // Generic structs
    p := Pair[string, int]{"Age", 30}
    fmt.Println(p)
    
    // Generic stack
    stack := &Stack[string]{}
    stack.Add("first")
    stack.Add("second")
    stack.Add("third")
    
    for stack.Len() > 0 {
        fmt.Println(stack.Pop())
    }
}
```

### Mini-Project: Generic Data Structure Library
Create reusable generic collections:

```go
package main

import "fmt"

// Queue
type Queue[T any] struct {
    items []T
}

func (q *Queue[T]) Enqueue(item T) {
    q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
    var zero T
    if len(q.items) == 0 {
        return zero, false
    }
    item := q.items[0]
    q.items = q.items[1:]
    return item, true
}

// Set using map
type Set[T comparable] struct {
    items map[T]bool
}

func NewSet[T comparable]() *Set[T] {
    return &Set[T]{items: make(map[T]bool)}
}

func (s *Set[T]) Add(item T) {
    s.items[item] = true
}

func (s *Set[T]) Contains(item T) bool {
    return s.items[item]
}

func (s *Set[T]) Size() int {
    return len(s.items)
}

// Filter function
func Filter[T any](items []T, predicate func(T) bool) []T {
    result := make([]T, 0)
    for _, item := range items {
        if predicate(item) {
            result = append(result, item)
        }
    }
    return result
}

func main() {
    // Queue example
    queue := &Queue[int]{}
    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)
    
    for {
        if val, ok := queue.Dequeue(); ok {
            fmt.Println("Dequeued:", val)
        } else {
            break
        }
    }
    
    // Set example
    set := NewSet[string]()
    set.Add("apple")
    set.Add("banana")
    set.Add("apple") // Duplicate
    
    fmt.Println("Set size:", set.Size())
    
    // Filter example
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
    fmt.Println("Even numbers:", evens)
}
```

---

## COMPREHENSIVE PROJECT: Build a REST API Server

Let's combine everything into a complete project:

```go
// main.go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "sync"
)

type Article struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Body  string `json:"body"`
}

type Server struct {
    articles map[int]Article
    mu       sync.RWMutex
    nextID   int
}

func NewServer() *Server {
    return &Server{
        articles: make(map[int]Article),
        nextID:   1,
    }
}

// GET /articles
func (s *Server) listArticles(w http.ResponseWriter, r *http.Request) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    articles := make([]Article, 0, len(s.articles))
    for _, article := range s.articles {
        articles = append(articles, article)
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(articles)
}

// POST /articles
func (s *Server) createArticle(w http.ResponseWriter, r *http.Request) {
    var article Article
    if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    s.mu.Lock()
    article.ID = s.nextID
    s.articles[article.ID] = article
    s.nextID++
    s.mu.Unlock()
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(article)
}

// Middleware
func logging(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
        handler(w, r)
    }
}

func (s *Server) setupRoutes() {
    http.HandleFunc("/articles", logging(s.listArticles))
    http.HandleFunc("/articles", logging(s.createArticle))
}

func main() {
    server := NewServer()
    server.setupRoutes()
    
    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

## LEARNING ROADMAP

1. **Week 1**: Variables, Functions, Structs
2. **Week 2**: Interfaces, Errors, Control Flow
3. **Week 3**: Slices, Maps, Pointers
4. **Week 4**: Advanced Functions, Local Development
5. **Week 5**: Goroutines, Channels, Concurrency
6. **Week 6**: Mutex, Advanced Concurrency
7. **Week 7**: Generics
8. **Week 8**: Build real projects combining all concepts

---

## RESOURCES

- Official Go docs: https://golang.org/doc
- Effective Go: https://golang.org/doc/effective_go
- Go by Example: https://gobyexample.com
- Go Weekly: https://golangweekly.com
