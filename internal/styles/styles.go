package styles

import "github.com/charmbracelet/lipgloss"

var HeaderText = lipgloss.NewStyle().
	Bold(true).
	Italic(true).
	Foreground(lipgloss.AdaptiveColor{Light: "#493d70", Dark: "#917bd1"}).Padding(1, 1, 1, 1)

var BoxBorder = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), true, true, true, true).
	BorderForeground(lipgloss.AdaptiveColor{Light: "#493d70", Dark: "#917bd1"}).
	Padding(1, 5, 1, 5).Width(35).Align(lipgloss.Center)

var HeaderBox = BoxBorder.Inherit(HeaderText)

var Text = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#3d3a45", Dark: "#e5defa"}).Padding(1, 1, 1, 1)

var Subtitle = lipgloss.NewStyle().Bold(false).Foreground(lipgloss.Color("#787878")).Padding(1, 2, 1, 2)

var TextInput = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), false, false, true, false).
	BorderForeground(lipgloss.AdaptiveColor{Light: "#493d70", Dark: "#917bd1"}).
	Width(35).Foreground(lipgloss.AdaptiveColor{Light: "#583e75", Dark: "#bf8df7"})

var SelectedOption = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{Light: "#583e75", Dark: "#bf8df7"})

var TableBorder = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#493d70", Dark: "#917bd1"})

var TableHeader = lipgloss.NewStyle().Bold(true).Align(lipgloss.Center)
var TableRowEven = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#3d3a45", Dark: "#e5defa"}).Align(lipgloss.Center)
var TableRowOdd = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#656073", Dark: "#9e99ad"}).Align(lipgloss.Center)
