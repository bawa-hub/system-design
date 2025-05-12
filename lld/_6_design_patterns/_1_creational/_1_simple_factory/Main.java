package _6_design_patterns._1_creational._1_simple_factory;

// Factory class
class AnimalFactory {
    public static Animal createAnimal(String type) {
        if (type.equals("Dog")) {
            return new Dog();
        } else if (type.equals("Cat")) {
            return new Cat();
        } else if (type.equals("Duck")) {
            return new Duck();
        }
        return null;
    }
}

// Animal interface (all animals will follow this)
interface Animal {
    void speak();
}

// Dog class
class Dog implements Animal {
    public void speak() {
        System.out.println("Woof!");
    }
}

// Cat class
class Cat implements Animal {
    public void speak() {
        System.out.println("Meow!");
    }
}

// Duck class
class Duck implements Animal {
    public void speak() {
        System.out.println("Quack!");
    }
}

// Main class
public class Main {
    public static void main(String[] args) {

        // Without a Factory:
        // You want a dog. You have to make it yourself:
        // "Okay, I’ll make a dog... I'll give it a name and teach it how to bark."
        // This takes time, and if you want a cat later, you have to make the cat yourself too.

        // Dog dog = new Dog();
        // dog.speak();
        // Cat cat = new Cat();
        // cat.speak();


        // With a Factory:
        // You have a magic factory (like a pizza shop). You just say:
        // "Hey factory, I want a dog!"
        // The factory gives you a ready-made dog.
        // Later, if you want a cat, you just say:
        // "Hey factory, I want a cat!"


        Animal dog = AnimalFactory.createAnimal("Dog");
        dog.speak(); // Factory gives you a Dog, and it barks!

        Animal cat = AnimalFactory.createAnimal("Cat");
        cat.speak(); // Factory gives you a Cat, and it meows!

    }
}
