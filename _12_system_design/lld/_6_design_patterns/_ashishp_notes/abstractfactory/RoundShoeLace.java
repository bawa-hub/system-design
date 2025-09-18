package _6_design_patterns._ashishp_notes.abstractfactory;
public class RoundShoeLace implements ShoeLace {

    @Override
    public String shoeLaceBuild() {
        return "Round";
    }

    @Override
    public String shoeLaceMaterial() {
        return "Synthetic Polyster";
    }
    
}
