const STUDENTS: [&str; 12] = [
    "Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph",
    "Kincaid", "Larry",
];

const PLANTS: [(char, &str); 4] = [
    ('C', "clover"),
    ('G', "grass"),
    ('R', "radishes"),
    ('V', "violets"),
];

pub fn plants(diagram: &str, student: &str) -> Vec<&'static str> {
    let section = match STUDENTS.binary_search(&student) {
        Err(_) => return vec![],
        Ok(student_section) => student_section,
    };

    let diagram: Vec<Vec<char>> = diagram
        .split("\n")
        .map(|x| x.chars().collect::<Vec<char>>())
        .collect();

    [
        diagram[0][section * 2],
        diagram[0][section * 2 + 1],
        diagram[1][section * 2],
        diagram[1][section * 2 + 1],
    ]
    .iter()
    .map(|plant_code| PLANTS.binary_search_by(|plant| plant.0.cmp(plant_code)))
    .flat_map(|x| match x {
        Ok(i) => Some(PLANTS[i].1),
        Err(_) => None,
    })
    .collect()
}
