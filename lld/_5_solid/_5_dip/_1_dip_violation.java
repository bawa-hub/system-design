package _5_solid._5_dip;

public class _1_dip_violation {
    
}

// PaymentProcessor that directly depends on a CreditCardProcessor:
class CreditCardProcessor {
    public void processPayment(double amount) {
        System.out.println("Processing credit card payment of " + amount);
    }
}

class PaymentProcessor {
    private CreditCardProcessor creditCardProcessor;

    public PaymentProcessor() {
        this.creditCardProcessor = new CreditCardProcessor(); // tight coupling
    }

    public void process(double amount) {
        creditCardProcessor.processPayment(amount);
    }
}

// 🚩 Problem:

//     Tight coupling between PaymentProcessor and CreditCardProcessor.
//     To support another payment method (e.g., PayPal), you must modify PaymentProcessor.