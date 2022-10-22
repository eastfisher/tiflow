-- things like this...
function removeTable(tableID)
	print("remove tableID:", tableID)
	local file = io.open( "./data/removeTable.txt", "w" )
    file:write( "here's a removeTable and tableID: " )
    file:write( tableID.."\r\n" )
    file:close()
end
