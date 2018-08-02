package sdl

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

func init() {
	// NOTE there is a known bug when using SDL2.dll version 2.0.8. Calling any
	// fmt.Print function results in a panic if no printing has been done
	// before. This is here to try to prevent this until we find the cause of
	// this problem.
	fmt.Print()
}

var (
	dll = syscall.NewLazyDLL("SDL2.dll")

	getError                = dll.NewProc("SDL_GetError")
	_init                   = dll.NewProc("SDL_Init")
	quit                    = dll.NewProc("SDL_Quit")
	createWindow            = dll.NewProc("SDL_CreateWindow")
	createWindowAndRenderer = dll.NewProc("SDL_CreateWindowAndRenderer")
	createWindowFrom        = dll.NewProc("SDL_CreateWindowFrom")
	destroyWindow           = dll.NewProc("SDL_DestroyWindow")
	destroyRenderer         = dll.NewProc("SDL_DestroyRenderer")
	pollEvent               = dll.NewProc("SDL_PollEvent")
	setWindowTitle          = dll.NewProc("SDL_SetWindowTitle")
	setWindowSize           = dll.NewProc("SDL_SetWindowSize")
	setWindowPosition       = dll.NewProc("SDL_SetWindowPosition")
	setWindowResizable      = dll.NewProc("SDL_SetWindowResizable")
	showWindow              = dll.NewProc("SDL_ShowWindow")
	hideWindow              = dll.NewProc("SDL_HideWindow")
	showCursor              = dll.NewProc("SDL_ShowCursor")
	getWindowBordersSize    = dll.NewProc("SDL_GetWindowBordersSize")
	getWindowBrightness     = dll.NewProc("SDL_GetWindowBrightness")
	getWindowData           = dll.NewProc("SDL_GetWindowData")
	getWindowDisplayIndex   = dll.NewProc("SDL_GetWindowDisplayIndex")
	getWindowDisplayMode    = dll.NewProc("SDL_GetWindowDisplayMode")
	getWindowFlags          = dll.NewProc("SDL_GetWindowFlags")
	getWindowFromID         = dll.NewProc("SDL_GetWindowFromID")
	getWindowGammaRamp      = dll.NewProc("SDL_GetWindowGammaRamp")
	getWindowGrab           = dll.NewProc("SDL_GetWindowGrab")
	getWindowID             = dll.NewProc("SDL_GetWindowID")
	getWindowMaximumSize    = dll.NewProc("SDL_GetWindowMaximumSize")
	getWindowMinimumSize    = dll.NewProc("SDL_GetWindowMinimumSize")
	getWindowOpacity        = dll.NewProc("SDL_GetWindowOpacity")
	getWindowPixelFormat    = dll.NewProc("SDL_GetWindowPixelFormat")
	getWindowPosition       = dll.NewProc("SDL_GetWindowPosition")
	getWindowSize           = dll.NewProc("SDL_GetWindowSize")
	getWindowSurface        = dll.NewProc("SDL_GetWindowSurface")
	getWindowTitle          = dll.NewProc("SDL_GetWindowTitle")
	getWindowWMInfo         = dll.NewProc("SDL_GetWindowWMInfo")
	maximizeWindow          = dll.NewProc("SDL_MaximizeWindow")
	minimizeWindow          = dll.NewProc("SDL_MinimizeWindow")
	raiseWindow             = dll.NewProc("SDL_RaiseWindow")
	restoreWindow           = dll.NewProc("SDL_RestoreWindow")
	setWindowBordered       = dll.NewProc("SDL_SetWindowBordered")
	setWindowBrightness     = dll.NewProc("SDL_SetWindowBrightness")
	setWindowData           = dll.NewProc("SDL_SetWindowData")
	setWindowDisplayMode    = dll.NewProc("SDL_SetWindowDisplayMode")
	setWindowFullscreen     = dll.NewProc("SDL_SetWindowFullscreen")
	setWindowGammaRamp      = dll.NewProc("SDL_SetWindowGammaRamp")
	setWindowGrab           = dll.NewProc("SDL_SetWindowGrab")
	setWindowHitTest        = dll.NewProc("SDL_SetWindowHitTest")
	setWindowIcon           = dll.NewProc("SDL_SetWindowIcon")
	setWindowInputFocus     = dll.NewProc("SDL_SetWindowInputFocus")
	setWindowMaximumSize    = dll.NewProc("SDL_SetWindowMaximumSize")
	setWindowMinimumSize    = dll.NewProc("SDL_SetWindowMinimumSize")
	setWindowModalFor       = dll.NewProc("SDL_SetWindowModalFor")
	setWindowOpacity        = dll.NewProc("SDL_SetWindowOpacity")
	setWindowsMessageHook   = dll.NewProc("SDL_SetWindowsMessageHook")
	getClipboardText        = dll.NewProc("SDL_GetClipboardText")
	setClipboardText        = dll.NewProc("SDL_SetClipboardText")
	hasClipboardText        = dll.NewProc("SDL_HasClipboardText")
	loadWAV_RW              = dll.NewProc("SDL_LoadWAV_RW")
	freeWAV                 = dll.NewProc("SDL_FreeWAV")
	rwFromFile              = dll.NewProc("SDL_RWFromFile")
	rwFromMem               = dll.NewProc("SDL_RWFromMem")
	rwFromConstMem          = dll.NewProc("SDL_RWFromConstMem")
	rwFromFP                = dll.NewProc("SDL_RWFromDP")
)

