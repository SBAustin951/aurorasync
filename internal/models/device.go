package models

import "github.com/google/uuid"

// ==========================
// Generic Types
// ==========================

type Capability struct {
	Type       string     `json:"type"`
	Instance   string     `json:"instance"`
	Parameters *Parameter `json:"parameters,omitempty"`
	Value      *int       `json:"value,omitempty"`
	State      *State     `json:"state,omitempty"`
}

type Parameter struct {
	DataType string    `json:"datatype"`
	Unit     *string   `json:"unit,omitempty"`
	Options  *[]Option `json:"options,omitempty"`
	Range    *Range    `json:"range,omitempty"`
	Fields   *[]Field  `json:"fields,omitempty"`
}

type Option struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Field struct {
	FieldName    string `json:"fieldName"`
	DataType     string `json:"dataType"`
	ElementType  string `json:"elementType"`
	Required     bool   `json:"required"`
	Size         Range  `json:"size"`
	ElementRange Range  `json:"elementRange"`
}

type Range struct {
	Min       int  `json:"min"`
	Max       int  `json:"max"`
	Precision *int `json:"precision,omitempty"`
}

type State struct {
	Status string `json:"status,omitempty"`
	Value  any    `json:"value,omitempty"`
}

// ==========================
// Device Response
// ==========================

type DeviceResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Devices []Device `json:"data"`
}

type Device struct {
	SKU          string       `json:"sku"`
	DeviceID     string       `json:"device"`
	Name         string       `json:"deviceName"`
	Type         string       `json:"type"`
	Capabilities []Capability `json:"capabilities"`
}

// ==========================
// Control Request/Response
// ==========================

type ControlRequest struct {
	RequestID uuid.UUID      `json:"requestId"`
	Payload   ControlPayload `json:"payload"`
}

type ControlPayload struct {
	SKU        string     `json:"sku"`
	DeviceID   string     `json:"device"`
	Capability Capability `json:"capability"`
}

type ControlResponse struct {
	Request    uuid.UUID  `json:"requestId"`
	Message    string     `json:"msg"`
	Code       int        `json:"code"`
	Capability Capability `json:"capability"`
}

// ==========================
// Device State Request/Response
// ==========================

type DeviceStateRequest struct {
	RequestID uuid.UUID    `json:"requestId"`
	Payload   StatePayload `json:"payload"`
}

type StatePayload struct {
	SKU      string `json:"sku"`
	DeviceID string `json:"device"`
}

type DeviceStateResponse struct {
	RequestID uuid.UUID           `json:"requestId"`
	Message   string              `json:"msg"`
	Code      int                 `json:"code"`
	Payload   *DeviceStatePayload `json:"data,omitempty"`
}

type DeviceStatePayload struct {
	SKU          string       `json:"sku"`
	DeviceID     string       `json:"device"`
	Capabilities []Capability `json:"capabilities"`
}
