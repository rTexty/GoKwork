package types

// EventType константы типов событий
const (
	EventTypeIsTyping         = "is_typing"
	EventTypeNotify           = "notify"
	EventTypeNewMessage       = "new_inbox"
	EventTypePopUpNotify      = "pop_up_notify"
	EventTypeMessageDelete    = "inbox_message_delete"
	EventTypeRemovePopUpNotify = "remove_pop_up_notify"
	EventTypeDialogUpdate     = "dialog_updated"
)

// Notify константы уведомлений
const (
	NotifyNewMessage = "new_message"
)

// BaseEvent представляет базовое событие
type BaseEvent struct {
	Event string                 `json:"event"`
	Data  map[string]interface{} `json:"data"`
}
