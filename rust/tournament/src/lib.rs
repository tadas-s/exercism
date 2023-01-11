use std::collections::BTreeMap;

#[derive(PartialEq)]
enum GameOutcome {
    Win,
    Draw,
    Loss
}

#[derive(Clone, Copy)]
struct TeamTally {
    played: u8,
    won: u8,
    drew: u8,
    lost: u8,
    points: u8
}

pub fn tally(match_results: &str) -> String {
    let mut totals: BTreeMap<String, TeamTally> = BTreeMap::new();

    match_results.split("\n").for_each(|line| {
        if let Some((team_a, outcome_a, team_b, outcome_b)) = parse_line(line) {
            log_team_result(&mut totals, team_a, outcome_a);
            log_team_result(&mut totals, team_b, outcome_b);
        }
    });

    let mut table: Vec<String> = vec!["Team                           | MP |  W |  D |  L |  P".to_string()];

    let mut totals_vec: Vec<(String, TeamTally)> = totals.into_iter().collect::<Vec<(String, TeamTally)>>();
    totals_vec.sort_by(|(_, tat), (_, tbt)| tbt.points.cmp(&tat.points));

    let mut formatted_totals = totals_vec.iter()
        .map(|(team, tally)| {
            format!(
                "{:30} | {:2} | {:2} | {:2} | {:2} | {:2}",
                team, tally.played, tally.won, tally.drew, tally.lost, tally.points
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

fn log_team_result(totals: &mut BTreeMap<String, TeamTally>, team: &str, game_outcome: GameOutcome) {
    match game_outcome {
        GameOutcome::Win => {
            totals.entry(team.to_string())
                  .and_modify(|t| { t.played += 1; t.won += 1; t.points += 3 })
                  .or_insert(TeamTally { played: 1, won: 1, drew: 0, lost: 0, points: 3 });
        },
        GameOutcome::Loss => {
            totals.entry(team.to_string())
                  .and_modify(|t| { t.played += 1; t.lost += 1 })
                  .or_insert(TeamTally { played: 1, won: 0, drew: 0, lost: 1, points: 0 });
        },
        GameOutcome::Draw => {
            totals.entry(team.to_string())
                  .and_modify(|t| { t.played += 1; t.drew += 1; t.points += 1 })
                  .or_insert(TeamTally { played: 1, won: 0, drew: 1, lost: 0, points: 1 });
        }
    };
}