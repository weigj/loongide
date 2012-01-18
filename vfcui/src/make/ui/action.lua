module("action")

name = "Action"
base = "object"

funcs = [[
+ Init()
@ SetText(text string)
@ Text() (text string)
* OnTriggered(fn func(bool))
]]

qtdrv = {
inc = "<QAction>",
name = "QAction",

Init = [[
drvNewObj(a0, new QAction(qApp));
]],

SetText = "setText",
Text = "text",

OnTriggered = [[
QObject::connect(self,SIGNAL(triggered(bool)),drvNewSignal(self,a1,a2),SLOT(call(bool)));
]]

}
