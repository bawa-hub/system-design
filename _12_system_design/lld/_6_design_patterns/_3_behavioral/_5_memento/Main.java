package _6_design_patterns._3_behavioral._5_memento;

import java.util.Stack;

// Step 1: Define the Memento

class TextMemento {
    private final String state;

    public TextMemento(String state) {
        this.state = state;
    }

    public String getState() {
        return state;
    }
}

// Step 2: Create the Originator

class TextEditor {
    private String text;

    public void setText(String text) {
        this.text = text;
    }

    public String getText() {
        return text;
    }

    public TextMemento save() {
        return new TextMemento(text);
    }

    public void restore(TextMemento memento) {
        this.text = memento.getState();
    }
}

// Step 3: Create the Caretaker

class TextHistory {
    private Stack<TextMemento> history = new Stack<>();

    public void save(TextMemento memento) {
        history.push(memento);
    }

    public TextMemento undo() {
        if (!history.isEmpty()) {
            return history.pop();
        }
        return null;
    }
}

// Step 4: Use the Memento

public class Main {
    public static void main(String[] args) {
        TextEditor editor = new TextEditor();
        TextHistory history = new TextHistory();

        // Set initial text and save it
        editor.setText("Hello, World!");
        history.save(editor.save());
        System.out.println("Text: " + editor.getText());

        // Modify text and save it
        editor.setText("Hello, Design Patterns!");
        history.save(editor.save());
        System.out.println("Text: " + editor.getText());

        // Modify text without saving
        editor.setText("Hello, Memento!");
        System.out.println("Text: " + editor.getText());

        // Undo last change
        editor.restore(history.undo());
        System.out.println("After Undo: " + editor.getText());

        // Undo again
        editor.restore(history.undo());
        System.out.println("After Undo: " + editor.getText());
    }
}

// Output
// Text: Hello, World!
// Text: Hello, Design Patterns!
// Text: Hello, Memento!
// After Undo: Hello, Design Patterns!
// After Undo: Hello, World!

/** 
Key Points of the Example

    Memento (TextMemento): Captures and stores the state of the object.
    Originator (TextEditor): Creates and restores mementos.
    Caretaker (TextHistory): Manages mementos and provides undo functionality.
    Encapsulation: The memento only exposes the state to the originator.
*/

