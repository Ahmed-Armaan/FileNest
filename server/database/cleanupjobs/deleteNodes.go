package cleanupjobs

import (
	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/storage"
)

// A bfs approach towards deletion. It lets us easily manage S3 deletion alongside database entry
// The load on database can also be managed, if needed, by changing the number of nodes selected in each batch
func deleteNodes() (bool, error) {
	nodesQueue, err := database.GetDeletedNodes(50)
	if err != nil {
		return false, err
	}

	for i := range len(nodesQueue) {
		currNode := nodesQueue[i]

		if database.NodeType(currNode.Type) == database.NodeTypeFile {
			if err := storage.DeleteFileFromS3(currNode.ID); err != nil {
				return false, err
			}
			database.HardDeletion(currNode.ID)
			continue
		}

		children, err := database.GetAllChildToDelete(&currNode.ID)
		if err != nil {
			return false, err
		}

		if len(children) == 0 {
			database.HardDeletion(currNode.ID)
		} else {
			for _, child := range children {
				if child.Type == string(database.NodeTypeFile) {
					if err := storage.DeleteFileFromS3(child.ID); err != nil {
						return false, err
					}
					database.HardDeletion(child.ID)
				}
			}
		}
	}

	if len(nodesQueue) < 50 {
		return true, nil
	}
	return false, nil
}
