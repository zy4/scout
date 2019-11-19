# scout

**scout** is a middleware for the **chi-router** to report events to sentry.io using the **sentry-go** client.

`sentry-go` client must be [initialized with `sentry.Init`](https://github.com/getsentry/sentry-go).

This middleware also replaces `middleware.Recoverer`.

Here is an example on how to use the middleware:

```go
r := chi.NewRouter()

// Use the middleware in the router
r.Use(scout.SentryRecovery)

r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    panic("caught")
})
```

## Attribution

**scout** was created to emulate the **raven-go** middleware [**raven-chi**](https://github.com/loikg/ravenchi) for the official **sentry-go** SDK
since **raven-go** is deprecated.
