use State::*;
/*
  Quite possibly slightly less usual solution - phrase is
  scaned using a state machine.
 */
#[derive(Clone, Copy, Debug)]
enum State {
    Separation,
    WordOrAbbr,
    Word,
    Abbr
}

pub fn abbreviate(phrase: &str) -> String {
    let (_last_state, abbreviation) =
        phrase
            .chars()
            .into_iter()
            .fold((Separation, "".to_string()), |(state, abbr), c| {
                match (state, c) {
                    (WordOrAbbr, c) if c.is_alphabetic() && c.is_uppercase() => (Abbr, abbr),
                    (WordOrAbbr, c) if c.is_alphabetic() && c.is_lowercase() => (Word, abbr),
                    (WordOrAbbr, c) if !c.is_alphabetic() => (Separation, abbr),
                    (Abbr, c) if c.is_alphabetic() && c.is_uppercase() => (Abbr, abbr),
                    (Abbr, c) if !c.is_alphabetic() => (Separation, abbr),
                    (Word, c) if c.is_alphabetic() && c.is_uppercase() => (Word, abbr + &c.to_string()),
                    (Word, c) if c.is_alphabetic() => (Word, abbr),
                    (Word, c) if c.is_whitespace() || (c == '-') => (Separation, abbr),
                    (Word, _) => (Word, abbr),
                    (Separation, c) if !c.is_alphabetic() => (Separation, abbr),
                    (Separation, c) if c.is_alphabetic() && c.is_uppercase() => (WordOrAbbr, abbr + &c.to_string()),
                    (Separation, c) if c.is_alphabetic() => (Word, abbr + &c.to_string()),
                    _ => panic!("Unhandled state: ({:?},'{}')", state, c),
                }
            });

    abbreviation.to_uppercase()
}
