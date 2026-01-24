export type NodeType = "file" | "directory"
export const NodeTypeFile = "file"
export const NodeTypeDirectory = "directory"

export class FileTreeNode {
	readonly nodeId: string
	readonly nodeName: string
	readonly nodeType: NodeType
	readonly updatedAt: string
	private readonly _parent: FileTreeNode | null
	private _children: FileTreeNode[] = []

	constructor(nodeId: string, nodeName: string, nodeType: NodeType, updatedAt: string, children: FileTreeNode[], parent: FileTreeNode | null) {
		this.nodeId = nodeId
		this.nodeName = nodeName
		this.nodeType = nodeType
		this.updatedAt = updatedAt
		this._children = children

		if (nodeName === "/") this._parent = null
		else this._parent = parent
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

	get parent() {
		return this._parent
	}
}
