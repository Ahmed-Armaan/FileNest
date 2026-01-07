function Logo(props: { className: string }) {
  return (
    <>
      <div className={`${props.className}`} style={{ fontFamily: 'audiowide, sans-serif', color: 'var(--text-primary)' }}>
        FileNest
      </div >
    </>
  )
}

export default Logo
