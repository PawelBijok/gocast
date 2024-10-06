package settings

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pafello/gocast/internal/geolocation"
	"github.com/pafello/gocast/internal/units"
)

type interviewModel struct {
	currentStep      int
	steps            map[int]string
	textInput        textinput.Model
	locationChoices  []geolocation.GeolocationResult
	selectedLocation geolocation.GeolocationResult
	unitChoices      []units.UnitSystem
	selectedUnit     units.UnitSystem
	cursor           int // which to-do list item our cursor is pointing at
}

func initialModel() interviewModel {
	li := textinput.New()
	li.Focus()
	li.Placeholder = "Warsaw"

	return interviewModel{
		currentStep:      0,
		textInput:        li,
		steps:            map[int]string{0: "What is your location", 1: "Select correct location", 2: "Select unit system"},
		locationChoices:  []geolocation.GeolocationResult{},
		selectedLocation: geolocation.GeolocationResult{},
		unitChoices:      []units.UnitSystem{units.Metric, units.Imperial},
		selectedUnit:     units.Metric,
		cursor:           0,
	}
}
func (m interviewModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m interviewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:

		if m.currentStep == 0 {
			switch msg.String() {
			case "esc", "ctrl+c":
				return m, tea.Quit

			case "enter":
				if m.textInput.Value() == "" {
					return m, nil
				}
				availableLocations, err := geolocation.GetGeolocationResults(m.textInput.Value())
				if err != nil {
					return m, tea.Quit
				}
				m.locationChoices = availableLocations
				m.currentStep++
			}
			m.textInput, cmd = m.textInput.Update(msg)

			return m, cmd

		}
		if m.currentStep == 1 {
			switch msg.String() {
			case "esc", "crtl+c":
				return m, tea.Quit

			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			case "down", "j":
				if m.cursor < len(m.locationChoices)-1 {
					m.cursor++
				}

			case "enter":
				m.selectedLocation = m.locationChoices[m.cursor]
				m.cursor = 0
				m.currentStep++
			}
			return m, cmd
		}

		if m.currentStep == 2 {
			switch msg.String() {
			case "esc", "ctrl+c":
				return m, tea.Quit

			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			case "down", "j":
				if m.cursor < len(m.unitChoices)-1 {
					m.cursor++
				}

			case "enter":
				m.selectedUnit = m.unitChoices[m.cursor]
				m.cursor = 0
				cmd = tea.Quit
			}
			return m, cmd
		}
	}

	return m, cmd
}

func (m interviewModel) View() string {
	// The header
	instructions := "(esc) quit, (enter) select"
	additionalInstruction := "(↑/↓) navigate"
	s := m.steps[m.currentStep] + "\n"

	if m.currentStep == 0 {
		s += m.textInput.View() + "\n"
	}
	if m.currentStep == 1 {
		for i, c := range m.locationChoices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, c.Describe())
		}
	}
	if m.currentStep == 2 {
		for i, u := range m.unitChoices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, u)
		}

	}

	s += instructions
	if m.currentStep != 0 {
		s += fmt.Sprintf(", %s", additionalInstruction)
	}

	return s
}

func InterviewUser() (UserSettings, error) {

	p := tea.NewProgram(initialModel())
	m, err := p.Run()

	model := m.(interviewModel)

	if err != nil {
		os.Exit(1)
	}

	userSettings := UserSettings{
		Location: model.selectedLocation,
		UnitSys:  model.selectedUnit,
	}

	return userSettings, nil

}
