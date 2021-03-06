//	This file is part of Fwew.
//	Fwew is free software: you can redistribute it and/or modify
// 	it under the terms of the GNU General Public License as published by
// 	the Free Software Foundation, either version 3 of the License, or
// 	(at your option) any later version.
//
//	Fwew is distributed in the hope that it will be useful,
//	but WITHOUT ANY WARRANTY; without even implied warranty of
//	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//	GNU General Public License for more details.
//
//	You should have received a copy of the GNU General Public License
//	along with Fwew.  If not, see http://gnu.org/licenses/

// Package main contains all the things. txt.go handles program strings.
package main

import (
	"fmt"
	"os/user"
	"path/filepath"
	"strings"
)

var usr, _ = user.Current()
var texts = map[string]string{}

func init() {
	// slash-commands Help
	texts["slashCommandHelp"] = "" +
		"/set [flag]\n" +
		"        set given options (separated by space)\n" +
		"        or show currently set options\n" +
		"/unset [flag]\n" +
		"        unset given options (separated by space)\n" +
		"        or show currently set options\n" +
		"/list <what> <cond> <spec> [and <what> <cond> <spec> ...]\n" +
		"        list all words that meet given criteria\n" +
		"/random <num> [where <what> <cond> <spec>]\n" +
		"        display given whole number > 0 of random entries\n" +
		"        for both /list and /random, <what>, <cond>, and <spec> work as follows:\n" +
		"        <what> is any one of: pos, word, words, syllables\n" +
		"        <cond> depends on the <what> used:\n" +
		"                <what>    | valid <cond>\n" +
		"                ----------|------------------------------------\n" +
		"                pos       | any one of: is, has, like\n" +
		"                word      | any one of: starts, ends, has, like\n" +
		"                words     | any one of: first, last\n" +
		"                syllables | any one of: <, <=, =, >=, >\n" +
		"        <spec> depends on the <cond> used:\n" +
		"                <cond>                       | valid <spec>\n" +
		"                -----------------------------|----------------------\n" +
		"                is, has, starts, ends        | any string of letter(s)\n" +
		"                <, <=, =, >=, >, first, last | any whole number > 0\n" +
		"                like                         | any string of letter(s) and\n" +
		"                                             |     wildcard asterisk(s)\n" +
		"/update\n" +
		"        download and update the dictionary file\n" +
		"/commands\n" +
		"        show this commands help text\n" +
		"/help\n" +
		"        show main help text\n" +
		"/exit\n" +
		"        exit/quit the program (aliases /quit /q /wc)\n"

	// <what> strings
	texts["w_pos"] = "pos"
	texts["w_word"] = "word"
	texts["w_words"] = "words"
	texts["w_syllables"] = "syllables"
	// <cond> strings
	texts["c_is"] = "is"
	texts["c_has"] = "has"
	texts["c_like"] = "like"
	texts["c_starts"] = "starts"
	texts["c_ends"] = "ends"
	texts["c_not-is"] = "not-is"
	texts["c_not-has"] = "not-has"
	texts["c_not-like"] = "not-like"
	texts["c_not-starts"] = "not-starts"
	texts["c_not-ends"] = "not-ends"
	texts["c_first"] = "first"
	texts["c_last"] = "last"

	// random
	texts["n_random"] = "random"

	// prompt suggest strings
	texts["/setDesc"] = "set option(s)"
	texts["/unsetDesc"] = "unset option(s)"
	texts["/listDesc"] = "list entries satisfying given condition(s)"
	texts["/randomDesc"] = "list random entries"
	texts["/updateDesc"] = "update the dictionary data file"
	texts["/commandsDesc"] = "show commands help"
	texts["/lenitionDesc"] = "show lenition table"
	texts["/configDesc"] = "edit configuration file"
	texts["/versionDesc"] = "show version info"
	texts["/helpDesc"] = "show usage help"
	texts["/exitDesc"] = "end program"
	texts["andDesc"] = "add condition to narrow search"
	texts["cDesc"] = texts["/configDesc"]
	texts["l=deDesc"] = "Deutsch"
	texts["l=engDesc"] = "English"
	texts["l=estDesc"] = "Eesti"
	texts["l=frDesc"] = "Français"
	texts["l=huDesc"] = "Magyar"
	texts["l=nlDesc"] = "Nederlands"
	texts["l=plDesc"] = "Polski"
	texts["l=ruDesc"] = "Русский"
	texts["l=svDesc"] = "Svenska"
	texts["posDesc"] = "part of speech"
	texts["wordDesc"] = texts["w_word"]
	texts["wordsDesc"] = texts["w_words"]
	texts["syllablesDesc"] = texts["w_syllables"]
	texts["randomDesc"] = "random number"
	texts["whereDesc"] = "add condition to random"
	texts["startsDesc"] = "field starts with"
	texts["endsDesc"] = "field ends with"
	texts["likeDesc"] = "field matches wildcard expression"
	texts["firstDesc"] = "list oldest words"
	texts["lastDesc"] = "list newest words"
	texts["hasDesc"] = "all matches of condition"
	texts["isDesc"] = "exact matches of condition"
	texts[">=Desc"] = "syllable count greater or equal"
	texts[">Desc"] = "syllable count greater"
	texts["<=Desc"] = "syllable count less or equal"
	texts["<Desc"] = "syllable count less"
	texts["=Desc"] = "syllable count equal"
	texts["languageDesc"] = "update config file: set language"
	texts["posFilterDesc"] = "update config file: set part of speech filter"
	texts["useAffixesDesc"] = "update config file: toggle affix parsing true/false"
	texts["debugModeDesc"] = "update config file: toggle debug mode true/false"
	texts["allDesc"] = "update config file: set part of speech filter to show all"
	texts["trueDesc"] = "update config file: set value to true"
	texts["falseDesc"] = "update config file: set value to false"
	texts["not-startsDesc"] = "field does not start with"
	texts["not-endsDesc"] = "field does not end with"
	texts["not-likeDesc"] = "field does not match wildcard expression"
	texts["not-hasDesc"] = "all matches of not condition"
	texts["not-isDesc"] = "exact matches of not condition"
	texts["!=Desc"] = "syllable count not equal"

	// file strings
	texts["homeDir"], _ = filepath.Abs(usr.HomeDir)
	texts["dataDir"] = filepath.Join(texts["homeDir"], ".fwew")
	texts["config"] = filepath.Join(texts["dataDir"], "config.json")
	texts["dictionary"] = filepath.Join(texts["dataDir"], "dictionary.txt")
	texts["dictURL"] = "https://tirea.learnnavi.org/dictionarydata/dictionary.txt"
	texts["dlSuccess"] = texts["dictURL"] + "\nsaved to\n" + texts["dictionary"] + "\n"

	// general message strings
	texts["cset"] = "currently set"
	texts["set"] = "set"
	texts["unset"] = "unset"
	texts["pre"] = "Prefixes"
	texts["inf"] = "Infixes"
	texts["suf"] = "Suffixes"
	texts["src"] = "source"
	texts["configSaved"] = "config file successfully updated\n"

	// error message strings
	texts["none"] = "no results\n"
	texts["noDataError"] = "err 1: failed to open dictionary file (" + texts["dictionary"] + ")"
	texts["fileError"] = "err 2: failed to open configuration file (" + texts["config"] + ")"
	texts["noOptionError"] = "err 3: invalid option"
	texts["invalidIntError"] = "err 4: input must be a decimal integer in range 0 <= n <= 32767 or octal integer in range 0 <= n <= 77777"
	texts["invalidOctalError"] = "err 5: invalid octal integer"
	texts["invalidDecimalError"] = "err 6: invalid decimal integer"
	texts["invalidLanguageError"] = "err 7: invalid language option"
	texts["invalidPOSFilterError"] = "err 8: invalid part of speech filter"
	texts["dictCloseError"] = "err 9: failed to close dictionary file (" + texts["dictionary"] + ")"
	texts["noFileError"] = "err 10: failed to open file"
	texts["fileCloseError"] = "err 11: failed to close input file"
	texts["configSyntaxError"] = "err 12: invalid syntax for config"
	texts["configOptionError"] = "err 13: invalid config option"
	texts["configValueError"] = "err 14: invalid config value for"
	texts["invalidNumericError"] = "err 15: invalid numeric digits"
	texts["downloadError"] = "err 16: could not download dictionary update"

	// main program strings
	texts["name"] = "fwew"
	texts["tip"] = "type \"/help\" or \"/commands\" for more info"
	texts["author"] = "Tirea Aean"
	Version.DictBuild = SHA1Hash(texts["dictionary"])
	texts["header"] = fmt.Sprintf("%s\n%s\n", Version, texts["tip"])
	texts["languages"] = "de, en, et, fr, hu, nl, pl, ru, sv"
	texts["POSFilters"] = "allvtr.n.num.pn.adv.adj.vin.v.inter.part.svin.adp.adv., n.vtrm.vim.conj.pn., sbd.n., intj.intj."
	texts["POSFilters"] += "vtrm., vtr.part., intj.vin., svin.prop.n.affixvin., intj.dem.dem., n.sbd.n., adv."
	texts["POSFilters"] += "adj., n.adj., adv.adj., intj.dem., pn.vtr., vin.adv., intj.pn., adv.ph.vin., vtr.adj.,  conj."
	texts["prompt"] = "~~> "

	// flag strings
	texts["usage"] = "usage"
	texts["bin"] = strings.ToLower(texts["name"])
	texts["options"] = "options"
	texts["usageV"] = "show program & dictionary version numbers"
	texts["usageL"] = "use specified language \n\tValid values: " + texts["languages"]
	texts["usageI"] = "display infix location data in bracket notation"
	texts["usageID"] = "display infix location data in dot notation"
	texts["usageIPA"] = "display IPA data"
	texts["usageS"] = "display syllable/stress breakdown"
	texts["usageSrc"] = "display source data"
	texts["usageP"] = "search for word(s) with specified part of speech"
	texts["usageR"] = "reverse the lookup direction from Na'vi->local to local->Na'vi"
	texts["usageA"] = "find all matches by using affixes to match the input word"
	texts["usageN"] = "convert numbers octal<->decimal"
	texts["usageM"] = "format output in markdown for bold and italic (mostly useful for fwew-discord bot)"
	texts["usageF"] = "filename of file to read as input"
	texts["usageC"] = "edit variable in configuration file"
	texts["usageD"] = "enable debug mode"
	texts["defaultFilter"] = "all"

	// lenition table
	texts["lenTable"] = "" +
		"px, tx, kx → p,  t,  k\n" +
		"p,  t,  k  → f,  s,  h\n" +
		"        ts → s\n" +
		"        '  → (disappears)"
}

// Text function is the accessor for []string texts
func Text(s string) string {
	return texts[s]
}
