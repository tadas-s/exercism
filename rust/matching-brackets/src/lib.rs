pub fn brackets_are_balanced(string: &str) -> bool {
    let mut stack: Vec<char> = vec![];

    for c in string.chars() {
        match c {
            '[' | '{' | '(' => stack.push(c),
            ']' | '}' | ')' => {
                if stack.pop() != Some(opening_bracket(c)) {
                    return false;
                }
            }
            _ => (),
        }
    }

    stack.is_empty()
}

fn opening_bracket(c: char) -> char {
    match c {
        ']' => '[',
        '}' => '{',
        ')' => '(',
        _ => panic!("No opening bracket for char '{c}'"),
    }
}