func GetError() error {
	if err := getError.Find(); err != nil {
		return err
	}
	ret, _, _ := getError.Call()
	return errors.New(toGoString(ret))
}

func toGoString(cStr uintptr) string {
	if cStr == 0 {
		return ""
	}
	buf := make([]byte, 0, 1024)
	for i := uintptr(0); ; i++ {
		b := *(*byte)(unsafe.Pointer(cStr + i))
		if b == 0 {
			break
		}
		buf = append(buf, b)
	}
	return string(buf)
}

func toSDLString(s string) []byte {
	buf := make([]byte, len(s)+1)
	copy(buf, s) // the last element will be 0
	return buf
}

func toSDLBool(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

func Init(flags uint) error {
	if err := dll.Load(); err != nil {
		return err
	}
	ret, _, _ := _init.Call(uintptr(flags))
	if ret == 0 {
		return nil
	}
	return GetError()
}

func Quit() {
	quit.Call()
}

func CreateWindow(title string, x, y, w, h int, flags uint) (*Window, error) {
	cTitle := toSDLString(title)
	ret, _, _ := createWindow.Call(
		uintptr(unsafe.Pointer(&cTitle[0])),
		uintptr(x),
		uintptr(y),
		uintptr(w),
		uintptr(h),
		uintptr(flags),
	)
	if ret == 0 {
		return nil, GetError()
	}
	return (*Window)(unsafe.Pointer(ret)), nil
}

func CreateWindowAndRenderer(width, height int, flags uint) (*Window, *Renderer, error) {
	var window, renderer uintptr
	ret, _, _ := createWindowAndRenderer.Call(
		uintptr(width),
		uintptr(height),
		uintptr(flags),
		uintptr(unsafe.Pointer(&window)),
		uintptr(unsafe.Pointer(&renderer)),
	)
	if ret == 0 {
		return (*Window)(unsafe.Pointer(window)), (*Renderer)(unsafe.Pointer(renderer)), nil
	}
	return nil, nil, GetError()
}

func CreateWindowFrom(data uintptr) (*Window, error) {
	ret, _, _ := createWindowFrom.Call(data)
	if ret == 0 {
		return nil, GetError()
	}
	return (*Window)(unsafe.Pointer(ret)), nil
}

func (r *Renderer) Destroy() {
	destroyRenderer.Call(uintptr(unsafe.Pointer(r)))
}

func PollEvent() Event {
	var event struct{}
	ret, _, _ := pollEvent.Call(uintptr(unsafe.Pointer(&event)))
	if ret == 0 {
		return nil
	}
	typ := *((*EventType)(unsafe.Pointer(&event)))
	switch typ {
	case WINDOWEVENT:
		return (*WindowEvent)(unsafe.Pointer(&event))
	case SYSWMEVENT:
		return (*SysWMEvent)(unsafe.Pointer(&event))
	case KEYDOWN, KEYUP:
		return (*KeyboardEvent)(unsafe.Pointer(&event))
	case TEXTEDITING:
		return (*TextEditingEvent)(unsafe.Pointer(&event))
	case TEXTINPUT:
		return (*TextInputEvent)(unsafe.Pointer(&event))
	case MOUSEMOTION:
		return (*MouseMotionEvent)(unsafe.Pointer(&event))
	case MOUSEBUTTONDOWN, MOUSEBUTTONUP:
		return (*MouseButtonEvent)(unsafe.Pointer(&event))
	case MOUSEWHEEL:
		return (*MouseWheelEvent)(unsafe.Pointer(&event))
	case JOYAXISMOTION:
		return (*JoyAxisEvent)(unsafe.Pointer(&event))
	case JOYBALLMOTION:
		return (*JoyBallEvent)(unsafe.Pointer(&event))
	case JOYHATMOTION:
		return (*JoyHatEvent)(unsafe.Pointer(&event))
	case JOYBUTTONDOWN, JOYBUTTONUP:
		return (*JoyButtonEvent)(unsafe.Pointer(&event))
	case JOYDEVICEADDED, JOYDEVICEREMOVED:
		return (*JoyDeviceEvent)(unsafe.Pointer(&event))
	case CONTROLLERAXISMOTION:
		return (*ControllerAxisEvent)(unsafe.Pointer(&event))
	case CONTROLLERBUTTONDOWN, CONTROLLERBUTTONUP:
		return (*ControllerButtonEvent)(unsafe.Pointer(&event))
	case CONTROLLERDEVICEADDED, CONTROLLERDEVICEREMOVED, CONTROLLERDEVICEREMAPPED:
		return (*ControllerDeviceEvent)(unsafe.Pointer(&event))
	case AUDIODEVICEADDED, AUDIODEVICEREMOVED:
		return (*AudioDeviceEvent)(unsafe.Pointer(&event))
	case FINGERMOTION, FINGERDOWN, FINGERUP:
		return (*TouchFingerEvent)(unsafe.Pointer(&event))
	case MULTIGESTURE:
		return (*MultiGestureEvent)(unsafe.Pointer(&event))
	case DOLLARGESTURE, DOLLARRECORD:
		return (*DollarGestureEvent)(unsafe.Pointer(&event))
	case DROPFILE, DROPTEXT, DROPBEGIN, DROPCOMPLETE:
		e := (*cDropEvent)(unsafe.Pointer(&event))
		return &DropEvent{
			Type:      e.Type,
			Timestamp: e.Timestamp,
			File:      toGoString(e.File),
			WindowID:  e.WindowID,
		}
	case RENDER_TARGETS_RESET, RENDER_DEVICE_RESET:
		return (*RenderEvent)(unsafe.Pointer(&event))
	case QUIT:
		return (*QuitEvent)(unsafe.Pointer(&event))
	case USEREVENT:
		return (*UserEvent)(unsafe.Pointer(&event))
	case CLIPBOARDUPDATE:
		return (*ClipboardEvent)(unsafe.Pointer(&event))
	default:
		return (*CommonEvent)(unsafe.Pointer(&event))
	}
}

func (w *Window) Destroy() {
	destroyWindow.Call(uintptr(unsafe.Pointer(w)))
}

func (w *Window) SetTitle(title string) {
	cTitle := toSDLString(title)
	setWindowTitle.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(&cTitle[0])),
	)
}

