package _6_design_patterns._ashishp_notes.abstractfactory;
public class CasualShoeFactory implements ShoeFactory {

    @Override
    public Sole createShoeSole() {
        return new ThinSole();
    }

    @Override
    public ShoeLace createShoeLace() {
        return new TapeShoeLace();
    }
    
}
