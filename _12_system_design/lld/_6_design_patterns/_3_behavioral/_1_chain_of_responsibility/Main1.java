package _6_design_patterns._3_behavioral._1_chain_of_responsibility;


/** 
Scenario

Imagine you’re building a support ticket system where tickets are categorized based on their severity. For example:

    Low Priority: Handled by a Level 1 Support Agent.
    Medium Priority: Escalated to a Level 2 Support Agent.
    High Priority: Escalated to a Manager.

The Chain of Responsibility Pattern can be used to process these tickets.

*/

// Step 1: Define the Handler Interface
abstract class SupportHandler {
    protected SupportHandler nextHandler;

    public void setNextHandler(SupportHandler nextHandler) {
        this.nextHandler = nextHandler;
    }

    public void handleRequest(String issue, int priority) {
        if (canHandle(priority)) {
            resolveIssue(issue);
        } else if (nextHandler != null) {
            nextHandler.handleRequest(issue, priority);
        } else {
            System.out.println("No one is able to handle this issue.");
        }
    }

    protected abstract boolean canHandle(int priority);
    protected abstract void resolveIssue(String issue);
}

// Step 2: Create Concrete Handlers
class Level1Support extends SupportHandler {
    @Override
    protected boolean canHandle(int priority) {
        return priority == 1; // Low Priority
    }

    @Override
    protected void resolveIssue(String issue) {
        System.out.println("Level 1 Support resolved issue: " + issue);
    }
}

class Level2Support extends SupportHandler {
    @Override
    protected boolean canHandle(int priority) {
        return priority == 2; // Medium Priority
    }

    @Override
    protected void resolveIssue(String issue) {
        System.out.println("Level 2 Support resolved issue: " + issue);
    }
}

class ManagerSupport extends SupportHandler {
    @Override
    protected boolean canHandle(int priority) {
        return priority == 3; // High Priority
    }

    @Override
    protected void resolveIssue(String issue) {
        System.out.println("Manager resolved critical issue: " + issue);
    }
}

// Step 3: Create the Chain
class SupportChain {
    public static SupportHandler createSupportChain() {
        SupportHandler level1 = new Level1Support();
        SupportHandler level2 = new Level2Support();
        SupportHandler manager = new ManagerSupport();

        level1.setNextHandler(level2);
        level2.setNextHandler(manager);

        return level1;
    }
}

// Step 4: Use the Chain
public class Main1 {
    public static void main(String[] args) {
        SupportHandler supportChain = SupportChain.createSupportChain();

        // Handle tickets with different priorities
        supportChain.handleRequest("Forgot password", 1); // Handled by Level 1
        supportChain.handleRequest("System crash", 3);    // Handled by Manager
        supportChain.handleRequest("Bug in app", 2);      // Handled by Level 2
    }
}

// Output
// Level 1 Support resolved issue: Forgot password
// Manager resolved critical issue: System crash
// Level 2 Support resolved issue: Bug in app

// Key Points of the Example

//     Handlers: Each handler (Level1Support, Level2Support, ManagerSupport) handles a specific type of request based on its priority.
//     Chain: If one handler can’t process the request, it forwards it to the next handler.
//     Flexibility: You can easily add more handlers (e.g., Level 3 Support) without modifying existing ones.
//     Real-World Application: Common in customer service ticket systems where escalation is required.