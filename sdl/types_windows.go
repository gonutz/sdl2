package sdl

import "unsafe"

// basic types

type Window struct{}
type Renderer struct{}
type Surface struct{}
type Texture struct{}
type RWops struct {
	size  uintptr
	seek  uintptr
	read  uintptr
	write uintptr
	close uintptr
}

// events

type Event interface{}

type EventType uint32

type CommonEvent struct {
	Type      EventType
	Timestamp uint32
}

type WindowEvent struct {
	Type      EventType
	Timestamp uint32
	WindowID  uint32
	Event     uint8
	padding1  uint8
	padding2  uint8
	padding3  uint8
	Data1     int32
	Data2     int32
}

type KeyboardEvent struct {
	Type      EventType
	Timestamp uint32
	WindowID  uint32
	State     uint8
	Repeat    uint8
	padding2  uint8
	padding3  uint8
	Keysym    Keysym
}

type TextEditingEvent struct {
	Type      EventType
	Timestamp uint32
	WindowID  uint32
	Text      [TEXTEDITINGEVENT_TEXT_SIZE]byte
	Start     int32
	Length    int32
}

type TextInputEvent struct {
	Type      EventType
	Timestamp uint32
	WindowID  uint32
	Text      [TEXTEDITINGEVENT_TEXT_SIZE]byte
}

type MouseMotionEvent struct {
	Type      EventType
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	State     uint32
	X         int32
	Y         int32
	XRel      int32
	YRel      int32
}

type MouseButtonEvent struct {
	Type      EventType
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	Button    uint8
	State     uint8
	Clicks    uint8
	padding1  uint8
	X         int32
	Y         int32
}

type MouseWheelEvent struct {
	Type      EventType
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	X         int32
	Y         int32
	Direction uint32
}

type JoyAxisEvent struct {
	Type      EventType
	Timestamp uint32
	Which     JoystickID
	Axis      uint8
	padding1  uint8
	padding2  uint8
	padding3  uint8
	Value     int16
	padding4  uint16
}

type JoyBallEvent struct {
	Type      EventType
	Timestamp uint32
	Which     JoystickID
	Ball      uint8
	padding1  uint8
	padding2  uint8
	padding3  uint8
	XRel      int16
	YRel      int16
}

type JoyHatEvent struct {
	Type      EventType
	Timestamp uint32
	Which     JoystickID
	Hat       uint8
	Value     uint8
	padding1  uint8
	padding2  uint8
}

type JoyButtonEvent struct {
	Type      EventType
	Timestamp uint32
	Which     JoystickID
	Button    uint8
	State     uint8
	padding1  uint8
	padding2  uint8
}

type JoyDeviceEvent struct {
	Type      EventType
	Timestamp uint32
	Which     int32
}

type ControllerAxisEvent struct {
	Type      EventType
	Timestamp uint32
	Which     JoystickID
	Axis      uint8
	padding1  uint8
	padding2  uint8
	padding3  uint8
	Value     int16
	padding4  uint16
}

type ControllerButtonEvent struct {
	Type      EventType
	Timestamp uint32
	Which     JoystickID
	Button    uint8
	State     uint8
	padding1  uint8
	padding2  uint8
}

type ControllerDeviceEvent struct {
	Type      EventType
	Timestamp uint32
	Which     int32
}

type AudioDeviceEvent struct {
	Type      EventType
	Timestamp uint32
	Which     uint32
	IsCapture uint8
	padding1  uint8
	padding2  uint8
	padding3  uint8
}

type TouchFingerEvent struct {
	Type      EventType
	Timestamp uint32
	TouchID   TouchID
	FingerID  FingerID
	X         float32
	Y         float32
	DX        float32
	DY        float32
	Pressure  float32
}

type MultiGestureEvent struct {
	Type       EventType
	Timestamp  uint32
	TouchID    TouchID
	DTheta     float32
	DDist      float32
	X          float32
	Y          float32
	NumFingers uint16
	padding    uint16
}

type DollarGestureEvent struct {
	Type       EventType
	Timestamp  uint32
	TouchID    TouchID
	GestureID  GestureID
	NumFingers uint32
	Error      float32
	X          float32
	Y          float32
}

type cDropEvent struct {
	Type      EventType
	Timestamp uint32
	File      uintptr
	WindowID  uint32
}

type DropEvent struct {
	Type      EventType
	Timestamp uint32
	File      string
	WindowID  uint32
}

type QuitEvent struct {
	Type      EventType
	Timestamp uint32
}

type OSEvent struct {
	Type      EventType
	Timestamp uint32
}

type UserEvent struct {
	Type      EventType
	Timestamp uint32
	WindowID  uint32
	Code      int32
	Data1     unsafe.Pointer
	Data2     unsafe.Pointer
}

type SysWMEvent struct {
	Type      EventType
	Timestamp uint32
	Msg       *SysWMmsg
}

type SysWMmsg struct {
	Version   Version
	Subsystem SYSWM_TYPE
	Hwnd      HWND
	Msg       uint
	WParam    WPARAM
	LParam    LPARAM
}

type RenderEvent CommonEvent
type ClipboardEvent CommonEvent

type DisplayMode struct {
	Format      uint32
	W           int32
	H           int32
	RefreshRate int32
	DriverData  uintptr
}

type SysWMInfo struct {
	Version   Version
	Subsystem SYSWM_TYPE
	Window    HWND
	Hdc       HDC
	Instance  HINSTANCE
}

type JoystickID int32

type TouchID int64
type FingerID int64

type GestureID int64

type HWND uintptr
type WPARAM uintptr
type LPARAM uintptr

type HDC uintptr
type HINSTANCE uintptr

type SYSWM_TYPE int32
type Scancode int32

type Version struct {
	Major uint8
	Minor uint8
	Patch uint8
}

type Keysym struct {
	Scancode Scancode
	Sym      Keycode
	Mod      uint16
	unused   uint32
}

type Keycode int32

type Point struct {
	X, Y int32
}

type Rect struct {
	X, Y, W, H int32
}

// These are what the 16 bits in SDL_AudioFormat currently mean, unspecified
// bits are always zero).
//
//   ++-----------------------sample is signed if set
//   ||
//   ||       ++-----------sample is bigendian if set
//   ||       ||
//   ||       ||          ++---sample is float if set
//   ||       ||          ||
//   ||       ||          || +---sample bit size---+
//   ||       ||          || |                     |
//   15 14 13 12 11 10 09 08 07 06 05 04 03 02 01 00
//
// There are macros in SDL 2.0 and later to query these bits.
type AudioFormat uint16

type Color struct {
	R, G, B, A uint8
}

type Palette struct {
	// TODO
}

type PixelFormat struct {
	Format        uint32
	Palette       *Palette
	BitsPerPixel  uint8
	BytesPerPixel uint8
	padding       [2]uint8
	Rmask         uint32
	Gmask         uint32
	Bmask         uint32
	Amask         uint32
	Rloss         uint8
	Gloss         uint8
	Bloss         uint8
	Aloss         uint8
	Rshift        uint8
	Gshift        uint8
	Bshift        uint8
	Ashift        uint8
	Refcount      int32
	Next          *PixelFormat
}

type AudioSpec struct{}
