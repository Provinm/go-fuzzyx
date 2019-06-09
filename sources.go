package fuzzyx


// InitialRank 声母分值表
var InitialRank = map[string]map[string]int {
	"ch": {"ch": 100, "c": 90, "zh": 80, "sh": 80, "r": 80, "z": 70, "s": 70, " ": 50},
 	"zh": {"zh": 100, "z": 90, "ch": 80, "sh": 80, "r": 80, "c": 70, "s": 70," ": 50,},
 	" ": {
		"zh": 50,
		"ch": 50,
		"sh": 50,
		"b": 50,
		"p": 50,
		"m": 50,
		"f": 50,
		"d": 50,
		"t": 50,
		"l": 50,
		"k": 50,
		"h": 50,
		"j": 50,
		"q": 50,
		"x": 50,
		"z": 50,
		"c": 50,
		"s": 50,
		"y": 50,
		"w": 50,
		"g": 50,
		"n": 50,
		" ": 100,
	},
	"r": {"r": 100, "zh": 80, "ch": 80, "sh": 80, " ": 50},
	"c": {"c": 100, "ch": 90, "z": 80, "s": 80, "zh": 70, "sh": 70, " ": 50},
	"b": {"b": 100, "p": 80, "m": 80, " ": 50},
	"d": {"d": 100, "t": 80, "l": 80, "n": 80, " ": 50},
	"g": {"g": 100, "k": 80, "h": 80, " ": 50},
	"f": {"f": 100, " ": 50},
	"h": {"h": 100, "k": 80, "g": 80, " ": 50},
	"k": {"k": 100, "h": 80, "g": 80, " ": 50},
	"j": {"j": 100, "q": 80, "x": 80, " ": 50},
	"m": {"m": 100, "b": 80, "p": 80, " ": 50},
	"l": {"l": 100, "d": 80, "t": 80, "n": 80, " ": 50},
	"n": {"n": 100, "d": 80, "t": 80, "l": 80, " ": 50},
	"q": {"q": 100, "j": 80, "x": 80, " ": 50},
	"p": {"p": 100, "b": 80, "m": 80, " ": 50},
	"s": {"s": 100, "sh": 90, "z": 80, "c": 80, "zh": 70, "ch": 70, " ": 50},
	"sh": {"sh": 100, "s": 90, "zh": 80, "ch": 80, "r": 80, "z": 70, "c": 70, " ": 50},
	"t": {"t": 100, "d": 80, "l": 80, "n": 80, " ": 50},
	"w": {"w": 100, " ": 50},
	"y": {"y": 100, " ": 50},
	"x": {"x": 100, "j": 80, "q": 80},
	"z": {"z": 100, "zh": 90, "c": 80, "s": 80, "ch": 70, "sh": 70, " ": 50}}


