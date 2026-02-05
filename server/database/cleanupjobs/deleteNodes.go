package cleanupjobs

import (
	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/storage"
)

// A bfs approach towards deletion. It lets us easily manage S3 deletion alongside database entry
// The load on database can also be managed, if needed, by changing the number of nodes selected in each batch
func deleteNodes(db database.DatabaseStore, s storage.StorageStore) (bool, error) {
	nodesQueue, err := db.ListDeletedNodes(50)
	if err != nil {
		return false, err
	}

	for i := range nodesQueue {
		currNode := nodesQueue[i]

		if database.NodeType(currNode.Type) == database.NodeTypeFile {
			if err := s.DeleteFileById(currNode.ID, db); err != nil {
				return false, err
			}
			if err := db.DeleteNodePermanently(currNode.ID); err != nil {
				return false, err
			}
			continue
		}

		children, err := db.ListChildrenForDeletion(&currNode.ID)
		if err != nil {
			return false, err
		}

		if len(children) == 0 {
			if err := db.DeleteNodePermanently(currNode.ID); err != nil {
				return false, err
			}
		} else {
			for _, child := range children {
				if child.Type == string(database.NodeTypeFile) {
					if err := s.DeleteFileById(child.ID, db); err != nil {
						return false, err
					}
					if err := db.DeleteNodePermanently(child.ID); err != nil {
						return false, err
					}
				}
			}
		}
	}

	if len(nodesQueue) < 50 {
		return true, nil
	}
	return false, nil
}
