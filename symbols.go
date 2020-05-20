package cli

// A collection of UTF-8 symbols that work by default in Gnome terminal on
// Debian AND that are specifically useful for UI design.

// https://en.wikipedia.org/wiki/Miscellaneous_Symbols
// https://www.w3schools.com/charsets/ref_html_utf8.asp

// TODO: This will be moved into its own package and called in optionally
var Settings = map[string]string{
	"brightness": "☼",
}

// ⋈ ⧓

// Could be useful for indicating corners
//  ◰ ◱ ◲ ◳

// tape drive ✇

// Stacked Lines  ▤ ☰ 𝌆

// Editing  ⁀ ⎁ ⎂ ⎃ �

// White space representation
//tab ↹ ⇄ ⇤ ⇥ ↤ ↦ ◁ ▷
//space · ␣ ˽  ␢
//paragraph, section¶ §

// Punctuation ‼

// Various
// ^ ⌃
// ✲
//  ⎇ ⌥ ✦ ✧   ⌤
// ⎋ ⌫  ⌦ ⎀
// ⌧  ⇞ ⇟   ⎉ ⎊ ⍰  ☾
// ⌂
// Info ℹ
// Letter ✉
// Cut ✂ ✄
// Reload ↶ ↷ ⟲ ⟳ ↺ ↻
