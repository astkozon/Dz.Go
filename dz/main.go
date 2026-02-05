package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	
	file, err := os.Open("url.txt")
	if err != nil {
		fmt.Printf("Ошибка : %v\n", err)
		return
	}
	defer file.Close()

	
	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		if url != "" {
			urls = append(urls, url)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		return
	}

	// Проверка каждой url 
	for _, url := range urls {
		checkURL(url)
	}
}

func checkURL(url string) {
	startTime := time.Now()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	duration := time.Since(startTime)

	if err != nil {
		fmt.Printf("[ERROR] %s (%v) %v\n", url, duration.Round(time.Millisecond), err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("[%d] %s (%v)\n", resp.StatusCode, url, duration.Round(time.Millisecond))
}