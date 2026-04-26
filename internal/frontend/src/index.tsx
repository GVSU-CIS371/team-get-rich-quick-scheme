import { render } from 'preact';
import { LocationProvider, Router, Route } from 'preact-iso';

import { Home } from './pages/Home';
import { NotFound } from './pages/_404.jsx';
import './style.css';
import {Login} from "./pages/Login";

export function App() {
	return (
		<div className="dark bg-background min-h-screen">
			<LocationProvider>
				<main>
					<Router>
						<Route path="/" component={Home} />
						<Route path="/login" component={Login} />
						<Route default component={NotFound} />
					</Router>
				</main>
			</LocationProvider>
		</div>
	);
}

render(<App />, document.getElementById('app'));