func (w *Window) SetSize(width, height int) {
	setWindowSize.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(width),
		uintptr(height),
	)
}

// SetPosition sets the client area's top-left position.
func (w *Window) SetPosition(x, y int) {
	setWindowPosition.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(x),
		uintptr(y),
	)
}

func (w *Window) SetResizable(resizable bool) {
	setWindowResizable.Call(
		uintptr(unsafe.Pointer(w)),
		toSDLBool(resizable),
	)
}

func (w *Window) Show() {
	showWindow.Call(uintptr(unsafe.Pointer(w)))
}

func (w *Window) Hide() {
	hideWindow.Call(uintptr(unsafe.Pointer(w)))
}

func ShowCursor(toggle int) int {
	ret, _, _ := showCursor.Call(uintptr(toggle))
	return int(ret)
}

func GetWindowFromID(id uint32) (*Window, error) {
	ret, _, _ := getWindowFromID.Call(uintptr(id))
	if ret == 0 {
		return nil, GetError()
	}
	return (*Window)(unsafe.Pointer(ret)), nil
}

func (w *Window) GetBordersSize() (top, left, bottom, right int, err error) {
	ret, _, _ := getWindowBordersSize.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(&top)),
		uintptr(unsafe.Pointer(&left)),
		uintptr(unsafe.Pointer(&bottom)),
		uintptr(unsafe.Pointer(&right)),
	)
	if ret != 0 {
		err = GetError()
	}
	return
}

