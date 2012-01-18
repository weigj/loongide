// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

// drvclass enums
const (
	DRVCLASS_NONE = iota
	DRVCLASS_OBJECT
	DRVCLASS_APP
	DRVCLASS_WIDGET
	DRVCLASS_MENUBAR
	DRVCLASS_MENU
	DRVCLASS_ACTION
	DRVCLASS_TABWIDGET
	DRVCLASS_LAYOUT
	DRVCLASS_BOXLAYOUT
	DRVCLASS_HBOXLAYOUT
	DRVCLASS_VBOXLAYOUT
	DRVCLASS_BASEBUTTON
	DRVCLASS_BUTTON
	DRVCLASS_CHECKBOX
	DRVCLASS_RADIO
	DRVCLASS_FRAME
	DRVCLASS_LABEL
	DRVCLASS_GROUPBOX
	DRVCLASS_DIALOG
	DRVCLASS_COMBOBOX
	DRVCLASS_PAINTER
)
// DRVCLASS_OBJECT drvid enums
const (
	OBJECT_NONE = iota
	OBJECT_DESTROY
	OBJECT_STARTTIMER
	OBJECT_KILLTIMER
	OBJECT_ONTIMEREVENT
)
// DRVCLASS_APP drvid enums
const (
	APP_NONE = iota
	APP_INIT
	APP_CLOSE
	APP_EXEC
	APP_EXIT
	APP_SETFONT
	APP_FONT
	APP_CLOSEALLWINDOWS
	APP_ONREMOVEITEM
)
// DRVCLASS_WIDGET drvid enums
const (
	WIDGET_NONE = iota
	WIDGET_INIT
	WIDGET_DESTROY
	WIDGET_SETAUTODESTROY
	WIDGET_AUTODESTROY
	WIDGET_SETLAYOUT
	WIDGET_LAYOUT
	WIDGET_SETPARENT
	WIDGET_PARENT
	WIDGET_SETVISIBLE
	WIDGET_ISVISIBLE
	WIDGET_SETWINDOWTITLE
	WIDGET_WINDOWTITLE
	WIDGET_SETPOS
	WIDGET_POS
	WIDGET_SETSIZE
	WIDGET_SIZE
	WIDGET_SETGEOMETRY
	WIDGET_GEOMETRY
	WIDGET_SETFONT
	WIDGET_FONT
	WIDGET_CLOSE
	WIDGET_UPDATE
	WIDGET_REPAINT
	WIDGET_ONSHOWEVENT
	WIDGET_ONHIDEEVENT
	WIDGET_ONCLOSEEVENT
	WIDGET_ONKEYPRESSEVENT
	WIDGET_ONKEYRELEASEEVENT
	WIDGET_ONMOUSEPRESSEVENT
	WIDGET_ONMOUSERELEASEEVENT
	WIDGET_ONMOUSEMOVEEVENT
	WIDGET_ONMOUSEDOUBLECLICKEVENT
	WIDGET_ONMOVEEVENT
	WIDGET_ONPAINTEVENT
	WIDGET_ONRESIZEEVENT
	WIDGET_ONENTEREVENT
	WIDGET_ONLEAVEEVENT
	WIDGET_ONFOCUSINEVENT
	WIDGET_ONFOCUSOUTEVENT
)
// DRVCLASS_MENUBAR drvid enums
const (
	MENUBAR_NONE = iota
	MENUBAR_INIT
	MENUBAR_ADDMENU
)
// DRVCLASS_MENU drvid enums
const (
	MENU_NONE = iota
	MENU_INIT
	MENU_SETTITLE
	MENU_TITLE
	MENU_ADDACTION
)
// DRVCLASS_ACTION drvid enums
const (
	ACTION_NONE = iota
	ACTION_INIT
	ACTION_SETTEXT
	ACTION_TEXT
	ACTION_ONTRIGGERED
)
// DRVCLASS_TABWIDGET drvid enums
const (
	TABWIDGET_NONE = iota
	TABWIDGET_INIT
	TABWIDGET_ADDTAB
	TABWIDGET_CLEAR
	TABWIDGET_COUNT
	TABWIDGET_CURRENTINDEX
	TABWIDGET_CURRENTWIDGET
	TABWIDGET_SETCURRENTINDEX
	TABWIDGET_SETCURRENTWIDGET
	TABWIDGET_INDEXOF
	TABWIDGET_INSERTTAB
	TABWIDGET_REMOVETAB
	TABWIDGET_SETTABTEXT
	TABWIDGET_SETTABTOOLTIP
	TABWIDGET_TABTEXT
	TABWIDGET_TABTOOLTIP
	TABWIDGET_WIDGETOF
	TABWIDGET_ONCURRENTCHANGED
)
// DRVCLASS_LAYOUT drvid enums
const (
	LAYOUT_NONE = iota
	LAYOUT_ADDLAYOUT
	LAYOUT_ADDWIDGET
	LAYOUT_SETSPACING
	LAYOUT_SPACING
	LAYOUT_SETMARGINS
	LAYOUT_MARGINS
	LAYOUT_SETMENUBAR
	LAYOUT_MENUBAR
)
// DRVCLASS_BOXLAYOUT drvid enums
const (
	BOXLAYOUT_NONE = iota
	BOXLAYOUT_ADDLAYOUTWITH
	BOXLAYOUT_ADDWIDGETWITH
	BOXLAYOUT_ADDSPACING
	BOXLAYOUT_ADDSTRETCH
)
// DRVCLASS_HBOXLAYOUT drvid enums
const (
	HBOXLAYOUT_NONE = iota
	HBOXLAYOUT_INIT
)
// DRVCLASS_VBOXLAYOUT drvid enums
const (
	VBOXLAYOUT_NONE = iota
	VBOXLAYOUT_INIT
)
// DRVCLASS_BASEBUTTON drvid enums
const (
	BASEBUTTON_NONE = iota
	BASEBUTTON_SETTEXT
	BASEBUTTON_TEXT
)
// DRVCLASS_BUTTON drvid enums
const (
	BUTTON_NONE = iota
	BUTTON_INIT
	BUTTON_SETFLAT
	BUTTON_ISFLAT
	BUTTON_SETDEFAULT
	BUTTON_ISDEFAULT
	BUTTON_SETMENU
	BUTTON_MENU
	BUTTON_ONCLICKED
)
// DRVCLASS_CHECKBOX drvid enums
const (
	CHECKBOX_NONE = iota
	CHECKBOX_INIT
	CHECKBOX_SETCHECK
	CHECKBOX_CHECK
	CHECKBOX_SETTRISTATE
	CHECKBOX_ISTRISTATE
	CHECKBOX_ONSTATECHANGED
)
// DRVCLASS_RADIO drvid enums
const (
	RADIO_NONE = iota
	RADIO_INIT
	RADIO_ONCLICKED
)
// DRVCLASS_FRAME drvid enums
const (
	FRAME_NONE = iota
	FRAME_INIT
	FRAME_SETFRAMESTYLE
	FRAME_FRAMESTYLE
	FRAME_SETFRAMERECT
	FRAME_FRAMERECT
)
// DRVCLASS_LABEL drvid enums
const (
	LABEL_NONE = iota
	LABEL_INIT
	LABEL_SETTEXT
	LABEL_TEXT
	LABEL_SETWORDWRAP
	LABEL_WORDWRAP
	LABEL_SETTEXTFORMAT
	LABEL_TEXTFORMAT
	LABEL_ONLINKACTIVATED
)
// DRVCLASS_GROUPBOX drvid enums
const (
	GROUPBOX_NONE = iota
	GROUPBOX_INIT
	GROUPBOX_SETTITLE
	GROUPBOX_TITLE
)
// DRVCLASS_DIALOG drvid enums
const (
	DIALOG_NONE = iota
	DIALOG_INIT
	DIALOG_SETMODAL
	DIALOG_ISMODAL
	DIALOG_SETRESULT
	DIALOG_RESULT
	DIALOG_EXEC
	DIALOG_DONE
	DIALOG_ACCEPT
	DIALOG_REJECT
	DIALOG_ONACCEPTED
	DIALOG_ONREJECTED
)
// DRVCLASS_COMBOBOX drvid enums
const (
	COMBOBOX_NONE = iota
	COMBOBOX_INIT
	COMBOBOX_COUNT
	COMBOBOX_SETCURRENTINDEX
	COMBOBOX_CURRENTINDEX
	COMBOBOX_CURRENTTEXT
	COMBOBOX_SETEDITABLE
	COMBOBOX_ISEDITABLE
	COMBOBOX_SETMAXCOUNT
	COMBOBOX_MAXCOUNT
	COMBOBOX_SETMAXVISIBLEITEMS
	COMBOBOX_MAXVISIBLEITEMS
	COMBOBOX_SETMINIMUMCONTENTSLENGTH
	COMBOBOX_MINIMUNCONTENTSLENGHT
	COMBOBOX_ADDITEM
	COMBOBOX_INSERTITEM
	COMBOBOX_REMOVEITEM
	COMBOBOX_ITEMTEXT
	COMBOBOX_ONCURRENTINDEXCHANGED
)
// DRVCLASS_PAINTER drvid enums
const (
	PAINTER_NONE = iota
	PAINTER_INIT
	PAINTER_DESTROY
	PAINTER_BEGIN
	PAINTER_END
	PAINTER_SETFONT
	PAINTER_FONT
	PAINTER_DRAWPOINT
	PAINTER_DRAWLINE
	PAINTER_DRAWLINES
	PAINTER_DRAWPOLYLINE
	PAINTER_DRAWTEXT
)
// object struct
type object struct {
	ItemHead
}

