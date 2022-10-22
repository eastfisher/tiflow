-- things like this...
function rowChanged(data)
	print("rowChanged object:", data)
	local file = io.open( "data/rowChanged.txt", "w" )
    file:write( "here's a rowChanged and data:\n" )
    file:write( data.."\r\n" )
    file:close()
end
