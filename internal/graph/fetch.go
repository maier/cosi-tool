// Copyright © 2018 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package graph

import (
	"regexp"
	"strings"

	"github.com/circonus-labs/circonus-gometrics/api"
	"github.com/pkg/errors"
)

// Fetch retrieves a graph using the Circonus API
func Fetch(client API, id, title string) (*api.Graph, error) {
	// logger := log.With().Str("cmd", "cosi graph fetch").Logger()

	if client == nil {
		return nil, errors.New("invalid state, nil client")
	}

	if id != "" {
		g, err := FetchByID(client, id)
		if err != nil {
			return nil, errors.Wrap(err, "graph by id")
		}
		return g, nil
	} else if title != "" {
		g, err := FetchByTitle(client, title)
		if err != nil {
			return nil, errors.Wrap(err, "graph by title")
		}
		return g, nil
	}

	return nil, errors.Errorf("missing required argument identifying which graph to fetch")
}

// FetchByID retrieves a graph by CID using Circonus API
func FetchByID(client API, id string) (*api.Graph, error) {
	if client == nil {
		return nil, errors.New("invalid state, nil client")
	}
	if id == "" {
		return nil, errors.New("invalid id (empty)")
	}
	cid := id
	if !strings.HasPrefix(cid, "/graph/") {
		cid = "/graph/" + id
	}
	if ok, err := regexp.MatchString(`^/graph/[0-9]+`, cid); err != nil {
		return nil, errors.Wrap(err, "compile graph id regexp")
	} else if !ok {
		return nil, errors.Errorf("invalid graph id (%s)", id)
	}
	db, err := client.FetchGraph(api.CIDType(&cid))
	if err != nil {
		return nil, errors.Wrap(err, "fetch api")
	}
	return db, nil
}

// FetchByTitle retrieves a graph by Display Title using Circonus API
func FetchByTitle(client API, name string) (*api.Graph, error) {
	if client == nil {
		return nil, errors.New("invalid state, nil client")
	}
	if name == "" {
		return nil, errors.Errorf("invalid title (empty)")
	}

	query := api.SearchQueryType("\"" + name + "\" (active:1)")

	dbs, err := client.SearchGraphs(&query, nil)
	if err != nil {
		return nil, errors.Wrap(err, "search api")
	}

	if dbs == nil {
		return nil, errors.Errorf("no graph or error returned")
	}

	if len(*dbs) == 0 {
		return nil, errors.Errorf("no graphs found matching (%s)", name)
	}

	if len(*dbs) > 1 {
		return nil, errors.Errorf("multiple graphs matching (%s)", name)
	}

	db := (*dbs)[0]
	return &db, nil
}
