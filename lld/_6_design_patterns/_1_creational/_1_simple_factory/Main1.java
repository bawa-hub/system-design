package _6_design_patterns._1_creational._1_simple_factory;

/***
 * A real-world example of the Simple Factory Design Pattern in a project can be seen in logging systems, such as in applications that use different logging levels (e.g., INFO, DEBUG, ERROR).
   
  Real-Project Example: Logging System

Problem:
In a real application, you need to log messages in different ways:
    Console Logs: Print messages to the console (for development).
    File Logs: Save messages to a file (for production).
    Database Logs: Store messages in a database (for analysis).

Instead of creating different logging objects (ConsoleLogger, FileLogger, DatabaseLogger) everywhere in your code, you can use a Simple Factory to manage this.

How the Simple Factory Helps:
    You ask the factory for the logger you need (e.g., console, file, or database).
    The factory creates the correct logger for you.
    Your main application doesn’t care how the loggers are implemented.
 */

interface Logger {
    void logMessage(String message);
}

class ConsoleLogger implements Logger {
    public void logMessage(String message) {
        System.out.println("Console Logger: " + message);
    }
}

class FileLogger implements Logger {
    public void logMessage(String message) {
        // Simulating writing to a file
        System.out.println("File Logger: Writing '" + message + "' to a file.");
    }
}

class DatabaseLogger implements Logger {
    public void logMessage(String message) {
        // Simulating writing to a database
        System.out.println("Database Logger: Inserting '" + message + "' into a database.");
    }
}

class LoggerFactory {
    public static Logger getLogger(String type) {
        if (type.equalsIgnoreCase("console")) {
            return new ConsoleLogger();
        } else if (type.equalsIgnoreCase("file")) {
            return new FileLogger();
        } else if (type.equalsIgnoreCase("database")) {
            return new DatabaseLogger();
        }
        return null; // Return null if no valid type is provided
    }
}


public class Main1 {
    public static void main(String[] args) {
        // Get a Console Logger
        Logger consoleLogger = LoggerFactory.getLogger("console");
        consoleLogger.logMessage("This is a console log.");

        // Get a File Logger
        Logger fileLogger = LoggerFactory.getLogger("file");
        fileLogger.logMessage("This is a file log.");

        // Get a Database Logger
        Logger databaseLogger = LoggerFactory.getLogger("database");
        databaseLogger.logMessage("This is a database log.");
    }
}

/**
 * 
 * What Happens in a Real Project:
    Flexibility: The factory lets you add new logger types (e.g., CloudLogger) without changing the rest of your code.
    Simplification: Your code doesn’t have to decide which logger to create—it just asks the factory.
    Centralization: All the logic for creating loggers is in one place (the factory), making your code easier to maintain.

    Example in a Real Library:
    This pattern is used in Java’s Logging Framework (like Log4j or SLF4J).
    You configure the logging type (e.g., console, file, or database) in a configuration file, and the framework’s internal factory creates the appropriate logger for your application.

    This approach ensures that your app can seamlessly switch between different logging implementations without rewriting any code.
 */