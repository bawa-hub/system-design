package _6_design_patterns._3_behavioral._7_visitor;

import java.util.ArrayList;
import java.util.List;

// Step 1: Define the Visitor Interface

interface Visitor {
    void visit(SalariedEmployee salariedEmployee);
    void visit(Contractor contractor);
}

// Step 2: Create the Element Interface

interface Employee {
    void accept(Visitor visitor);
}

// Step 3: Create Concrete Elements

class SalariedEmployee implements Employee {
    private String name;
    private double salary;

    public SalariedEmployee(String name, double salary) {
        this.name = name;
        this.salary = salary;
    }

    public String getName() {
        return name;
    }

    public double getSalary() {
        return salary;
    }

    @Override
    public void accept(Visitor visitor) {
        visitor.visit(this);
    }
}

class Contractor implements Employee {
    private String name;
    private double hourlyRate;
    private int hoursWorked;

    public Contractor(String name, double hourlyRate, int hoursWorked) {
        this.name = name;
        this.hourlyRate = hourlyRate;
        this.hoursWorked = hoursWorked;
    }

    public String getName() {
        return name;
    }

    public double getHourlyRate() {
        return hourlyRate;
    }

    public int getHoursWorked() {
        return hoursWorked;
    }

    @Override
    public void accept(Visitor visitor) {
        visitor.visit(this);
    }
}

// Step 4: Create a Concrete Visitor

class TaxCalculatorVisitor implements Visitor {
    @Override
    public void visit(SalariedEmployee salariedEmployee) {
        double tax = salariedEmployee.getSalary() * 0.2; // Example tax rate
        System.out.println("Tax for salaried employee " + salariedEmployee.getName() + ": $" + tax);
    }

    @Override
    public void visit(Contractor contractor) {
        double tax = contractor.getHourlyRate() * contractor.getHoursWorked() * 0.1; // Example tax rate
        System.out.println("Tax for contractor " + contractor.getName() + ": $" + tax);
    }
}

// Step 5: Use the Visitor Pattern

public class Main {
    public static void main(String[] args) {
        List<Employee> employees = new ArrayList<>();
        employees.add(new SalariedEmployee("Alice", 50000));
        employees.add(new Contractor("Bob", 50, 100));

        Visitor taxCalculator = new TaxCalculatorVisitor();

        for (Employee employee : employees) {
            employee.accept(taxCalculator);
        }
    }
}

// Output
// Tax for salaried employee Alice: $10000.0
// Tax for contractor Bob: $500.0

/** 
Key Points of the Example

    Elements (SalariedEmployee, Contractor): The objects being visited.
    Visitor (TaxCalculatorVisitor): Performs operations on the elements without changing their structure.
    Dynamic Behavior: Adding a new operation is as simple as creating another visitor, e.g., PayrollVisitor.

*/