# Architecture Overview

The service follows a layered architecture: an HTTP handler layer thin enough to only unmarshal requests and dispatch to a domain service.

The domain service coordinates business rules and delegates persistence to a repository layer sitting behind an interface.

State lives in Postgres, cache lives in Redis, and async work is fanned out via a light-weight worker pool consuming from a durable queue.

Observability is uniform: every request carries an X-Request-ID that flows through logs, traces, and outbound calls to downstream services.

Failures are surfaced as typed errors and mapped to HTTP responses at the boundary, never mid-stack.
