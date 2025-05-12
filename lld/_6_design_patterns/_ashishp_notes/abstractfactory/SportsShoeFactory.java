package _6_design_patterns._ashishp_notes.abstractfactory;
public class SportsShoeFactory implements ShoeFactory {

    @Override
    public Sole createShoeSole() {
        return new BumpySole();
    }

    @Override
    public ShoeLace createShoeLace() {
        return new RoundShoeLace();
     }
    
}
