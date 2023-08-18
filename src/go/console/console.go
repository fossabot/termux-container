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

package console

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err       error
}

func initialModel() model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.Prompt = "> "
	ti.Cursor.Blink = false
	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	view := strings.Fields(m.textInput.View())
	value := strings.Fields(m.textInput.Value())
	var ret string
	var list []string = []string{"help", "exit", "new", "search", "pull", "rmi", "cp", "ls", "rm", "login", "import", "export", "info"}
	if len(value) == 0 {
		ret = m.textInput.View()
	} else {
		ret += "\033[0m" + view[0] + " "
		for i := 0; i < len(value); i++ {
			highlight := false
			for j := 0; j < len(list); j++ {
				if value[i] == list[j] {
					ret += "\033[32m" + strings.ReplaceAll(view[i+1], "\033[0m", "\033[0m\033[32m") + " " + "\033[0m"
					highlight = true
					break
				}
			}
			if highlight == false {
				ret += view[i+1] + " " + "\033[0m"
			}
		}
		if m.textInput.Value()[len(m.textInput.Value())-1] == ' ' {
			ret += "\033[7m \033[0m"
		}
	}
	return ret
}
func Console() string {
	print("\n")
	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return m.(model).textInput.Value()
}
