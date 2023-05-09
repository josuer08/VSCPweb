module VSCPweb

go 1.19

require github.com/go-chi/chi/v5 v5.0.7

require healthHandlers v0.0.0

replace healthHandlers v0.0.0 => ./handlers/health/

require controllerHandlers v0.0.0

replace controllerHandlers v0.0.0 => ./handlers/controller/

require managementHandlers v0.0.0

replace managementHandlers v0.0.0 => ./handlers/management/

require virtualAPIHandlers v0.0.0

replace virtualAPIHandlers v0.0.0 => ./handlers/virtualapi/

require staticHandlers v0.0.0

replace staticHandlers v0.0.0 => ./handlers/static/
