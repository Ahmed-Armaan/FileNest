package database

import (
	"os"
	"testing"
)

// The following besnch mark is used to measure database query performance.
// Any change in the Db queries, if made for performance basis should be compared agains previous implementation
func BenchmarkUserBrowseFlow(b *testing.B) {
	googleId := os.Getenv("TEST_GOOGLE_ID")
	if googleId == "" {
		b.Skip("BENCH_GOOGLE_ID not set")
	}

	if err := DbInit(); err != nil {
		b.Fatal(err)
	}

	// Warm-up (not measured)
	_, err := GetRootNodeId(googleId)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for _ = range b.N {
		_, err := GetUserDataByGoogleId(
			googleId,
			UserDbColums.ID,
			UserDbColums.UserName,
			UserDbColums.ProfileImage,
		)
		if err != nil {
			b.Fatal(err)
		}

		root, err := GetRootNodeId(googleId)
		if err != nil {
			b.Fatal(err)
		}

		children, err := GetAllChild(&root.ID, googleId)
		if err != nil {
			b.Fatal(err)
		}

		if len(children) > 0 {
			_, err = GetAllChild(&children[0].ID, googleId)
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}
