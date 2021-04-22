
mod quicksort;
mod heapsort;


pub use quicksort::QuickSort;
pub use heapsort::HeapSort;

pub trait Sorter<T>{
    fn sort(&self, slice: &mut [T])
    where
        T: Ord;
}

