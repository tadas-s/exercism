#[derive(Debug)]
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

#[derive(Debug)]
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

impl PartialEq for RibonucleicAcid {
    fn eq(&self, other: &RibonucleicAcid) -> bool {
        self.strand == other.strand
    }
}
