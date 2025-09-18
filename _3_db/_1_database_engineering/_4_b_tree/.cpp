// Non‑Clustered B+Tree Index on `city` with clustered-key pointers (id)
// ----------------------------------------------------------------------
// What this demonstrates
//  - A simple in‑memory base table (clustered on id via std::map for clarity)
//  - A B+Tree *non‑clustered index* on `city`, whose leaf entries store
//      (city, id) — i.e., the *row locator* is the clustered key
//  - Queries: WHERE city = ?
//      Step 1: probe non‑clustered B+Tree to get list of ids
//      Step 2: use ids to fetch full rows from the base table (clustered index)
//  - ASCII tree printout + Graphviz DOT export for visualization
//
// Build:   g++ -std=c++17 -O2 nci_bptree.cpp -o nci_bptree
// Run:     ./nci_bptree
// Visual:  dot -Tpng idx_city.dot -o idx_city.png   # requires Graphviz
// ----------------------------------------------------------------------
#include <bits/stdc++.h>
using namespace std;

// ------------------------ Base table (clustered on id) -----------------------
struct Row {
    int id; string name; string city;
};

class BaseTable {
    // Simulate clustered index on id using std::map
    map<int, Row> clustered; // id -> Row
public:
    void insert(Row r) {
        if (!clustered.emplace(r.id, r).second) throw runtime_error("duplicate id");
    }
    optional<Row> get(int id) const {
        auto it = clustered.find(id);
        if (it == clustered.end()) return nullopt;
        return it->second;
    }
    void print_all() const {
        cout << "BaseTable (clustered on id)\n";
        for (auto &p: clustered) {
            cout << "  (id=" << p.second.id << ", name=" << p.second.name
                 << ", city=" << p.second.city << ")\n";
        }
        cout << "\n";
    }
};

// ------------------------ Non‑Clustered B+Tree on city -----------------------
// Simpler, educational B+Tree with small order for compact output.
// Node max_keys = ORDER*2; min_keys = ORDER (except root)

struct NciNode {
    bool is_leaf;
    vector<string> keys;               // separator keys (city)
    vector<int> children;              // child node ids (for internal nodes)

    // Leaf payload: for each key[i], a vector of clustered keys (ids)
    vector<vector<int>> values;        // only used if is_leaf
    int next_leaf = -1;                // leaf-level linked list
};

class NciBPlusTree {
public:
    explicit NciBPlusTree(int order = 2): ORDER(order) {
        // create root as empty leaf
        NciNode root; root.is_leaf = true; nodes.push_back(root); // id=0
    }

    // Insert (city -> id)
    void insert(const string& city, int id) {
        int new_child = -1; string promote_key;
        bool grew = insert_rec(0, city, id, new_child, promote_key);
        if (grew) {
            // create a new root
            NciNode new_root; new_root.is_leaf = false;
            new_root.keys.push_back(promote_key);
            new_root.children.push_back(0);
            new_root.children.push_back(new_child);
            nodes.push_back(new_root); // at end
            int new_root_id = (int)nodes.size()-1;
            // move new_root to index 0, swap with old root
            swap(nodes[0], nodes[new_root_id]);
        }
    }

    // Search exact city -> list of ids
    vector<int> search(const string& city) const {
        int nid = 0;
        while (true) {
            auto const &n = nodes[nid];
            if (n.is_leaf) {
                int i = lower_bound(n.keys.begin(), n.keys.end(), city) - n.keys.begin();
                if (i < (int)n.keys.size() && n.keys[i] == city) return n.values[i];
                return {};
            } else {
                int i = lower_bound(n.keys.begin(), n.keys.end(), city) - n.keys.begin();
                nid = n.children[i];
            }
        }
    }

