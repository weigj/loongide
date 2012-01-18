require "ui.object"

require "ui.app"


require "ui.widget"
require "ui.tabwidget"

require "ui.dialog"

require "ui.layout"
require "ui.boxlayout"
require "ui.hboxlayout"
require "ui.vboxlayout"

require "ui.basebutton"
require "ui.button"
require "ui.checkbox"
require "ui.radio"

require "ui.frame"
require "ui.label"

require "ui.groupbox"

require "ui.painter"

require "ui.menubar"
require "ui.menu"
require "ui.action"

require "ui.combobox"

require "makelib"

function make()
	cdrv_type = "qtdrv"
	begin_def()
		def(object)

		def(app)

		--def(font)

		def(widget)

		def(menubar)
		def(menu)
		def(action)

		def(tabwidget)

		def(layout)
		def(boxlayout)
		def(hboxlayout)
		def(vboxlayout)

		def(basebutton)
		def(button)
		def(checkbox)
		def(radio)

		def(frame)
		def(label)
		def(groupbox)
		def(dialog)

		def(combobox)

		def(painter)
	end_def()

	local ui = io.open("../ui/ui.go","w")
	ui:write(table.concat(go_def.heads,"\n"))
	ui:write("\n")
	ui:write(table.concat(go_def.drvenum,"\n"))
	ui:write("\n")
	ui:write(table.concat(go_def.enums,"\n"))
	ui:write("\n")
	ui:write(table.concat(go_def.funcs,"\n"))
	ui:write("\n")
	ui:close()

	local cdrv = io.open(string.format("../%s/cdrv.cpp",cdrv_type),"w")
	cdrv:write(table.concat(c_def.include,"\n"))
	cdrv:write("\n")
	cdrv:write(table.concat(c_def.drvenum,"\n"))
	cdrv:write("\n")
	cdrv:write(table.concat(c_def.enums,"\n"))
	cdrv:write("\n")
	cdrv:write(table.concat(c_def.funcs,"\n"))
	cdrv:write("\n")
	cdrv:write(table.concat(c_def.drvfunc,"\n"))
	cdrv:write("\n")
	cdrv:close()
end

make()
