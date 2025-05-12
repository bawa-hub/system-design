package _6_design_patterns._3_behavioral._9_state;

// Step 1: Define the State Interface

interface State {
    void doAction(Order order);
}

// Step 2: Create Concrete States

class PendingState implements State {
    @Override
    public void doAction(Order order) {
        System.out.println("Order is in pending state.");
        order.setState(this); // Set the current state to Pending
    }
}

class ShippedState implements State {
    @Override
    public void doAction(Order order) {
        System.out.println("Order has been shipped.");
        order.setState(this); // Set the current state to Shipped
    }
}

class DeliveredState implements State {
    @Override
    public void doAction(Order order) {
        System.out.println("Order has been delivered.");
        order.setState(this); // Set the current state to Delivered
    }
}

// Step 3: Create the Context (Order)

class Order {
    private State state;

    public Order() {
        state = new PendingState(); // Default state is Pending
    }

    public void setState(State state) {
        this.state = state;
    }

    public void processOrder() {
        state.doAction(this); // Delegate action to the current state
    }
}

// Step 4: Use the State Pattern

public class Main {
    public static void main(String[] args) {
        Order order = new Order();
        
        // Initial state is Pending
        order.processOrder();

        // Change state to Shipped
        order.setState(new ShippedState());
        order.processOrder();

        // Change state to Delivered
        order.setState(new DeliveredState());
        order.processOrder();
    }
}

// Output
// Order is in pending state.
// Order has been shipped.
// Order has been delivered.

/** 
Key Points of the Example
    State Interface (State): Defines the behavior for different states.
    Concrete States (PendingState, ShippedState, DeliveredState): Implement specific behavior for each state.
    Context (Order): The object that holds the current state and delegates behavior to it.
    State Transitions: The context changes its state by calling setState().
*/