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
    receivedString = "Превед: "..str
    print (receivedString)
end

--5
sumNumbers = function (a,b)
	sq = squareGO (a) -- function squareGO should be set before
    return a + b, sq
end

function concat(a, b)
	return a .. " + " .. b
end

function getpage (url)
	print("wait http..")
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
	--print(response["body"])
	return response["status_code"]
end