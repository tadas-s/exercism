use std::collections::BTreeMap;

#[derive(PartialEq)]
enum GameOutcome {
    Win,
    Draw,
    Loss
}

#[derive(Clone, Copy)]
struct Team<'a> {
    name: &'a str,
    played: u8,
    won: u8,
    drew: u8,
    lost: u8
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
        if let Some((team_a, outcome_a, team_b, outcome_b)) = parse_line(line) {
            log_team_result(&mut totals, team_a, outcome_a);
            log_team_result(&mut totals, team_b, outcome_b);
        }
    });

    let mut table: Vec<String> = vec!["Team                           | MP |  W |  D |  L |  P".to_string()];

    let mut totals_vec: Vec<Team> = totals.values().cloned().collect();
    totals_vec.sort_by(|ta, tb| tb.points().cmp(&ta.points()));

    let mut formatted_totals = totals_vec.iter()
        .map(|team| {
            format!(
                "{:30} | {:2} | {:2} | {:2} | {:2} | {:2}",
                team.name, team.played, team.won, team.drew, team.lost, team.points()
            ).to_string()
        })
        .collect();

    table.append(&mut formatted_totals);

    table.join("\n")
}

fn parse_line(l: &str) -> Option<(&str, GameOutcome, &str, GameOutcome)> {
    match &l.split(";").collect::<Vec<&str>>()[..] {
        &[team_a, team_b, "win"] => Some((team_a, GameOutcome::Win, team_b, GameOutcome::Loss)),
        &[team_a, team_b, "loss"] => Some((team_a, GameOutcome::Loss, team_b, GameOutcome::Win)),
        &[team_a, team_b, "draw"] => Some((team_a, GameOutcome::Draw, team_b, GameOutcome::Draw)),
        _ => None
    }
}

fn log_team_result<'a>(totals: &mut BTreeMap<&'a str, Team<'a>>, team: &'a str, game_outcome: GameOutcome) {
    let entry = totals.entry(team).or_insert(
        Team { name: team, played: 0, won: 0, drew: 0, lost: 0 }
    );

    entry.played += 1;

    match game_outcome {
        GameOutcome::Win => entry.won += 1,
        GameOutcome::Loss => entry.lost += 1,
        GameOutcome::Draw => entry.drew += 1
    };
}