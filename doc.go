// The sgf library contains various tools for working with a SGF files and the
// SGF trees within them. Trees are simply a collection of nodes connected
// together via parent and child relationships. Each node can have properties,
// which are keys with a slice of values. All keys and values are stored as
// strings.
//
// In general, functions or methods which require a coordinate expect that
// coordinate to be supplied as an SGF string. For example, the string "dd" is
// the top left hoshi point. Such strings can be generated by the Point() utility
// function; e.g. Point(3, 3) returns "dd".
//
// Nodes can be used to generate boards via node.Board(), but editing a board has
// no effect on the node that created it. Creating boards is relatively
// expensive, but often unnecessary for simple tasks.
package sgf
