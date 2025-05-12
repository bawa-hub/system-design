package _6_design_patterns._ashishp_notes.builder;
interface PizzaBuilder {
    void buildDough();
    void buildSauce();
    void buildTopping();
    Pizza getPizza();
}