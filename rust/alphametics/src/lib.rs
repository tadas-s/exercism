use itertools::Itertools;
use std::collections::{HashMap, HashSet};
use std::iter::zip;

struct Num {
    digits: Vec<char>,
}

impl Num {
    fn new(number: &str) -> Num {
        Num {
            digits: number.trim().chars().collect(),
        }
    }

    fn to_u64(&self, mapping: &HashMap<&char, u8>) -> u64 {
        self.digits
            .iter()
            .rev()
            .enumerate()
            .map(|(i, &d)| (mapping[&d] as u64) * 10u64.pow(i as u32) )
            .sum()
    }
}

fn get_unique_digits(input: &str) -> Vec<char> {
    input
        .chars()
        .filter(|&c| c.is_alphabetic())
        .unique()
        .collect()
}

fn get_non_zero_digits(input: &str) -> HashSet<char> {
    let mut parts = input.split(" == ");
    let mut members = parts
        .next()
        .unwrap()
        .split(" + ")
        .map(|m| m.trim())
        .collect::<Vec<_>>();
    members.push(parts.next().unwrap().trim());
    members.iter().map(|m| m.chars().next().unwrap()).collect()
}

fn parse_input_into_nums(input: &str) -> (Vec<Num>, Num) {
    let mut parts = input.split(" == ");
    let sum_members = parts
        .next()
        .unwrap()
        .split(" + ")
        .map(Num::new)
        .collect::<Vec<_>>();
    let final_sum = Num::new(parts.next().unwrap());

    (sum_members, final_sum)
}

pub fn solve(input: &str) -> Option<HashMap<char, u8>> {
    let unique_digits = get_unique_digits(input);
    let non_zero_digits = get_non_zero_digits(input);
    let (sum_members, final_sum) = parse_input_into_nums(input);

    for v in (0..=9u8).rev().permutations(unique_digits.len()) {
        let mapping: HashMap<&char, u8> = zip(&unique_digits, v).collect();

        let variant_sum = sum_members.iter().map(|m| m.to_u64(&mapping)).sum::<u64>();
        let variant_final_sum = final_sum.to_u64(&mapping);

        if variant_sum == variant_final_sum {
            // If any of leading digits is zero - continue searching
            if non_zero_digits.iter().any(|d| mapping[d] == 0) {
                continue;
            }

            return Some(mapping.iter().map(|(&&character, &b)| (character, b)).collect());
        }
    }

    None
}
