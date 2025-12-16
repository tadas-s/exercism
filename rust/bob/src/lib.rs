pub fn reply(message: &str) -> &str {
    let shouting = message.contains(char::is_alphabetic) && message == message.to_uppercase();
    let trimmed = message.trim();
    let questioning = trimmed.ends_with("?");

    match trimmed {
        "" => "Fine. Be that way!",
        _ if questioning && shouting => "Calm down, I know what I'm doing!",
        _ if shouting => "Whoa, chill out!",
        _ if questioning => "Sure.",
        _ => "Whatever.",
    }
}
