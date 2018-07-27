package sdl

const (
	INIT_TIMER          = 0x00000001
	INIT_AUDIO          = 0x00000010
	INIT_VIDEO          = 0x00000020 // automatically inits events subsystem
	INIT_JOYSTICK       = 0x00000200 // automatically inits events subsystem
	INIT_HAPTIC         = 0x00001000 // force feedback
	INIT_GAMECONTROLLER = 0x00002000 // automatically inits events subsystem
	INIT_EVENTS         = 0x00004000
	INIT_NOPARACHUTE    = 0x00100000 // compatibility; this flag is ignored

	INIT_EVERYTHING = INIT_TIMER | INIT_AUDIO | INIT_VIDEO | INIT_EVENTS | INIT_JOYSTICK | INIT_HAPTIC | INIT_GAMECONTROLLER
)

const TEXTEDITINGEVENT_TEXT_SIZE = 32

const (
	SCANCODE_UNKNOWN            Scancode = 0
	SCANCODE_A                           = 4
	SCANCODE_B                           = 5
	SCANCODE_C                           = 6
	SCANCODE_D                           = 7
	SCANCODE_E                           = 8
	SCANCODE_F                           = 9
	SCANCODE_G                           = 10
	SCANCODE_H                           = 11
	SCANCODE_I                           = 12
	SCANCODE_J                           = 13
	SCANCODE_K                           = 14
	SCANCODE_L                           = 15
	SCANCODE_M                           = 16
	SCANCODE_N                           = 17
	SCANCODE_O                           = 18
	SCANCODE_P                           = 19
	SCANCODE_Q                           = 20
	SCANCODE_R                           = 21
	SCANCODE_S                           = 22
	SCANCODE_T                           = 23
	SCANCODE_U                           = 24
	SCANCODE_V                           = 25
	SCANCODE_W                           = 26
	SCANCODE_X                           = 27
	SCANCODE_Y                           = 28
	SCANCODE_Z                           = 29
	SCANCODE_1                           = 30
	SCANCODE_2                           = 31
	SCANCODE_3                           = 32
	SCANCODE_4                           = 33
	SCANCODE_5                           = 34
	SCANCODE_6                           = 35
	SCANCODE_7                           = 36
	SCANCODE_8                           = 37
	SCANCODE_9                           = 38
	SCANCODE_0                           = 39
	SCANCODE_RETURN                      = 40
	SCANCODE_ESCAPE                      = 41
	SCANCODE_BACKSPACE                   = 42
	SCANCODE_TAB                         = 43
	SCANCODE_SPACE                       = 44
	SCANCODE_MINUS                       = 45
	SCANCODE_EQUALS                      = 46
	SCANCODE_LEFTBRACKET                 = 47
	SCANCODE_RIGHTBRACKET                = 48
	SCANCODE_BACKSLASH                   = 49
	SCANCODE_NONUSHASH                   = 50
	SCANCODE_SEMICOLON                   = 51
	SCANCODE_APOSTROPHE                  = 52
	SCANCODE_GRAVE                       = 53
	SCANCODE_COMMA                       = 54
	SCANCODE_PERIOD                      = 55
	SCANCODE_SLASH                       = 56
	SCANCODE_CAPSLOCK                    = 57
	SCANCODE_F1                          = 58
	SCANCODE_F2                          = 59
	SCANCODE_F3                          = 60
	SCANCODE_F4                          = 61
	SCANCODE_F5                          = 62
	SCANCODE_F6                          = 63
	SCANCODE_F7                          = 64
	SCANCODE_F8                          = 65
	SCANCODE_F9                          = 66
	SCANCODE_F10                         = 67
	SCANCODE_F11                         = 68
	SCANCODE_F12                         = 69
	SCANCODE_PRINTSCREEN                 = 70
	SCANCODE_SCROLLLOCK                  = 71
	SCANCODE_PAUSE                       = 72
	SCANCODE_INSERT                      = 73
	SCANCODE_HOME                        = 74
	SCANCODE_PAGEUP                      = 75
	SCANCODE_DELETE                      = 76
	SCANCODE_END                         = 77
	SCANCODE_PAGEDOWN                    = 78
	SCANCODE_RIGHT                       = 79
	SCANCODE_LEFT                        = 80
	SCANCODE_DOWN                        = 81
	SCANCODE_UP                          = 82
	SCANCODE_NUMLOCKCLEAR                = 83
	SCANCODE_KP_DIVIDE                   = 84
	SCANCODE_KP_MULTIPLY                 = 85
	SCANCODE_KP_MINUS                    = 86
	SCANCODE_KP_PLUS                     = 87
	SCANCODE_KP_ENTER                    = 88
	SCANCODE_KP_1                        = 89
	SCANCODE_KP_2                        = 90
	SCANCODE_KP_3                        = 91
	SCANCODE_KP_4                        = 92
	SCANCODE_KP_5                        = 93
	SCANCODE_KP_6                        = 94
	SCANCODE_KP_7                        = 95
	SCANCODE_KP_8                        = 96
	SCANCODE_KP_9                        = 97
	SCANCODE_KP_0                        = 98
	SCANCODE_KP_PERIOD                   = 99
	SCANCODE_NONUSBACKSLASH              = 100
	SCANCODE_APPLICATION                 = 101
	SCANCODE_POWER                       = 102
	SCANCODE_KP_EQUALS                   = 103
	SCANCODE_F13                         = 104
	SCANCODE_F14                         = 105
	SCANCODE_F15                         = 106
	SCANCODE_F16                         = 107
	SCANCODE_F17                         = 108
	SCANCODE_F18                         = 109
	SCANCODE_F19                         = 110
	SCANCODE_F20                         = 111
	SCANCODE_F21                         = 112
	SCANCODE_F22                         = 113
	SCANCODE_F23                         = 114
	SCANCODE_F24                         = 115
	SCANCODE_EXECUTE                     = 116
	SCANCODE_HELP                        = 117
	SCANCODE_MENU                        = 118
	SCANCODE_SELECT                      = 119
	SCANCODE_STOP                        = 120
	SCANCODE_AGAIN                       = 121
	SCANCODE_UNDO                        = 122
	SCANCODE_CUT                         = 123
	SCANCODE_COPY                        = 124
	SCANCODE_PASTE                       = 125
	SCANCODE_FIND                        = 126
	SCANCODE_MUTE                        = 127
	SCANCODE_VOLUMEUP                    = 128
	SCANCODE_VOLUMEDOWN                  = 129
	SCANCODE_KP_COMMA                    = 133
	SCANCODE_KP_EQUALSAS400              = 134
	SCANCODE_INTERNATIONAL1              = 135
	SCANCODE_INTERNATIONAL2              = 136
	SCANCODE_INTERNATIONAL3              = 137
	SCANCODE_INTERNATIONAL4              = 138
	SCANCODE_INTERNATIONAL5              = 139
	SCANCODE_INTERNATIONAL6              = 140
	SCANCODE_INTERNATIONAL7              = 141
	SCANCODE_INTERNATIONAL8              = 142
	SCANCODE_INTERNATIONAL9              = 143
	SCANCODE_LANG1                       = 144
	SCANCODE_LANG2                       = 145
	SCANCODE_LANG3                       = 146
	SCANCODE_LANG4                       = 147
	SCANCODE_LANG5                       = 148
	SCANCODE_LANG6                       = 149
	SCANCODE_LANG7                       = 150
	SCANCODE_LANG8                       = 151
	SCANCODE_LANG9                       = 152
	SCANCODE_ALTERASE                    = 153
	SCANCODE_SYSREQ                      = 154
	SCANCODE_CANCEL                      = 155
	SCANCODE_CLEAR                       = 156
	SCANCODE_PRIOR                       = 157
	SCANCODE_RETURN2                     = 158
	SCANCODE_SEPARATOR                   = 159
	SCANCODE_OUT                         = 160
	SCANCODE_OPER                        = 161
	SCANCODE_CLEARAGAIN                  = 162
	SCANCODE_CRSEL                       = 163
	SCANCODE_EXSEL                       = 164
	SCANCODE_KP_00                       = 176
	SCANCODE_KP_000                      = 177
	SCANCODE_THOUSANDSSEPARATOR          = 178
	SCANCODE_DECIMALSEPARATOR            = 179
	SCANCODE_CURRENCYUNIT                = 180
	SCANCODE_CURRENCYSUBUNIT             = 181
	SCANCODE_KP_LEFTPAREN                = 182
	SCANCODE_KP_RIGHTPAREN               = 183
	SCANCODE_KP_LEFTBRACE                = 184
	SCANCODE_KP_RIGHTBRACE               = 185
	SCANCODE_KP_TAB                      = 186
	SCANCODE_KP_BACKSPACE                = 187
	SCANCODE_KP_A                        = 188
	SCANCODE_KP_B                        = 189
	SCANCODE_KP_C                        = 190
	SCANCODE_KP_D                        = 191
	SCANCODE_KP_E                        = 192
	SCANCODE_KP_F                        = 193
	SCANCODE_KP_XOR                      = 194
	SCANCODE_KP_POWER                    = 195
	SCANCODE_KP_PERCENT                  = 196
	SCANCODE_KP_LESS                     = 197
	SCANCODE_KP_GREATER                  = 198
	SCANCODE_KP_AMPERSAND                = 199
	SCANCODE_KP_DBLAMPERSAND             = 200
	SCANCODE_KP_VERTICALBAR              = 201
	SCANCODE_KP_DBLVERTICALBAR           = 202
	SCANCODE_KP_COLON                    = 203
	SCANCODE_KP_HASH                     = 204
	SCANCODE_KP_SPACE                    = 205
	SCANCODE_KP_AT                       = 206
	SCANCODE_KP_EXCLAM                   = 207
	SCANCODE_KP_MEMSTORE                 = 208
	SCANCODE_KP_MEMRECALL                = 209
	SCANCODE_KP_MEMCLEAR                 = 210
	SCANCODE_KP_MEMADD                   = 211
	SCANCODE_KP_MEMSUBTRACT              = 212
	SCANCODE_KP_MEMMULTIPLY              = 213
	SCANCODE_KP_MEMDIVIDE                = 214
	SCANCODE_KP_PLUSMINUS                = 215
	SCANCODE_KP_CLEAR                    = 216
	SCANCODE_KP_CLEARENTRY               = 217
	SCANCODE_KP_BINARY                   = 218
	SCANCODE_KP_OCTAL                    = 219
	SCANCODE_KP_DECIMAL                  = 220
	SCANCODE_KP_HEXADECIMAL              = 221
	SCANCODE_LCTRL                       = 224
	SCANCODE_LSHIFT                      = 225
	SCANCODE_LALT                        = 226
	SCANCODE_LGUI                        = 227
	SCANCODE_RCTRL                       = 228
	SCANCODE_RSHIFT                      = 229
	SCANCODE_RALT                        = 230
	SCANCODE_RGUI                        = 231
	SCANCODE_MODE                        = 257
	SCANCODE_AUDIONEXT                   = 258
	SCANCODE_AUDIOPREV                   = 259
	SCANCODE_AUDIOSTOP                   = 260
	SCANCODE_AUDIOPLAY                   = 261
	SCANCODE_AUDIOMUTE                   = 262
	SCANCODE_MEDIASELECT                 = 263
	SCANCODE_WWW                         = 264
	SCANCODE_MAIL                        = 265
	SCANCODE_CALCULATOR                  = 266
	SCANCODE_COMPUTER                    = 267
	SCANCODE_AC_SEARCH                   = 268
	SCANCODE_AC_HOME                     = 269
	SCANCODE_AC_BACK                     = 270
	SCANCODE_AC_FORWARD                  = 271
	SCANCODE_AC_STOP                     = 272
	SCANCODE_AC_REFRESH                  = 273
	SCANCODE_AC_BOOKMARKS                = 274
	SCANCODE_BRIGHTNESSDOWN              = 275
	SCANCODE_BRIGHTNESSUP                = 276
	SCANCODE_DISPLAYSWITCH               = 277
	SCANCODE_KBDILLUMTOGGLE              = 278
	SCANCODE_KBDILLUMDOWN                = 279
	SCANCODE_KBDILLUMUP                  = 280
	SCANCODE_EJECT                       = 281
	SCANCODE_SLEEP                       = 282
	SCANCODE_APP1                        = 283
	SCANCODE_APP2                        = 284
	SCANCODE_AUDIOREWIND                 = 285
	SCANCODE_AUDIOFASTFORWARD            = 286
)

