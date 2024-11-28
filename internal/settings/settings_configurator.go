package settings

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pafello/gocast/internal/geolocation"
	"github.com/pafello/gocast/internal/styles"
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
		steps:            map[int]string{0: "What is your location", 1: "Select correct location", 2: "Select unit system", 3: "", 4: "Not now? Sure, come back when you are ready :)"},
		locationChoices:  []geolocation.GeolocationResult{},
		selectedLocation: geolocation.GeolocationResult{},
		unitChoices:      []units.UnitSystem{units.Metric, units.Imperial},
		selectedUnit:     units.Metric,
		cursor:           0,
	}
}
func (m interviewModel) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m interviewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:

		if m.currentStep == 0 {
			switch msg.String() {
			case "esc", "ctrl+c":
				m.currentStep = 4
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
			case "esc", "ctrl+c":

				m.currentStep = 4
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
				m.currentStep = 4
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
				m.currentStep++
				cmd = tea.Quit
			}
			return m, cmd
		}
	}

	return m, cmd
}

func (m interviewModel) View() string {
	instructions := "(esc) quit, (enter) select"
	additionalInstruction := "(↑/↓) navigate"
	s := ""

	if m.currentStep == 0 {
		s += styles.HeaderBox.Render("Welcome to Go Cast!") + "\n"
	}
	if len(m.steps[m.currentStep]) > 0 {
		s += styles.Text.Render(m.steps[m.currentStep]) + "\n"
	}

	if m.currentStep == 0 {
		s += styles.TextInput.Render(m.textInput.View()) + "\n"
	}
	if m.currentStep == 1 {
		for i, c := range m.locationChoices {
			cursor := " "
			t := ""
			if m.cursor == i {
				cursor = "•"
				t = styles.SelectedOption.Render(c.Describe())
			} else {
				t = c.Describe()
			}
			s += fmt.Sprintf("%s %s\n", cursor, t)
		}
	}
	if m.currentStep == 2 {
		for i, u := range m.unitChoices {
			t := ""
			cursor := " "
			if m.cursor == i {
				cursor = "•"
				t = styles.SelectedOption.Render(string(u))
			} else {
				t = string(u)
			}
			s += fmt.Sprintf("%s %s\n", cursor, t)
		}

	}
	if m.currentStep != 3 && m.currentStep != 4 {
		s += styles.Subtitle.Render(instructions)
	}
	if m.currentStep == 1 || m.currentStep == 2 {
		s += styles.Subtitle.Render(additionalInstruction)
	}

	return s
}

type Status int

const (
	Success Status = iota
	Canceled
)

func InterviewUser() (UserSettings, Status, error) {

	p := tea.NewProgram(initialModel())
	data, err := p.Run()

	model := data.(interviewModel)

	if err != nil {
		return UserSettings{}, Canceled, err
	}

	if model.currentStep != 3 {
		return UserSettings{}, Canceled, nil
	}

	userSettings := UserSettings{
		Location: model.selectedLocation,
		UnitSys:  model.selectedUnit,
	}

	return userSettings, Success, nil

}
