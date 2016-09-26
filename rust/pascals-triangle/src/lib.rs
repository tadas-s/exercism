pub struct PascalsTriangle {
    rows: Vec<Vec<u32>>,
}

impl PascalsTriangle {
    pub fn new(row_count: usize) -> Self {
        let mut rows: Vec<Vec<u32>> = Vec::with_capacity(row_count);

        if row_count > 0 {
            rows.push(vec![1]);

            for row in 1..row_count {
                let mut r: Vec<u32> = Vec::with_capacity(row + 1); 
                for i in 0..(row + 1) {
                    r.push(match i {
                        0 => rows[row - 1][i],
                        _ if i == row => rows[row - 1][i - 1],
                        _ => rows[row - 1][i - 1] + rows[row - 1][i],
                    });
                }
                rows.push(r);
            }
        }

        PascalsTriangle {
            rows: rows,
        }
    }

    pub fn rows(&self) -> Vec<Vec<u32>> {
        self.rows.clone()
    }
}
