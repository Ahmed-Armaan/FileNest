function SideBar(props: { currDirId: string }) {

  return (
    <>
      <div className="h-full p-5" style={{ background: "var(--bg-secondary)" }}>
        <div>
          Upload File
        </div>
        <div>
          Create Folder
        </div>
      </div>
    </>
  )
}

export default SideBar
