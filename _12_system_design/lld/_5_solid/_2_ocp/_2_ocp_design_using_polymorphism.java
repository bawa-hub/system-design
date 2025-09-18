package _5_solid._2_ocp;

public class _2_ocp_design_using_polymorphism {

    // usage
    DiscountCalculator calculator = new DiscountCalculator(new PremiumCustomerDiscount());
    double discounted = calculator.calculate(500);
}

// Define a base abstraction:
interface DiscountStrategy {
    double getDiscount(double amount);
}

class RegularCustomerDiscount implements DiscountStrategy {
    public double getDiscount(double amount) {
        return amount * 0.1;
    }
}

// Implement specific strategies:
class PremiumCustomerDiscount implements DiscountStrategy {
    public double getDiscount(double amount) {
        return amount * 0.2;
    }
}

// Calculator now works with abstraction:
class DiscountCalculator {
    private DiscountStrategy discountStrategy;

    public DiscountCalculator(DiscountStrategy discountStrategy) {
        this.discountStrategy = discountStrategy;
    }

    public double calculate(double amount) {
        return discountStrategy.getDiscount(amount);
    }
}

// Now, if you add a GoldCustomerDiscount, you just create a new class — no modification to existing code!