package os_select

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var choices = []string{"[] Almalinux", "[] Alpine", "[] Amazonlinux", "[] Archlinux", "[] Busybox", "[] CentOS", "[] Debian", "[] Devuan", "[] Fedora", "[] Gentoo", "[] Kali", "[] Manjaro", "[] OpenSUSE", "[] ParrotOS", "[] Rockylinux", "[] Ubuntu", "[] Voidlinux"}

type model struct {
	cursor int
	choice string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			m.choice = choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("\033[1;38;2;254;228;208mChoose the OS to install:\n\n")
	if m.cursor > 3 && m.cursor < len(choices)-3 {
		for i := m.cursor - 3; i < m.cursor+3; i++ {
			if m.cursor == i {
				s.WriteString("\033[1;38;2;152;245;225m->\033[1;38;2;254;228;208m ")
			} else {
				s.WriteString("   ")
			}
			s.WriteString(choices[i])
			s.WriteString("\n")
		}
	} else if m.cursor >= len(choices)-3 {
		for i := m.cursor - 3; i < m.cursor+3; i++ {
			if i < len(choices) {
				if m.cursor == i {
					s.WriteString("\033[1;38;2;152;245;225m->\033[1;38;2;254;228;208m ")
				} else {
					s.WriteString("   ")
				}
				s.WriteString(choices[i])
				s.WriteString("\n")
			} else {
				s.WriteString("\n")
			}
		}
	} else {
		for i := 0; i < 3-m.cursor; i++ {
			s.WriteString("\n")
		}
		for i := 0; i < m.cursor+3; i++ {
			if m.cursor == i {
				s.WriteString("\033[1;38;2;152;245;225m->\033[1;38;2;254;228;208m ")
			} else {
				s.WriteString("   ")
			}
			s.WriteString(choices[i])
			s.WriteString("\n")
		}
	}
	s.WriteString("\nPress Up/Down to move, Enter to select\n")
	return s.String()
}
func Os_select() string {
	p := tea.NewProgram(model{})
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
	}
	return m.(model).choice
}
