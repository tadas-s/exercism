extern crate regex;

use regex::Regex;

pub fn abbreviate(name: &str) -> String {
    let re = Regex::new("([A-Z]([a-z]+|[A-Z]+))|([a-z]+)").unwrap();
    
    return re.captures_iter(&name)
        .map(|c| c.at(1).unwrap_or(c.at(3).unwrap_or("")))
        .map(|c| c.chars().next().unwrap())
        .collect::<String>()
        .to_uppercase();
}
