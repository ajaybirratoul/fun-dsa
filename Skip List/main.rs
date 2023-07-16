use rand::Rng;

// Node struct representing a single node in the skip list
struct Node<T> {
    value: T,
    next: Vec<Option<Box<Node<T>>>>,
}

impl<T> Node<T> {
    fn new(value: T, height: usize) -> Node<T> {
        Node {
            value,
            next: vec![None; height],
        }
    }
}

// SkipList struct representing the skip list data structure
pub struct SkipList<T> {
    head: Node<T>,
    height: usize,
    len: usize,
}

impl<T: PartialOrd> SkipList<T> {
    pub fn new() -> SkipList<T> {
        let head = Node::new(Default::default(), 32);
        SkipList {
            head,
            height: 1,
            len: 0,
        }
    }

    pub fn is_empty(&self) -> bool {
        self.len == 0
    }

    pub fn len(&self) -> usize {
        self.len
    }

    pub fn insert(&mut self, value: T) {
        let mut rng = rand::thread_rng();
        let mut update = vec![None; self.height];
        let mut x = &mut self.head;

        for i in (0..self.height).rev() {
            while let Some(ref mut next) = x.next[i] {
                if next.value < value {
                    x = next;
                } else {
                    break;
                }
            }
            update[i] = Some(Box::new(x.clone()));
        }

        let height = self.random_height();
        if height > self.height {
            for i in self.height..height {
                update.push(Some(Box::new(self.head.clone())));
            }
            self.height = height;
        }

        let node = Node::new(value, height);
        for i in 0..height {
            if let Some(ref mut next) = update[i].as_mut().unwrap().next[i] {
                node.next[i] = next.next[i].take();
                next.next[i] = Some(Box::new(node.clone()));
            } else {
                node.next[i] = None;
                update[i].as_mut().unwrap().next[i] = Some(Box::new(node.clone()));
            }
        }

        self.len += 1;
    }

    pub fn contains(&self, value: T) -> bool {
        let mut x = &self.head;

        for i in (0..self.height).rev() {
            while let Some(ref next) = x.next[i] {
                if next.value <= value {
                    if next.value == value {
                        return true;
                    }
                    x = next;
                } else {
                    break;
                }
            }
        }

        false
    }

    fn random_height(&mut self) -> usize {
        let mut height = 1;
        while height < self.height && rand::random::<bool>() {
            height += 1;
        }
        height
    }
}
