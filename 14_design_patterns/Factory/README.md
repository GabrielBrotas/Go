# Factories

Ways of controlling how an object is constructed

- Object creation logic becomes too convoluted
- Struct has too many fields, need to initialize all correctly
- Wholesale object creation (non-piecewise, unlike Builder) can be outsourced to
    - a separate function (Factory Function, a.k.a Constructor)
    - that may exist in a separate struct (Factory)
    
A component responsible solely for the wholesale (not piecewise) creation of objects

- A factory function (a.k.a) is a helper function for making struct instances
- A factory is any entity that can take care of object creation
- Can be a function or dedicated struct