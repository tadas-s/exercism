use std::collections::BTreeMap;

pub struct School {
    grades_names: BTreeMap<u32, Vec<String>>,
}

impl School {
    pub fn new() -> School {
        School { grades_names: BTreeMap::new() }
    }

    pub fn add(&mut self, grade: u32, student: &str) {
        self.grades_names
            .entry(grade)
            .and_modify(|list| {
                list.push(student.to_string());
                list.sort();
            })
            .or_insert(vec![student.to_string()]);
    }

    pub fn grades(&self) -> Vec<u32> {
        self.grades_names.keys().cloned().collect()
    }

    // If `grade` returned a reference, `School` would be forced to keep a `Vec<String>`
    // internally to lend out. By returning an owned vector of owned `String`s instead,
    // the internal structure can be completely arbitrary. The tradeoff is that some data
    // must be copied each time `grade` is called.
    pub fn grade(&self, grade: u32) -> Vec<String> {
        match self.grades_names.get(&grade) {
            Some(list) => list.clone(),
            None => vec![],
        }
    }
}