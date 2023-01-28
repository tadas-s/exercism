pub fn reply(phrase: &str) -> &str {
    let shouting = phrase.contains(char::is_alphabetic) && phrase == phrase.to_uppercase();
    let trimmed = phrase.trim();
    let questioning = trimmed.ends_with("?");

    match trimmed {
        "" => "Fine. Be that way!",
        _ if questioning && shouting => "Calm down, I know what I'm doing!",
        _ if shouting => "Whoa, chill out!",
        _ if questioning => "Sure.",
        _ => "Whatever.",
    }
}
