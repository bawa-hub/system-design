package _6_design_patterns._1_creational._5_prototype;

// Step 1: Define the Prototype Interface

// An interface with a clone method.
interface GameCharacter extends Cloneable {
    GameCharacter clone();
}

// Step 2: Create a Concrete Prototype Class

// Define the character's properties and cloning logic.
class Warrior implements GameCharacter {
    private String name;
    private int health;
    private int attackPower;

    public Warrior(String name, int health, int attackPower) {
        this.name = name;
        this.health = health;
        this.attackPower = attackPower;
    }

    // Clone method to create a copy
    @Override
    public Warrior clone() {
        return new Warrior(this.name, this.health, this.attackPower);
    }

    // Setters for customization
    public void setName(String name) {
        this.name = name;
    }

    @Override
    public String toString() {
        return "Warrior{name='" + name + "', health=" + health + ", attackPower=" + attackPower + "}";
    }
}

// Step 3: Use the Prototype to Create and Customize Characters

// Clone the prototype and customize the copies.
public class Main {
    public static void main(String[] args) {
        // Create a prototype
        Warrior prototypeWarrior = new Warrior("Prototype Warrior", 100, 50);

        // Clone the prototype and customize it
        Warrior warrior1 = prototypeWarrior.clone();
        warrior1.setName("Knight of Light");

        Warrior warrior2 = prototypeWarrior.clone();
        warrior2.setName("Shadow Assassin");

        // Print the customized warriors
        System.out.println(warrior1); // Output: Warrior{name='Knight of Light', health=100, attackPower=50}
        System.out.println(warrior2); // Output: Warrior{name='Shadow Assassin', health=100, attackPower=50}
    }
}
