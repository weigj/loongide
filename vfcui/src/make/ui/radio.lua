module("radio")

name = "Radio"
base = "baseButton"

funcs = [[
+Init()
* OnClicked(fn func())
]]


qtdrv = {
inc = "<QRadioButton>",
name = "QRadioButton",

Init = [[
drvNewObj(a0,new QRadioButton);
]],

OnClicked = [[
QObject::connect(self,SIGNAL(clicked()),drvNewSignal(self,a1,a2),SLOT(call()));
]],

}