// GetBrightness returns the brightness for the display where 0.0 is completely
// dark and 1.0 is normal brightness.
// TODO this does not seem to work, returns values in ...E6
func (w *Window) GetBrightness() float32 {
	ret, _, _ := getWindowBrightness.Call(uintptr(unsafe.Pointer(w)))
	return float32(ret)
}

func (w *Window) GetData(name string) uintptr {
	cName := toSDLString(name)
	ret, _, _ := getWindowData.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(&cName[0])),
	)
	return ret
}

func (w *Window) GetDisplayIndex() (int, error) {
	ret, _, _ := getWindowDisplayIndex.Call(uintptr(unsafe.Pointer(w)))
	if int(ret) < 0 {
		return -1, GetError()
	}
	return int(ret), nil
}

func (w *Window) GetDisplayMode() (*DisplayMode, error) {
	var mode DisplayMode
	ret, _, _ := getWindowDisplayMode.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(&mode)),
	)
	if ret == 0 {
		return &mode, nil
	}
	return nil, GetError()
}

func (w *Window) GetFlags() uint32 {
	ret, _, _ := getWindowFlags.Call(uintptr(unsafe.Pointer(w)))
	return uint32(ret)
}

func (w *Window) GetGammaRamp() (red, green, blue [256]uint16, err error) {
	ret, _, _ := getWindowGammaRamp.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(&red[0])),
		uintptr(unsafe.Pointer(&green[0])),
		uintptr(unsafe.Pointer(&blue[0])),
	)
	if ret != 0 {
		err = GetError()
	}
	return
}

func (w *Window) GetGrab() bool {
	ret, _, _ := getWindowGrab.Call(uintptr(unsafe.Pointer(w)))
	return ret != 0
}

func (w *Window) GetID() (uint32, error) {
	ret, _, _ := getWindowID.Call(uintptr(unsafe.Pointer(w)))
	if ret == 0 {
		return 0, GetError()
	}
	return uint32(ret), nil
}

func (w *Window) GetMaximumSize() (width, height int) {
	getWindowMaximumSize.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(&width)),
		uintptr(unsafe.Pointer(&height)),
	)
	return
}

func (w *Window) GetMinimumSize() (width, height int) {
	getWindowMinimumSize.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(&width)),
		uintptr(unsafe.Pointer(&height)),
	)
	return
}

// GetOpacity returns 0.0 for transparent and 1.0f for opaque.
func (w *Window) GetOpacity() (float32, error) {
	var opacity float32
	ret, _, _ := getWindowOpacity.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(&opacity)),
	)
	if ret == 0 {
		return opacity, nil
	}
	return 0, GetError()
}

func (w *Window) GetPixelFormat() (uint32, error) {
	ret, _, _ := getWindowPixelFormat.Call(uintptr(unsafe.Pointer(w)))
	if ret == 0 {
		return 0, GetError()
	}
	return uint32(ret), nil
}

