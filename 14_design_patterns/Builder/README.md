# Builder


- Some objects are simple and can be created in a single constructor call;
- Other objects require a lot of ceremony to create;
- Having a factory function with 10 arguments is not productive
- Instead, opt for piecewise (piece-by-piece) construction
- Builder provides an API for constructing an object step-by-step

> Builder
When piecewise object construction is complicated, provide an API for doing it succinctly

- A Builder is a separate component used for building an object
- To make builder fluent, return the receiver - allows chaining
- Different facets of an object can be built with different builders working in tandem via a common struct