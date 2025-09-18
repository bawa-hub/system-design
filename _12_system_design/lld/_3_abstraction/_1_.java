package _3_abstraction;

public class _1_ {
    public static void main(String[] args) {
        Animal a = new Dog();
        a.makeSound();   // Output: Bark

    }
}

/**
 
 What is Abstraction?
   Abstraction is the process of hiding internal details and showing only essential features to the user.

    It answers:
    "What an object does, not how it does it."

    Real-Life Analogy

    Car:
    You press the accelerator to speed up, but you don’t know the internal mechanism (fuel injection, combustion, etc.).
    That’s abstraction.

    ATM Machine:
    You interact with buttons and screen, not the internal transaction logic.

    Why Use Abstraction?

    To reduce complexity
    To increase reusability
    To isolate changes (only internal implementation changes)
    To enforce contracts via interfaces or abstract classes

    How to Achieve Abstraction in Java?

    There are two primary ways:

    1. Abstract Classes

    A class declared with the abstract keyword.
    Can have abstract methods (without body) and concrete methods (with body).
    Cannot be instantiated directly.

    2. Interfaces

    A completely abstract type.
    All methods are implicitly public and abstract (until Java 7).
    From Java 8+, interfaces can have default and static methods.
    A class can implement multiple interfaces.

 */


// using abstract class 
abstract class Animal {
    abstract void makeSound();   // abstract method

    void breathe() {
        System.out.println("Breathing...");
    }
}

class Dog extends Animal {
    void makeSound() {
        System.out.println("Bark");
    }
}

// using interface
interface Vehicle {
    void drive();   // implicitly public abstract
}

class Car implements Vehicle {
    public void drive() {
        System.out.println("Driving car...");
    }
}