// Keycode values

const K_SCANCODE_MASK = 1 << 30
const (
	K_UNKNOWN            Keycode = 0
	K_RETURN                     = '\r'
	K_ESCAPE                     = '\033'
	K_BACKSPACE                  = '\b'
	K_TAB                        = '\t'
	K_SPACE                      = ' '
	K_EXCLAIM                    = '!'
	K_QUOTEDBL                   = '"'
	K_HASH                       = '#'
	K_PERCENT                    = '%'
	K_DOLLAR                     = '$'
	K_AMPERSAND                  = '&'
	K_QUOTE                      = '\''
	K_LEFTPAREN                  = '('
	K_RIGHTPAREN                 = ')'
	K_ASTERISK                   = '*'
	K_PLUS                       = '+'
	K_COMMA                      = ','
	K_MINUS                      = '-'
	K_PERIOD                     = '.'
	K_SLASH                      = '/'
	K_0                          = '0'
	K_1                          = '1'
	K_2                          = '2'
	K_3                          = '3'
	K_4                          = '4'
	K_5                          = '5'
	K_6                          = '6'
	K_7                          = '7'
	K_8                          = '8'
	K_9                          = '9'
	K_COLON                      = ':'
	K_SEMICOLON                  = ';'
	K_LESS                       = '<'
	K_EQUALS                     = '='
	K_GREATER                    = '>'
	K_QUESTION                   = '?'
	K_AT                         = '@'
	K_LEFTBRACKET                = '['
	K_BACKSLASH                  = '\\'
	K_RIGHTBRACKET               = ']'
	K_CARET                      = '^'
	K_UNDERSCORE                 = '_'
	K_BACKQUOTE                  = '`'
	K_a                          = 'a'
	K_b                          = 'b'
	K_c                          = 'c'
	K_d                          = 'd'
	K_e                          = 'e'
	K_f                          = 'f'
	K_g                          = 'g'
	K_h                          = 'h'
	K_i                          = 'i'
	K_j                          = 'j'
	K_k                          = 'k'
	K_l                          = 'l'
	K_m                          = 'm'
	K_n                          = 'n'
	K_o                          = 'o'
	K_p                          = 'p'
	K_q                          = 'q'
	K_r                          = 'r'
	K_s                          = 's'
	K_t                          = 't'
	K_u                          = 'u'
	K_v                          = 'v'
	K_w                          = 'w'
	K_x                          = 'x'
	K_y                          = 'y'
	K_z                          = 'z'
	K_CAPSLOCK                   = K_SCANCODE_MASK | SCANCODE_CAPSLOCK
	K_F1                         = K_SCANCODE_MASK | SCANCODE_F1
	K_F2                         = K_SCANCODE_MASK | SCANCODE_F2
	K_F3                         = K_SCANCODE_MASK | SCANCODE_F3
	K_F4                         = K_SCANCODE_MASK | SCANCODE_F4
	K_F5                         = K_SCANCODE_MASK | SCANCODE_F5
	K_F6                         = K_SCANCODE_MASK | SCANCODE_F6
	K_F7                         = K_SCANCODE_MASK | SCANCODE_F7
	K_F8                         = K_SCANCODE_MASK | SCANCODE_F8
	K_F9                         = K_SCANCODE_MASK | SCANCODE_F9
	K_F10                        = K_SCANCODE_MASK | SCANCODE_F10
	K_F11                        = K_SCANCODE_MASK | SCANCODE_F11
	K_F12                        = K_SCANCODE_MASK | SCANCODE_F12
	K_PRINTSCREEN                = K_SCANCODE_MASK | SCANCODE_PRINTSCREEN
	K_SCROLLLOCK                 = K_SCANCODE_MASK | SCANCODE_SCROLLLOCK
	K_PAUSE                      = K_SCANCODE_MASK | SCANCODE_PAUSE
	K_INSERT                     = K_SCANCODE_MASK | SCANCODE_INSERT
	K_HOME                       = K_SCANCODE_MASK | SCANCODE_HOME
	K_PAGEUP                     = K_SCANCODE_MASK | SCANCODE_PAGEUP
	K_DELETE                     = '\177'
	K_END                        = K_SCANCODE_MASK | SCANCODE_END
	K_PAGEDOWN                   = K_SCANCODE_MASK | SCANCODE_PAGEDOWN
	K_RIGHT                      = K_SCANCODE_MASK | SCANCODE_RIGHT
	K_LEFT                       = K_SCANCODE_MASK | SCANCODE_LEFT
	K_DOWN                       = K_SCANCODE_MASK | SCANCODE_DOWN
	K_UP                         = K_SCANCODE_MASK | SCANCODE_UP
	K_NUMLOCKCLEAR               = K_SCANCODE_MASK | SCANCODE_NUMLOCKCLEAR
	K_KP_DIVIDE                  = K_SCANCODE_MASK | SCANCODE_KP_DIVIDE
	K_KP_MULTIPLY                = K_SCANCODE_MASK | SCANCODE_KP_MULTIPLY
	K_KP_MINUS                   = K_SCANCODE_MASK | SCANCODE_KP_MINUS
	K_KP_PLUS                    = K_SCANCODE_MASK | SCANCODE_KP_PLUS
	K_KP_ENTER                   = K_SCANCODE_MASK | SCANCODE_KP_ENTER
	K_KP_1                       = K_SCANCODE_MASK | SCANCODE_KP_1
	K_KP_2                       = K_SCANCODE_MASK | SCANCODE_KP_2
	K_KP_3                       = K_SCANCODE_MASK | SCANCODE_KP_3
	K_KP_4                       = K_SCANCODE_MASK | SCANCODE_KP_4
	K_KP_5                       = K_SCANCODE_MASK | SCANCODE_KP_5
	K_KP_6                       = K_SCANCODE_MASK | SCANCODE_KP_6
	K_KP_7                       = K_SCANCODE_MASK | SCANCODE_KP_7
	K_KP_8                       = K_SCANCODE_MASK | SCANCODE_KP_8
	K_KP_9                       = K_SCANCODE_MASK | SCANCODE_KP_9
	K_KP_0                       = K_SCANCODE_MASK | SCANCODE_KP_0
	K_KP_PERIOD                  = K_SCANCODE_MASK | SCANCODE_KP_PERIOD
	K_APPLICATION                = K_SCANCODE_MASK | SCANCODE_APPLICATION
	K_POWER                      = K_SCANCODE_MASK | SCANCODE_POWER
	K_KP_EQUALS                  = K_SCANCODE_MASK | SCANCODE_KP_EQUALS
	K_F13                        = K_SCANCODE_MASK | SCANCODE_F13
	K_F14                        = K_SCANCODE_MASK | SCANCODE_F14
	K_F15                        = K_SCANCODE_MASK | SCANCODE_F15
	K_F16                        = K_SCANCODE_MASK | SCANCODE_F16
	K_F17                        = K_SCANCODE_MASK | SCANCODE_F17
	K_F18                        = K_SCANCODE_MASK | SCANCODE_F18
	K_F19                        = K_SCANCODE_MASK | SCANCODE_F19
	K_F20                        = K_SCANCODE_MASK | SCANCODE_F20
	K_F21                        = K_SCANCODE_MASK | SCANCODE_F21
	K_F22                        = K_SCANCODE_MASK | SCANCODE_F22
	K_F23                        = K_SCANCODE_MASK | SCANCODE_F23
	K_F24                        = K_SCANCODE_MASK | SCANCODE_F24
	K_EXECUTE                    = K_SCANCODE_MASK | SCANCODE_EXECUTE
	K_HELP                       = K_SCANCODE_MASK | SCANCODE_HELP
	K_MENU                       = K_SCANCODE_MASK | SCANCODE_MENU
	K_SELECT                     = K_SCANCODE_MASK | SCANCODE_SELECT
	K_STOP                       = K_SCANCODE_MASK | SCANCODE_STOP
	K_AGAIN                      = K_SCANCODE_MASK | SCANCODE_AGAIN
	K_UNDO                       = K_SCANCODE_MASK | SCANCODE_UNDO
	K_CUT                        = K_SCANCODE_MASK | SCANCODE_CUT
	K_COPY                       = K_SCANCODE_MASK | SCANCODE_COPY
	K_PASTE                      = K_SCANCODE_MASK | SCANCODE_PASTE
	K_FIND                       = K_SCANCODE_MASK | SCANCODE_FIND
	K_MUTE                       = K_SCANCODE_MASK | SCANCODE_MUTE
	K_VOLUMEUP                   = K_SCANCODE_MASK | SCANCODE_VOLUMEUP
	K_VOLUMEDOWN                 = K_SCANCODE_MASK | SCANCODE_VOLUMEDOWN
	K_KP_COMMA                   = K_SCANCODE_MASK | SCANCODE_KP_COMMA
	K_KP_EQUALSAS400             = K_SCANCODE_MASK | SCANCODE_KP_EQUALSAS400
	K_ALTERASE                   = K_SCANCODE_MASK | SCANCODE_ALTERASE
	K_SYSREQ                     = K_SCANCODE_MASK | SCANCODE_SYSREQ
	K_CANCEL                     = K_SCANCODE_MASK | SCANCODE_CANCEL
	K_CLEAR                      = K_SCANCODE_MASK | SCANCODE_CLEAR
	K_PRIOR                      = K_SCANCODE_MASK | SCANCODE_PRIOR
	K_RETURN2                    = K_SCANCODE_MASK | SCANCODE_RETURN2
	K_SEPARATOR                  = K_SCANCODE_MASK | SCANCODE_SEPARATOR
	K_OUT                        = K_SCANCODE_MASK | SCANCODE_OUT
	K_OPER                       = K_SCANCODE_MASK | SCANCODE_OPER
	K_CLEARAGAIN                 = K_SCANCODE_MASK | SCANCODE_CLEARAGAIN
	K_CRSEL                      = K_SCANCODE_MASK | SCANCODE_CRSEL
	K_EXSEL                      = K_SCANCODE_MASK | SCANCODE_EXSEL
	K_KP_00                      = K_SCANCODE_MASK | SCANCODE_KP_00
	K_KP_000                     = K_SCANCODE_MASK | SCANCODE_KP_000
	K_THOUSANDSSEPARATOR         = K_SCANCODE_MASK | SCANCODE_THOUSANDSSEPARATOR
	K_DECIMALSEPARATOR           = K_SCANCODE_MASK | SCANCODE_DECIMALSEPARATOR
	K_CURRENCYUNIT               = K_SCANCODE_MASK | SCANCODE_CURRENCYUNIT
	K_CURRENCYSUBUNIT            = K_SCANCODE_MASK | SCANCODE_CURRENCYSUBUNIT
	K_KP_LEFTPAREN               = K_SCANCODE_MASK | SCANCODE_KP_LEFTPAREN
	K_KP_RIGHTPAREN              = K_SCANCODE_MASK | SCANCODE_KP_RIGHTPAREN
	K_KP_LEFTBRACE               = K_SCANCODE_MASK | SCANCODE_KP_LEFTBRACE
	K_KP_RIGHTBRACE              = K_SCANCODE_MASK | SCANCODE_KP_RIGHTBRACE
	K_KP_TAB                     = K_SCANCODE_MASK | SCANCODE_KP_TAB
	K_KP_BACKSPACE               = K_SCANCODE_MASK | SCANCODE_KP_BACKSPACE
	K_KP_A                       = K_SCANCODE_MASK | SCANCODE_KP_A
	K_KP_B                       = K_SCANCODE_MASK | SCANCODE_KP_B
	K_KP_C                       = K_SCANCODE_MASK | SCANCODE_KP_C
	K_KP_D                       = K_SCANCODE_MASK | SCANCODE_KP_D
	K_KP_E                       = K_SCANCODE_MASK | SCANCODE_KP_E
	K_KP_F                       = K_SCANCODE_MASK | SCANCODE_KP_F
	K_KP_XOR                     = K_SCANCODE_MASK | SCANCODE_KP_XOR
	K_KP_POWER                   = K_SCANCODE_MASK | SCANCODE_KP_POWER
	K_KP_PERCENT                 = K_SCANCODE_MASK | SCANCODE_KP_PERCENT
	K_KP_LESS                    = K_SCANCODE_MASK | SCANCODE_KP_LESS
	K_KP_GREATER                 = K_SCANCODE_MASK | SCANCODE_KP_GREATER
	K_KP_AMPERSAND               = K_SCANCODE_MASK | SCANCODE_KP_AMPERSAND
	K_KP_DBLAMPERSAND            = K_SCANCODE_MASK | SCANCODE_KP_DBLAMPERSAND
	K_KP_VERTICALBAR             = K_SCANCODE_MASK | SCANCODE_KP_VERTICALBAR
	K_KP_DBLVERTICALBAR          = K_SCANCODE_MASK | SCANCODE_KP_DBLVERTICALBAR
	K_KP_COLON                   = K_SCANCODE_MASK | SCANCODE_KP_COLON
	K_KP_HASH                    = K_SCANCODE_MASK | SCANCODE_KP_HASH
	K_KP_SPACE                   = K_SCANCODE_MASK | SCANCODE_KP_SPACE
	K_KP_AT                      = K_SCANCODE_MASK | SCANCODE_KP_AT
	K_KP_EXCLAM                  = K_SCANCODE_MASK | SCANCODE_KP_EXCLAM
	K_KP_MEMSTORE                = K_SCANCODE_MASK | SCANCODE_KP_MEMSTORE
	K_KP_MEMRECALL               = K_SCANCODE_MASK | SCANCODE_KP_MEMRECALL
	K_KP_MEMCLEAR                = K_SCANCODE_MASK | SCANCODE_KP_MEMCLEAR
	K_KP_MEMADD                  = K_SCANCODE_MASK | SCANCODE_KP_MEMADD
	K_KP_MEMSUBTRACT             = K_SCANCODE_MASK | SCANCODE_KP_MEMSUBTRACT
	K_KP_MEMMULTIPLY             = K_SCANCODE_MASK | SCANCODE_KP_MEMMULTIPLY
	K_KP_MEMDIVIDE               = K_SCANCODE_MASK | SCANCODE_KP_MEMDIVIDE
	K_KP_PLUSMINUS               = K_SCANCODE_MASK | SCANCODE_KP_PLUSMINUS
	K_KP_CLEAR                   = K_SCANCODE_MASK | SCANCODE_KP_CLEAR
	K_KP_CLEARENTRY              = K_SCANCODE_MASK | SCANCODE_KP_CLEARENTRY
	K_KP_BINARY                  = K_SCANCODE_MASK | SCANCODE_KP_BINARY
	K_KP_OCTAL                   = K_SCANCODE_MASK | SCANCODE_KP_OCTAL
	K_KP_DECIMAL                 = K_SCANCODE_MASK | SCANCODE_KP_DECIMAL
	K_KP_HEXADECIMAL             = K_SCANCODE_MASK | SCANCODE_KP_HEXADECIMAL
	K_LCTRL                      = K_SCANCODE_MASK | SCANCODE_LCTRL
	K_LSHIFT                     = K_SCANCODE_MASK | SCANCODE_LSHIFT
	K_LALT                       = K_SCANCODE_MASK | SCANCODE_LALT
	K_LGUI                       = K_SCANCODE_MASK | SCANCODE_LGUI
	K_RCTRL                      = K_SCANCODE_MASK | SCANCODE_RCTRL
	K_RSHIFT                     = K_SCANCODE_MASK | SCANCODE_RSHIFT
	K_RALT                       = K_SCANCODE_MASK | SCANCODE_RALT
	K_RGUI                       = K_SCANCODE_MASK | SCANCODE_RGUI
	K_MODE                       = K_SCANCODE_MASK | SCANCODE_MODE
	K_AUDIONEXT                  = K_SCANCODE_MASK | SCANCODE_AUDIONEXT
	K_AUDIOPREV                  = K_SCANCODE_MASK | SCANCODE_AUDIOPREV
	K_AUDIOSTOP                  = K_SCANCODE_MASK | SCANCODE_AUDIOSTOP
	K_AUDIOPLAY                  = K_SCANCODE_MASK | SCANCODE_AUDIOPLAY
	K_AUDIOMUTE                  = K_SCANCODE_MASK | SCANCODE_AUDIOMUTE
	K_MEDIASELECT                = K_SCANCODE_MASK | SCANCODE_MEDIASELECT
	K_WWW                        = K_SCANCODE_MASK | SCANCODE_WWW
	K_MAIL                       = K_SCANCODE_MASK | SCANCODE_MAIL
	K_CALCULATOR                 = K_SCANCODE_MASK | SCANCODE_CALCULATOR
	K_COMPUTER                   = K_SCANCODE_MASK | SCANCODE_COMPUTER
	K_AC_SEARCH                  = K_SCANCODE_MASK | SCANCODE_AC_SEARCH
	K_AC_HOME                    = K_SCANCODE_MASK | SCANCODE_AC_HOME
	K_AC_BACK                    = K_SCANCODE_MASK | SCANCODE_AC_BACK
	K_AC_FORWARD                 = K_SCANCODE_MASK | SCANCODE_AC_FORWARD
	K_AC_STOP                    = K_SCANCODE_MASK | SCANCODE_AC_STOP
	K_AC_REFRESH                 = K_SCANCODE_MASK | SCANCODE_AC_REFRESH
	K_AC_BOOKMARKS               = K_SCANCODE_MASK | SCANCODE_AC_BOOKMARKS
	K_BRIGHTNESSDOWN             = K_SCANCODE_MASK | SCANCODE_BRIGHTNESSDOWN
	K_BRIGHTNESSUP               = K_SCANCODE_MASK | SCANCODE_BRIGHTNESSUP
	K_DISPLAYSWITCH              = K_SCANCODE_MASK | SCANCODE_DISPLAYSWITCH
	K_KBDILLUMTOGGLE             = K_SCANCODE_MASK | SCANCODE_KBDILLUMTOGGLE
	K_KBDILLUMDOWN               = K_SCANCODE_MASK | SCANCODE_KBDILLUMDOWN
	K_KBDILLUMUP                 = K_SCANCODE_MASK | SCANCODE_KBDILLUMUP
	K_EJECT                      = K_SCANCODE_MASK | SCANCODE_EJECT
	K_SLEEP                      = K_SCANCODE_MASK | SCANCODE_SLEEP
	K_APP1                       = K_SCANCODE_MASK | SCANCODE_APP1
	K_APP2                       = K_SCANCODE_MASK | SCANCODE_APP2
	K_AUDIOREWIND                = K_SCANCODE_MASK | SCANCODE_AUDIOREWIND
	K_AUDIOFASTFORWARD           = K_SCANCODE_MASK | SCANCODE_AUDIOFASTFORWARD
)

