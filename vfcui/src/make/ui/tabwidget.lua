module("tabwidget")

name = "TabWidget"
base = "Widget"

funcs = [[
+Init()
AddTab(w WidgetItem, label string)
Clear()
@ Count() (n int)
@ CurrentIndex() (index int)
@ CurrentWidget() (w WidgetItem)
@ SetCurrentIndex(index int)
@ SetCurrentWidget(w WidgetItem)
IndexOf(w WidgetItem) (index int)
InsertTab(index int, w WidgetItem, label string)
RemoveTab(index int)
SetTabText(index int, label string)
SetTabToolTip(index int, tip string)
TabText(index int) (label string)
TabToolTip(index int) (tip string)
WidgetOf(index int) (w WidgetItem)
* OnCurrentChanged(fn func(int))
]]


qtdrv = {
inc = "<QTabWidget>",
name = "QTabWidget",

Init = [[
drvNewObj(a0,new QTabWidget);
]],

AddTab = "addTab",
Clear = "clear",
Count = "count",
CurrentIndex = "currentIndex",
CurrentWidget = "currentWidget",
SetCurrentIndex = "setCurrentIndex",
SetCurrentWidget = "setCurrentWidget",
IndexOf = "indexOf",
InsertTab = "insertTab",
RemoveTab = "removeTab",
SetTabText = "setTabText",
TabText = "tabText",
SetTabToolTip = "setTabToolTip",
TabToolTip = "tabToolTip",

WidgetOf = [[
drvSetHandle(a2,self->widget(drvGetInt(a1)));
]],

OnCurrentChanged = [[
QObject::connect(self,SIGNAL(currentChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
]],

}
