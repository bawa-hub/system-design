package _5_solid._5_dip;

public class _2_dip_complaint_refactor {
    public static void main(String[] args) {
        PaymentMethod creditCard = new CreditCardProcessor();
        PaymentProcessor processor = new PaymentProcessor(creditCard);
        processor.process(100.0);

        PaymentMethod paypal = new PayPalProcessor();
        processor = new PaymentProcessor(paypal);
        processor.process(200.0);
    }
}

// Define a PaymentMethod interface for abstraction.
// Make PaymentProcessor depend on the interface, not a concrete class.

interface PaymentMethod {
    void processPayment(double amount);
}

class CreditCardProcessor implements PaymentMethod {
    public void processPayment(double amount) {
        System.out.println("Processing credit card payment of " + amount);
    }
}

class PayPalProcessor implements PaymentMethod {
    public void processPayment(double amount) {
        System.out.println("Processing PayPal payment of " + amount);
    }
}

class PaymentProcessor {
    private PaymentMethod paymentMethod;

    // Injecting dependency via constructor (dependency injection)
    public PaymentProcessor(PaymentMethod paymentMethod) {
        this.paymentMethod = paymentMethod;
    }

    public void process(double amount) {
        paymentMethod.processPayment(amount);
    }
}

// ✅ Why This Works:

//     High-level module (PaymentProcessor) no longer depends on low-level modules (e.g., CreditCardProcessor, PayPalProcessor).
//     It depends on abstractions (PaymentMethod), which makes it extensible and flexible.
//     To add a new payment method (like BitcoinProcessor), you just create a new class that implements PaymentMethod and inject it.