CFEcho
======
Collection of middleware for use within Cloud foundry

CorrelationMiddleware
======================
This Middleware forwards the X-Vcap-Request-ID to downstream calls via X-Correlation-ID header. It also makes available the CorrelationId in the http.Request Context so downstream middleware/handlers can extract it and use it for e.g. logging

LICENSE
=======
MIT
