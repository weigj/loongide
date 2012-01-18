module("boxlayout")

name = "boxLayout"
base = "layout"

funcs = [[
AddLayoutWith(layout LayoutItem, stetch int)
AddWidgetWith(widget WidgetItem, stretch int, alignment int)
AddSpacing(size int)
AddStretch(size int)
]]

expands = [[
]]

qtdrv = {
inc = "<QBoxLayout>",
name = "QBoxLayout",

AddLayoutWith = [[
self->addLayout((QLayout*)drvGetNative(a1),drvGetInt(a2));
]],
AddWidgetWith = [[
self->addWidget((QWidget*)drvGetNative(a1),drvGetInt(a2),Qt::Alignment(drvGetInt(a3)));
]],
AddSpacing = "addSpacing",
AddStretch = "addStretch",
}
