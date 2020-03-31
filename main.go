package main

import (
	restservice "restful_api/conf"
)

func main() {
	restservice.RestfulAPIServiceInit("HTTP")
	restservice.RestfulAPIServiceInit("HTTPS")
}

