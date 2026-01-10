function SideBar(props: { currDirId: string }) {
  console.log(props.currDirId)

  const uploadFile = async (file: File) => {
    const res = await fetch(
      `${import.meta.env.VITE_BACKEND_URL}/api/upload`,
      { credentials: "include" }
    )

    if (!res.ok) return alert("Failed to get upload URL")

    const { uploadUrl } = await res.json()

    await fetch(uploadUrl, {
      method: "PUT",
      body: file,
    })
  }

  return (
    <div
      className="h-full p-5"
      style={{ background: "var(--bg-secondary)" }}
    >
      <div>
        <input
          type="file"
          onChange={(e) => {
            if (e.target.files?.[0]) {
              uploadFile(e.target.files[0])
            }
          }}
        />
      </div>

      <div className="mt-3">
        <button>Create Folder</button>
      </div>
    </div>
  )
}

export default SideBar
