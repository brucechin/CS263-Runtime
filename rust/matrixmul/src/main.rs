
type Matrix = Vec<Vec<f32>>;
use nalgebra::{ArrayStorage, Dynamic, Matrix as NMatrix, VecStorage};
use nalgebra::{U100, U20};
// use ndarray::Array;
use tch::{Device, Kind, Tensor};

fn get_matrices(n: usize) -> (Matrix, Matrix) {
    let a: Matrix = vec![vec![0.1; n]; n];
    let b: Matrix = vec![vec![0.1; n]; n];
    (a, b)
}
fn matmul_torch(n: i64) {
    let a = Tensor::ones(&[n, n], (Kind::Float, Device::Cpu)) * 0.1;
    let b = Tensor::ones(&[n, n], (Kind::Float, Device::Cpu)) * 0.1;
    let o = a.matmul(&b);
}


// fn matmul_ndarray(n: usize) {
//     let a = Array::<f32, _>::from_elem((n, n), 0.1);
//     let b = Array::<f32, _>::from_elem((n, n), 0.1);
//     let out = a.dot(&b);
// }

fn matmul(n: usize) {
    let (a, b) = get_matrices(n);
    let mut out: Matrix = vec![vec![0.0; n]; n];
    for i in 0..n {
        for j in 0..n {
            for k in 0..n {
                out[i][j] += a[i][k] * b[k][j];
            }
        }
    }
}


fn base_matmul() {
    let retries = 5;
    for n in [ 500, 1000, 2000, 4000, 8000].iter() {
        for j in 0..retries{
            // let time = std::time::Instant::now();
            // matmul_ndarray(*n as usize);
            // let took = time.elapsed();
            // println!("{} -- length {} -- {} seconds", "NDarray", n, took.as_secs_f64());
    
            let time = std::time::Instant::now();
            matmul_torch(*n as i64);
            let took = time.elapsed();
            println!("{} -- length {} -- {} seconds", "Torch", n, took.as_secs_f64());
        }

    }

}


fn main() {
    base_matmul();
}
