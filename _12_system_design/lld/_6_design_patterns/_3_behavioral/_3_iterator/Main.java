package _6_design_patterns._3_behavioral._3_iterator;

import java.util.List;
import java.util.ArrayList;


// Step 1: Define the Iterator Interface

interface Iterator<T> {
    boolean hasNext();
    T next();
}

// Step 2: Define the Collection Interface

interface Playlist {
    Iterator<Song> createIterator();
}

// Step 3: Create the Concrete Iterator

class PlaylistIterator implements Iterator<Song> {
    private List<Song> songs;
    private int position = 0;

    public PlaylistIterator(List<Song> songs) {
        this.songs = songs;
    }

    @Override
    public boolean hasNext() {
        return position < songs.size();
    }

    @Override
    public Song next() {
        if (this.hasNext()) {
            return songs.get(position++);
        }
        return null;
    }
}

// Step 4: Create the Concrete Collection

class SongPlaylist implements Playlist {
    private List<Song> songs = new ArrayList<>();

    public void addSong(Song song) {
        songs.add(song);
    }

    @Override
    public Iterator<Song> createIterator() {
        return new PlaylistIterator(songs);
    }
}

// Step 5: Use the Iterator

class Song {
    private String title;

    public Song(String title) {
        this.title = title;
    }

    @Override
    public String toString() {
        return title;
    }
}

public class Main {
    public static void main(String[] args) {
        // Create a playlist and add songs
        SongPlaylist playlist = new SongPlaylist();
        playlist.addSong(new Song("Song A"));
        playlist.addSong(new Song("Song B"));
        playlist.addSong(new Song("Song C"));

        // Get an iterator
        Iterator<Song> iterator = playlist.createIterator();

        // Iterate through songs
        while (iterator.hasNext()) {
            System.out.println("Playing: " + iterator.next());
        }
    }
}

// Output
// Playing: Song A
// Playing: Song B
// Playing: Song C

/**
 * 
 * Key Points of the Example

    Iterator Interface: Defines methods to traverse the collection (hasNext, next).
    Concrete Iterator: Implements the traversal logic for the specific collection.
    Collection Interface: Provides a method to get an iterator.
    Concrete Collection: Contains the elements and provides the iterator.
    
 */