func (w *Window) GetPosition() (x, y int) {
	getWindowPosition.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(&x)),
		uintptr(unsafe.Pointer(&y)),
	)
	return
}

func (w *Window) GetSize() (width, height int) {
	getWindowSize.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(&width)),
		uintptr(unsafe.Pointer(&height)),
	)
	return
}

func (w *Window) GetSurface() (*Surface, error) {
	ret, _, _ := getWindowSurface.Call(uintptr(unsafe.Pointer(w)))
	if ret == 0 {
		return nil, GetError()
	}
	return (*Surface)(unsafe.Pointer(ret)), nil
}

func (w *Window) GetTitle() string {
	ret, _, _ := getWindowTitle.Call(uintptr(unsafe.Pointer(w)))
	return toGoString(ret)
}

func (w *Window) GetWMInfo() (*SysWMInfo, error) {
	var info SysWMInfo
	ret, _, _ := getWindowWMInfo.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(&info)),
	)
	if ret == 0 {
		return nil, GetError()
	}
	return &info, nil
}

func (w *Window) Maximize() {
	maximizeWindow.Call(uintptr(unsafe.Pointer(w)))
}

func (w *Window) Minimize() {
	minimizeWindow.Call(uintptr(unsafe.Pointer(w)))
}

func (w *Window) Raise() {
	raiseWindow.Call(uintptr(unsafe.Pointer(w)))
}

func (w *Window) Restore() {
	restoreWindow.Call(uintptr(unsafe.Pointer(w)))
}

func (w *Window) SetBordered(bordered bool) {
	setWindowBordered.Call(
		uintptr(unsafe.Pointer(w)),
		toSDLBool(bordered),
	)
}

func (w *Window) SetBrightness(brightness float32) error {
	ret, _, _ := setWindowBrightness.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(brightness),
	)
	if ret == 0 {
		return nil
	}
	return GetError()
}

func (w *Window) SetData(name string, data uintptr) uintptr {
	cName := toSDLString(name)
	ret, _, _ := setWindowData.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(&cName[0])),
		data,
	)
	return ret
}

func (w *Window) SetDisplayMode(mode *DisplayMode) error {
	ret, _, _ := setWindowDisplayMode.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(mode)),
	)
	if ret == 0 {
		return nil
	}
	return GetError()
}

func (w *Window) SetFullscreen(flags uint32) error {
	ret, _, _ := setWindowFullscreen.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(flags),
	)
	if ret == 0 {
		return nil
	}
	return GetError()
}

func (w *Window) SetGammaRamp(red, green, blue [256]uint16) error {
	ret, _, _ := setWindowGammaRamp.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(&red[0])),
		uintptr(unsafe.Pointer(&green[0])),
		uintptr(unsafe.Pointer(&blue[0])),
	)
	if ret == 0 {
		return nil
	}
	return GetError()
}

func (w *Window) SetGrab(grabbed bool) {
	setWindowGrab.Call(
		uintptr(unsafe.Pointer(w)),
		toSDLBool(grabbed),
	)
}

func (w *Window) SetHitTest(callback func(window *Window, area *Point, data uintptr), data uintptr) error {
	var ret uintptr
	if callback == nil {
		ret, _, _ = setWindowHitTest.Call(
			uintptr(unsafe.Pointer(w)),
			0,
			data,
		)
	} else {
		ret, _, _ = setWindowHitTest.Call(
			uintptr(unsafe.Pointer(w)),
			syscall.NewCallback(func(window *Window, area *Point, data uintptr) uintptr {
				callback(window, area, data)
				return 0
			}),
			data,
		)
	}
	if ret == 0 {
		return nil
	}
	return GetError()
}

func (w *Window) SetIcon(icon *Surface) {
	setWindowIcon.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(icon)),
	)
}

func (w *Window) SetInputFocus() error {
	ret, _, _ := setWindowInputFocus.Call(uintptr(unsafe.Pointer(w)))
	if ret == 0 {
		return nil
	}
	return GetError()
}

func (w *Window) SetMaximumSize(width, height int) {
	setWindowMaximumSize.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(width),
		uintptr(height),
	)
}

