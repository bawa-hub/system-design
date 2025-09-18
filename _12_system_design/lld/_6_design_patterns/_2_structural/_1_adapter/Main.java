package _6_design_patterns._2_structural._1_adapter;

// Step 1: Define the Target Interface

// This is the interface your system understands.
interface PaymentProcessor {
    void pay(int amount);
}

// Step 2: Create the Adaptee

// This represents the third-party library with a different interface.
class ThirdPartyPayment {
    public void makePayment(double amount) {
        System.out.println("Payment of $" + amount + " made using ThirdPartyPayment.");
    }
}

// Step 3: Create the Adapter

// This bridges the gap between your system and the third-party library.
class PaymentAdapter implements PaymentProcessor {
    private ThirdPartyPayment thirdPartyPayment;

    public PaymentAdapter(ThirdPartyPayment thirdPartyPayment) {
        this.thirdPartyPayment = thirdPartyPayment;
    }

    @Override
    public void pay(int amount) {
        // Convert int amount to double and delegate to the third-party method
        thirdPartyPayment.makePayment((double) amount);
    }
}

// Step 4: Use the Adapter in Your System

// Now you can use the third-party library through your system’s interface.
public class Main {
    public static void main(String[] args) {
        // Create an instance of the third-party payment system
        ThirdPartyPayment thirdPartyPayment = new ThirdPartyPayment();

        // Use the adapter to integrate with your system
        PaymentProcessor paymentProcessor = new PaymentAdapter(thirdPartyPayment);

        // Process a payment
        paymentProcessor.pay(100); // Output: Payment of $100.0 made using ThirdPartyPayment.
    }
}