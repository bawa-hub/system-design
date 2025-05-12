package _6_design_patterns._ashishp_notes.abstractfactory;
public class ThinSole implements Sole {

    @Override
    public String soleBuild() {
        return "Thin Plated";
     }

    @Override
    public String soleMaterial() {
        return "Rubber";
    }
    
    
}
