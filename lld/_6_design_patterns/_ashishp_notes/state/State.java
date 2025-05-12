package _6_design_patterns._ashishp_notes.state;
interface State {
    void insertDollar(VendingMachine context);
    void ejectMoney(VendingMachine context);
    void dispense(VendingMachine context);
}
