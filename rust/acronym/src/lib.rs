use State::*;

#[derive(Clone, Copy, Debug)]
enum State {
    Seeking,
    WordOrAbbr,
    Word,
    Abbr
}

pub fn abbreviate(phrase: &str) -> String {
    let (_last_state, abbreviation) =
        phrase
            .chars()
            .fold((Seeking, "".to_string()), |(state, abbr), c| {
                match (state, c) {
                    (WordOrAbbr, c) if c.is_alphabetic() && c.is_uppercase() => (Abbr, abbr),
                    (WordOrAbbr, c) if c.is_alphabetic() && c.is_lowercase() => (Word, abbr),
                    (WordOrAbbr, c) if !c.is_alphabetic() => (Seeking, abbr),
                    (Abbr, c) if c.is_alphabetic() && c.is_uppercase() => (Abbr, abbr),
                    (Abbr, c) if !c.is_alphabetic() => (Seeking, abbr),
                    (Word, c) if c.is_alphabetic() && c.is_uppercase() => (Word, abbr + &c.to_string()),
                    (Word, c) if c.is_alphabetic() => (Word, abbr),
                    (Word, c) if c.is_whitespace() || (c == '-') => (Seeking, abbr),
                    (Word, _) => (Word, abbr),
                    (Seeking, c) if !c.is_alphabetic() => (Seeking, abbr),
                    (Seeking, c) if c.is_alphabetic() && c.is_uppercase() => (WordOrAbbr, abbr + &c.to_string()),
                    (Seeking, c) if c.is_alphabetic() => (Word, abbr + &c.to_string()),
                    _ => panic!("Unhandled state: ({state:?},'{c}')"),
                }
            });

    abbreviation.to_uppercase()
}