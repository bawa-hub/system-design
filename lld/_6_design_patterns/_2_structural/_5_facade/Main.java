package _6_design_patterns._2_structural._5_facade;

// Step 1: Create Subsystems

class DVDPlayer {
    public void turnOn() {
        System.out.println("DVD Player is ON.");
    }

    public void playMovie(String movie) {
        System.out.println("Playing movie: " + movie);
    }

    public void turnOff() {
        System.out.println("DVD Player is OFF.");
    }
}

class Projector {
    public void turnOn() {
        System.out.println("Projector is ON.");
    }

    public void setInput(String input) {
        System.out.println("Projector input set to: " + input);
    }

    public void turnOff() {
        System.out.println("Projector is OFF.");
    }
}

class SoundSystem {
    public void turnOn() {
        System.out.println("Sound System is ON.");
    }

    public void setVolume(int level) {
        System.out.println("Sound System volume set to: " + level);
    }

    public void turnOff() {
        System.out.println("Sound System is OFF.");
    }
}

// Step 2: Create the Facade

class HomeTheaterFacade {
    private DVDPlayer dvdPlayer;
    private Projector projector;
    private SoundSystem soundSystem;

    public HomeTheaterFacade(DVDPlayer dvdPlayer, Projector projector, SoundSystem soundSystem) {
        this.dvdPlayer = dvdPlayer;
        this.projector = projector;
        this.soundSystem = soundSystem;
    }

    public void watchMovie(String movie) {
        System.out.println("Get ready to watch a movie...");
        dvdPlayer.turnOn();
        projector.turnOn();
        projector.setInput("DVD");
        soundSystem.turnOn();
        soundSystem.setVolume(10);
        dvdPlayer.playMovie(movie);
    }

    public void stopMovie() {
        System.out.println("Shutting down the home theater...");
        dvdPlayer.turnOff();
        projector.turnOff();
        soundSystem.turnOff();
    }
}

// Step 3: Use the Facade

public class Main {
    public static void main(String[] args) {
        // Create subsystems
        DVDPlayer dvdPlayer = new DVDPlayer();
        Projector projector = new Projector();
        SoundSystem soundSystem = new SoundSystem();

        // Create facade
        HomeTheaterFacade homeTheater = new HomeTheaterFacade(dvdPlayer, projector, soundSystem);

        // Use facade
        homeTheater.watchMovie("Inception");
        homeTheater.stopMovie();
    }
}

// Output
// Get ready to watch a movie...
// DVD Player is ON.
// Projector is ON.
// Projector input set to: DVD
// Sound System is ON.
// Sound System volume set to: 10
// Playing movie: Inception
// Shutting down the home theater...
// DVD Player is OFF.
// Projector is OFF.
// Sound System is OFF.


// Key Points of the Example
//     Subsystems (DVDPlayer, Projector, SoundSystem): These are the complex components of the system.
//     Facade (HomeTheaterFacade): Provides a unified interface to the subsystems.
//     Client Code: Interacts only with the facade, not the subsystems directly.
