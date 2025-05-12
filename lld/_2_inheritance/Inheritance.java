package _2_inheritance;
// Creating a new class based on an existing class to promote code reuse. 
// The new class, called the derived or child class, inherits attributes and behaviors (methods) from the base or parent class. 
// This allows the child class to use or extend the functionality of the parent class.

class Vehicle {
    public void start() {
        System.out.println("Vehicle started");
    }

    public void stop() {
        System.out.println("Vehicle stopped");
    }
}

class Car extends Vehicle {
    public void honk() {
        System.out.println("Car honking");
    }
}

public class Inheritance {
    public static void main(String[] args) {
        Car myCar = new Car();
        myCar.start();  // Inherited method from Vehicle class
        myCar.honk();   // Method from Car class
        myCar.stop();   // Inherited method from Vehicle class
    }
}
