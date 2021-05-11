package services

import (
	"encoding/binary"

	"github.com/Damillora/Shioriko/pkg/database"
	"github.com/Damillora/Shioriko/pkg/models"
	"github.com/corona10/goimagehash"
)

func SimilaritySearch(originalHashInt uint64) ([]models.PostSimilarityListItem, error) {
	originalHash := goimagehash.NewImageHash(originalHashInt, goimagehash.PHash)

	hashSlice := make([]byte, 8)
	binary.LittleEndian.PutUint64(hashSlice, originalHashInt)

	var blobs []database.Blob
	database.DB.
		Joins("inner join posts on posts.blob_id = blobs.id").
		Where("hash1 = ?", hashSlice[0:2]).
		Or("hash2 = ?", hashSlice[2:4]).
		Or("hash3 = ?", hashSlice[4:6]).
		Or("hash4 = ?", hashSlice[6:8]).
		Find(&blobs)

	posts := make([]models.PostSimilarityListItem, 0)
	for _, blob := range blobs {
		hash2 := append(blob.Hash1, blob.Hash2...)
		hash3 := append(hash2, blob.Hash3...)
		hash4 := append(hash3, blob.Hash4...)

		hashInt := binary.LittleEndian.Uint64(hash4)

		compareHash := goimagehash.NewImageHash(hashInt, goimagehash.PHash)

		distance, _ := compareHash.Distance(originalHash)
		if distance < 1 {
			var post database.Post
			database.DB.Where("blob_id = ?", blob.ID).Find(&post)
			posts = append(posts, models.PostSimilarityListItem{
				ID:       post.ID,
				Distance: distance,
			})
		}
	}
	return posts, nil

}
