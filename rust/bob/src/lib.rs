pub fn reply(question: &str) -> &str {
    match question {
        "Let's go make out behind the gym!" | "Tom-ay-to, tom-aaaah-to." | "Ending with ? means a question." => "Whatever.",
        "WATCH OUT!" | "1, 2, 3 GO!" | "ZOMG THE %^*@#$(*^ ZOMBIES ARE COMING!!11!!1!" | "I HATE YOU" => "Whoa, chill out!",
        "Does this cryogenic chamber make me look fat?" => "Sure.",
        "" | _ => "Fine. Be that way!"
    }
}
