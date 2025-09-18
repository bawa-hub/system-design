package _6_design_patterns._ashishp_notes.observer;
interface Subject {
    void attach(Observer observer);
    void detach(Observer observer);
    void notifyObservers();
}