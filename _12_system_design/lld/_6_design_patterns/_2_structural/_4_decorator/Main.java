package _6_design_patterns._2_structural._4_decorator;


// Step 1: Define the Component Interface

interface Text {
    String render(); // Common operation
}

// Step 2: Create the Concrete Component

class PlainText implements Text {
    private String content;

    public PlainText(String content) {
        this.content = content;
    }

    @Override
    public String render() {
        return content;
    }
}

// Step 3: Create the Abstract Decorator

abstract class TextDecorator implements Text {
    protected Text text; // Composition

    public TextDecorator(Text text) {
        this.text = text;
    }

    @Override
    public String render() {
        return text.render();
    }
}

// Step 4: Create Concrete Decorators

class BoldText extends TextDecorator {
    public BoldText(Text text) {
        super(text);
    }

    @Override
    public String render() {
        return "<b>" + text.render() + "</b>";
    }
}

class ItalicText extends TextDecorator {
    public ItalicText(Text text) {
        super(text);
    }

    @Override
    public String render() {
        return "<i>" + text.render() + "</i>";
    }
}

class UnderlineText extends TextDecorator {
    public UnderlineText(Text text) {
        super(text);
    }

    @Override
    public String render() {
        return "<u>" + text.render() + "</u>";
    }
}

// Step 5: Use the Decorator Pattern

public class Main {
    public static void main(String[] args) {
        // Base text
        Text plainText = new PlainText("Hello, World!");

        // Apply bold decoration
        Text boldText = new BoldText(plainText);

        // Apply italic decoration on top of bold
        Text italicBoldText = new ItalicText(boldText);

        // Apply underline decoration on top of italic + bold
        Text finalText = new UnderlineText(italicBoldText);

        System.out.println(finalText.render()); // Output: <u><i><b>Hello, World!</b></i></u>
    }
}

// Output
// <u><i><b>Hello, World!</b></i></u>


/** 
Key Points of the Example

    Component Interface (Text): Defines common operations for all components and decorators.
    Concrete Component (PlainText): The base object that can be wrapped by decorators.
    Abstract Decorator (TextDecorator): Wraps a component and forwards operations to it.
    Concrete Decorators (BoldText, ItalicText, UnderlineText): Add specific behaviors to the component.
*/    