package _5_solid._3_lsp;

public class _1_lsp_violation {
    
}

class Bird {
    public void fly() {
        System.out.println("Flying...");
    }
}

class Sparrow extends Bird {
    public void fly() {
        System.out.println("Sparrow flying");
    }
}

class Ostrich extends Bird {
    public void fly() {
        throw new UnsupportedOperationException("Ostrich can't fly!");
    }
}

// 🚩 Problem:
// Ostrich is a Bird, but calling fly() breaks expectations — not substitutable. This violates LSP.