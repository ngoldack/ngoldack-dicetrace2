package core

const (
	ApplicationJson = "application/json"

	ErrorRequestInvalidContentType   = "errors#invalid-content-type"
	ErrorRequestMalformedRequestBody = "errors#malformed-request-body"
	ErrorRequestInvalidRequestBody   = "errors#invalid-request-body"
	ErrorRequestInvalidUuidProvided  = "errors#invalid-uuid-provided"

	ErrorDatabaseNodeNotCreated = "errors#database-node-not-created"
	ErrorDatabaseNodeNotFound   = "errors#database-node-not-found"
	ErrorDatabaseNodeNotDeleted = "errors#database-node-not-deleted"
	ErrorDatabaseInternal       = "errors#database-internal"
)
