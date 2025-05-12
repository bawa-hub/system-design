package _3_abstraction;
// Hiding complex implementation details and showing only the essential features of the object. 
// This allows developers to interact with objects through simplified interfaces without needing to understand all the internal workings. 
// Abstraction is achieved through abstract classes and interfaces.

import _4_polymorphism.Animal;

abstract class Animal {
    // Abstract method (no implementation)
    public abstract void sound();

    // Regular method
    public void sleep() {
        System.out.println("The animal is sleeping");
    }
}

class Dog extends Animal {
    // Provide implementation for the abstract method
    public void sound() {
        System.out.println("Dog barks");
    }
}

public class Abstraction {
    public static void main(String[] args) {
        Animal dog = new Dog();
        dog.sound();  // Calls the implementation in the Dog class
        dog.sleep();  // Calls the common method from the Animal class
    }
}
