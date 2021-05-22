//Purpose: Implementation of Dijkstra's algorithm which finds the shortest
//path from a start node to every other node in a weighted graph. 
//Time complexity: O(n^2)
#include <iostream>
#include <limits>
#include <chrono>
using namespace std;
#define MAXV 160000

class EdgeNode{
    public:
        int key;
        int weight;
        EdgeNode *next;
        EdgeNode(int, int);
};

EdgeNode::EdgeNode(int key, int weight){
    this->key = key;
    this->weight = weight;
    this->next = NULL;
}

class Graph{
    bool directed;
    public:
        EdgeNode *edges[MAXV + 1];
        Graph(bool);
        ~Graph();
        void insert_edge(int, int, int, bool);
        void print();
};

Graph::Graph(bool directed){
    this->directed = directed;
    for(int i = 1; i < (MAXV + 1); i ++){
        this->edges[i] = NULL;
    }
}

Graph::~Graph(){
    //to do
}

void Graph::insert_edge(int x, int y, int weight, bool directed){
    if(x > 0 && x < (MAXV + 1) && y > 0 && y < (MAXV + 1)){
        EdgeNode *edge = new EdgeNode(y, weight);
        edge->next = this->edges[x];
        this->edges[x] = edge;
        if(!directed){
            insert_edge(y, x, weight, true);
        }
    }
}

void Graph::print(){
    for(int v = 1; v < (MAXV + 1); v ++){
        if(this->edges[v] != NULL){
            cout << "Vertex " << v << " has neighbors: " << endl;
            EdgeNode *curr = this->edges[v];
            while(curr != NULL){
                cout << curr->key << endl;
                curr = curr->next;
            }
        }
    }
}

void init_vars(bool discovered[], int distance[], int parent[]){
    for(int i = 1; i < (MAXV + 1); i ++){
        discovered[i] = false;
        distance[i] = std::numeric_limits<int>::max();
        parent[i] = -1;
    }
}

void dijkstra_shortest_path(Graph *g, int parent[], int distance[], int start){

    bool discovered[MAXV + 1];
    EdgeNode *curr;
    int v_curr;
    int v_neighbor;
    int weight;
    int smallest_dist;

    init_vars(discovered, distance, parent);

    distance[start] = 0;
    v_curr = start;

    while(discovered[v_curr] == false){

        discovered[v_curr] = true;
        curr = g->edges[v_curr];

        while(curr != NULL){

            v_neighbor = curr->key;
            weight = curr->weight;

            if((distance[v_curr] + weight) < distance[v_neighbor]){
                distance[v_neighbor] = distance[v_curr] + weight;
                parent[v_neighbor] = v_curr;
            }
            curr = curr->next;
        }

        //set the next current vertex to the vertex with the smallest distance
        smallest_dist = std::numeric_limits<int>::max();
        for(int i = 1; i < (MAXV + 1); i ++){
            if(!discovered[i] && (distance[i] < smallest_dist)){
                v_curr = i;
                smallest_dist = distance[i];
            }
        }
    }
}

void print_shortest_path(int v, int parent[]){
    if(v > 0 && v < (MAXV + 1) && parent[v] != -1){
        cout << parent[v] << " ";
        print_shortest_path(parent[v], parent);
    }
}

void print_distances(int start, int distance[]){
    for(int i = 1; i < (MAXV + 1); i ++){
        if(distance[i] != std::numeric_limits<int>::max()){
            cout << "Shortest distance from " << start << " to " << i << " is: " << distance[i] << endl;
        }
    }
}

int main(){
    for(int i =10000; i <=160000; i*= 4){

        Graph *g = new Graph(false);
        int parent[MAXV + 1];
        int distance[MAXV + 1];
        int start = 1;
        for(int k = 0; k < i; k++){
            for(int j = 0; j < i/100; j++){
                g->insert_edge(k, rand() % i, rand() % 100, false);

            }
        }

        auto startTime = std::chrono::system_clock::now();                                           
        dijkstra_shortest_path(g, parent, distance, start);


        auto endTime = std::chrono::system_clock::now();                                             
        std::chrono::duration<double> elapsedSeconds = endTime - startTime;                                
        std::cout << ">>> Dijkstra " << i <<  " nodes completed in " << elapsedSeconds.count() << " seconds.\n"; 
        std::cout << std::endl;    
        delete g;
    }
    
    return 0;
}
