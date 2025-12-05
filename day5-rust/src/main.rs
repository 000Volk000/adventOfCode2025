mod day5;
mod day5_2;
use std::io::{Write, stdin, stdout};

fn main() -> Result<(), String> {
    let mut opt = String::new();
    print!("Do you want to execute ex 1 or 2? [1,2]: ");
    stdout().flush().unwrap();

    stdin().read_line(&mut opt).expect("Failed to read line");

    match opt.trim() {
        "1" => {
            day5::ex1();
        }
        "2" => {
            day5_2::ex2();
        }
        _ => {
            return Err("Please pick 1 or 2".to_string());
        }
    }

    Ok(())
}
