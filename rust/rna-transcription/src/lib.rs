#[derive(PartialEq, Debug)]
pub struct DeoxyribonucleicAcid {
    pub strand: String,
}

impl DeoxyribonucleicAcid {
    pub fn new(strand: &str) -> DeoxyribonucleicAcid {
        DeoxyribonucleicAcid {
            strand: strand.to_string()
        }
    }

    pub fn to_rna(&self) -> RibonucleicAcid {
        RibonucleicAcid::new(
            self.strand.chars().map(|nucleotide| {
                match nucleotide {
                    'G' => 'C',
                    'C' => 'G',
                    'T' => 'A',
                    'A' => 'U',
                    _ => '?'
                }
            }).collect::<String>().as_str()
        )
    }
}

#[derive(PartialEq, Debug)]
pub struct RibonucleicAcid {
    pub strand: String,
}

impl RibonucleicAcid {
    pub fn new(strand: &str) -> RibonucleicAcid {
        RibonucleicAcid {
            strand: strand.to_string()
        }
    }
}
