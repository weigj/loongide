-- tag +,-,@,*
function get_func(s)
	local tag,name,input,output = string.match(s,"([%+%-%@%*]?)%s*([_%w%*%[%]]+)%s*(%b())%s*(%b())")
	if name == nil then
		tag,name,input = string.match(s,"([%+%-%@%*]?)%s*([_%w%*%[%]]+)%s*(%b())")
	end
	if name == nil then
		tag,name = string.match(s,"([%+%-%@%*]?)%s*([_%w%*%[%]]+)")
	end

	if name == nil then
		return nil
	end
	input = input or "()"
	output = output or "()"

	input = string.sub(input,2,-2)..","
	output = string.sub(output,2,-2)..","

	local func = {}
	func.tag = tag or ""
	func.name = name
	func.input = {}
	func.output = {}

	for v,t in string.gmatch(input,"([_%w]+)%s+([_%w%*%[%]%(%),]+),") do
		func.input[#func.input+1] = {var=v,type=t}
	end

	for v,t in string.gmatch(output,"([_%w]+)%s+([_%w%*%[%]%(%),]+),") do
		func.output[#func.output+1] = {var=v,type=t}
	end

	return func
end

function get_lines(s)
	local i = 0
	local j = 0
	local lines = {}
	while true do
		j = string.find(s.."\n","\n",i)
		if j == nil then
			break
		end
		local line = string.sub(s,i,j-1)
		if line ~= "" then
			lines[#lines+1] = line
		end
		i = j+1
	end
	return lines
end

function get_funcs(s)
	local lines = get_lines(s)
	local funcs = {}
	for k,v in ipairs(lines) do
		local func = get_func(v)
		if func ~= nil then
			funcs[#funcs+1] = func
		end
	end
	return funcs
end

function make_params(t)
	local o = {}
	for k,v in ipairs(t) do
		o[k] = v.var .." ".. v.type
	end
	return "("..table.concat(o,",")..")"
end

function print_funcs(funcs)
	for k,v in ipairs(funcs) do
		print(k,v.tag,v.name,make_params(v.input),make_params(v.output))
	end
end

--cdrv_type = "qtdrv"

go_def = {}
c_def = {}

function begin_def()
	go_def.heads = {}
	go_def.drvenum = {}
	go_def.enums = {}
	go_def.defs = {}
	go_def.funcs = {}

	c_def.include = {}
	c_def.drvenum = {}
	c_def.drvfunc = {}
	c_def.enums = {}
	c_def.funcs = {}

	go_def.heads[#go_def.heads+1] = [[
// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
]]
	go_def.heads[#go_def.heads+1] = [[
package ui
]]

	go_def.drvenum[#go_def.drvenum+1] = "// drvclass enums"
	go_def.drvenum[#go_def.drvenum+1] = "const ("
	go_def.drvenum[#go_def.drvenum+1] = "\tDRVCLASS_NONE = iota"

	c_def.drvenum[#c_def.drvenum+1] = "// drvclass enums"
	c_def.drvenum[#c_def.drvenum+1] = "enum DRVCLASS_ENUMS {"
	c_def.drvenum[#c_def.drvenum+1] = "    DRVCLASS_NONE = 0,"

	c_def.include[#c_def.include+1] = [[#include "cdrv.h"]]

	c_def.drvfunc[#c_def.drvfunc+1] = [[
typedef int (*DRV_CALLBACK)(void* fn, void *a1,void* a2,void* a3,void* a4);

static DRV_CALLBACK pfn_drv_callback;

int drv_callback(void* fn, void *a1,void* a2,void* a3,void* a4)
{
    return pfn_drv_callback(fn,a1,a2,a3,a4);
}	

int drv(int drvcls, int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    switch(drvcls) {
    case -1:
        pfn_drv_callback = (DRV_CALLBACK)a0;
        return 1;]]

end

function end_def()
	go_def.drvenum[#go_def.drvenum+1] = ")"
	c_def.drvenum[#c_def.drvenum+1] = "    DRVCLASS_LAST"
	c_def.drvenum[#c_def.drvenum+1] = "};"
	c_def.drvfunc[#c_def.drvfunc+1] = [[
    default:
        return 0;
    }
    return 1;
}
]]
end

function def(item)
	local go = {}
	local c = {}
	local name = item.name
	local base = item.base or ""
	local name_tag = string.upper(name)
	local name_cls = "DRVCLASS_"..string.upper(name)

	assert(item.funcs ~= nil,string.format("%s funcs not defined!",name))

	local funcs = get_funcs(item.funcs)
	local cdrv = item[cdrv_type]


	go.enums = {}
	go.defs = {}
	go.funcs = {}
	go.setattr = {}
	go.getattr = {}

	c.enums = {}
	c.funcs = {}

	go_def.drvenum[#go_def.drvenum+1] = "\t"..name_cls
	c_def.drvenum[#c_def.drvenum+1] = "    "..name_cls..","

	go.enums[#go.enums+1] = string.format("// %s drvid enums",name_cls)
	go.enums[#go.enums+1] = "const ("
	go.enums[#go.enums+1] = string.format("\t%s = iota",name_tag.."_NONE")

	c.enums[#c.enums+1] = string.format("// %s drvid enums",name_cls)
	c.enums[#c.enums+1] = string.format("enum %s {",name_tag.."_ENUMS")
	c.enums[#c.enums+1] = string.format("    %s = 0,",name_tag.."_NONE")

	c.funcs[#c.funcs+1] = string.format([[
int drv_%s(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)]],string.lower(name))
	c.funcs[#c.funcs+1] = "{"
	c.funcs[#c.funcs+1] = cdrv.head or string.format([[
    handle_head* head = (handle_head*)a0;
    %s *self = (%s*)head->native;]],cdrv.name,cdrv.name)

	c.funcs[#c.funcs+1] = "    switch (drvid) {"

	-- go struct def
	go.defs[#go.defs+1] = string.format([[
// %s struct
type %s struct {
	%s
}
]],name,name,base)
	go.defs[#go.defs+1] = string.format([[
func (p *%s) Name() string {
	return "%s"
}
]],name,name)
	go.defs[#go.defs+1] = string.format([[
func (p *%s) String() string {
	return itemString(p)
}]],name)

	go.setattr[#go.setattr+1] = string.format([[
func (p *%s) SetAttr(attr string, value interface{}) bool {
	switch attr {]],name)
	go.getattr[#go.getattr+1] = string.format([[
func (p *%s) Attr(attr string) interface{} {
	switch attr {]],name)
	
	if item.defs ~= nil then
		go.defs[#go.defs+1] = item.defs
	end
	-- funcs enums
	for k,v in ipairs(funcs) do
		local func = v.name
		local tag = v.tag
		local enum = name_tag.."_"..string.upper(v.name)
		go.enums[#go.enums+1] = "\t"..enum
		c.enums[#c.enums+1] = "    "..enum..","
		local go_in = {}
		local go_out = {}
		local go_var = {}
		local go_body1 = {}
		local go_body2 = {}
		local c_in = {}
		local c_out = {}
		local go_in_type = ""
		local go_out_type = ""
		local index = 0
		for k,v in pairs(v.input) do
			index = index+1
			local var = v.var
			local type = v.type
			go_in[#go_in+1] = var.." "..type
			go_var[#go_var+1] = var
			go_in_type = v.type
			local c_type = "drvGet"..string.upper(string.sub(type,1,1))..string.sub(type,2,-1)
			if type == "Item" then
				c_type = "drvGetNative"
			elseif type == "WidgetItem" then
				c_type = "drvGetWidget"
			elseif type == "LayoutItem" then
				c_type = "drvGetLayout"
			elseif type == "MenuItem" then
				c_type = "drvGetMenu"
			elseif type == "*FontInfo" then
				c_type = "drvGetFont"
			elseif type == "*Menu" then
				c_type = "drvGetMenu"
			elseif type == "*Action" then
				c_type = "drvGetAction"
			elseif type == "[]Point" then
				c_type = "drvGetPoints"
			end
			c_in[#c_in+1] = c_type
		end				
		for k,v in pairs(v.output) do
			index = index+1
			local var = v.var
			local type = v.type
			local c_type = "drvSet"..string.upper(string.sub(type,1,1))..string.sub(type,2,-1)			
			go_out[#go_out+1] = var.." "..type
			go_out_type = type
			if type == "string" then
				var = "sh_"..var
				go_body1[#go_body1+1] = string.format("\tvar %s StringHead",var)
				go_body2[#go_body2+1] = string.format("\t%s = %s.String()",v.var,var)
			elseif type == "bool" then
				var = "b_"..var
				go_body1[#go_body1+1] = string.format("\tvar %s int",var)
				go_body2[#go_body2+1] = string.format("\t%s = %s != 0",v.var,var)
			elseif type == "*FontInfo" then
				var = "fh_"..var
				go_body1[#go_body1+1] = string.format("\tvar %s font_head",var)
				go_body2[#go_body2+1] = string.format("\t%s = %s.Font()",v.var,var)
				c_type = "drvSetFont"
			elseif type == "Item" or 
				   type == "WidgetItem" or 
				   type == "LayoutItem" or 
				   type == "MenuItem" then
				var = "hh_"..var
				go_body1[#go_body1+1] = string.format("\tvar %s ItemHead",var)
				go_body2[#go_body2+1] = string.format([[
	if %s.native != 0 {
		item := FindItem(%s.native)
		if item != nil {
			%s = item.(%s)
		}
	}]],var,var,v.var,v.type)
				c_type = "drvSetHandle"
			end
			go_var[#go_var+1] = "&"..var
			c_out[#c_out+1] = c_type
		end

		local go_head = string.format("func (p *%s) %s(%s)",name,func,table.concat(go_in,","))
		if #go_out > 0 then
			go_head = go_head .. string.format("(%s)",table.concat(go_out,","))
		end
		local go_new_head = nil
		local body = ""
		if tag == "+" then
			--"init and new"
			go_body1[#go_body1+1] = string.format("\tp.classid = %s",name_cls)
			go_body2[#go_body2+1] = string.format("\tAddItem(p)")
			go_body = string.format("\tdrv(%s,%s,p",name_cls,enum)
			go_head = go_head .. " *"..name
			local new = "new_"
			if string.sub(name,1,1) == string.upper(string.sub(name,1,1)) then
				new = "New"
			end
			go_new_head = "func "..string.gsub(func,"%w[%l]+",new..name,1) .. "("..table.concat(go_in,",") .. ") *"..name
		elseif tag == "-" then
			-- "delete"
			go_body = string.format("\tdrv(%s,%s,p",name_cls,enum)
		elseif tag == "@" then
			-- "attribute"
			go_body = string.format("\tdrv(%s,%s,p",name_cls,enum)
			local set = string.match(func,"^Set(%u%w+)")
			if set ~= nil and #go_in == 1 then
				local attr = string.lower(set)
				go.setattr[#go.setattr+1] = string.format([[
	case %q:
		if v, ok := value.(%s); ok {
			p.%s(v)
			return true
		}
		return false]],attr,go_in_type,func)
			elseif #go_out == 1 then
				local attr = string.lower(func)
				local is = string.match(func,"^Is(%u%w+)")
				if is ~= nil then
					attr = string.lower(is)
				end
				go.getattr[#go.getattr+1] = string.format([[
	case %q:
		return p.%s()]],attr,func)
			end
		elseif tag == "*" then
			-- "event"
			go_body = string.format("\tdrv_event(%s,%s,p",name_cls,enum)
		else
			-- "func"
			go_body = string.format("\tdrv(%s,%s,p",name_cls,enum)
		end
		if #go_var > 0 then
			go_body = go_body..","..table.concat(go_var,",")
		end
		go_body = go_body..")"

		local go_func = {}
		go_func[#go_func+1] = go_head .. " {"
		if #go_body1 > 0 then
			go_func[#go_func+1] = table.concat(go_body1,"\n")
		end
		go_func[#go_func+1] = go_body
		if #go_body2 > 0 then
			go_func[#go_func+1] = table.concat(go_body2,"\n")
		end
		if tag == "+" then
			go_func[#go_func+1] = "\treturn p"
		else
			go_func[#go_func+1] = "\treturn"
		end
		go_func[#go_func+1] = "}"
		if go_new_head ~= nil then
			local go_new = {}
			go_new[#go_new+1] = go_new_head .. "{"
			go_new[#go_new+1] = string.format("\treturn new(%s).%s(%s)",name,func,table.concat(go_var,","))
			go_new[#go_new+1] = "}"
			go.funcs[#go.funcs+1] = table.concat(go_new,"\n").."\n"
		end
		go.funcs[#go.funcs+1] = table.concat(go_func,"\n").."\n"

		assert(cdrv[func] ~= nil,string.format("cdrv <%s> class %s : %s not defined!",cdrv_type,name,func))
		local cbody = cdrv[func]
		if string.find(cbody,"[%;%(%)]") == nil then		
			if #c_out == 0 then
				local c_type = {}
				for k,v in ipairs(c_in) do
					c_type[#c_type+1] = v .. "(a"..k..")"
				end
				cbody = "self->"..cbody.."("..table.concat(c_type,",")..");"
			elseif #c_out == 1 then
				local c_type = {}
				for k,v in ipairs(c_in) do
					c_type[#c_type+1] = v .. "(a"..k..")"
				end
				cbody = c_out[1].."(a"..(#c_in+1)..",".."self->"..cbody.."("..table.concat(c_type,",").."));"
			end			
		end
		c.funcs[#c.funcs+1] = string.format([[
    case %s: {]],enum)
		--c.funcs[#c.funcs+1] = table.concat(c_var,"\n")
		local c_lines = get_lines(cbody)
		for k, v in ipairs(c_lines) do
			c.funcs[#c.funcs+1] = "        "..v
		end
		c.funcs[#c.funcs+1] = [[
        break;
    }]]
	end

	go.enums[#go.enums+1] = ")"
	go.setattr[#go.setattr+1] = string.format([[
	default:
		return p.%s.SetAttr(attr,value)
	}
	return false
}]],base)
	go.getattr[#go.getattr+1] = string.format([[
	default:
		return p.%s.Attr(attr)
	}
	return nil
}]],base)
	
	c.enums[#c.enums+1] = "    "..name_tag.."_LAST"
	c.enums[#c.enums+1] = "};"	

	if item.expands ~= nil then
		go.funcs[#go.funcs+1] = item.expands
	end
	
	c_def.enums[#c_def.enums+1] = table.concat(c.enums,"\n")
	if cdrv.inc ~= nil then
		c_def.include[#c_def.include+1] = "#include "..cdrv.inc
	end

	go_def.enums[#go_def.enums+1] = table.concat(go.enums,"\n")

	go_def.defs[#go_def.defs+1] = table.concat(go.defs,"\n")
	go_def.funcs[#go_def.funcs+1] = table.concat(go.defs,"\n")
	go_def.funcs[#go_def.funcs+1] = table.concat(go.setattr,"\n")
	go_def.funcs[#go_def.funcs+1] = table.concat(go.getattr,"\n")
	go_def.funcs[#go_def.funcs+1] = table.concat(go.funcs,"\n")
	
	c.funcs[#c.funcs+1] = [[
    default:
        return 0;
    }
    return 1;
}
]]
	c_def.funcs[#c_def.funcs+1] = table.concat(c.funcs,"\n")
	c_def.drvfunc[#c_def.drvfunc+1] = string.format([[
    case %s:
        return drv_%s(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);]],name_cls,string.lower(name))
end
