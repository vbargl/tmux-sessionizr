package tmux

import (
	"fmt"

	"barglvojtech.net/tmux-sessionizr/sdk/shell"
)

const sessionizrID = "53579"

func SetHook(hook Hook, command string) error {
	scopedHook := fmt.Sprintf("%s[%s]", hook, sessionizrID)
	cmd := []string{"tmux", "set-hook", "-g", scopedHook, command}
	return shell.Execute(cmd...)
}

type Hook string

func IsHook(hook string) bool {
	switch hook {
	case
		string(AfterBindKeyHook),
		string(AfterCapturePaneHook),
		string(AfterCopyModeHook),
		string(AfterDisplayMessageHook),
		string(AfterDisplayPanesHook),
		string(AfterKillPaneHook),
		string(AfterListBuffersHook),
		string(AfterListClientsHook),
		string(AfterListKeysHook),
		string(AfterListPanesHook),
		string(AfterListSessionsHook),
		string(AfterListWindowsHook),
		string(AfterLoadBufferHook),
		string(AfterLockServerHook),
		string(AfterNewSessionHook),
		string(AfterNewWindowHook),
		string(AfterPasteBufferHook),
		string(AfterPipePaneHook),
		string(AfterQueueHook),
		string(AfterRefreshClientHook),
		string(AfterRenameSessionHook),
		string(AfterRenameWindowHook),
		string(AfterResizePaneHook),
		string(AfterResizeWindowHook),
		string(AfterSaveBufferHook),
		string(AfterSelectLayoutHook),
		string(AfterSelectPaneHook),
		string(AfterSelectWindowHook),
		string(AfterSendKeysHook),
		string(AfterSetBufferHook),
		string(AfterSetEnvironmentHook),
		string(AfterSetHookHook),
		string(AfterSetOptionHook),
		string(AfterShowEnvironmentHook),
		string(AfterShowMessagesHook),
		string(AfterShowOptionsHook),
		string(AfterSplitWindowHook),
		string(AfterUnbindKeyHook),
		string(AlertActivityHook),
		string(AlertBellHook),
		string(AlertSilenceHook),
		string(ClientActiveHook),
		string(ClientAttachedHook),
		string(ClientDetachedHook),
		string(ClientFocusInHook),
		string(ClientFocusOutHook),
		string(ClientResizedHook),
		string(ClientSessionChangedHook),
		string(CommandErrorHook),
		string(SessionClosedHook),
		string(SessionCreatedHook),
		string(SessionRenamedHook),
		string(SessionWindowChangedHook),
		string(WindowLinkedHook),
		string(WindowUnlinkedHook):
		return true
	default:
		return false
	}
}

const (
	AfterBindKeyHook         Hook = "after-bind-key"
	AfterCapturePaneHook     Hook = "after-capture-pane"
	AfterCopyModeHook        Hook = "after-copy-mode"
	AfterDisplayMessageHook  Hook = "after-display-message"
	AfterDisplayPanesHook    Hook = "after-display-panes"
	AfterKillPaneHook        Hook = "after-kill-pane"
	AfterListBuffersHook     Hook = "after-list-buffers"
	AfterListClientsHook     Hook = "after-list-clients"
	AfterListKeysHook        Hook = "after-list-keys"
	AfterListPanesHook       Hook = "after-list-panes"
	AfterListSessionsHook    Hook = "after-list-sessions"
	AfterListWindowsHook     Hook = "after-list-windows"
	AfterLoadBufferHook      Hook = "after-load-buffer"
	AfterLockServerHook      Hook = "after-lock-server"
	AfterNewSessionHook      Hook = "after-new-session"
	AfterNewWindowHook       Hook = "after-new-window"
	AfterPasteBufferHook     Hook = "after-paste-buffer"
	AfterPipePaneHook        Hook = "after-pipe-pane"
	AfterQueueHook           Hook = "after-queue"
	AfterRefreshClientHook   Hook = "after-refresh-client"
	AfterRenameSessionHook   Hook = "after-rename-session"
	AfterRenameWindowHook    Hook = "after-rename-window"
	AfterResizePaneHook      Hook = "after-resize-pane"
	AfterResizeWindowHook    Hook = "after-resize-window"
	AfterSaveBufferHook      Hook = "after-save-buffer"
	AfterSelectLayoutHook    Hook = "after-select-layout"
	AfterSelectPaneHook      Hook = "after-select-pane"
	AfterSelectWindowHook    Hook = "after-select-window"
	AfterSendKeysHook        Hook = "after-send-keys"
	AfterSetBufferHook       Hook = "after-set-buffer"
	AfterSetEnvironmentHook  Hook = "after-set-environment"
	AfterSetHookHook         Hook = "after-set-hook"
	AfterSetOptionHook       Hook = "after-set-option"
	AfterShowEnvironmentHook Hook = "after-show-environment"
	AfterShowMessagesHook    Hook = "after-show-messages"
	AfterShowOptionsHook     Hook = "after-show-options"
	AfterSplitWindowHook     Hook = "after-split-window"
	AfterUnbindKeyHook       Hook = "after-unbind-key"
	AlertActivityHook        Hook = "alert-activity"
	AlertBellHook            Hook = "alert-bell"
	AlertSilenceHook         Hook = "alert-silence"
	ClientActiveHook         Hook = "client-active"
	ClientAttachedHook       Hook = "client-attached"
	ClientDetachedHook       Hook = "client-detached"
	ClientFocusInHook        Hook = "client-focus-in"
	ClientFocusOutHook       Hook = "client-focus-out"
	ClientResizedHook        Hook = "client-resized"
	ClientSessionChangedHook Hook = "client-session-changed"
	CommandErrorHook         Hook = "command-error"
	SessionClosedHook        Hook = "session-closed"
	SessionCreatedHook       Hook = "session-created"
	SessionRenamedHook       Hook = "session-renamed"
	SessionWindowChangedHook Hook = "session-window-changed"
	WindowLinkedHook         Hook = "window-linked"
	WindowUnlinkedHook       Hook = "window-unlinked"
)
