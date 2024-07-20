package models

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type InputModel struct {
	textInput textinput.Model
}

type ValueSubmittedMsg struct {
	Value string
}

func NewInputModel(placeholder string) *InputModel {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	return &InputModel{
		textInput: ti,
	}
}

func (m *InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			value := m.textInput.Value()
			m.textInput.SetValue("")
			m.textInput.Blur()
			outFunc := func() tea.Msg {
				return ValueSubmittedMsg{Value: value}
			}
			return GetMasterModel().GoBack(outFunc)
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *InputModel) View() string { return m.textInput.View() }
