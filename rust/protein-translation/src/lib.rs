use std::collections::BTreeMap;

const STOP_CODON: &str = "stop codon";

pub struct CodonsInfo<'a> {
    map: BTreeMap<&'a str, &'a str>,
}

impl<'a> CodonsInfo<'a> {
    pub fn name_for(&self, codon: &str) -> Option<&'a str> {
        self.map.get(codon).cloned()
    }

    pub fn of_rna(&self, rna: &str) -> Option<Vec<&'a str>> {
        let mut names = vec![];

        for offset in (0..rna.len()).step_by(3) {
            let codon = &rna[offset..(offset + 3).min(rna.len())];

            match self.name_for(codon) {
                Some(STOP_CODON) => break,
                Some(name) => names.push(name),
                None => return None,
            }
        }

        Some(names)
    }
}

pub fn parse<'a>(pairs: Vec<(&'a str, &'a str)>) -> CodonsInfo<'a> {
    CodonsInfo {
        map: pairs.iter().cloned().collect(),
    }
}
