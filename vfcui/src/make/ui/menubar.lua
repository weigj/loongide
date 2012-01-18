module("menubar")

name = "MenuBar"
base = "Widget"

funcs = [[
+ Init()
AddMenu(menu MenuItem)
]]


qtdrv = {
inc = "<QMenuBar>",
name = "QMenuBar",

Init = [[
drvNewObj(a0,new QMenuBar);
]],

AddMenu = "addMenu",

}
