package _5_solid._4_isp;

public class _2_isp_complaint_refactor {
    
}

// Split the interfaces:
interface Printer {
    void print();
}

interface Scanner {
    void scan();
}

interface Fax {
    void fax();
}

// Now classes implement only what they need:
class OldPrinter implements Printer {
    public void print() {
        System.out.println("Printing...");
    }
}

class AllInOnePrinter implements Printer, Scanner, Fax {
    public void print() {
        System.out.println("Printing...");
    }

    public void scan() {
        System.out.println("Scanning...");
    }

    public void fax() {
        System.out.println("Faxing...");
    }
}
