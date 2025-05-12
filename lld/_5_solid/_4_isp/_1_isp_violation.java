package _5_solid._4_isp;

public class _1_isp_violation {
    
}

interface MultiFunctionDevice {
    void print();
    void scan();
    void fax();
}

class OldPrinter implements MultiFunctionDevice {
    public void print() {
        System.out.println("Printing...");
    }

    public void scan() {
        throw new UnsupportedOperationException("Scan not supported");
    }

    public void fax() {
        throw new UnsupportedOperationException("Fax not supported");
    }
}

// 🚩 Problem:
// OldPrinter is forced to implement scan() and fax() it doesn't support. This violates ISP.