func (w *Window) SetMinimumSize(width, height int) {
	setWindowMinimumSize.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(width),
		uintptr(height),
	)
}

func (w *Window) SetModalFor(parent *Window) error {
	ret, _, _ := setWindowModalFor.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(parent)),
	)
	if ret == 0 {
		return nil
	}
	return GetError()
}

func (w *Window) SetOpacity(opacity float32) error {
	ret, _, _ := setWindowOpacity.Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(opacity),
	)
	if ret == 0 {
		return nil
	}
	return GetError()
}

func SetWindowsMessageHook(callback func(userdata uintptr, hwnd uintptr, msg uint32, w uint64, l int64), userdata uintptr) {
	if callback == nil {
		setWindowsMessageHook.Call(0, userdata)
	} else {
		setWindowsMessageHook.Call(
			syscall.NewCallback(func(userdata uintptr, hwnd uintptr, msg uint32, w uint64, l int64) uintptr {
				callback(userdata, hwnd, msg, w, l)
				return 0
			}),
			userdata,
		)
	}
}

func PointInRect(p *Point, r *Rect) bool {
	return p.X >= r.X && p.X < r.X+r.W && p.Y >= r.Y && p.Y < r.Y+r.H
}

func (r *Rect) Empty() bool {
	return r == nil || r.W <= 0 || r.H <= 0
}

func (r *Rect) Equals(s *Rect) bool {
	return r != nil && s != nil &&
		r.X == s.X && r.Y == s.Y && r.W == s.W && r.H == s.H
}

func GetClipboardText() string {
	ret, _, _ := getClipboardText.Call()
	return toGoString(ret)
}

func SetClipboardText(s string) error {
	c := toSDLString(s)
	ret, _, _ := setClipboardText.Call(uintptr(unsafe.Pointer(&c[0])))
	if ret == 0 {
		return nil
	}
	return GetError()
}

func HasClipboardText() bool {
	ret, _, _ := hasClipboardText.Call()
	return ret != 0
}

//#define SDL_LoadWAV(file, spec, audio_buf, audio_len) \
//    SDL_LoadWAV_RW(SDL_RWFromFile(file, "rb"),1, spec,audio_buf,audio_len)
func LoadWav(file string, spec *AudioSpec) (*WAV, *AudioSpec, error) {
	ops, err := RWFromFile(file, "rb")
	if err != nil {
		return nil, nil, err
	}
	return LoadWAV_RW(ops, true, spec)
}

func LoadWAV_RW(src *RWops, freesrc bool, spec *AudioSpec) (*WAV, *AudioSpec, error) {
	var wav WAV
	ret, _, _ := loadWAV_RW.Call(
		uintptr(unsafe.Pointer(src)),
		toSDLBool(freesrc),
		uintptr(unsafe.Pointer(spec)),
		uintptr(unsafe.Pointer(&wav.buf)),
		uintptr(unsafe.Pointer(&wav.len)),
	)
	if ret == 0 {
		return nil, nil, GetError()
	}
	return &wav, (*AudioSpec)(unsafe.Pointer(ret)), nil
}

type WAV struct {
	buf *uint8
	len uint32
}

func (wav *WAV) Free() {
	freeWAV.Call(uintptr(unsafe.Pointer(wav.buf)))
}

func RWFromFile(file, mode string) (*RWops, error) {
	cFile := toSDLString(file)
	cMode := toSDLString(mode)
	ret, _, _ := rwFromFile.Call(
		uintptr(unsafe.Pointer(&cFile[0])),
		uintptr(unsafe.Pointer(&cMode[0])),
	)
	if ret == 0 {
		return nil, GetError()
	}
	return (*RWops)(unsafe.Pointer(ret)), nil
}

func RWFromMem(mem unsafe.Pointer, size int) (*RWops, error) {
	ret, _, _ := rwFromMem.Call(uintptr(mem), uintptr(size))
	if ret == 0 {
		return nil, GetError()
	}
	return (*RWops)(unsafe.Pointer(ret)), nil
}