const (
	SYSWM_UNKNOWN SYSWM_TYPE = iota
	SYSWM_WINDOWS
	SYSWM_X11
	SYSWM_DIRECTFB
	SYSWM_COCOA
	SYSWM_UIKIT
	SYSWM_WAYLAND
	SYSWM_MIR
	SYSWM_WINRT
	SYSWM_ANDROID
	SYSWM_VIVANTE
	SYSWM_OS2
)

const (
	FIRSTEVENT               EventType = 0
	QUIT                     EventType = 0x100
	APP_TERMINATING          EventType = 0x101
	APP_LOWMEMORY            EventType = 0x102
	APP_WILLENTERBACKGROUND  EventType = 0x103
	APP_DIDENTERBACKGROUND   EventType = 0x104
	APP_WILLENTERFOREGROUND  EventType = 0x105
	APP_DIDENTERFOREGROUND   EventType = 0x106
	WINDOWEVENT              EventType = 0x200
	SYSWMEVENT               EventType = 0x201
	KEYDOWN                  EventType = 0x300
	KEYUP                    EventType = 0x301
	TEXTEDITING              EventType = 0x302
	TEXTINPUT                EventType = 0x303
	KEYMAPCHANGED            EventType = 0x304
	MOUSEMOTION              EventType = 0x400
	MOUSEBUTTONDOWN          EventType = 0x401
	MOUSEBUTTONUP            EventType = 0x402
	MOUSEWHEEL               EventType = 0x403
	JOYAXISMOTION            EventType = 0x600
	JOYBALLMOTION            EventType = 0x601
	JOYHATMOTION             EventType = 0x602
	JOYBUTTONDOWN            EventType = 0x603
	JOYBUTTONUP              EventType = 0x604
	JOYDEVICEADDED           EventType = 0x605
	JOYDEVICEREMOVED         EventType = 0x606
	CONTROLLERAXISMOTION     EventType = 0x650
	CONTROLLERBUTTONDOWN     EventType = 0x651
	CONTROLLERBUTTONUP       EventType = 0x652
	CONTROLLERDEVICEADDED    EventType = 0x653
	CONTROLLERDEVICEREMOVED  EventType = 0x654
	CONTROLLERDEVICEREMAPPED EventType = 0x655
	FINGERDOWN               EventType = 0x700
	FINGERUP                 EventType = 0x701
	FINGERMOTION             EventType = 0x702
	DOLLARGESTURE            EventType = 0x800
	DOLLARRECORD             EventType = 0x801
	MULTIGESTURE             EventType = 0x802
	CLIPBOARDUPDATE          EventType = 0x900
	DROPFILE                 EventType = 0x1000
	DROPTEXT                 EventType = 0x1001
	DROPBEGIN                EventType = 0x1002
	DROPCOMPLETE             EventType = 0x1003
	AUDIODEVICEADDED         EventType = 0x1100
	AUDIODEVICEREMOVED       EventType = 0x1101
	RENDER_TARGETS_RESET     EventType = 0x2000
	RENDER_DEVICE_RESET      EventType = 0x2001
	USEREVENT                EventType = 0x8000
	LASTEVENT                EventType = 0xFFFF
)

