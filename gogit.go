package gogit

import (
	"fmt"
)

func Run() {
  snowflake := Repo{"IoraHealth", "IoraHealth"}
  statuses, _ := Open(&snowflake)

  fmt.Printf("size: %d\n", len(statuses))

  for _, status := range statuses {
	  fmt.Printf("[number: %d, comments: %d, status: %s, octocatted: %v]\n",
	             status.Number,
	             status.CommentCount,
	             status.Status,
	             status.Octocatted)
  }
}
