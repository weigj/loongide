module("hboxlayout")

name = "HBoxLayout"
base = "boxLayout"


funcs = [[
+Init()
]]

qtdrv = {
inc = "<QHBoxLayout>",	
name = "QHBoxLayout",

Init = [[
drvNewObj(a0,new QHBoxLayout);
]],

}