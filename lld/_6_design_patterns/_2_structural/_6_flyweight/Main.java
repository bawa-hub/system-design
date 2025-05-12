package _6_design_patterns._2_structural._6_flyweight;

import java.util.HashMap;
import java.util.Map;

// Step 1: Define the Flyweight Interface

interface TreeType {
    void render(int x, int y); // Render method includes unique extrinsic state
}

// Step 2: Create the Concrete Flyweight

class Tree implements TreeType {
    private String name;
    private String color;
    private String texture; // Intrinsic state (shared)

    public Tree(String name, String color, String texture) {
        this.name = name;
        this.color = color;
        this.texture = texture;
    }

    @Override
    public void render(int x, int y) {
        System.out.println("Rendering tree [" + name + "] at (" + x + ", " + y + ") with color: " + color);
    }
}

// Step 3: Create the Flyweight Factory

// The factory ensures shared objects are reused.

class TreeFactory {
    private static Map<String, Tree> treeMap = new HashMap<>();

    public static Tree getTree(String name, String color, String texture) {
        String key = name + "-" + color + "-" + texture;
        if (!treeMap.containsKey(key)) {
            treeMap.put(key, new Tree(name, color, texture)); // Create new if not present
            System.out.println("Creating new tree type: " + key);
        }
        return treeMap.get(key); // Return shared instance
    }
}

// Step 4: Use the Flyweight Pattern

public class Main {
    public static void main(String[] args) {
        // Shared tree types
        TreeType oakTree = TreeFactory.getTree("Oak", "Green", "Rough");
        TreeType pineTree = TreeFactory.getTree("Pine", "Dark Green", "Smooth");

        // Render trees with unique positions (extrinsic state)
        oakTree.render(10, 20);
        oakTree.render(15, 25);

        pineTree.render(30, 40);
        pineTree.render(35, 45);

        // Reuse shared tree types
        TreeType anotherOakTree = TreeFactory.getTree("Oak", "Green", "Rough");
        anotherOakTree.render(50, 60);
    }
}

// Output
// Creating new tree type: Oak-Green-Rough
// Creating new tree type: Pine-Dark Green-Smooth
// Rendering tree [Oak] at (10, 20) with color: Green
// Rendering tree [Oak] at (15, 25) with color: Green
// Rendering tree [Pine] at (30, 40) with color: Dark Green
// Rendering tree [Pine] at (35, 45) with color: Dark Green
// Rendering tree [Oak] at (50, 60) with color: Green

/** 
Key Points of the Example
    Flyweight (Tree): Represents the shared intrinsic state (e.g., tree name, color, texture).
    Factory (TreeFactory): Manages shared objects and ensures reuse.
    Extrinsic State: Unique attributes like position (x, y) are passed during operations.
    Client Code: Interacts with the factory and flyweight objects.
*/