package _6_design_patterns._ashishp_notes.facade;
class Memory {
    public void load(long position, byte[] data) {
        System.out.println("Memory: Loading data at position " + position);
    }
}