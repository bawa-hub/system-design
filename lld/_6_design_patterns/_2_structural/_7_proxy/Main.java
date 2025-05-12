package _6_design_patterns._2_structural._7_proxy;

// Step 1: Define the Subject Interface

interface Database {
    void query(String sql);
}

// Step 2: Create the Real Subject

class RealDatabase implements Database {
    @Override
    public void query(String sql) {
        System.out.println("Executing SQL query on the real database: " + sql);
    }
}

// Step 3: Create the Proxy

class DatabaseProxy implements Database {
    private RealDatabase realDatabase;
    private boolean isAuthenticated;

    public DatabaseProxy(boolean isAuthenticated) {
        this.isAuthenticated = isAuthenticated;
    }

    @Override
    public void query(String sql) {
        if (!isAuthenticated) {
            System.out.println("Access Denied: Authentication required!");
            return;
        }

        if (realDatabase == null) {
            realDatabase = new RealDatabase(); // Lazy initialization
        }

        System.out.println("Proxy: Adding logging and access control...");
        realDatabase.query(sql);
    }
}

// Step 4: Use the Proxy

public class Main {
    public static void main(String[] args) {
        // Without authentication
        Database db1 = new DatabaseProxy(false);
        db1.query("SELECT * FROM users");

        System.out.println();

        // With authentication
        Database db2 = new DatabaseProxy(true);
        db2.query("SELECT * FROM orders");
    }
}

// Output
// Access Denied: Authentication required!
// Proxy: Adding logging and access control...
// Executing SQL query on the real database: SELECT * FROM orders


/** 
Key Points of the Example
    Subject (Database): Defines the common interface for both the real object and proxy.
    Real Subject (RealDatabase): The actual object that performs the operation.
    Proxy (DatabaseProxy): Controls access to the real object and adds additional functionality, such as authentication and logging.
    Client Code: Interacts with the proxy, which decides whether and how to forward the request to the real object.
*/