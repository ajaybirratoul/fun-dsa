#include <iostream>
#include <vector>

struct Point {
    int x, y;

    Point(int x, int y) : x(x), y(y) {}
};

class Quadtree {
    static const int CAPACITY = 4;

    struct Node {
        bool isLeaf;
        std::vector<Point> points;
        Node* children[4];

        Node() : isLeaf(true), points(), children{} {}
    };

    Node* root;

    void subdivide(Node* node) {
        int x = node->points[0].x;
        int y = node->points[0].y;
        int halfWidth = (node->points[1].x - x) / 2;
        int halfHeight = (node->points[2].y - y) / 2;

        node->children[0] = new Node();
        node->children[0]->points.push_back(Point(x + halfWidth, y));
        node->children[0]->points.push_back(Point(x + halfWidth * 2, y + halfHeight));
        node->children[0]->points.push_back(Point(x + halfWidth, y + halfHeight * 2));

        node->children[1] = new Node();
        node->children[1]->points.push_back(Point(x + halfWidth, y + halfHeight));
        node->children[1]->points.push_back(Point(x + halfWidth * 2, y + halfHeight));
        node->children[1]->points.push_back(Point(x + halfWidth * 2, y + halfHeight * 2));

        node->children[2] = new Node();
        node->children[2]->points.push_back(Point(x, y + halfHeight));
        node->children[2]->points.push_back(Point(x + halfWidth, y + halfHeight));
        node->children[2]->points.push_back(Point(x, y + halfHeight * 2));

        node->children[3] = new Node();
        node->children[3]->points.push_back(Point(x, y));
        node->children[3]->points.push_back(Point(x + halfWidth, y));
        node->children[3]->points.push_back(Point(x, y + halfHeight));
    }

    void insert(Node* node, const Point& point) {
        if (node->isLeaf) {
            node->points.push_back(point);
            if (node->points.size() > CAPACITY) {
                subdivide(node);
                node->isLeaf = false;

                // Reinsert points to children
                for (const auto& p : node->points) {
                    for (int i = 0; i < 4; ++i) {
                        if (contains(node->children[i], p)) {
                            insert(node->children[i], p);
                            break;
                        }
                    }
                }
                node->points.clear();
            }
        } else {
            for (int i = 0; i < 4; ++i) {
                if (contains(node->children[i], point)) {
                    insert(node->children[i], point);
                    break;
                }
            }
        }
    }

    bool contains(Node* node, const Point& point) {
        return (point.x >= node->points[0].x && point.x <= node->points[1].x &&
                point.y >= node->points[0].y && point.y <= node->points[2].y);
    }

public:
    Quadtree(int width, int height) {
        root = new Node();
        root->points.push_back(Point(0, 0));
        root->points.push_back(Point(width, 0));
        root->points.push_back(Point(0, height));
        root->points.push_back(Point(width, height));
    }

    void insert(const Point& point) {
        insert(root, point);
    }
};

int main() {
    Quadtree quadtree(800, 600);

    quadtree.insert(Point(100, 100));
    quadtree.insert(Point(200, 200));
    quadtree.insert(Point(300, 300));
    quadtree.insert(Point(400, 400));

    return 0;
}