// ShowCursor parameters
const (
	QUERY   = -1
	DISABLE = 0
	ENABLE  = 1
)

const (
	WINDOW_FULLSCREEN         = 0x00000001
	WINDOW_OPENGL             = 0x00000002
	WINDOW_SHOWN              = 0x00000004
	WINDOW_HIDDEN             = 0x00000008
	WINDOW_BORDERLESS         = 0x00000010
	WINDOW_RESIZABLE          = 0x00000020
	WINDOW_MINIMIZED          = 0x00000040
	WINDOW_MAXIMIZED          = 0x00000080
	WINDOW_INPUT_GRABBED      = 0x00000100
	WINDOW_INPUT_FOCUS        = 0x00000200
	WINDOW_MOUSE_FOCUS        = 0x00000400
	WINDOW_FULLSCREEN_DESKTOP = (WINDOW_FULLSCREEN | 0x00001000)
	WINDOW_FOREIGN            = 0x00000800
	WINDOW_ALLOW_HIGHDPI      = 0x00002000
	WINDOW_MOUSE_CAPTURE      = 0x00004000
	WINDOW_ALWAYS_ON_TOP      = 0x00008000
	WINDOW_SKIP_TASKBAR       = 0x00010000
	WINDOW_UTILITY            = 0x00020000
	WINDOW_TOOLTIP            = 0x00040000
	WINDOW_POPUP_MENU         = 0x00080000
	WINDOW_VULKAN             = 0x10000000
)

