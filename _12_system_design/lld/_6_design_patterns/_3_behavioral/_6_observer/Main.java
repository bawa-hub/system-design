package _6_design_patterns._3_behavioral._6_observer;

import java.util.ArrayList;
import java.util.List;

// Step 1: Define the Observer Interface

interface Observer {
    void update(String stockName, double price);
}

// Step 2: Define the Subject Interface

interface Subject {
    void addObserver(Observer observer);
    void removeObserver(Observer observer);
    void notifyObservers();
}

// Step 3: Create the Concrete Subject

class Stock implements Subject {
    private String name;
    private double price;
    private List<Observer> observers = new ArrayList<>();

    public Stock(String name, double price) {
        this.name = name;
        this.price = price;
    }

    public void setPrice(double price) {
        this.price = price;
        notifyObservers();
    }

    public double getPrice() {
        return price;
    }

    @Override
    public void addObserver(Observer observer) {
        observers.add(observer);
    }

    @Override
    public void removeObserver(Observer observer) {
        observers.remove(observer);
    }

    @Override
    public void notifyObservers() {
        for (Observer observer : observers) {
            observer.update(name, price);
        }
    }
}

// Step 4: Create Concrete Observers

class Investor implements Observer {
    private String name;

    public Investor(String name) {
        this.name = name;
    }

    @Override
    public void update(String stockName, double price) {
        System.out.println(name + " received update: " + stockName + " is now $" + price);
    }
}

// Step 5: Use the Observer Pattern

public class Main {
    public static void main(String[] args) {
        // Create a stock
        Stock googleStock = new Stock("Google", 1500.00);

        // Create observers
        Investor alice = new Investor("Alice");
        Investor bob = new Investor("Bob");

        // Add observers to the stock
        googleStock.addObserver(alice);
        googleStock.addObserver(bob);

        // Change stock price
        googleStock.setPrice(1520.50);
        googleStock.setPrice(1550.00);

        // Remove an observer
        googleStock.removeObserver(bob);

        // Change stock price again
        googleStock.setPrice(1580.00);
    }
}

// Output
// Alice received update: Google is now $1520.5
// Bob received update: Google is now $1520.5
// Alice received update: Google is now $1550.0
// Bob received update: Google is now $1550.0
// Alice received update: Google is now $1580.0

/** 
Key Points of the Example
    Subject (Stock): The object being observed. It manages observers and notifies them of changes.
    Observers (Investors): Depend on the subject for updates and react to changes.
    Dynamic Management: Observers can be added or removed at runtime.
    Loose Coupling: The subject doesn’t know about the implementation of the observers.
*/
