package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func loadEnv(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }

        parts := strings.SplitN(line, "=", 2)
        if len(parts) != 2 {
            continue
        }

        key := strings.TrimSpace(parts[0])
        value := strings.TrimSpace(parts[1])
        value = strings.Trim(value, "\"'")

        os.Setenv(key, value)
    }

    return scanner.Err()
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run envloader.go <envfile>")
        fmt.Println("Example: go run envloader.go .env")
        os.Exit(1)
    }

    err := loadEnv(os.Args[1])
    if err != nil {
        fmt.Printf("❌ Error: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("✅ Environment variables loaded:")
    for _, env := range os.Environ() {
        if strings.HasPrefix(env, "APP_") {
            fmt.Printf("  %s\n", env)
        }
    }
}
