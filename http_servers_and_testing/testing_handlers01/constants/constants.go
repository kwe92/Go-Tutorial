package constants

type EndPoints struct {
	IsPrime string
}

type HttpMethods struct {
	GET string
}

var Endpoints = EndPoints{
	IsPrime: "/check-is-prime",
}

var HTTPMethods = HttpMethods{
	GET: "GET",
}
