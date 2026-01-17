export type NodeType = "file" | "directory"

export class FileTreeNode {
	readonly nodeId: string
	readonly nodeName: string
	readonly nodeType: NodeType
	private _children: FileTreeNode[] = []

	constructor(nodeId: string, nodeName: string, nodeType: NodeType, children: FileTreeNode[]) {
		this.nodeId = nodeId
		this.nodeName = nodeName
		this.nodeType = nodeType
		this._children = children
	}

	set children(children: FileTreeNode[]) {
		if (this.nodeType === "directory") {
			this._children = children
		}
		else {
			throw new Error("Files cannot have children")
		}
	}

	get children() {
		if (this.nodeType === "directory") {
			return this._children
		}
		else {
			throw new Error("Files cannot have children")
		}
	}
}
