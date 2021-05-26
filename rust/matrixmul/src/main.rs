fn matmul(n: usize) {
    let (a, b) = get_matrices(n);
    let mut out: Matrix = vec![vec![0; n]; n];
    let time = std::time::Instant::now();
    let retries = 5;
    for i in 0..retries {
        for i in 0..n {
            for j in 0..n {
                let mut tmp = 0i32;
                for k in 0..n {
                    tmp += a[i][k] * b[k][j];
                }
                out[i][j] = tmp;
            }
        }
    }
    let took = time.elapsed();
    println!("{} -- length {} -- {} seconds", "Naive", n, took.as_secs_f64() / retries);
}


fn base_matmul() {
    
    for n in [ 500, 1000, 2000, 4000, 8000].iter() {
            matmul(*n as usize);

    
            // let time = std::time::Instant::now();
            // matmul_torch(*n as i64);
            // let took = time.elapsed();
            // println!("{} -- length {} -- {} seconds", "Torch", n, took.as_secs_f64());
        

    }

}
