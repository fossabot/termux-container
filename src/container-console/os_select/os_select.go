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

// Edit this list to use it in your own program.
var choices = []string{"[\033[0m\033[1;38;2;254;228;208m] Almalinux", "[\033[0m\033[1;38;2;254;228;208m] Alpine", "[\033[0m\033[1;38;2;254;228;208m] Amazonlinux", "[\033[0m\033[1;38;2;254;228;208m] Archlinux", "[\033[0m\033[1;38;2;254;228;208m] Busybox", "[\033[0m\033[1;38;2;254;228;208m] CentOS", "[\033[0m\033[1;38;2;254;228;208m] Debian", "[\033[0m\033[1;38;2;254;228;208m] Devuan", "[\033[0m\033[1;38;2;254;228;208m] Fedora", "[\033[0m\033[1;38;2;254;228;208m] Gentoo", "[\033[0m\033[1;38;2;254;228;208m] Kali", "[\033[0m\033[1;38;2;254;228;208m] Manjaro", "[\033[0m\033[1;38;2;254;228;208m] OpenSUSE", "[\033[0m\033[1;38;2;254;228;208m] ParrotOS", "[\033[0m\033[1;38;2;254;228;208m] Rockylinux", "[\033[0m\033[1;38;2;254;228;208m] Ubuntu", "[\033[0m\033[1;38;2;254;228;208m] Voidlinux"}
var chosen = []string{"almalinux", "alpine", "amazonlinux", "archlinux", "busybox", "centos", "debian", "devuan", "fedora", "gentoo", "kali", "manjaro", "opensuse", "parrotos", "rockylinux", "ubuntu", "voidlinux"}

type model struct {
	cursor int
	choice string
	chosen string
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
			m.chosen = chosen[m.cursor]
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
	// The value `3` means that 3+1+3 lines will be shown in terminal,
	// the cursor will always be at the center of the them.
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

// You might also need to rename this function to use it in your own program.
func Os_select() string {
	p := tea.NewProgram(model{})
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
	}
	return m.(model).chosen
}