const (
	HINT_FRAMEBUFFER_ACCELERATION                 = "SDL_FRAMEBUFFER_ACCELERATION"
	HINT_RENDER_DRIVER                            = "SDL_RENDER_DRIVER"
	HINT_RENDER_OPENGL_SHADERS                    = "SDL_RENDER_OPENGL_SHADERS"
	HINT_RENDER_DIRECT3D_THREADSAFE               = "SDL_RENDER_DIRECT3D_THREADSAFE"
	HINT_RENDER_DIRECT3D11_DEBUG                  = "SDL_RENDER_DIRECT3D11_DEBUG"
	HINT_RENDER_LOGICAL_SIZE_MODE                 = "SDL_RENDER_LOGICAL_SIZE_MODE"
	HINT_RENDER_SCALE_QUALITY                     = "SDL_RENDER_SCALE_QUALITY"
	HINT_RENDER_VSYNC                             = "SDL_RENDER_VSYNC"
	HINT_VIDEO_ALLOW_SCREENSAVER                  = "SDL_VIDEO_ALLOW_SCREENSAVER"
	HINT_VIDEO_X11_XVIDMODE                       = "SDL_VIDEO_X11_XVIDMODE"
	HINT_VIDEO_X11_XINERAMA                       = "SDL_VIDEO_X11_XINERAMA"
	HINT_VIDEO_X11_XRANDR                         = "SDL_VIDEO_X11_XRANDR"
	HINT_VIDEO_X11_NET_WM_PING                    = "SDL_VIDEO_X11_NET_WM_PING"
	HINT_VIDEO_X11_NET_WM_BYPASS_COMPOSITOR       = "SDL_VIDEO_X11_NET_WM_BYPASS_COMPOSITOR"
	HINT_WINDOW_FRAME_USABLE_WHILE_CURSOR_HIDDEN  = "SDL_WINDOW_FRAME_USABLE_WHILE_CURSOR_HIDDEN"
	HINT_WINDOWS_INTRESOURCE_ICON                 = "SDL_WINDOWS_INTRESOURCE_ICON"
	HINT_WINDOWS_INTRESOURCE_ICON_SMALL           = "SDL_WINDOWS_INTRESOURCE_ICON_SMALL"
	HINT_WINDOWS_ENABLE_MESSAGELOOP               = "SDL_WINDOWS_ENABLE_MESSAGELOOP"
	HINT_GRAB_KEYBOARD                            = "SDL_GRAB_KEYBOARD"
	HINT_MOUSE_NORMAL_SPEED_SCALE                 = "SDL_MOUSE_NORMAL_SPEED_SCALE"
	HINT_MOUSE_RELATIVE_SPEED_SCALE               = "SDL_MOUSE_RELATIVE_SPEED_SCALE"
	HINT_MOUSE_RELATIVE_MODE_WARP                 = "SDL_MOUSE_RELATIVE_MODE_WARP"
	HINT_MOUSE_FOCUS_CLICKTHROUGH                 = "SDL_MOUSE_FOCUS_CLICKTHROUGH"
	HINT_TOUCH_MOUSE_EVENTS                       = "SDL_TOUCH_MOUSE_EVENTS"
	HINT_VIDEO_MINIMIZE_ON_FOCUS_LOSS             = "SDL_VIDEO_MINIMIZE_ON_FOCUS_LOSS"
	HINT_IDLE_TIMER_DISABLED                      = "SDL_IOS_IDLE_TIMER_DISABLED"
	HINT_ORIENTATIONS                             = "SDL_IOS_ORIENTATIONS"
	HINT_APPLE_TV_CONTROLLER_UI_EVENTS            = "SDL_APPLE_TV_CONTROLLER_UI_EVENTS"
	HINT_APPLE_TV_REMOTE_ALLOW_ROTATION           = "SDL_APPLE_TV_REMOTE_ALLOW_ROTATION"
	HINT_IOS_HIDE_HOME_INDICATOR                  = "SDL_IOS_HIDE_HOME_INDICATOR"
	HINT_ACCELEROMETER_AS_JOYSTICK                = "SDL_ACCELEROMETER_AS_JOYSTICK"
	HINT_TV_REMOTE_AS_JOYSTICK                    = "SDL_TV_REMOTE_AS_JOYSTICK"
	HINT_XINPUT_ENABLED                           = "SDL_XINPUT_ENABLED"
	HINT_XINPUT_USE_OLD_JOYSTICK_MAPPING          = "SDL_XINPUT_USE_OLD_JOYSTICK_MAPPING"
	HINT_GAMECONTROLLERCONFIG                     = "SDL_GAMECONTROLLERCONFIG"
	HINT_GAMECONTROLLER_IGNORE_DEVICES            = "SDL_GAMECONTROLLER_IGNORE_DEVICES"
	HINT_GAMECONTROLLER_IGNORE_DEVICES_EXCEPT     = "SDL_GAMECONTROLLER_IGNORE_DEVICES_EXCEPT"
	HINT_JOYSTICK_ALLOW_BACKGROUND_EVENTS         = "SDL_JOYSTICK_ALLOW_BACKGROUND_EVENTS"
	HINT_ALLOW_TOPMOST                            = "SDL_ALLOW_TOPMOST"
	HINT_TIMER_RESOLUTION                         = "SDL_TIMER_RESOLUTION"
	HINT_QTWAYLAND_CONTENT_ORIENTATION            = "SDL_QTWAYLAND_CONTENT_ORIENTATION"
	HINT_QTWAYLAND_WINDOW_FLAGS                   = "SDL_QTWAYLAND_WINDOW_FLAGS"
	HINT_THREAD_STACK_SIZE                        = "SDL_THREAD_STACK_SIZE"
	HINT_VIDEO_HIGHDPI_DISABLED                   = "SDL_VIDEO_HIGHDPI_DISABLED"
	HINT_MAC_CTRL_CLICK_EMULATE_RIGHT_CLICK       = "SDL_MAC_CTRL_CLICK_EMULATE_RIGHT_CLICK"
	HINT_VIDEO_WIN_D3DCOMPILER                    = "SDL_VIDEO_WIN_D3DCOMPILER"
	HINT_VIDEO_WINDOW_SHARE_PIXEL_FORMAT          = "SDL_VIDEO_WINDOW_SHARE_PIXEL_FORMAT"
	HINT_WINRT_PRIVACY_POLICY_URL                 = "SDL_WINRT_PRIVACY_POLICY_URL"
	HINT_WINRT_PRIVACY_POLICY_LABEL               = "SDL_WINRT_PRIVACY_POLICY_LABEL"
	HINT_WINRT_HANDLE_BACK_BUTTON                 = "SDL_WINRT_HANDLE_BACK_BUTTON"
	HINT_VIDEO_MAC_FULLSCREEN_SPACES              = "SDL_VIDEO_MAC_FULLSCREEN_SPACES"
	HINT_MAC_BACKGROUND_APP                       = "SDL_MAC_BACKGROUND_APP"
	HINT_ANDROID_APK_EXPANSION_MAIN_FILE_VERSION  = "SDL_ANDROID_APK_EXPANSION_MAIN_FILE_VERSION"
	HINT_ANDROID_APK_EXPANSION_PATCH_FILE_VERSION = "SDL_ANDROID_APK_EXPANSION_PATCH_FILE_VERSION"
	HINT_IME_INTERNAL_EDITING                     = "SDL_IME_INTERNAL_EDITING"
	HINT_ANDROID_SEPARATE_MOUSE_AND_TOUCH         = "SDL_ANDROID_SEPARATE_MOUSE_AND_TOUCH"
	HINT_RETURN_KEY_HIDES_IME                     = "SDL_RETURN_KEY_HIDES_IME"
	HINT_EMSCRIPTEN_KEYBOARD_ELEMENT              = "SDL_EMSCRIPTEN_KEYBOARD_ELEMENT"
	HINT_NO_SIGNAL_HANDLERS                       = "SDL_NO_SIGNAL_HANDLERS"
	HINT_WINDOWS_NO_CLOSE_ON_ALT_F4               = "SDL_WINDOWS_NO_CLOSE_ON_ALT_F4"
	HINT_BMP_SAVE_LEGACY_FORMAT                   = "SDL_BMP_SAVE_LEGACY_FORMAT"
	HINT_WINDOWS_DISABLE_THREAD_NAMING            = "SDL_WINDOWS_DISABLE_THREAD_NAMING"
	HINT_RPI_VIDEO_LAYER                          = "SDL_RPI_VIDEO_LAYER"
	HINT_VIDEO_DOUBLE_BUFFER                      = "SDL_VIDEO_DOUBLE_BUFFER"
	HINT_OPENGL_ES_DRIVER                         = "SDL_OPENGL_ES_DRIVER"
	HINT_AUDIO_RESAMPLING_MODE                    = "SDL_AUDIO_RESAMPLING_MODE"
	HINT_AUDIO_CATEGORY                           = "SDL_AUDIO_CATEGORY"
)

