use std::collections::BTreeMap;

enum Result {
    Win,
    Draw,
    Loss,
}

struct Team<'a> {
    name: &'a str,
    played: u8,
    won: u8,
    drew: u8,
    lost: u8,
}

impl PartialEq for Team<'_> {
    fn eq(&self, other: &Self) -> bool {
        self.name == other.name
    }
}

impl Team<'_> {
    pub fn points(&self) -> u8 {
        self.won * 3 + self.drew
    }
}

pub fn tally(match_results: &str) -> String {
    let mut totals: BTreeMap<&str, Team> = BTreeMap::new();

    match_results.split("\n").for_each(|line| {
        if let Some((team_a, result_a, team_b, result_b)) = parse_line(line) {
            log_team_result(&mut totals, team_a, result_a);
            log_team_result(&mut totals, team_b, result_b);
        }
    });

    let mut table: Vec<String> =
        vec!["Team                           | MP |  W |  D |  L |  P".to_string()];

    let mut totals_vec: Vec<Team> = totals.into_values().collect();
    totals_vec.sort_by(|ta, tb| tb.points().cmp(&ta.points()));

    let mut formatted_totals = totals_vec
        .iter()
        .map(|team| {
            format!(
                "{:30} | {:2} | {:2} | {:2} | {:2} | {:2}",
                team.name,
                team.played,
                team.won,
                team.drew,
                team.lost,
                team.points()
            )
            .to_string()
        })
        .collect();

    table.append(&mut formatted_totals);

    table.join("\n")
}

fn parse_line(l: &str) -> Option<(&str, Result, &str, Result)> {
    match &l.split(";").collect::<Vec<&str>>()[..] {
        &[team_a, team_b, "win"] =>
            Some((team_a, Result::Win, team_b, Result::Loss)),
        &[team_a, team_b, "loss"] =>
            Some((team_a, Result::Loss, team_b, Result::Win)),
        &[team_a, team_b, "draw"] =>
            Some((team_a, Result::Draw, team_b, Result::Draw)),
        _ => None,
    }
}

fn log_team_result<'a>(
    totals: &mut BTreeMap<&'a str, Team<'a>>,
    team: &'a str,
    game_result: Result,
) {
    let entry = totals.entry(team).or_insert(Team {
        name: team,
        played: 0,
        won: 0,
        drew: 0,
        lost: 0,
    });

    entry.played += 1;

    match game_result {
        Result::Win => entry.won += 1,
        Result::Loss => entry.lost += 1,
        Result::Draw => entry.drew += 1,
    };
}
