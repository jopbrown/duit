// +build windows

package draw

import (
	draw "github.com/ktye/duitdraw"
)

// type alias
type (
	Color       = draw.Color
	Display     = draw.Display
	Cursor      = draw.Cursor
	Font        = draw.Font
	Image       = draw.Image
	Keyboardctl = draw.Keyboardctl
	Mouse       = draw.Mouse
	Mousectl    = draw.Mousectl
	Pix         = draw.Pix
)

// package exported function alias
var (
	Init    = draw.Init
	MakePix = draw.MakePix
)

// color const alias
const (
	Opaque        Color = draw.Opaque
	Transparent   Color = draw.Transparent
	Black         Color = draw.Black
	White         Color = draw.White
	Red           Color = draw.Red
	Green         Color = draw.Green
	Blue          Color = draw.Blue
	Cyan          Color = draw.Cyan
	Magenta       Color = draw.Magenta
	Yellow        Color = draw.Yellow
	Paleyellow    Color = draw.Paleyellow
	Darkyellow    Color = draw.Darkyellow
	Darkgreen     Color = draw.Darkgreen
	Palegreen     Color = draw.Palegreen
	Medgreen      Color = draw.Medgreen
	Darkblue      Color = draw.Darkblue
	Palebluegreen Color = draw.Palebluegreen
	Paleblue      Color = draw.Paleblue
	Bluegreen     Color = draw.Bluegreen
	Greygreen     Color = draw.Greygreen
	Palegreygreen Color = draw.Palegreygreen
	Yellowgreen   Color = draw.Yellowgreen
	Medblue       Color = draw.Medblue
	Greyblue      Color = draw.Greyblue
	Palegreyblue  Color = draw.Palegreyblue
	Purpleblue    Color = draw.Purpleblue

	Notacolor Color = draw.Notacolor
	Nofill    Color = draw.Nofill
)

// diaplay const alias
const (
	Refbackup       = draw.Refbackup
	Refnone         = draw.Refnone
	Refmesg         = draw.Refmesg
	DefaultDPI      = draw.DefaultDPI
	DefaultFontSize = draw.DefaultFontSize
)

// keyboard const alias
const (
	KeyFn = draw.KeyFn

	KeyHome   = draw.KeyHome
	KeyUp     = draw.KeyUp
	KeyPageUp = draw.KeyPageUp
	//KeyPrint     = draw.KeyPrint
	KeyLeft  = draw.KeyLeft
	KeyRight = draw.KeyRight
	KeyDown  = draw.KeyDown
	//KeyView      = draw.KeyView
	KeyPageDown = draw.KeyPageDown
	KeyInsert   = draw.KeyInsert
	KeyEnd      = draw.KeyEnd
	//KeyAlt       = draw.KeyAlt
	//KeyShift     = draw.KeyShift
	//KeyCtl       = draw.KeyCtl
	//KeyBackspace = draw.KeyBackspace
	KeyDelete = draw.KeyDelete
	KeyEscape = draw.KeyEscape
	//KeyEOF       = draw.KeyEOF
	KeyCmd = draw.KeyCmd
)

// pix const alias
const (
	CRed    = draw.CRed
	CGreen  = draw.CGreen
	CBlue   = draw.CBlue
	CGrey   = draw.CGrey
	CAlpha  = draw.CAlpha
	CMap    = draw.CMap
	CIgnore = draw.CIgnore
	NChan   = draw.NChan
)

// pix variable alias
var (
	GREY8  = draw.GREY8
	ARGB32 = draw.ARGB32
	ABGR32 = draw.ABGR32
)
