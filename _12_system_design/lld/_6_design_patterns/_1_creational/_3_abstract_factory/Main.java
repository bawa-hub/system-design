package _6_design_patterns._1_creational._3_abstract_factory;

// Step 1: Define Common Interfaces for Product Families

// Chair interface
interface Chair {
    void sitOn();
}

// Table interface
interface Table {
    void use();
}


// Step 2: Create Concrete Product Implementations

// Modern Furniture
class ModernChair implements Chair {
    public void sitOn() {
        System.out.println("Sitting on a modern chair.");
    }
}

class ModernTable implements Table {
    public void use() {
        System.out.println("Using a modern table.");
    }
}

// Victorian Furniture
class VictorianChair implements Chair {
    public void sitOn() {
        System.out.println("Sitting on a Victorian chair.");
    }
}

class VictorianTable implements Table {
    public void use() {
        System.out.println("Using a Victorian table.");
    }
}

// Art Deco Furniture
class ArtDecoChair implements Chair {
    public void sitOn() {
        System.out.println("Sitting on an Art Deco chair.");
    }
}

class ArtDecoTable implements Table {
    public void use() {
        System.out.println("Using an Art Deco table.");
    }
}


// Step 3: Define the Abstract Factory

interface FurnitureFactory {
    Chair createChair();
    Table createTable();
}


// Step 4: Create Concrete Factories

// Modern Furniture Factory
class ModernFurnitureFactory implements FurnitureFactory {
    public Chair createChair() {
        return new ModernChair();
    }

    public Table createTable() {
        return new ModernTable();
    }
}

// Victorian Furniture Factory
class VictorianFurnitureFactory implements FurnitureFactory {
    public Chair createChair() {
        return new VictorianChair();
    }

    public Table createTable() {
        return new VictorianTable();
    }
}

// Art Deco Furniture Factory
class ArtDecoFurnitureFactory implements FurnitureFactory {
    public Chair createChair() {
        return new ArtDecoChair();
    }

    public Table createTable() {
        return new ArtDecoTable();
    }
}


// Step 5: Use the Abstract Factory in Your Application

public class Main {
    public static void main(String[] args) {
        // Create a Modern Furniture Factory
        FurnitureFactory modernFactory = new ModernFurnitureFactory();
        Chair modernChair = modernFactory.createChair();
        Table modernTable = modernFactory.createTable();
        modernChair.sitOn(); // Output: Sitting on a modern chair.
        modernTable.use();   // Output: Using a modern table.

        // Create a Victorian Furniture Factory
        FurnitureFactory victorianFactory = new VictorianFurnitureFactory();
        Chair victorianChair = victorianFactory.createChair();
        Table victorianTable = victorianFactory.createTable();
        victorianChair.sitOn(); // Output: Sitting on a Victorian chair.
        victorianTable.use();   // Output: Using a Victorian table.

        // Create an Art Deco Furniture Factory
        FurnitureFactory artDecoFactory = new ArtDecoFurnitureFactory();
        Chair artDecoChair = artDecoFactory.createChair();
        Table artDecoTable = artDecoFactory.createTable();
        artDecoChair.sitOn(); // Output: Sitting on an Art Deco chair.
        artDecoTable.use();   // Output: Using an Art Deco table.
    }
}

