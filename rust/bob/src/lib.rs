pub fn reply(phrase: &str) -> &str {
    match phrase {
        "" => "Fine. Be that way!",
        _ if is_shouting(phrase) => "Whoa, chill out!",
        _ if phrase.ends_with("?") => "Sure.",
        _ => "Whatever."
    }
}

fn is_shouting(phrase: &str) -> bool {
    !phrase.is_empty() && phrase == phrase.to_uppercase()
}
