use std::collections::HashSet;
use unicode_segmentation::UnicodeSegmentation;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &'a [&str]) -> HashSet<&'a str> {
    let mut anagrams: HashSet<&str> = HashSet::new();
    let sorted_word = sorted_word_graphemes(word);

    for possible_anagram in possible_anagrams {
        let sorted_possible_anagram = sorted_word_graphemes(possible_anagram);

        if word == *possible_anagram || word.to_lowercase() == possible_anagram.to_lowercase() {
            continue;
        }

        if sorted_word == sorted_possible_anagram {
            anagrams.insert(possible_anagram);
        }
    }

    anagrams
}

fn sorted_word_graphemes(word: &str) -> String {
    let lowercased = word.to_lowercase();
    let mut graphemes = lowercased.graphemes(true).collect::<Vec<&str>>();
    graphemes.sort_unstable();
    graphemes.join("")
}