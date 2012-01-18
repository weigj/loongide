module("vboxlayout")

name = "VBoxLayout"
base = "boxLayout"

funcs = [[
+Init()
]]

qtdrv = {
inc = "<QVBoxLayout>",	
name = "QVBoxLayout",

Init = [[
drvNewObj(a0,new QVBoxLayout);
]],
}