    // Pretty ASCII dump (levels) + leaf contents
    void dump_ascii() const {
        cout << "\nNon‑Clustered B+Tree on city (leaf stores city + id list)" << "\n";
        // BFS levels
        deque<int> q{0};
        while (!q.empty()) {
            size_t sz = q.size();
            while (sz--) {
                int nid = q.front(); q.pop_front();
                auto const &n = nodes[nid];
                cout << (n.is_leaf?"(L)":"(I)") << "#" << nid << "[";
                for (size_t i=0;i<n.keys.size();++i) {
                    cout << n.keys[i];
                    if (!n.is_leaf && i+1<n.keys.size()) cout << " | ";
                    if (n.is_leaf) {
                        cout << ":";
                        cout << "{";
                        for (size_t k=0;k<n.values[i].size();++k) {
                            cout << n.values[i][k];
                            if (k+1<n.values[i].size()) cout << ",";
                        }
                        cout << "}";
                        if (i+1<n.keys.size()) cout << "  ";
                    }
                }
                cout << "]  ";
                if (!n.is_leaf) for (int c: n.children) q.push_back(c);
            }
            cout << "\n";
        }
        // Leaf chain
        cout << "Leaf chain: ";
        int leaf = leftmost_leaf();
        while (leaf != -1) {
            auto const &n = nodes[leaf];
            cout << "#" << leaf << " -> ";
            leaf = n.next_leaf;
        }
        cout << "nil\n\n";
    }

    // Dump Graphviz DOT to a file
    void dump_dot(const string& filename) const {
        ofstream out(filename);
        out << "digraph NCI {\n  node [shape=record];\n";
        for (size_t i=0;i<nodes.size();++i) {
            auto const &n = nodes[i];
            out << "  n" << i << " [label=\"" << (n.is_leaf?"L":"I") << "#" << i << ": ";
            for (size_t k=0;k<n.keys.size();++k) {
                out << n.keys[k];
                if (n.is_leaf) {
                    out << " (";
                    for (size_t j=0;j<n.values[k].size();++j){ out << n.values[k][j]; if (j+1<n.values[k].size()) out << ","; }
                    out << ")";
                }
                if (k+1<n.keys.size()) out << " | ";
            }
            out << "\"];\n";
            if (!n.is_leaf) {
                for (size_t c=0;c<n.children.size();++c) {
                    out << "  n"<<i<<" -> n"<<n.children[c]<<";\n";
                }
            } else if (n.next_leaf!=-1) {
                out << "  n"<<i<<" -> n"<<n.next_leaf<<" [style=dashed,color=gray,label=\"next\"];\n";
            }
        }
        out << "}\n";
        out.close();
    }

private:
    int ORDER;                     // min degree
    vector<NciNode> nodes;         // node 0 is root (may be moved during splits)

    int leftmost_leaf() const {
        int nid = 0; while (!nodes[nid].is_leaf) nid = nodes[nid].children.front(); return nid; }

    // Insert recursively; returns true if root must grow
    bool insert_rec(int nid, const string& city, int id, int& new_child_out, string& promote_key) {
        NciNode &n = nodes[nid];
        if (n.is_leaf) {
            // insert into leaf
            int i = lower_bound(n.keys.begin(), n.keys.end(), city) - n.keys.begin();
            if (i < (int)n.keys.size() && n.keys[i] == city) {
                // append id (keep sorted, avoid dups)
                auto &vec = n.values[i];
                if (!binary_search(vec.begin(), vec.end(), id)) {
                    vec.insert(lower_bound(vec.begin(), vec.end(), id), id);
                }
            } else {
                n.keys.insert(n.keys.begin()+i, city);
                n.values.insert(n.values.begin()+i, vector<int>{id});
            }
            if ((int)n.keys.size() > 2*ORDER) {
                // split leaf
                int right_id = split_leaf(nid);
                new_child_out = right_id;
                promote_key = nodes[right_id].keys.front();
                return true; // caller must add separator to parent
            }
            return false;
        } else {
            // descend
            int i = lower_bound(n.keys.begin(), n.keys.end(), city) - n.keys.begin();
            bool grew = insert_rec(n.children[i], city, id, new_child_out, promote_key);
            if (!grew) return false;
            // insert separator key + new child
            n.keys.insert(n.keys.begin()+i, promote_key);
            n.children.insert(n.children.begin()+i+1, new_child_out);
            if ((int)n.keys.size() > 2*ORDER) {
                int right_id = split_internal(nid);
                new_child_out = right_id;
                promote_key = nodes[right_id].keys.front(); // after split_internal we will adjust
                // Note: split_internal returns with the *first* key of right already set as correct promote
                return true;
            }
            return false;
        }
    }

