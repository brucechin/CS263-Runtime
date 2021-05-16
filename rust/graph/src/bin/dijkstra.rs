use petgraph::prelude::*;
use std::cmp::{max, min};
use std::time::Instant;
use rand::Rng;

use petgraph::algo::dijkstra;
fn main(){
    let runTimes = 10; //obtain the average time
    let mut g = Graph::new_undirected();
    let mut rng = rand::thread_rng();

    for j in [10000, 100000, 1000000].iter() {
        let nodes: Vec<NodeIndex<_>> = (0..*j).into_iter().map(|i| g.add_node(i)).collect();
        let neighbor_bound = j / 100;

        for i in 0..*j {
            let n1 = nodes[i];
            let neighbour_count = i % neighbor_bound;
            let j_from = max(0, i as i32 - neighbour_count as i32 / 2) as usize;
            let j_to = min(*j, j_from + neighbour_count);
            for j in j_from..j_to {
                let n2 = nodes[j];
                let distance = (i + 3) % 10;
                g.add_edge(n1, n2, distance);
            }
        }
        let begin = Instant::now();

        for k in 0..runTimes{
            let _scores = dijkstra(&g, nodes[rng.gen::<usize>() % nodes.len()], Some(nodes[rng.gen::<usize>() % nodes.len()]), |e| *e.weight());
        }
        let end = Instant::now();
        println!("dijkstra on {} nodes time {:?}", j,  end.duration_since(begin));
    }

}