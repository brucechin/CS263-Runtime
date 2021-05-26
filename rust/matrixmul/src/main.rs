type Matrix = Vec<Vec<i32>>;
use nalgebra::{ArrayStorage, Dynamic, Matrix as NMatrix, VecStorage};
use nalgebra::{U100, U20};
use ndarray::Array;
use tch::{Device, Kind, Tensor};

fn get_matrices(n: usize) -> (Matrix, Matrix) {
    let a: Matrix = vec![vec![1i32; n]; n];
    let b: Matrix = vec![vec![1i32; n]; n];
    (a, b)
}
fn matmul_torch(n: i64) {
    let a = Tensor::ones(&[n, n], (Kind::Float, Device::Cpu)) * 0.1;
    let b = Tensor::ones(&[n, n], (Kind::Float, Device::Cpu)) * 0.1;
    let o = a.matmul(&b);
}


fn matmul_ndarray(n: usize) {
    let a = Array::<f32, _>::from_elem((n, n), 0.1);
    let b = Array::<f32, _>::from_elem((n, n), 0.1);
    let out = a.dot(&b);
}

fn matmul(n: usize) {
    let (a, b) = get_matrices(n);
    let mut out: Matrix = vec![vec![0; n]; n];
    let time = std::time::Instant::now();
    for i in 0..n {
        for j in 0..n {
            let mut tmp = 0i32;
            for k in 0..n {
                tmp += a[i][k] * b[k][j];
            }
            out[i][j] = tmp;
        }
    }
    let took = time.elapsed();
    println!("{} -- length {} -- {} seconds", "Naive", n, took.as_secs_f64());
}


fn base_matmul() {
    let retries = 2;
    for n in [ 500, 1000, 2000, 4000, 8000].iter() {
        for j in 0..retries{
            matmul(*n as usize);

    
            // let time = std::time::Instant::now();
            // matmul_torch(*n as i64);
            // let took = time.elapsed();
            // println!("{} -- length {} -- {} seconds", "Torch", n, took.as_secs_f64());
        }

    }

}


fn main() {
    base_matmul();
}
