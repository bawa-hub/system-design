package _5_solid._3_lsp;

public class _2_lsp_compliant_design {
    
}

// 👉 Split the hierarchy to ensure substitutability.
interface Bird {
    void layEggs();
}

interface FlyingBird extends Bird {
    void fly();
}

// Now define only appropriate behavior:
class Sparrow implements FlyingBird {
    public void fly() {
        System.out.println("Sparrow flying");
    }

    public void layEggs() {
        System.out.println("Sparrow laying eggs");
    }
}

class Ostrich implements Bird {
    public void layEggs() {
        System.out.println("Ostrich laying eggs");
    }
}

// ✅ Now, no one expects Ostrich to fly.
// ✅ Substitution works safely.