package ui

// TODO: Can start putting them here and break them off and move them to
// pacakges from here
var Symbols = map[string]map[string]string{
	"status": map[string]string{
		"info":    "ℹ ",
		"success": "✔ ",
		"warning": "⚠ ",
		"error":   "✖ ",
	},
	"emoticons": map[string]string{
		"sad":   "☹",
		"happy": "☺",
	},
	"shogi": map[string]string{
		"white": "☖",
		"black": "☗",
	},
	"diamond_checkbox": map[string]string{
		"false": "◇",
		"true":  "◈",
		"solid": "◆",
	},
	"radio": map[string]string{
		"true":   "◉",
		"false":  "◯",
		"filled": "●",
		"dotted": "◌",
	},
	"checkbox": map[string]string{
		"xed":     "☒",
		"rounded": "▢",
		"checked": "☑",
		"filled":  "◼",
		"false":   "☐",
	},
	"arrows": map[string]string{
		"up":    "↑",
		"down":  "↓",
		"left":  "←",
		"right": "→",
	},
	"fraction": map[string]string{
		"one_half":       "½",
		"one_third":      "⅓",
		"one_quarter":    "¼",
		"one_fifth":      "⅕",
		"one_sixth":      "⅙",
		"one_seventh":    "⅐",
		"one_eighth":     "⅛",
		"one_ninth":      "⅑",
		"one_tenth":      "⅒",
		"two_thirds":     "⅔",
		"two_fifths":     "⅖",
		"three_quarters": "¾",
		"three_fifths":   "⅗",
		"three_eighths":  "⅜",
		"four_fifths":    "⅘",
		"five_sixths":    "⅚",
		"five_eighths":   "⅝",
		"seven_eighths":  "⅞",
	},
	"player": map[string]string{
		"stop":                "◼",
		"play":                "▶",
		"repeat":              "🔁",
		"repeat_once":         "🔂",
		"random":              "🔀",
		"mute":                "🔇",
		"mute_volume":         "🔈",
		"decrease_volume":     "🔉",
		"increase_volume":     "🔊 ",
		"decrease_brightness": "🔅",
		"increase_brightness": "🔆",
	},
	"math": map[string]string{
		"pi":       "π ",
		"omega":    "Ω ",
		"theta":    "Θ ",
		"beta":     "β ",
		"delta":    "δ ",
		"multiply": "✕",
	},
	"star": map[string]string{
		"filled": "★",
		"empty":  "☆",
	},
	"flag": map[string]string{
		"filled": "⚑",
		"empty":  "⚐",
	},
	"heart": map[string]string{
		"filled": "♥",
		"empty":  "♡",
	},
	"calendar": map[string]string{
		"spiral":   "🗓",
		"standard": "📅",
		"tear_off": "📆",
	},
	"security": map[string]string{
		"key":          "🔑",
		"lock_and_key": "🔐",
		"lock_closed":  "🔒",
		"lock_open":    "🔓",
	},
	"filesystem": map[string]string{
		"directory":      "📁",
		"directory_open": "📂",
		"editor":         "📝",
		"label":          "🏷",
	},
	"prompt": map[string]string{
		"menu":                     "☰",
		"rook_filled":              "♜",
		"root_empty":               "♖",
		"notes":                    "♫",
		"flag_empty":               "⚐",
		"flag_filled":              "⚑",
		"voltage":                  "⚡",
		"wavy_arrow":               "↝",
		"tailed_arrow":             "↣",
		"barred_arrow":             "↦",
		"hooked_arrow":             "↪",
		"angled_arrow":             "↳",
		"curved_arrow":             "↺",
		"downward_harpoon":         "⇁",
		"thin_filled_arrow_point":  " ►",
		"thin_empty_arrow_point":   "▻",
		"filled_arrow_point":       "▶",
		"empty_arrow_point":        "▷",
		"small_filled_arrow_point": "▸",
		"small_empty_arrow_point":  "▹",
		"filled_square":            "◼",
		"empty_square":             "◻",
		"sum":                      "∑",
		"for_all":                  "∀",
		"complement":               "∁",
		"empty_set":                "∅",
		"element_of":               "∈",
		"not_element_of":           "∉",
		"small_element_of":         "∊",
		"end_of_proof":             "∎",
		"plus_minux":               "∓",
		"plus_dot":                 "∔",
		"asterisk":                 "∗",
		"bullet_filled":            "∙",
		"bullet_empty":             "∘",
		"infinity":                 "∞",
		"measured_angle":           "∡",
		"parallel_to":              "∥",
		"union":                    "∪",
		"intergral":                "∫",
		"therefore":                "∴",
		"proportional":             "∷",
		"homothetic":               "∻",
		"inverted_lazy_s":          "∾",
		"sine_wave":                "∿",
		"minux_tilde":              "≂",
		"approx_equal_to":          "≅",
		"geometrically_equivilent": "≎",
		"equal_colon":              "≕",
		"defined":                  "≝",
		"measured":                 "≞",
		"delta_equal_to":           "≜",
		"star_equals":              "≛",
		"long_menu":                "≣",
		"succeeds":                 "≻",
		"subset_of":                "⊂",
		"superset_of":              "⊃",
		"square_original_of":       "⊐",
		"circled_plus":             "⊕",
		"circled_times":            "⊗",
		"circled_asterisk":         "⊛",
		"circled_ring":             "⊚",
		"square_dot":               "⊡",
		"left_tack":                "⊣",
		"succeeds_under_relation":  "⊱",
		"original_of":              "⊶",
		"multimap":                 "⊸",
		"conjugate_matrix":         "⊹",
		"right_angle":              "⊾",
		"diamond":                  "⋄",
		"star":                     "⋆",
		"double_superset":          "⋑",
		"much_greater_than":        "⋙",
		"equal_or_succeeds":        "⋟",
		"vertical_ellipsis":        "⋮",
		"midline_ellipsis":         "⋯",
		"cursor_line":              "▁",
		"numero":                   "№",
		"prescription":             "℞",
		"ohm":                      "Ω",
		"clock":                    "⏰",
		"pencil":                   "✏",
		"x":                        "✖",
		"thick_arrow":              "➡",
		"id":                       "🆔",
		"ok":                       "🆗",
		"new":                      "🆕",
		"spiral":                   "🌀",
		"night":                    "🌃",
		"sunrise":                  "🌄",
		"sunset":                   "🌅",
		"wave":                     "🌊",
		"nightsky":                 "🌌",
		"europe_and_africa_globe":  "🌍",
		"americas_globe":           "🌎",
		"asia_globe":               "🌏",
		"moon":                     "🌑",
		"mushroom":                 "🍄",
		"candy":                    "🍬",
		"backpack":                 "🎒",
		"ticket":                   "🎟",
		"video":                    "🎞",
		"music":                    "🎧",
		"game":                     "🎮",
		"dice":                     "🎲",
		"house":                    "🏠",
		"house_with_tree":          "🏡",
		"powerplant":               "🏭",
		"capital":                  "🏢",
		"island":                   "🏝",
		"penguin":                  "🐧",
		"ghost":                    "👻",
		"laptop":                   "💻",
		"disk":                     "💽",
		"floppy":                   "💾",
		"directory":                "📁",
		"directory_open":           "📂",
		"bar_chart":                "📊",
		"calendar":                 "📅",
		"text_file":                "📄",
		"pin":                      "📌",
		"paperclip":                "📎",
		"fire":                     "🔥",
		"incognito":                "🕵",
		"pen":                      "🖊",
		"trash":                    "🗑",
		"archive":                  "🗄",
		"gray_key":                 "🗝",
		"notification":             "🔔",
		"random":                   "🔀",
		"label":                    "🏷",
		"image":                    "🖼",
		"link":                     "🔗",
		"email":                    "📧",
		"desktop":                  "🖥",
		"keyboard":                 "⌨",
		"battery":                  "🔋",
		"plug":                     "🔌",
		"knob":                     "🎛",
		"slider":                   "🎚",
		"mobile":                   "📱",
		"mouse":                    "🖱",
		"joystick":                 "🕹",
		"broadcast":                "📢",
		"satallite_antenna":        "📡",
		"tv":                       "📺",
		"radio":                    "📻",
		"projector":                "📽",
		"movie_camera":             "🎥",
		"printer":                  "🖨",
		"pill":                     "💊",
		"compression":              "🗜",
		"crystal_ball":             "🔮",
		"receipt":                  "🧾",
		"lock_with_ink_pen":        "🔏",
		"telescope":                "🔭",
		"microscope":               "🔬",
		"test_tube":                "🧪",
		"petri_dish":               "🧫",
		"gear":                     "⚙",
		"magnet":                   "🧲",
		"puzzle_piece":             "🧩",
		"nazar_amulet":             "🧿",
		"wrench":                   "🔧",
		"anchor":                   "⚓",
		"star_filled":              "★",
		"star_empty":               "☆",
		"eight_point_asterisk":     "✳",
		"snowflake":                "❄",
		"office_building":          "🏢",
		"cityscape":                "🏙",
		"tree":                     "🌳",
		"sprout":                   "🌱",
		"owl":                      "🦉",
		"red_circle":               "⭕",
		"knife":                    "🗡",
		"vote":                     "🗳",
		"map":                      "🗺",
		"drive":                    "🚨",
		"forbidden":                "🚫",
		"wrench_and_hammer":        "🛠",
		"satallite":                "🛰",
		"ufo":                      "🛸",
		"robot":                    "🤖",
		"heart":                    "🧡",
		"debian":                   "🍥",
		"bread":                    "🍞",
		"tornado":                  "🌪",
		"lock":                     "🔒",
		"open_lock":                "🔓",
		"lock_and_key":             "🔐",
		"key":                      "🔑",
	},
	"desktop": map[string]string{
		"home":               "🏠",
		"trash":              "🗑",
		"compress":           "🗜",
		"visible":            "👁",
		"pencil":             "✏",
		"writing":            "✍",
		"lower_right_pencil": "✎",
		"upper_right_pencil": " ✐",
		"stopped_hourglass":  "⌛",
		"active_hourglass":   "⏳",
	},
	"www": map[string]string{
		"link":                   "🔗",
		"europe_and_africa":      "🌍",
		"americas":               "🌎",
		"asia_and_oceania":       "🌏",
		"globe_with_meridians":   "🌐",
		"world_map":              "🗺",
		"anchor":                 "⚓",
		"wifi":                   "📶",
		"image":                  "🖼",
		"missing_image":          "⌧",
		"notifications":          "🔔",
		"disabled_notifications": "🔕",
		"bookmark":               "🔖",
		"tabs":                   "📑",
		"attach":                 "📎",
		"archive":                "🗃",
		"news":                   "📰",
		"pin":                    "📌",
		"email":                  "📧",
		"outbox":                 "📤",
		"inbox":                  "📥",
		"send":                   "📨",
		"receive":                "📩",
	},
	"book": map[string]string{
		"open":   "📖",
		"closed": "📕",
	},
	"keyboard": map[string]string{
		"windows":   "❖",
		"meta":      "◆",
		"shift":     "⇧",
		"helm":      "⎈",
		"capslock":  "⇪",
		"option":    "⌥",
		"alt":       "⎇",
		"command":   "⌘",
		"control":   "✲",
		"backspace": "⌫",
		"enter":     "⏎",
		"return":    "⏎",
		"menu":      "𝌆",
	},
	"sun": map[string]string{
		"filled": "☀",
		"empty":  "☼",
	},
	"musical": map[string]string{
		"quarter":        "♩",
		"eigth":          "♪",
		"violin":         "🎻",
		"saxophone":      "🎷",
		"guitar":         "🎸",
		"drum":           "🥁",
		"trumpet":        "🎺",
		"beamed_eigth":   "♫",
		"score":          "🎼",
		"slider":         "🎚",
		"knobs":          "🎛",
		"microphone":     "🎤",
		"headphone":      "🎧",
		"radio":          "📻",
		"multiple_notes": "🎶",
		"beamed_sixth":   "♬",
		"flat":           "♭",
		"natural":        "♮",
		"sharp":          "♯",
	},
	"weather": map[string]string{
		"sunny":        "☀",
		"cloudy":       "☁",
		"rainy":        "☂",
		"snowy":        "☃",
		"snowy2":       "🌨",
		"tornado":      "🌪",
		"foggy":        " 🌫",
		"rainbow":      "🌈",
		"stormy":       "⛈",
		"meteorshower": "☄",
	},
}