func (p *object) Name() string {
	return "object"
}

func (p *object) String() string {
	return itemString(p)
}
func (p *object) SetAttr(attr string, value interface{}) bool {
	switch attr {
	default:
		return p.ItemHead.SetAttr(attr,value)
	}
	return false
}
func (p *object) Attr(attr string) interface{} {
	switch attr {
	default:
		return p.ItemHead.Attr(attr)
	}
	return nil
}
func (p *object) Destroy() {
	drv(DRVCLASS_OBJECT,OBJECT_DESTROY,p)
	return
}

func (p *object) StartTimer(millisecond int)(id int) {
	drv(DRVCLASS_OBJECT,OBJECT_STARTTIMER,p,millisecond,&id)
	return
}

func (p *object) KillTimer(id int) {
	drv(DRVCLASS_OBJECT,OBJECT_KILLTIMER,p,id)
	return
}

func (p *object) OnTimerEvent(fn func(*TimerEvent)) {
	drv_event(DRVCLASS_OBJECT,OBJECT_ONTIMEREVENT,p,fn)
	return
}

// app struct
type app struct {
	object
}

func (p *app) Name() string {
	return "app"
}

func (p *app) String() string {
	return itemString(p)
}

func (p *app) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "font":
		if v, ok := value.(*FontInfo); ok {
			p.SetFont(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr,value)
	}
	return false
}
func (p *app) Attr(attr string) interface{} {
	switch attr {
	case "font":
		return p.Font()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func new_app() *app{
	return new(app).Init()
}

func (p *app) Init() *app {
	p.classid = DRVCLASS_APP
	drv(DRVCLASS_APP,APP_INIT,p)
	AddItem(p)
	return p
}

func (p *app) Close() {
	drv(DRVCLASS_APP,APP_CLOSE,p)
	return
}

func (p *app) Exec()(code int) {
	drv(DRVCLASS_APP,APP_EXEC,p,&code)
	return
}

func (p *app) Exit(code int) {
	drv(DRVCLASS_APP,APP_EXIT,p,code)
	return
}

func (p *app) SetFont(font *FontInfo) {
	drv(DRVCLASS_APP,APP_SETFONT,p,font)
	return
}

func (p *app) Font()(font *FontInfo) {
	var fh_font font_head
	drv(DRVCLASS_APP,APP_FONT,p,&fh_font)
	font = fh_font.Font()
	return
}

func (p *app) CloseAllWindows() {
	drv(DRVCLASS_APP,APP_CLOSEALLWINDOWS,p)
	return
}

func (p *app) onRemoveItem(fn func(uintptr)) {
	drv_event(DRVCLASS_APP,APP_ONREMOVEITEM,p,fn)
	return
}


var theApp *app

func Init() {
	theApp = new_app()
	theApp.onRemoveItem(RemoveItem)
}

func Close() {
	theApp.Close()
}

func RunLoop() int {
	return theApp.Exec()
}

func Exit(code int) {
	theApp.Exit(code)
}

func CloseAllWindows() {
	theApp.CloseAllWindows()
}

func DefaultFont() (font *FontInfo) {
	return theApp.Font()
}

func SetDefaultFont(font *FontInfo) {
	theApp.SetFont(font)
}


// Widget struct
type Widget struct {
	object
}

func (p *Widget) Name() string {
	return "Widget"
}

func (p *Widget) String() string {
	return itemString(p)
}

func (p *Widget) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "autodestroy":
		if v, ok := value.(bool); ok {
			p.SetAutoDestroy(v)
			return true
		}
		return false
	case "layout":
		if v, ok := value.(LayoutItem); ok {
			p.SetLayout(v)
			return true
		}
		return false
	case "parent":
		if v, ok := value.(WidgetItem); ok {
			p.SetParent(v)
			return true
		}
		return false
	case "visible":
		if v, ok := value.(bool); ok {
			p.SetVisible(v)
			return true
		}
		return false
	case "windowtitle":
		if v, ok := value.(string); ok {
			p.SetWindowTitle(v)
			return true
		}
		return false
	case "pos":
		if v, ok := value.(Point); ok {
			p.SetPos(v)
			return true
		}
		return false
	case "size":
		if v, ok := value.(Size); ok {
			p.SetSize(v)
			return true
		}
		return false
	case "geometry":
		if v, ok := value.(Rect); ok {
			p.SetGeometry(v)
			return true
		}
		return false
	case "font":
		if v, ok := value.(*FontInfo); ok {
			p.SetFont(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr,value)
	}
	return false
}
func (p *Widget) Attr(attr string) interface{} {
	switch attr {
	case "autodestroy":
		return p.AutoDestroy()
	case "layout":
		return p.Layout()
	case "parent":
		return p.Parent()
	case "visible":
		return p.IsVisible()
	case "windowtitle":
		return p.WindowTitle()
	case "pos":
		return p.Pos()
	case "size":
		return p.Size()
	case "geometry":
		return p.Geometry()
	case "font":
		return p.Font()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewWidget() *Widget{
	return new(Widget).Init()
}

func (p *Widget) Init() *Widget {
	p.classid = DRVCLASS_WIDGET
	drv(DRVCLASS_WIDGET,WIDGET_INIT,p)
	AddItem(p)
	return p
}

func (p *Widget) Destroy() {
	drv(DRVCLASS_WIDGET,WIDGET_DESTROY,p)
	return
}

func (p *Widget) SetAutoDestroy(b bool) {
	drv(DRVCLASS_WIDGET,WIDGET_SETAUTODESTROY,p,b)
	return
}

func (p *Widget) AutoDestroy()(b bool) {
	var b_b int
	drv(DRVCLASS_WIDGET,WIDGET_AUTODESTROY,p,&b_b)
	b = b_b != 0
	return
}

func (p *Widget) SetLayout(layout LayoutItem) {
	drv(DRVCLASS_WIDGET,WIDGET_SETLAYOUT,p,layout)
	return
}

func (p *Widget) Layout()(layout LayoutItem) {
	var hh_layout ItemHead
	drv(DRVCLASS_WIDGET,WIDGET_LAYOUT,p,&hh_layout)
	if hh_layout.native != 0 {
		item := FindItem(hh_layout.native)
		if item != nil {
			layout = item.(LayoutItem)
		}
	}
	return
}

func (p *Widget) SetParent(parent WidgetItem) {
	drv(DRVCLASS_WIDGET,WIDGET_SETPARENT,p,parent)
	return
}

func (p *Widget) Parent()(parent WidgetItem) {
	var hh_parent ItemHead
	drv(DRVCLASS_WIDGET,WIDGET_PARENT,p,&hh_parent)
	if hh_parent.native != 0 {
		item := FindItem(hh_parent.native)
		if item != nil {
			parent = item.(WidgetItem)
		}
	}
	return
}

func (p *Widget) SetVisible(b bool) {
	drv(DRVCLASS_WIDGET,WIDGET_SETVISIBLE,p,b)
	return
}

func (p *Widget) IsVisible()(b bool) {
	var b_b int
	drv(DRVCLASS_WIDGET,WIDGET_ISVISIBLE,p,&b_b)
	b = b_b != 0
	return
}

func (p *Widget) SetWindowTitle(text string) {
	drv(DRVCLASS_WIDGET,WIDGET_SETWINDOWTITLE,p,text)
	return
}

func (p *Widget) WindowTitle()(text string) {
	var sh_text StringHead
	drv(DRVCLASS_WIDGET,WIDGET_WINDOWTITLE,p,&sh_text)
	text = sh_text.String()
	return
}

func (p *Widget) SetPos(pt Point) {
	drv(DRVCLASS_WIDGET,WIDGET_SETPOS,p,pt)
	return
}

func (p *Widget) Pos()(pt Point) {
	drv(DRVCLASS_WIDGET,WIDGET_POS,p,&pt)
	return
}

func (p *Widget) SetSize(sz Size) {
	drv(DRVCLASS_WIDGET,WIDGET_SETSIZE,p,sz)
	return
}

func (p *Widget) Size()(sz Size) {
	drv(DRVCLASS_WIDGET,WIDGET_SIZE,p,&sz)
	return
}

func (p *Widget) SetGeometry(rc Rect) {
	drv(DRVCLASS_WIDGET,WIDGET_SETGEOMETRY,p,rc)
	return
}

func (p *Widget) Geometry()(rc Rect) {
	drv(DRVCLASS_WIDGET,WIDGET_GEOMETRY,p,&rc)
	return
}

func (p *Widget) SetFont(font *FontInfo) {
	drv(DRVCLASS_WIDGET,WIDGET_SETFONT,p,font)
	return
}

func (p *Widget) Font()(font *FontInfo) {
	var fh_font font_head
	drv(DRVCLASS_WIDGET,WIDGET_FONT,p,&fh_font)
	font = fh_font.Font()
	return
}

func (p *Widget) Close() {
	drv(DRVCLASS_WIDGET,WIDGET_CLOSE,p)
	return
}

func (p *Widget) Update() {
	drv(DRVCLASS_WIDGET,WIDGET_UPDATE,p)
	return
}

func (p *Widget) Repaint() {
	drv(DRVCLASS_WIDGET,WIDGET_REPAINT,p)
	return
}

func (p *Widget) OnShowEvent(fn func(*ShowEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONSHOWEVENT,p,fn)
	return
}

func (p *Widget) OnHideEvent(fn func(*HideEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONHIDEEVENT,p,fn)
	return
}

func (p *Widget) OnCloseEvent(fn func(*CloseEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONCLOSEEVENT,p,fn)
	return
}

func (p *Widget) OnKeyPressEvent(fn func(*KeyEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONKEYPRESSEVENT,p,fn)
	return
}

func (p *Widget) OnKeyReleaseEvent(fn func(*KeyEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONKEYRELEASEEVENT,p,fn)
	return
}

func (p *Widget) OnMousePressEvent(fn func(*MouseEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONMOUSEPRESSEVENT,p,fn)
	return
}

func (p *Widget) OnMouseReleaseEvent(fn func(*MouseEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONMOUSERELEASEEVENT,p,fn)
	return
}

func (p *Widget) OnMouseMoveEvent(fn func(*MouseEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONMOUSEMOVEEVENT,p,fn)
	return
}

func (p *Widget) OnMouseDoubleClickEvent(fn func(*MouseEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONMOUSEDOUBLECLICKEVENT,p,fn)
	return
}

func (p *Widget) OnMoveEvent(fn func(*MoveEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONMOVEEVENT,p,fn)
	return
}

func (p *Widget) OnPaintEvent(fn func(*PaintEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONPAINTEVENT,p,fn)
	return
}

func (p *Widget) OnResizeEvent(fn func(*ResizeEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONRESIZEEVENT,p,fn)
	return
}

func (p *Widget) OnEnterEvent(fn func(*EnterEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONENTEREVENT,p,fn)
	return
}

func (p *Widget) OnLeaveEvent(fn func(*LeaveEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONLEAVEEVENT,p,fn)
	return
}

func (p *Widget) OnFocusInEvent(fn func(*FocusEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONFOCUSINEVENT,p,fn)
	return
}

func (p *Widget) OnFocusOutEvent(fn func(*FocusEvent)) {
	drv_event(DRVCLASS_WIDGET,WIDGET_ONFOCUSOUTEVENT,p,fn)
	return
}

func (p *Widget) Show() {
	p.SetVisible(true)
}

func (p *Widget) Hide() {
	p.SetVisible(false)
}

func (p *Widget) SetPosv(x,y int) {
	p.SetPos(Pt(x,y))
}

func (p *Widget) Posv() (int,int) {
	return p.Pos().Unpack()
}

func (p *Widget) SetSizev(width,height int) {
	p.SetSize(Sz(width,height))
}

func (p *Widget) Sizev() (int,int) {
	return p.Size().Unpack()
}

func (p *Widget) SetGeometryv(x,y,width,height int) {
	p.SetGeometry(Rc(x,y,width,height))
}

func (p *Widget) Geometryv() (int,int,int,int) {
	return p.Geometry().Unpack()
}


// MenuBar struct
type MenuBar struct {
	Widget
}

func (p *MenuBar) Name() string {
	return "MenuBar"
}

func (p *MenuBar) String() string {
	return itemString(p)
}
func (p *MenuBar) SetAttr(attr string, value interface{}) bool {
	switch attr {
	default:
		return p.Widget.SetAttr(attr,value)
	}
	return false
}
func (p *MenuBar) Attr(attr string) interface{} {
	switch attr {
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewMenuBar() *MenuBar{
	return new(MenuBar).Init()
}

func (p *MenuBar) Init() *MenuBar {
	p.classid = DRVCLASS_MENUBAR
	drv(DRVCLASS_MENUBAR,MENUBAR_INIT,p)
	AddItem(p)
	return p
}

func (p *MenuBar) AddMenu(menu MenuItem) {
	drv(DRVCLASS_MENUBAR,MENUBAR_ADDMENU,p,menu)
	return
}

// Menu struct
type Menu struct {
	Widget
}

func (p *Menu) Name() string {
	return "Menu"
}

func (p *Menu) String() string {
	return itemString(p)
}
func (p *Menu) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "title":
		if v, ok := value.(string); ok {
			p.SetTitle(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr,value)
	}
	return false
}
func (p *Menu) Attr(attr string) interface{} {
	switch attr {
	case "title":
		return p.Title()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewMenu() *Menu{
	return new(Menu).Init()
}

func (p *Menu) Init() *Menu {
	p.classid = DRVCLASS_MENU
	drv(DRVCLASS_MENU,MENU_INIT,p)
	AddItem(p)
	return p
}

func (p *Menu) SetTitle(text string) {
	drv(DRVCLASS_MENU,MENU_SETTITLE,p,text)
	return
}

func (p *Menu) Title()(text string) {
	var sh_text StringHead
	drv(DRVCLASS_MENU,MENU_TITLE,p,&sh_text)
	text = sh_text.String()
	return
}

func (p *Menu) AddAction(act *Action) {
	drv(DRVCLASS_MENU,MENU_ADDACTION,p,act)
	return
}

// Action struct
type Action struct {
	object
}

func (p *Action) Name() string {
	return "Action"
}

func (p *Action) String() string {
	return itemString(p)
}
func (p *Action) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "text":
		if v, ok := value.(string); ok {
			p.SetText(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr,value)
	}
	return false
}
func (p *Action) Attr(attr string) interface{} {
	switch attr {
	case "text":
		return p.Text()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewAction() *Action{
	return new(Action).Init()
}

func (p *Action) Init() *Action {
	p.classid = DRVCLASS_ACTION
	drv(DRVCLASS_ACTION,ACTION_INIT,p)
	AddItem(p)
	return p
}

func (p *Action) SetText(text string) {
	drv(DRVCLASS_ACTION,ACTION_SETTEXT,p,text)
	return
}

func (p *Action) Text()(text string) {
	var sh_text StringHead
	drv(DRVCLASS_ACTION,ACTION_TEXT,p,&sh_text)
	text = sh_text.String()
	return
}

func (p *Action) OnTriggered(fn func(bool)) {
	drv_event(DRVCLASS_ACTION,ACTION_ONTRIGGERED,p,fn)
	return
}

// TabWidget struct
type TabWidget struct {
	Widget
}

func (p *TabWidget) Name() string {
	return "TabWidget"
}

func (p *TabWidget) String() string {
	return itemString(p)
}
func (p *TabWidget) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "currentindex":
		if v, ok := value.(int); ok {
			p.SetCurrentIndex(v)
			return true
		}
		return false
	case "currentwidget":
		if v, ok := value.(WidgetItem); ok {
			p.SetCurrentWidget(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr,value)
	}
	return false
}
func (p *TabWidget) Attr(attr string) interface{} {
	switch attr {
	case "count":
		return p.Count()
	case "currentindex":
		return p.CurrentIndex()
	case "currentwidget":
		return p.CurrentWidget()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewTabWidget() *TabWidget{
	return new(TabWidget).Init()
}

func (p *TabWidget) Init() *TabWidget {
	p.classid = DRVCLASS_TABWIDGET
	drv(DRVCLASS_TABWIDGET,TABWIDGET_INIT,p)
	AddItem(p)
	return p
}

func (p *TabWidget) AddTab(w WidgetItem,label string) {
	drv(DRVCLASS_TABWIDGET,TABWIDGET_ADDTAB,p,w,label)
	return
}

func (p *TabWidget) Clear() {
	drv(DRVCLASS_TABWIDGET,TABWIDGET_CLEAR,p)
	return
}

func (p *TabWidget) Count()(n int) {
	drv(DRVCLASS_TABWIDGET,TABWIDGET_COUNT,p,&n)
	return
}

func (p *TabWidget) CurrentIndex()(index int) {
	drv(DRVCLASS_TABWIDGET,TABWIDGET_CURRENTINDEX,p,&index)
	return
}

func (p *TabWidget) CurrentWidget()(w WidgetItem) {
	var hh_w ItemHead
	drv(DRVCLASS_TABWIDGET,TABWIDGET_CURRENTWIDGET,p,&hh_w)
	if hh_w.native != 0 {
		item := FindItem(hh_w.native)
		if item != nil {
			w = item.(WidgetItem)
		}
	}
	return
}

func (p *TabWidget) SetCurrentIndex(index int) {
	drv(DRVCLASS_TABWIDGET,TABWIDGET_SETCURRENTINDEX,p,index)
	return
}

func (p *TabWidget) SetCurrentWidget(w WidgetItem) {
	drv(DRVCLASS_TABWIDGET,TABWIDGET_SETCURRENTWIDGET,p,w)
	return
}

func (p *TabWidget) IndexOf(w WidgetItem)(index int) {
	drv(DRVCLASS_TABWIDGET,TABWIDGET_INDEXOF,p,w,&index)
	return
}

func (p *TabWidget) InsertTab(index int,w WidgetItem,label string) {
	drv(DRVCLASS_TABWIDGET,TABWIDGET_INSERTTAB,p,index,w,label)
	return
}

func (p *TabWidget) RemoveTab(index int) {
	drv(DRVCLASS_TABWIDGET,TABWIDGET_REMOVETAB,p,index)
	return
}

func (p *TabWidget) SetTabText(index int,label string) {
	drv(DRVCLASS_TABWIDGET,TABWIDGET_SETTABTEXT,p,index,label)
	return
}

func (p *TabWidget) SetTabToolTip(index int,tip string) {
	drv(DRVCLASS_TABWIDGET,TABWIDGET_SETTABTOOLTIP,p,index,tip)
	return
}

func (p *TabWidget) TabText(index int)(label string) {
	var sh_label StringHead
	drv(DRVCLASS_TABWIDGET,TABWIDGET_TABTEXT,p,index,&sh_label)
	label = sh_label.String()
	return
}

func (p *TabWidget) TabToolTip(index int)(tip string) {
	var sh_tip StringHead
	drv(DRVCLASS_TABWIDGET,TABWIDGET_TABTOOLTIP,p,index,&sh_tip)
	tip = sh_tip.String()
	return
}

func (p *TabWidget) WidgetOf(index int)(w WidgetItem) {
	var hh_w ItemHead
	drv(DRVCLASS_TABWIDGET,TABWIDGET_WIDGETOF,p,index,&hh_w)
	if hh_w.native != 0 {
		item := FindItem(hh_w.native)
		if item != nil {
			w = item.(WidgetItem)
		}
	}
	return
}

func (p *TabWidget) OnCurrentChanged(fn func(int)) {
	drv_event(DRVCLASS_TABWIDGET,TABWIDGET_ONCURRENTCHANGED,p,fn)
	return
}

// layout struct
type layout struct {
	ItemHead
}

func (p *layout) Name() string {
	return "layout"
}

func (p *layout) String() string {
	return itemString(p)
}
func (p *layout) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "spacing":
		if v, ok := value.(int); ok {
			p.SetSpacing(v)
			return true
		}
		return false
	case "margins":
		if v, ok := value.(Margins); ok {
			p.SetMargins(v)
			return true
		}
		return false
	case "menubar":
		if v, ok := value.(WidgetItem); ok {
			p.SetMenuBar(v)
			return true
		}
		return false
	default:
		return p.ItemHead.SetAttr(attr,value)
	}
	return false
}
func (p *layout) Attr(attr string) interface{} {
	switch attr {
	case "spacing":
		return p.Spacing()
	case "margins":
		return p.Margins()
	case "menubar":
		return p.MenuBar()
	default:
		return p.ItemHead.Attr(attr)
	}
	return nil
}
func (p *layout) AddLayout(layout LayoutItem) {
	drv(DRVCLASS_LAYOUT,LAYOUT_ADDLAYOUT,p,layout)
	return
}

func (p *layout) AddWidget(widget WidgetItem) {
	drv(DRVCLASS_LAYOUT,LAYOUT_ADDWIDGET,p,widget)
	return
}

func (p *layout) SetSpacing(spac int) {
	drv(DRVCLASS_LAYOUT,LAYOUT_SETSPACING,p,spac)
	return
}

func (p *layout) Spacing()(spac int) {
	drv(DRVCLASS_LAYOUT,LAYOUT_SPACING,p,&spac)
	return
}

func (p *layout) SetMargins(m Margins) {
	drv(DRVCLASS_LAYOUT,LAYOUT_SETMARGINS,p,m)
	return
}

func (p *layout) Margins()(m Margins) {
	drv(DRVCLASS_LAYOUT,LAYOUT_MARGINS,p,&m)
	return
}

func (p *layout) SetMenuBar(widget WidgetItem) {
	drv(DRVCLASS_LAYOUT,LAYOUT_SETMENUBAR,p,widget)
	return
}

func (p *layout) MenuBar()(widget WidgetItem) {
	var hh_widget ItemHead
	drv(DRVCLASS_LAYOUT,LAYOUT_MENUBAR,p,&hh_widget)
	if hh_widget.native != 0 {
		item := FindItem(hh_widget.native)
		if item != nil {
			widget = item.(WidgetItem)
		}
	}
	return
}

func (p *layout) SetMargin(v int) {
	p.SetMargins(Mr(v,v,v,v))
}
func (p *layout) SetMarginsv(left,top,right,bottom int) {
	p.SetMargins(Mr(left,top,right,bottom))
}

func (p *layout) Marginsv() (int,int,int,int) {
	return p.Margins().Unpack()
}

// boxLayout struct
type boxLayout struct {
	layout
}

func (p *boxLayout) Name() string {
	return "boxLayout"
}

func (p *boxLayout) String() string {
	return itemString(p)
}
func (p *boxLayout) SetAttr(attr string, value interface{}) bool {
	switch attr {
	default:
		return p.layout.SetAttr(attr,value)
	}
	return false
}
func (p *boxLayout) Attr(attr string) interface{} {
	switch attr {
	default:
		return p.layout.Attr(attr)
	}
	return nil
}
func (p *boxLayout) AddLayoutWith(layout LayoutItem,stetch int) {
	drv(DRVCLASS_BOXLAYOUT,BOXLAYOUT_ADDLAYOUTWITH,p,layout,stetch)
	return
}

func (p *boxLayout) AddWidgetWith(widget WidgetItem,stretch int,alignment int) {
	drv(DRVCLASS_BOXLAYOUT,BOXLAYOUT_ADDWIDGETWITH,p,widget,stretch,alignment)
	return
}

func (p *boxLayout) AddSpacing(size int) {
	drv(DRVCLASS_BOXLAYOUT,BOXLAYOUT_ADDSPACING,p,size)
	return
}

func (p *boxLayout) AddStretch(size int) {
	drv(DRVCLASS_BOXLAYOUT,BOXLAYOUT_ADDSTRETCH,p,size)
	return
}


// HBoxLayout struct
type HBoxLayout struct {
	boxLayout
}

func (p *HBoxLayout) Name() string {
	return "HBoxLayout"
}

func (p *HBoxLayout) String() string {
	return itemString(p)
}
func (p *HBoxLayout) SetAttr(attr string, value interface{}) bool {
	switch attr {
	default:
		return p.boxLayout.SetAttr(attr,value)
	}
	return false
}
func (p *HBoxLayout) Attr(attr string) interface{} {
	switch attr {
	default:
		return p.boxLayout.Attr(attr)
	}
	return nil
}
func NewHBoxLayout() *HBoxLayout{
	return new(HBoxLayout).Init()
}

func (p *HBoxLayout) Init() *HBoxLayout {
	p.classid = DRVCLASS_HBOXLAYOUT
	drv(DRVCLASS_HBOXLAYOUT,HBOXLAYOUT_INIT,p)
	AddItem(p)
	return p
}

// VBoxLayout struct
type VBoxLayout struct {
	boxLayout
}

func (p *VBoxLayout) Name() string {
	return "VBoxLayout"
}

func (p *VBoxLayout) String() string {
	return itemString(p)
}
func (p *VBoxLayout) SetAttr(attr string, value interface{}) bool {
	switch attr {
	default:
		return p.boxLayout.SetAttr(attr,value)
	}
	return false
}
func (p *VBoxLayout) Attr(attr string) interface{} {
	switch attr {
	default:
		return p.boxLayout.Attr(attr)
	}
	return nil
}
func NewVBoxLayout() *VBoxLayout{
	return new(VBoxLayout).Init()
}

func (p *VBoxLayout) Init() *VBoxLayout {
	p.classid = DRVCLASS_VBOXLAYOUT
	drv(DRVCLASS_VBOXLAYOUT,VBOXLAYOUT_INIT,p)
	AddItem(p)
	return p
}

// baseButton struct
type baseButton struct {
	Widget
}

func (p *baseButton) Name() string {
	return "baseButton"
}

func (p *baseButton) String() string {
	return itemString(p)
}
func (p *baseButton) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "text":
		if v, ok := value.(string); ok {
			p.SetText(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr,value)
	}
	return false
}
func (p *baseButton) Attr(attr string) interface{} {
	switch attr {
	case "text":
		return p.Text()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func (p *baseButton) SetText(text string) {
	drv(DRVCLASS_BASEBUTTON,BASEBUTTON_SETTEXT,p,text)
	return
}

func (p *baseButton) Text()(text string) {
	var sh_text StringHead
	drv(DRVCLASS_BASEBUTTON,BASEBUTTON_TEXT,p,&sh_text)
	text = sh_text.String()
	return
}

// Button struct
type Button struct {
	baseButton
}

func (p *Button) Name() string {
	return "Button"
}

func (p *Button) String() string {
	return itemString(p)
}
func (p *Button) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "flat":
		if v, ok := value.(bool); ok {
			p.SetFlat(v)
			return true
		}
		return false
	case "default":
		if v, ok := value.(bool); ok {
			p.SetDefault(v)
			return true
		}
		return false
	case "menu":
		if v, ok := value.(MenuItem); ok {
			p.SetMenu(v)
			return true
		}
		return false
	default:
		return p.baseButton.SetAttr(attr,value)
	}
	return false
}
func (p *Button) Attr(attr string) interface{} {
	switch attr {
	case "flat":
		return p.IsFlat()
	case "default":
		return p.IsDefault()
	case "menu":
		return p.Menu()
	default:
		return p.baseButton.Attr(attr)
	}
	return nil
}
func NewButton() *Button{
	return new(Button).Init()
}

func (p *Button) Init() *Button {
	p.classid = DRVCLASS_BUTTON
	drv(DRVCLASS_BUTTON,BUTTON_INIT,p)
	AddItem(p)
	return p
}

func (p *Button) SetFlat(b bool) {
	drv(DRVCLASS_BUTTON,BUTTON_SETFLAT,p,b)
	return
}

func (p *Button) IsFlat()(b bool) {
	var b_b int
	drv(DRVCLASS_BUTTON,BUTTON_ISFLAT,p,&b_b)
	b = b_b != 0
	return
}

func (p *Button) SetDefault(b bool) {
	drv(DRVCLASS_BUTTON,BUTTON_SETDEFAULT,p,b)
	return
}

func (p *Button) IsDefault()(b bool) {
	var b_b int
	drv(DRVCLASS_BUTTON,BUTTON_ISDEFAULT,p,&b_b)
	b = b_b != 0
	return
}

func (p *Button) SetMenu(menu MenuItem) {
	drv(DRVCLASS_BUTTON,BUTTON_SETMENU,p,menu)
	return
}

func (p *Button) Menu()(menu MenuItem) {
	var hh_menu ItemHead
	drv(DRVCLASS_BUTTON,BUTTON_MENU,p,&hh_menu)
	if hh_menu.native != 0 {
		item := FindItem(hh_menu.native)
		if item != nil {
			menu = item.(MenuItem)
		}
	}
	return
}

func (p *Button) OnClicked(fn func()) {
	drv_event(DRVCLASS_BUTTON,BUTTON_ONCLICKED,p,fn)
	return
}

// CheckBox struct
type CheckBox struct {
	baseButton
}

func (p *CheckBox) Name() string {
	return "CheckBox"
}

func (p *CheckBox) String() string {
	return itemString(p)
}
func (p *CheckBox) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "check":
		if v, ok := value.(int); ok {
			p.SetCheck(v)
			return true
		}
		return false
	case "tristate":
		if v, ok := value.(bool); ok {
			p.SetTristate(v)
			return true
		}
		return false
	default:
		return p.baseButton.SetAttr(attr,value)
	}
	return false
}
func (p *CheckBox) Attr(attr string) interface{} {
	switch attr {
	case "check":
		return p.Check()
	case "tristate":
		return p.IsTristate()
	default:
		return p.baseButton.Attr(attr)
	}
	return nil
}
func NewCheckBox() *CheckBox{
	return new(CheckBox).Init()
}

func (p *CheckBox) Init() *CheckBox {
	p.classid = DRVCLASS_CHECKBOX
	drv(DRVCLASS_CHECKBOX,CHECKBOX_INIT,p)
	AddItem(p)
	return p
}

func (p *CheckBox) SetCheck(state int) {
	drv(DRVCLASS_CHECKBOX,CHECKBOX_SETCHECK,p,state)
	return
}

func (p *CheckBox) Check()(state int) {
	drv(DRVCLASS_CHECKBOX,CHECKBOX_CHECK,p,&state)
	return
}

func (p *CheckBox) SetTristate(b bool) {
	drv(DRVCLASS_CHECKBOX,CHECKBOX_SETTRISTATE,p,b)
	return
}

func (p *CheckBox) IsTristate()(b bool) {
	var b_b int
	drv(DRVCLASS_CHECKBOX,CHECKBOX_ISTRISTATE,p,&b_b)
	b = b_b != 0
	return
}

func (p *CheckBox) OnStateChanged(fn func(int)) {
	drv_event(DRVCLASS_CHECKBOX,CHECKBOX_ONSTATECHANGED,p,fn)
	return
}

// Radio struct
type Radio struct {
	baseButton
}

func (p *Radio) Name() string {
	return "Radio"
}

func (p *Radio) String() string {
	return itemString(p)
}
func (p *Radio) SetAttr(attr string, value interface{}) bool {
	switch attr {
	default:
		return p.baseButton.SetAttr(attr,value)
	}
	return false
}
func (p *Radio) Attr(attr string) interface{} {
	switch attr {
	default:
		return p.baseButton.Attr(attr)
	}
	return nil
}
func NewRadio() *Radio{
	return new(Radio).Init()
}

func (p *Radio) Init() *Radio {
	p.classid = DRVCLASS_RADIO
	drv(DRVCLASS_RADIO,RADIO_INIT,p)
	AddItem(p)
	return p
}

func (p *Radio) OnClicked(fn func()) {
	drv_event(DRVCLASS_RADIO,RADIO_ONCLICKED,p,fn)
	return
}

// Frame struct
type Frame struct {
	Widget
}

func (p *Frame) Name() string {
	return "Frame"
}

func (p *Frame) String() string {
	return itemString(p)
}
func (p *Frame) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "framestyle":
		if v, ok := value.(int); ok {
			p.SetFrameStyle(v)
			return true
		}
		return false
	case "framerect":
		if v, ok := value.(Rect); ok {
			p.SetFrameRect(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr,value)
	}
	return false
}
func (p *Frame) Attr(attr string) interface{} {
	switch attr {
	case "framestyle":
		return p.FrameStyle()
	case "framerect":
		return p.FrameRect()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewFrame() *Frame{
	return new(Frame).Init()
}

func (p *Frame) Init() *Frame {
	p.classid = DRVCLASS_FRAME
	drv(DRVCLASS_FRAME,FRAME_INIT,p)
	AddItem(p)
	return p
}

func (p *Frame) SetFrameStyle(style int) {
	drv(DRVCLASS_FRAME,FRAME_SETFRAMESTYLE,p,style)
	return
}

func (p *Frame) FrameStyle()(style int) {
	drv(DRVCLASS_FRAME,FRAME_FRAMESTYLE,p,&style)
	return
}

func (p *Frame) SetFrameRect(rc Rect) {
	drv(DRVCLASS_FRAME,FRAME_SETFRAMERECT,p,rc)
	return
}

func (p *Frame) FrameRect()(rc Rect) {
	drv(DRVCLASS_FRAME,FRAME_FRAMERECT,p,&rc)
	return
}

func (p *Frame) SetFrameRectv(x,y,width,height int) {
	p.SetFrameRect(Rc(x,y,width,height))
}

func (p *Frame) FrameRectv() (int,int,int,int) {
	return p.FrameRect().Unpack()
}

// Label struct
type Label struct {
	Frame
}

func (p *Label) Name() string {
	return "Label"
}

func (p *Label) String() string {
	return itemString(p)
}
func (p *Label) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "text":
		if v, ok := value.(string); ok {
			p.SetText(v)
			return true
		}
		return false
	case "wordwrap":
		if v, ok := value.(bool); ok {
			p.SetWordWrap(v)
			return true
		}
		return false
	case "textformat":
		if v, ok := value.(int); ok {
			p.SetTextFormat(v)
			return true
		}
		return false
	default:
		return p.Frame.SetAttr(attr,value)
	}
	return false
}
func (p *Label) Attr(attr string) interface{} {
	switch attr {
	case "text":
		return p.Text()
	case "wordwrap":
		return p.WordWrap()
	case "textformat":
		return p.TextFormat()
	default:
		return p.Frame.Attr(attr)
	}
	return nil
}
func NewLabel() *Label{
	return new(Label).Init()
}

func (p *Label) Init() *Label {
	p.classid = DRVCLASS_LABEL
	drv(DRVCLASS_LABEL,LABEL_INIT,p)
	AddItem(p)
	return p
}

func (p *Label) SetText(text string) {
	drv(DRVCLASS_LABEL,LABEL_SETTEXT,p,text)
	return
}

func (p *Label) Text()(text string) {
	var sh_text StringHead
	drv(DRVCLASS_LABEL,LABEL_TEXT,p,&sh_text)
	text = sh_text.String()
	return
}

func (p *Label) SetWordWrap(b bool) {
	drv(DRVCLASS_LABEL,LABEL_SETWORDWRAP,p,b)
	return
}

func (p *Label) WordWrap()(b bool) {
	var b_b int
	drv(DRVCLASS_LABEL,LABEL_WORDWRAP,p,&b_b)
	b = b_b != 0
	return
}

func (p *Label) SetTextFormat(format int) {
	drv(DRVCLASS_LABEL,LABEL_SETTEXTFORMAT,p,format)
	return
}

func (p *Label) TextFormat()(format int) {
	drv(DRVCLASS_LABEL,LABEL_TEXTFORMAT,p,&format)
	return
}

func (p *Label) OnLinkActivated(fn func(string)) {
	drv_event(DRVCLASS_LABEL,LABEL_ONLINKACTIVATED,p,fn)
	return
}

// GroupBox struct
type GroupBox struct {
	Widget
}

func (p *GroupBox) Name() string {
	return "GroupBox"
}

func (p *GroupBox) String() string {
	return itemString(p)
}
func (p *GroupBox) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "title":
		if v, ok := value.(string); ok {
			p.SetTitle(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr,value)
	}
	return false
}
func (p *GroupBox) Attr(attr string) interface{} {
	switch attr {
	case "title":
		return p.Title()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewGroupBox() *GroupBox{
	return new(GroupBox).Init()
}

func (p *GroupBox) Init() *GroupBox {
	p.classid = DRVCLASS_GROUPBOX
	drv(DRVCLASS_GROUPBOX,GROUPBOX_INIT,p)
	AddItem(p)
	return p
}

func (p *GroupBox) SetTitle(text string) {
	drv(DRVCLASS_GROUPBOX,GROUPBOX_SETTITLE,p,text)
	return
}

func (p *GroupBox) Title()(text string) {
	var sh_text StringHead
	drv(DRVCLASS_GROUPBOX,GROUPBOX_TITLE,p,&sh_text)
	text = sh_text.String()
	return
}

// Dialog struct
type Dialog struct {
	Widget
}

func (p *Dialog) Name() string {
	return "Dialog"
}

func (p *Dialog) String() string {
	return itemString(p)
}
func (p *Dialog) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "modal":
		if v, ok := value.(bool); ok {
			p.SetModal(v)
			return true
		}
		return false
	case "result":
		if v, ok := value.(int); ok {
			p.SetResult(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr,value)
	}
	return false
}
func (p *Dialog) Attr(attr string) interface{} {
	switch attr {
	case "modal":
		return p.IsModal()
	case "result":
		return p.Result()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewDialog() *Dialog{
	return new(Dialog).Init()
}

func (p *Dialog) Init() *Dialog {
	p.classid = DRVCLASS_DIALOG
	drv(DRVCLASS_DIALOG,DIALOG_INIT,p)
	AddItem(p)
	return p
}

func (p *Dialog) SetModal(b bool) {
	drv(DRVCLASS_DIALOG,DIALOG_SETMODAL,p,b)
	return
}

func (p *Dialog) IsModal()(b bool) {
	var b_b int
	drv(DRVCLASS_DIALOG,DIALOG_ISMODAL,p,&b_b)
	b = b_b != 0
	return
}

func (p *Dialog) SetResult(r int) {
	drv(DRVCLASS_DIALOG,DIALOG_SETRESULT,p,r)
	return
}

func (p *Dialog) Result()(r int) {
	drv(DRVCLASS_DIALOG,DIALOG_RESULT,p,&r)
	return
}

func (p *Dialog) Exec()(r int) {
	drv(DRVCLASS_DIALOG,DIALOG_EXEC,p,&r)
	return
}

func (p *Dialog) Done(r int) {
	drv(DRVCLASS_DIALOG,DIALOG_DONE,p,r)
	return
}

func (p *Dialog) Accept() {
	drv(DRVCLASS_DIALOG,DIALOG_ACCEPT,p)
	return
}

func (p *Dialog) Reject() {
	drv(DRVCLASS_DIALOG,DIALOG_REJECT,p)
	return
}

func (p *Dialog) OnAccepted(fn func()) {
	drv_event(DRVCLASS_DIALOG,DIALOG_ONACCEPTED,p,fn)
	return
}

func (p *Dialog) OnRejected(fn func()) {
	drv_event(DRVCLASS_DIALOG,DIALOG_ONREJECTED,p,fn)
	return
}

// ComboBox struct
type ComboBox struct {
	Widget
}

func (p *ComboBox) Name() string {
	return "ComboBox"
}

func (p *ComboBox) String() string {
	return itemString(p)
}
func (p *ComboBox) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "currentindex":
		if v, ok := value.(int); ok {
			p.SetCurrentIndex(v)
			return true
		}
		return false
	case "editable":
		if v, ok := value.(bool); ok {
			p.SetEditable(v)
			return true
		}
		return false
	case "maxcount":
		if v, ok := value.(int); ok {
			p.SetMaxCount(v)
			return true
		}
		return false
	case "maxvisibleitems":
		if v, ok := value.(int); ok {
			p.SetMaxVisibleItems(v)
			return true
		}
		return false
	case "minimumcontentslength":
		if v, ok := value.(int); ok {
			p.SetMinimumContentsLength(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr,value)
	}
	return false
}
func (p *ComboBox) Attr(attr string) interface{} {
	switch attr {
	case "count":
		return p.Count()
	case "currentindex":
		return p.CurrentIndex()
	case "currenttext":
		return p.CurrentText()
	case "editable":
		return p.IsEditable()
	case "maxcount":
		return p.MaxCount()
	case "maxvisibleitems":
		return p.MaxVisibleItems()
	case "minimuncontentslenght":
		return p.MinimunContentsLenght()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewComboBox() *ComboBox{
	return new(ComboBox).Init()
}

func (p *ComboBox) Init() *ComboBox {
	p.classid = DRVCLASS_COMBOBOX
	drv(DRVCLASS_COMBOBOX,COMBOBOX_INIT,p)
	AddItem(p)
	return p
}

func (p *ComboBox) Count()(count int) {
	drv(DRVCLASS_COMBOBOX,COMBOBOX_COUNT,p,&count)
	return
}

func (p *ComboBox) SetCurrentIndex(index int) {
	drv(DRVCLASS_COMBOBOX,COMBOBOX_SETCURRENTINDEX,p,index)
	return
}

func (p *ComboBox) CurrentIndex()(index int) {
	drv(DRVCLASS_COMBOBOX,COMBOBOX_CURRENTINDEX,p,&index)
	return
}

func (p *ComboBox) CurrentText()(text string) {
	var sh_text StringHead
	drv(DRVCLASS_COMBOBOX,COMBOBOX_CURRENTTEXT,p,&sh_text)
	text = sh_text.String()
	return
}

func (p *ComboBox) SetEditable(b bool) {
	drv(DRVCLASS_COMBOBOX,COMBOBOX_SETEDITABLE,p,b)
	return
}

func (p *ComboBox) IsEditable()(b bool) {
	var b_b int
	drv(DRVCLASS_COMBOBOX,COMBOBOX_ISEDITABLE,p,&b_b)
	b = b_b != 0
	return
}

func (p *ComboBox) SetMaxCount(count int) {
	drv(DRVCLASS_COMBOBOX,COMBOBOX_SETMAXCOUNT,p,count)
	return
}

func (p *ComboBox) MaxCount()(count int) {
	drv(DRVCLASS_COMBOBOX,COMBOBOX_MAXCOUNT,p,&count)
	return
}

func (p *ComboBox) SetMaxVisibleItems(max int) {
	drv(DRVCLASS_COMBOBOX,COMBOBOX_SETMAXVISIBLEITEMS,p,max)
	return
}

func (p *ComboBox) MaxVisibleItems()(max int) {
	drv(DRVCLASS_COMBOBOX,COMBOBOX_MAXVISIBLEITEMS,p,&max)
	return
}

func (p *ComboBox) SetMinimumContentsLength(characters int) {
	drv(DRVCLASS_COMBOBOX,COMBOBOX_SETMINIMUMCONTENTSLENGTH,p,characters)
	return
}

func (p *ComboBox) MinimunContentsLenght()(characters int) {
	drv(DRVCLASS_COMBOBOX,COMBOBOX_MINIMUNCONTENTSLENGHT,p,&characters)
	return
}

func (p *ComboBox) AddItem(text string) {
	drv(DRVCLASS_COMBOBOX,COMBOBOX_ADDITEM,p,text)
	return
}

func (p *ComboBox) InsertItem(index int,text string) {
	drv(DRVCLASS_COMBOBOX,COMBOBOX_INSERTITEM,p,index,text)
	return
}

func (p *ComboBox) RemoveItem(index int) {
	drv(DRVCLASS_COMBOBOX,COMBOBOX_REMOVEITEM,p,index)
	return
}

func (p *ComboBox) ItemText(index int)(text string) {
	var sh_text StringHead
	drv(DRVCLASS_COMBOBOX,COMBOBOX_ITEMTEXT,p,index,&sh_text)
	text = sh_text.String()
	return
}

func (p *ComboBox) OnCurrentIndexChanged(fn func(int)) {
	drv_event(DRVCLASS_COMBOBOX,COMBOBOX_ONCURRENTINDEXCHANGED,p,fn)
	return
}

// Painter struct
type Painter struct {
	ItemHead
}

func (p *Painter) Name() string {
	return "Painter"
}

func (p *Painter) String() string {
	return itemString(p)
}
func (p *Painter) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "font":
		if v, ok := value.(*FontInfo); ok {
			p.SetFont(v)
			return true
		}
		return false
	default:
		return p.ItemHead.SetAttr(attr,value)
	}
	return false
}
func (p *Painter) Attr(attr string) interface{} {
	switch attr {
	case "font":
		return p.Font()
	default:
		return p.ItemHead.Attr(attr)
	}
	return nil
}
func NewPainter() *Painter{
	return new(Painter).Init()
}

func (p *Painter) Init() *Painter {
	p.classid = DRVCLASS_PAINTER
	drv(DRVCLASS_PAINTER,PAINTER_INIT,p)
	AddItem(p)
	return p
}

func (p *Painter) destroy() {
	drv(DRVCLASS_PAINTER,PAINTER_DESTROY,p)
	return
}

func (p *Painter) Begin(w WidgetItem) {
	drv(DRVCLASS_PAINTER,PAINTER_BEGIN,p,w)
	return
}

func (p *Painter) End() {
	drv(DRVCLASS_PAINTER,PAINTER_END,p)
	return
}

func (p *Painter) SetFont(font *FontInfo) {
	drv(DRVCLASS_PAINTER,PAINTER_SETFONT,p,font)
	return
}

func (p *Painter) Font()(font *FontInfo) {
	var fh_font font_head
	drv(DRVCLASS_PAINTER,PAINTER_FONT,p,&fh_font)
	font = fh_font.Font()
	return
}

func (p *Painter) DrawPoint(pt Point) {
	drv(DRVCLASS_PAINTER,PAINTER_DRAWPOINT,p,pt)
	return
}

func (p *Painter) DrawLine(pt1 Point,pt2 Point) {
	drv(DRVCLASS_PAINTER,PAINTER_DRAWLINE,p,pt1,pt2)
	return
}

func (p *Painter) DrawLines(pts []Point) {
	drv(DRVCLASS_PAINTER,PAINTER_DRAWLINES,p,pts)
	return
}

func (p *Painter) DrawPolyline(pts []Point) {
	drv(DRVCLASS_PAINTER,PAINTER_DRAWPOLYLINE,p,pts)
	return
}

func (p *Painter) DrawText(pt Point,text string) {
	drv(DRVCLASS_PAINTER,PAINTER_DRAWTEXT,p,pt,text)
	return
}


func (p *Painter) Destroy() {	
	native := p.native
	p.destroy()
	RemoveItem(native)
}


