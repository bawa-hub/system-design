package _5_solid._1_srp;

public class _1_srp_violation {
    
}

class Invoice {
    private int amount;

    public Invoice(int amount) {
        this.amount = amount;
    }

    public void printInvoice() {
        System.out.println("Printing invoice with amount: " + amount);
    }

    public void saveToDatabase() {
        System.out.println("Saving invoice to database...");
    }

    public int calculateTotal() {
        return amount + (amount * 18 / 100); // Add 18% tax
    }
}

/**
 * What's wrong here?

The class Invoice is doing too many things:
    Business logic: calculateTotal()
    Presentation logic: printInvoice()
    Persistence logic: saveToDatabase()

This violates SRP. If any of these responsibilities change (e.g., saving logic changes to use a NoSQL DB), we must touch this class.
 */