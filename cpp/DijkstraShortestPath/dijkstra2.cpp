#include <iostream>
#include <map>
#include <queue>

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

void dijkstra_shortest_path(
    Graph *g,
    int parent[],
    EdgeVal distance[],
    int start
) {
    int nv = g->nv;
    // initialize
    for (int i = 0; i < nv; i++ ) {
        parent[i] = -1;
        distance[i] = std::numeric_limits<int64_t>::max();
    }
    std::vector<bool> discovered(nv+1, false);

    distance[start] = 0;

    typedef std::tuple<EdgeVal, int> Q;
    std::priority_queue<Q> pq;
    pq.push( Q(0, start) );

    int v; EdgeVal w;

    // algorithm
 
    while (!pq.empty()) {
        // take queue header
        std::tie(w, v) = pq.top();
        pq.pop();
        
        if (discovered[v]) continue;
        discovered[v] = true;

        // traverse its edges
        Graph::EdgeSet &edgelist = g->edges[v];
        EdgeVal mydist = distance[v];
        for (auto it = edgelist.begin(); it != edgelist.end(); it++ ) {
            int u = it->first;
            EdgeVal w = it->second;

            EdgeVal dist_new = mydist + w;
            if (dist_new < distance[u])  // find a shorter path
            {
                distance[u] = dist_new;
                parent[u] = v;    
                
                // we do not clear old records of u, because the priority queue will put it in the front, and the discovered will be marked. 
                // When the edge distance are all 1, their will be no path update, only newly discovered paths, so this is no difference from an algorithm that removes old values in priority queue.

                pq.push( Q(dist_new, u) );
            }
        }
    }
}

int main() {
    for(int i =10000; i <=160000; i*= 4){

        Graph *g = new Graph(i, false);
        int *parent = new int [MAXV + 1];
        EdgeVal *distance = new EdgeVal[MAXV + 1];
        int start = 1;
        for(int k = 0; k < i; k++){
            for(int j = 0; j < i/100; j++){
                // set all edges to 1, the algorithm turns similar to BFS. 
                // This ensures this c++ priority queue implementation is same complexity with go's priority queue 
                // (c++ does not implement outdated element remove) 
                g->insert_edge(k, rand() % i, 1, false); 

            }
        }

        auto startTime = std::chrono::system_clock::now();                                           
        dijkstra_shortest_path(g, parent, distance, start);


        auto endTime = std::chrono::system_clock::now();                                             
        std::chrono::duration<double> elapsedSeconds = endTime - startTime;                                
        std::cout << ">>> Dijkstra " << i <<  " nodes completed in " << elapsedSeconds.count() << " seconds.\n"; 
        std::cout << std::endl;    
        delete g;
        delete[] parent;
        delete[] distance;
    }
    
    return 0;
}