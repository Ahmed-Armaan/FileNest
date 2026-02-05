import { useState } from "react"

export type PasswordPromptProps = {
	code: string
	setPassword: (password: string) => void
}

export default function PasswordPromt(props: PasswordPromptProps) {
	const [password, setLocalPassword] = useState("")

	return (
		<div
			className="fixed inset-0 flex items-center justify-center z-[9999]"
			style={{ background: "rgba(0,0,0,0.6)" }} // backdrop
		>
			<div
				className="w-full max-w-sm rounded-xl border p-6 shadow-xl"
				style={{
					background: "var(--bg-secondary)",
					borderColor: "var(--border-primary)",
				}}
			>
				<h2
					className="mb-2 text-lg font-medium"
					style={{ color: "var(--text-primary)" }}
				>
					Protected Share
				</h2>

				<p
					className="mb-4 text-sm"
					style={{ color: "var(--text-secondary)" }}
				>
					This shared item is password protected.
				</p>

				<input
					type="password"
					placeholder="Enter password"
					value={password}
					onChange={(e) => setLocalPassword(e.target.value)}
					className="w-full rounded-lg border px-3 py-2 text-sm focus:outline-none"
					style={{
						background: "var(--bg-tertiary)",
						borderColor: "var(--border-primary)",
						color: "var(--text-primary)",
					}}
				/>

				<button
					onClick={() => props.setPassword(password)}
					disabled={!password}
					className="mt-4 w-full rounded-lg px-3 py-2 text-sm font-medium disabled:opacity-50"
					style={{
						background: "var(--accent-primary)",
						color: "var(--text-primary)",
					}}
				>
					Unlock
				</button>
			</div>
		</div>
	)
}
