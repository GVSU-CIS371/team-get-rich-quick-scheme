import { render } from 'preact';
import { LocationProvider, Router, Route } from 'preact-iso';

import { HomePage } from './pages/Home';
import { NotFound } from './pages/_404.jsx';
import './style.css';
import {LoginPage} from "@/pages/Login";
import {RegisterPage} from "./pages/Register";
import {AuthProvider} from "@/components/auth/Auth";
import {DashboardHomePage} from "@/pages/Dashboard";
import {LogoutPage} from "@/pages/Logout";
import {AddOrganizationPage} from "@/pages/Dashboard/organizations/add";

export function App() {
	return (
		<AuthProvider>
			<div className="dark bg-background min-h-screen">
				<LocationProvider>
					<main>
						<Router>
							<Route path="/" component={HomePage} />
							<Route path="/login" component={LoginPage} />
							<Route path="/register" component={RegisterPage} />
							<Route path="/logout" component={LogoutPage} />
							<Route path="/dashboard" component={DashboardHomePage} />
							<Route path="/dashboard/organizations/add" component={AddOrganizationPage} />
							<Route default component={NotFound} />
						</Router>
					</main>
				</LocationProvider>
			</div>
		</AuthProvider>
	);
}

render(<App />, document.getElementById('app'));
