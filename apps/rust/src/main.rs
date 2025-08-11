use addsix::add_six;
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() != 2 {
        eprintln!("Usage: addsix <number>");
        std::process::exit(1);
    }

    let input = match args[1].parse::<f64>() {
        Ok(num) => num,
        Err(_) => {
            eprintln!("Error: '{}' is not a valid number", args[1]);
            std::process::exit(1);
        }
    };

    let result = add_six(input);
    println!("{}", result);
}