const (
	ASSERTION_RETRY = iota
	ASSERTION_BREAK
	ASSERTION_ABORT
	ASSERTION_IGNORE
	ASSERTION_ALWAYS_IGNORE
)

const (
	BLENDFACTOR_ZERO                = 0x1
	BLENDFACTOR_ONE                 = 0x2
	BLENDFACTOR_SRC_COLOR           = 0x3
	BLENDFACTOR_ONE_MINUS_SRC_COLOR = 0x4
	BLENDFACTOR_SRC_ALPHA           = 0x5
	BLENDFACTOR_ONE_MINUS_SRC_ALPHA = 0x6
	BLENDFACTOR_DST_COLOR           = 0x7
	BLENDFACTOR_ONE_MINUS_DST_COLOR = 0x8
	BLENDFACTOR_DST_ALPHA           = 0x9
	BLENDFACTOR_ONE_MINUS_DST_ALPHA = 0xA
)

const (
	BLENDMODE_NONE    = 0x00000000
	BLENDMODE_BLEND   = 0x00000001
	BLENDMODE_ADD     = 0x00000002
	BLENDMODE_MOD     = 0x00000004
	BLENDMODE_INVALID = 0x7FFFFFFF
)

const (
	BLENDOPERATION_ADD          = 0x1
	BLENDOPERATION_SUBTRACT     = 0x2
	BLENDOPERATION_REV_SUBTRACT = 0x3
	BLENDOPERATION_MINIMUM      = 0x4
	BLENDOPERATION_MAXIMUM      = 0x5
)

