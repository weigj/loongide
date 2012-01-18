module("painter")

name = "Painter"
base = "ItemHead"

funcs = [[
+ Init()
- destroy()
Begin(w WidgetItem)
End()
@ SetFont(font *FontInfo)
@ Font() (font *FontInfo)
DrawPoint(pt Point)
DrawLine(pt1 Point, pt2 Point)
DrawLines(pts []Point)
DrawPolyline(pts []Point)
DrawText(pt Point,text string)
]]

expands = [[

func (p *Painter) Destroy() {	
	native := p.native
	p.destroy()
	RemoveItem(native)
}

]]

qtdrv = {
inc = "<QPainter>",
name = "QPainter",

Init = [[
drvNewHead(a0,new QPainter);
]],

destroy = [[
drvDelObj(a0,self);
]],

Begin = "begin",
End = "end",
SetFont = "setFont",
Font = "font",
DrawPoint = "drawPoint",
DrawLine = "drawLine",
DrawLines = "drawLines",
DrawPolyline = "drawPolyline",
DrawText = "drawText",

}