module("menu")

name = "Menu"
base = "Widget"

funcs = [[
+ Init()
@ SetTitle(text string)
@ Title() (text string)
AddAction(act *Action)
]]


qtdrv = {
inc = "<QMenu>",
name = "QMenu",

Init = [[
drvNewObj(a0,new QMenu);
]],

SetTitle = "setTitle",
Title = "title",
AddAction = "addAction",

}
