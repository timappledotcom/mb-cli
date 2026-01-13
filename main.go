package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"mb/internal/config"
	"mb/internal/poster"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func main() {
	logger := log.New(os.Stderr)
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Could not load config", "error", err)
	}

	// Check if setup is needed
	if cfg.MicroBlogToken == "" {
		setupConfig(cfg)
	}

	var content string
	// usage: mb "my text"
	if len(os.Args) > 1 {
		content = strings.Join(os.Args[1:], " ")
	} else {
		// usage: mb (opens editor)
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewText().
					Title("Valid Post?").
					Value(&content).
					Placeholder("What's on your mind?").
					CharLimit(280).
					Lines(5),
			),
		)
		err := form.Run()
		if err != nil {
			os.Exit(1)
		}
	}

	if strings.TrimSpace(content) == "" {
		fmt.Println("Empty post, cancelling.")
		return
	}

	// Double check or just send? "Quick to use" -> just send.
	// But let's give a nice spinner.

	var wg sync.WaitGroup
	
	// Create channels to capture results
	type result struct {
		service string
		err     error
	}
	results := make(chan result, 1)

	spinner := func(name string) {
		// Just a placeholder for a real spinner if we used bubbletea directly, 
		// but since we are just running functions, we'll print status.
		// For a "Pretty" CLI, let's just print a nice starting message.
	}
	_ = spinner

	style := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	fmt.Println(style.Render("Posting..."))

	if cfg.MicroBlogToken != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := poster.PostToMicroBlog(cfg, content)
			results <- result{"Micro.blog", err}
		}()
	}

	wg.Wait()
	close(results)

	hasError := false
	for res := range results {
		if res.err != nil {
			hasError = true
			fmt.Printf("%s %s: %v\n", lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("❌"), res.service, res.err)
		} else {
			fmt.Printf("%s %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Render("✅"), res.service)
		}
	}

	if !hasError {
		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("86")).Italic(true).Render("Done!"))
	} else {
		os.Exit(1)
	}
}

func setupConfig(cfg *config.Config) {
	fmt.Println("Welcome! Let's set up your keys.")
	
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Micro.blog Token").
				Value(&cfg.MicroBlogToken).
				Password(true),
		),
	)

	err := form.Run()

	if err != nil {
		log.Fatal("Setup cancelled")
	}

	if err := cfg.Save(); err != nil {
		log.Fatal("Could not save config", "error", err)
	}
	fmt.Println("Configuration saved!")
}
