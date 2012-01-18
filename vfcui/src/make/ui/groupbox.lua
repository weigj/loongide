module("groupbox")

name = "GroupBox"
base = "Widget"

funcs = [[
+ Init()
@ SetTitle(text string)
@ Title() (text string)
]]


qtdrv = {
inc = "<QGroupBox>",
name = "QGroupBox",

Init = [[
drvNewObj(a0,new QGroupBox);
]],

SetTitle = "setTitle",
Title = "title",

}
