#include <iostream>
#include <map>
#include <queue>
#include <limits>
#include <chrono>

#include "heap.hpp"

#define MAXV 160000

typedef int64_t EdgeVal;

class Graph {
    public:

    using EdgeSet = std::map<int, EdgeVal> ;
    using KV = std::pair<int, EdgeVal> ;

    std::vector<EdgeSet> edges;
    int nv;
    bool directedGraph;

    Graph(int nv_, bool directed_=false) 
    : nv(nv_), directedGraph(directed_) 
    {
        edges = std::vector<EdgeSet>(nv_);
    } 

    void _add(int src, int dst, EdgeVal weight) {
        edges[src].insert( KV(dst, weight) );
    }

    void insert_edge(int src, int dst, EdgeVal weight, bool directed);
};

void Graph::insert_edge(int src, int dst, EdgeVal weight=0, bool directed=false)  {
    this->_add(src, dst, weight);
    if (!this->directedGraph || !directed ) {
        this->_add(dst, src, weight);
    }
}

int dijkstra_shortest_path(
    Graph *g,
    int parent[],
    EdgeVal distance[],
    int start
) {
    int nv = g->nv;
    // initialize
    for (int i = 0; i < nv; i++ ) {
        parent[i] = -1;
        distance[i] = -1;
    }
    std::vector<bool> discovered(nv+1, false);

    distance[start] = 0;

    // typedef std::tuple<EdgeVal, int> Q;
    // std::priority_queue<Q> pq;
    // pq.push( Q(0, start) );

    PriorQueue *pq = emptyPriorQueue(nv, distance);
    pq->Push(start);


    int v; EdgeVal w;

    // algorithm
 
    // while (!pq.empty()) {
    //     // take queue header
    //     std::tie(w, v) = pq.top();
    //     pq.pop();

    int op_count = 0;

    while(pq->cap > 0) {
        v = pq->Pop();
        
        if (discovered[v]) continue;
        discovered[v] = true;

        // traverse its edges
        Graph::EdgeSet &edgelist = g->edges[v];
        EdgeVal mydist = distance[v];

        for (auto it = edgelist.begin(); it != edgelist.end(); it++ ) {

            op_count++;


            int u = it->first;
            EdgeVal w = it->second;

            EdgeVal dist_new = mydist + w;
            if (distance[u] == -1) {
                distance[u]=dist_new;
                parent[u]=v;
                pq->Push(u);
            }
            else if (distance[u] > dist_new) {
                distance[u]=dist_new;
                parent[u]=v;
                pq->Fix(u);
            }
            // if (dist_new < distance[u])  // find a shorter path
            // {
            //     distance[u] = dist_new;
            //     parent[u] = v;    
                
            //     // we do not clear old records of u, because the priority queue will put it in the front, and the discovered will be marked. 
            //     // When the edge distance are all 1, their will be no path update, only newly discovered paths, so this is no difference from an algorithm that removes old values in priority queue.

            //     pq.push( Q(dist_new, u) );
            // }
        }
    }
    delete[] pq;

    // printf("Op count: %d\n", op_count);

    return op_count;
}

int main(int argc, char **argv) {

    if (argc>1) {
        int i = atoi(argv[1]);
        Graph *g = new Graph(i, true);
        int *parent = new int [MAXV + 1];
        EdgeVal *distance = new EdgeVal[MAXV + 1];
        int start = 1;
        for(int k = 0; k < i; k++){
            // for(int j = 0; j < i/100; j++){
            for(int j = 0; j < 20; j++){
                // set all edges to 1, the algorithm turns similar to BFS. 
                // This ensures this c++ priority queue implementation is same complexity with go's priority queue 
                // (c++ does not implement outdated element remove) 
                g->insert_edge(k, rand() % i, 0, false); 

            }
        }

        auto startTime = std::chrono::system_clock::now();  

        int traverse_edges =          dijkstra_shortest_path(g, parent, distance, start);;
        
        for (int i =0; i < 50+5; i++) {
            if (i==5) // start timing
                startTime = std::chrono::system_clock::now();    

        dijkstra_shortest_path(g, parent, distance, start);
        }


        auto endTime = std::chrono::system_clock::now();                                             
        std::chrono::duration<double> elapsedSeconds = endTime - startTime;                                
        std::cout << ">>> Dijkstra " << i <<  " nodes traversing " << traverse_edges << " Edges completed in " << elapsedSeconds.count() /50 << " seconds.\n"; 
        std::cout << std::endl;    
        delete g;
        delete[] parent;
        delete[] distance;

        return 0;
    }

    for(int i =2500; i <=160000; i*= 4){

        Graph *g = new Graph(i, true);
        int *parent = new int [MAXV + 1];
        EdgeVal *distance = new EdgeVal[MAXV + 1];
        int start = 1;
        for(int k = 0; k < i; k++){
            // for(int j = 0; j < i/100; j++){
            for(int j = 0; j < 20; j++){
                // set all edges to 1, the algorithm turns similar to BFS. 
                // This ensures this c++ priority queue implementation is same complexity with go's priority queue 
                // (c++ does not implement outdated element remove) 
                g->insert_edge(k, rand() % i, 0, true); 

            }
        }

        auto startTime = std::chrono::system_clock::now();  

        int traverse_edges =          dijkstra_shortest_path(g, parent, distance, start);;
        
        for (int i =0; i < 50+5; i++) {
            if (i==5) // start timing
                startTime = std::chrono::system_clock::now();    

        dijkstra_shortest_path(g, parent, distance, start);
        }


        auto endTime = std::chrono::system_clock::now();                                             
        std::chrono::duration<double> elapsedSeconds = endTime - startTime;                                
        std::cout << ">>> Dijkstra " << i <<  " nodes traversing " << traverse_edges << " Edges completed in " << elapsedSeconds.count() /50 << " seconds.\n"; 
        std::cout << std::endl;    
        delete g;
        delete[] parent;
        delete[] distance;
    }
    
    return 0;
}