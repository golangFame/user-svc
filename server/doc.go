//for handlers
//var (
//	encodeSpan, _ = tracer.StartSpanFromContext(dCtx, utils.Function())
//)
//defer encodeSpan.Finish(tracer.WithError(err))

//for services
//encodeSpan, dCtx := tracer.StartSpanFromContext(dCtx, "svc."+utils.FunctionName())
//defer func() {
//	encodeSpan.Finish(tracer.WithError(err))
//}()

/*
no need of environment variable for testing since prod uses different db.

for handlers sending response with updated time ask to put requests with time
*/

package server
