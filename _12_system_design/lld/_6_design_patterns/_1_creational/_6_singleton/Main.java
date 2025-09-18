package _6_design_patterns._1_creational._6_singleton;

// Step 1: Implement the Singleton Class

// Here’s how you create a Singleton in Java.
class DatabaseConnection {
    // Static variable to hold the single instance
    private static DatabaseConnection instance;

    // Private constructor to prevent instantiation from outside
    private DatabaseConnection() {
        System.out.println("Database Connection Created.");
    }

    // Public method to provide access to the instance
    public static synchronized DatabaseConnection getInstance() {
        if (instance == null) {
            instance = new DatabaseConnection();
        }
        return instance;
    }

    public void connect() {
        System.out.println("Connecting to the database...");
    }
}

// Step 2: Use the Singleton

public class Main {
    public static void main(String[] args) {
        // Get the single instance of DatabaseConnection
        DatabaseConnection connection1 = DatabaseConnection.getInstance();
        connection1.connect(); // Output: Connecting to the database...

        // Try to create another instance
        DatabaseConnection connection2 = DatabaseConnection.getInstance();
        connection2.connect(); // Output: Connecting to the database...

        // Verify both references point to the same instance
        System.out.println(connection1 == connection2); // Output: true
    }
}

// Key Points in the Code:
//     Static Variable (instance): Holds the single instance of the class.
//     Private Constructor: Prevents direct instantiation from outside the class.
//     Static Method (getInstance): Ensures only one instance is created and provides access to it.
//     Synchronized Access: Ensures thread safety when creating the instance.
