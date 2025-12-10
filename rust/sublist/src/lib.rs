#[derive(Debug, PartialEq, Eq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist(first_list: &[i32], second_list: &[i32]) -> Comparison {
    let first_list_sublist_of_second_list =
        first_list.is_empty() ||
        second_list.windows(first_list.len()).any(|w| w == first_list);

    let second_list_sublist_of_first_list =
        second_list.is_empty() ||
        first_list.windows(second_list.len()).any(|w| w == second_list);

    match (first_list_sublist_of_second_list, second_list_sublist_of_first_list) {
        (true, true) => Comparison::Equal,
        (true, false) => Comparison::Sublist,
        (false, true) => Comparison::Superlist,
        (false, false) => Comparison::Unequal
    }
}
