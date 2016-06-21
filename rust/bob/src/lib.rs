pub fn reply(phrase: &str) -> &str {
    match phrase {
        "" => "Fine. Be that way!",
        _ if is_shouting(phrase) => "Whoa, chill out!",
        _ if is_question(phrase) => "Sure.",
        _ => "Whatever."
    }
}

fn is_shouting(phrase: &str) -> bool {
    !phrase.is_empty() && phrase == phrase.to_uppercase()
}

fn is_question(phrase: &str) -> bool {
    phrase.chars().last().unwrap_or('-') == '?'
}
