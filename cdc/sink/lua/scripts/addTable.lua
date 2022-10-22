-- things like this...
function addTable(tableID)
	print("add tableID:", tableID)
	local file = io.open( "./data/addTable.txt", "w" )
    file:write( "here's a addTable and tableID: " )
    file:write( tableID.."\r\n" )
    file:close()
end
