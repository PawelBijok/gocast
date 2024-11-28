package styles

import "github.com/charmbracelet/lipgloss"

var HeaderText = lipgloss.NewStyle().
	Bold(true).
	Italic(true).
	Foreground(lipgloss.AdaptiveColor{Light: "#493d70", Dark: "#917bd1"})

var Border = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), true, true, true, true).
	BorderForeground(lipgloss.AdaptiveColor{Light: "#493d70", Dark: "#917bd1"}).
	Padding(1, 5, 1, 5)

var HeaderBox = Border.Inherit(HeaderText)

var Subtitle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#3d3a45", Dark: "#e5defa"}).Padding(1, 1, 1, 1)
