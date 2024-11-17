package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Msg Types
type statusMsg int
type errMsg struct{ err error }

// Function to check a URL status
func checkUrl(url string) tea.Cmd {
	return func() tea.Msg {
		client := &http.Client{Timeout: 10 * time.Second}
		res, err := client.Get(url)
		if err != nil {
			return errMsg{err}
		}
		return statusMsg(res.StatusCode)
	}
}

// BubbleTea Model
type model struct {
	textInput textinput.Model
	status    int
	err       error
	loading   bool
}

// Initializing the model
func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter the URL"
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 30

	return model{
		textInput: ti,
		status:    0,
		err:       nil,
		loading:   false,
	}
}

// Init - Initialize Commands
func (m model) Init() tea.Cmd {
	return nil
}

// Update - Handle messages
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			url := m.textInput.Value()

			// Handle empty URL
			if url == "" {
				m.err = fmt.Errorf("URL can't be empty")
				m.textInput.Focus()
				return m, nil
			}

			// Start loading and check server
			m.loading = true
			m.err = nil
			return m, checkUrl(url)

		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}

	case statusMsg:
		// Display status and reset input
		m.status = int(msg)
		m.loading = false
		m.textInput.Reset()
		m.textInput.Focus()
		return m, nil

	case errMsg:
		// Display error and reset input
		m.err = msg.err
		m.loading = false
		m.textInput.Reset()
		m.textInput.Focus()
		return m, nil
	}

	// Update the text input field
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

// View - Render the UI
func (m model) View() string {
	if m.loading {
		return "Checking URL...\n"
	}

	if m.err != nil {
		return fmt.Sprintf("Error: %v\n\n%s", m.err, m.textInput.View())
	}

	if m.status > 0 {
		return fmt.Sprintf("Status: %d %s\n\n%s", m.status, http.StatusText(m.status), m.textInput.View())
	}

	return fmt.Sprintf(
		"Enter a URL to check its status:\n\n%s\n\nPress Enter to submit or Ctrl+C to quit.",
		m.textInput.View(),
	)
}

func main() {
	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		log.Fatal(err)
	}
}
