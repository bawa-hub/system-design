package _6_design_patterns._3_behavioral._4_mediator;

import java.util.ArrayList;
import java.util.List;


// Step 1: Define the Mediator Interface

interface ChatMediator {
    void sendMessage(String message, User user);
    void addUser(User user);
}

// Step 2: Create the Concrete Mediator

class ChatRoom implements ChatMediator {
    private List<User> users = new ArrayList<>();

    @Override
    public void addUser(User user) {
        users.add(user);
    }

    @Override
    public void sendMessage(String message, User sender) {
        for (User user : users) {
            if (user != sender) { // Don't send the message to the sender
                user.receiveMessage(message);
            }
        }
    }
}

// Step 3: Define the User Class

abstract class User {
    protected ChatMediator mediator;
    protected String name;

    public User(ChatMediator mediator, String name) {
        this.mediator = mediator;
        this.name = name;
    }

    public abstract void sendMessage(String message);
    public abstract void receiveMessage(String message);
}

// Step 4: Create Concrete Users

class ChatUser extends User {
    public ChatUser(ChatMediator mediator, String name) {
        super(mediator, name);
    }

    @Override
    public void sendMessage(String message) {
        System.out.println(name + " sends: " + message);
        mediator.sendMessage(message, this);
    }

    @Override
    public void receiveMessage(String message) {
        System.out.println(name + " receives: " + message);
    }
}

// Step 5: Use the Mediator

public class Main {
    public static void main(String[] args) {
        ChatMediator chatRoom = new ChatRoom();

        User user1 = new ChatUser(chatRoom, "Alice");
        User user2 = new ChatUser(chatRoom, "Bob");
        User user3 = new ChatUser(chatRoom, "Charlie");

        chatRoom.addUser(user1);
        chatRoom.addUser(user2);
        chatRoom.addUser(user3);

        user1.sendMessage("Hello everyone!");
        user2.sendMessage("Hi Alice!");
    }
}

// Output
// Alice sends: Hello everyone!
// Bob receives: Hello everyone!
// Charlie receives: Hello everyone!
// Bob sends: Hi Alice!
// Alice receives: Hi Alice!
// Charlie receives: Hi Alice!

/** 
Key Points of the Example
    Mediator (ChatRoom): Handles communication between users.
    Colleagues (Users): Communicate indirectly via the mediator.
    Decoupled Communication: Users don’t know about other users; they only interact with the mediator.
    Flexibility: New users or communication rules can be added without changing existing code.
*/