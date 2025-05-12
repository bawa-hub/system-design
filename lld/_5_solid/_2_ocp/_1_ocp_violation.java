package _5_solid._2_ocp;

public class _1_ocp_violation {
    
}

class DiscountCalculator {
    public double calculateDiscount(String customerType, double amount) {
        if (customerType.equals("Regular")) {
            return amount * 0.1;
        } else if (customerType.equals("Premium")) {
            return amount * 0.2;
        } else {
            return 0;
        }
    }
}

// Problem:
// If you add a new customer type ("Gold"), you have to modify the calculateDiscount() method. This violates OCP.