package _6_design_patterns._ashishp_notes.abstractfactory;
public class FormalShoeFactory implements ShoeFactory {

    @Override
    public Sole createShoeSole() {
        return new FlatSole();
     }

    @Override
    public ShoeLace createShoeLace() {
        return new RoundShoeLace();
     }
    
}
