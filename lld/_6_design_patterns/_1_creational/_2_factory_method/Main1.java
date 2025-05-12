package _6_design_patterns._1_creational._2_factory_method;

/**
 * 
 * Example from a Real Project: Database Connections

In real applications, different databases (e.g., MySQL, PostgreSQL, Oracle) require different connection setups. The Factory Method Pattern is often used to create database connections.
Problem:

Your application needs to connect to different databases, but the connection process for each database is slightly different. Instead of managing the connection logic for each database, you use a factory method to handle it.
 */

//  Step 1: Create a Common Interface for Database Connections
interface DatabaseConnection {
    void connect();
}

// Step 2: Create Specific Connection Classes
class MySQLConnection implements DatabaseConnection {
    public void connect() {
        System.out.println("Connecting to MySQL database...");
    }
}

class PostgreSQLConnection implements DatabaseConnection {
    public void connect() {
        System.out.println("Connecting to PostgreSQL database...");
    }
}

class OracleConnection implements DatabaseConnection {
    public void connect() {
        System.out.println("Connecting to Oracle database...");
    }
}

// Step 3: Create a Factory Interface
interface DatabaseConnectionFactory {
    DatabaseConnection createConnection();
}

// Step 4: Create Specific Factories for Databases
class MySQLConnectionFactory implements DatabaseConnectionFactory {
    public DatabaseConnection createConnection() {
        return new MySQLConnection();
    }
}

class PostgreSQLConnectionFactory implements DatabaseConnectionFactory {
    public DatabaseConnection createConnection() {
        return new PostgreSQLConnection();
    }
}

class OracleConnectionFactory implements DatabaseConnectionFactory {
    public DatabaseConnection createConnection() {
        return new OracleConnection();
    }
}

// Step 5: Use the Factories in Your Application
public class Main1 {
    public static void main(String[] args) {
        // Get a MySQL connection
        DatabaseConnectionFactory mysqlFactory = new MySQLConnectionFactory();
        DatabaseConnection mysqlConnection = mysqlFactory.createConnection();
        mysqlConnection.connect(); // Output: Connecting to MySQL database...

        // Get a PostgreSQL connection
        DatabaseConnectionFactory postgresFactory = new PostgreSQLConnectionFactory();
        DatabaseConnection postgresConnection = postgresFactory.createConnection();
        postgresConnection.connect(); // Output: Connecting to PostgreSQL database...

        // Get an Oracle connection
        DatabaseConnectionFactory oracleFactory = new OracleConnectionFactory();
        DatabaseConnection oracleConnection = oracleFactory.createConnection();
        oracleConnection.connect(); // Output: Connecting to Oracle database...
    }
}

