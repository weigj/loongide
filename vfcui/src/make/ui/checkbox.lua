module("checkbox")

name = "CheckBox"
base = "baseButton"

funcs = [[
+Init()
@ SetCheck(state int)
@ Check() (state int)
@ SetTristate(b bool)
@ IsTristate() (b bool)

* OnStateChanged(fn func(int))
]]


qtdrv = {
inc = "<QCheckBox>",
name = "QCheckBox",

Init = [[
drvNewObj(a0,new QCheckBox);
]],

SetCheck = [[
self->setCheckState((Qt::CheckState)drvGetInt(a1));
]],

Check = [[
drvSetInt(a1,self->checkState());
]],

SetTristate = "setTristate",
IsTristate = "isTristate",

OnStateChanged = [[
QObject::connect(self,SIGNAL(stateChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
]]

}
