pub fn build_proverb(list: &[&str]) -> String {
    if list.len() > 0 {
        let mut proverb: String = list
            .windows(2)
            .map(|pieces| format!("For want of a {} the {} was lost.\n", pieces[0], pieces[1]))
            .collect::<Vec<_>>()
            .join("");

        proverb.push_str(&format!("And all for the want of a {}.", list[0]));

        proverb
    } else {
        String::from("")
    }
}
