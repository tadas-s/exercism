package dotdsl

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Properties holds the properties of a node or edge.
// The values can be int, bool or string.
type Properties map[string]any

// Graph stores the parts of a dot graph.
// All entities are stored as a Properties map (`nil` Properties when none set)
// attrs is the Properties for the entire Graph, vs a specific node or edge.
type Graph struct {
	nodes map[string]Properties
	edges map[string]Properties
	attrs Properties
}

// Parse creates a Graph from a text blob.
func Parse(data string) (*Graph, error) {
	tokens, err := tokenizer(data)

	if err != nil {
		return nil, err
	}

	graph, err := parser(tokens)

	if err != nil {
		return nil, err
	}

	return graph, nil
}

type LangTokenKind int

const (
	LTNone LangTokenKind = iota
	LTBracket
	LTSeparator
	LTAssignment
	LTEdge
	LTName
	LTNumber
	LTQuotedString
)

func (lt LangTokenKind) String() string {
	switch lt {
	case LTBracket:
		return "bracket"
	case LTSeparator:
		return "separator"
	case LTAssignment:
		return "assignment"
	case LTEdge:
		return "edge"
	case LTName:
		return "name"
	case LTNumber:
		return "number"
	case LTQuotedString:
		return "quoted string"
	default:
		log.Panicf("String representation not defined for token '%d'", lt)
	}

	return ""
}

type LangToken struct {
	kind  LangTokenKind
	value string
}

func (lt LangToken) Is(kind LangTokenKind, value string) bool {
	return lt.kind == kind && lt.value == value
}

type LangTokens []LangToken

// Get retrieves token at index or "blank" if n is out of bounds
func (lts LangTokens) Get(n int) LangToken {
	if n < len(lts) {
		return lts[n]
	}

	return LangToken{LTNone, ""}
}

// Kind retrieves token kind at index or LTNone if n is out of bounds
func (lts LangTokens) Kind(n int) LangTokenKind {
	return lts.Get(n).kind
}

// Value retrieves token value at index or "" if n is out of bounds
func (lts LangTokens) Value(n int) string {
	return lts.Get(n).value
}

func tokenizer(data string) (LangTokens, error) {
	runes := []rune(data)
	pos := 0
	var tokens LangTokens

	for pos < len(runes) {
		r := runes[pos]

		if r == '{' || r == '}' || r == '[' || r == ']' {
			tokens = append(tokens, LangToken{LTBracket, string(r)})
			pos++
		} else if r == ';' {
			tokens = append(tokens, LangToken{LTSeparator, string(r)})
			pos++
		} else if r == '=' {
			tokens = append(tokens, LangToken{LTAssignment, string(r)})
			pos++
		} else if r == '-' {
			pos++
			r = runes[pos]

			if r == '-' {
				tokens = append(tokens, LangToken{LTEdge, "--"})
				pos++
			} else {
				return tokens, fmt.Errorf("bad character: '%s' at position %d", string(r), pos)
			}
		} else if r == '/' {
			pos++
			r = runes[pos]

			if r == '/' {
				// skip comment until end of line
				for r != '\n' {
					pos++
					r = runes[pos]
				}
			} else {
				return tokens, fmt.Errorf("bad character: '%s' at position %d", string(r), pos)
			}
		} else if r == '#' {
			// skip comment until end of line
			for r != '\n' {
				pos++
				r = runes[pos]
			}
		} else if r >= 'a' && r <= 'z' {
			var value strings.Builder

			for r >= 'a' && r <= 'z' {
				value.WriteRune(r)
				pos++
				r = runes[pos]
			}

			tokens = append(tokens, LangToken{LTName, value.String()})
		} else if r >= '0' && r <= '9' {
			var value strings.Builder

			for r >= '0' && r <= '9' {
				value.WriteRune(r)
				pos++
				r = runes[pos]
			}

			tokens = append(tokens, LangToken{LTNumber, value.String()})
		} else if r == '"' {
			var value strings.Builder

			pos++
			r = runes[pos]

			for r != '"' {
				value.WriteRune(r)
				pos++
				r = runes[pos]
			}

			pos++

			tokens = append(tokens, LangToken{LTQuotedString, value.String()})
		} else if r == ' ' || r == '\n' {
			pos++
		} else {
			return tokens, fmt.Errorf("bad character: '%s' at position %d", string(r), pos)
		}
	}

	return tokens, nil
}

