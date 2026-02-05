# Golang Mini-Projects & Exercises

## Project 1: Personal Finance Tracker

Learn: Variables, Structs, Slices, Functions

```go
package main

import (
    "fmt"
    "time"
)

type Transaction struct {
    ID        int
    Amount    float64
    Category  string
    Date      time.Time
    Type      string // "income" or "expense"
}

type Account struct {
    Name         string
    Balance      float64
    Transactions []Transaction
    nextID       int
}

func NewAccount(name string, initialBalance float64) *Account {
    return &Account{
        Name:         name,
        Balance:      initialBalance,
        Transactions: make([]Transaction, 0),
        nextID:       1,
    }
}

func (a *Account) AddIncome(amount float64, category string) {
    if amount <= 0 {
        fmt.Println("Amount must be positive")
        return
    }
    
    a.Balance += amount
    a.Transactions = append(a.Transactions, Transaction{
        ID:       a.nextID,
        Amount:   amount,
        Category: category,
        Date:     time.Now(),
        Type:     "income",
    })
    a.nextID++
}

func (a *Account) AddExpense(amount float64, category string) error {
    if amount <= 0 {
        return fmt.Errorf("amount must be positive")
    }
    if amount > a.Balance {
        return fmt.Errorf("insufficient balance")
    }
    
    a.Balance -= amount
    a.Transactions = append(a.Transactions, Transaction{
        ID:       a.nextID,
        Amount:   amount,
        Category: category,
        Date:     time.Now(),
        Type:     "expense",
    })
    a.nextID++
    return nil
}

func (a *Account) GetBalance() float64 {
    return a.Balance
}

func (a *Account) GetExpensesByCategory(category string) float64 {
    total := 0.0
    for _, t := range a.Transactions {
        if t.Type == "expense" && t.Category == category {
            total += t.Amount
        }
    }
    return total
}

func (a *Account) PrintStatement() {
    fmt.Printf("=== Account: %s ===\n", a.Name)
    fmt.Printf("Balance: $%.2f\n\n", a.Balance)
    fmt.Println("Transactions:")
    for _, t := range a.Transactions {
        fmt.Printf("[%d] %s - %s: $%.2f (%s)\n", 
            t.ID, t.Date.Format("2006-01-02"), t.Category, t.Amount, t.Type)
    }
}

func main() {
    account := NewAccount("Savings", 5000)
    
    account.AddIncome(2000, "salary")
    account.AddExpense(500, "food")
    account.AddExpense(200, "entertainment")
    account.AddExpense(100, "food")
    
    account.PrintStatement()
    
    foodSpend := account.GetExpensesByCategory("food")
    fmt.Printf("\nTotal food expenses: $%.2f\n", foodSpend)
}
```

---

## Project 2: Task Management System

Learn: Interfaces, Methods, Slices, Error Handling

