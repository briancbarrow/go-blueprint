package advanced

import (
	_ "embed"
)

//go:embed files/htmx/hello.templ.tmpl
var helloTemplTemplate []byte

//go:embed files/htmx/base.templ.tmpl
var baseTemplTemplate []byte

//go:embed files/htmx/htmx.min.js.tmpl
var htmxMinJsTemplate []byte

//go:embed files/htmx/efs.go.tmpl
var efsTemplate []byte

//go:embed files/htmx/hello.go.tmpl
var helloGoTemplate []byte

//go:embed files/htmx/routes/http_router.tmpl
var httpRouterHtmxTemplRoutes []byte

//go:embed files/htmx/imports/http_router.tmpl
var httpRouterHtmxTemplImports []byte

//go:embed files/htmx/routes/standard_library.tmpl
var stdLibHtmxTemplRoutes []byte

//go:embed files/htmx/imports/standard_library.tmpl
var stdLibHtmxTemplImports []byte

func HttpRouterHtmxTemplImportsTemplate() []byte {
	return httpRouterHtmxTemplImports
}

func HttpRouterHtmxTemplRoutesTemplate() []byte {
	return httpRouterHtmxTemplRoutes
}

func StdLibHtmxTemplRoutesTemplate() []byte {
	return stdLibHtmxTemplRoutes
}

func StdLibHtmxTemplImportsTemplate() []byte {
	return stdLibHtmxTemplImports
}

func HelloTemplTemplate() []byte {
	return helloTemplTemplate
}

func BaseTemplTemplate() []byte {
	return baseTemplTemplate
}

func HtmxJSTemplate() []byte {
	return htmxMinJsTemplate
}

func EfsTemplate() []byte {
	return efsTemplate
}

func HelloGoTemplate() []byte {
	return helloGoTemplate
}
