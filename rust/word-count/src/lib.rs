extern crate regex;

use std::collections::HashMap;
use regex::Regex;

pub fn word_count(string: &str) -> HashMap<String, u32> {
    let re = Regex::new(r"[^A-Za-z0-9]+").unwrap();
    let mut counts: HashMap<String, u32> = HashMap::new();
    
    for word in re.split(&string) {
        if !word.is_empty() {
            let counter = counts.entry(word.to_lowercase().to_string()).or_insert(0);
            *counter += 1;
        }
    }
    
    return counts;
}
