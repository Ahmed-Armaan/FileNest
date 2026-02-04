import type { LoaderFunctionArgs } from "react-router"

type HomePageType = "normal" | "share"

export type Loaderdata = {
	mode: HomePageType,
	code?: string,
}

export async function HomeLoader({ params }: LoaderFunctionArgs): Promise<Loaderdata> {
	if (params.share_code) {
		return {
			mode: "share",
			code: params.share_code,
		}
	}
	else {
		return {
			mode: "normal"
		}
	}
}
