I — Interface Segregation Principle (ISP)

    Definition:
    "Clients should not be forced to depend on interfaces they do not use."

In other words:
👉 Split large, bloated interfaces into smaller, more specific ones so that classes only need to implement what they actually use.

🏢 Real-World Analogy:

Imagine an interface Worker with the method attendMeeting().

    A software engineer can implement it.
    But a janitor shouldn't be forced to implement attendMeeting() — it's not part of their role. That’s a violation of ISP.
