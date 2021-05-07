
use std::time::Instant;
use snap::raw::*;

fn main(){
    
    //TODO add large files to compress and decompress
    let begin = Instant::now();
    const CORPUS_ALICE29: &'static [u8] = include_bytes!("../data/alice29.txt");
    let mut output = vec![0; snap::raw::max_compress_len(CORPUS_ALICE29.len())];
    snap::raw::Encoder::new().compress(CORPUS_ALICE29, &mut output).unwrap();
    let end = Instant::now();
    println!("compress time {:?}", end.duration_since(begin));


    let compressed =  snap::raw::Encoder::new().compress_vec(CORPUS_ALICE29).unwrap();
    let begin = Instant::now();
    let mut dst = vec![0; CORPUS_ALICE29.len()];
    snap::raw::Decoder::new().decompress(&compressed, &mut dst);
    let end = Instant::now();
    println!("decompress time {:?}", end.duration_since(begin));


}