func parser(tokens LangTokens) (*Graph, error) {
	pos := 0

	graph := &Graph{
		nodes: map[string]Properties{},
		edges: map[string]Properties{},
		attrs: Properties{},
	}

	if tokens.Get(pos).Is(LTName, "graph") {
		pos++

		if tokens.Get(pos).Is(LTBracket, "{") {
			pos++

			for !tokens.Get(pos).Is(LTBracket, "}") {
				err := parseGraphDeclaration(tokens, &pos, graph)

				if err != nil {
					return nil, err
				}
			}
		} else {
			return nil, errors.New("invalid graph")
		}

		return graph, nil
	}

	return nil, errors.New("invalid graph")
}

func parseGraphDeclaration(tokens LangTokens, pos *int, graph *Graph) error {
	if tokens.Get(*pos).Is(LTBracket, "[") {
		attribute, value, err := parseGraphAttribute(tokens, pos)

		if err != nil {
			return err
		}

		graph.attrs[attribute] = value

		if tokens.Get(*pos).Is(LTSeparator, ";") {
			*pos++
			return nil
		}
		return fmt.Errorf("graph attribute: unexpected %s %s", tokens.Kind(*pos), tokens.Value(*pos))
	} else if tokens.Kind(*pos) == LTName {
		return parseGraphNode(tokens, pos, graph)
	}

	return fmt.Errorf("graph declaration item: unexpected %s %s", tokens.Kind(*pos), tokens.Value(*pos))
}

func parseGraphAttribute(tokens LangTokens, pos *int) (string, any, error) {
	var attribute string
	var value any

	if tokens.Get(*pos).Is(LTBracket, "[") &&
		tokens.Kind(*pos+1) == LTName &&
		tokens.Kind(*pos+2) == LTAssignment &&
		(tokens.Kind(*pos+3) == LTNumber || tokens.Kind(*pos+3) == LTName || tokens.Kind(*pos+3) == LTQuotedString) &&
		tokens.Get(*pos+4).Is(LTBracket, "]") {

		attribute = tokens.Value(*pos + 1)

		if tokens.Kind(*pos+3) == LTNumber {
			intValue, _ := strconv.Atoi(tokens.Value(*pos + 3))
			value = intValue
		} else if tokens.Value(*pos+3) == "true" {
			value = true
		} else if tokens.Value(*pos+3) == "false" {
			value = false
		} else {
			value = tokens.Value(*pos + 3)
		}

		*pos += 5
	} else {
		return "", "", fmt.Errorf("attribute parsing, unexpected %s %s", tokens.Kind(*pos), tokens.Value(*pos))
	}

	return attribute, value, nil
}

func parseGraphNode(tokens LangTokens, pos *int, graph *Graph) error {
	var nodes []string
	var edges [][]string
	var attribute string
	var value any

	for {
		if len(nodes) > 0 && tokens.Kind(*pos) == LTEdge && tokens.Kind(*pos+1) == LTName {
			nodes = append(nodes, tokens.Value(*pos+1))
			*pos += 2
		} else if len(nodes) == 0 && tokens.Kind(*pos) == LTName {
			nodes = append(nodes, tokens.Value(*pos))
			*pos++
		} else {
			break
		}
	}

	if tokens.Get(*pos).Is(LTBracket, "[") {
		var err error

		attribute, value, err = parseGraphAttribute(tokens, pos)

		if err != nil {
			return err
		}
	}

	if tokens.Get(*pos).Is(LTSeparator, ";") {
		*pos++
	} else {
		return fmt.Errorf("expected line end, found %s '%s'", tokens.Kind(*pos), tokens.Value(*pos))
	}

	if len(nodes) > 1 {
		for i := 0; i < len(nodes)-1; i++ {
			edges = append(edges, []string{nodes[i], nodes[i+1]})
		}
	}

	for _, nodeName := range nodes {
		if _, exists := graph.nodes[nodeName]; !exists {
			graph.nodes[nodeName] = Properties{}
		}
	}

	if len(nodes) == 1 && attribute != "" {
		graph.nodes[nodes[0]][attribute] = value
	}

	for _, edgePair := range edges {
		edge := fmt.Sprintf("{%s %s}", edgePair[0], edgePair[1])
		_, exists := graph.edges[edge]

		alternativeEdge := fmt.Sprintf("{%s %s}", edgePair[1], edgePair[0])
		_, alternativeExists := graph.edges[alternativeEdge]

		if !exists && !alternativeExists {
			graph.edges[edge] = Properties{}

			if attribute != "" {
				graph.edges[edge][attribute] = value
			}
		} else if !exists {
			if attribute != "" {
				graph.edges[alternativeEdge][attribute] = value
			}
		}
	}

	return nil
}
