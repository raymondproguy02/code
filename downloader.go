package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "sync"
)

func downloadFile(url string, wg *sync.WaitGroup) {
    defer wg.Done()
    
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("❌ Error: %s\n", url)
        return
    }
    defer resp.Body.Close()

    filename := fmt.Sprintf("download_%d", wg)
    file, err := os.Create(filename)
    if err != nil {
        fmt.Printf("❌ File error: %s\n", url)
        return
    }
    defer file.Close()

    _, err = io.Copy(file, resp.Body)
    if err != nil {
        fmt.Printf("❌ Copy error: %s\n", url)
        return
    }

    fmt.Printf("✅ Downloaded: %s\n", url)
}

func main() {
    urls := []string{
        "https://example.com/file1.jpg",
        "https://example.com/file2.pdf",
        "https://example.com/file3.zip",
    }

    var wg sync.WaitGroup
    for _, url := range urls {
        wg.Add(1)
        go downloadFile(url, &wg)
    }
    wg.Wait()
    fmt.Println("🎉 All downloads complete")
}
