package _6_design_patterns._3_behavioral._10_template_method;

// Step 1: Define the Template Method in the Base Class

abstract class OrderProcessor {
    
    // Template method
    public final void processOrder() {
        validateOrder();
        processPayment();
        shipOrder();
        sendNotification();
    }

    // Steps to be implemented by subclasses
    abstract void validateOrder();
    abstract void processPayment();
    abstract void shipOrder();
    abstract void sendNotification();
}

// Step 2: Implement Concrete Subclasses with Specific Behavior

class CreditCardOrderProcessor extends OrderProcessor {

    @Override
    void validateOrder() {
        System.out.println("Validating credit card order...");
    }

    @Override
    void processPayment() {
        System.out.println("Processing payment through Credit Card...");
    }

    @Override
    void shipOrder() {
        System.out.println("Shipping order via standard delivery...");
    }

    @Override
    void sendNotification() {
        System.out.println("Sending credit card payment confirmation...");
    }
}

class PayPalOrderProcessor extends OrderProcessor {

    @Override
    void validateOrder() {
        System.out.println("Validating PayPal order...");
    }

    @Override
    void processPayment() {
        System.out.println("Processing payment through PayPal...");
    }

    @Override
    void shipOrder() {
        System.out.println("Shipping order via express delivery...");
    }

    @Override
    void sendNotification() {
        System.out.println("Sending PayPal payment confirmation...");
    }
}

// Step 3: Use the Template Method Pattern

public class Main {
    public static void main(String[] args) {
        // Using the CreditCardOrderProcessor
        OrderProcessor creditCardOrder = new CreditCardOrderProcessor();
        creditCardOrder.processOrder(); // Follows the template method

        System.out.println();

        // Using the PayPalOrderProcessor
        OrderProcessor paypalOrder = new PayPalOrderProcessor();
        paypalOrder.processOrder(); // Follows the template method
    }
}

// Output

// Validating credit card order...
// Processing payment through Credit Card...
// Shipping order via standard delivery...
// Sending credit card payment confirmation...

// Validating PayPal order...
// Processing payment through PayPal...
// Shipping order via express delivery...
// Sending PayPal payment confirmation...


/** 
Key Points of the Example
    Template Method (processOrder): Defines the algorithm for processing the order, but leaves the specific steps to be implemented by subclasses.
    Concrete Subclasses (CreditCardOrderProcessor, PayPalOrderProcessor): Implement specific behaviors for payment validation, processing, shipping, and notification.
    Code Reusability: The processOrder() method provides a common structure for all order processing, eliminating duplicate code.
*/