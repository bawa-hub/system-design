package _6_design_patterns._ashishp_notes.command;
class Light {
    private boolean isOn = false;
    private String location;

    public Light(String location) {
        this.location = location;
    }

    public void turnOn() {
        isOn = true;
        System.out.println(location + " light is now ON");
    }

    public void turnOff() {
        isOn = false;
        System.out.println(location + " light is now OFF");
    }
}