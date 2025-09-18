package _1_encapsulation;
// Bundling the data (variables) and the methods (functions) that operate on the data into a single unit or class. 
// It restricts direct access to some of an object's components, which helps to prevent accidental interference and misuse of the data. 
// Encapsulation provides control over data by using access specifiers such as private, protected, and public.

class Employee {
    private String name;
    private int age;

    // Constructor
    public Employee(String name, int age) {
        this.name = name;
        this.age = age;
    }

    // Getter for name
    public String getName() {
        return name;
    }

    // Setter for name
    public void setName(String name) {
        this.name = name;
    }

    // Getter for age
    public int getAge() {
        return age;
    }

    // Setter for age
    public void setAge(int age) {
        this.age = age;
    }
}

public class Encapsulation {
    public static void main(String[] args) {
        Employee emp = new Employee("John", 30);
        
        // Accessing private variables using getter methods
        System.out.println("Employee Name: " + emp.getName());
        System.out.println("Employee Age: " + emp.getAge());

        // Modifying private variables using setter methods
        emp.setAge(31);
        System.out.println("Updated Age: " + emp.getAge());
    }
}
