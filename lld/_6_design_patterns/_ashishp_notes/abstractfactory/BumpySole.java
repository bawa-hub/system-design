package _6_design_patterns._ashishp_notes.abstractfactory;
public class BumpySole implements Sole {

    @Override
    public String soleBuild() {
        return "Bummpy";
    }

    @Override
    public String soleMaterial() {
        return "Plastic Rubber";
    }
    
}
