local http = require("http")

testString = "LuaBridge works Превед!"
receivedString = ""
answer = 126
local qu = 77

--2
window = {
    title = "Window v.0.1",
    width = 400,
    height = 500
}
--3
window2 = {
    title = "Window v.0.2",
    size = {
        w = 400,
        h = 500
    }
}
--4
printMessageLua = function (str)
    receivedString = "Lua.Превед: "..str
    print (receivedString)
end

--5
squareNumberL = function (a)
	sq = _square (a) -- function _square should be registered
    return sq
end

sumNumbersL = function (a,b)
	sq = _summa (a, b) -- function _summa should be registered
    return sq, tostring(sq) .. "^"
end

function concatL(a, b)
	return a .. " & " .. b
end

function getpageL (url)
	print("get http..."..url)
     response, error_message = http.request("GET", url, {
        query="page=1",
        headers={Accept="*/*"},
		cookies={Vanilla="111"}
     })
--	print(error_message)
--	print(response["status_code"])
--	print(response["url"])
--	print(response["body_size"])
--	print(response["cookies"]["Vanilla"])
--	print(response["headers"]["X-Garden-Version"])
--	print(response["headers"]["Cache-Control"])
--	print(response["body"])
	return response["status_code"], response["body_size"] 
end