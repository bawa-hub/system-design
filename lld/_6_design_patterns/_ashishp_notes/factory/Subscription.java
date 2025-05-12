package _6_design_patterns._ashishp_notes.factory;

public interface Subscription {
     public String subscriptionType();
     public boolean addSubscription(Customer customer);
     public boolean removeSubscription(Customer customer);
     public boolean updateSubscription(Customer customer);

}