const (
	GL_RED_SIZE = iota
	GL_GREEN_SIZE
	GL_BLUE_SIZE
	GL_ALPHA_SIZE
	GL_BUFFER_SIZE
	GL_DOUBLEBUFFER
	GL_DEPTH_SIZE
	GL_STENCIL_SIZE
	GL_ACCUM_RED_SIZE
	GL_ACCUM_GREEN_SIZE
	GL_ACCUM_BLUE_SIZE
	GL_ACCUM_ALPHA_SIZE
	GL_STEREO
	GL_MULTISAMPLEBUFFERS
	GL_MULTISAMPLESAMPLES
	GL_ACCELERATED_VISUAL
	GL_RETAINED_BACKING
	GL_CONTEXT_MAJOR_VERSION
	GL_CONTEXT_MINOR_VERSION
	GL_CONTEXT_EGL
	GL_CONTEXT_FLAGS
	GL_CONTEXT_PROFILE_MASK
	GL_SHARE_WITH_CURRENT_CONTEXT
	GL_FRAMEBUFFER_SRGB_CAPABLE
	GL_CONTEXT_RELEASE_BEHAVIOR
	GL_CONTEXT_RESET_NOTIFICATION
	GL_CONTEXT_NO_ERROR
)

const (
	GL_CONTEXT_DEBUG_FLAG              = 0x0001
	GL_CONTEXT_FORWARD_COMPATIBLE_FLAG = 0x0002
	GL_CONTEXT_ROBUST_ACCESS_FLAG      = 0x0004
	GL_CONTEXT_RESET_ISOLATION_FLAG    = 0x0008
)

const (
	GL_CONTEXT_PROFILE_CORE          = 0x0001
	GL_CONTEXT_PROFILE_COMPATIBILITY = 0x0002
	GL_CONTEXT_PROFILE_ES            = 0x0004
)

const (
	CONTROLLER_AXIS_INVALID = iota - 1
	CONTROLLER_AXIS_LEFTX
	CONTROLLER_AXIS_LEFTY
	CONTROLLER_AXIS_RIGHTX
	CONTROLLER_AXIS_RIGHTY
	CONTROLLER_AXIS_TRIGGERLEFT
	CONTROLLER_AXIS_TRIGGERRIGHT
	CONTROLLER_AXIS_MAX
)

