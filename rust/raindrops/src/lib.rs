pub fn raindrops(number: u64) -> String {
    let mut drops : Vec<String> = Vec::new();

    if number % 3 == 0 {
        drops.push("Pling".to_string());
    }

    if number % 5 == 0 {
        drops.push("Plang".to_string());
    }

    if number % 7 == 0 {
        drops.push("Plong".to_string());
    }

    if drops.len() > 0 {
        drops.join("")
    } else {
        number.to_string()
    }
}
