use CalculatorInput::*;

pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    let mut stack: Vec<i32> = Vec::new();

    for input in inputs {
        match input {
            Value(v) => stack.push(*v),
            operation @ (Add | Subtract | Multiply | Divide) => {
                if let (Some(op2), Some(op1)) = (stack.pop(), stack.pop()) {
                    match operation {
                        Add => stack.push(op1 + op2),
                        Subtract => stack.push(op1 - op2),
                        Multiply => stack.push(op1 * op2),
                        Divide => stack.push(op1 / op2),
                        _ => unreachable!()
                    }
                } else {
                    return None;
                }
            }
        }
    }

    if stack.len() == 1 {
        stack.pop()
    } else {
        None
    }
}
