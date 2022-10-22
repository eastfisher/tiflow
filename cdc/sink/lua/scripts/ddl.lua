-- things like this...
function ddl(data)
	print("ddl object:", data)
	local file = io.open( "data/ddl.txt", "w" )
    file:write( "here's a ddl and data:\n" )
    file:write( data.."\r\n" )
    file:close()
end
