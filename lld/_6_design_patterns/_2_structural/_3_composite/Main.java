package _6_design_patterns._2_structural._3_composite;

import java.util.ArrayList;
import java.util.List;

// Step 1: Define the Component Interface

interface Graphic {
    void render(); // Common operation
}

// Step 2: Create Leaf Components (Individual Objects)

class Circle implements Graphic {
    @Override
    public void render() {
        System.out.println("Rendering Circle");
    }
}

class Rectangle implements Graphic {
    @Override
    public void render() {
        System.out.println("Rendering Rectangle");
    }
}

// Step 3: Create Composite Class (Collection)

class Group implements Graphic {
    private List<Graphic> graphics = new ArrayList<>();

    public void add(Graphic graphic) {
        graphics.add(graphic);
    }

    public void remove(Graphic graphic) {
        graphics.remove(graphic);
    }

    @Override
    public void render() {
        System.out.println("Rendering Group");
        for (Graphic graphic : graphics) {
            graphic.render(); // Delegate rendering to child components
        }
    }
}

// Step 4: Use the Composite Pattern

public class Main {
    public static void main(String[] args) {
        // Create individual shapes
        Graphic circle = new Circle();
        Graphic rectangle = new Rectangle();

        // Create a group and add shapes to it
        Group group1 = new Group();
        group1.add(circle);
        group1.add(rectangle);

        // Create another group and add the first group
        Group group2 = new Group();
        group2.add(group1);
        group2.add(new Circle());

        // Render all shapes and groups
        group2.render();
    }
}

// Output

// Rendering Group
// Rendering Group
// Rendering Circle
// Rendering Rectangle
// Rendering Circle


/** 
Key Points of the Example

    Component Interface (Graphic): Defines common operations for all components (e.g., render()).
    Leaf Components (Circle, Rectangle): Represent individual objects with specific behaviors.
    Composite Class (Group): Represents a collection of components (can include both leaf and composite objects).
    Client Code: Interacts with both leaf and composite objects through the same interface.
*/    