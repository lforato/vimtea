package vimtea

import tea "github.com/charmbracelet/bubbletea"

// QuitMsg is emitted by QuitCmd to signal the embedding parent model that the
// user is done editing. The parent is responsible for consuming this message
// (e.g. closing the editor modal) and reading Content.
type QuitMsg struct {
	Content string
}

// QuitCmd matches the KeyBinding.Handler signature and returns a tea.Cmd that
// dispatches a QuitMsg carrying the current buffer text. Use it when wiring a
// quit key (typically esc in ModeNormal) so parent models can react.
//
//	editor.AddBinding(vimtea.KeyBinding{
//		Key:     "esc",
//		Mode:    vimtea.ModeNormal,
//		Handler: vimtea.QuitCmd,
//	})
func QuitCmd(buf Buffer) tea.Cmd {
	content := buf.Text()
	return func() tea.Msg {
		return QuitMsg{Content: content}
	}
}