    int split_leaf(int nid) {
        NciNode right; right.is_leaf = true;
        NciNode &left = nodes[nid];
        int total = (int)left.keys.size();
        int mid = total/2; // move [mid..end) to right
        right.keys.assign(left.keys.begin()+mid, left.keys.end());
        right.values.assign(left.values.begin()+mid, left.values.end());
        left.keys.erase(left.keys.begin()+mid, left.keys.end());
        left.values.erase(left.values.begin()+mid, left.values.end());
        // link leaves
        right.next_leaf = left.next_leaf; left.next_leaf = (int)nodes.size();
        nodes.push_back(right);
        return (int)nodes.size()-1; // id of right
    }

    // Split internal node; return new right node id; set its first key
    int split_internal(int nid) {
        NciNode right; right.is_leaf = false;
        NciNode &left = nodes[nid];
        int total = (int)left.keys.size();
        int mid = total/2; // promote left.keys[mid]
        string promote = left.keys[mid];

        // right gets keys [mid+1..end) and children [mid+1..end]
        right.keys.assign(left.keys.begin()+mid+1, left.keys.end());
        right.children.assign(left.children.begin()+mid+1, left.children.end());

        // shrink left to [0..mid-1] and children [0..mid]
        left.keys.erase(left.keys.begin()+mid, left.keys.end());
        left.children.erase(left.children.begin()+mid+1, left.children.end());

        nodes.push_back(right);
        int rid = (int)nodes.size()-1;

        // The caller expects the *first key of right* to be used as the promote.
        // Make sure right has at least one key; if empty, reuse promote.
        if (nodes[rid].keys.empty()) nodes[rid].keys.push_back(promote);
        else nodes[rid].keys.front() = promote;
        return rid;
    }
};

// ------------------------------ Demo ----------------------------------------
int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    BaseTable table;
    // Insert base rows (clustered on id)
    vector<Row> rows = {
        {5, "Bob",    "Delhi"},
        {1, "Alice",  "Paris"},
        {3, "Kumar",  "Delhi"},
        {8, "David",  "London"},
        {2, "Charlie","Paris"},
        {7, "Riya",   "Mumbai"},
        {4, "Eve",    "Berlin"},
        {6, "Raj",    "Mumbai"}
    };
    for (auto &r: rows) table.insert(r);

    // Build Non‑Clustered Index on city (leaf stores (city -> list of ids))
    NciBPlusTree idx_city(2); // ORDER=2 => max_keys=4 per node
    for (auto &r: rows) idx_city.insert(r.city, r.id);

    table.print_all();
    idx_city.dump_ascii();
    idx_city.dump_dot("idx_city.dot");
    cout << "Graphviz DOT saved to idx_city.dot (render with: dot -Tpng idx_city.dot -o idx_city.png)\n\n";

    // ----------------- Query Flow Demo -----------------
    auto run_query = [&](const string& city){
        cout << "QUERY: SELECT * FROM Employees WHERE city='" << city << "'\n";
        // Step 1: probe non‑clustered index to get clustered keys (ids)
        auto ids = idx_city.search(city);
        cout << "  -> idx_city leaf gives ids: ";
        if (ids.empty()) cout << "{}\n"; else { cout << "{"; for (size_t i=0;i<ids.size();++i){ cout<<ids[i]<<(i+1<ids.size()?", ":""); } cout << "}" << "\n"; }
        // Step 2: use clustered keys to fetch full rows from base table
        cout << "  -> fetch rows by id from clustered index (table)\n";
        for (int id: ids) {
            auto r = table.get(id);
            if (r) cout << "     (id="<<r->id<<", name="<<r->name<<", city="<<r->city<<")\n";
        }
        cout << "\n";
    };

    run_query("Delhi");
    run_query("Mumbai");
    run_query("Tokyo");

    return 0;
}
