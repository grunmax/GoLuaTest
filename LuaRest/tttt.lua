--local io = require("io")
local http = require("socket.http")
--local ltn12 = require("ltn12")

-- connect to server "www.cs.princeton.edu" and retrieves this manual
-- file from "~diego/professional/luasocket/http.html" and print it to stdout
--http.request{ 
--    url = "http://w3.impa.br/~diego/software/luasocket/http.html", 
--    sink = ltn12.sink.file(io.stdout)
--}

r, c, h = http.request {
  method = "GET",
  url = "http://w3.impa.br/~diego/software/luasocket/http.html"
}
print(h["date"])
print(h["server"])
print(h["content-length"])
print(h["content-type"])
print(h["connection"])
print(h["last-modified"])