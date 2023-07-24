use std::collections::{HashMap, VecDeque};

#[derive(Default)]
struct TrieNode {
    children: HashMap<char, usize>,
    is_word: bool,
    fail: usize,
    output: Vec<String>,
}

struct AhoCorasick {
    trie: Vec<TrieNode>,
}

impl AhoCorasick {
    fn new() -> Self {
        AhoCorasick { trie: vec![TrieNode::default()] }
    }

    fn insert(&mut self, word: &str) {
        let mut node_idx = 0; // Start from the root node
        for &ch in word.chars() {
            let child = self.trie[node_idx].children.entry(ch).or_insert(self.trie.len());
            if *child == self.trie.len() {
                self.trie.push(TrieNode::default());
            }
            node_idx = *child;
        }
        self.trie[node_idx].is_word = true;
    }

    fn build_failure_links(&mut self) {
        let mut queue = VecDeque::new();

        // Initialize the root's children failure links
        for &(_, child_idx) in &self.trie[0].children {
            self.trie[child_idx].fail = 0;
            queue.push_back(child_idx);
        }

        while let Some(node_idx) = queue.pop_front() {
            let node = &mut self.trie[node_idx];

            for (&ch, &child_idx) in &node.children {
                queue.push_back(child_idx);

                let mut failure_link = node.fail;
                while failure_link != 0 && !self.trie[failure_link].children.contains_key(&ch) {
                    failure_link = self.trie[failure_link].fail;
                }
                if let Some(&target_idx) = self.trie[failure_link].children.get(&ch) {
                    self.trie[child_idx].fail = target_idx;
                } else {
                    self.trie[child_idx].fail = 0;
                }
            }

            if self.trie[node_idx].is_word {
                let mut output_link = node.fail;
                while output_link != 0 {
                    self.trie[node_idx].output.extend(&self.trie[output_link].output);
                    output_link = self.trie[output_link].fail;
                }
            }
        }
    }

    fn search(&self, text: &str) -> Vec<String> {
        let mut result = Vec::new();
        let mut node_idx = 0;

        for (i, ch) in text.chars().enumerate() {
            while node_idx != 0 && !self.trie[node_idx].children.contains_key(&ch) {
                node_idx = self.trie[node_idx].fail;
            }

            if let Some(&target_idx) = self.trie[node_idx].children.get(&ch) {
                node_idx = target_idx;
                if self.trie[node_idx].is_word {
                    result.extend(&self.trie[node_idx].output);
                }
            }
        }
        result
    }
}

fn main() {
    let mut ac = AhoCorasick::new();

    // Add keywords to search for
    ac.insert("he");
    ac.insert("she");
    ac.insert("his");
    ac.insert("hers");

    // Build failure links
    ac.build_failure_links();

    let text = "ushers";

    // Search for keywords in the text
    let found_words = ac.search(text);
    println!("Found words: {:?}", found_words);
}
