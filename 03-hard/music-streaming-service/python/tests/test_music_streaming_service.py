from artist import Artist
from music_streaming_system import MusicStreamingSystem


def test_music_streaming_system() -> None:
    system = MusicStreamingSystem()
    artist = Artist("art1", "Daft Punk")
    system.add_artist(artist)
    song = system.add_song("s1", "One More Time", "art1", 320)

    assert song.title == "One More Time"
    assert system.get_player() is not None


def test_search_songs_by_title() -> None:
    system = MusicStreamingSystem()
    system.add_artist(Artist("art1", "Daft Punk"))
    system.add_song("s1", "One More Time", "art1", 320)
    results = system.search_songs_by_title("One")
    assert len(results) == 1
