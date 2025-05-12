package _5_solid._1_srp;

public class _2_refactor_using_srp {
    
}

// Core business logic
class Invoice {
    private int amount;

    public Invoice(int amount) {
        this.amount = amount;
    }

    public int calculateTotal() {
        return amount + (amount * 18 / 100);
    }

    public int getAmount() {
        return amount;
    }
}

// Responsible only for printing
class InvoicePrinter {
    public void print(Invoice invoice) {
        System.out.println("Printing invoice with amount: " + invoice.getAmount());
    }
}

// Responsible only for persistence
class InvoicePersistence {
    public void save(Invoice invoice) {
        System.out.println("Saving invoice with amount: " + invoice.getAmount() + " to DB...");
    }
}

// Each class now has one reason to change:
//     Invoice: Business logic changes
//     InvoicePrinter: Output format changes
//     InvoicePersistence: Storage mechanism changes