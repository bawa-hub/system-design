package _6_design_patterns._ashishp_notes.singleton;
class EagerSingleton {
    // The single instance, created immediately
    private static final EagerSingleton instance = new EagerSingleton();
    
    // Private constructor to prevent instantiation
    private EagerSingleton() {}
    
    // Public method to get the instance
    public static EagerSingleton getInstance() {
        return instance;
    }
}