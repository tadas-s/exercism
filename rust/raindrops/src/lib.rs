pub fn raindrops(n: u32) -> String {
    let pling = n % 3 == 0;
    let plang = n % 5 == 0;
    let plong = n % 7 == 0;

    if pling || plang || plong {
        let mut rainspeak = "".to_string();
        if pling { rainspeak.push_str("Pling") }
        if plang { rainspeak.push_str("Plang") }
        if plong { rainspeak.push_str("Plong") }

        rainspeak
    } else {
        n.to_string()
    }
}