```go
package main

import (
    "fmt"
    "time"
)

type Priority int

const (
    Low Priority = iota
    Medium
    High
)

type Status int

const (
    Todo Status = iota
    InProgress
    Done
)

type Task struct {
    ID       int
    Title    string
    Priority Priority
    Status   Status
    DueDate  time.Time
    Created  time.Time
}

type TaskManager struct {
    tasks []Task
    count int
}

func NewTaskManager() *TaskManager {
    return &TaskManager{
        tasks: make([]Task, 0),
        count: 0,
    }
}

func (tm *TaskManager) AddTask(title string, priority Priority, dueDate time.Time) int {
    task := Task{
        ID:       tm.count + 1,
        Title:    title,
        Priority: priority,
        Status:   Todo,
        DueDate:  dueDate,
        Created:  time.Now(),
    }
    tm.tasks = append(tm.tasks, task)
    tm.count++
    return task.ID
}

func (tm *TaskManager) UpdateStatus(id int, status Status) error {
    for i := range tm.tasks {
        if tm.tasks[i].ID == id {
            tm.tasks[i].Status = status
            return nil
        }
    }
    return fmt.Errorf("task %d not found", id)
}

func (tm *TaskManager) GetTasksByStatus(status Status) []Task {
    result := make([]Task, 0)
    for _, task := range tm.tasks {
        if task.Status == status {
            result = append(result, task)
        }
    }
    return result
}

func (tm *TaskManager) GetTasksByPriority(priority Priority) []Task {
    result := make([]Task, 0)
    for _, task := range tm.tasks {
        if task.Priority == priority {
            result = append(result, task)
        }
    }
    return result
}

func (t Task) String() string {
    statusStr := []string{"TODO", "IN PROGRESS", "DONE"}[t.Status]
    priorityStr := []string{"LOW", "MEDIUM", "HIGH"}[t.Priority]
    return fmt.Sprintf("[%d] %s | Priority: %s | Status: %s | Due: %s",
        t.ID, t.Title, priorityStr, statusStr, t.DueDate.Format("2006-01-02"))
}

func (tm *TaskManager) PrintAll() {
    fmt.Println("=== All Tasks ===")
    for _, task := range tm.tasks {
        fmt.Println(task)
    }
}

func main() {
    tm := NewTaskManager()
    
    tm.AddTask("Learn Go", High, time.Now().AddDate(0, 0, 7))
    tm.AddTask("Build REST API", High, time.Now().AddDate(0, 0, 14))
    tm.AddTask("Read documentation", Low, time.Now().AddDate(0, 0, 30))
    
    tm.UpdateStatus(1, InProgress)
    tm.UpdateStatus(2, InProgress)
    
    tm.PrintAll()
    
    fmt.Println("\n=== High Priority Tasks ===")
    for _, task := range tm.GetTasksByPriority(High) {
        fmt.Println(task)
    }
    
    fmt.Println("\n=== In Progress Tasks ===")
    for _, task := range tm.GetTasksByStatus(InProgress) {
        fmt.Println(task)
    }
}
```

---

## Project 3: Concurrent Web Scraper

Learn: Goroutines, Channels, Concurrency, Error Handling

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Page struct {
    URL   string
    Title string
    Error error
}

func fetchPage(url string) *Page {
    // Simulating HTTP request
    time.Sleep(time.Duration(100) * time.Millisecond)
    
    if url == "http://error.com" {
        return &Page{
            URL:   url,
            Error: fmt.Errorf("failed to fetch %s", url),
        }
    }
    
    return &Page{
        URL:   url,
        Title: fmt.Sprintf("Page from %s", url),
    }
}

func ScrapeConcurrent(urls []string, maxConcurrent int) []*Page {
    semaphore := make(chan struct{}, maxConcurrent)
    results := make(chan *Page, len(urls))
    var wg sync.WaitGroup
    
    for _, url := range urls {
        wg.Add(1)
        go func(u string) {
            defer wg.Done()
            
            // Acquire semaphore
            semaphore <- struct{}{}
            defer func() { <-semaphore }()
            
            page := fetchPage(u)
            results <- page
        }(url)
    }
    
    wg.Wait()
    close(results)
    
    pages := make([]*Page, 0, len(urls))
    for page := range results {
        pages = append(pages, page)
    }
    return pages
}

func main() {
    urls := []string{
        "http://example.com/1",
        "http://example.com/2",
        "http://error.com",
        "http://example.com/3",
        "http://example.com/4",
    }
    
    start := time.Now()
    pages := ScrapeConcurrent(urls, 2)
    elapsed := time.Since(start)
    
    fmt.Printf("Scraped %d pages in %v\n\n", len(pages), elapsed)
    
    for _, page := range pages {
        if page.Error != nil {
            fmt.Printf("❌ %s: %v\n", page.URL, page.Error)
        } else {
            fmt.Printf("✓ %s: %s\n", page.URL, page.Title)
        }
    }
}
```

---

## Project 4: Configuration Management System

Learn: Structs, JSON, Error Handling, Generics

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Database struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    User     string `json:"user"`
    Password string `json:"password"`
}

type Server struct {
    Host string `json:"host"`
    Port int    `json:"port"`
}

type Config struct {
    AppName  string   `json:"app_name"`
    Version  string   `json:"version"`
    Database Database `json:"database"`
    Server   Server   `json:"server"`
}

func LoadConfig(filename string) (*Config, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }
    
    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("failed to parse config: %w", err)
    }
    
    return &config, nil
}

func (c *Config) SaveConfig(filename string) error {
    data, err := json.MarshalIndent(c, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal config: %w", err)
    }
    
    if err := os.WriteFile(filename, data, 0644); err != nil {
        return fmt.Errorf("failed to write config file: %w", err)
    }
    
    return nil
}

func (c *Config) GetDatabaseURL() string {
    return fmt.Sprintf("postgres://%s:%s@%s:%d/mydb",
        c.Database.User,
        c.Database.Password,
        c.Database.Host,
        c.Database.Port,
    )
}

func main() {
    // Create sample config
    config := &Config{
        AppName: "MyApp",
        Version: "1.0.0",
        Database: Database{
            Host:     "localhost",
            Port:     5432,
            User:     "admin",
            Password: "secret",
        },
        Server: Server{
            Host: "0.0.0.0",
            Port: 8080,
        },
    }
    
    // Save config
    config.SaveConfig("config.json")
    fmt.Println("Config saved")
    
    // Load config
    loaded, err := LoadConfig("config.json")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("App: %s v%s\n", loaded.AppName, loaded.Version)
    fmt.Printf("Database URL: %s\n", loaded.GetDatabaseURL())
    fmt.Printf("Server: %s:%d\n", loaded.Server.Host, loaded.Server.Port)
}
```

