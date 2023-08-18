// SPDX-License-Identifier: MIT
/*
 *
 * This file is part of termux-container, with ABSOLUTELY NO WARRANTY.
 *
 * MIT License
 *
 * Copyright (c) 2022-2023 Moe-hacker
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 *
 *
 */
// It works, so do not touch anything.
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
