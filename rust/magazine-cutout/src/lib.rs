use std::collections::HashMap;

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mut got_words: HashMap<&str, usize> = HashMap::new();

    for magazine_word in magazine {
        let count = got_words.entry(magazine_word).or_insert(0);
        *count += 1;
    }

    for note_word in note {
        let count = got_words.entry(note_word).or_insert(0);

        if *count == 0 { return false };
        *count -= 1;
    }

    true
}