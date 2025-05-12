package _6_design_patterns._2_structural._1_adapter;


// Scenario: Legacy Logger to a Modern Logging Framework

// Imagine you have an older system that uses a legacy logger with the following interface:
// class LegacyLogger {
//     public void logMessage(String message) {
//         System.out.println("Legacy Log: " + message);
//     }
// }

// Now, your modern application uses a new logging framework with this interface:
// interface ModernLogger {
//     void log(String level, String message);
// }

// The interfaces are incompatible. To use the legacy logger with the modern system, you can create an Adapter.


// Step 1: Define the ModernLogger Interface

// This is the new interface your modern application expects.
interface ModernLogger {
    void log(String level, String message);
}

// Step 2: Create the Adaptee (LegacyLogger)

// This represents the old logger with its unique interface.
class LegacyLogger {
    public void logMessage(String message) {
        System.out.println("Legacy Log: " + message);
    }
}

// Step 3: Create the Adapter

// This bridges the legacy logger with the modern logging interface.
class LoggerAdapter implements ModernLogger {
    private LegacyLogger legacyLogger;

    public LoggerAdapter(LegacyLogger legacyLogger) {
        this.legacyLogger = legacyLogger;
    }

    @Override
    public void log(String level, String message) {
        // Convert modern log levels to legacy log messages
        String formattedMessage = "[" + level.toUpperCase() + "] " + message;
        legacyLogger.logMessage(formattedMessage);
    }
}

// Step 4: Use the Adapter in Your Modern System

public class Main1 {
    public static void main(String[] args) {
        // Create an instance of the legacy logger
        LegacyLogger legacyLogger = new LegacyLogger();

        // Use the adapter to integrate the legacy logger with the modern system
        ModernLogger modernLogger = new LoggerAdapter(legacyLogger);

        // Log messages with different levels
        modernLogger.log("INFO", "This is an informational message.");
        modernLogger.log("ERROR", "An error occurred!");
    }
}

// Output:
// Legacy Log: [INFO] This is an informational message.
// Legacy Log: [ERROR] An error occurred!

/** 
Key Points of the Example:
    LegacyLogger: Represents the old system you can’t change.
    ModernLogger: Represents the modern interface your system expects.
    LoggerAdapter: Adapts the LegacyLogger to the ModernLogger interface by adding the necessary translation logic.
    Benefits:
        You can reuse the LegacyLogger without modifying its code.
        Your modern system can use it as if it were a ModernLogger.
*/        