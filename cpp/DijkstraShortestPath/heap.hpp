// a c++ implementation of cost heap
// similar to CS263-Runtime/go/DijkstraShortestPath/heap.go

// API:
// EmptyPriorQueue(int length, int64_t* cost) 
// pq.Push(int v)
// pq.Pop()
// pq.Fix(int v)

#include <vector>
#include <assert.h>

struct PriorQueue {
    using Cost = int64_t;
    int length;
    int cap=0;
    std::vector<int> heap;
    int* index=NULL;
    Cost* cost=NULL;
    
    ~PriorQueue() {
        if (index!=NULL) delete[] index;
        // if (cost!=NULL) delete[] cost;
    }

    bool less(int i, int j) { 
        return cost[heap[i]] < cost[heap[j]]; }

    void swap(int i, int j) {
        int t1 = heap[i]; int t2 = heap[j]; 
        heap[j] = t1; heap[i] = t2;
        index[t2] = i; index[t1] = j; }

    void up(int n) {
        while(true) {
            int i = (n-1)/2;
            if (i==n || !less(n, i)) break;
            swap(i, n);
            n = i;
        }
    }

    bool down(int from, int len) {
        int i = from;
        while(true) {
            int left = 2*i+1;
            if (left>=len || left<0) break;
            int j = left;
            int right = left + 1;
            if (right<len && less(right, left)) {
                j = right;
            }
            if (!less(j, i)) break;
            swap(i, j);
            i = j;
        }
        return (i>from);
    }

public:
    void Push(int v);
    int Pop();
    void Fix(int v);

};

PriorQueue* emptyPriorQueue(int length, PriorQueue::Cost cost[]) {
    PriorQueue *p = new PriorQueue[1];
    p->length = length;
    p->cost = cost;
    p->index = new int[length];
    return p;
}

void PriorQueue::Push(int v) {
    int n = this->cap;
    this->heap.push_back(v);
    this->cap = n+1;
    this->index[v] = n;
    this->up(n);
}

int PriorQueue::Pop() {
    int n = this->cap - 1;
    if (n<0) assert(0);
    this->swap(0, n);
    this->down(0, n);

    int v = this->heap[n];
    this->index[v] = -1;
    this->heap.pop_back();
    this->cap = n;
    return v;
}

void PriorQueue::Fix(int v) {
    int i = this->index[v];
    if (!this->down(i, this->cap)) {
        this->up(i);
    }
}
