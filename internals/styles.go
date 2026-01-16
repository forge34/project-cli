package internals

import "github.com/charmbracelet/lipgloss"

var Header = lipgloss.NewStyle().
	Bold(true)

var Success = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#43A047")).
	Bold(true)

var InfoBox = lipgloss.NewStyle().
	Padding(0, 2).
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#7B61FF"))
