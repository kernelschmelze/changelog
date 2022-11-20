package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Commit struct {
	Time         time.Time
	ID           string
	Author       string
	Message      string
	ChangedFiles []string
}

func checkIfError(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "%s\n", e)
		os.Exit(1)
	}
}

func main() {

	path := "."

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	r, err := git.PlainOpen(path)
	checkIfError(err)

	ref, err := r.Head()
	checkIfError(err)

	history, err := r.Log(&git.LogOptions{From: ref.Hash()})
	checkIfError(err)

	commits := []Commit{}

	err = history.ForEach(func(c *object.Commit) error {

		commit := Commit{
			ID:      c.ID().String(),
			Author:  c.Author.Name,
			Time:    c.Author.When,
			Message: strings.TrimRight(c.Message, "\n"),
		}

		if len(c.ParentHashes) > 0 {

			if parent, err := r.CommitObject(c.ParentHashes[0]); err == nil {

				if patch, err := parent.Patch(c); err == nil {

					for _, p := range patch.FilePatches() {

						if from, to := p.Files(); from != nil {
							commit.ChangedFiles = append(commit.ChangedFiles, from.Path())
						} else if to != from {
							commit.ChangedFiles = append(commit.ChangedFiles, to.Path())
						}

					}

				}

			}

		}

		commits = append(commits, commit)

		return nil
	})

	checkIfError(err)

	sort.Slice(commits, func(i, j int) bool {
		return commits[i].Time.Unix() > commits[j].Time.Unix()
	})

	if out, err := json.MarshalIndent(commits, "", "  "); err == nil {
		fmt.Println(string(out))
	}

}
