# Atmo + NATS Example Project

This repo is an example of using [Atmo](https://github.com/suborbital/atmo) with NATS as a streaming messaging layer. In this example, Atmo connects to NATS and handles messages on the `com.suborbital.atmo/hello` topic, responding on the `com.suborbital.atmo/respond` topic.

To try it out, clone this repo and run `make run`. You must have [Subo](https://github.com/suborbital/subo), Docker and docker-compose installed.

Then, in another terminal, run `cd tester; go run main.go`. The tester will send two messages, and show you the responses.

Everything about the application is described in the `Directive.yaml` file, and the handler is in `helloworld/src/lib.rs`. Atmo is designed to declaratively set up your application and all of the connections it needs to operate. Your application code does not need to concern itself with **how** connections, servers, and runtimes are set up, it only needs to handle your business logic. This is a very basic example, but illustrates how Atmo can be used to create declarative, event-driven, WebAssembly-powered applications with very minimal boilerplate code!