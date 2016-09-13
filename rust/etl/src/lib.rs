use std::collections::BTreeMap;

pub type LegacyFormat = BTreeMap<i32, Vec<String>>;
pub type NewFormat = BTreeMap<String, i32>;

pub fn transform(input: &LegacyFormat) -> NewFormat {
    let mut transformed: NewFormat = NewFormat::new();

    for (points, letters) in input {
        for letter in letters {
            transformed.insert(letter.to_lowercase().to_string(), *points);
        }
    }

    return transformed;
}
