L — Liskov Substitution Principle (LSP)

    Barbara Liskov:
    “If S is a subtype of T, then objects of type T may be replaced with objects of type S without altering any of the desirable properties of the program.”

In short:
👉 Subtypes must be substitutable for their base types
👉 A subclass should not break the behavior expected from the parent class

Real-World Analogy:
Imagine a Bird base class. All birds should fly, right?
Now if you introduce a Penguin as a subclass that can't fly, and still call bird.fly(), you'll get runtime surprises.
That's an LSP violation: the subclass doesn't behave like the base.
