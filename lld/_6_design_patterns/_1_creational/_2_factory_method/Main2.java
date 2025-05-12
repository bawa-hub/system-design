package _6_design_patterns._1_creational._2_factory_method;

// Why Use a Factory When It Returns a Single Type?

// 1. Encapsulation of Object Creation Logic
// The factory encapsulates the logic for creating the object. While in this example, creating a MySQLConnection might look simple (new MySQLConnection()), in a real project, the process of creating an instance might involve:

//     Reading configuration files (e.g., database credentials).
//     Setting up logging or monitoring hooks.
//     Handling dependency injection.

// The factory hides this complexity from the rest of your application.

// Example:
class MySQLConnectionFactory implements DatabaseConnectionFactory {
    public DatabaseConnection createConnection() {
        System.out.println("Reading MySQL configuration...");
        // Imagine this is reading from a config file or environment variables
        String host = "localhost";
        String user = "root";
        String password = "password";
        return new MySQLConnection(host, user, password);
    }
}

// Without the factory, every part of your application would need to handle this logic, leading to duplicate code and potential errors.

// 2. Ease of Extensibility
// Suppose the way you create MySQLConnection changes in the future:

//     Maybe you need to add connection pooling.
//     Or you need to set additional parameters like timeouts.

// By using a factory, you only need to update the factory code. All parts of your application that use the factory automatically benefit from the new logic.

// 3. Polymorphism and Runtime Flexibility
// Factories make it easy to switch between different implementations at runtime.

// Example: Let's say your app connects to different databases based on the environment:

//     Use MySQL in development.
//     Use PostgreSQL in production.

// With factories, you can decide at runtime which connection to create.

class DatabaseConnectionFactoryProvider {
    public static DatabaseConnectionFactory getFactory(String dbType) {
        if (dbType.equalsIgnoreCase("mysql")) {
            return new MySQLConnectionFactory();
        } else if (dbType.equalsIgnoreCase("postgresql")) {
            return new PostgreSQLConnectionFactory();
        } else {
            throw new IllegalArgumentException("Unsupported database type");
        }
    }
}

public class Main2 {
    public static void main(String[] args) {
        String environment = "development"; // Or "production"
        String dbType = environment.equals("development") ? "mysql" : "postgresql";

        DatabaseConnectionFactory factory = DatabaseConnectionFactoryProvider.getFactory(dbType);
        DatabaseConnection connection = factory.createConnection();
        connection.connect();
    }
}

// This flexibility is much harder to achieve if you don’t use factories.


// 4. Testing and Mocking
// Using a factory makes unit testing easier. You can:

//     Replace the factory with a mock factory that returns mock objects for testing.
//     Simulate different environments or scenarios without touching the application logic.

// Example:
class MockMySQLConnectionFactory implements DatabaseConnectionFactory {
    public DatabaseConnection createConnection() {
        return new MockMySQLConnection(); // A mock implementation for testing
    }
}

// 5. Consistency Across Multiple Object Types
// In some systems, you might have several factories for different objects (e.g., MySQLConnection, PostgreSQLConnection, etc.). Using factories provides a consistent pattern for object creation, which makes the codebase easier to understand and maintain.