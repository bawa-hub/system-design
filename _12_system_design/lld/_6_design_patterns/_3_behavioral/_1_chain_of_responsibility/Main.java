package _6_design_patterns._3_behavioral._1_chain_of_responsibility;


// Step 1: Define the Handler Interface

// Each handler defines how to process the request and forward it to the next handler.
abstract class Logger {
    protected Logger nextLogger;

    public void setNextLogger(Logger nextLogger) {
        this.nextLogger = nextLogger;
    }

    public void logMessage(int level, String message) {
        if (canHandle(level)) {
            write(message);
        }
        if (nextLogger != null) {
            nextLogger.logMessage(level, message);
        }
    }

    protected abstract boolean canHandle(int level);
    protected abstract void write(String message);
}

// Step 2: Create Concrete Handlers

// Each handler processes specific log levels.
class InfoLogger extends Logger {
    @Override
    protected boolean canHandle(int level) {
        return level == 1; // INFO level
    }

    @Override
    protected void write(String message) {
        System.out.println("INFO: " + message);
    }
}

class DebugLogger extends Logger {
    @Override
    protected boolean canHandle(int level) {
        return level == 2; // DEBUG level
    }

    @Override
    protected void write(String message) {
        System.out.println("DEBUG: " + message);
    }
}

class ErrorLogger extends Logger {
    @Override
    protected boolean canHandle(int level) {
        return level == 3; // ERROR level
    }

    @Override
    protected void write(String message) {
        System.out.println("ERROR: " + message);
    }
}

// Step 3: Create the Chain

class LoggerChain {
    public static Logger getChainOfLoggers() {
        Logger errorLogger = new ErrorLogger();
        Logger debugLogger = new DebugLogger();
        Logger infoLogger = new InfoLogger();

        infoLogger.setNextLogger(debugLogger);
        debugLogger.setNextLogger(errorLogger);

        return infoLogger;
    }
}

// Step 4: Use the Chain

public class Main {
    public static void main(String[] args) {
        Logger loggerChain = LoggerChain.getChainOfLoggers();

        loggerChain.logMessage(1, "This is an informational message.");
        loggerChain.logMessage(2, "This is a debug message.");
        loggerChain.logMessage(3, "This is an error message.");
    }
}

// Output:
// INFO: This is an informational message.
// DEBUG: This is a debug message.
// ERROR: This is an error message.



// Key Points of the Example

//     Logger Chain: Represents the chain of handlers (InfoLogger, DebugLogger, ErrorLogger).
//     Request: The logMessage method processes the request based on the level.
//     Decoupling: The client (Main class) doesn’t know which logger will handle the request.