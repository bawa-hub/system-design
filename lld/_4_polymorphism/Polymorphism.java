package _4_polymorphism;
// Allowing objects of different types to be treated as instances of the same class through a common interface. Polymorphism can be achieved in two ways:
//     Compile-time polymorphism (method overloading): Using the same method name with different parameter lists in the same class.
//     Runtime polymorphism (method overriding): When a child class provides a specific implementation of a method that is already defined in its parent class.

// Compile-time polymorphism (Method Overloading):

class Calculator {
    // Overloaded add methods
    public int add(int a, int b) {
        return a + b;
    }

    public double add(double a, double b) {
        return a + b;
    }
}

// Runtime polymorphism (Method Overriding):
class Animal {
    public void sound() {
        System.out.println("Some generic animal sound");
    }
}

class Cat extends Animal {
    @Override
    public void sound() {
        System.out.println("Cat meows");
    }
}

public class Polymorphism {
    public static void main(String[] args) {
        
        Calculator calc = new Calculator();
        System.out.println("Sum (int): " + calc.add(10, 20));
        System.out.println("Sum (double): " + calc.add(10.5, 20.5));

        Animal myAnimal = new Cat(); // Runtime polymorphism
        myAnimal.sound();  // Calls the overridden method in the Cat class
    }
}