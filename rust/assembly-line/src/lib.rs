pub fn production_rate_per_hour(speed: u8) -> f64 {
    let rate = (speed as f64) * 221.0;

    match speed {
        1..=4 => rate,
        5..=8 => rate * 0.9,
        9..=10 => rate * 0.77,
        _ => 0.0 // presumably higher speeds make line fall apart ;)
    }
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    (production_rate_per_hour(speed) / 60.0).floor() as u32
}
