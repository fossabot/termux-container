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
package cap_select

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	sel := make(map[int]struct{})
	var on [22]int = [22]int{2, 9, 11, 12, 14, 15, 16, 17, 19, 20, 21, 22, 23, 24, 25, 26, 28, 30, 32, 33, 34, 35}
	for i := 0; i < len(on); i++ {
		sel[on[i]] = struct{}{}
	}
	return model{
		choices:  []string{"cap_chown", "cap_dac_override", "cap_dac_read_search", "cap_fowner", "cap_fsetid", "cap_kill", "cap_setgid", "cap_setuid", "cap_setpcap", "cap_linux_immutable", "cap_net_bind_service", "cap_net_broadcast", "cap_net_admin", "cap_net_raw", "cap_ipc_lock", "cap_ipc_owner", "cap_sys_module", "cap_sys_rawio", "cap_sys_chroot", "cap_sys_ptrace", "cap_sys_pacct", "cap_sys_admin", "cap_sys_boot", "cap_sys_nice", "cap_sys_resource", "cap_sys_time", "cap_sys_tty_config", "cap_mknod", "cap_lease", "cap_audit_write", "cap_audit_control", "cap_setfcap", "cap_mac_override", "cap_mac_admin", "cap_syslog", "cap_wake_alarm", "cap_block_suspend", "cap_audit_read", "cap_perfmon", "cap_bpf", "cap_checkpoint_restore"},
		selected: sel,
	}
}
func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "enter":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}
func (m model) View() string {
	s := "\033[1;38;2;254;228;208mChoose the capabilities to be dropped\n\n"
	if m.cursor > 3 && m.cursor < 37 {
		for i, choice := range m.choices {
			if i-3 <= m.cursor && i+3 >= m.cursor {
				cursor := "  "
				if m.cursor == i {
					cursor = "->"
				}
				checked := " "
				if _, ok := m.selected[i]; ok {
					checked = "\033[1;38;2;152;245;225mx"
				}
				s += fmt.Sprintf("\033[1;38;2;254;228;208m%s [%s\033[1;38;2;254;228;208m] %s\n", cursor, checked, choice)
			}
		}
	} else if m.cursor >= 37 {
		for i := m.cursor - 3; i < m.cursor+3; i++ {
			if i <= 40 {
				cursor := "  "
				if m.cursor == i {
					cursor = "->"
				}
				checked := " "
				if _, ok := m.selected[i]; ok {
					checked = "\033[1;38;2;152;245;225mx"
				}
				s += fmt.Sprintf("\033[1;38;2;254;228;208m%s [%s\033[1;38;2;254;228;208m] %s\n", cursor, checked, m.choices[i])
			} else {
				s += "\n"
			}
		}
		s += "\n"
	} else {
		for i := 0; i < 3-m.cursor; i++ {
			s += "\n"
		}
		for i := 0; i < m.cursor+4; i++ {
			cursor := "  "
			if m.cursor == i {
				cursor = "->"
			}
			checked := " "
			if _, ok := m.selected[i]; ok {
				checked = "\033[1;38;2;152;245;225mx"
			}
			s += fmt.Sprintf("\033[1;38;2;254;228;208m%s [%s\033[1;38;2;254;228;208m] %s\n", cursor, checked, m.choices[i])
		}
	}
	s += "\nPress Up/Down to move, Space to select, Enter to apply\n"
	return s
}
func Cap_select() int {
	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	for i := 0; i < len(m.(model).choices); i++ {
		if _, ok := m.(model).selected[i]; ok {
			fmt.Print(i)
		}
	}
	fmt.Print(m.(model).selected)
	return 0
}
