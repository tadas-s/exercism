pub fn reply(phrase: &str) -> &str {
    match phrase {
        "" => "Fine. Be that way!",
        _ if phrase == phrase.to_uppercase() => "Whoa, chill out!",
        _ if phrase.ends_with("?") => "Sure.",
        _ => "Whatever."
    }
}

