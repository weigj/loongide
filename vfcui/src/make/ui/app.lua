module("app")

name = "app"
base = "object"

defs = [[
]]

funcs = [[
+ Init()
- Close()
Exec() (code int)
Exit(code int)
@ SetFont(font *FontInfo)
@ Font() (font *FontInfo)
CloseAllWindows()

* onRemoveItem(fn func(uintptr))
]]

expands = [[

var theApp *app

func Init() {
	theApp = new_app()
	theApp.onRemoveItem(RemoveItem)
}

func Close() {
	theApp.Close()
}

func RunLoop() int {
	return theApp.Exec()
}

func Exit(code int) {
	theApp.Exit(code)
}

func CloseAllWindows() {
	theApp.CloseAllWindows()
}

func DefaultFont() (font *FontInfo) {
	return theApp.Font()
}

func SetDefaultFont(font *FontInfo) {
	theApp.SetFont(font)
}

]]

qtdrv = {
inc = [["qtapp.h"]],
name = "QtApp",

Init = [[
drvNewObj(a0,new QtApp);
]],

Close = [[
drvDelObj(a0,self);
]],

Exec = "exec",
Exit = "exit",
SetFont = "setFont",
Font = "font",
CloseAllWindows = "closeAllWindows",

onRemoveItem = [[
self->pfnDeleteObj = a1;
]],

}
