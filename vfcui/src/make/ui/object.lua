module("object")

name = "object"
base = "ItemHead"

funcs = [[
- Destroy()
StartTimer(millisecond int) (id int)
KillTimer(id int)

*OnTimerEvent(fn func(*TimerEvent))
]]

qtdrv = {
inc = "<QtCore/QObject>",
name = "QObject",

Destroy = [[
drvDelObj(a0,self);
]],

StartTimer = "startTimer",
KillTimer = "killTimer",

OnTimerEvent = [[
drvNewEvent(QEvent::Timer,a0,a1,a2);
]]
}
