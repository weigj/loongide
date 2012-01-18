module("layout")

name = "layout"
base = "ItemHead"

funcs = [[
AddLayout(layout LayoutItem)
AddWidget(widget WidgetItem)
@ SetSpacing(spac int)
@ Spacing() (spac int)
@ SetMargins(m Margins)
@ Margins() (m Margins)
@ SetMenuBar(widget WidgetItem)
@ MenuBar() (widget WidgetItem)
]]

expands = [[
func (p *layout) SetMargin(v int) {
	p.SetMargins(Mr(v,v,v,v))
}
func (p *layout) SetMarginsv(left,top,right,bottom int) {
	p.SetMargins(Mr(left,top,right,bottom))
}

func (p *layout) Marginsv() (int,int,int,int) {
	return p.Margins().Unpack()
}
]]

qtdrv = {
inc = "<QLayout>",
name = "QLayout",

AddLayout = "addItem",
AddWidget = "addWidget",
SetSpacing = "setSpacing",
Spacing = "spacing",
SetMargins = "setContentsMargins",
Margins = "contentsMargins",
SetMenuBar = "setMenuBar",
MenuBar = "menuBar",
}
