extern crate regex;

use regex::Regex;

pub fn abbreviate(name: &str) -> String {
    let re = Regex::new("([A-Z]([a-z]+|[A-Z]+))|([a-z]+)").unwrap();
    
    return re.captures_iter(&name)
        .map(|c| c.at(0).unwrap().chars().next().unwrap())
        .collect::<String>()
        .to_uppercase();
}
