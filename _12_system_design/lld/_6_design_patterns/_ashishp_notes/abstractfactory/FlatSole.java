package _6_design_patterns._ashishp_notes.abstractfactory;
public class FlatSole implements Sole {

    @Override
    public String soleBuild() {
        return "Flat";
    }

    @Override
    public String soleMaterial() {
        return "Synthetic Rubber";
    }
    
}
