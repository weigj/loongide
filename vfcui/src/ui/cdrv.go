// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

/*
extern int drv(int id, int action, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void*a9);

void init_callback(int classid, int drvid)
{
	extern int drv_callback(void*,void*,void*,void*,void*);
	drv(classid,drvid,&drv_callback,0,0,0,0,0,0,0,0,0);
}

*/
import "C"
import "unsafe"
import "fmt"
import "os"
import "reflect"

type ItemHead struct {
	native  uintptr
	classid int
}

func (p *ItemHead) Head() *ItemHead {
	return p
}

func (p *ItemHead) Native() uintptr {
	return p.native
}

func itemString(item Item) string {
	return fmt.Sprintf("&{%s{%p}}", item.Name(), item)
}

func (p *ItemHead) String() string {
	return itemString(p)
}

func (p *ItemHead) Name() string {
	return "ItemHead"
}

func (p *ItemHead) SetAttr(attr string, value interface{}) bool {
	return false
}

func (p *ItemHead) Attr(attr string) interface{} {
	return nil
}

type Item interface {
	Head() *ItemHead
	Native() uintptr
	Name() string
	String() string
}

type WidgetItem interface {
	Item
	SetLayout(LayoutItem)
	Layout() LayoutItem
}

type LayoutItem interface {
	Item
	AddWidget(WidgetItem)
	AddWidgetWith(widget WidgetItem, stretch int, alignment int)
}

type MenuItem interface {
	Item
}

type StringHead struct {
	Data uintptr
	Len  int
}

func (s *StringHead) String() string {
	if s.Data == 0 {
		return ""
	}
	return string((*[1 << 30]byte)(unsafe.Pointer(s.Data))[0:s.Len])
}

type SliceHead struct {
	Data uintptr
	Len  int
	Cap  int
}

type FontInfo struct {
	Family     string
	PointSize  int
	Weight     int
	Stretch    int
	Bold       bool
	Italic     bool
	Strikeout  bool
	Underline  bool
	Overline   bool
	FixedPitch bool
}

func (p *FontInfo) Init() *FontInfo {
	p.Weight = -1
	p.PointSize = 0
	p.Stretch = 100
	p.Bold = false
	p.Italic = false
	p.Strikeout = false
	p.Underline = false
	p.Overline = false
	p.FixedPitch = false
	return p
}

func Font() *FontInfo {
	return new(FontInfo).Init()
}

func FontWith(family string, size int, weight int) *FontInfo {
	return &FontInfo{Family: family, PointSize: size, Weight: weight, Stretch: 100}
}

func _b(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func font_to_head(f *FontInfo) *font_head {
	if f.Weight > 50 {
		f.Bold = true
	}
	return &font_head{(*StringHead)(unsafe.Pointer(&f.Family)), f.PointSize, f.Weight, f.Stretch,
		_b(f.Bold), _b(f.Italic), _b(f.Strikeout), _b(f.Underline), _b(f.Overline), _b(f.FixedPitch)}
}

type font_head struct {
	family     *StringHead
	pointsize  int
	weight     int
	stretch    int
	bold       byte
	italic     byte
	strikeout  byte
	underline  byte
	overline   byte
	fixedpitch byte
}

func (f *font_head) Font() *FontInfo {
	if f.family == nil {
		return nil
	}
	return &FontInfo{f.family.String(), f.pointsize, f.weight, f.stretch,
		f.bold != 0, f.italic != 0, f.strikeout != 0, f.underline != 0, f.overline != 0, f.fixedpitch != 0}
}

func _p(a interface{}) uintptr {
	if a == nil {
		return 0
	}
	switch v := a.(type) {
	case bool:
		n := 0
		if v {
			n = 1
		}
		return uintptr(unsafe.Pointer(&n))
	case uintptr:
		return v
	case *uintptr:
		return uintptr(unsafe.Pointer(v))
	case int:
		return uintptr(unsafe.Pointer(&v))
	case *int:
		return uintptr(unsafe.Pointer(v))
	case uint:
		return uintptr(unsafe.Pointer(&v))
	case *uint:
		return uintptr(unsafe.Pointer(v))
	case string:
		return uintptr(unsafe.Pointer((*StringHead)(unsafe.Pointer(&v))))
	case *StringHead:
		return uintptr(unsafe.Pointer(v))
	case *ItemHead:
		return uintptr(unsafe.Pointer(v))
	case Item:
		return uintptr(unsafe.Pointer(v.Head()))
	case *Item:
		return uintptr(unsafe.Pointer((*v).Head()))
	case Point:
		return uintptr(unsafe.Pointer(&v))
	case *Point:
		return uintptr(unsafe.Pointer(v))
	case Size:
		return uintptr(unsafe.Pointer(&v))
	case *Size:
		return uintptr(unsafe.Pointer(v))
	case Rect:
		return uintptr(unsafe.Pointer(&v))
	case *Rect:
		return uintptr(unsafe.Pointer(v))
	case Margins:
		return uintptr(unsafe.Pointer(&v))
	case *Margins:
		return uintptr(unsafe.Pointer(v))
	case []Point:
		return uintptr(unsafe.Pointer((*SliceHead)(unsafe.Pointer(&v))))
	case FontInfo:
		return uintptr(unsafe.Pointer(font_to_head(&v)))
	case *FontInfo:
		return uintptr(unsafe.Pointer(font_to_head(v)))
	case *font_head:
		return uintptr(unsafe.Pointer(v))
	default:
		warning("Warning drv, param type \"%s\" not match!", reflect.TypeOf(v))
	}
	return 0
}

func drv10(id int, act int, a0, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) int {
	return int(C.drv(C.int(id), C.int(act), unsafe.Pointer(a0), unsafe.Pointer(a1), unsafe.Pointer(a2), unsafe.Pointer(a3), unsafe.Pointer(a4), unsafe.Pointer(a5), unsafe.Pointer(a6), unsafe.Pointer(a7), unsafe.Pointer(a8), unsafe.Pointer(a9)))
}

func drv(id int, act int, a ...interface{}) int {
	switch len(a) {
	case 0:
		return drv10(id, act, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	case 1:
		return drv10(id, act, _p(a[0]), 0, 0, 0, 0, 0, 0, 0, 0, 0)
	case 2:
		return drv10(id, act, _p(a[0]), _p(a[1]), 0, 0, 0, 0, 0, 0, 0, 0)
	case 3:
		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), 0, 0, 0, 0, 0, 0, 0)
	case 4:
		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), _p(a[3]), 0, 0, 0, 0, 0, 0)
	case 5:
		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), _p(a[3]), _p(a[4]), 0, 0, 0, 0, 0)
	case 6:
		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), _p(a[3]), _p(a[4]), _p(a[5]), 0, 0, 0, 0)
	case 7:
		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), _p(a[3]), _p(a[4]), _p(a[5]), _p(a[6]), 0, 0, 0)
	case 8:
		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), _p(a[3]), _p(a[4]), _p(a[5]), _p(a[6]), _p(a[7]), 0, 0)
	case 9:
		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), _p(a[3]), _p(a[4]), _p(a[5]), _p(a[6]), _p(a[7]), _p(a[9]), 0)
	case 10:
		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), _p(a[3]), _p(a[4]), _p(a[5]), _p(a[6]), _p(a[7]), _p(a[9]), _p(a[10]))
	default:
		panic("drv param count must <= 10")
	}
	return 0
}

