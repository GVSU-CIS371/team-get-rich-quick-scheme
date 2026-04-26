import {useLocation} from "preact-iso";

export function Home() {
	const location = useLocation();
	location.route("/login");

	return (
		<div class="home">
			<p>Redirecting...</p>
		</div>
	);
}

