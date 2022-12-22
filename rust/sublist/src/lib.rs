#[derive(Debug, PartialEq, Eq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(first_list: &[T], second_list: &[T]) -> Comparison {
    let first_sublist_of_second =
        first_list.is_empty() ||
        second_list.windows(first_list.len()).any(|w| first_list == w);

    let second_sublist_of_first =
        second_list.is_empty() ||
        first_list.windows(second_list.len()).any(|w| second_list == w);

    match (first_sublist_of_second, second_sublist_of_first) {
        (true, true) => Comparison::Equal,
        (true, false) => Comparison::Sublist,
        (false, true) => Comparison::Superlist,
        (false, false) => Comparison::Unequal,
    }
}
