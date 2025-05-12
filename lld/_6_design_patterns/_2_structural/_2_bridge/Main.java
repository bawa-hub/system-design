package _6_design_patterns._2_structural._2_bridge;


// Step 1: Define the Implementation Interface

// Implementor interface
interface Color {
    void applyColor();
}

// Step 2: Create Concrete Implementations

class RedColor implements Color {
    @Override
    public void applyColor() {
        System.out.println("Applying red color.");
    }
}

class BlueColor implements Color {
    @Override
    public void applyColor() {
        System.out.println("Applying blue color.");
    }
}

// Step 3: Define the Abstraction

abstract class Shape {
    protected Color color; // Bridge to implementation

    public Shape(Color color) {
        this.color = color;
    }

    abstract void draw();
}

// Step 4: Create Concrete Abstractions

class Circle extends Shape {
    public Circle(Color color) {
        super(color);
    }

    @Override
    void draw() {
        System.out.print("Drawing Circle - ");
        color.applyColor(); // Delegate color application
    }
}

class Square extends Shape {
    public Square(Color color) {
        super(color);
    }

    @Override
    void draw() {
        System.out.print("Drawing Square - ");
        color.applyColor(); // Delegate color application
    }
}

// Step 5: Use the Bridge Pattern

public class Main {
    public static void main(String[] args) {
        // Create red circle
        Shape redCircle = new Circle(new RedColor());
        redCircle.draw();

        // Create blue square
        Shape blueSquare = new Square(new BlueColor());
        blueSquare.draw();
    }
}

// Output

// Drawing Circle - Applying red color.
// Drawing Square - Applying blue color.

/** 
Key Points of the Example

    Implementor Interface (Color): Defines the low-level details.
    Concrete Implementations (RedColor, BlueColor): Provide specific implementations.
    Abstraction (Shape): Provides the high-level interface, bridging to the implementor.
    Concrete Abstractions (Circle, Square): Implement specific behavior while using the implementor.
*/    