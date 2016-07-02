local request = require "http.request"

-- This endpoint returns a never-ending stream of chunks containing the current time
local req = request.new_from_uri("http://w3.impa.br/~diego/software/luasocket/http.html")
local _, stream = assert(req:go())
for chunk in stream:each_chunk() do
	io.write(chunk)
end