package index_controller

import "larago/routes"

func Index(request routes.Request) interface{} {

	return "Hello World"
}

type TestStruct struct {
	Name, Email string
}

func JsonTest(request routes.Request) interface{} {

	test := TestStruct{
		Name:  "Test Name",
		Email: "TestEmail",
	}

	return routes.JsonResponse(test)

}
