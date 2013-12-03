package gogit

import (
	"fmt"
)

func Run() {
  repo := Repo{"IoraHealth", "IoraHealth"}
  statuses, _ := repo.Open()

  fmt.Printf("size: %d\n", len(statuses))

  for _, status := range statuses {
	  fmt.Printf("[number: %d, comments: %d, status: %s, octocatted: %v]\n",
	             status.Number,
	             status.CommentCount,
	             status.Status,
	             status.Octocatted)
  }
}
