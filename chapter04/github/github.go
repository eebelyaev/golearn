package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues запрашивает GitHub.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	s := IssuesURL + "?q=" + q
	resp, err := http.Get(s)
	if err != nil {
		return nil, err
	}
	// Необходимо закрыть resp.Body на всех путях выполнения.
	// (В главе Б вы познакомитесь с более простым решением: ’defer’.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("сбой запроса: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