---

## Project 5: Expression Calculator

Learn: Recursion, Interfaces, Functions, Error Handling

```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

type Expression interface {
    Evaluate() (float64, error)
    String() string
}

type NumberExpr struct {
    Value float64
}

func (n NumberExpr) Evaluate() (float64, error) {
    return n.Value, nil
}

func (n NumberExpr) String() string {
    return fmt.Sprintf("%.2f", n.Value)
}

type BinaryOp struct {
    Left  Expression
    Right Expression
    Op    string
}

func (b BinaryOp) Evaluate() (float64, error) {
    left, err := b.Left.Evaluate()
    if err != nil {
        return 0, err
    }
    
    right, err := b.Right.Evaluate()
    if err != nil {
        return 0, err
    }
    
    switch b.Op {
    case "+":
        return left + right, nil
    case "-":
        return left - right, nil
    case "*":
        return left * right, nil
    case "/":
        if right == 0 {
            return 0, fmt.Errorf("division by zero")
        }
        return left / right, nil
    default:
        return 0, fmt.Errorf("unknown operator: %s", b.Op)
    }
}

func (b BinaryOp) String() string {
    return fmt.Sprintf("(%s %s %s)", b.Left, b.Op, b.Right)
}

// Simple parser
func ParseExpression(expr string) (Expression, error) {
    tokens := strings.Fields(expr)
    
    if len(tokens) == 0 {
        return nil, fmt.Errorf("empty expression")
    }
    
    if len(tokens) == 1 {
        val, err := strconv.ParseFloat(tokens[0], 64)
        if err != nil {
            return nil, fmt.Errorf("invalid number: %s", tokens[0])
        }
        return NumberExpr{val}, nil
    }
    
    if len(tokens) == 3 {
        left, err := ParseExpression(tokens[0])
        if err != nil {
            return nil, err
        }
        
        right, err := ParseExpression(tokens[2])
        if err != nil {
            return nil, err
        }
        
        return BinaryOp{left, right, tokens[1]}, nil
    }
    
    return nil, fmt.Errorf("invalid expression")
}

func main() {
    tests := []string{
        "10 + 5",
        "20 - 8",
        "4 * 7",
        "15 / 3",
    }
    
    for _, test := range tests {
        expr, err := ParseExpression(test)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            continue
        }
        
        result, err := expr.Evaluate()
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            continue
        }
        
        fmt.Printf("%s = %.2f\n", expr, result)
    }
}
```

---

## Project 6: URL Shortener with Goroutines

Learn: Maps, Goroutines, Channels, Concurrency Control

