// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

const (
	AlignLeft            = 0x0001
	AlignLeading         = AlignLeft
	AlignRight           = 0x0002
	AlignTrailing        = AlignRight
	AlignHCenter         = 0x0004
	AlignJustify         = 0x0008
	AlignAbsolute        = 0x0010
	AlignHorizontal_Mask = AlignLeft | AlignRight | AlignHCenter | AlignJustify | AlignAbsolute

	AlignTop           = 0x0020
	AlignBottom        = 0x0040
	AlignVCenter       = 0x0080
	AlignVertical_Mask = AlignTop | AlignBottom | AlignVCenter

	AlignCenter = AlignVCenter | AlignHCenter
)

const (
	Unchecked        = 0
	PartiallyChecked = 1
	Checked          = 2
)

const (
	PlainText = 0
	RichText  = 1
	AutoText  = 2
	LogText   = 3
)

const (
	LightWeight    = 25
	NormalWeight   = 50
	DemiBoldWeight = 63
	BoldWeight     = 75
	BlackWeight    = 87
)

const (
	NoModifier          = 0x00000000
	ShiftModifier       = 0x02000000
	ControlModifier     = 0x04000000
	AltModifier         = 0x08000000
	MetaModifier        = 0x10000000
	KeypadModifier      = 0x20000000
	GroupSwitchModifier = 0x40000000
	// Do not extend the mask to include 0x01000000
	KeyboardModifierMask = 0xfe000000
)

//shorter names for shortcuts
const (
	META          = MetaModifier
	SHIFT         = ShiftModifier
	CTRL          = ControlModifier
	ALT           = AltModifier
	MODIFIER_MASK = KeyboardModifierMask
	UNICODE_ACCEL = 0x00000000
)

const (
	NoButton        = 0x00000000
	LeftButton      = 0x00000001
	RightButton     = 0x00000002
	MidButton       = 0x00000004 // ### Qt 5: remove me
	MiddleButton    = MidButton
	XButton1        = 0x00000008
	XButton2        = 0x00000010
	MouseButtonMask = 0x000000ff
)

type Point struct {
	X, Y int
}

func (p Point) Unpack() (int, int) {
	return p.X, p.Y
}

type Size struct {
	Width, Height int
}

func (p Size) Unpack() (int, int) {
	return p.Width, p.Height
}

type Rect struct {
	X, Y          int
	Width, Height int
}

func (p Rect) Unpack() (int, int, int, int) {
	return p.X, p.Y, p.Width, p.Height
}

func (r Rect) Point() Point {
	return Point{r.X, r.Y}
}

func (r Rect) Size() Size {
	return Size{r.Width, r.Height}
}

type Margins struct {
	Left, Top, Right, Bottom int
}

func (p Margins) Unpack() (int, int, int, int) {
	return p.Left, p.Top, p.Right, p.Bottom
}

func Pt(x, y int) Point {
	return Point{x, y}
}

func Sz(w, h int) Size {
	return Size{w, h}
}

func Rc(x, y, w, h int) Rect {
	return Rect{x, y, w, h}
}

func Mr(left, right, top, bottom int) Margins {
	return Margins{left, right, top, bottom}
}
