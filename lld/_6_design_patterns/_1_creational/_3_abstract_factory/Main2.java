package _6_design_patterns._1_creational._3_abstract_factory;


/**
 * 
 * Real-World Example: Cross-Platform UI Components
Problem:

You are building a cross-platform UI framework that supports different operating systems (Windows, macOS, Linux). Each OS has its own style of UI components (e.g., Buttons and Checkboxes). Your application should be able to work with these components seamlessly.
Abstract Factory for UI Components:

    Product Families: Button, Checkbox.
    Concrete Families: WindowsButton, WindowsCheckbox; MacButton, MacCheckbox.

 */

 // Common Interfaces
interface Button {
    void click();
}

interface Checkbox {
    void check();
}

// Windows Implementations
class WindowsButton implements Button {
    public void click() {
        System.out.println("Windows Button clicked!");
    }
}

class WindowsCheckbox implements Checkbox {
    public void check() {
        System.out.println("Windows Checkbox checked!");
    }
}

// Mac Implementations
class MacButton implements Button {
    public void click() {
        System.out.println("Mac Button clicked!");
    }
}

class MacCheckbox implements Checkbox {
    public void check() {
        System.out.println("Mac Checkbox checked!");
    }
}

// Abstract Factory
interface GUIFactory {
    Button createButton();
    Checkbox createCheckbox();
}

// Concrete Factories
class WindowsFactory implements GUIFactory {
    public Button createButton() {
        return new WindowsButton();
    }

    public Checkbox createCheckbox() {
        return new WindowsCheckbox();
    }
}

class MacFactory implements GUIFactory {
    public Button createButton() {
        return new MacButton();
    }

    public Checkbox createCheckbox() {
        return new MacCheckbox();
    }
}

// Client Code
public class Main2 {
    public static void main(String[] args) {
        // Create a Windows Factory
        GUIFactory windowsFactory = new WindowsFactory();
        Button windowsButton = windowsFactory.createButton();
        Checkbox windowsCheckbox = windowsFactory.createCheckbox();
        windowsButton.click(); // Output: Windows Button clicked!
        windowsCheckbox.check(); // Output: Windows Checkbox checked!

        // Create a Mac Factory
        GUIFactory macFactory = new MacFactory();
        Button macButton = macFactory.createButton();
        Checkbox macCheckbox = macFactory.createCheckbox();
        macButton.click(); // Output: Mac Button clicked!
        macCheckbox.check(); // Output: Mac Checkbox checked!
    }
}
