use std::collections::{BTreeMap, BTreeSet};

pub struct School {
    grades_names: BTreeMap<u32, BTreeSet<String>>,
}

impl School {
    pub fn new() -> School {
        School {
            grades_names: BTreeMap::new(),
        }
    }

    pub fn add(&mut self, grade: u32, student: &str) {
        let student = student.to_string();

        if self
            .grades_names
            .values()
            .any(|students| students.contains(&student))
        {
            return;
        }

        self.grades_names.entry(grade).or_default().insert(student);
    }

    pub fn grades(&self) -> Vec<u32> {
        self.grades_names.keys().cloned().collect()
    }

    // If `grade` returned a reference, `School` would be forced to keep a `Vec<String>`
    // internally to lend out. By returning an owned vector of owned `String`s instead,
    // the internal structure can be completely arbitrary. The tradeoff is that some data
    // must be copied each time `grade` is called.
    pub fn grade(&self, grade: u32) -> Vec<String> {
        self.grades_names
            .get(&grade)
            .map(|students| students.iter().cloned().collect())
            .unwrap_or_default()
    }
}
