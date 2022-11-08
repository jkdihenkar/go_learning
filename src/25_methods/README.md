# Methods

Methods are the functions with recievers.

## Why pointer recievers for methods?
- The first is so that the method can modify the value that its receiver points to.
- The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
