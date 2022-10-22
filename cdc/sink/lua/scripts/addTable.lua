-- things like this...
function addTable(tableID)
	print("add tableID:", tableID)
	local file = io.open( "./data/addTable.txt", "w" )
    file:write( "here's a addTable and tableID: " )
    file:write( tableID.."\r\n" )
    file:close()

    local http = require("http")

    -- POST
    data = "{\"operation\":\"sink_add_table\",  \"data\": {\"table_id\": " .. tableID .. "}}"

    local response, error_message = http.request("POST", "http://localhost:5005/sink_sync", {
        form=data
    })
    print("res: ", response, error_message)
end
