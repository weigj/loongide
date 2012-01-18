module("basebutton")

name = "baseButton"
base = "Widget"

funcs = [[
@ SetText(text string)
@ Text() (text string)
]]

qtdrv = {
inc = "<QAbstractButton>",
name = "QAbstractButton",

SetText = "setText",
Text = "text",

}