func RWFromBytes(b []byte) (*RWops, error) {
	ret, _, _ := rwFromMem.Call(
		uintptr(unsafe.Pointer(&b[0])),
		uintptr(len(b)),
	)
	if ret == 0 {
		return nil, GetError()
	}
	return (*RWops)(unsafe.Pointer(ret)), nil
}

func RWFromConstMem(mem unsafe.Pointer, size int) (*RWops, error) {
	ret, _, _ := rwFromConstMem.Call(uintptr(mem), uintptr(size))
	if ret == 0 {
		return nil, GetError()
	}
	return (*RWops)(unsafe.Pointer(ret)), nil
}

func RWFromConstBytes(b []byte) (*RWops, error) {
	ret, _, _ := rwFromConstMem.Call(
		uintptr(unsafe.Pointer(&b[0])),
		uintptr(len(b)),
	)
	if ret == 0 {
		return nil, GetError()
	}
	return (*RWops)(unsafe.Pointer(ret)), nil
}

func RWFromFP(fp unsafe.Pointer, autoClose bool) (*RWops, error) {
	ret, _, _ := rwFromFP.Call(uintptr(fp), toSDLBool(autoClose))
	if ret == 0 {
		return nil, GetError()
	}
	return (*RWops)(unsafe.Pointer(ret)), nil
}

func (rw *RWops) Size() (int64, error) {
	ret, _, _ := syscall.Syscall(rw.size, 1, uintptr(unsafe.Pointer(rw)), 0, 0)
	size := int64(ret)
	if size < -1 {
		return size, GetError()
	}
	return size, nil
}
func (rw *RWops) Seek(offset int64, whence int) (int64, error) {
	ret, _, _ := syscall.Syscall(
		rw.seek, 3,
		uintptr(unsafe.Pointer(rw)),
		uintptr(offset),
		uintptr(whence),
	)
	pos := int64(ret)
	if pos == -1 {
		return -1, GetError()
	}
	return pos, nil
}

func (rw *RWops) Tell() (int64, error) {
	return rw.Seek(0, RW_SEEK_CUR)
}

func (rw *RWops) Read(dest unsafe.Pointer, size uint, maxnum uint) (uint, error) {
	ret, _, _ := syscall.Syscall6(
		rw.read, 4,
		uintptr(unsafe.Pointer(rw)),
		uintptr(dest),
		uintptr(size),
		uintptr(maxnum),
		0, 0,
	)
	if ret == 0 {
		return 0, GetError()
	}
	return uint(ret), nil
}

func (rw *RWops) ReadBytes(dest []byte) (uint, error) {
	ret, _, _ := syscall.Syscall6(
		rw.read, 4,
		uintptr(unsafe.Pointer(rw)),
		uintptr(unsafe.Pointer(&dest[0])),
		uintptr(1),
		uintptr(len(dest)),
		0, 0,
	)
	if ret == 0 {
		return 0, GetError()
	}
	return uint(ret), nil
}

func (rw *RWops) Write(src unsafe.Pointer, size uint, num uint) (uint, error) {
	ret, _, _ := syscall.Syscall6(
		rw.write, 4,
		uintptr(unsafe.Pointer(rw)),
		uintptr(src),
		uintptr(size),
		uintptr(num),
		0, 0,
	)
	written := uint(ret)
	if written < num {
		return written, GetError()
	}
	return written, nil
}

func (rw *RWops) WriteBytes(b []byte) (uint, error) {
	ret, _, _ := syscall.Syscall6(
		rw.write, 4,
		uintptr(unsafe.Pointer(rw)),
		uintptr(unsafe.Pointer(&b[0])),
		uintptr(1),
		uintptr(len(b)),
		0, 0,
	)
	written := uint(ret)
	if written < uint(len(b)) {
		return written, GetError()
	}
	return written, nil
}

func (rw *RWops) Close() error {
	ret, _, _ := syscall.Syscall(rw.close, 1, uintptr(unsafe.Pointer(rw)), 0, 0)
	if ret == 0 {
		return nil
	}
	return GetError()
}
