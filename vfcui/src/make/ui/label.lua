module("label")

name = "Label"
base = "Frame"

funcs = [[
+ Init()
@ SetText(text string)
@ Text() (text string)
@ SetWordWrap(b bool)
@ WordWrap() (b bool)
@ SetTextFormat(format int)
@ TextFormat() (format int)
* OnLinkActivated(fn func(string))
]]


qtdrv = {
inc = "<QLabel>",
name = "QLabel",

Init = [[
drvNewObj(a0,new QLabel);
]],

SetText = "setText",
Text = "text",
SetWordWrap = "setWordWrap",
WordWrap = "wordWrap",
SetTextFormat = [[
self->setTextFormat((Qt::TextFormat)drvGetInt(a1));
]],
TextFormat = [[
drvSetInt(a1,self->textFormat());
]],

OnLinkActivated = [[
QObject::connect(self,SIGNAL(linkActivated(QString)),drvNewSignal(self,a1,a2),SLOT(call(QString)));
]],

}
