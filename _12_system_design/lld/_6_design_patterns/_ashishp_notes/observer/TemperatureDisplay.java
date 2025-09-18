package _6_design_patterns._ashishp_notes.observer;
class TemperatureDisplay implements Observer {
    private String name;

    public TemperatureDisplay(String name) {
        this.name = name;
    }

    @Override
    public void update(float temperature) {
        System.out.println(name + " displays Temperature: " + temperature + " degrees Celsius");
    }
}
