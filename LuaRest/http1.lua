local http = require("http")

        response, error_message = http.request("GET", "http://vsi.org.ua", {
            query="page=1",
            headers={Accept="*/*"},
			cookies={Vanilla="111"}
        })
		print(error_message)
		print(response["status_code"])
		print(response["url"])
		print(response["body_size"])
		print(response["cookies"]["Vanilla"])
		print(response["headers"]["X-Garden-Version"])
		print(response["headers"]["Cache-Control"])
		
		--print(response["body"])