```go
package main

import (
    "crypto/md5"
    "fmt"
    "math/rand"
    "sync"
)

type URLShortener struct {
    mu           sync.RWMutex
    urlMap       map[string]string // short -> long
    reverseMap   map[string]string // long -> short
    stats        map[string]int    // short -> hits
}

func NewURLShortener() *URLShortener {
    return &URLShortener{
        urlMap:     make(map[string]string),
        reverseMap: make(map[string]string),
        stats:      make(map[string]int),
    }
}

func generateShortCode(longURL string) string {
    hash := md5.Sum([]byte(longURL))
    return fmt.Sprintf("%x", hash)[:6]
}

func (us *URLShortener) Shorten(longURL string) string {
    us.mu.Lock()
    defer us.mu.Unlock()
    
    // Check if already shortened
    if short, exists := us.reverseMap[longURL]; exists {
        return short
    }
    
    short := generateShortCode(longURL)
    us.urlMap[short] = longURL
    us.reverseMap[longURL] = short
    us.stats[short] = 0
    
    return short
}

func (us *URLShortener) Expand(shortURL string) (string, error) {
    us.mu.RLock()
    defer us.mu.RUnlock()
    
    longURL, exists := us.urlMap[shortURL]
    if !exists {
        return "", fmt.Errorf("short URL not found")
    }
    
    return longURL, nil
}

func (us *URLShortener) Click(shortURL string) error {
    us.mu.Lock()
    defer us.mu.Unlock()
    
    if _, exists := us.urlMap[shortURL]; !exists {
        return fmt.Errorf("short URL not found")
    }
    
    us.stats[shortURL]++
    return nil
}

func (us *URLShortener) Stats() map[string]int {
    us.mu.RLock()
    defer us.mu.RUnlock()
    
    // Create copy to avoid race conditions
    statsCopy := make(map[string]int)
    for k, v := range us.stats {
        statsCopy[k] = v
    }
    return statsCopy
}

func main() {
    shortener := NewURLShortener()
    
    // Shorten URLs
    url1 := "https://www.example.com/very/long/url/path"
    url2 := "https://github.com/golang/go"
    
    short1 := shortener.Shorten(url1)
    short2 := shortener.Shorten(url2)
    
    fmt.Printf("Shortened: %s -> %s\n", url1, short1)
    fmt.Printf("Shortened: %s -> %s\n", url2, short2)
    
    // Simulate concurrent clicks
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(code string) {
            defer wg.Done()
            shortener.Click(code)
        }(short1)
    }
    
    for i := 0; i < 50; i++ {
        wg.Add(1)
        go func(code string) {
            defer wg.Done()
            shortener.Click(code)
        }(short2)
    }
    
    wg.Wait()
    
    // Expand URLs
    expanded, _ := shortener.Expand(short1)
    fmt.Printf("\nExpanded: %s -> %s\n", short1, expanded)
    
    // Show stats
    fmt.Println("\nClick Stats:")
    for short, clicks := range shortener.Stats() {
        fmt.Printf("%s: %d clicks\n", short, clicks)
    }
}
```

---

## Project 7: Generic Logger with Different Outputs

Learn: Generics, Interfaces, Variadic Functions

```go
package main

import (
    "fmt"
    "os"
    "time"
)

type LogLevel int

const (
    DEBUG LogLevel = iota
    INFO
    WARNING
    ERROR
)

type Logger[T any] interface {
    Log(level LogLevel, message string, data T)
}

type ConsoleLogger[T any] struct {
    minLevel LogLevel
}

func (cl *ConsoleLogger[T]) Log(level LogLevel, message string, data T) {
    levelStr := []string{"DEBUG", "INFO", "WARNING", "ERROR"}[level]
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    fmt.Printf("[%s] %s: %s (data: %v)\n", timestamp, levelStr, message, data)
}

type FileLogger[T any] struct {
    filename string
    minLevel LogLevel
}

func (fl *FileLogger[T]) Log(level LogLevel, message string, data T) {
    levelStr := []string{"DEBUG", "INFO", "WARNING", "ERROR"}[level]
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    line := fmt.Sprintf("[%s] %s: %s (data: %v)\n", timestamp, levelStr, message, data)
    
    file, _ := os.OpenFile(fl.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    defer file.Close()
    file.WriteString(line)
}

func main() {
    // String logger
    consoleLogger := &ConsoleLogger[string]{INFO}
    consoleLogger.Log(INFO, "Application started", "version 1.0")
    consoleLogger.Log(WARNING, "High memory usage", "85%")
    
    // Int logger
    intLogger := &ConsoleLogger[int]{DEBUG}
    intLogger.Log(DEBUG, "Request count", 42)
    intLogger.Log(ERROR, "Failed requests", 3)
}
```

These projects progressively build your skills from basic to advanced Go concepts. After completing these, you'll be well-equipped to build production applications!