const (
	CONTROLLER_BUTTON_INVALID = iota - 1
	CONTROLLER_BUTTON_A
	CONTROLLER_BUTTON_B
	CONTROLLER_BUTTON_X
	CONTROLLER_BUTTON_Y
	CONTROLLER_BUTTON_BACK
	CONTROLLER_BUTTON_GUIDE
	CONTROLLER_BUTTON_START
	CONTROLLER_BUTTON_LEFTSTICK
	CONTROLLER_BUTTON_RIGHTSTICK
	CONTROLLER_BUTTON_LEFTSHOULDER
	CONTROLLER_BUTTON_RIGHTSHOULDER
	CONTROLLER_BUTTON_DPAD_UP
	CONTROLLER_BUTTON_DPAD_DOWN
	CONTROLLER_BUTTON_DPAD_LEFT
	CONTROLLER_BUTTON_DPAD_RIGHT
	CONTROLLER_BUTTON_MAX
)

const (
	HINT_DEFAULT = iota
	HINT_NORMAL
	HINT_OVERRIDE
)

const (
	HITTEST_NORMAL = iota
	HITTEST_DRAGGABLE
	HITTEST_RESIZE_TOPLEFT
	HITTEST_RESIZE_TOP
	HITTEST_RESIZE_TOPRIGHT
	HITTEST_RESIZE_RIGHT
	HITTEST_RESIZE_BOTTOMRIGHT
	HITTEST_RESIZE_BOTTOM
	HITTEST_RESIZE_BOTTOMLEFT
	HITTEST_RESIZE_LEFT
)

const (
	JOYSTICK_POWER_UNKNOWN = iota - 1
	JOYSTICK_POWER_EMPTY
	JOYSTICK_POWER_LOW
	JOYSTICK_POWER_MEDIUM
	JOYSTICK_POWER_FULL
	JOYSTICK_POWER_WIRED
	JOYSTICK_POWER_MAX
)

const (
	KMOD_NONE     = 0x0000
	KMOD_LSHIFT   = 0x0001
	KMOD_RSHIFT   = 0x0002
	KMOD_LCTRL    = 0x0040
	KMOD_RCTRL    = 0x0080
	KMOD_LALT     = 0x0100
	KMOD_RALT     = 0x0200
	KMOD_LGUI     = 0x0400
	KMOD_RGUI     = 0x0800
	KMOD_NUM      = 0x1000
	KMOD_CAPS     = 0x2000
	KMOD_MODE     = 0x4000
	KMOD_RESERVED = 0x8000
)

const (
	LOG_CATEGORY_APPLICATION = iota
	LOG_CATEGORY_ERROR
	LOG_CATEGORY_ASSERT
	LOG_CATEGORY_SYSTEM
	LOG_CATEGORY_AUDIO
	LOG_CATEGORY_VIDEO
	LOG_CATEGORY_RENDER
	LOG_CATEGORY_INPUT
	LOG_CATEGORY_TEST

	// Reserved for future SDL library use
	LOG_CATEGORY_RESERVED1
	LOG_CATEGORY_RESERVED2
	LOG_CATEGORY_RESERVED3
	LOG_CATEGORY_RESERVED4
	LOG_CATEGORY_RESERVED5
	LOG_CATEGORY_RESERVED6
	LOG_CATEGORY_RESERVED7
	LOG_CATEGORY_RESERVED8
	LOG_CATEGORY_RESERVED9
	LOG_CATEGORY_RESERVED10

	// Beyond this point is reserved for application use, e.g.
	// const (
	//     MYAPP_CATEGORY_AWESOME1 = iota + SDL_LOG_CATEGORY_CUSTOM
	//     MYAPP_CATEGORY_AWESOME2
	//     MYAPP_CATEGORY_AWESOME3
	//     ...
	// )
	LOG_CATEGORY_CUSTOM
)

const (
	LOG_PRIORITY_VERBOSE = iota + 1
	LOG_PRIORITY_DEBUG
	LOG_PRIORITY_INFO
	LOG_PRIORITY_WARN
	LOG_PRIORITY_ERROR
	LOG_PRIORITY_CRITICAL
	NUM_LOG_PRIORITIES
)

const (
	MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT = 0x00000001
	MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT = 0x00000002
)

const (
	MESSAGEBOX_COLOR_BACKGROUND = iota
	MESSAGEBOX_COLOR_TEXT
	MESSAGEBOX_COLOR_BUTTON_BORDER
	MESSAGEBOX_COLOR_BUTTON_BACKGROUND
	MESSAGEBOX_COLOR_BUTTON_SELECTED
	MESSAGEBOX_COLOR_MAX
)

const (
	MESSAGEBOX_ERROR       = 0x00000010
	MESSAGEBOX_WARNING     = 0x00000020
	MESSAGEBOX_INFORMATION = 0x00000040
)

const (
	POWERSTATE_UNKNOWN = iota
	POWERSTATE_ON_BATTERY
	POWERSTATE_NO_BATTERY
	POWERSTATE_CHARGING
	POWERSTATE_CHARGED
)

const (
	RENDERER_SOFTWARE      = 0x00000001
	RENDERER_ACCELERATED   = 0x00000002
	RENDERER_PRESENTVSYNC  = 0x00000004
	RENDERER_TARGETTEXTURE = 0x00000008
)

const (
	FLIP_NONE       = 0x00000000
	FLIP_HORIZONTAL = 0x00000001
	FLIP_VERTICAL   = 0x00000002
)

const (
	TEXTUREACCESS_STATIC = iota
	TEXTUREACCESS_STREAMING
	TEXTUREACCESS_TARGET
)

const (
	TEXTUREMODULATE_NONE  = 0x00000000
	TEXTUREMODULATE_COLOR = 0x00000001
	TEXTUREMODULATE_ALPHA = 0x00000002
)

const (
	THREAD_PRIORITY_LOW = iota
	THREAD_PRIORITY_NORMAL
	THREAD_PRIORITY_HIGH
)

const (
	WINRT_PATH_INSTALLED_LOCATION = iota
	WINRT_PATH_LOCAL_FOLDER
	WINRT_PATH_ROAMING_FOLDER
	WINRT_PATH_TEMP_FOLDER
)

const (
	WINDOWEVENT_NONE = iota
	WINDOWEVENT_SHOWN
	WINDOWEVENT_HIDDEN
	WINDOWEVENT_EXPOSED
	WINDOWEVENT_MOVED
	WINDOWEVENT_RESIZED
	WINDOWEVENT_SIZE_CHANGED
	WINDOWEVENT_MINIMIZED
	WINDOWEVENT_MAXIMIZED
	WINDOWEVENT_RESTORED
	WINDOWEVENT_ENTER
	WINDOWEVENT_LEAVE
	WINDOWEVENT_FOCUS_GAINED
	WINDOWEVENT_FOCUS_LOST
	WINDOWEVENT_CLOSE
	WINDOWEVENT_TAKE_FOCUS
	WINDOWEVENT_HIT_TEST
)