var FinalRank = map[string]map[string]int {
	"ang": {
		"ang": 100,
		"an": 80,
		"eng": 80,
		"uang": 80,
		"iang": 80,
		"a": 50,
		"ai": 50,
		"en": 50,
		"ong": 50,
		"uan": 40,
		"e": 30,
		"ao": 30,
	},
	"ei": {
		"ei": 100,
		"ui": 80,
		"ai": 60,
		"en": 60,
		"ie": 50,
		"an": 50,
		"in": 40,
		"eng": 30,
		"i": 20,
		"a": 10,
	},
	"ai": {
		"ai": 100,
		"an": 80,
		"uai": 80,
		"a": 60,
		"ei": 60,
		"e": 50,
		"ao": 50,
		"ang": 50,
	},
	"uan": {
		"uan": 100,
		"uang": 90,
		"an": 80,
		"ua": 80,
		"un": 60,
		"u": 50,
		"uai": 50,
		"uo": 50,
		"ui": 40,
		"ang": 40,
		"ian": 40,
		"ue": 20,
	},
	"iu": {
		"iu": 100,
		"ou": 80,
		"i": 30,
		"ie": 30,
		"iong": 30,
		"ia": 20,
		"iang": 20,
	},
	"ong": {
		"ong": 100,
		"iong": 90,
		"o": 60,
		"eng": 60,
		"ang": 50,
		"uo": 20,
	},
	"ao": {
		"ao": 100,
		"o": 80,
		"iao": 80,
		"e": 60,
		"ou": 60,
		"a": 50,
		"ai": 50,
		"an": 30,
		"ang": 30,
		"en": 20,
		"eng": 20,
	},
	"an": {
		"an": 100,
		"ai": 80,
		"ang": 80,
		"uan": 80,
		"ian": 80,
		"en": 70,
		"eng": 70,
		"a": 60,
		"ei": 50,
		"un": 40,
		"e": 30,
		"ao": 30,
		"uang": 20,
	},
	"uai": {
		"uai": 100,
		"ai": 80,
		"ua": 70,
		"ui": 60,
		"u": 50,
		"uan": 50,
		"uang": 40,
		"uo": 40,
		"un": 30,
	},
	"en": {
		"en": 100,
		"eng": 90,
		"in": 80,
		"un": 80,
		"e": 70,
		"an": 70,
		"ei": 60,
		"ang": 50,
		"a": 30,
		"ao": 20,
	},
	"iong": {
		"iong": 100,
		"ong": 90,
		"iang": 70,
		"i": 50,
		"ing": 50,
		"ian": 50,
		"iao": 50,
		"ia": 50,
		"in": 40,
		"iu": 30,
		"ie": 30,
	},
	"in": {
		"in": 100,
		"en": 80,
		"ing": 80,
		"ian": 80,
		"i": 50,
		"ei": 40,
		"ia": 40,
		"iong": 40,
		"iang": 40,
		"ie": 30,
		"eng": 30,
	},
	"ia": {
		"ia": 100,
		"ian": 80,
		"iao": 70,
		"iang": 60,
		"a": 50,
		"i": 50,
		"ie": 50,
		"ing": 50,
		"iong": 50,
		"in": 40,
		"iu": 20,
	},
	"ing": {
		"ing": 100,
		"in": 80,
		"eng": 80,
		"ian": 80,
		"i": 50,
		"ia": 50,
		"iong": 50,
		"iang": 50,
		"iao": 40,
		"ie": 30,
	},
	"ie": {
		"ie": 100,
		"i": 50,
		"ei": 50,
		"ia": 50,
		"ue": 50,
		"ian": 40,
		"iao": 40,
		"iu": 30,
		"in": 30,
		"ing": 30,
		"iong": 30,
		"iang": 30,
		"e": 10,
	},
	"er": {
		"er": 100,
		"e": 80,
	},
	"iao": {
		"iao": 100,
		"ao": 80,
		"ian": 80,
		"ia": 70,
		"i": 50,
		"iong": 50,
		"iang": 50,
		"ie": 40,
		"ing": 40,
		"ou": 30,
	},
	"ian": {
		"ian": 100,
		"an": 80,
		"in": 80,
		"ing": 80,
		"iao": 80,
		"ia": 80,
		"iang": 80,
		"i": 50,
		"iong": 50,
		"ie": 40,
		"uan": 40,
	},
	"eng": {
		"eng": 100,
		"en": 90,
		"ang": 80,
		"ing": 80,
		"an": 70,
		"un": 70,
		"ong": 60,
		"e": 30,
		"ei": 30,
		"in": 30,
		"a": 20,
		"ao": 20,
	},
	"iang": {
		"iang": 100,
		"ang": 80,
		"ian": 80,
		"iong": 70,
		"ia": 60,
		"i": 50,
		"ing": 50,
		"iao": 50,
		"in": 40,
		"ie": 30,
		"iu": 20,
	},
	"u": {
		"u": 100,
		"ui": 50,
		"un": 50,
		"uan": 50,
		"ua": 50,
		"uai": 50,
		"uang": 50,
		"uo": 50,
		"ue": 50,
		"a": 20,
		"e": 10,
	},
	"uang": {
		"uang": 100,
		"uan": 90,
		"ang": 80,
		"ua": 70,
		"u": 50,
		"un": 50,
		"uai": 40,
		"uo": 40,
		"ui": 30,
		"an": 20,
	},
	"a": {
		"a": 100,
		"ua": 80,
		"ai": 60,
		"an": 60,
		"ao": 50,
		"ang": 50,
		"ia": 50,
		"en": 30,
		"o": 20,
		"u": 20,
		"ou": 20,
		"eng": 20,
		"ei": 10,
	},
	"e": {
		"e": 100,
		"o": 80,
		"er": 80,
		"uo": 80,
		"en": 70,
		"ao": 60,
		"ou": 60,
		"ai": 50,
		"ue": 50,
		"an": 30,
		"ang": 30,
		"eng": 30,
		"u": 10,
		"ie": 10,
	},
	"i": {
		"i": 100,
		"v": 50,
		"ie": 50,
		"in": 50,
		"ing": 50,
		"ian": 50,
		"iao": 50,
		"ia": 50,
		"iong": 50,
		"iang": 50,
		"iu": 30,
		"ei": 20,
	},
	"o": {
		"o": 100,
		"e": 80,
		"ao": 80,
		"ou": 80,
		"uo": 80,
		"ong": 60,
		"a": 20,
	},
	"uo": {
		"uo": 100,
		"o": 80,
		"e": 80,
		"ua": 70,
		"ou": 60,
		"u": 50,
		"un": 50,
		"uan": 50,
		"uai": 40,
		"uang": 40,
		"ui": 30,
		"ong": 20,
	},
	"un": {
		"un": 100,
		"en": 80,
		"eng": 70,
		"ui": 60,
		"uan": 60,
		"u": 50,
		"uang": 50,
		"uo": 50,
		"an": 40,
		"ua": 30,
		"uai": 30,
	},
	"ui": {
		"ui": 100,
		"ei": 80,
		"un": 60,
		"uai": 60,
		"u": 50,
		"uan": 40,
		"ua": 30,
		"uang": 30,
		"uo": 30,
		"ue": 30,
	},
	"v": {
		"v": 100,
		"i": 50,
		"ue": 50,
	},
	"ue": {
		"ue": 100,
		"e": 50,
		"u": 50,
		"v": 50,
		"ie": 50,
		"ui": 30,
		"uan": 20,
	},
	"ou": {
		"ou": 100,
		"o": 80,
		"iu": 80,
		"e": 60,
		"ao": 60,
		"uo": 60,
		"iao": 30,
		"a": 20,
	},
	"ua": {
		"ua": 100,
		"a": 80,
		"uan": 80,
		"uai": 70,
		"uang": 70,
		"uo": 70,
		"u": 50,
		"ui": 30,
		"un": 30,
	},
}