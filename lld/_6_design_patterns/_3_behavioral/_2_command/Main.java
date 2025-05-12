package _6_design_patterns._3_behavioral._2_command;

import java.util.Stack;

// Step 1: Define the Command Interface

interface Command {
    void execute();
    void undo();
}

// Step 2: Create Concrete Commands

class WriteCommand implements Command {
    private String text;
    private StringBuilder document;

    public WriteCommand(StringBuilder document, String text) {
        this.document = document;
        this.text = text;
    }

    @Override
    public void execute() {
        document.append(text);
    }

    @Override
    public void undo() {
        document.delete(document.length() - text.length(), document.length());
    }
}

class DeleteCommand implements Command {
    private StringBuilder document;
    private String deletedText;

    public DeleteCommand(StringBuilder document, int length) {
        this.document = document;
        this.deletedText = document.substring(document.length() - length);
    }

    @Override
    public void execute() {
        document.delete(document.length() - deletedText.length(), document.length());
    }

    @Override
    public void undo() {
        document.append(deletedText);
    }
}

// Step 3: Create the Invoker

class CommandInvoker {
    private Stack<Command> history = new Stack<>();

    public void executeCommand(Command command) {
        command.execute();
        history.push(command);
    }

    public void undoCommand() {
        if (!history.isEmpty()) {
            history.pop().undo();
        }
    }
}

// Step 4: Use the Command Pattern

public class Main {
    public static void main(String[] args) {
        StringBuilder document = new StringBuilder();
        CommandInvoker invoker = new CommandInvoker();

        // Write text
        Command writeCommand1 = new WriteCommand(document, "Hello ");
        invoker.executeCommand(writeCommand1);

        Command writeCommand2 = new WriteCommand(document, "World!");
        invoker.executeCommand(writeCommand2);

        System.out.println("Document: " + document); // Output: Hello World!

        // Undo last write
        invoker.undoCommand();
        System.out.println("After Undo: " + document); // Output: Hello 

        // Delete text
        Command deleteCommand = new DeleteCommand(document, 6);
        invoker.executeCommand(deleteCommand);

        System.out.println("After Delete: " + document); // Output: (empty)

        // Undo delete
        invoker.undoCommand();
        System.out.println("After Undo Delete: " + document); // Output: Hello
    }
}

/** 
Key Points of the Example
    Command Interface: Defines the structure for all commands (execute and undo methods).
    Concrete Commands: Implement specific operations (WriteCommand, DeleteCommand).
    Invoker: Executes and keeps track of commands to support undo functionality.
    Receiver: The StringBuilder acts as the receiver performing the actual operations.
*/    