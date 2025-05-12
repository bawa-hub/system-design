package _6_design_patterns._ashishp_notes.abstractfactory;
public class TapeShoeLace implements ShoeLace{

    @Override
    public String shoeLaceBuild() {
        return "Tape Flat";
    }

    @Override
    public String shoeLaceMaterial() {
        return "Synthetic Cotton";
    }
    
}
