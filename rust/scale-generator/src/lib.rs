use core::fmt;
use std::{fmt::Display, str::FromStr};

use Note::*;
use Key::*;

#[derive(Debug)]
pub enum Error {
    BadNote,
    BadInterval
}

#[derive(Clone, Copy, PartialEq, Eq)]
pub enum Note {
    C,
    Csharp,
    Db,
    D,
    Dsharp,
    Eb,
    E,
    F,
    Fsharp,
    Gb,
    G,
    Gsharp,
    Ab,
    A,
    Asharp,
    Bb,
    B
}

impl Display for Note {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            C => write!(f, "C"),
            Csharp => write!(f, "C#"),
            Db => write!(f, "Db"),
            D => write!(f, "D"),
            Dsharp => write!(f, "D#"),
            Eb => write!(f, "Eb"),
            E => write!(f, "E"),
            F => write!(f, "F"),
            Fsharp => write!(f, "F#"),
            Gb => write!(f, "Gb"),
            G => write!(f, "G"),
            Gsharp => write!(f, "G#"),
            Ab => write!(f, "Ab"),
            A => write!(f, "A"),
            Asharp => write!(f, "A#"),
            Bb => write!(f, "Bb"),
            B => write!(f, "B"),
        }
    }
}

impl FromStr for Note {
    type Err = Error;

    fn from_str(note_str: &str) -> Result<Note, Error> {
        match note_str {
            "c"  | "C"  => Ok(C),
            "c#" | "C#" => Ok(Csharp),
            "db" | "Db" => Ok(Db),
            "d"  | "D"  => Ok(D),
            "d#" | "D#" => Ok(Dsharp),
            "eb" | "Eb" => Ok(Eb),
            "e"  | "E"  => Ok(E),
            "f"  | "F"  => Ok(F),
            "f#" | "F#" => Ok(Fsharp),
            "gb" | "Gb" => Ok(Gb),
            "g"  | "G"  => Ok(G),
            "g#" | "G#" => Ok(Gsharp),
            "ab" | "Ab" => Ok(Ab),
            "a"  | "A"  => Ok(A),
            "a#" | "A#" => Ok(Asharp),
            "bb" | "Bb" => Ok(Bb),
            "b"  | "B"  => Ok(B),
            _ => Err(Error::BadNote)
        }
    }
}

#[derive(Clone, Copy, PartialEq, Eq)]
pub enum Key {
    Major, Minor
}

#[derive(Clone, Copy, PartialEq, Eq)]
struct Tonic {
    note: Note,
    key: Key
}

impl FromStr for Tonic {
    type Err = Error;

    fn from_str(tonic_str: &str) -> Result<Tonic, Error> {
        let note = Note::from_str(tonic_str);

        let key = match tonic_str.chars().nth(0) {
            Some('a'..='g') => Ok(Minor),
            Some('A'..='G') => Ok(Major),
            _ => Err(Error::BadNote)
        };

        match (note, key) {
            (Ok(note), Ok(key)) => Ok(Tonic { note, key }),
            _ => Err(Error::BadNote)
        }
    }
}

static SHARPS: &'static [Note] = &[
    A,
    Asharp,
    B,
    C,
    Csharp,
    D,
    Dsharp,
    E,
    F,
    Fsharp,
    G,
    Gsharp,
];

static FLATS: &'static [Note] = &[
    A,
    Bb,
    B,
    C,
    Db,
    D,
    Eb,
    E,
    F,
    Gb,
    G,
    Ab,
];

#[derive(Clone, Copy)]
struct ScaleIterator<'a> {
    index: usize,
    scale: &'static [Note],
    intervals: &'a str,
    interval_index: usize,
}

impl ScaleIterator<'_> {
    pub fn new<'a>(tonic: &str, intervals: &'a str) -> Result<ScaleIterator<'a>, Error> {
        let tonic = Tonic::from_str(tonic)?;

        let scale = match (tonic.note, tonic.key) {
            (C, Major) | (G, Major) | (D, Major) | (A, _) | (E, Major) | (B, Major) |
            (Fsharp, Major) | (E, Minor) | (B, Minor) | (Fsharp, Minor) |
            (Csharp, Minor) | (Gsharp, Minor) | (Dsharp, Minor) => SHARPS,
            _ => FLATS
        };

        if let Some(index) = scale.iter().position(|n| tonic.note == *n) {
            Ok(ScaleIterator { index, scale, intervals, interval_index: 0 })
        } else {
            Err(Error::BadNote)
        }
    }
}

impl<'a> Iterator for ScaleIterator<'a> {
    type Item = &'a Note;

    fn next(&mut self) -> Option<Self::Item> {
        let note = &self.scale[self.index];

        if let Some(interval) = self.intervals.chars().nth(self.interval_index) {
            match interval {
                'm' => self.index += 1,
                'M' => self.index += 2,
                'A' => self.index += 3,
                _ => panic!("Bad interval '{}'", interval)
            }
        } else if self.interval_index > self.intervals.len() {
            return None;
        }

        self.index = self.index % self.scale.len();
        self.interval_index += 1;

        Some(note)
    }
}

pub struct Scale<'a> {
    iterator: ScaleIterator<'a>,
}

impl Scale<'_> {
    pub fn new<'a>(tonic: &'a str, intervals: &'a str) -> Result<Scale<'a>, Error> {
        let iterator = ScaleIterator::new(tonic, intervals)?;
        Ok(Scale { iterator })
    }

    pub fn chromatic<'a>(tonic: &'a str) -> Result<Scale<'a>, Error> {
        Scale::new(tonic, "mmmmmmmmmmmm")
    }

    pub fn enumerate(&self) -> Vec<String> {
        self.iterator.map(|note| note.to_string()).collect()
    }
}