var func_map = make(map[unsafe.Pointer]interface{})
var item_map = make(map[uintptr]Item)

var fnOnAddItem func(Item)
var fnOnRemoveItem func(Item)

func OnAddItem(fn func(Item)) {
	fnOnAddItem = fn
}

func OnRemoveItem(fn func(Item)) {
	fnOnRemoveItem = fn
}

func AddItem(item Item) {
	item_map[item.Native()] = item
	if fnOnAddItem != nil {
		fnOnAddItem(item)
	}
}

func FindItem(native uintptr) Item {
	return item_map[native]
}

func RemoveItem(native uintptr) {
	item := FindItem(native)
	if item == nil {
		return
	}
	item.Head().native = 0
	if fnOnRemoveItem != nil && item != nil {
		fnOnRemoveItem(item)
	}
	item_map[native] = nil
	delete(item_map, native)
}

func init() {
	C.init_callback(-1, 0)
}

func drv_event(id int, event int, hi Item, fn interface{}) {
	var pfn unsafe.Pointer = unsafe.Pointer(&fn)
	func_map[pfn] = fn
	drv10(id, event, uintptr(unsafe.Pointer(hi.Head())), uintptr(pfn), 0, 0, 0, 0, 0, 0, 0, 0)
}

//export drv_callback
func drv_callback(pfn unsafe.Pointer, a1, a2, a3, a4 unsafe.Pointer) int {
	fn, ok := func_map[pfn]
	if !ok {
		return 0
	}
	switch v := (fn).(type) {
	case func():
		v()
	case func(int):
		v(*(*int)(a1))
	case func(uint):
		v(*(*uint)(a1))
	case func(bool):
		v(*(*int)(a1) != 0)
	case func(uintptr):
		v(uintptr(a1))
	case func(string):
		v(((*StringHead)(a1)).String())
	case func(*ShowEvent):
		v((*ShowEvent)(a1))
	case func(*HideEvent):
		v((*HideEvent)(a1))
	case func(*CloseEvent):
		v((*CloseEvent)(a1))
	case func(*KeyEvent):
		v((*KeyEvent)(a1))
	case func(*MouseEvent):
		v((*MouseEvent)(a1))
	case func(*MoveEvent):
		v((*MoveEvent)(a1))
	case func(*ResizeEvent):
		v((*ResizeEvent)(a1))
	case func(*EnterEvent):
		v((*EnterEvent)(a1))
	case func(*LeaveEvent):
		v((*LeaveEvent)(a1))
	case func(*FocusEvent):
		v((*FocusEvent)(a1))
	case func(*PaintEvent):
		v((*PaintEvent)(a1))
	case func(*TimerEvent):
		v((*TimerEvent)(a1))
	default:
		warning("Warning drv_callback, func type \"%s\" not match!", reflect.TypeOf(v))
	}
	return 1
}

func warning(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
