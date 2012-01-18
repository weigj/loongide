module("widget")

name = "Widget"
base = "object"

defs = [[
]]

funcs = [[
+ Init()
- Destroy()
@ SetAutoDestroy(b bool)
@ AutoDestroy() (b bool)
@ SetLayout(layout LayoutItem)
@ Layout() (layout LayoutItem)
@ SetParent(parent WidgetItem)
@ Parent() (parent WidgetItem)
@ SetVisible(b bool)
@ IsVisible() (b bool)
@ SetWindowTitle(text string)
@ WindowTitle()(text string)
@ SetPos(pt Point)
@ Pos() (pt Point)
@ SetSize(sz Size)
@ Size() (sz Size)
@ SetGeometry(rc Rect)
@ Geometry()(rc Rect)
@ SetFont(font *FontInfo)
@ Font() (font *FontInfo)
Close()
Update()
Repaint()
* OnShowEvent(fn func(*ShowEvent))
* OnHideEvent(fn func(*HideEvent))
* OnCloseEvent(fn func(*CloseEvent))
* OnKeyPressEvent(fn func(*KeyEvent))
* OnKeyReleaseEvent(fn func(*KeyEvent))
* OnMousePressEvent(fn func(*MouseEvent))
* OnMouseReleaseEvent(fn func(*MouseEvent))
* OnMouseMoveEvent(fn func(*MouseEvent))
* OnMouseDoubleClickEvent(fn func(*MouseEvent))
* OnMoveEvent(fn func(*MoveEvent))
* OnPaintEvent(fn func(*PaintEvent))
* OnResizeEvent(fn func(*ResizeEvent))
* OnEnterEvent(fn func(*EnterEvent))
* OnLeaveEvent(fn func(*LeaveEvent))
* OnFocusInEvent(fn func(*FocusEvent))
* OnFocusOutEvent(fn func(*FocusEvent))
]]

expands = [[
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

]]

qtdrv = {
inc = "<QWidget>",
name = "QWidget",

Init = [[
drvNewObj(a0,new QWidget);
]],
Destroy = [[
drvDelObj(a0,self);
]],

SetAutoDestroy = [[
self->setAttribute(Qt::WA_DeleteOnClose,drvGetBool(a1));
]],
AutoDestroy = [[
drvSetBool(a1,self->testAttribute(Qt::WA_DeleteOnClose));
]],
SetLayout = "setLayout",
Layout = "layout",
SetVisible = "setVisible",
IsVisible = "isVisible",
SetParent = "setParent",
Parent = "parent",
SetWindowTitle = "setWindowTitle",
WindowTitle = "windowTitle",
SetPos = "move",
Pos = "pos",
SetSize = "resize",
Size = "size",
SetGeometry = "setGeometry",
Geometry = "geometry",
SetFont = "setFont",
Font = "font", 
Close = "close",
Update = "update",
Repaint = "repaint",

OnShowEvent = [[
drvNewEvent(QEvent::Show,a0,a1,a2);
]],

OnHideEvent = [[
drvNewEvent(QEvent::Hide,a0,a1,a2);
]],

OnCloseEvent = [[
drvNewEvent(QEvent::Close,a0,a1,a2);
]],

OnKeyPressEvent = [[
drvNewEvent(QEvent::KeyPress,a0,a1,a2);
]],

OnKeyReleaseEvent = [[
drvNewEvent(QEvent::KeyRelease,a0,a1,a2);
]],

OnMousePressEvent = [[
drvNewEvent(QEvent::MouseButtonPress,a0,a1,a2);
]],

OnMouseReleaseEvent = [[
drvNewEvent(QEvent::MouseButtonRelease,a0,a1,a2);
]],

OnMouseMoveEvent = [[
drvNewEvent(QEvent::MouseMove,a0,a1,a2);
]],

OnMouseDoubleClickEvent = [[
drvNewEvent(QEvent::MouseButtonDblClick,a0,a1,a2);
]],

OnMoveEvent = [[
drvNewEvent(QEvent::Move,a0,a1,a2);
]],

OnPaintEvent = [[
drvNewEvent(QEvent::Paint,a0,a1,a2);
]],

OnResizeEvent = [[
drvNewEvent(QEvent::Resize,a0,a1,a2);
]],

OnEnterEvent = [[
drvNewEvent(QEvent::Enter,a0,a1,a2);
]],

OnLeaveEvent = [[
drvNewEvent(QEvent::Leave,a0,a1,a2);
]],

OnFocusInEvent = [[
drvNewEvent(QEvent::FocusIn,a0,a1,a2);
]],

OnFocusOutEvent = [[
drvNewEvent(QEvent::FocusOut,a0,a1,a2);
]],

}
